package iec104

import (
	"encoding/json"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	log "github.com/hashicorp/go-hclog"
	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/thinkgos/go-iecp5/cs104"
)

type Worker struct {
	ctx       dm.Context
	driver    *Driver
	job       Job
	slave     *Slave
	log       log.Logger
	isStarted bool
}

func NewWorker(d *Driver, ctx dm.Context, job Job, slave *Slave, logger log.Logger) *Worker {
	w := &Worker{
		driver: d,
		slave:  slave,
		job:    job,
		ctx:    ctx,
		log:    logger,
	}
	w.slave.handler.Listener = w
	// 必须！激活iec104。
	slave.client.SendStartDt()
	slave.client.SetOnDtActiveConfirm(func(c *cs104.Client) {
		w.OnStartConfirm()
	})
	return w
}

// Execute interval时间间隔周期执行测点采集上报任务。采集测点信息通过总召实现。
func (w *Worker) Execute() error {
	// 激活iec104后，才可以发送总召唤。isStarted为true时，iec104激活。
	if w.isStarted {
		// 总召唤命令。返回通过Listener接口方法MeasuredValueFloat（遥测值）、SinglePoint（遥信值）。
		return w.slave.client.InterrogationCmd(asdu.CauseOfTransmission{
			Cause: asdu.Activation,
		}, 0x1, asdu.QOIStation)
	}
	return nil
}

// SetPointCmdFloat 遥调设点,预设。
func (w *Worker) SetPointCmdFloat(point Point, val float32) error {
	return asdu.SetpointCmdFloat(w.slave.client, asdu.C_SE_NC_1, asdu.CauseOfTransmission{
		Cause: asdu.Activation,
	}, 0x1, asdu.SetpointCommandFloatInfo{
		Ioa:   asdu.InfoObjAddr(point.PointNum),
		Value: val,
		Qos: asdu.QualifierOfSetpointCmd{
			InSelect: true,
		},
	})
}

// SetPointCmdFloatExecute 遥调设点,执行。
func (w *Worker) SetPointCmdFloatExecute(info asdu.SetpointCommandFloatInfo, du *asdu.ASDU) {
	err := asdu.SetpointCmdFloat(w.slave.client, asdu.C_SE_NC_1, asdu.CauseOfTransmission{
		Cause: asdu.Activation,
	}, du.CommonAddr, asdu.SetpointCommandFloatInfo{
		Ioa:   info.Ioa,
		Value: info.Value,
		Qos: asdu.QualifierOfSetpointCmd{
			InSelect: false,
		},
	})
	if err != nil {
		w.log.Error("failed to set point", "error", err)
	}
}

func (w *Worker) SetPointCmdFloatPreseted(info asdu.SetpointCommandFloatInfo, du *asdu.ASDU) {
	w.SetPointCmdFloatExecute(info, du)
}

// SingleCmdExecute 单点遥控,执行。
func (w *Worker) SingleCmdExecute(info asdu.SingleCommandInfo, du *asdu.ASDU) {
	err := asdu.SingleCmd(w.slave.client, asdu.C_SC_NA_1, asdu.CauseOfTransmission{
		Cause: asdu.Activation,
	}, du.CommonAddr, asdu.SingleCommandInfo{
		Ioa:   info.Ioa,
		Value: info.Value,
		Qoc: asdu.QualifierOfCommand{
			InSelect: false,
		},
	})
	if err != nil {
		w.log.Error("failed to set point", "error", err)
	}
}

func (w *Worker) SingleCmdPreseted(info asdu.SingleCommandInfo, du *asdu.ASDU) {
	w.SingleCmdExecute(info, du)
}

// SingleCmd 单点遥控
func (w *Worker) SingleCmd(point Point, val bool) error {
	return asdu.SingleCmd(w.slave.client, asdu.C_SC_NA_1, asdu.CauseOfTransmission{
		Cause: asdu.Activation,
	}, 0x1, asdu.SingleCommandInfo{
		Ioa:   asdu.InfoObjAddr(point.PointNum),
		Value: val,
		Qoc: asdu.QualifierOfCommand{
			InSelect: true,
		},
	})
}

func (w *Worker) OnStartConfirm() {
	w.isStarted = true
}

func (w *Worker) MeasuredValueFloat(floatInfos []asdu.MeasuredValueFloatInfo) {
	// 获取从站返回的采集点信息，全遥测报文（总召）。
	points := make(map[string]interface{})
	for _, info := range floatInfos {
		pointName := getPointName(w.job.Points, uint(info.Ioa))
		if pointName == "" {
			continue
		}
		points[pointName] = info.Value
	}
	w.report(points)
}

func (w *Worker) SinglePoint(singlePointInfos []asdu.SinglePointInfo) {
	// 获取从站返回的采集点信息，全遥信报文（总召）。
	points := make(map[string]interface{})
	for _, info := range singlePointInfos {
		pointName := getPointName(w.job.Points, uint(info.Ioa))
		if pointName == "" {
			continue
		}
		points[pointName] = info.Value
	}
	w.report(points)
}

func (w *Worker) ReportInfo(infoObjAddr asdu.InfoObjAddr, value interface{}) {
	points := make(map[string]interface{})
	pointName := getPointName(w.job.Points, uint(infoObjAddr))
	w.log.Info("server execution confirmation", "Message type", pointName)
	if pointName != "" {
		points[pointName] = value
		w.report(points)
		w.log.Info("report measuring point", "name", pointName, "value", value)
	}
}

func (w *Worker) report(points map[string]interface{}) {
	msg := v1.Message{}

	msg.Kind = v1.MessageDeviceReport
	msg.Metadata = make(map[string]string)
	msg.Metadata[dm.KeyDriverName] = w.driver.driverName
	msg.Metadata[dm.KeyDeviceName] = w.slave.info.Name
	msg.Content = v1.LazyValue{Value: points}

	msgData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	res, err := w.driver.report.Post(&plugin.Request{Req: string(msgData)})
	if err != nil {
		return
	}
	w.log.Debug("iec104 driver report message", "msgdata", string(msgData), "rsp", res.Data)

	if err = w.slave.UpdateStatus(SlaveOnline); err != nil {
		w.log.Error("failed to update slave status", "error", err, "status", "online")
	}
}

func getPointName(points []Point, ioa uint) string {
	for _, point := range points {
		if point.PointNum == ioa {
			return point.Name
		}
	}
	return ""
}

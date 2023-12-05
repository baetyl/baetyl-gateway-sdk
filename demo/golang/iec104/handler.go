package iec104

import (
	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/thinkgos/go-iecp5/cs104"
)

type Handler struct {
	Listener Listener
}

type Listener interface {
	OnStartConfirm()                                                            // iec104激活确认
	MeasuredValueFloat(floatInfos []asdu.MeasuredValueFloatInfo)                // 全遥测报文，测量值 float32
	SinglePoint(singlePointInfos []asdu.SinglePointInfo)                        // 全遥信报文，单点遥信 true or false
	SetPointCmdFloatPreseted(info asdu.SetpointCommandFloatInfo, du *asdu.ASDU) // 遥调设点预设命令回调
	SingleCmdPreseted(info asdu.SingleCommandInfo, du *asdu.ASDU)               // 单点遥控设点命令回调
	ReportInfo(infoObjAddr asdu.InfoObjAddr, value interface{})                 // 上报AO和DO的值
}

func (c Handler) StartDtConfirm(client *cs104.Client) {

}

func (c Handler) InterrogationHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}

func (c Handler) CounterInterrogationHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}
func (c Handler) ReadHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}

func (c Handler) TestCommandHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}

func (c Handler) ClockSyncHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}
func (c Handler) ResetProcessHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}
func (c Handler) DelayAcquisitionHandler(con asdu.Connect, asdu *asdu.ASDU) error {
	return nil
}
func (c Handler) ASDUHandler(con asdu.Connect, du *asdu.ASDU) error {
	if du.Identifier.Type == asdu.M_ME_NC_1 {
		info := du.GetMeasuredValueFloat()
		if c.Listener != nil {
			c.Listener.MeasuredValueFloat(info)
		}
	} else if du.Identifier.Type == asdu.M_SP_NA_1 {
		info := du.GetSinglePoint()
		if c.Listener != nil {
			c.Listener.SinglePoint(info)
		}
	} else if du.Identifier.Type == asdu.C_SE_NC_1 && du.Coa.Cause == asdu.ActivationCon {
		cmd := du.GetSetpointFloatCmd()
		if cmd.Qos.InSelect {
			if c.Listener != nil {
				c.Listener.SetPointCmdFloatPreseted(cmd, du)
			}
		} else {
			if c.Listener != nil {
				c.Listener.ReportInfo(cmd.Ioa, cmd.Value)
			}
		}
	} else if du.Identifier.Type == asdu.C_SC_NA_1 {
		cmd := du.GetSingleCmd()
		if cmd.Qoc.InSelect {
			if c.Listener != nil {
				c.Listener.SingleCmdPreseted(cmd, du)
			}
		} else {
			if c.Listener != nil {
				c.Listener.ReportInfo(cmd.Ioa, cmd.Value)
			}
		}
	}
	return nil
}

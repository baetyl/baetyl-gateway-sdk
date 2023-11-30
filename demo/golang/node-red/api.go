package httpin

import (
	"encoding/json"

	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
)

type API struct {
	driverName string
	report     plugin.Report
}

func NewAPI(driverName string, report plugin.Report) *API {
	a := &API{
		driverName: driverName,
		report:     report,
	}
	return a
}

func (a *API) Report(c *gin.Context) (any, error) {
	L().Debug("api report start")
	deviceName := c.Param("name")

	props := make(map[string]any)
	err := c.ShouldBindBodyWith(&props, binding.JSON)
	if err != nil {
		return nil, E(ErrInvalid, F("error", err.Error()))
	}

	msg := v1.Message{
		Kind: v1.MessageDeviceReport,
		Metadata: map[string]string{
			dm.KeyDriverName: a.driverName,
			dm.KeyDeviceName: deviceName,
			"clientIP":       c.ClientIP(),
		},
		Content: v1.LazyValue{Value: props},
	}

	dt, err := json.Marshal(msg)
	if err != nil {
		return nil, E(ErrRunning, F("error", err.Error()))
	}

	res, err := a.report.Post(&plugin.Request{Req: string(dt)})
	if err != nil {
		return nil, E(ErrRunning, F("error", err.Error()))
	}
	L().Debug("custom driver report message", string(dt), res.Data)

	return res, nil
}

func (a *API) State(c *gin.Context) (any, error) {
	L().Debug("api status start")
	deviceName := c.Param("name")

	props := make(map[string]bool)
	err := c.ShouldBindBodyWith(&props, binding.JSON)
	if err != nil {
		return nil, E(ErrInvalid, F("error", err.Error()))
	}
	state, ok := props["state"]
	if !ok {
		return nil, E(ErrInvalid, F("error", "status not found"))
	}

	msg := v1.Message{
		Kind: v1.MessageDeviceLifecycleReport,
		Metadata: map[string]string{
			dm.KeyDriverName: a.driverName,
			dm.KeyDeviceName: deviceName,
			"clientIP":       c.ClientIP(),
		},
		Content: v1.LazyValue{Value: state},
	}

	dt, err := json.Marshal(msg)
	if err != nil {
		return nil, E(ErrRunning, F("error", err.Error()))
	}

	res, err := a.report.State(&plugin.Request{Req: string(dt)})
	if err != nil {
		return nil, E(ErrRunning, F("error", err.Error()))
	}
	L().Debug("custom driver report message", string(dt), res.Data)

	return res, nil
}

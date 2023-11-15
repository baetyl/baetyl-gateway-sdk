package httpin

import (
	"encoding/json"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	L().Debug("working start")

	props := make(map[string]any)
	err := c.ShouldBindBodyWith(&props, binding.JSON)
	if err != nil {
		return nil, E(ErrInvalid, F("error", err.Error()))
	}

	msg := v1.Message{
		Kind: v1.MessageDeviceReport,
		Metadata: map[string]string{
			dm.KeyDriverName: a.driverName,
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

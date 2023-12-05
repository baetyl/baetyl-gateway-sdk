package main

import (
	_ "net/http/pprof"
	"os"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/demo/modbus"
)

func main() {
	if err := plugin.Serve(&plugin.ServeOpts{
		FactoryFunc: modbus.NewDriver,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})
		logger.Error("plugin modbus shutting down", "error", err)
		os.Exit(1)
	}
}

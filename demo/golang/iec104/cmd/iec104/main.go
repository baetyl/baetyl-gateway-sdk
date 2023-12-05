package main

import (
	"os"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/demo/iec104"
)

func main() {
	if err := plugin.Serve(&plugin.ServeOpts{
		FactoryFunc: iec104.NewDriver,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})
		logger.Error("plugin iec104 shutting down", "error", err)
		os.Exit(1)
	}
}

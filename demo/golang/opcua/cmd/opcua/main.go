package main

import (
	"os"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/demo/opcua"
)

func main() {
	if err := plugin.Serve(&plugin.ServeOpts{
		FactoryFunc: opcua.NewDriver,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})
		logger.Error("plugin opcua shutting down", "error", err)
		os.Exit(1)
	}
}

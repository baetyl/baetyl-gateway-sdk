package main

import (
	"os"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/demo/bacnet"
)

func main() {
	if err := plugin.Serve(&plugin.ServeOpts{
		FactoryFunc: bacnet.NewDriver,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})
		logger.Error("plugin bacnet shutting down", "error", err)
		os.Exit(1)
	}
}

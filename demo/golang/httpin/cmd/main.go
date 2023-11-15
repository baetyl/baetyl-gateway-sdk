package main

import (
	"os"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/demo/httpin"
)

//go:generate go build -o ../../../../test/driver/http-in/http-in .

func main() {
	if err := plugin.Serve(&plugin.ServeOpts{
		FactoryFunc: httpin.NewDriver,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})
		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}

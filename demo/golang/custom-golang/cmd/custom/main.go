package main

import (
	"os"

	plugin "github.com/baetyl/baetyl-gateway-sdk/sdk/golang"
	"github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/demo/golang"
)

//go:generate go build -o ../../../../../test/driver/custom-golang/custom-golang .

func main() {
	if err := plugin.Serve(&plugin.ServeOpts{
		FactoryFunc: custom.NewDriver,
	}); err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})
		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}

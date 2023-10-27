package plugin

import (
	"crypto/tls"
	"fmt"
	"math"
	"os"
	"time"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var parentPid int

func init() {
	parentPid = os.Getppid()
}

type TLSProviderFunc func() (*tls.Config, error)

type ServeOpts struct {
	FactoryFunc     Factory
	TLSProviderFunc TLSProviderFunc
	Logger          log.Logger
}

func checkParentAlive() {
	go func() {
		for {
			if os.Getppid() != parentPid {
				fmt.Println("parent no alive, exit")
				os.Exit(0)
			}
			_, err := os.FindProcess(parentPid)
			if err != nil {
				fmt.Println("parent no alive, exit")
				os.Exit(0)
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func Serve(opts *ServeOpts) error {
	checkParentAlive()
	logger := opts.Logger
	if logger == nil {
		logger = log.New(&log.LoggerOptions{
			Level:      log.Debug,
			Output:     os.Stderr,
			JSONFormat: true,
		})
	}

	pluginSets := map[int]plugin.PluginSet{
		1: {
			PluginName: &DriverGRPCPlugin{
				Factory: opts.FactoryFunc,
				Log:     logger,
			},
		},
	}
	serveOpts := &plugin.ServeConfig{
		HandshakeConfig:  Handshake,
		VersionedPlugins: pluginSets,
		TLSProvider:      opts.TLSProviderFunc,
		Logger:           logger,
		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
			opts = append(opts, grpc.MaxRecvMsgSize(math.MaxInt32))
			opts = append(opts, grpc.MaxSendMsgSize(math.MaxInt32))
			return plugin.DefaultGRPCServer(opts)
		},
	}

	plugin.Serve(serveOpts)

	return nil
}

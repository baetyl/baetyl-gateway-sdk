package plugin

import (
	"context"
	"net"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang/proto"
)

var _ Driver = &gRPCClient{}

type gRPCClient struct {
	client proto.DriverClient
	broker *plugin.GRPCBroker
	logger log.Logger
}

func (c *gRPCClient) GetDriverInfo(req *Request) (*Response, error) {
	res, err := c.client.GetDriverInfo(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (c *gRPCClient) SetConfig(req *Request) (*Response, error) {
	res, err := c.client.SetConfig(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		Data: res.Data,
	}, nil
}

func (c *gRPCClient) Setup(config *BackendConfig) (*Response, error) {
	reportImpl := config.ReportSvc
	report := &gRPCReportServer{
		Impl: reportImpl,
	}
	// 如果这里使用基于 go-plugin 框架中的 broker.proto 生成并构建的 broker 来做通信
	// 则需要在客户端从 starStream 流中读取并记录当前链接的配置信息
	// 配置信息包括：brokerID address 等
	// 其中 address 为 unix/windows 系统的 socket 文件
	// 需要客户端基于此构建出连接服务端的 client
	// 但是在 java 中，直接基于 socket 文件打开链接需要考虑平台，并容易出现未知错误
	// 所以这里不使用 broker 来做配置通信，转为直接启动 report server
	// 并将端口通过 setup 函数的 brokerID 字段传给驱动插件
	//
	// //Register the server in this closure.
	// serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
	// 	 s := grpc.NewServer(opts...)
	//	 proto.RegisterReportServer(s, report)
	//	 return s
	// }
	// brokerID := c.broker.NextId()
	// go c.broker.AcceptAndServe(brokerID, serverFunc)

	// 直接通过 grpc 启动 report server
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}
	brokerID := lis.Addr().(*net.TCPAddr).Port
	grpcServer := grpc.NewServer()
	proto.RegisterReportServer(grpcServer, report)
	go func() {
		er := grpcServer.Serve(lis)
		if er != nil {
			c.logger.Error("failed to start grpc report server", er)
		}
	}()

	res, err := c.client.Setup(context.Background(), &proto.RequestArgs{
		Request:  config.DriverName,
		Brokerid: uint32(brokerID),
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Data: res.Data,
	}, nil
}

func (c *gRPCClient) Start(req *Request) (*Response, error) {
	_, err := c.client.Start(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *gRPCClient) Restart(req *Request) (*Response, error) {
	_, err := c.client.Restart(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *gRPCClient) Stop(req *Request) (*Response, error) {
	_, err := c.client.Stop(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *gRPCClient) Get(req *Request) (*Response, error) {
	_, err := c.client.Get(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *gRPCClient) Set(req *Request) (*Response, error) {
	_, err := c.client.Set(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

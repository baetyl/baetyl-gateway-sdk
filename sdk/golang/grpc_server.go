package golang

import (
	"context"
	"fmt"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang/proto"
)

type gRPCServer struct {
	// This is the real implementation
	factory Factory
	driver  Driver

	log    log.Logger
	broker *plugin.GRPCBroker
}

func (s *gRPCServer) GetDriverInfo(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	res, err := s.driver.GetDriverInfo(&Request{
		Req: req.Request,
	})
	if err != nil {
		return nil, err
	}

	return &proto.ResponseResult{
		Data: res.Data,
	}, nil
}

func (s *gRPCServer) SetConfig(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	res, err := s.driver.SetConfig(&Request{
		Req: req.Request,
	})
	if err != nil {
		return nil, err
	}

	return &proto.ResponseResult{Data: res.Data}, nil
}

func (s *gRPCServer) Setup(ctx context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	// 详见 grpc_client.go 文件关于不用 go-plugin 框架 broker server 说明
	// conn, err := s.broker.Dial(req.Brokerid)

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("0.0.0.0:%d", req.Brokerid),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &proto.ResponseResult{}, err
	}
	report := &gRPCReportClient{client: proto.NewReportClient(conn)}

	config := &BackendConfig{
		ReportSvc:  report,
		DriverName: req.Request,
		Log:        s.log,
	}
	driver, err := s.factory(ctx, config)
	if err != nil {
		return &proto.ResponseResult{}, err
	}
	s.driver = driver

	res, err := driver.Setup(config)
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{Data: res.Data}, nil
}

func (s *gRPCServer) Start(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := s.driver.Start(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (s *gRPCServer) Restart(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := s.driver.Restart(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (s *gRPCServer) Stop(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := s.driver.Stop(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (s *gRPCServer) Get(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := s.driver.Get(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

func (s *gRPCServer) Set(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	_, err := s.driver.Set(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}

	return &proto.ResponseResult{}, nil
}

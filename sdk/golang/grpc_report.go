package plugin

import (
	"context"

	log "github.com/hashicorp/go-hclog"

	"github.com/baetyl/baetyl-gateway-sdk/sdk/golang/proto"
)

var _ Report = &gRPCReportClient{}

type gRPCReportClient struct {
	client proto.ReportClient
}

func (m *gRPCReportClient) Post(req *Request) (*Response, error) {
	res, err := m.client.Post(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{Data: res.Data}, nil
}

func (m *gRPCReportClient) State(req *Request) (*Response, error) {
	res, err := m.client.State(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{Data: res.Data}, nil
}

type gRPCReportServer struct {
	Impl Report
	log  log.Logger
}

func (m *gRPCReportServer) Post(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	if req != nil {
		m.log.Debug("req request invalid")
		return &proto.ResponseResult{}, nil
	}
	m.log.Debug("req request", req.Request)
	res, err := m.Impl.Post(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}
	return &proto.ResponseResult{Data: res.Data}, nil
}

func (m *gRPCReportServer) State(_ context.Context, req *proto.RequestArgs) (resp *proto.ResponseResult, err error) {
	res, err := m.Impl.State(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}
	return &proto.ResponseResult{Data: res.Data}, nil
}

package golang

import (
	"context"

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
}

func (m *gRPCReportServer) Post(_ context.Context, req *proto.RequestArgs) (resp *proto.ResponseResult, err error) {
	res, err := m.Impl.Post(&Request{
		Req: req.Request,
	})

	return &proto.ResponseResult{Data: res.Data}, err
}

func (m *gRPCReportServer) State(_ context.Context, req *proto.RequestArgs) (resp *proto.ResponseResult, err error) {
	res, err := m.Impl.State(&Request{
		Req: req.Request,
	})

	return &proto.ResponseResult{Data: res.Data}, err
}

# !/usr/bin/env python3
from proto import driver_pb2_grpc
from proto.driver_pb2 import ResponseResult


class gRPCClient(driver_pb2_grpc.Report):
    client: driver_pb2_grpc.ReportStub

    @staticmethod
    def NewClient(report: driver_pb2_grpc.ReportStub) :
        c = gRPCClient()
        c.client = report
        return c

    def Post(self, requestArgs,** kwargs):
        response = self.client.Post(requestArgs)
        return ""

    def State(self, requestArgs,** kwargs):
        response = self.client.State(requestArgs.request)
        return ""

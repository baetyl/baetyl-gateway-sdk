# !/usr/bin/env python3
import grpc
from proto import driver_pb2_grpc
from proto import driver_pb2
from interface_driver import IDriver
from grpc_client import gRPCClient


class gRPCServer(driver_pb2_grpc.Driver):
    driver :IDriver

    @staticmethod
    def NewDriver(driver : IDriver):
        d = gRPCServer()
        d.driver = driver
        return d

    def GetDriverInfo(self,requestArgs,context,** kwargs):
        data = self.driver.GetDriverInfo(requestArgs.request)
        return driver_pb2.ResponseResult(data)

    def Setup(self,requestArgs: driver_pb2.RequestArgs,context,** kwargs):
        channel = grpc.insecure_channel(target="127.0.0.1:"+str(requestArgs.brokerid))
        print("接收到requestArgs.brokerid"+str(requestArgs.brokerid))
        client = driver_pb2_grpc.ReportStub(channel)
        report = gRPCClient.NewClient(client)
        data = self.driver.Setup(requestArgs.request,report)
        return driver_pb2.ResponseResult(data=data)

    def SetConfig(self, requestArgs, context,** kwargs):
        data = self.driver.SetConfig(requestArgs.request)
        return driver_pb2.ResponseResult(data=data)

    def Set(self,requestArgs,context,** kwargs):
        data= self.driver.Set(requestArgs.request)
        return driver_pb2.ResponseResult(data)

    def Get(self,requestArgs,context,** kwargs):
        data=  self.driver.Get(requestArgs.request)
        return driver_pb2.ResponseResult(data)

    def Start(self, requestArgs, context,** kwargs):
        data = self.driver.Start(requestArgs.request)
        return  driver_pb2.ResponseResult(data=data)


    def Restart(self, requestArgs, context,** kwargs):
        data = self.driver.Restart(requestArgs.request)
        return driver_pb2.ResponseResult(data)

    def Stop(self, requestArgs, context,** kwargs):
        data = self.driver.stop(requestArgs.request)
        return driver_pb2.ResponseResult(data)
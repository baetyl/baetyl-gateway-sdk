package com.baidu.bce.sdk.plugin;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import proto.DriverOuterClass;
import proto.ReportGrpc;

public class GrpcReport implements IReport {
    private final ManagedChannel channel;
    private final ReportGrpc.ReportBlockingStub blockingStub;

    public GrpcReport(String host, int port) {
        this.channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();
        this.blockingStub = ReportGrpc.newBlockingStub(channel);
    }

    public void shutdown() {
        channel.shutdown();
    }

    @Override
    public String post(String data) {
        DriverOuterClass.RequestArgs args = DriverOuterClass.RequestArgs.newBuilder().setRequest(data).build();
        DriverOuterClass.ResponseResult response = blockingStub.post(args);
        return response.getData();
    }

    @Override
    public String state(String data) {
        DriverOuterClass.RequestArgs args = DriverOuterClass.RequestArgs.newBuilder().setRequest(data).build();
        DriverOuterClass.ResponseResult response = blockingStub.state(args);
        return response.getData();
    }
}

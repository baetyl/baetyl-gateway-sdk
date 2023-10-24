package com.baidu.bce.sdk.plugin;

import io.grpc.stub.StreamObserver;
import proto.DriverGrpc;
import proto.DriverOuterClass;

public class GrpcDriver extends DriverGrpc.DriverImplBase {
    private final IDriver driver;

    public GrpcDriver(IDriver driver) {
        this.driver = driver;
    }

    @Override
    public void getDriverInfo(DriverOuterClass.RequestArgs request,
                              StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.getDriverInfo(request.getRequest());
        this.send(data, responseObserver);
    }

    @Override
    public void setConfig(DriverOuterClass.RequestArgs request,
                          StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.setConfig(request.getRequest());
        this.send(data, responseObserver);
    }

    @Override
    public void setup(DriverOuterClass.RequestArgs request,
                      StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        // 根据传递来的信息，初始化一个 grpc client 来连接宿主的 report server
        try {
            IReport report = new GrpcReport("0.0.0.0", request.getBrokerid());
            String data = this.driver.setup(request.getRequest(), report);
            this.send(data, responseObserver);
        } catch (Exception e) {
            System.err.println(e);
        }

    }

    @Override
    public void start(DriverOuterClass.RequestArgs request,
                      StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.start(request.getRequest());
        this.send(data, responseObserver);
    }

    @Override
    public void restart(DriverOuterClass.RequestArgs request,
                        StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.restart(request.getRequest());
        this.send(data, responseObserver);
    }

    @Override
    public void stop(DriverOuterClass.RequestArgs request,
                     StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.stop(request.getRequest());
        this.send(data, responseObserver);
    }

    @Override
    public void get(DriverOuterClass.RequestArgs request,
                    StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.get(request.getRequest());
        this.send(data, responseObserver);
    }

    @Override
    public void set(DriverOuterClass.RequestArgs request,
                    StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        String data = this.driver.set(request.getRequest());
        this.send(data, responseObserver);
    }

    private void send(String data, StreamObserver<DriverOuterClass.ResponseResult> responseObserver) {
        DriverOuterClass.ResponseResult res = DriverOuterClass.ResponseResult.newBuilder().setData(data).build();
        responseObserver.onNext(res);
        responseObserver.onCompleted();
    }
}

package com.baidu.bce.sdk.plugin;

import io.grpc.health.v1.HealthCheckRequest;
import io.grpc.health.v1.HealthCheckResponse;
import io.grpc.health.v1.HealthGrpc;
import io.grpc.protobuf.services.HealthStatusManager;
import io.grpc.stub.StreamObserver;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class GrpcHealth extends HealthGrpc.HealthImplBase {
    private final HealthStatusManager healthStatusManager;
    private Map<String, HealthCheckResponse.ServingStatus> statusMap;

    public GrpcHealth() {
        this.healthStatusManager = new HealthStatusManager();
        this.statusMap = new ConcurrentHashMap<>();

        // go-plugin 通过下面这个来检查插件健康状态
        this.statusMap.put("plugin", HealthCheckResponse.ServingStatus.SERVING);
    }

    public void setServingStatus(String serviceName, HealthCheckResponse.ServingStatus status) {
        this.statusMap.put(serviceName, status);
    }

    public void deleteServingStatus(String serviceName) {
        this.statusMap.remove(serviceName);
    }

    @Override
    public void check(HealthCheckRequest request, StreamObserver<HealthCheckResponse> responseObserver) {
        String serviceName = request.getService();
        HealthCheckResponse.ServingStatus status = HealthCheckResponse.ServingStatus.UNRECOGNIZED;
        if (this.statusMap.containsKey(serviceName)) {
            status = this.statusMap.get(serviceName);
        }
        HealthCheckResponse response = HealthCheckResponse.newBuilder()
                .setStatus(status)
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void watch(HealthCheckRequest request, StreamObserver<HealthCheckResponse> responseObserver) {
        super.watch(request, responseObserver);
    }
}

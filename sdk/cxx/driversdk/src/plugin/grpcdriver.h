#pragma once

#include "idriver.h"
#include "driver.pb.h"
#include "driver.grpc.pb.h"

class GrpcDriver final : public proto::Driver::Service {
public:
    explicit GrpcDriver(IDriver &idriver);

    grpc::Status GetDriverInfo(grpc::ServerContext *context, const proto::RequestArgs *request,
                               proto::ResponseResult *response) override;

    grpc::Status SetConfig(grpc::ServerContext *context, const proto::RequestArgs *request,
                           proto::ResponseResult *response) override;

    grpc::Status
    Setup(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) override;

    grpc::Status
    Start(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) override;

    grpc::Status
    Restart(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) override;

    grpc::Status
    Stop(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) override;

    grpc::Status
    Get(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) override;

    grpc::Status
    Set(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) override;

private:
    IDriver &driver;
};

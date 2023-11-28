#pragma once

#include <grpcpp/create_channel.h>
#include <grpcpp/client_context.h>
#include "ireport.h"
#include "driver.pb.h"
#include "driver.grpc.pb.h"

class GrpcReport final : public IReport {
public:
    explicit GrpcReport(std::string &target);

    std::string post(const std::string &data) override;

    std::string state(const std::string &data) override;

private:
    std::string addr;
    grpc::ClientContext ctx;
    std::unique_ptr<proto::Report::Stub> stub;
};
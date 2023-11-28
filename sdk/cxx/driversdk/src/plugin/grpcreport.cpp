#include "grpcreport.h"

std::string GrpcReport::post(const std::string &data) {
    proto::ResponseResult res;
    proto::RequestArgs req;
    req.set_request(data);
    grpc::Status status = stub->Post(&ctx, req, &res);
    if (status.ok()) {
        return res.data();
    }
    std::string msg = "failed to post, grpc status code :" + std::to_string(status.error_code()) + ", status reason:" +
                      status.error_details();
    return msg;
}

std::string GrpcReport::state(const std::string &data) {
    proto::ResponseResult res;
    proto::RequestArgs req;
    req.set_request(data);
    grpc::Status status = stub->State(&ctx, req, &res);
    if (status.ok()) {
        return res.data();
    }
    std::string msg = "failed to state, grpc status code :" + std::to_string(status.error_code()) + ", status reason:" +
                      status.error_details();
    return msg;
}

GrpcReport::GrpcReport(std::string &target) : addr(target) {
    auto channel = grpc::CreateChannel(addr, grpc::InsecureChannelCredentials());
    stub = proto::Report::NewStub(channel);
}

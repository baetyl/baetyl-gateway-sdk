#include "grpcdriver.h"
#include "grpcreport.h"

GrpcDriver::GrpcDriver(IDriver &idriver) : driver(idriver) {}

grpc::Status GrpcDriver::GetDriverInfo(grpc::ServerContext *context, const proto::RequestArgs *request,
                                       proto::ResponseResult *response) {
    std::string data = driver.getDriverInfo(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status GrpcDriver::SetConfig(grpc::ServerContext *context, const proto::RequestArgs *request,
                                   proto::ResponseResult *response) {
    std::string data = driver.setConfig(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status
GrpcDriver::Setup(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) {
    std::string port = std::to_string(request->brokerid());
    std::string addr = "0.0.0.0:" + port;
    GrpcReport report(addr);

    std::string data = driver.setup(request->request(), &report);
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status
GrpcDriver::Start(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) {
    std::string data = driver.start(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status
GrpcDriver::Restart(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) {
    std::string data = driver.restart(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status
GrpcDriver::Stop(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) {
    std::string data = driver.stop(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status
GrpcDriver::Get(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) {
    std::string data = driver.get(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}

grpc::Status
GrpcDriver::Set(grpc::ServerContext *context, const proto::RequestArgs *request, proto::ResponseResult *response) {
    std::string data = driver.set(request->request());
    response->set_data(data);
    return grpc::Status::OK;
}


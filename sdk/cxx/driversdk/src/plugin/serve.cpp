#include "serve.h"
#include <thread>
#include <iostream>
#include <csignal>
#include <grpcpp/server_builder.h>

const grpc::string kHealthyService("plugin");

Serve::Serve(IDriver *driver) : grpcDriver(*driver) {
    ppid = getppid();
}

void Serve::checker() {
    std::thread([&]() {
        while (true) {
            if (getppid() != ppid) {
                std::cout << "parent not alive, exit" << std::endl;
                exit(0);
            }

            if (kill(ppid, 0) != 0) {
                std::cout << "parent not alive, exit" << std::endl;
                exit(0);
            }
            std::this_thread::sleep_for(std::chrono::seconds(5));
        }
    }).detach();
}

void Serve::start() {
    checker();

    grpc::EnableDefaultHealthCheckService(true);
    grpc::ServerBuilder builder;
    int selectedPort;
    builder.AddListeningPort("0.0.0.0:0", grpc::InsecureServerCredentials(),&selectedPort);
    builder.RegisterService(&grpcDriver);
    std::unique_ptr<grpc::Server> svr(builder.BuildAndStart());

    grpc::HealthCheckServiceInterface* health = svr->GetHealthCheckService();
    health->SetServingStatus(kHealthyService, true);

    std::cout << "1|1|tcp|127.0.0.1:" << selectedPort << "|grpc" << std::endl;
    svr->Wait();
}

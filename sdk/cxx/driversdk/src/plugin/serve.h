#pragma once

#include "grpcdriver.h"
#include "idriver.h"

class Serve {
public:
    explicit Serve(IDriver* driver);
    void start();

private:
    GrpcDriver grpcDriver;
    pid_t ppid;

    void checker();
};

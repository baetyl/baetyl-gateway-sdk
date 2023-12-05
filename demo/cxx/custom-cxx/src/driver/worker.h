#pragma once

#include <iostream>
#include <map>
#include <thread>
#include <chrono>

#include "models/job.h"
#include "models/message.h"
#include "device.h"
#include "plugin/ireport.h"
#include "utils/l.h"

class Worker {
public:
    // Constructor
    Worker(const Job &j, const Device &dev, const std::string &dn, IReport &r);

    // Member functions
    void working();

    void requestExit();

    void reportProperty();

    Job& getJob();

    Device& getDevice();

private:
    Job job;
    Device device;
    std::string driverName;
    IReport &report;
    bool exitFlag = false;

    std::string msg2JSON(DRIVERSDK::Message<std::map<std::string, float>> &msg);
};

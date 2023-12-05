#pragma once

#include <iostream>
#include <map>
#include <vector>
#include <thread>

#include "context/context.h"
#include "models/driverconfig.h"
#include "plugin/ireport.h"
#include "utils/l.h"
#include "device.h"
#include "worker.h"
#include "models/event.h"

class Custom {
public:
    // Constructor
    Custom(DRIVERSDK::Context& ctx,  DriverConfig& cfg, IReport& report);

    // Member functions
    void start();
    void restart();
    void stop();
    void set(const DRIVERSDK::DeviceInfo& info, const std::map<std::string, float>& props);
    void event(const DRIVERSDK::DeviceInfo& info, const DRIVERSDK::Event<std::string>& event);
    void propertyGet(const DRIVERSDK::DeviceInfo& info);

    DRIVERSDK::Context& getCtx();

private:
    DRIVERSDK::Context& ctx;
    DriverConfig& cfg;
    std::map<std::string, Worker> ws;
    std::map<std::string, Device> devs;
    IReport& report;
    std::map<std::string, std::thread> threads;
};

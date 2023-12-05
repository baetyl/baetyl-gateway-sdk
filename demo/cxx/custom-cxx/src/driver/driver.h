#pragma once

#include <iostream>
#include <vector>
#include <map>

#include "plugin/idriver.h"
#include "context/context.h"
#include "models/driverconfig.h"
#include "custom.h"
#include "models/message.h"
#include "models/event.h"
#include "models/property.h"
#include "utils/l.h"
#include "plugin//ireport.h"

class Driver : public IDriver {
public:
    std::string getDriverInfo(const std::string &data) override;
    std::string setConfig(const std::string &data) override;
    std::string setup(const std::string &driver, IReport *report) override;
    std::string start(const std::string &data) override;
    std::string restart(const std::string &data) override;
    std::string stop(const std::string &data) override;
    std::string get(const std::string &data) override;
    std::string set(const std::string &data) override;

private:
    std::string driverName;
    std::string configPath;
    DriverConfig config;
    IReport *report;
    Custom *custom;

    DriverConfig loadConfig(DRIVERSDK::Context &dm, const std::string &dr);
};

#pragma once

#include <string>
#include "ireport.h"

class IDriver {
public:
    virtual std::string getDriverInfo(const std::string& data) = 0;
    virtual std::string setConfig(const std::string& data) = 0;
    virtual std::string setup(const std::string& driver, IReport* report) = 0;
    virtual std::string start(const std::string& data) = 0;
    virtual std::string restart(const std::string& data) = 0;
    virtual std::string stop(const std::string& data) = 0;
    virtual std::string get(const std::string& data) = 0;
    virtual std::string set(const std::string& data) = 0;

    virtual ~IDriver() = default;
};
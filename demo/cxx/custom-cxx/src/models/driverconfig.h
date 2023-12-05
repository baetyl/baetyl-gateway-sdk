#pragma once

#include <string>
#include <vector>

#include "deviceconfig.h"
#include "job.h"

class DriverConfig {
public:
    // Constructors
    DriverConfig();

    DriverConfig(const std::string &driverName, const std::vector <DeviceConfig> &devices,
                 const std::vector <Job> &jobs);

    // Getters
    const std::string &getDriverName() const;

    const std::vector <DeviceConfig> &getDevices() const;

    const std::vector <Job> &getJobs() const;

    // Setters
    void setDriverName(const std::string &driverName);

    void setDevices(const std::vector <DeviceConfig> &devices);

    void setJobs(const std::vector <Job> &jobs);

private:
    std::string driverName;
    std::vector <DeviceConfig> devices;
    std::vector <Job> jobs;
};

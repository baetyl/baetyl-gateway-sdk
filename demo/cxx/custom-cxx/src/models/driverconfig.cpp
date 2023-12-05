#include "driverconfig.h"

// Default Constructor
DriverConfig::DriverConfig() {}

// Parameterized Constructor
DriverConfig::DriverConfig(const std::string& dname, const std::vector<DeviceConfig>& devs, const std::vector<Job>& js)
        : driverName(dname), devices(devs), jobs(js) {}

// Getters
const std::string& DriverConfig::getDriverName() const {
    return driverName;
}

const std::vector<DeviceConfig>& DriverConfig::getDevices() const {
    return devices;
}

const std::vector<Job>& DriverConfig::getJobs() const {
    return jobs;
}

// Setters
void DriverConfig::setDriverName(const std::string& dname) {
    this->driverName = dname;
}

void DriverConfig::setDevices(const std::vector<DeviceConfig>& devs) {
    this->devices = devs;
}

void DriverConfig::setJobs(const std::vector<Job>& js) {
    this->jobs = js;
}

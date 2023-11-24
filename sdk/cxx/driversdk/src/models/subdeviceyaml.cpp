#include "subdeviceyaml.h"

namespace DRIVERSDK {

// Default constructor
    SubDeviceYaml::SubDeviceYaml() {}

// Parameterized constructor
    SubDeviceYaml::SubDeviceYaml(const std::vector<DeviceInfo>& subDevices, const std::string& subDriver)
            : devices(subDevices), driver(subDriver) {}

// Getter and setter methods
    const std::vector<DeviceInfo>& SubDeviceYaml::getDevices() const {
        return devices;
    }

    void SubDeviceYaml::setDevices(const std::vector<DeviceInfo>& subDevices) {
        devices = subDevices;
    }

    const std::string& SubDeviceYaml::getDriver() const {
        return driver;
    }

    void SubDeviceYaml::setDriver(const std::string& subDriver) {
        driver = subDriver;
    }

} // namespace DRIVERSDK

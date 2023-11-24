#pragma once

#include <list>
#include "deviceinfo.h" // Include the DeviceInfo header

namespace DRIVERSDK {

    class SubDeviceYaml {
    private:
        std::vector<DeviceInfo> devices;
        std::string driver;

    public:
        // Default constructor
        SubDeviceYaml();

        // Parameterized constructor
        SubDeviceYaml(const std::vector<DeviceInfo>& subDevices, const std::string& subDriver);

        // Getter and setter methods
        const std::vector<DeviceInfo>& getDevices() const;
        void setDevices(const std::vector<DeviceInfo>& subDevices);

        const std::string& getDriver() const;
        void setDriver(const std::string& subDriver);
    };

} // namespace DRIVERSDK


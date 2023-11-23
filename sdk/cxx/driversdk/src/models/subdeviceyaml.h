// subdeviceyaml.h

#ifndef DRIVERSDK_SUBDEVICEYAML_H
#define DRIVERSDK_SUBDEVICEYAML_H

#include <list>
#include "deviceinfo.h" // Include the DeviceInfo header

namespace DRIVERSDK {

    class SubDeviceYaml {
    private:
        std::list<DeviceInfo> devices;
        std::string driver;

    public:
        // Default constructor
        SubDeviceYaml();

        // Parameterized constructor
        SubDeviceYaml(const std::list<DeviceInfo>& subDevices, const std::string& subDriver);

        // Getter and setter methods
        const std::list<DeviceInfo>& getDevices() const;
        void setDevices(const std::list<DeviceInfo>& subDevices);

        const std::string& getDriver() const;
        void setDriver(const std::string& subDriver);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_SUBDEVICEYAML_H

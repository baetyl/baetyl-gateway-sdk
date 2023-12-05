#pragma once

#include <string>

class DeviceConfig {
public:
    // Constructors
    DeviceConfig();
    DeviceConfig(const std::string& device);

    // Getter
    const std::string& getDevice() const;

    // Setter
    void setDevice(const std::string& device);

private:
    std::string device;
};

#pragma once

#include "models/deviceconfig.h"
#include "models/deviceinfo.h"
#include "simulator.h"

class Device {
public:
    // Constructor
    Device();
    Device(const DeviceConfig& cfg, DRIVERSDK::DeviceInfo&  info);

    // Member functions
    float get(int index);
    void set(int index, float val);

    // Getters and setters
    const DRIVERSDK::DeviceInfo& getInfo() const;
    void setInfo(const DRIVERSDK::DeviceInfo& info);

private:
    DeviceConfig cfg;
    DRIVERSDK::DeviceInfo info;
    Simulator cli;
};

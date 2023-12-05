#pragma once

#include <string>
#include <vector>

#include "property.h"

class Job {
public:
    // Constructors
    Job();
    Job(const std::string& device, long interval, const std::vector<Property>& properties);

    // Getters
    const std::string& getDevice() const;
    long getInterval() const;
    const std::vector<Property>& getProperties() const;

    // Setters
    void setDevice(const std::string& dev);
    void setInterval(long val);
    void setProperties(const std::vector<Property>& props);

private:
    std::string device;
    long interval;
    std::vector<Property> properties;
};

#include "job.h"

// Default Constructor
Job::Job() : interval(0) {}

// Parameterized Constructor
Job::Job(const std::string& dev, long val, const std::vector<Property>& props)
        : device(dev), interval(val), properties(props) {}

// Getters
const std::string& Job::getDevice() const {
    return device;
}

long Job::getInterval() const {
    return interval;
}

const std::vector<Property>& Job::getProperties() const {
    return properties;
}

// Setters
void Job::setDevice(const std::string& dev) {
    this->device = dev;
}

void Job::setInterval(long val) {
    this->interval = val;
}

void Job::setProperties(const std::vector<Property>& props) {
    this->properties = props;
}

#include "worker.h"
#include "context/context.h"
#include <nlohmann/json.hpp>

// Constructor
Worker::Worker(const Job &j, const Device &dev, const std::string &dn, IReport &r)
        : job(j), device(dev), driverName(dn), report(r) {
}

void Worker::working() {
    L::debug("Worker interval " + std::to_string(job.getInterval()) + " ns");

    std::thread([this]() {
        while (!exitFlag) {
            reportProperty();
            std::this_thread::sleep_for(std::chrono::nanoseconds(job.getInterval()));
        }
    }).detach();
}

void Worker::requestExit() {
    exitFlag = true;
}

void Worker::reportProperty() {
    L::debug("Worker reportProperty");

    std::map<std::string, float> props;
    for (const Property &item: job.getProperties()) {
        float val = device.get(item.getIndex());
        props[item.getName()] = val;
    }

    std::map<std::string, std::string> meta;
    meta[DRIVERSDK::Context::KEY_DRIVER_NAME] = driverName;
    meta[DRIVERSDK::Context::KEY_DEVICE_NAME] = device.getInfo().getName();
    DRIVERSDK::Message<std::map<std::string, float>> msg(DRIVERSDK::Context::MESSAGE_DEVICE_REPORT, meta, props);

    std::string jStr = msg2JSON(msg);

    L::debug("Worker reportProperty msg " + jStr);
    report.post(jStr);
}

Job &Worker::getJob() {
    return job;
}

Device &Worker::getDevice() {
    return device;
}

std::string Worker::msg2JSON(DRIVERSDK::Message<std::map<std::string, float>> &msg) {
    nlohmann::json j = nlohmann::json{{"kind",    msg.getKind()},
                                      {"meta",    msg.getMeta()},
                                      {"content", msg.getContent()}};
    return j.dump();
}
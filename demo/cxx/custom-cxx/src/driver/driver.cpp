#include <nlohmann/json.hpp>
#include "driver.h"

// Implementing IDriver interface
std::string Driver::getDriverInfo(const std::string &data) {
    nlohmann::json j = nlohmann::json{{"name",       driverName},
                                      {"configPath", configPath},
                                      {"driverName", config.getDriverName()}};
    return j.dump();
}

std::string Driver::setConfig(const std::string &data) {
    L::debug("Driver setConfig " + data);
    configPath = data;
    return "plugin " + driverName + " setConfig success";
}

std::string Driver::setup(const std::string &driver, IReport *r) {
    L::debug("Driver setup " + driver);
    driverName = driver;
    this->report = r;
    return "plugin " + driver + " setup success";
}

std::string Driver::start(const std::string &data) {
    L::debug("Driver start " + data);
    DRIVERSDK::Context ctx;
    ctx.loadYamlConfig(configPath, driverName);
    config = loadConfig(ctx, driverName);
    Custom c(ctx, config, *report);
    custom = &c;
    custom->start();
    return "plugin " + driverName + " start success";
}

std::string Driver::restart(const std::string &data) {
    L::debug("Driver restart " + data);
    custom->restart();
    return "plugin " + driverName + " restart success";
}

std::string Driver::stop(const std::string &data) {
    L::debug("Driver stop " + data);
    custom->stop();
    return "plugin " + driverName + " stop success";
}

std::string Driver::get(const std::string &data) {
    L::debug("Driver get " + data);

    nlohmann::json j = nlohmann::json::parse(data);

    std::string dr = j[DRIVERSDK::Context::KEY_DRIVER_NAME];
    std::string deviceName = j[DRIVERSDK::Context::KEY_DEVICE_NAME];

    if (dr.empty() || deviceName.empty()) {
        return "plugin " + dr + " get failed";
    }

    DRIVERSDK::DeviceInfo info = custom->getCtx().getDevice(dr, deviceName);

    std::string kd = j["kind"];

    if (kd == DRIVERSDK::Context::MESSAGE_DEVICE_EVENT) {
        DRIVERSDK::Event<std::string> event;
        event.setType(j["content"]["type"]);
        event.setPayload(j["content"]["payload"]);
        custom->event(info, event);
    } else if (kd == DRIVERSDK::Context::MESSAGE_DEVICE_PROPERTY_GET) {
        custom->propertyGet(info);
    } else {
        L::debug("driver get unsupported message type");
    }
    return "plugin " + dr + " get success";
}

std::string Driver::set(const std::string &data) {
    L::debug("Driver set " + data);

    nlohmann::json j = nlohmann::json::parse(data);
    std::string dr = j[DRIVERSDK::Context::KEY_DRIVER_NAME];
    std::string deviceName = j[DRIVERSDK::Context::KEY_DEVICE_NAME];

    if (dr.empty() || deviceName.empty()) {
        return "plugin " + dr + " get failed";
    }

    DRIVERSDK::DeviceInfo info = custom->getCtx().getDevice(dr, deviceName);

    custom->set(info, j["content"]);
    return "plugin " + driverName + " set success";
}

DriverConfig Driver::loadConfig(DRIVERSDK::Context &dm, const std::string &dr) {
    std::vector<DeviceConfig> devices;
    std::vector<Job> jobs;

    for (const DRIVERSDK::DeviceInfo &info: dm.getAllDevices(dr)) {
        DRIVERSDK::AccessConfig accessConfig = info.getAccessConfig();

        DeviceConfig device;
        device.setDevice(info.getName());
        devices.push_back(device);

        std::vector<Property> jobProps;

        DRIVERSDK::AccessTemplate tpl = dm.getAccessTemplate(dr, info.getAccessTemplate());
        for (const DRIVERSDK::DeviceProperty &prop: tpl.getProperties()) {
            std::string visitor = prop.getVisitor().getCustom();
            if (!visitor.empty()) {
                nlohmann::json j = nlohmann::json::parse(visitor);
                Property jobProp(j["name"], j["type"], j["index"]);
                jobProps.push_back(jobProp);
            }
        }

        Job job;
        nlohmann::json j = nlohmann::json::parse(accessConfig.getCustom());
        job.setDevice(j["device"]);
        job.setInterval(j["interval"]);
        for (const auto &item: j["properties"]) {
            Property prop;
            prop.setName(item["name"]);
            prop.setType(item["type"]);
            prop.setIndex(item["index"]);
        }

        job.setDevice(info.getName());
        job.setProperties(jobProps);

        jobs.push_back(job);
    }
}


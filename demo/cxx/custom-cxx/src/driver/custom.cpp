#include "custom.h"

// Constructor
Custom::Custom(DRIVERSDK::Context &ctx, DriverConfig &cfg, IReport &report)
        : ctx(ctx), cfg(cfg), report(report), devs{}, ws{}, threads{} {

    std::map<std::string, DRIVERSDK::DeviceInfo> infos;
    for (const DRIVERSDK::DeviceInfo &info: ctx.getAllDevices(cfg.getDriverName())) {
        infos.insert(std::make_pair(info.getName(), info));
    }

    for (const DeviceConfig &item: cfg.getDevices()) {
        if (infos.count(item.getDevice()) > 0) {
            devs.insert(std::make_pair(item.getDevice(), Device(item, infos[item.getDevice()])));
        }
    }

    for (const Job &job: cfg.getJobs()) {
        auto it = devs.find(job.getDevice());
        if (it != devs.end()) {
            ws.insert(std::make_pair(it->second.getInfo().getName(),
                                     Worker(job, it->second, cfg.getDriverName(), report)));
        }
    }
}

DRIVERSDK::Context& Custom::getCtx(){
    return ctx;
}

// Member function implementations...
void Custom::start() {
    for (const auto &pair: ws) {
        const std::string key = pair.first;
        const Worker &w = pair.second;
        threads.insert(std::make_pair(key, std::thread([this, key, w]() {
            this->ws.find(key)->second.working();
            threads.erase(key);
        })));
        threads.find(key)->second.detach();
    }
}

void Custom::restart() {
    stop();
    start();
}

void Custom::stop() {
    for (auto &threadPair: threads) {
        ws.find(threadPair.first)->second.requestExit();
        threadPair.second.join();
    }
    threads.clear();
}

void Custom::set(const DRIVERSDK::DeviceInfo &info, const std::map<std::string, float> &props) {
    L::debug("Custom set (name = " + info.getName() + ")");
    auto it = ws.find(info.getName());
    if (it != ws.end()) {
        Worker &w = it->second;
        for (const auto &entry: props) {
            for (const Property &p: w.getJob().getProperties()) {
                if (entry.first == p.getName()) {
                    w.getDevice().set(p.getIndex(), entry.second);
                }
            }
        }
    }
}

void Custom::event(const DRIVERSDK::DeviceInfo &info, const DRIVERSDK::Event<std::string> &event) {
    if (event.getType() == DRIVERSDK::Context::TYPE_REPORT_EVENT) {
        propertyGet(info);
    } else {
        L::debug("custom event " + info.getName());
    }
}

void Custom::propertyGet(const DRIVERSDK::DeviceInfo &info) {
    L::debug("Custom get (name = " + info.getName() + ")");
    auto it = ws.find(info.getName());
    if (it != ws.end()) {
        Worker &w = it->second;
        try {
            w.reportProperty();
        } catch (const std::exception &e) {
            e.what();
        }
    }
}

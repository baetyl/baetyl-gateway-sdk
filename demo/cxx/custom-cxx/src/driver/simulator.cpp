#include "simulator.h"
#include <cstdlib>

Simulator::Simulator(const std::string &n)
        : name(n) {
    srand((unsigned int) time(nullptr));
    point.setTemperature(static_cast<float>(rand() % 100));
    point.setHumidity(static_cast<float>(rand() % 100));
    point.setPressure(static_cast<float>(rand() % 100));

    generateSimulateData();
}

void Simulator::set(int index, float val) {
    switch (index) {
        case Point::INDEX_TEMPERATURE:
            point.setTemperature(val);
            break;
        case Point::INDEX_HUMIDITY:
            point.setHumidity(val);
            break;
        case Point::INDEX_PRESSURE:
            point.setPressure(val);
            break;
        default:
            return;
    }
}

float Simulator::get(int index) {
    switch (index) {
        case Point::INDEX_TEMPERATURE:
            return point.getTemperature();
        case Point::INDEX_HUMIDITY:
            return point.getHumidity();
        case Point::INDEX_PRESSURE:
            return point.getPressure();
        default:
            return 0.0f;
    }
}

void Simulator::generateSimulateData() {
    std::thread([this]() {
        while (true) {
            set(Point::INDEX_TEMPERATURE, static_cast<float>(random() % 100));
            set(Point::INDEX_HUMIDITY, static_cast<float>(random() % 100));
            set(Point::INDEX_PRESSURE, static_cast<float>(random() % 100));

            L::debug("Simulate generate random temperature = " + std::to_string(get(Point::INDEX_TEMPERATURE)));
            L::debug("Simulate generate random humidity = " + std::to_string(get(Point::INDEX_HUMIDITY)));
            L::debug("Simulate generate random pressure = " + std::to_string(get(Point::INDEX_PRESSURE)));

            std::this_thread::sleep_for(std::chrono::seconds(SIMULATOR_INTERVAL));
        }
    }).detach();
}

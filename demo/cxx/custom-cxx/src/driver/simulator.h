#pragma once

#include <string>
#include <random>
#include <iostream>
#include <thread>
#include <chrono>

#include "models/point.h"
#include "utils/l.h"

class Simulator {
public:
    explicit Simulator(const std::string &n);
    void set(int index, float val);
    float get(int index);

private:
    static const int SIMULATOR_INTERVAL = 10;
    std::string name;
    Point point;

    void generateSimulateData();
};

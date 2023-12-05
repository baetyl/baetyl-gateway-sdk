#include "point.h"

// Default Constructor
Point::Point() : temperature(0.0f), humidity(0.0f), pressure(0.0f) {}

// Parameterized Constructor
Point::Point(float temperature, float humidity, float pressure)
        : temperature(temperature), humidity(humidity), pressure(pressure) {}

// Getters
float Point::getTemperature() const {
    return temperature;
}

float Point::getHumidity() const {
    return humidity;
}

float Point::getPressure() const {
    return pressure;
}

// Setters
void Point::setTemperature(float tp) {
    this->temperature = tp;
}

void Point::setHumidity(float hu) {
    this->humidity = hu;
}

void Point::setPressure(float ps) {
    this->pressure = ps;
}

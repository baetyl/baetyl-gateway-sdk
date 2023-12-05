#pragma once

class Point {
public:
    // Constants
    static const int INDEX_TEMPERATURE = 0;
    static const int INDEX_HUMIDITY = 1;
    static const int INDEX_PRESSURE = 2;

    // Constructors
    Point();
    Point(float temperature, float humidity, float pressure);

    // Getters
    float getTemperature() const;
    float getHumidity() const;
    float getPressure() const;

    // Setters
    void setTemperature(float tp);
    void setHumidity(float hu);
    void setPressure(float ps);

private:
    float temperature;
    float humidity;
    float pressure;
};

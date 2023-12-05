#pragma once

#include <iostream>
#include <variant>
#include <string>

namespace DRIVERSDK {

    template <typename T>
    class Event {
    public:
        // Constructors
        Event();
        Event(const std::string& type, const T& payload);

        // Getters
        std::string getType() const;
        T getPayload() const;

        // Setters
        void setType(const std::string& type);
        void setPayload(const T& payload);

    private:
        std::string type;
        T payload;
    };

} // namespace DRIVERSDK

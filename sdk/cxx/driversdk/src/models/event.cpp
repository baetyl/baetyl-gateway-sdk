#include "event.h"

namespace DRIVERSDK {
    Event::Event(const std::string &eventType, void *eventPayload) : type(eventType), payload(eventPayload) {}

    const std::string &Event::getType() const {
        return type;
    }

    void Event::setType(const std::string &eventType) {
        type = eventType;
    }

    const void *Event::getPayload() const {
        return payload;
    }

    void Event::setPayload(void *eventPayload) {
        payload = eventPayload;
    }
}
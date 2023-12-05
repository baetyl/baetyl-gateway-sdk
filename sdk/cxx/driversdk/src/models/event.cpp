#include "event.h"

namespace DRIVERSDK {

    // Template implementation
    template <typename T>
    Event<T>::Event() {}

    template <typename T>
    Event<T>::Event(const std::string& type, const T& payload)
            : type(type), payload(payload) {}

    template <typename T>
    std::string Event<T>::getType() const {
        return type;
    }

    template <typename T>
    T Event<T>::getPayload() const {
        return payload;
    }

    template <typename T>
    void Event<T>::setType(const std::string& type) {
        this->type = type;
    }

    template <typename T>
    void Event<T>::setPayload(const T& payload) {
        this->payload = payload;
    }

    // Explicit instantiation for commonly used types
    template class Event<int>;
    template class Event<std::string>;
    template class Event<double>;

} // namespace DRIVERSDK

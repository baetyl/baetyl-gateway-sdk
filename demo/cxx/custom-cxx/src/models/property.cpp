#include "property.h"

Property::Property() : index(0) {}

Property::Property(const std::string& n, const std::string& tp, int idx)
        : name(n), type(tp), index(idx) {}

// Getters
const std::string& Property::getName() const {
    return name;
}

const std::string& Property::getType() const {
    return type;
}

int Property::getIndex() const {
    return index;
}

// Setters
void Property::setName(const std::string& n) {
    this->name = n;
}

void Property::setType(const std::string& tp) {
    this->type = tp;
}

void Property::setIndex(int idx) {
    this->index = idx;
}

#pragma once

#include <string>

class Property {
public:
    // Constructors
    Property();
    Property(const std::string& name, const std::string& type, int index);

    // Getters
    const std::string& getName() const;
    const std::string& getType() const;
    int getIndex() const;

    // Setters
    void setName(const std::string& n);
    void setType(const std::string& tp);
    void setIndex(int idx);

private:
    std::string name;
    std::string type;
    int index;
};

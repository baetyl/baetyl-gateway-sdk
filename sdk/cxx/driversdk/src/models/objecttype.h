#pragma once

#include <string>

namespace DRIVERSDK {

    class ObjectType {
    private:
        std::string displayName;
        std::string type;
        std::string format;

    public:
        // Default constructor
        ObjectType();

        // Parameterized constructor
        ObjectType(const std::string& display, const std::string& objType, const std::string& objFormat);

        // Getter and setter methods
        const std::string& getDisplayName() const;
        void setDisplayName(const std::string& display);

        const std::string& getType() const;
        void setType(const std::string& objType);

        const std::string& getFormat() const;
        void setFormat(const std::string& objFormat);
    };

} // namespace DRIVERSDK


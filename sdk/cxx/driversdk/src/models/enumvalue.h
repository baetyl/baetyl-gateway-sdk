#ifndef DRIVERSDK_ENUMVALUE_H
#define DRIVERSDK_ENUMVALUE_H

#include <string>

namespace DRIVERSDK {

    class EnumValue {
    private:
        std::string name;
        std::string value;
        std::string displayName;

    public:
        // Default constructor
        EnumValue();

        // Parameterized constructor
        EnumValue(const std::string& enumName, const std::string& enumValue, const std::string& display);

        // Getter and setter methods
        const std::string& getName() const;
        void setName(const std::string& enumName);

        const std::string& getValue() const;
        void setValue(const std::string& enumValue);

        const std::string& getDisplayName() const;
        void setDisplayName(const std::string& display);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_ENUMVALUE_H

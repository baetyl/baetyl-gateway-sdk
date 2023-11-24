#pragma once

#include <string>

namespace DRIVERSDK {
    class ArrayType {
    private:
        std::string type;
        int min;
        int max;
        std::string format;

    public:
        ArrayType();

        ArrayType(const std::string &arrayType, int arrayMin, int arrayMax, const std::string &arrayFormat);

        const std::string &getType() const;

        void setType(const std::string &arrayType);

        int getMin() const;

        void setMin(int arrayMin);

        int getMax() const;

        void setMax(int arrayMax);

        const std::string &getFormat() const;

        void setFormat(const std::string &arrayFormat);
    };
}


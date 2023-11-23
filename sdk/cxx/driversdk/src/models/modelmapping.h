#ifndef DRIVERSDK_MODELMAPPING_H
#define DRIVERSDK_MODELMAPPING_H

#include <string>

namespace DRIVERSDK {

    class ModelMapping {
    private:
        std::string attribute;
        std::string type;
        std::string expression;
        int precision;
        double deviation;
        int silentWin;

    public:
        // Default constructor
        ModelMapping();

        // Parameterized constructor
        ModelMapping(const std::string& attr, const std::string& tp, const std::string& expr,
                     int prec, double dev, int silentWindow);

        // Getter and setter methods
        const std::string& getAttribute() const;
        void setAttribute(const std::string& attr);

        const std::string& getType() const;
        void setType(const std::string& tp);

        const std::string& getExpression() const;
        void setExpression(const std::string& expr);

        int getPrecision() const;
        void setPrecision(int prec);

        double getDeviation() const;
        void setDeviation(double dev);

        int getSilentWindow() const;
        void setSilentWindow(int silentWin);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_MODELMAPPING_H

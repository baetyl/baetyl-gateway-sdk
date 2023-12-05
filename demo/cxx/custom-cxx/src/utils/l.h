#pragma once

#include <iostream>
#include <fstream>
#include <ctime>
#include <iomanip>

class L {
public:
    static void info(const std::string& s);
    static void debug(const std::string& s);
    static void warn(const std::string& s);
    static void error(const std::string& s);

private:
    L();  // Private constructor to prevent instantiation
    static void text(const std::string& s);

    static std::ofstream bufferedWriter;
    static std::time_t currentTime;
    static std::tm* localTime;
    static char timeBuffer[80];

    static const std::string filePath;
    static const std::string logInfo;
    static const std::string logDebug;
    static const std::string logWarn;
    static const std::string logError;
};

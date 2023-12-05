#include "l.h"

std::ofstream L::bufferedWriter;
std::time_t L::currentTime;
std::tm* L::localTime;
char L::timeBuffer[80];

const std::string L::filePath = "output/custom-cpp.txt";
const std::string L::logInfo = "[info] ";
const std::string L::logDebug = "[debug] ";
const std::string L::logWarn = "[warn] ";
const std::string L::logError = "[error] ";

L::L() {}

void L::text(const std::string& s) {
    currentTime = std::time(nullptr);
    localTime = std::localtime(&currentTime);
    std::strftime(timeBuffer, sizeof(timeBuffer), "%Y-%m-%d %H:%M:%S.000", localTime);

    std::cerr << "[" << timeBuffer << "] " << s << std::endl;

    if (bufferedWriter.is_open()) {
        bufferedWriter << "[" << timeBuffer << "] " << s << std::endl;
        bufferedWriter.flush();
    }
}

void L::info(const std::string& s) {
    text(logInfo + s);
}

void L::debug(const std::string& s) {
    text(logDebug + s);
}

void L::warn(const std::string& s) {
    text(logWarn + s);
}

void L::error(const std::string& s) {
    text(logError + s);
}


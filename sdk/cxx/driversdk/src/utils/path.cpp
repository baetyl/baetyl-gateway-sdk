// PathUtils.cpp
#include "path.h"

#ifdef _WIN32
#include <windows.h>
#else
#include <unistd.h>
#endif

namespace PathUtils {
    std::string joinPaths(const std::string& path1, const std::string& path2) {
        if (path1.empty() || path1.back() == '/' || path1.back() == '\\') {
            return path1 + path2;
        } else {
#ifdef _WIN32
            return path1 + "\\" + path2;
#else
            return path1 + "/" + path2;
#endif
        }
    }
}

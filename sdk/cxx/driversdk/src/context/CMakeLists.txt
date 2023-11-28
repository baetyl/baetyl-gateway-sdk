find_package(yaml-cpp REQUIRED)
if (NOT yaml-cpp_FOUND)
        message(FATAL_ERROR "yaml-cpp not found. Please install YAML-CPP by running:\n"
                "Ubuntu/Debian:\n"
                "  sudo apt-get install libyaml-cpp-dev\n"
                "Homebrew (macOS):\n"
                "  brew install yaml-cpp\n"
                "Or visit: 3rdparty/README.md and include 3rdparty/3rdfinder/yaml-cpp.cmake\n"
                "Or visit: https://github.com/jbeder/yaml-cpp\n")
endif ()
message("-- Found YAML-CPP version: ${yaml-cpp_VERSION}")

set(DRIVER_CONTEXT_SRC_LIST
        context.cpp
)

add_library(drivercontext STATIC ${DRIVER_CONTEXT_SRC_LIST})

target_include_directories(drivercontext PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
        )

target_link_libraries(drivercontext
        yaml-cpp
        drivermodel
        driverutil
)
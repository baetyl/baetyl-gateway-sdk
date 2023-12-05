find_package(nlohmann_json REQUIRED)
if(NOT nlohmann_json_FOUND)
    message(FATAL_ERROR "nlohmann/json not found. Please install it. You can use the following commands on Linux:\n"
            "  sudo apt-get install nlohmann-json3-dev\n"
            "On macOS with brew:\n"
            "  brew install nlohmann-json")
else()
    message(STATUS "Found nlohmann/json version: ${nlohmann_json_VERSION}")
endif()

set(DRIVER_SRC_LIST
        custom.cpp
        device.cpp
        driver.cpp
        simulator.cpp
        worker.cpp
)

add_library(driver STATIC ${DRIVER_SRC_LIST})

target_include_directories(driver PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
        ${nlohmann_json_INCLUDE_DIRS}
)

target_link_libraries(driver
        model
        util
        nlohmann_json::nlohmann_json
)
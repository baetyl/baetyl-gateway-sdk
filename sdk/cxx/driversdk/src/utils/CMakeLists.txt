set(DRIVER_UTIL_SRC_LIST
        path.cpp
)

add_library(driverutil STATIC ${DRIVER_UTIL_SRC_LIST})

target_include_directories(driverutil PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
        )

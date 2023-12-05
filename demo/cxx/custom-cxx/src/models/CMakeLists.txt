set(MODEL_SRC_LIST
        deviceconfig.cpp
        driverconfig.cpp
        job.cpp
        point.cpp
        property.cpp
)

add_library(model STATIC ${MODEL_SRC_LIST})

target_include_directories(model PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
)

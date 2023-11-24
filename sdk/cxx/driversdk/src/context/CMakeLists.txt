set(DRIVER_CONTEXT_SRC_LIST
        context.cpp
)

add_library(drivercontext STATIC ${DRIVER_CONTEXT_SRC_LIST})

target_include_directories(drivercontext PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
        )

target_link_libraries(drivercontext
        ${YAML_CPP_LIBRARIES}
        drivermodel
        driverutil
)
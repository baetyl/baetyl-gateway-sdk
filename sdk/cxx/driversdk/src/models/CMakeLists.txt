set(DRIVER_MODEL_SRC_LIST
        accessconfig.cpp
        accesstemplate.cpp
        arraytype.cpp
        deviceinfo.cpp
        deviceproperty.cpp
        enumtype.cpp
        enumvalue.cpp
        event.cpp
        modelmapping.cpp
        objecttype.cpp
        propertyvisitor.cpp
        subdeviceyaml.cpp
        ../context/yamlparser.cpp
        ../context/yamlparser.h
)

add_library(drivermodel STATIC ${DRIVER_MODEL_SRC_LIST})

target_include_directories(drivermodel PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
)
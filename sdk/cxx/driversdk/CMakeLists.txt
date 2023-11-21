cmake_minimum_required(VERSION 3.13)
project(driversdk)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# set output path
set(OUTPUT_PATH "${CMAKE_SOURCE_DIR}/output-${CMAKE_SYSTEM_PROCESSOR}")
set(EXECUTABLE_OUTPUT_PATH "${OUTPUT_PATH}")
set(LIBRARY_OUTPUT_PATH "${OUTPUT_PATH}")

# set root path
set(ROOT_INCLUDE "${CMAKE_SOURCE_DIR}/include")
set(ROOT_DIR "${CMAKE_SOURCE_DIR}")

# set 3rd path
set(SUFFIX "${CMAKE_SYSTEM_PROCESSOR}")
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
set(THIRDPART_DIR "${CMAKE_SOURCE_DIR}/3rdparty")
set(3RD_FINDER "${THIRDPART_DIR}/3rdfinder")

# set proto path
set(PROTO_DIR "${CMAKE_SOURCE_DIR}/proto")

# include 3rd finder
include("${3RD_FINDER}/yaml-cpp.cmake")
include("${PROTO_DIR}/driver-proto.cmake")

# add subdirectory
add_subdirectory(${PROTO_DIR})
#add_subdirectory(${ROOT_DIR}/src/utils)
#add_subdirectory(${ROOT_DIR}/src/models)
#add_subdirectory(${ROOT_DIR}/src/context)
#add_subdirectory(${ROOT_DIR}/src/plugin)

if(CMAKE_BUILD_TYPE STREQUAL Debug)
    set(CMAKE_C_FLAGS "-DDEBUG ${CMAKE_C_FLAGS}")
    set(CMAKE_CXX_FLAGS "-DDEBUG ${CMAKE_CXX_FLAGS}")
    add_definitions("-DDEBUG")
endif()
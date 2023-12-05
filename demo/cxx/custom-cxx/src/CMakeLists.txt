set(SOURCE_LIST
        main.cpp
)

add_executable(main ${SOURCE_LIST})

target_include_directories(main PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
)

target_link_libraries(main
        driver
)
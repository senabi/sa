cmake_minimum_required(VERSION 3.5)
project(project_name VERSION 0.0.1)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
find_package(OpenCV REQUIRED)

#add_subdirectory(glad)

add_executable(elf main.cc)
target_link_libraries(
    elf PUBLIC
    ${OpenCV_LIBS}
)

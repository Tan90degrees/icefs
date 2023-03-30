cmake_minimum_required(VERSION 3.22)
project(IcefsClient)

if(DEBUG)
    message(STATUS "Build icefs debug")
    add_compile_definitions(ICEFS_DEBUG)
endif()

set(3RD_PARTY_DIR ../../../3rdparty)
set(YYJSON_DIR_NAME yyjson)
set(ICEFS_RPC_DIR_NAME rpc)
set(ICEFS_UTILS_DIR_NAME utils)
set(ICEFS_OPERS_DIR_NAME icefsoperators)

find_package(PkgConfig REQUIRED)

pkg_check_modules(LIBFUSE3 REQUIRED IMPORTED_TARGET fuse3)

set(ICEFS_UTILS_DIR ${PROJECT_SOURCE_DIR}/${ICEFS_UTILS_DIR_NAME})
set(ICEFS_OPERS_DIR ${PROJECT_SOURCE_DIR}/${ICEFS_OPERS_DIR_NAME})
set(ICEFS_RPC_SUB_PRJ icefsClientRpc)

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

add_subdirectory(${ICEFS_RPC_DIR_NAME})

include_directories(${3RD_PARTY_DIR}/${YYJSON_DIR_NAME}/src
    ${CMAKE_CURRENT_BINARY_DIR}/${ICEFS_RPC_DIR_NAME}
    ${LIBFUSE3_INCLUDE_DIRS}
    include)

aux_source_directory(${3RD_PARTY_DIR}/${YYJSON_DIR_NAME}/src yyjsonSrc)
aux_source_directory(${ICEFS_UTILS_DIR} icefsClientUtilsSrc)
aux_source_directory(${ICEFS_OPERS_DIR} icefsClientOperatorsSrc)

add_executable(${PROJECT_NAME} icefsClient.cpp
    ${yyjsonSrc}
    ${icefsClientUtilsSrc}
    ${icefsClientOperatorsSrc})

target_link_libraries(${PROJECT_NAME}
    PkgConfig::LIBFUSE3
    ${ICEFS_RPC_SUB_PRJ})

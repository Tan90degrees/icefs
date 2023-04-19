project(${ICEFS_THRIFT_SUB_PRJ})

find_package(Boost REQUIRED)
include_directories(SYSTEM "${Boost_INCLUDE_DIR}")
# find_package(OpenSSL REQUIRED)
# include_directories(SYSTEM "${OPENSSL_INCLUDE_DIR}")
find_package(ZLIB REQUIRED)
include_directories(SYSTEM ${ZLIB_INCLUDE_DIRS})
# find_package(Thrift REQUIRED)
# include_directories(SYSTEM ${THRIFT_INCLUDE_DIRS})

set(ICEFS_THRIFTS_DIR ../../../../../thrifts/lowlevel)

# get_filename_component(icefsThrift ${ICEFS_THRIFTS_DIR}/icefs.thrift ABSOLUTE)
get_filename_component(icefsServicesThrift ${ICEFS_THRIFTS_DIR}/icefsServices.thrift ABSOLUTE)

set(icefsThriftHdr ${CMAKE_CURRENT_BINARY_DIR}/IcefsThrift.h)
set(icefsThriftSrc ${CMAKE_CURRENT_BINARY_DIR}/IcefsThrift.cpp)
set(icefsServicesTypeHdr ${CMAKE_CURRENT_BINARY_DIR}/icefsServices_types.h)
set(icefsServicesTypeSrc ${CMAKE_CURRENT_BINARY_DIR}/icefsServices_types.cpp)

find_program(_THRIFT_PROGGRAM thrift REQUIRED)

add_custom_command(
    OUTPUT ${icefsThriftHdr} ${icefsThriftSrc} ${icefsServicesTypeHdr} ${icefsServicesTypeSrc}
    COMMAND ${_THRIFT_PROGGRAM}
    ARGS -gen "cpp:no_skeleton"
      -out "${CMAKE_CURRENT_BINARY_DIR}"
      "${ICEFS_THRIFTS_DIR}/icefsServices.thrift"
    WORKING_DIRECTORY ${PROJECT_SOURCE_DIR}
    DEPENDS "${icefsServicesThrift}")

add_library(${PROJECT_NAME} STATIC
    ${icefsThriftHdr}
    ${icefsThriftSrc}
    ${icefsServicesTypeHdr}
    ${icefsServicesTypeSrc})

target_link_libraries(${PROJECT_NAME}
    ${ZLIB_LIBRARIES}
    # ${THRIFT_LIBRARIES}
    thrift)

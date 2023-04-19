project(${ICEFS_GRPC_SUB_PRJ})

find_package(Protobuf CONFIG REQUIRED)
find_package(gRPC CONFIG REQUIRED)

message(STATUS "Now using gRPC@${gRPC_VERSION}")
message(STATUS "Now using protobuf@${Protobuf_VERSION}")

set(_GRPC_GRPCPP_REFLECTION gRPC::grpc++_reflection)
set(_GRPC_GRPCPP gRPC::grpc++)
set(_PROTOBUF_LIBPROTOBUF protobuf::libprotobuf)

set(ICEFS_PROTOS_DIR ../../../../../protos/lowlevel)

get_filename_component(fuseKernelProto ${ICEFS_PROTOS_DIR}/fuseKernel.proto ABSOLUTE)
get_filename_component(icefsProto ${ICEFS_PROTOS_DIR}/icefs.proto ABSOLUTE)
get_filename_component(icefsServicesProto ${ICEFS_PROTOS_DIR}/icefsServices.proto ABSOLUTE)

set(fuseKernelPBHdr ${CMAKE_CURRENT_BINARY_DIR}/fuseKernel.pb.h)
set(icefsPBHdr ${CMAKE_CURRENT_BINARY_DIR}/icefs.pb.h)
set(icefsServicesPBHdr ${CMAKE_CURRENT_BINARY_DIR}/icefsServices.pb.h)

set(fuseKernelRpcHdr ${CMAKE_CURRENT_BINARY_DIR}/fuseKernel.grpc.pb.h)
set(icefsRpcHdr ${CMAKE_CURRENT_BINARY_DIR}/icefs.grpc.pb.h)
set(icefsServicesRpcHdr ${CMAKE_CURRENT_BINARY_DIR}/icefsServices.grpc.pb.h)

set(fuseKernelPBSrc ${CMAKE_CURRENT_BINARY_DIR}/fuseKernel.pb.cc)
set(icefsPBSrc ${CMAKE_CURRENT_BINARY_DIR}/icefs.pb.cc)
set(icefsServicesPBSrc ${CMAKE_CURRENT_BINARY_DIR}/icefsServices.pb.cc)

set(fuseKernelRpcSrc ${CMAKE_CURRENT_BINARY_DIR}/fuseKernel.grpc.pb.cc)
set(icefsRpcSrc ${CMAKE_CURRENT_BINARY_DIR}/icefs.grpc.pb.cc)
set(icefsServicesRpcSrc ${CMAKE_CURRENT_BINARY_DIR}/icefsServices.grpc.pb.cc)

if(CMAKE_CROSSCOMPILING)
    find_program(_PROTOBUF_PROTOC protoc REQUIRED)
else()
    set(_PROTOBUF_PROTOC $<TARGET_FILE:protobuf::protoc>)
endif()

if(CMAKE_CROSSCOMPILING)
    find_program(_GRPC_GRPCPP_PLUGIN grpc_cpp_plugin REQUIRED)
else()
    set(_GRPC_GRPCPP_PLUGIN $<TARGET_FILE:gRPC::grpc_cpp_plugin>)
endif()

add_custom_command(
    OUTPUT ${fuseKernelPBHdr} ${icefsPBHdr} ${icefsServicesPBHdr} ${fuseKernelRpcHdr} ${icefsRpcHdr} ${icefsServicesRpcHdr} ${fuseKernelPBSrc} ${icefsPBSrc} ${icefsServicesPBSrc} ${fuseKernelRpcSrc} ${icefsRpcSrc} ${icefsServicesRpcSrc}
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS -I "${ICEFS_PROTOS_DIR}"
      --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
      --plugin=protoc-gen-grpc="${_GRPC_GRPCPP_PLUGIN}"
      --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
      "${ICEFS_PROTOS_DIR}/*.proto"
    WORKING_DIRECTORY ${PROJECT_SOURCE_DIR}
    DEPENDS "${fuseKernelProto}" "${icefsProto}" "${icefsServicesProto}")

add_library(${PROJECT_NAME} STATIC
    ${fuseKernelPBHdr}
    ${icefsPBHdr}
    ${icefsServicesPBHdr}
    ${fuseKernelRpcHdr}
    ${icefsRpcHdr}
    ${icefsServicesRpcHdr}
    ${fuseKernelPBSrc}
    ${icefsPBSrc}
    ${icefsServicesPBSrc}
    ${fuseKernelRpcSrc}
    ${icefsRpcSrc}
    ${icefsServicesRpcSrc})

target_link_libraries(${PROJECT_NAME}
    ${_GRPC_GRPCPP_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})

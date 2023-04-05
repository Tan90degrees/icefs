#!/bin/bash

PWD=`pwd`
LOW_LEVEL_PROTO_PATH=$PWD/../../../protos/lowlevel
LOW_LEVEL_RPC_PATH=$PWD/rpc

protoc -I $LOW_LEVEL_PROTO_PATH --grpc_out=$LOW_LEVEL_RPC_PATH --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` --cpp_out=$LOW_LEVEL_RPC_PATH $LOW_LEVEL_PROTO_PATH/*.proto

#!/bin/bash
###
 # @Author: Tan90degrees tangentninetydegrees@gmail.com
 # @Date: 2023-03-30 05:57:50
 # @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 # @LastEditTime: 2023-04-19 06:01:50
 # @FilePath: /icefs/src/lowlevel/client/protoGenCode.sh
 # @Description: 
### 

PWD=`pwd`
LOW_LEVEL_PROTO_PATH=$PWD/../../../protos/lowlevel
LOW_LEVEL_RPC_PATH=$PWD/rpc/grpc

protoc -I $LOW_LEVEL_PROTO_PATH --grpc_out=$LOW_LEVEL_RPC_PATH --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` --cpp_out=$LOW_LEVEL_RPC_PATH $LOW_LEVEL_PROTO_PATH/*.proto

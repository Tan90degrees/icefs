#!/bin/bash
###
 # @Author: Tan90degrees tangentninetydegrees@gmail.com
 # @Date: 2023-03-30 05:57:50
 # @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 # @LastEditTime: 2023-04-19 08:09:30
 # @FilePath: /icefs/src/lowlevel/server/protoGenCode.sh
 # @Description: 
### 

PWD=`pwd`
LOW_LEVEL_PROTO_PATH=$PWD/../../../protos/lowlevel
LOW_LEVEL_RPC_PATH=$PWD/icefsgrpc

protoc -I $LOW_LEVEL_PROTO_PATH --go-grpc_out=$LOW_LEVEL_RPC_PATH --go_out=$LOW_LEVEL_RPC_PATH $LOW_LEVEL_PROTO_PATH/*.proto

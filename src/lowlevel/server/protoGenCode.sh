#!/bin/bash

PWD=`pwd`
LOW_LEVEL_PROTO_PATH=$PWD/../../../protos/lowlevel
LOW_LEVEL_RPC_PATH=$PWD/icefsrpc

protoc -I $LOW_LEVEL_PROTO_PATH --go-grpc_out=$LOW_LEVEL_RPC_PATH --go_out=$LOW_LEVEL_RPC_PATH $LOW_LEVEL_PROTO_PATH/*.proto

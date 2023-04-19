#!/bin/bash
###
 # @Author: Tan90degrees tangentninetydegrees@gmail.com
 # @Date: 2023-04-19 05:53:19
 # @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 # @LastEditTime: 2023-04-19 08:24:28
 # @FilePath: /icefs/src/lowlevel/client/thriftGenCode.sh
 # @Description: 
### 

PWD=`pwd`
LOW_LEVEL_THRIFT_PATH=$PWD/../../../thrifts/lowlevel
LOW_LEVEL_THRIFT_RPC_PATH=$PWD/rpc/thrift

thrift -gen cpp:no_skeleton -out $LOW_LEVEL_THRIFT_RPC_PATH $LOW_LEVEL_THRIFT_PATH/icefsServices.thrift
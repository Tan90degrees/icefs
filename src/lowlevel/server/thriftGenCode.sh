#!/bin/bash
###
 # @Author: Tan90degrees tangentninetydegrees@gmail.com
 # @Date: 2023-04-17 13:03:24
 # @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 # @LastEditTime: 2023-04-19 08:24:19
 # @FilePath: /icefs/src/lowlevel/server/thriftGenCode.sh
 # @Description: 
### 

PWD=`pwd`
LOW_LEVEL_THRIFT_PATH=$PWD/../../../thrifts/lowlevel
LOW_LEVEL_THRIFT_RPC_PATH=$PWD

thrift -gen go:package=icefsthrift -out $LOW_LEVEL_THRIFT_RPC_PATH $LOW_LEVEL_THRIFT_PATH/icefsServices.thrift
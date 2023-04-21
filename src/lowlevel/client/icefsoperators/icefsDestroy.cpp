/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:41
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsDestroy.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsDestroy(void *userData) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsDestroyReq req;
      icefsgrpc::IcefsDestroyRes res;
      grpc::ClientContext ctx;

      req.set_host_name("Hello");
      req.set_info("I'm the client.");

      grpc::Status status = gRpcClient->DoIcefsDestroy(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        std::cout << res.info() << std::endl;
      } else {
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsDestroyReq req;
      icefsthrift::IcefsDestroyRes res;

      req.__set_host_name("Hello");
      req.__set_info("I'm the client.");

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsDestroy(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        std::cout << res.info << std::endl;
      } else {
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    default:
      break;
  }
}

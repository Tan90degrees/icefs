/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:49
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsInit.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsInit(void *userData, struct fuse_conn_info *conn) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsInitReq req;
      icefsgrpc::IcefsInitRes res;
      grpc::ClientContext ctx;

      if (conn->capable & FUSE_CAP_EXPORT_SUPPORT)
        conn->want |= FUSE_CAP_EXPORT_SUPPORT;

      if (this->clientConfig.cacheTimeout &&
          conn->capable & FUSE_CAP_WRITEBACK_CACHE) {
        conn->want |= FUSE_CAP_WRITEBACK_CACHE;
      }
      if (conn->capable & FUSE_CAP_FLOCK_LOCKS) {
        conn->want |= FUSE_CAP_FLOCK_LOCKS;
      }

      req.set_uuid(this->clientConfig.uuid);
      req.set_info("I'm the client.");
      req.set_want(conn->want);
      req.set_timeout(this->clientConfig.cacheTimeout);

      grpc::Status status = gRpcClient->DoIcefsInit(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        std::cout << res.info() << std::endl;
        conn->want = res.can();
      } else {
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsInitReq req;
      icefsthrift::IcefsInitRes res;

      if (conn->capable & FUSE_CAP_EXPORT_SUPPORT)
        conn->want |= FUSE_CAP_EXPORT_SUPPORT;

      if (this->clientConfig.cacheTimeout &&
          conn->capable & FUSE_CAP_WRITEBACK_CACHE) {
        conn->want |= FUSE_CAP_WRITEBACK_CACHE;
      }
      if (conn->capable & FUSE_CAP_FLOCK_LOCKS) {
        conn->want |= FUSE_CAP_FLOCK_LOCKS;
      }

      req.__set_uuid(this->clientConfig.uuid);
      req.__set_info("I'm the client.");
      req.__set_want(conn->want);
      req.__set_timeout(this->clientConfig.cacheTimeout);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsInit(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        std::cout << res.info << std::endl;
        conn->want = res.can;
      } else {
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    default:
      break;
  }
}
/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:44
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsForget.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsForget(fuse_req_t fuseReq, fuse_ino_t inode,
                                uint64_t nLookup) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsForgetReq req;
      icefsgrpc::IcefsForgetRes res;
      grpc::ClientContext ctx;

      req.set_inode(inode);
      req.set_nlookup(nLookup);

      grpc::Status status = gRpcClient->DoIcefsForget(&ctx, req, &res);
      if (!status.ok() || res.status()) {
        ICEFS_PR_ERR_STATUS;
      }
      fuse_reply_none(fuseReq);
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsForgetReq req;
      icefsthrift::IcefsForgetRes res;

      req.__set_inode(inode);
      req.__set_nlookup(nLookup);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsForget(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (res.status) {
        ICEFS_PR_ERR_STATUS;
      }
      fuse_reply_none(fuseReq);
      break;
    }

    default:
      fuse_reply_none(fuseReq);
      break;
  }
}
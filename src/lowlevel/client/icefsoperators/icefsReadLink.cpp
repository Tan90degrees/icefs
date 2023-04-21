/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:58
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsReadLink.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsReadLink(fuse_req_t fuseReq, fuse_ino_t inode) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsReadLinkReq req;
      icefsgrpc::IcefsReadLinkRes res;
      grpc::ClientContext ctx;

      req.set_inode(inode);

      grpc::Status status = gRpcClient->DoIcefsReadLink(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_readlink(fuseReq, res.path().c_str());
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsReadLinkReq req;
      icefsthrift::IcefsReadLinkRes res;

      req.__set_inode(inode);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsReadLink(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        fuse_reply_readlink(fuseReq, res.path.c_str());
      } else {
        fuse_reply_err(fuseReq, res.status);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}
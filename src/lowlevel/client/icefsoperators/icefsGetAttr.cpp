/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:47
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsGetAttr.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsGetAttr(fuse_req_t fuseReq, fuse_ino_t inode,
                                 struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsGetAttrReq req;
      icefsgrpc::IcefsGetAttrRes res;
      grpc::ClientContext ctx;
      struct stat attr;

      req.set_inode(inode);

      grpc::Status status = gRpcClient->DoIcefsGetAttr(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        IcefsGRpcFillAttrIn(&attr, res.stat());
        fuse_reply_attr(fuseReq, &attr, this->clientConfig.cacheTimeout);
      } else {
        ICEFS_PR_ERR_STATUS;
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsGetAttrReq req;
      icefsthrift::IcefsGetAttrRes res;
      struct stat attr;

      req.__set_inode(inode);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsGetAttr(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        IcefsThriftFillAttrIn(&attr, res.stat);
        fuse_reply_attr(fuseReq, &attr, this->clientConfig.cacheTimeout);
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

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:39
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsAccess.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsAccess(fuse_req_t fuseReq, fuse_ino_t inode,
                                int mask) {
  ICEFS_PR_FUNCTION;
  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsAccessReq req;
      icefsgrpc::IcefsAccessRes res;
      grpc::ClientContext ctx;

      req.set_inode(inode);
      req.set_mask(mask);

      grpc::Status status = gRpcClient->DoIcefsAccess(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_err(fuseReq, ICEFS_EOK);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsAccessReq req;
      icefsthrift::IcefsAccessRes res;

      req.__set_inode(inode);
      req.__set_mask(mask);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsAccess(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      fuse_reply_err(fuseReq, res.status);
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}

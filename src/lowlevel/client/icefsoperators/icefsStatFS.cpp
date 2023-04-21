/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:41:05
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsStatFS.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsStatFS(fuse_req_t fuseReq, fuse_ino_t inode) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsStatFSReq req;
      icefsgrpc::IcefsStatFSRes res;
      grpc::ClientContext ctx;

      struct statvfs statData;
      req.set_inode(inode);

      grpc::Status status = gRpcClient->DoIcefsStatFS(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        IcefsGRpcFillStatvfsIn(&statData, res.statvfs());
        fuse_reply_statfs(fuseReq, &statData);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsStatFSReq req;
      icefsthrift::IcefsStatFSRes res;
      struct statvfs statData;

      req.__set_inode(inode);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsStatFS(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        IcefsThriftFillStatvfsIn(&statData, res.statvfs);
        fuse_reply_statfs(fuseReq, &statData);
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

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:53
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsMknod.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsMknod(fuse_req_t fuseReq, fuse_ino_t parentInode,
                               const char *name, mode_t mode, dev_t rdev) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsMknodReq req;
      icefsgrpc::IcefsMknodRes res;
      grpc::ClientContext ctx;
      struct fuse_entry_param entry;

      req.set_parent_inode(parentInode);
      req.set_name(name);
      req.set_mode(mode);
      req.set_rdev(rdev);

      grpc::Status status = gRpcClient->DoIcefsMknod(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        IcefsGRpcFillFuseEntryParamIn(&entry, res.entry());
        fuse_reply_entry(fuseReq, &entry);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsMknodReq req;
      icefsthrift::IcefsMknodRes res;
      struct fuse_entry_param entry;

      req.__set_parent_inode(parentInode);
      req.__set_name(name);
      req.__set_mode(mode);
      req.__set_rdev(rdev);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsMknod(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        IcefsThriftFillFuseEntryParamIn(&entry, res.entry);
        fuse_reply_entry(fuseReq, &entry);
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

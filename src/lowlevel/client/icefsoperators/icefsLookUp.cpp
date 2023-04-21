/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:51
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsLookUp.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsLookUp(fuse_req_t fuseReq, fuse_ino_t parentInode,
                                const char *name) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsLookUpReq req;
      icefsgrpc::IcefsLookUpRes res;
      grpc::ClientContext ctx;
      struct fuse_entry_param entry;

      req.set_parent_inode(parentInode);
      req.set_name(name);

      grpc::Status status = gRpcClient->DoIcefsLookUp(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        IcefsGRpcFillFuseEntryParamIn(&entry, res.entry());
        fuse_reply_entry(fuseReq, &entry);
      } else if (res.status() == ENOENT) {
        entry.attr_timeout = this->clientConfig.cacheTimeout;
        entry.entry_timeout = this->clientConfig.cacheTimeout;
        entry.ino = 0;
        entry.attr.st_ino = 0;
        fuse_reply_entry(fuseReq, &entry);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsLookUpReq req;
      icefsthrift::IcefsLookUpRes res;
      struct fuse_entry_param entry;

      req.__set_parent_inode(parentInode);
      req.__set_name(name);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsLookUp(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        IcefsThriftFillFuseEntryParamIn(&entry, res.entry);
        fuse_reply_entry(fuseReq, &entry);
      } else if (res.status == ENOENT) {
        entry.attr_timeout = this->clientConfig.cacheTimeout;
        entry.entry_timeout = this->clientConfig.cacheTimeout;
        entry.ino = 0;
        entry.attr.st_ino = 0;
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
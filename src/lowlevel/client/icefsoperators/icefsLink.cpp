/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:50
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsLink.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsLink(fuse_req_t fuseReq, fuse_ino_t inode,
                              fuse_ino_t newParent, const char *newName) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsLinkReq req;
      icefsgrpc::IcefsLinkRes res;
      grpc::ClientContext ctx;
      struct fuse_entry_param entry;

      req.set_inode(inode);
      req.set_new_parent_inode(newParent);
      req.set_new_name(newName);

      grpc::Status status = gRpcClient->DoIcefsLink(&ctx, req, &res);
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
      icefsthrift::IcefsLinkReq req;
      icefsthrift::IcefsLinkRes res;
      struct fuse_entry_param entry;

      req.__set_inode(inode);
      req.__set_new_parent_inode(newParent);
      req.__set_new_name(newName);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsLink(res, req);
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
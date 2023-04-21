/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:41:06
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsSymLink.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsSymLink(fuse_req_t fuseReq, const char *link,
                                 fuse_ino_t parentInode, const char *name) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsSymLinkReq req;
      icefsgrpc::IcefsSymLinkRes res;
      grpc::ClientContext ctx;
      struct fuse_entry_param entry;

      req.set_link(link);
      req.set_parent_inode(parentInode);
      req.set_name(name);

      grpc::Status status = gRpcClient->DoIcefsSymLink(&ctx, req, &res);
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
      icefsthrift::IcefsSymLinkReq req;
      icefsthrift::IcefsSymLinkRes res;
      struct fuse_entry_param entry;

      req.__set_link(link);
      req.__set_parent_inode(parentInode);
      req.__set_name(name);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsSymLink(res, req);
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

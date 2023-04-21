/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:41:01
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsRename.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsRename(fuse_req_t fuseReq, fuse_ino_t parentInode,
                                const char *name, fuse_ino_t newParent,
                                const char *newName, unsigned int flags) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsRenameReq req;
      icefsgrpc::IcefsRenameRes res;
      grpc::ClientContext ctx;

      req.set_parent_inode(parentInode);
      req.set_name(name);
      req.set_new_parent_inode(newParent);
      req.set_new_name(newName);
      req.set_flags(flags);

      grpc::Status status = gRpcClient->DoIcefsRename(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_err(fuseReq, ICEFS_EOK);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsRenameReq req;
      icefsthrift::IcefsRenameRes res;

      req.__set_parent_inode(parentInode);
      req.__set_name(name);
      req.__set_new_parent_inode(newParent);
      req.__set_new_name(newName);
      req.__set_flags(flags);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsRename(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      fuse_reply_err(fuseReq, res.status);
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}
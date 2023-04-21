/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:40
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsCreate.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsCreate(fuse_req_t fuseReq, fuse_ino_t parentInode,
                                const char *name, mode_t mode,
                                struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsCreateReq req;
      icefsgrpc::IcefsCreateRes res;
      grpc::ClientContext ctx;
      struct fuse_entry_param entry;

      req.set_parent_inode(parentInode);
      req.set_name(name);
      req.set_mode(mode);
      req.set_flags(fi->flags);

      grpc::Status status = gRpcClient->DoIcefsCreate(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fi->fh = res.fh();
        if (this->clientConfig.cacheMode == ICEFS_CACHE_NEVER) {
          fi->direct_io = 1;
        } else {
          fi->keep_cache = 1;
        }
        IcefsGRpcFillFuseEntryParamIn(&entry, res.entry());
        fuse_reply_create(fuseReq, &entry, fi);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsCreateReq req;
      icefsthrift::IcefsCreateRes res;
      struct fuse_entry_param entry;

      req.__set_parent_inode(parentInode);
      req.__set_name(name);
      req.__set_mode(mode);
      req.__set_flags(fi->flags);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsCreate(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        fi->fh = res.fh;
        if (this->clientConfig.cacheMode == ICEFS_CACHE_NEVER) {
          fi->direct_io = 1;
        } else {
          fi->keep_cache = 1;
        }
        IcefsThriftFillFuseEntryParamIn(&entry, res.entry);
        fuse_reply_create(fuseReq, &entry, fi);
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

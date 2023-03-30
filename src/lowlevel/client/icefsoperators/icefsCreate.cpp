/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:22:54
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
  IcefsCreateReq req;
  IcefsCreateRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  FuseFileInfo *fileInfo = new FuseFileInfo();
  struct fuse_entry_param entry;
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_parent_inode(parentInode);
  req.set_name(name);
  req.set_mode(mode);
  IcefsFillFuseFileInfoOut(fileInfo, fi);
  req.set_allocated_file_info(fileInfo);

  grpc::Status status = stub_->DoIcefsCreate(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fi->fh = res.fh();
    if (this->config.cacheMode == ICEFS_CACHE_NEVER) {
      fi->direct_io = 1;
    } else {
      fi->keep_cache = 1;
    }
    IcefsFillFuseEntryParamIn(&entry, res.entry());
    fuse_reply_create(fuseReq, &entry, fi);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}

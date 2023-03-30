/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:24:58
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
  IcefsRenameReq req;
  IcefsRenameRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_parent_inode(parentInode);
  req.set_name(name);
  req.set_new_parent_inode(newParent);
  req.set_new_name(newName);
  req.set_flags(flags);

  grpc::Status status = stub_->DoIcefsRename(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_err(fuseReq, ICEFS_EOK);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
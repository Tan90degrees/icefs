/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:24:35
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
  IcefsMknodReq req;
  IcefsMknodRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  struct fuse_entry_param entry;
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_parent_inode(parentInode);
  req.set_name(name);
  req.set_mode(mode);
  req.set_rdev(rdev);

  grpc::Status status = stub_->DoIcefsMknod(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    IcefsFillFuseEntryParamIn(&entry, res.entry());
    fuse_reply_entry(fuseReq, &entry);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
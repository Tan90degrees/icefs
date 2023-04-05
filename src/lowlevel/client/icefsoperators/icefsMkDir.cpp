/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:37
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsMkDir.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsMkDir(fuse_req_t fuseReq, fuse_ino_t parentInode,
                               const char *name, mode_t mode) {
  IcefsMkDirReq req;
  IcefsMkDirRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  struct fuse_entry_param entry;
  req.set_parent_inode(parentInode);
  req.set_name(name);
  req.set_mode(mode);

  grpc::Status status = stub_->DoIcefsMkDir(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    IcefsFillFuseEntryParamIn(&entry, res.entry());
    fuse_reply_entry(fuseReq, &entry);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
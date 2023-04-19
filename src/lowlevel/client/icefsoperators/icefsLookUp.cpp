/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 06:08:00
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
  IcefsLookUpReq req;
  IcefsLookUpRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  struct fuse_entry_param entry;
  req.set_parent_inode(parentInode);
  req.set_name(name);

  grpc::Status status = stub_->DoIcefsLookUp(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    IcefsFillFuseEntryParamIn(&entry, res.entry());
    fuse_reply_entry(fuseReq, &entry);
  } else if (res.status() == ENOENT) {
    entry.attr_timeout = this->config.cacheTimeout;
    entry.entry_timeout = this->config.cacheTimeout;
    entry.ino = 0;
    entry.attr.st_ino = 0;
    fuse_reply_entry(fuseReq, &entry);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:34
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
  IcefsLinkReq req;
  IcefsLinkRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  struct fuse_entry_param entry;
  req.set_inode(inode);
  req.set_new_parent_inode(newParent);
  req.set_new_name(newName);

  grpc::Status status = stub_->DoIcefsLink(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    IcefsFillFuseEntryParamIn(&entry, res.entry());
    fuse_reply_entry(fuseReq, &entry);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
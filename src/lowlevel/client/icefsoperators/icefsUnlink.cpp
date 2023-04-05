/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:55
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsUnlink.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsUnlink(fuse_req_t fuseReq, fuse_ino_t parentInode,
                                const char *name) {
  IcefsUnlinkReq req;
  IcefsUnlinkRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_parent_inode(parentInode);
  req.set_name(name);

  grpc::Status status = stub_->DoIcefsUnlink(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_err(fuseReq, ICEFS_EOK);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
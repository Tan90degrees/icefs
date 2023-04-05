/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:53
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsStatFS.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsStatFS(fuse_req_t fuseReq, fuse_ino_t inode) {
  IcefsStatFSReq req;
  IcefsStatFSRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  struct statvfs statData;
  req.set_inode(inode);

  grpc::Status status = stub_->DoIcefsStatFS(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    IcefsFillStatvfsIn(&statData, res.statvfs());
    fuse_reply_statfs(fuseReq, &statData);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}

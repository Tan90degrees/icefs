/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:24
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsFlock.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsFlock(fuse_req_t fuseReq, fuse_ino_t inode,
                               struct fuse_file_info *fi, int op) {
  IcefsFlockReq req;
  IcefsFlockRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_fh(fi->fh);
  req.set_op(op);

  grpc::Status status = stub_->DoIcefsFlock(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_err(fuseReq, ICEFS_EOK);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
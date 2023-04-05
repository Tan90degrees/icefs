/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:36
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsLseek.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsLseek(fuse_req_t fuseReq, fuse_ino_t inode,
                               off_t offset, int whence,
                               struct fuse_file_info *fi) {
  IcefsLseekReq req;
  IcefsLseekRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_offset(offset);
  req.set_whence(whence);
  req.set_fh(fi->fh);

  grpc::Status status = stub_->DoIcefsLseek(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_lseek(fuseReq, res.offset());
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
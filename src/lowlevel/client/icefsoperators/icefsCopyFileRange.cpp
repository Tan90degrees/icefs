/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:21
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsCopyFileRange.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsCopyFileRange(fuse_req_t fuseReq, fuse_ino_t inodeIn,
                                       off_t offsetIn,
                                       struct fuse_file_info *fiIn,
                                       fuse_ino_t inodeOut, off_t offsetOut,
                                       struct fuse_file_info *fiOut, size_t len,
                                       int flags) {
  IcefsCopyFileRangeReq req;
  IcefsCopyFileRangeRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseFileInfo *fileInfoIn = new FuseFileInfo();
  FuseFileInfo *fileInfoOut = new FuseFileInfo();
  req.set_offset_in(offsetIn);
  req.set_fh_in(fiIn->fh);
  req.set_offset_out(offsetOut);
  req.set_fh_out(fiOut->fh);
  req.set_len(len);
  req.set_flags(flags);

  grpc::Status status = stub_->DoIcefsCopyFileRange(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_write(fuseReq, res.size());
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
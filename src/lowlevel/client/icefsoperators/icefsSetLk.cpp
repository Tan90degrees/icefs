/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:25:06
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsSetLk.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsSetLk(fuse_req_t fuseReq, fuse_ino_t inode,
                               struct fuse_file_info *fi, struct flock *lock,
                               int sleep) {
  //   IcefsSetLkReq req;
  //   IcefsSetLkRes res;
  //   grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  //   FuseReq *fuseReqToSend = new FuseReq();
  //   FuseCtx *fuseCtx = new FuseCtx();
  //   IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  //   req.set_allocated_req(fuseReqToSend);

  //   grpc::Status status = stub_->DoIcefsSetLk(&ctx, req, &res);
  //   if (status.ok() && !res.status()) {
  //   } else {
  //     fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  //     ICEFS_PR_ERR_STATUS;
  //   }
  fuse_reply_err(fuseReq, ENOTSUP);
}
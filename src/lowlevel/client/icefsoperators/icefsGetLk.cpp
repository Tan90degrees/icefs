/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:23:24
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsGetLk.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsGetLk(fuse_req_t fuseReq, fuse_ino_t inode,
                               struct fuse_file_info *fi, struct flock *lock) {
  // IcefsGetLkReq req;
  // IcefsGetLkRes res;
  // grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  // FuseReq *fuseReqToSend = new FuseReq();
  // FuseCtx *fuseCtx = new FuseCtx();
  // FuseFileInfo *fileInfo = new FuseFileInfo();
  // flockStruct flock;
  // fuse_reply_err(fuseReq, ENOTSUP);
  // IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  // req.set_allocated_req(fuseReqToSend);
  // req.set_inode(inode);
  // IcefsFillFuseFileInfoOut(fileInfo, fi);
  // req.set_allocated_file_info(fileInfo);
  // IcefsFillFlockStructIn(flock, lock);
  // req.set_allocated_lock(&flock);

  // grpc::Status status = stub_->DoIcefsGetLk(&ctx, req, &res);
  // if (status.ok() && !res.status()) {
  // } else {
  //   fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  //   ICEFS_PR_ERR_STATUS;
  // }
  fuse_reply_err(fuseReq, ENOTSUP);
}

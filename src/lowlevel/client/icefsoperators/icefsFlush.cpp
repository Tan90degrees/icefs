/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:23:03
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsFlush.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsFlush(fuse_req_t fuseReq, fuse_ino_t inode,
                               struct fuse_file_info *fi) {
  IcefsFlushReq req;
  IcefsFlushRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  FuseFileInfo *fileInfo = new FuseFileInfo();
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_inode(inode);
  IcefsFillFuseFileInfoOut(fileInfo, fi);
  req.set_allocated_file_info(fileInfo);

  grpc::Status status = stub_->DoIcefsFlush(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_err(fuseReq, ICEFS_EOK);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}

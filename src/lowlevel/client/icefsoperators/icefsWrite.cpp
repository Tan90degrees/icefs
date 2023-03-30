/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:25:22
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsWrite.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsWrite(fuse_req_t fuseReq, fuse_ino_t inode,
                               const char *buf, size_t size, off_t offset,
                               struct fuse_file_info *fi) {
  IcefsWriteReq req;
  IcefsWriteRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  FuseFileInfo *fileInfo = new FuseFileInfo();
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_inode(inode);
  // std::string dataCopy(buf, size);
  // req.set_allocated_buf(&dataCopy);
  req.set_buf(buf);
  req.set_size(size);
  req.set_offset(offset);
  IcefsFillFuseFileInfoOut(fileInfo, fi);
  req.set_allocated_file_info(fileInfo);

  grpc::Status status = stub_->DoIcefsWrite(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_write(fuseReq, res.size());
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}

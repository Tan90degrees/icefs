/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:33
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsIoctl.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsIoctl(fuse_req_t fuseReq, fuse_ino_t inode,
                               unsigned int cmd, void *arg,
                               struct fuse_file_info *fi, unsigned flags,
                               const void *inBuf, size_t inBufSize,
                               size_t outBufSize) {
  //   IcefsIoctlReq req;
  //   IcefsIoctlRes res;
  //   grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  //   FuseReq *fuseReqToSend = new FuseReq();
  //   FuseCtx *fuseCtx = new FuseCtx();
  //   FuseFileInfo *fileInfo = new FuseFileInfo();
  //   IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  //   req.set_allocated_req(fuseReqToSend);
  //   req.set_inode(inode);
  //   req.set_cmd(cmd);
  //   req.set_arg(arg);
  //   IcefsFillFuseFileInfoOut(fileInfo, fi);
  //   req.set_allocated_file_info(fileInfo);
  //   req.set_flags(flags);
  //   req.set_in_buf_size(inBufSize);
  //   req.set_out_buf_size(outBufSize);

  //   grpc::Status status = stub_->DoIcefsIoctl(&ctx, req, &res);
  //   if (status.ok() && !res.status()) {
  //   } else {
  //     fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  //     ICEFS_PR_ERR_STATUS;
  //   }
  fuse_reply_err(fuseReq, ENOTSUP);
}
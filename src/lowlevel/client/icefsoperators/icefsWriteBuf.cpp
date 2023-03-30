/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:25:25
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsWriteBuf.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsWriteBuf(fuse_req_t fuseReq, fuse_ino_t inode,
                                  struct fuse_bufvec *bufVector, off_t offset,
                                  struct fuse_file_info *fi) {
  // IcefsWriteBufReq req;
  // IcefsWriteBufRes res;
  // grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  // FuseReq *fuseReqToSend = new FuseReq();
  // FuseCtx *fuseCtx = new FuseCtx();
  // IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  // req.set_allocated_req(fuseReqToSend);
  // req.set_inode(inode);
  // size_t size = 0;
  // ioVector * buf = nullptr;
  // for (size_t i = 0; i < bufVector->count; ++i) {
  //   buf = req.add_buf();
  //   buf->set_size(bufVector->buf[i].size);
  //   buf->set_data(bufVector->buf[i].mem);
  // }

  // grpc::Status status = stub_->DoIcefsWriteBuf(&ctx, req, &res);
  // if (status.ok() && !res.status()) {
  // } else {
  //   fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  //   ICEFS_PR_ERR_STATUS;
  // }
  fuse_reply_err(fuseReq, ENOTSUP);
}
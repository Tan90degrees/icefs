/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:49
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsRetrieveReply.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsRetrieveReply(fuse_req_t fuseReq, void *cookie,
                                       fuse_ino_t inode, off_t offset,
                                       struct fuse_bufvec *bufv) {
  //   IcefsRetrieveReplyReq req;
  //   IcefsRetrieveReplyRes res;
  //   grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  //   FuseReq *fuseReqToSend = new FuseReq();
  //   FuseCtx *fuseCtx = new FuseCtx();
  //   IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  //   req.set_allocated_req(fuseReqToSend);
  //   req.set_cookie(cookie);
  //   req.set_inode(inode);
  //   req.set_offset(offset);

  //   grpc::Status status = stub_->DoIcefsRetrieveReply(&ctx, req, &res);
  //   if (status.ok() && !res.status()) {
  //     fuse_reply_none(fuseReq);
  //   } else {
  //     fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  //     ICEFS_PR_ERR_STATUS;
  //   }
  fuse_reply_err(fuseReq, ENOTSUP);
}
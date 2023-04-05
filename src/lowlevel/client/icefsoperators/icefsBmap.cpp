/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:20:11
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:20
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsBmap.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsBmap(fuse_req_t fuseReq, fuse_ino_t inode,
                              size_t blockSize, uint64_t index) {
  //   IcefsBmapReq req;
  //   IcefsBmapRes res;
  //   grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  //   FuseReq *fuseReqToSend = new FuseReq();
  //   FuseCtx *fuseCtx = new FuseCtx();
  //   IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  //   req.set_allocated_req(fuseReqToSend);

  //   grpc::Status status = stub_->DoIcefsBmap(&ctx, req, &res);
  //   if (status.ok() && !res.status()) {
  //   } else {
  //     fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  //     ICEFS_PR_ERR_STATUS;
  //   }
  fuse_reply_err(fuseReq, ENOTSUP);
}
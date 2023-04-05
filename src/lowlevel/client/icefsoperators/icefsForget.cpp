/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:26
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsForget.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsForget(fuse_req_t fuseReq, fuse_ino_t inode,
                                uint64_t nLookup) {
  IcefsForgetReq req;
  IcefsForgetRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_inode(inode);
  req.set_nlookup(nLookup);

  grpc::Status status = stub_->DoIcefsForget(&ctx, req, &res);
  if (!status.ok() || res.status()) {
    ICEFS_PR_ERR_STATUS;
  }
  fuse_reply_none(fuseReq);
}
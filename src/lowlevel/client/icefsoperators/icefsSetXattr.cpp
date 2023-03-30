/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:25:14
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsSetXattr.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsSetXattr(fuse_req_t fuseReq, fuse_ino_t inode,
                                  const char *name, const char *value,
                                  size_t size, int flags) {
  IcefsSetXattrReq req;
  IcefsSetXattrRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_inode(inode);
  req.set_name(name);
  req.set_value(value);
  req.set_size(size);
  req.set_flags(flags);

  grpc::Status status = stub_->DoIcefsSetXattr(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fuse_reply_err(fuseReq, ICEFS_EOK);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}

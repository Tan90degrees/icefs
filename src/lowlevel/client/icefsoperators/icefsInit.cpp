/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 13:43:31
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsInit.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsInit(void *userData, struct fuse_conn_info *conn) {
  IcefsInitReq req;
  IcefsInitRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  if (conn->capable & FUSE_CAP_EXPORT_SUPPORT)
    conn->want |= FUSE_CAP_EXPORT_SUPPORT;

  if (this->config.cacheTimeout && conn->capable & FUSE_CAP_WRITEBACK_CACHE) {
    conn->want |= FUSE_CAP_WRITEBACK_CACHE;
  }
  if (conn->capable & FUSE_CAP_FLOCK_LOCKS) {
    conn->want |= FUSE_CAP_FLOCK_LOCKS;
  }

  req.set_uuid(this->config.uuid);
  req.set_info("I'm the client.");
  req.set_want(conn->want);
  req.set_timeout(this->config.cacheTimeout);
  grpc::Status status = stub_->DoIcefsInit(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    std::cout << res.info() << std::endl;
    conn->want = res.can();
  } else {
    ICEFS_PR_ERR_STATUS;
  }
}
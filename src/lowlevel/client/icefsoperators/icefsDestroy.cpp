/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:23
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsDestroy.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsDestroy(void *userData) {
  IcefsDestroyReq req;
  IcefsDestroyRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_host_name("Hello");
  req.set_info("I'm the client.");

  grpc::Status status = stub_->DoIcefsDestroy(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    std::cout << res.info() << std::endl;
  } else {
    ICEFS_PR_ERR_STATUS;
  }
}

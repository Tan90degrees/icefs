/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:27
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsForgetMulti.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsForgetMulti(fuse_req_t fuseReq, size_t count,
                                     struct fuse_forget_data *forgets) {
  IcefsForgetMultiReq req;
  IcefsForgetMultiRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  IcefsForgetMultiReq_forget_data *forgetOne;
  req.set_count(count);
  for (size_t i = 0; i < count; ++i) {
    forgetOne = req.add_to_forget();
    forgetOne->set_inode(forgets[i].ino);
    forgetOne->set_nlookup(forgets[i].nlookup);
  }

  grpc::Status status = stub_->DoIcefsForgetMulti(&ctx, req, &res);
  if (!status.ok() || res.status()) {
    ICEFS_PR_ERR_STATUS;
  }
  fuse_reply_none(fuseReq);
}
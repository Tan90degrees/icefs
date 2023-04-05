/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:31
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsGetXattr.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsGetXattr(fuse_req_t fuseReq, fuse_ino_t inode,
                                  const char *name, size_t size) {
  IcefsGetXattrReq req;
  IcefsGetXattrRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_inode(inode);
  req.set_name(name);
  req.set_size(size);

  grpc::Status status = stub_->DoIcefsGetXattr(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    if (size) {
      int32_t ret = fuse_reply_buf(fuseReq, res.value().c_str(), res.size());
      if (ret != ICEFS_EOK) {
        fuse_reply_err(fuseReq, ret);
        std::cout << "ERR:" << __FUNCTION__ << "fuse_reply_buf" << ret
                  << std::endl;
      }
    } else {
      fuse_reply_xattr(fuseReq, res.size());
    }
  } else if (res.status() == ENODATA) {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
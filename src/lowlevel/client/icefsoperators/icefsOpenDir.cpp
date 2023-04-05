/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:40
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsOpenDir.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsOpenDir(fuse_req_t fuseReq, fuse_ino_t inode,
                                 struct fuse_file_info *fi) {
  IcefsOpenDirReq req;
  IcefsOpenDirRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_inode(inode);

  grpc::Status status = stub_->DoIcefsOpenDir(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fi->fh = res.fh();
    if (this->config.cacheMode) {
      fi->keep_cache = 1;
      fi->cache_readdir = 1;
    }
    fuse_reply_open(fuseReq, fi);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
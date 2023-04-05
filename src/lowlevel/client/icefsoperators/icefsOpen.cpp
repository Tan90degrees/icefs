/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:57:32
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsOpen.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsOpen(fuse_req_t fuseReq, fuse_ino_t inode,
                              struct fuse_file_info *fi) {
  IcefsOpenReq req;
  IcefsOpenRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_inode(inode);
  if (this->config.cacheMode && (fi->flags & O_ACCMODE) == O_WRONLY) {
    fi->flags &= ~O_ACCMODE;
    fi->flags |= O_RDWR;
  }

  if (this->config.cacheMode && fi->flags & O_APPEND) {
    fi->flags &= ~O_APPEND;
  }
  req.set_flags(fi->flags);

  grpc::Status status = stub_->DoIcefsOpen(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    fi->fh = res.fh();
    if (this->config.cacheMode) {
      fi->keep_cache = 1;
    } else if ((fi->flags & O_ACCMODE) == O_RDONLY) {
      fi->noflush = 1;
      fi->direct_io = 1;
    } else {
      fi->direct_io = 1;
    }
    fuse_reply_open(fuseReq, fi);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
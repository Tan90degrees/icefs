/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-20 09:40:11
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsGetLk.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsGetLk(fuse_req_t fuseReq, fuse_ino_t inode,
                               struct fuse_file_info *fi, struct flock *lock) {
  ICEFS_PR_FUNCTION;

  fuse_reply_err(fuseReq, ENOTSUP);
}

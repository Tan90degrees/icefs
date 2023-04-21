/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-20 09:52:17
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsIoctl.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsIoctl(fuse_req_t fuseReq, fuse_ino_t inode,
                               unsigned int cmd, void *arg,
                               struct fuse_file_info *fi, unsigned flags,
                               const void *inBuf, size_t inBufSize,
                               size_t outBufSize) {
  ICEFS_PR_FUNCTION;

  fuse_reply_err(fuseReq, ENOTSUP);
}
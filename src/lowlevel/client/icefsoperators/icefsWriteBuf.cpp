/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-20 11:26:46
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsWriteBuf.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsWriteBuf(fuse_req_t fuseReq, fuse_ino_t inode,
                                  struct fuse_bufvec *bufVector, off_t offset,
                                  struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  fuse_reply_err(fuseReq, ENOTSUP);
}
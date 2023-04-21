/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:20:11
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-19 14:59:52
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsBmap.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsBmap(fuse_req_t fuseReq, fuse_ino_t inode,
                              size_t blockSize, uint64_t index) {
  ICEFS_PR_FUNCTION;

  fuse_reply_err(fuseReq, ENOTSUP);
}
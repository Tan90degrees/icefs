/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-20 10:59:48
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsRetrieveReply.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsRetrieveReply(fuse_req_t fuseReq, void *cookie,
                                       fuse_ino_t inode, off_t offset,
                                       struct fuse_bufvec *bufv) {
  ICEFS_PR_FUNCTION;

  fuse_reply_err(fuseReq, ENOTSUP);
}
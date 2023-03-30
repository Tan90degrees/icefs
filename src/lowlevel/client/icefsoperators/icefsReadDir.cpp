/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:24:45
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsReadDir.cpp
 * @Description: 
 * 
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsReadDir(fuse_req_t fuseReq, fuse_ino_t inode,
                                 size_t size, off_t offset,
                                 struct fuse_file_info *fi) {
  IcefsReadDirReq req;
  IcefsReadDirRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  FuseFileInfo *fileInfo = new FuseFileInfo();
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_inode(inode);
  req.set_size(size);
  req.set_offset(offset);
  IcefsFillFuseFileInfoOut(fileInfo, fi);
  req.set_allocated_file_info(fileInfo);

  grpc::Status status = stub_->DoIcefsReadDir(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    size_t remain = size;
    char *buf = new (std::nothrow) char[size];
    char *p = buf;
    size_t entrySize;
    int dataSize = res.data_size();
    struct fuse_entry_param entry = {0};
    for (int i = 0; i < dataSize; ++i) {
      entry.attr.st_ino = res.data(i).ino();
      entry.attr.st_mode = res.data(i).type() << 12;
      entrySize =
          fuse_add_direntry(fuseReq, p, remain, res.data(i).name().c_str(),
                            &entry.attr, res.data(i).off());
      if (entrySize > remain) {
        break;
      }
      p += entrySize;
      remain -= entrySize;
    }
    int32_t ret = fuse_reply_buf(fuseReq, buf, size - remain);
    if (ret != ICEFS_EOK) {
      fuse_reply_err(fuseReq, ret);
      std::cout << "ERR:" << __FUNCTION__ << "fuse_reply_buf" << ret
                << std::endl;
    }
    delete[] buf;
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
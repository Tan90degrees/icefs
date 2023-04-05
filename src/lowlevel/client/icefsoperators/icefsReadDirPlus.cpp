/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:53:44
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsReadDirPlus.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsReadDirPlus(fuse_req_t fuseReq, fuse_ino_t inode,
                                     size_t size, off_t offset,
                                     struct fuse_file_info *fi) {
  IcefsReadDirPlusReq req;
  IcefsReadDirPlusRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  req.set_inode(inode);
  req.set_offset(offset);
  req.set_fh(fi->fh);

  grpc::Status status = stub_->DoIcefsReadDirPlus(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    size_t remain = size;
    char *buf = new (std::nothrow) char[size];
    char *p = buf;
    size_t entrySize;
    int dataSize = res.data_size();
    struct fuse_entry_param entry = {0};
    for (int i = 0; i < dataSize; ++i) {
      IcefsFillFuseEntryParamIn(&entry, res.data(i).entry());
      entrySize = fuse_add_direntry_plus(fuseReq, p, remain,
                                         res.data(i).dir_entry().name().c_str(),
                                         &entry, res.data(i).dir_entry().off());
      if (entrySize > remain) {
        // TODO: forget one
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
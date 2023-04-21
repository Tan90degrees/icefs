/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:57
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
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsReadDirReq req;
      icefsgrpc::IcefsReadDirRes res;
      grpc::ClientContext ctx;

      req.set_offset(offset);
      req.set_fh(fi->fh);

      grpc::Status status = gRpcClient->DoIcefsReadDir(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        size_t remain = size;
        char *buf = new (std::nothrow) char[size];
        if (buf == nullptr) {
          fuse_reply_err(fuseReq, ENOMEM);
          return;
        }
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
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsReadDirReq req;
      icefsthrift::IcefsReadDirRes res;

      req.__set_offset(offset);
      req.__set_fh(fi->fh);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsReadDir(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        size_t remain = size;
        char *buf = new (std::nothrow) char[size];
        if (buf == nullptr) {
          fuse_reply_err(fuseReq, ENOMEM);
          return;
        }
        char *p = buf;
        size_t entrySize;
        int dataSize = res.data.size();
        struct fuse_entry_param entry = {0};
        for (int i = 0; i < dataSize; ++i) {
          entry.attr.st_ino = res.data[i].ino;
          entry.attr.st_mode = res.data[i].type << 12;
          entrySize =
              fuse_add_direntry(fuseReq, p, remain, res.data[i].name.c_str(),
                                &entry.attr, res.data[i].off);
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
        fuse_reply_err(fuseReq, res.status);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}
/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:52
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsLseek.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsLseek(fuse_req_t fuseReq, fuse_ino_t inode,
                               off_t offset, int whence,
                               struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsLseekReq req;
      icefsgrpc::IcefsLseekRes res;
      grpc::ClientContext ctx;

      req.set_offset(offset);
      req.set_whence(whence);
      req.set_fh(fi->fh);

      grpc::Status status = gRpcClient->DoIcefsLseek(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_lseek(fuseReq, res.offset());
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsLseekReq req;
      icefsthrift::IcefsLseekRes res;

      req.__set_offset(offset);
      req.__set_whence(whence);
      req.__set_fh(fi->fh);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsLseek(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        fuse_reply_lseek(fuseReq, res.offset);
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
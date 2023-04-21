/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:42
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsFallocate.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsFallocate(fuse_req_t fuseReq, fuse_ino_t inode,
                                   int mode, off_t offset, off_t length,
                                   struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsFallocateReq req;
      icefsgrpc::IcefsFallocateRes res;
      grpc::ClientContext ctx;

      req.set_mode(mode);
      req.set_offset(offset);
      req.set_length(length);
      req.set_fh(fi->fh);

      grpc::Status status = gRpcClient->DoIcefsFallocate(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_err(fuseReq, ICEFS_EOK);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsFallocateReq req;
      icefsthrift::IcefsFallocateRes res;

      req.__set_mode(mode);
      req.__set_offset(offset);
      req.__set_length(length);
      req.__set_fh(fi->fh);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsFallocate(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      fuse_reply_err(fuseReq, res.status);
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}
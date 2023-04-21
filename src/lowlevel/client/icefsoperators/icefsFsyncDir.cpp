/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:46
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsFsyncDir.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsFsyncDir(fuse_req_t fuseReq, fuse_ino_t inode,
                                  int dataSync, struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsFsyncDirReq req;
      icefsgrpc::IcefsFsyncDirRes res;
      grpc::ClientContext ctx;

      req.set_data_sync(dataSync);
      req.set_fh(fi->fh);

      grpc::Status status = gRpcClient->DoIcefsFsyncDir(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_err(fuseReq, ICEFS_EOK);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsFsyncDirReq req;
      icefsthrift::IcefsFsyncDirRes res;

      req.__set_data_sync(dataSync);
      req.__set_fh(fi->fh);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsFsyncDir(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      fuse_reply_err(fuseReq, res.status);
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}
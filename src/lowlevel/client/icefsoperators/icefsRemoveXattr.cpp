/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:41:01
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsRemoveXattr.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsRemoveXattr(fuse_req_t fuseReq, fuse_ino_t inode,
                                     const char *name) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsRemoveXattrReq req;
      icefsgrpc::IcefsRemoveXattrRes res;
      grpc::ClientContext ctx;

      req.set_inode(inode);
      req.set_name(name);

      grpc::Status status = gRpcClient->DoIcefsRemoveXattr(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_err(fuseReq, ICEFS_EOK);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsRemoveXattrReq req;
      icefsthrift::IcefsRemoveXattrRes res;

      req.__set_inode(inode);
      req.__set_name(name);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsRemoveXattr(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      fuse_reply_err(fuseReq, res.status);
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}

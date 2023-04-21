/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:41:04
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsSetXattr.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsSetXattr(fuse_req_t fuseReq, fuse_ino_t inode,
                                  const char *name, const char *value,
                                  size_t size, int flags) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsSetXattrReq req;
      icefsgrpc::IcefsSetXattrRes res;
      grpc::ClientContext ctx;

      req.set_inode(inode);
      req.set_name(name);
      req.set_value(value);
      req.set_flags(flags);

      grpc::Status status = gRpcClient->DoIcefsSetXattr(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        fuse_reply_err(fuseReq, ICEFS_EOK);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsSetXattrReq req;
      icefsthrift::IcefsSetXattrRes res;

      req.__set_inode(inode);
      req.__set_name(name);
      req.__set_value(value);
      req.__set_flags(flags);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsSetXattr(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      fuse_reply_err(fuseReq, res.status);
      break;
    }

    default:
      fuse_reply_err(fuseReq, EIO);
      break;
  }
}

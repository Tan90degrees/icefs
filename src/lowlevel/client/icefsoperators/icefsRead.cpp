/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:56
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsRead.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsRead(fuse_req_t fuseReq, fuse_ino_t inode, size_t size,
                              off_t offset, struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsReadReq req;
      icefsgrpc::IcefsReadRes res;
      grpc::ClientContext ctx;

      req.set_size(size);
      req.set_offset(offset);
      req.set_fh(fi->fh);

      grpc::Status status = gRpcClient->DoIcefsRead(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        int32_t ret = fuse_reply_buf(fuseReq, res.data().c_str(), res.size());
        if (ret != ICEFS_EOK) {
          fuse_reply_err(fuseReq, ret);
          std::cout << "ERR:" << __FUNCTION__ << "fuse_reply_buf" << ret
                    << std::endl;
        }
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsReadReq req;
      icefsthrift::IcefsReadRes res;

      req.__set_size(size);
      req.__set_offset(offset);
      req.__set_fh(fi->fh);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsRead(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        int32_t ret = fuse_reply_buf(fuseReq, res.data.c_str(), res.size);
        if (ret != ICEFS_EOK) {
          fuse_reply_err(fuseReq, ret);
          std::cout << "ERR:" << __FUNCTION__ << "fuse_reply_buf" << ret
                    << std::endl;
        }
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

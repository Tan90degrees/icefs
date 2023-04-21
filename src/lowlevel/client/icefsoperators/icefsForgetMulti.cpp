/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:40:45
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsForgetMulti.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsForgetMulti(fuse_req_t fuseReq, size_t count,
                                     struct fuse_forget_data *forgets) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsForgetMultiReq req;
      icefsgrpc::IcefsForgetMultiRes res;
      grpc::ClientContext ctx;
      icefsgrpc::IcefsForgetMultiReq_forget_data *forgetOne;

      req.set_count(count);
      for (size_t i = 0; i < count; ++i) {
        forgetOne = req.add_to_forget();
        forgetOne->set_inode(forgets[i].ino);
        forgetOne->set_nlookup(forgets[i].nlookup);
      }

      grpc::Status status = gRpcClient->DoIcefsForgetMulti(&ctx, req, &res);
      if (!status.ok() || res.status()) {
        ICEFS_PR_ERR_STATUS;
      }
      fuse_reply_none(fuseReq);
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsForgetMultiReq req;
      icefsthrift::IcefsForgetMultiRes res;
      std::vector<icefsthrift::icefsForgetData> forgetMultiData(count);

      for (size_t i = 0; i < count; ++i) {
        forgetMultiData[i].__set_inode(forgets[i].ino);
        forgetMultiData[i].__set_nlookup(forgets[i].nlookup);
      }

      req.__set_count(count);
      req.__set_to_forget(forgetMultiData);

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsForgetMulti(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (res.status) {
        ICEFS_PR_ERR_STATUS;
      }
      fuse_reply_none(fuseReq);
      break;
    }

    default:
      fuse_reply_none(fuseReq);
      break;
  }
}
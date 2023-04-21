/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-21 02:41:03
 * @FilePath: /icefs/src/lowlevel/client/icefsoperators/icefsSetAttr.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include <stdio.h>

#include <string>

#include "icefsClient.hpp"
#include "icefsClientUtils.hpp"

void IcefsClient::DoIcefsSetAttr(fuse_req_t fuseReq, fuse_ino_t inode,
                                 struct stat *attr, int toSet,
                                 struct fuse_file_info *fi) {
  ICEFS_PR_FUNCTION;

  switch (this->clientConfig.linkType) {
    case ICEFS_LINK_USE_GRPC: {
      icefsgrpc::IcefsSetAttrReq req;
      icefsgrpc::IcefsSetAttrRes res;
      grpc::ClientContext ctx;
      icefsgrpc::statStruct *attrToSet = new icefsgrpc::statStruct();
      icefsgrpc::timeStruct *stAtime = new icefsgrpc::timeStruct();
      icefsgrpc::timeStruct *stMtime = new icefsgrpc::timeStruct();
      icefsgrpc::timeStruct *stCtime = new icefsgrpc::timeStruct();
      struct stat attrGot;

      req.set_inode(inode);
      IcefsGRpcFillAttrOut(attrToSet, *attr, stAtime, stMtime, stCtime);
      req.set_allocated_stat(attrToSet);
      req.set_to_set(toSet);
      if (fi) {
        req.set_fh(fi->fh);
        req.set_has_fh(true);
      } else {
        req.set_has_fh(false);
      }

      grpc::Status status = gRpcClient->DoIcefsSetAttr(&ctx, req, &res);
      if (status.ok() && !res.status()) {
        IcefsGRpcFillAttrIn(&attrGot, res.stat());
        fuse_reply_attr(fuseReq, &attrGot, this->clientConfig.cacheTimeout);
      } else {
        fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
        ICEFS_PR_ERR_STATUS;
      }
      break;
    }

    case ICEFS_LINK_USE_THRIFT: {
      icefsthrift::IcefsSetAttrReq req;
      icefsthrift::IcefsSetAttrRes res;
      icefsthrift::statStruct attrToSet;
      icefsthrift::timeStruct stAtime;
      icefsthrift::timeStruct stMtime;
      icefsthrift::timeStruct stCtime;

      struct stat attrGot;

      req.__set_inode(inode);
      IcefsThriftFillAttrOut(&attrToSet, *attr, stAtime, stMtime, stCtime);
      req.__set_stat(attrToSet);
      req.__set_to_set(toSet);
      if (fi) {
        req.__set_fh(fi->fh);
        req.__set_has_fh(true);
      } else {
        req.__set_has_fh(false);
      }

      icefsThriftConn *thriftConn = thriftClientPool->GetIcefsThriftConn();
      thriftConn->thriftClient->DoIcefsSetAttr(res, req);
      thriftClientPool->PutIcefsThriftConn(thriftConn);
      if (!res.status) {
        IcefsThriftFillAttrIn(&attrGot, res.stat);
        fuse_reply_attr(fuseReq, &attrGot, this->clientConfig.cacheTimeout);
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
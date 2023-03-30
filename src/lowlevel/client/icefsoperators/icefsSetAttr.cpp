/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-30 04:19:29
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:25:04
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
  IcefsSetAttrReq req;
  IcefsSetAttrRes res;
  grpc::ClientContext ctx;
  ICEFS_PR_FUNCTION;
  FuseReq *fuseReqToSend = new FuseReq();
  FuseCtx *fuseCtx = new FuseCtx();
  statStruct *attrToSet = new statStruct();
  timeStruct *stAtime = new timeStruct();
  timeStruct *stMtime = new timeStruct();
  timeStruct *stCtime = new timeStruct();
  struct stat attrGot;
  IcefsFillFuseReq(fuseReqToSend, fuseCtx, fuseReq);
  req.set_allocated_req(fuseReqToSend);
  req.set_inode(inode);
  IcefsFillAttrOut(attrToSet, *attr, stAtime, stMtime, stCtime);
  req.set_allocated_stat(attrToSet);
  req.set_to_set(toSet);
  if (fi) {
    FuseFileInfo *fileInfo = new FuseFileInfo();
    IcefsFillFuseFileInfoOut(fileInfo, fi);
    req.set_allocated_file_info(fileInfo);
  }

  grpc::Status status = stub_->DoIcefsSetAttr(&ctx, req, &res);
  if (status.ok() && !res.status()) {
    IcefsFillAttrIn(&attrGot, res.stat());
    fuse_reply_attr(fuseReq, &attrGot, this->config.cacheTimeout);
  } else {
    fuse_reply_err(fuseReq, res.status() ? res.status() : EIO);
    ICEFS_PR_ERR_STATUS;
  }
}
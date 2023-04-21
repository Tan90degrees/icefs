/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-06 07:53:02
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-20 17:26:30
 * @FilePath: /icefs/src/lowlevel/client/include/icefsClientUtils.hpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#ifndef ICEFS_CLIENT_UTILS_HPP
#define ICEFS_CLIENT_UTILS_HPP

#include <stdio.h>

#include "icefsClient.hpp"

#ifdef ICEFS_DEBUG

#ifndef ICEFS_PR_FUNCTION
#define ICEFS_PR_FUNCTION                     \
  do {                                        \
    printf("ICEFS_FUNC: %s\n", __FUNCTION__); \
  } while (0);
#endif

#ifndef ICEFS_PR_ERR_STATUS
#define ICEFS_PR_ERR_STATUS                  \
  do {                                       \
    printf("ICEFS_ERR: %s\n", __FUNCTION__); \
  } while (0);
#endif

#else

#ifndef ICEFS_PR_FUNCTION
#define ICEFS_PR_FUNCTION
#endif

#ifndef ICEFS_PR_ERR_STATUS
#define ICEFS_PR_ERR_STATUS
#endif

#endif

void IcefsFillFuseReq(icefsgrpc::FuseReq *dstReq, icefsgrpc::FuseCtx *dstCtx,
                      const fuse_req_t srcReq);

void IcefsFillFuseFileInfoOut(icefsgrpc::FuseFileInfo *dstFileInfo,
                              const struct fuse_file_info *srcFileInfo);

void IcefsGRpcFillAttrOut(icefsgrpc::statStruct *dstAttr,
                          const struct stat &srcAttr,
                          icefsgrpc::timeStruct *stAtime,
                          icefsgrpc::timeStruct *stMtime,
                          icefsgrpc::timeStruct *stCtime);

void IcefsThriftFillAttrOut(icefsthrift::statStruct *dstAttr,
                            const struct stat &srcAttr,
                            icefsthrift::timeStruct &stAtime,
                            icefsthrift::timeStruct &stMtime,
                            icefsthrift::timeStruct &stCtime);

void IcefsFillFlockStructOut(icefsgrpc::flockStruct &dstFlock,
                             const struct flock *lock);

void IcefsGRpcFillAttrIn(struct stat *dstAttr,
                         const icefsgrpc::statStruct &srcAttr);

void IcefsThriftFillAttrIn(struct stat *dstAttr,
                           const icefsthrift::statStruct &srcAttr);

void IcefsGRpcFillFuseEntryParamIn(struct fuse_entry_param *dstEntry,
                                   const icefsgrpc::FuseEntryParam &srcEntry);

void IcefsThriftFillFuseEntryParamIn(
    struct fuse_entry_param *dstEntry,
    const icefsthrift::FuseEntryParam &srcEntry);

void IcefsGRpcFillStatvfsIn(struct statvfs *dstStat,
                            const icefsgrpc::statvfsStruct &srcStat);

void IcefsThriftFillStatvfsIn(struct statvfs *dstStat,
                              const icefsthrift::statvfsStruct &srcStat);

#endif
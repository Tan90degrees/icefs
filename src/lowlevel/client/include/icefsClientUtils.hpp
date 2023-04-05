/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-06 07:53:02
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 16:03:27
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
#define ICEFS_PR_ERR_STATUS                                    \
  do {                                                         \
    printf("ICEFS_ERR: %s: %d\n", __FUNCTION__, res.status()); \
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

void IcefsFillFuseReq(FuseReq *dstReq, FuseCtx *dstCtx,
                      const fuse_req_t srcReq);
void IcefsFillAttrIn(struct stat *dstAttr, const statStruct &srcAttr);
void IcefsFillAttrOut(statStruct *dstAttr, const struct stat &srcAttr,
                      timeStruct *stAtime, timeStruct *stMtime,
                      timeStruct *stCtime);
void IcefsFillFuseEntryParamIn(struct fuse_entry_param *dstEntry,
                               const FuseEntryParam &srcEntry);
void IcefsFillFuseFileInfoOut(FuseFileInfo *dstFileInfo,
                              const struct fuse_file_info *srcFileInfo);
void IcefsFillStatvfsIn(struct statvfs *dstStat, const statvfsStruct &srcStat);
void IcefsFillFlockStructIn(flockStruct &dstFlock, const struct flock *lock);

#endif
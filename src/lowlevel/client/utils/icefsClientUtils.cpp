/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-06 07:53:14
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-20 13:57:39
 * @FilePath: /icefs/src/lowlevel/client/utils/icefsClientUtils.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include "icefsClientUtils.hpp"

struct fuse_req {
  struct fuse_session *se;
  uint64_t unique;
  int ctr;
  pthread_mutex_t lock;
  struct fuse_ctx ctx;
  struct fuse_chan *ch;
  int interrupted;
  unsigned int ioctl_64bit : 1;
  union {
    struct {
      uint64_t unique;
    } i;
    struct {
      fuse_interrupt_func_t func;
      void *data;
    } ni;
  } u;
  struct fuse_req *next;
  struct fuse_req *prev;
};

static inline void icefsFillFuseCtx(icefsgrpc::FuseCtx *dstCtx,
                                    const struct fuse_ctx &srcCtx) {
  dstCtx->set_uid(srcCtx.uid);
  dstCtx->set_gid(srcCtx.gid);
  dstCtx->set_pid(srcCtx.pid);
  dstCtx->set_umask(srcCtx.umask);
}

void IcefsFillFuseReq(icefsgrpc::FuseReq *dstReq, icefsgrpc::FuseCtx *dstCtx,
                      const fuse_req_t srcReq) {
  icefsFillFuseCtx(dstCtx, srcReq->ctx);
  dstReq->set_allocated_ctx(dstCtx);
  dstReq->set_unique(srcReq->u.i.unique);
}

void IcefsFillFuseFileInfoOut(icefsgrpc::FuseFileInfo *dstFileInfo,
                              const struct fuse_file_info *srcFileInfo) {
  dstFileInfo->set_flags(srcFileInfo->flags);
  dstFileInfo->set_writepage(srcFileInfo->writepage);
  dstFileInfo->set_direct_io(srcFileInfo->direct_io);
  dstFileInfo->set_keep_cache(srcFileInfo->keep_cache);
  dstFileInfo->set_flush(srcFileInfo->flush);
  dstFileInfo->set_nonseekable(srcFileInfo->nonseekable);
  dstFileInfo->set_flock_release(srcFileInfo->flock_release);
  dstFileInfo->set_cache_readdir(srcFileInfo->cache_readdir);
  dstFileInfo->set_noflush(srcFileInfo->noflush);
  dstFileInfo->set_fh(srcFileInfo->fh);
  dstFileInfo->set_lock_owner(srcFileInfo->lock_owner);
  dstFileInfo->set_poll_events(srcFileInfo->poll_events);
}

static inline void icefsGRpcFillTimeSructOut(icefsgrpc::timeStruct *dstTime,
                                             const struct timespec &srcTime) {
  dstTime->set_time_sec(srcTime.tv_sec);
  dstTime->set_time_n_sec(srcTime.tv_nsec);
}

void IcefsGRpcFillAttrOut(icefsgrpc::statStruct *dstAttr,
                          const struct stat &srcAttr,
                          icefsgrpc::timeStruct *stAtime,
                          icefsgrpc::timeStruct *stMtime,
                          icefsgrpc::timeStruct *stCtime) {
  dstAttr->set_st_dev(srcAttr.st_dev);
  dstAttr->set_st_ino(srcAttr.st_ino);
  dstAttr->set_st_mode(srcAttr.st_mode);
  dstAttr->set_st_nlink(srcAttr.st_nlink);
  dstAttr->set_st_uid(srcAttr.st_uid);
  dstAttr->set_st_gid(srcAttr.st_gid);
  dstAttr->set_st_rdev(srcAttr.st_rdev);
  dstAttr->set_st_size(srcAttr.st_size);
  icefsGRpcFillTimeSructOut(stAtime, srcAttr.st_atim);
  icefsGRpcFillTimeSructOut(stMtime, srcAttr.st_mtim);
  icefsGRpcFillTimeSructOut(stCtime, srcAttr.st_ctim);
  dstAttr->set_allocated_st_atim(stAtime);
  dstAttr->set_allocated_st_mtim(stMtime);
  dstAttr->set_allocated_st_ctim(stCtime);
  dstAttr->set_st_blksize(srcAttr.st_blksize);
  dstAttr->set_st_blocks(srcAttr.st_blocks);
}

static inline void icefsThriftFillTimeSructOut(icefsthrift::timeStruct &dstTime,
                                               const struct timespec &srcTime) {
  dstTime.__set_time_sec(srcTime.tv_sec);
  dstTime.__set_time_n_sec(srcTime.tv_nsec);
}

void IcefsThriftFillAttrOut(icefsthrift::statStruct *dstAttr,
                            const struct stat &srcAttr,
                            icefsthrift::timeStruct &stAtime,
                            icefsthrift::timeStruct &stMtime,
                            icefsthrift::timeStruct &stCtime) {
  dstAttr->__set_st_dev(srcAttr.st_dev);
  dstAttr->__set_st_ino(srcAttr.st_ino);
  dstAttr->__set_st_mode(srcAttr.st_mode);
  dstAttr->__set_st_nlink(srcAttr.st_nlink);
  dstAttr->__set_st_uid(srcAttr.st_uid);
  dstAttr->__set_st_gid(srcAttr.st_gid);
  dstAttr->__set_st_rdev(srcAttr.st_rdev);
  dstAttr->__set_st_size(srcAttr.st_size);
  icefsThriftFillTimeSructOut(stAtime, srcAttr.st_atim);
  icefsThriftFillTimeSructOut(stMtime, srcAttr.st_mtim);
  icefsThriftFillTimeSructOut(stCtime, srcAttr.st_ctim);
  dstAttr->__set_st_atim(stAtime);
  dstAttr->__set_st_mtim(stMtime);
  dstAttr->__set_st_ctim(stCtime);
  dstAttr->__set_st_blksize(srcAttr.st_blksize);
  dstAttr->__set_st_blocks(srcAttr.st_blocks);
}

void IcefsFillFlockStructOut(icefsgrpc::flockStruct &dstFlock,
                             const struct flock *lock) {
  dstFlock.set_lock_type(lock->l_type);
  dstFlock.set_lock_whence(lock->l_whence);
  dstFlock.set_lock_start(lock->l_start);
  dstFlock.set_lock_len(lock->l_len);
  dstFlock.set_lock_pid(lock->l_pid);
}

static inline void icefsGRpcFillTimeSructIn(
    struct timespec *dstTime, const icefsgrpc::timeStruct &srcTime) {
  dstTime->tv_sec = srcTime.time_sec();
  dstTime->tv_nsec = srcTime.time_n_sec();
}

static inline void icefsThriftFillTimeSructIn(
    struct timespec *dstTime, const icefsthrift::timeStruct &srcTime) {
  dstTime->tv_sec = srcTime.time_sec;
  dstTime->tv_nsec = srcTime.time_n_sec;
}

void IcefsGRpcFillAttrIn(struct stat *dstAttr,
                         const icefsgrpc::statStruct &srcAttr) {
  dstAttr->st_dev = srcAttr.st_dev();
  dstAttr->st_ino = srcAttr.st_ino();
  dstAttr->st_mode = srcAttr.st_mode();
  dstAttr->st_nlink = srcAttr.st_nlink();
  dstAttr->st_uid = srcAttr.st_uid();
  dstAttr->st_gid = srcAttr.st_gid();
  dstAttr->st_rdev = srcAttr.st_rdev();
  dstAttr->st_size = srcAttr.st_size();
  icefsGRpcFillTimeSructIn(&dstAttr->st_atim, srcAttr.st_atim());
  icefsGRpcFillTimeSructIn(&dstAttr->st_mtim, srcAttr.st_mtim());
  icefsGRpcFillTimeSructIn(&dstAttr->st_ctim, srcAttr.st_ctim());
  dstAttr->st_blksize = srcAttr.st_blksize();
  dstAttr->st_blocks = srcAttr.st_blocks();
}

void IcefsThriftFillAttrIn(struct stat *dstAttr,
                           const icefsthrift::statStruct &srcAttr) {
  dstAttr->st_dev = srcAttr.st_dev;
  dstAttr->st_ino = srcAttr.st_ino;
  dstAttr->st_mode = srcAttr.st_mode;
  dstAttr->st_nlink = srcAttr.st_nlink;
  dstAttr->st_uid = srcAttr.st_uid;
  dstAttr->st_gid = srcAttr.st_gid;
  dstAttr->st_rdev = srcAttr.st_rdev;
  dstAttr->st_size = srcAttr.st_size;
  icefsThriftFillTimeSructIn(&dstAttr->st_atim, srcAttr.st_atim);
  icefsThriftFillTimeSructIn(&dstAttr->st_mtim, srcAttr.st_mtim);
  icefsThriftFillTimeSructIn(&dstAttr->st_ctim, srcAttr.st_ctim);
  dstAttr->st_blksize = srcAttr.st_blksize;
  dstAttr->st_blocks = srcAttr.st_blocks;
}

void IcefsGRpcFillFuseEntryParamIn(struct fuse_entry_param *dstEntry,
                                   const icefsgrpc::FuseEntryParam &srcEntry) {
  dstEntry->ino = srcEntry.inode();
  dstEntry->generation = srcEntry.generation();
  IcefsGRpcFillAttrIn(&dstEntry->attr, srcEntry.attr());
  dstEntry->attr_timeout = srcEntry.attr_timeout();
  dstEntry->entry_timeout = srcEntry.entry_timeout();
}

void IcefsThriftFillFuseEntryParamIn(
    struct fuse_entry_param *dstEntry,
    const icefsthrift::FuseEntryParam &srcEntry) {
  dstEntry->ino = srcEntry.inode;
  dstEntry->generation = srcEntry.generation;
  IcefsThriftFillAttrIn(&dstEntry->attr, srcEntry.attr);
  dstEntry->attr_timeout = srcEntry.attr_timeout;
  dstEntry->entry_timeout = srcEntry.entry_timeout;
}

void IcefsGRpcFillStatvfsIn(struct statvfs *dstStat,
                            const icefsgrpc::statvfsStruct &srcStat) {
  dstStat->f_bsize = srcStat.f_bsize();
  dstStat->f_frsize = srcStat.f_frsize();
  dstStat->f_blocks = srcStat.f_blocks();
  dstStat->f_bfree = srcStat.f_bfree();
  dstStat->f_bavail = srcStat.f_bavail();
  dstStat->f_files = srcStat.f_files();
  dstStat->f_ffree = srcStat.f_ffree();
  dstStat->f_favail = srcStat.f_favail();
  dstStat->f_fsid = srcStat.f_fsid();
  // dstStat->__f_unused = srcStat.unused();
  dstStat->f_flag = srcStat.f_flag();
  dstStat->f_namemax = srcStat.f_namemax();
}

void IcefsThriftFillStatvfsIn(struct statvfs *dstStat,
                              const icefsthrift::statvfsStruct &srcStat) {
  dstStat->f_bsize = srcStat.f_bsize;
  dstStat->f_frsize = srcStat.f_frsize;
  dstStat->f_blocks = srcStat.f_blocks;
  dstStat->f_bfree = srcStat.f_bfree;
  dstStat->f_bavail = srcStat.f_bavail;
  dstStat->f_files = srcStat.f_files;
  dstStat->f_ffree = srcStat.f_ffree;
  dstStat->f_favail = srcStat.f_favail;
  dstStat->f_fsid = srcStat.f_fsid;
  // dstStat->__f_unused = srcStat.unused;
  dstStat->f_flag = srcStat.f_flag;
  dstStat->f_namemax = srcStat.f_namemax;
}

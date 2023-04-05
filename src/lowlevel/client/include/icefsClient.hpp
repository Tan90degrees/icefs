/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-02 06:41:52
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-05 05:44:47
 * @FilePath: /icefs/src/lowlevel/client/include/icefsClient.hpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#ifndef ICEFS_CLIENT_HPP
#define ICEFS_CLIENT_HPP

#ifndef FUSE_USE_VERSION
#define FUSE_USE_VERSION FUSE_MAKE_VERSION(3, 14)
#endif

#include <errno.h>
#include <fuse_lowlevel.h>
#include <grpcpp/grpcpp.h>

#include <string>

#include "icefsServices.grpc.pb.h"

#define ICEFS_CACHE_NEVER_TIMEOUT (0.0)
#define ICEFS_CACHE_NORMAL_TIMEOUT (1.0)
#define ICEFS_CACHE_ALWAYS_TIMEOUT (86400.0)

#define ICEFS_EOK (0)
#define ICEFS_ERR (-1)

enum {
  ICEFS_CACHE_NEVER,   // 0.0s
  ICEFS_CACHE_NORMAL,  // 1.0s
  ICEFS_CACHE_ALWAYS,  // 86400.0s
};

static const double IcefsCacheMode[] = {ICEFS_CACHE_NEVER_TIMEOUT,
                                        ICEFS_CACHE_NORMAL_TIMEOUT,
                                        ICEFS_CACHE_ALWAYS_TIMEOUT};

struct IcefsClientConfig {
  std::string serverAddressFull;
  std::string uuid;
  std::string serverAddress;
  uint16_t port;
  int cacheMode;
  double cacheTimeout;
};

class IcefsClient {
 private:
  IcefsClientConfig config;
  std::unique_ptr<Icefs::Stub> stub_;

 public:
  IcefsClient(std::shared_ptr<grpc::Channel::ChannelInterface> channel,
              const IcefsClientConfig *config)
      : stub_(Icefs::NewStub(channel)) {
    this->config = *config;
  }
  ~IcefsClient();

  void DoIcefsInit(void *userData, struct fuse_conn_info *conn);
  void DoIcefsDestroy(void *userData);
  void DoIcefsLookUp(fuse_req_t fuseReq, fuse_ino_t parentInode,
                     const char *name);
  void DoIcefsForget(fuse_req_t fuseReq, fuse_ino_t inode, uint64_t nLookUp);
  void DoIcefsGetAttr(fuse_req_t fuseReq, fuse_ino_t inode,
                      struct fuse_file_info *fi);
  void DoIcefsSetAttr(fuse_req_t fuseReq, fuse_ino_t inode, struct stat *attr,
                      int toSet, struct fuse_file_info *fi);
  void DoIcefsReadLink(fuse_req_t fuseReq, fuse_ino_t inode);
  void DoIcefsMknod(fuse_req_t fuseReq, fuse_ino_t parentInode,
                    const char *name, mode_t mode, dev_t rdev);
  void DoIcefsMkDir(fuse_req_t fuseReq, fuse_ino_t parentInode,
                    const char *name, mode_t mode);
  void DoIcefsUnlink(fuse_req_t fuseReq, fuse_ino_t parentInode,
                     const char *name);
  void DoIcefsRmDir(fuse_req_t fuseReq, fuse_ino_t parentInode,
                    const char *name);
  void DoIcefsSymLink(fuse_req_t fuseReq, const char *link,
                      fuse_ino_t parentInode, const char *name);
  void DoIcefsRename(fuse_req_t fuseReq, fuse_ino_t parentInode,
                     const char *name, fuse_ino_t newParent,
                     const char *newName, unsigned int flags);
  void DoIcefsLink(fuse_req_t fuseReq, fuse_ino_t inode, fuse_ino_t newParent,
                   const char *newName);
  void DoIcefsOpen(fuse_req_t fuseReq, fuse_ino_t inode,
                   struct fuse_file_info *fi);
  void DoIcefsRead(fuse_req_t fuseReq, fuse_ino_t inode, size_t size, off_t off,
                   struct fuse_file_info *fi);
  void DoIcefsWrite(fuse_req_t fuseReq, fuse_ino_t inode, const char *buf,
                    size_t size, off_t off, struct fuse_file_info *fi);
  void DoIcefsFlush(fuse_req_t fuseReq, fuse_ino_t inode,
                    struct fuse_file_info *fi);
  void DoIcefsRelease(fuse_req_t fuseReq, fuse_ino_t inode,
                      struct fuse_file_info *fi);
  void DoIcefsFsync(fuse_req_t fuseReq, fuse_ino_t inode, int dataSync,
                    struct fuse_file_info *fi);
  void DoIcefsOpenDir(fuse_req_t fuseReq, fuse_ino_t inode,
                      struct fuse_file_info *fi);
  void DoIcefsReadDir(fuse_req_t fuseReq, fuse_ino_t inode, size_t size,
                      off_t off, struct fuse_file_info *fi);
  void DoIcefsReleaseDir(fuse_req_t fuseReq, fuse_ino_t inode,
                         struct fuse_file_info *fi);
  void DoIcefsFsyncDir(fuse_req_t fuseReq, fuse_ino_t inode, int dataSync,
                       struct fuse_file_info *fi);
  void DoIcefsStatFS(fuse_req_t fuseReq, fuse_ino_t inode);
  void DoIcefsSetXattr(fuse_req_t fuseReq, fuse_ino_t inode, const char *name,
                       const char *value, size_t size, int flags);
  void DoIcefsGetXattr(fuse_req_t fuseReq, fuse_ino_t inode, const char *name,
                       size_t size);
  void DoIcefsListXattr(fuse_req_t fuseReq, fuse_ino_t inode, size_t size);
  void DoIcefsRemoveXattr(fuse_req_t fuseReq, fuse_ino_t inode,
                          const char *name);
  void DoIcefsAccess(fuse_req_t fuseReq, fuse_ino_t inode, int mask);
  void DoIcefsCreate(fuse_req_t fuseReq, fuse_ino_t parentInode,
                     const char *name, mode_t mode, struct fuse_file_info *fi);
  void DoIcefsGetLk(fuse_req_t fuseReq, fuse_ino_t inode,
                    struct fuse_file_info *fi, struct flock *lock);
  void DoIcefsSetLk(fuse_req_t fuseReq, fuse_ino_t inode,
                    struct fuse_file_info *fi, struct flock *lock, int sleep);
  void DoIcefsBmap(fuse_req_t fuseReq, fuse_ino_t inode, size_t blockSize,
                   uint64_t idx);
  void DoIcefsIoctl(fuse_req_t fuseReq, fuse_ino_t inode, unsigned int cmd,
                    void *arg, struct fuse_file_info *fi, unsigned flags,
                    const void *inBuf, size_t inBufSize, size_t outBufSize);
  void DoIcefsPoll(fuse_req_t fuseReq, fuse_ino_t inode,
                   struct fuse_file_info *fi, struct fuse_pollhandle *ph);
  void DoIcefsWriteBuf(fuse_req_t fuseReq, fuse_ino_t inode,
                       struct fuse_bufvec *bufVector, off_t off,
                       struct fuse_file_info *fi);
  void DoIcefsRetrieveReply(fuse_req_t fuseReq, void *cookie, fuse_ino_t inode,
                            off_t offset, struct fuse_bufvec *bufVector);
  void DoIcefsForgetMulti(fuse_req_t fuseReq, size_t count,
                          struct fuse_forget_data *forgets);
  void DoIcefsFlock(fuse_req_t fuseReq, fuse_ino_t inode,
                    struct fuse_file_info *fi, int op);
  void DoIcefsFallocate(fuse_req_t fuseReq, fuse_ino_t inode, int mode,
                        off_t offset, off_t length, struct fuse_file_info *fi);
  void DoIcefsReadDirPlus(fuse_req_t fuseReq, fuse_ino_t inode, size_t size,
                          off_t off, struct fuse_file_info *fi);
  void DoIcefsCopyFileRange(fuse_req_t fuseReq, fuse_ino_t inodeIn, off_t offIn,
                            struct fuse_file_info *fiIn, fuse_ino_t inodeOut,
                            off_t offOut, struct fuse_file_info *fiOut,
                            size_t len, int flags);
  void DoIcefsLseek(fuse_req_t fuseReq, fuse_ino_t inode, off_t off, int whence,
                    struct fuse_file_info *fi);
};

#endif
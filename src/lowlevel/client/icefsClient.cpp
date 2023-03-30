/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-21 07:32:56
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 14:22:20
 * @FilePath: /icefs/src/lowlevel/client/icefsClient.cpp
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
#include "icefsClient.hpp"

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include <string>

#include "yyjson.h"

#define ICEFS_CLIENT_VERSION 1.0

static IcefsClient *g_icefsClient = nullptr;

static void icefsInit(void *userdata, struct fuse_conn_info *conn) {
  g_icefsClient->DoIcefsInit(userdata, conn);
}

static void icefsDestroy(void *userdata) {
  g_icefsClient->DoIcefsDestroy(userdata);
}

static void icefsLookUp(fuse_req_t req, fuse_ino_t parentInode,
                        const char *name) {
  g_icefsClient->DoIcefsLookUp(req, parentInode, name);
}

static void icefsForget(fuse_req_t req, fuse_ino_t ino, uint64_t nlookup) {
  g_icefsClient->DoIcefsForget(req, ino, nlookup);
}

static void icefsGetAttr(fuse_req_t req, fuse_ino_t ino,
                         struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsGetAttr(req, ino, fi);
}

static void icefsSetAttr(fuse_req_t req, fuse_ino_t ino, struct stat *attr,
                         int to_set, struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsSetAttr(req, ino, attr, to_set, fi);
}

static void icefsReadLink(fuse_req_t req, fuse_ino_t ino) {
  g_icefsClient->DoIcefsReadLink(req, ino);
}

static void icefsMknod(fuse_req_t req, fuse_ino_t parentInode, const char *name,
                       mode_t mode, dev_t rdev) {
  g_icefsClient->DoIcefsMknod(req, parentInode, name, mode, rdev);
}

static void icefsMkDir(fuse_req_t req, fuse_ino_t parentInode, const char *name,
                       mode_t mode) {
  g_icefsClient->DoIcefsMkDir(req, parentInode, name, mode);
}

static void icefsUnlink(fuse_req_t req, fuse_ino_t parentInode,
                        const char *name) {
  g_icefsClient->DoIcefsUnlink(req, parentInode, name);
}

static void icefsRmDir(fuse_req_t req, fuse_ino_t parentInode,
                       const char *name) {
  g_icefsClient->DoIcefsRmDir(req, parentInode, name);
}

static void icefsSymLink(fuse_req_t req, const char *link,
                         fuse_ino_t parentInode, const char *name) {
  g_icefsClient->DoIcefsSymLink(req, link, parentInode, name);
}

static void icefsRename(fuse_req_t req, fuse_ino_t parentInode,
                        const char *name, fuse_ino_t newparent,
                        const char *newname, unsigned int flags) {
  g_icefsClient->DoIcefsRename(req, parentInode, name, newparent, newname,
                               flags);
}

static void icefsLink(fuse_req_t req, fuse_ino_t ino, fuse_ino_t newparent,
                      const char *newname) {
  g_icefsClient->DoIcefsLink(req, ino, newparent, newname);
}

static void icefsOpen(fuse_req_t req, fuse_ino_t ino,
                      struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsOpen(req, ino, fi);
}

static void icefsRead(fuse_req_t req, fuse_ino_t ino, size_t size, off_t off,
                      struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsRead(req, ino, size, off, fi);
}

static void icefsWrite(fuse_req_t req, fuse_ino_t ino, const char *buf,
                       size_t size, off_t off, struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsWrite(req, ino, buf, size, off, fi);
}

static void icefsFlush(fuse_req_t req, fuse_ino_t ino,
                       struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsFlush(req, ino, fi);
}

static void icefsRelease(fuse_req_t req, fuse_ino_t ino,
                         struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsRelease(req, ino, fi);
}

static void icefsFsync(fuse_req_t req, fuse_ino_t ino, int datasync,
                       struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsFsync(req, ino, datasync, fi);
}

static void icefsOpenDir(fuse_req_t req, fuse_ino_t ino,
                         struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsOpenDir(req, ino, fi);
}

static void icefsReadDir(fuse_req_t req, fuse_ino_t ino, size_t size, off_t off,
                         struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsReadDir(req, ino, size, off, fi);
}

static void icefsReleaseDir(fuse_req_t req, fuse_ino_t ino,
                            struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsReleaseDir(req, ino, fi);
}

static void icefsFsyncDir(fuse_req_t req, fuse_ino_t ino, int datasync,
                          struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsFsyncDir(req, ino, datasync, fi);
}

static void icefsStatFS(fuse_req_t req, fuse_ino_t ino) {
  g_icefsClient->DoIcefsStatFS(req, ino);
}

static void icefsSetXattr(fuse_req_t req, fuse_ino_t ino, const char *name,
                          const char *value, size_t size, int flags) {
  g_icefsClient->DoIcefsSetXattr(req, ino, name, value, size, flags);
}

static void icefsGetXattr(fuse_req_t req, fuse_ino_t ino, const char *name,
                          size_t size) {
  g_icefsClient->DoIcefsGetXattr(req, ino, name, size);
}

static void icefsListXattr(fuse_req_t req, fuse_ino_t ino, size_t size) {
  g_icefsClient->DoIcefsListXattr(req, ino, size);
}

static void icefsRemoveXattr(fuse_req_t req, fuse_ino_t ino, const char *name) {
  g_icefsClient->DoIcefsRemoveXattr(req, ino, name);
}

static void icefsAccess(fuse_req_t req, fuse_ino_t ino, int mask) {
  g_icefsClient->DoIcefsAccess(req, ino, mask);
}

static void icefsCreate(fuse_req_t req, fuse_ino_t parentInode,
                        const char *name, mode_t mode,
                        struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsCreate(req, parentInode, name, mode, fi);
}

static void icefsGetLk(fuse_req_t req, fuse_ino_t ino,
                       struct fuse_file_info *fi, struct flock *lock) {
  g_icefsClient->DoIcefsGetLk(req, ino, fi, lock);
}

static void icefsSetLk(fuse_req_t req, fuse_ino_t ino,
                       struct fuse_file_info *fi, struct flock *lock,
                       int sleep) {
  g_icefsClient->DoIcefsSetLk(req, ino, fi, lock, sleep);
}

static void icefsBmap(fuse_req_t req, fuse_ino_t ino, size_t blocksize,
                      uint64_t idx) {
  g_icefsClient->DoIcefsBmap(req, ino, blocksize, idx);
}

static void icefsIoctl(fuse_req_t req, fuse_ino_t ino, unsigned int cmd,
                       void *arg, struct fuse_file_info *fi, unsigned flags,
                       const void *in_buf, size_t in_bufsz, size_t out_bufsz) {
  g_icefsClient->DoIcefsIoctl(req, ino, cmd, arg, fi, flags, in_buf, in_bufsz,
                              out_bufsz);
}

static void icefsPoll(fuse_req_t req, fuse_ino_t ino, struct fuse_file_info *fi,
                      struct fuse_pollhandle *ph) {
  g_icefsClient->DoIcefsPoll(req, ino, fi, ph);
}

static void icefsWriteBuf(fuse_req_t req, fuse_ino_t ino,
                          struct fuse_bufvec *bufv, off_t off,
                          struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsWriteBuf(req, ino, bufv, off, fi);
}

static void icefsRetrieveReply(fuse_req_t req, void *cookie, fuse_ino_t ino,
                               off_t offset, struct fuse_bufvec *bufv) {
  g_icefsClient->DoIcefsRetrieveReply(req, cookie, ino, offset, bufv);
}

static void icefsForgetMulti(fuse_req_t req, size_t count,
                             struct fuse_forget_data *forgets) {
  g_icefsClient->DoIcefsForgetMulti(req, count, forgets);
}

static void icefsFlock(fuse_req_t req, fuse_ino_t ino,
                       struct fuse_file_info *fi, int op) {
  g_icefsClient->DoIcefsFlock(req, ino, fi, op);
}

static void icefsFallocate(fuse_req_t req, fuse_ino_t ino, int mode,
                           off_t offset, off_t length,
                           struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsFallocate(req, ino, mode, offset, length, fi);
}

static void icefsReadDirPlus(fuse_req_t req, fuse_ino_t ino, size_t size,
                             off_t off, struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsReadDirPlus(req, ino, size, off, fi);
}

static void icefsCopyFileRange(fuse_req_t req, fuse_ino_t ino_in, off_t off_in,
                               struct fuse_file_info *fi_in, fuse_ino_t ino_out,
                               off_t off_out, struct fuse_file_info *fi_out,
                               size_t len, int flags) {
  g_icefsClient->DoIcefsCopyFileRange(req, ino_in, off_in, fi_in, ino_out,
                                      off_out, fi_out, len, flags);
}

static void icefsLseek(fuse_req_t req, fuse_ino_t ino, off_t off, int whence,
                       struct fuse_file_info *fi) {
  g_icefsClient->DoIcefsLseek(req, ino, off, whence, fi);
}

static const struct fuse_lowlevel_ops icefs_ll_oper = {
    .init = icefsInit,
    .destroy = icefsDestroy,
    .lookup = icefsLookUp,
    .forget = icefsForget,
    .getattr = icefsGetAttr,
    .setattr = icefsSetAttr,
    .readlink = icefsReadLink,
    .mknod = icefsMknod,
    .mkdir = icefsMkDir,
    .unlink = icefsUnlink,
    .rmdir = icefsRmDir,
    .symlink = icefsSymLink,
    .rename = icefsRename,
    .link = icefsLink,
    .open = icefsOpen,
    .read = icefsRead,
    .write = icefsWrite,
    .flush = icefsFlush,
    .release = icefsRelease,
    .fsync = icefsFsync,
    .opendir = icefsOpenDir,
    .readdir = icefsReadDir,
    .releasedir = icefsReleaseDir,
    .fsyncdir = icefsFsyncDir,
    .statfs = icefsStatFS,
    .setxattr = icefsSetXattr,
    .getxattr = icefsGetXattr,
    .listxattr = icefsListXattr,
    .removexattr = icefsRemoveXattr,
    .access = icefsAccess,
    .create = icefsCreate,
    .getlk = icefsGetLk,
    .setlk = icefsSetLk,
    .bmap = icefsBmap,
    .ioctl = icefsIoctl,
    .poll = icefsPoll,
    // .write_buf = icefsWriteBuf, // TODO: 无意义
    .retrieve_reply = icefsRetrieveReply,
    .forget_multi = icefsForgetMulti,
    .flock = icefsFlock,
    .fallocate = icefsFallocate,
    .readdirplus = icefsReadDirPlus,
#ifdef HAVE_COPY_FILE_RANGE
    .copy_file_range = icefsCopyFileRange,
#endif
    .lseek = icefsLseek,
};

int icefsParseConfig(IcefsClientConfig *config) {
  int ret = ICEFS_ERR;
  yyjson_read_flag flag =
      YYJSON_READ_ALLOW_COMMENTS | YYJSON_READ_ALLOW_TRAILING_COMMAS;
  yyjson_read_err err;
  yyjson_doc *doc = yyjson_read_file(ICEFS_CONFIG_PATH, flag, NULL, &err);

  if (doc != nullptr) {
    yyjson_val *root = yyjson_doc_get_root(doc);
    yyjson_val *value = yyjson_obj_get(root, ICEFS_CONFIG_SERV_ADDR);
    if (value != nullptr) {
      config->serverAddress = yyjson_get_str(value);
      ret = ICEFS_EOK;
    } else {
      printf("icefsParseConfig: server_address is not found in config.json.\n");
      ret = ICEFS_ERR;
    }

    value = yyjson_obj_get(root, ICEFS_CONFIG_CACHE_MODE);
    if (value != nullptr) {
      config->cacheMode = yyjson_get_int(value);
      if (config->cacheMode >= 0 &&
          config->cacheMode < sizeof(IcefsCacheMode) / sizeof(double)) {
        config->cacheTimeout = IcefsCacheMode[config->cacheMode];
        ret = ICEFS_EOK;
      } else {
        printf("icefsParseConfig: cache_mode should be 0, 1 or 2.\n");
        ret = ICEFS_ERR;
      }

      value = yyjson_obj_get(root, ICEFS_CONFIG_UUID);
      if (value != nullptr) {
        config->uuid = yyjson_get_str(value);
        ret = ICEFS_EOK;
      } else {
        printf("icefsParseConfig: uuid is not found in config.json.\n");
        ret = ICEFS_ERR;
      }
    } else {
      printf("icefsParseConfig: cache_mode is not found in config.json.\n");
      ret = ICEFS_ERR;
    }
  } else {
    printf("icefsParseConfig: read error (%u): %s at position: %ld\n", err.code,
           err.msg, err.pos);
  }

  yyjson_doc_free(doc);
  return ret;
}

int icefsCreateClient(const IcefsClientConfig *config) {
  g_icefsClient =
      new IcefsClient(grpc::CreateChannel(config->serverAddress,
                                          grpc::InsecureChannelCredentials()),
                      config);
  return g_icefsClient != nullptr ? ICEFS_EOK : ICEFS_ERR;
}

int main(int argc, char *argv[]) {
  struct fuse_args args = FUSE_ARGS_INIT(argc, argv);
  struct fuse_session *se;
  struct fuse_cmdline_opts opts;
  IcefsClientConfig config;
  struct fuse_loop_config *loopConfig = fuse_loop_cfg_create();
  int ret = ICEFS_ERR;

  if (fuse_parse_cmdline(&args, &opts) != ICEFS_EOK) return ICEFS_ERR;

  if (opts.show_help) {
    printf("usage: %s [options] <mountpoint>\n", argv[0]);
    fuse_cmdline_help();
    fuse_lowlevel_help();
    ret = ICEFS_EOK;
    goto errOut1;
  } else if (opts.show_version) {
    printf("Icefs client version: %f\n", ICEFS_CLIENT_VERSION);
    printf("Libfuse version: %s\n", fuse_pkgversion());
    fuse_lowlevel_version();
    ret = ICEFS_EOK;
    goto errOut1;
  }

  if (opts.mountpoint == NULL) {
    printf("usage: %s [options] <mountpoint>\n", argv[0]);
    printf("       %s --help\n", argv[0]);
    ret = ICEFS_ERR;
    goto errOut1;
  }

  se = fuse_session_new(&args, &icefs_ll_oper, sizeof(icefs_ll_oper), NULL);
  if (se == NULL) goto errOut1;

  if (fuse_set_signal_handlers(se) != 0) goto errOut2;

  umask(0);

  if (fuse_session_mount(se, opts.mountpoint) != 0) goto errOut3;

  ret = icefsParseConfig(&config);

  if (ret != ICEFS_EOK) goto errOut4;

  fuse_daemonize(opts.foreground);

  ret = icefsCreateClient(&config);

  if (ret != ICEFS_EOK) goto errOut4;

  if (opts.singlethread)
    ret = fuse_session_loop(se);
  else {
    fuse_loop_cfg_set_clone_fd(loopConfig, opts.clone_fd);
    fuse_loop_cfg_set_idle_threads(loopConfig, opts.max_idle_threads);
    ret = fuse_session_loop_mt(se, loopConfig);
  }

errOut4:
  fuse_session_unmount(se);
errOut3:
  fuse_remove_signal_handlers(se);
errOut2:
  fuse_session_destroy(se);
errOut1:
  free(opts.mountpoint);
  fuse_loop_cfg_destroy(loopConfig);
  fuse_opt_free_args(&args);

  return ret ? ICEFS_ERR : ICEFS_EOK;
}

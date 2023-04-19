/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-04-17 09:26:19
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 12:43:14
 * @FilePath: /icefs/thrift/icefsServices.thrift
 * @Description: 
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */

typedef i8 ui8
typedef i16  ui16
typedef i32  ui32
typedef i64  ui64

struct FuseCtx {
  1: ui32 uid;
  2: ui32 gid;
  3: i32 pid;
  4: ui32 umask;
}

struct FuseReq {
  1: ui64 unique;
  2: FuseCtx ctx;
}

struct FuseFileInfo {
  1: i32 flags;
  2: bool writepage;
  3: bool direct_io;
  4: bool keep_cache;
  5: bool flush;
  6: bool nonseekable;
  7: bool flock_release;
  8: bool cache_readdir;
  9: bool noflush;
  10: ui64 fh;
  11: ui64 lock_owner;
  12: ui32 poll_events;
}

struct flockStruct {
  1: i32 lock_type;
  2: i32 lock_whence;
  3: i64 lock_start;
  4: i64 lock_len;
  5: i32 lock_pid;
}

struct timeStruct {
  1: i64 time_sec;
  2: i64 time_n_sec;
}

struct statStruct {
  1: ui64 st_dev;
  2: ui64 st_ino;
  3: ui32 st_mode;
  4: ui64 st_nlink;
  5: ui32 st_uid;
  6: ui32 st_gid;
  7: ui64 st_rdev;
  8: i64 st_size;
  9: timeStruct st_atim;
  10: timeStruct st_mtim;
  11: timeStruct st_ctim;
  12: i64 st_blksize;
  13: i64 st_blocks;
}

struct FuseEntryParam {
  1: ui64 inode;
  2: ui64 generation;
  3: statStruct attr;
  4: double attr_timeout;
  5: double entry_timeout;
}

struct DirentStruct {
  1: ui64 ino;
  2: i64 off;
  3: ui32 reclen;
  4: ui32 type;
  5: string name;
}

struct statvfsStruct {
  1: ui64 f_bsize;
  2: ui64 f_frsize;
  3: ui64 f_blocks;
  4: ui64 f_bfree;
  5: ui64 f_bavail;
  6: ui64 f_files;
  7: ui64 f_ffree;
  8: ui64 f_favail;
  9: ui64 f_fsid;
  10: i32 unused;
  11: ui64 f_flag;
  12: ui64 f_namemax;
}

struct ioVector {
  1: ui64 size;
  2: binary data;
}

// Request
struct IcefsAccessReq {
  1: ui64 inode;
  2: i32 mask;
}

struct IcefsBmapReq {
  1: FuseReq req;
  2: ui64 inode;
  3: ui64 block_size;
  4: ui64 index;
}

struct IcefsCopyFileRangeReq {
  1: i64 offset_in;
  2: ui64 fh_in;
  3: i64 offset_out;
  4: ui64 fh_out;
  5: ui64 len;
  6: i32 flags;
}

struct IcefsCreateReq {
  1: ui64 parent_inode;
  2: string name;
  3: ui32 mode;
  4: i32 flags;
}

struct IcefsDestroyReq {
  1: string host_name;
  2: string info;
}

struct IcefsFallocateReq {
  1: i32 mode;
  2: i64 offset;
  3: i64 length;
  4: ui64 fh;
}

struct IcefsFlockReq {
  1: ui64 fh;
  2: i32 op;
}

struct IcefsFlushReq {
  1: ui64 fh;
}

struct icefsForgetData {
    1: ui64 inode;
    2: ui64 nlookup;
}

struct IcefsForgetMultiReq {
  1: ui64 count;
  2: list<icefsForgetData> to_forget;
}

struct IcefsForgetReq {
  1: ui64 inode;
  2: ui64 nlookup;
}

struct IcefsFsyncReq {
  1: i32 data_sync;
  2: ui64 fh;
}

struct IcefsFsyncDirReq {
  1: i32 data_sync;
  2: ui64 fh;
}

struct IcefsGetAttrReq {
  1: ui64 inode;
}

struct IcefsGetLkReq {
  1: FuseReq req;
  2: ui64 inode;
  3: FuseFileInfo file_info;
  4: flockStruct lock;
}

struct IcefsGetXattrReq {
  1: ui64 inode;
  2: string name;
  3: ui64 size;
}

struct IcefsInitReq {
  1: string uuid;
  2: string info;
  3: ui32 want;
  4: double timeout;
}

struct IcefsIoctlReq {
  1: FuseReq req;
  2: ui64 inode;
  3: ui32 cmd;
  4: binary arg;
  5: FuseFileInfo file_info;
  6: ui32 flags;
  7: ui64 in_buf_size;
  8: ui64 out_buf_size;
}

struct IcefsLinkReq {
  1: ui64 inode;
  2: ui64 new_parent_inode;
  3: string new_name;
}

struct IcefsListXattrReq {
  1: ui64 inode;
  2: ui64 size;
}

struct IcefsLookUpReq {
  1: ui64 parent_inode;
  2: string name;
}

struct IcefsLseekReq {
  1: ui64 fh;
  2: i64 offset;
  3: i32 whence;
}

struct IcefsMkDirReq {
  1: ui64 parent_inode;
  2: string name;
  3: ui32 mode;
}

struct IcefsMknodReq {
  1: ui64 parent_inode;
  2: string name;
  3: ui32 mode;
  4: ui64 rdev;
}

struct IcefsOpenReq {
  1: ui64 inode;
  2: i32 flags;
}

struct IcefsOpenDirReq {
  1: ui64 inode;
}

struct IcefsPollReq {
  1: FuseReq req;
  2: ui64 inode;
  3: FuseFileInfo file_info;
}

struct IcefsReadReq {
  1: ui64 size;
  2: i64 offset;
  3: ui64 fh;
}

struct IcefsReadDirReq {
  1: i64 offset;
  2: ui64 fh;
}

struct IcefsReadDirPlusReq {
  1: ui64 inode;
  2: i64 offset;
  3: ui64 fh;
}

struct IcefsReadLinkReq {
  1: ui64 inode;
}

struct IcefsReleaseReq {
  1: ui64 inode;
  2: ui64 fh;
}

struct IcefsReleaseDirReq {
  1: ui64 fh;
}

struct IcefsRemoveXattrReq {
  1: ui64 inode;
  2: string name;
}

struct IcefsRenameReq {
  1: ui64 parent_inode;
  2: string name;
  3: ui64 new_parent_inode;
  4: string new_name;
  5: ui32 flags;
}

struct IcefsRetrieveReplyReq {
  1: FuseReq req;
  2: binary cookie;
  3: ui64 inode;
  4: i64 offset;
  5: list<ioVector> data;
}

struct IcefsRmDirReq {
  1: ui64 parent_inode;
  2: string name;
}

struct IcefsSetAttrReq {
  1: ui64 inode;
  2: statStruct stat;
  3: i32 to_set;
  4: ui64 fh;
  5: bool has_fh;
}

struct IcefsSetLkReq {
  1: FuseReq req;
  2: ui64 inode;
  3: FuseFileInfo file_info;
  4: flockStruct lock;
  5: i32 sleep;
}

struct IcefsSetXattrReq {
  1: ui64 inode;
  2: string name;
  3: string value;
  4: i32 flags;
}

struct IcefsStatFSReq {
  1: ui64 inode;
}

struct IcefsSymLinkReq {
  1: string link;
  2: ui64 parent_inode;
  3: string name;
}

struct IcefsUnlinkReq {
  1: ui64 parent_inode;
  2: string name;
}

struct IcefsWriteBufReq {
  1: FuseReq req;
  2: ui64 inode;
  3: list<ioVector> buf;
  4: i64 offset;
  5: FuseFileInfo file_info;
}

struct IcefsWriteReq {
  1: binary buf;
  2: i64 offset;
  3: ui64 fh;
  4: ui64 size;
}

// Response
struct IcefsAccessRes {
  1: i32 status;
}

struct IcefsBmapRes {
  1: i32 status;
}

struct IcefsCopyFileRangeRes {
  1: i32 status;
  2: ui64 size;
}

struct IcefsCreateRes {
  1: i32 status;
  2: ui64 fh;
  3: FuseEntryParam entry;
}

struct IcefsDestroyRes {
  1: i32 status;
  2: string info;
}

struct IcefsFallocateRes {
  1: i32 status;
}

struct IcefsFlockRes {
  1: i32 status;
}

struct IcefsFlushRes {
  1: i32 status;
}

struct IcefsForgetMultiRes {
  1: i32 status;
}

struct IcefsForgetRes {
  1: i32 status;
}

struct IcefsFsyncRes {
  1: i32 status;
}

struct IcefsFsyncDirRes {
  1: i32 status;
}

struct IcefsGetAttrRes {
  1: i32 status;
  2: statStruct stat;
}

struct IcefsGetLkRes {
  1: i32 status;
}

struct IcefsGetXattrRes {
  1: i32 status;
  2: i64 size;
  3: binary value;
}

struct IcefsInitRes {
  1: i32 status;
  2: string info;
  3: ui32 can;
}

struct IcefsIoctlRes {
  1: i32 status;
 // 2: FuseIoctlOut out;
  2: list<ioVector> data;
}

struct IcefsLinkRes {
  1: i32 status;
  2: FuseEntryParam entry;
}

struct IcefsListXattrRes {
  1: i32 status;
  2: i64 size;
  3: binary value;
}

struct IcefsLookUpRes {
  1: i32 status;
  2: FuseEntryParam entry;
}

struct IcefsLseekRes {
  1: i32 status;
  2: i64 offset;
}

struct IcefsMkDirRes {
  1: i32 status;
  2: FuseEntryParam entry;
}

struct IcefsMknodRes {
  1: i32 status;
  2: FuseEntryParam entry;
}

struct IcefsOpenRes {
  1: i32 status;
  2: ui64 fh;
}

struct IcefsOpenDirRes {
  1: i32 status;
  2: ui64 fh;
}

struct IcefsPollRes {
  1: i32 status;
  2: ui32 revents;
}

struct IcefsReadRes {
  1: i32 status;
  2: ui64 size;
  3: binary data;
}

struct IcefsReadDirRes {
  1: i32 status;
  2: list<DirentStruct> data;
}

struct IcefsReadDirPlusData {
    1: FuseEntryParam entry;
    2: DirentStruct dir_entry;
}

struct IcefsReadDirPlusRes {
  1: i32 status;
  2: list<IcefsReadDirPlusData> data;
}

struct IcefsReadLinkRes {
  1: i32 status;
  2: string path;
}

struct IcefsReleaseRes {
  1: i32 status;
}

struct IcefsReleaseDirRes {
  1: i32 status;
}

struct IcefsRemoveXattrRes {
  1: i32 status;
}

struct IcefsRenameRes {
  1: i32 status;
}

struct IcefsRetrieveReplyRes {
  1: i32 status;
}

struct IcefsRmDirRes {
  1: i32 status;
}

struct IcefsSetAttrRes {
  1: i32 status;
  2: statStruct stat;
}

struct IcefsSetLkRes {
  1: i32 status;
}

struct IcefsSetXattrRes {
  1: i32 status;
}

struct IcefsStatFSRes {
  1: i32 status;
  2: statvfsStruct statvfs;
}

struct IcefsSymLinkRes {
  1: i32 status;
  2: FuseEntryParam entry;
}

struct IcefsUnlinkRes {
  1: i32 status;
}

struct IcefsWriteBufRes {
  1: i32 status;
  2: ui64 size;
}

struct IcefsWriteRes {
  1: i32 status;
  2: ui64 size;
}

// include "icefs.thrift"

service IcefsThrift {
  IcefsInitRes DoIcefsInit(1:IcefsInitReq req);
  IcefsDestroyRes DoIcefsDestroy(1:IcefsDestroyReq req);
  IcefsLookUpRes DoIcefsLookUp(1:IcefsLookUpReq req);
  IcefsForgetRes DoIcefsForget(1:IcefsForgetReq req);
  IcefsGetAttrRes DoIcefsGetAttr(1:IcefsGetAttrReq req);
  IcefsSetAttrRes DoIcefsSetAttr(1:IcefsSetAttrReq req);
  IcefsReadLinkRes DoIcefsReadLink(1:IcefsReadLinkReq req);
  IcefsMknodRes DoIcefsMknod(1:IcefsMknodReq req);
  IcefsMkDirRes DoIcefsMkDir(1:IcefsMkDirReq req);
  IcefsUnlinkRes DoIcefsUnlink(1:IcefsUnlinkReq req);
  IcefsRmDirRes DoIcefsRmDir(1:IcefsRmDirReq req);
  IcefsSymLinkRes DoIcefsSymLink(1:IcefsSymLinkReq req);
  IcefsRenameRes DoIcefsRename(1:IcefsRenameReq req);
  IcefsLinkRes DoIcefsLink(1:IcefsLinkReq req);
  IcefsOpenRes DoIcefsOpen(1:IcefsOpenReq req);
  IcefsReadRes DoIcefsRead(1:IcefsReadReq req);
  IcefsWriteRes DoIcefsWrite(1:IcefsWriteReq req);
  IcefsFlushRes DoIcefsFlush(1:IcefsFlushReq req);
  IcefsReleaseRes DoIcefsRelease(1:IcefsReleaseReq req);
  IcefsFsyncRes DoIcefsFsync(1:IcefsFsyncReq req);
  IcefsOpenDirRes DoIcefsOpenDir(1:IcefsOpenDirReq req);
  IcefsReadDirRes DoIcefsReadDir(1:IcefsReadDirReq req);
  IcefsReleaseDirRes DoIcefsReleaseDir(1:IcefsReleaseDirReq req);
  IcefsFsyncDirRes DoIcefsFsyncDir(1:IcefsFsyncDirReq req);
  IcefsStatFSRes DoIcefsStatFS(1:IcefsStatFSReq req);
  IcefsSetXattrRes DoIcefsSetXattr(1:IcefsSetXattrReq req);
  IcefsGetXattrRes DoIcefsGetXattr(1:IcefsGetXattrReq req);
  IcefsListXattrRes DoIcefsListXattr(1:IcefsListXattrReq req);
  IcefsRemoveXattrRes DoIcefsRemoveXattr(1:IcefsRemoveXattrReq req);
  IcefsAccessRes DoIcefsAccess(1:IcefsAccessReq req);
  IcefsCreateRes DoIcefsCreate(1:IcefsCreateReq req);
  IcefsGetLkRes DoIcefsGetLk(1:IcefsGetLkReq req);
  IcefsSetLkRes DoIcefsSetLk(1:IcefsSetLkReq req);
  IcefsBmapRes DoIcefsBmap(1:IcefsBmapReq req);
  IcefsIoctlRes DoIcefsIoctl(1:IcefsIoctlReq req);
  IcefsPollRes DoIcefsPoll(1:IcefsPollReq req);
  IcefsWriteBufRes DoIcefsWriteBuf(1:IcefsWriteBufReq req);
  IcefsRetrieveReplyRes DoIcefsRetrieveReply(1:IcefsRetrieveReplyReq req);
  IcefsForgetMultiRes DoIcefsForgetMulti(1:IcefsForgetMultiReq req);
  IcefsFlockRes DoIcefsFlock(1:IcefsFlockReq req);
  IcefsFallocateRes DoIcefsFallocate(1:IcefsFallocateReq req);
  IcefsReadDirPlusRes DoIcefsReadDirPlus(1:IcefsReadDirPlusReq req);
  IcefsCopyFileRangeRes DoIcefsCopyFileRange(1:IcefsCopyFileRangeReq req);
  IcefsLseekRes DoIcefsLseek(1:IcefsLseekReq req);
}
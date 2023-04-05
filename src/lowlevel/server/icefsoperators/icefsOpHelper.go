/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 14:47:23
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:31:50
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpHelper.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"log"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	ICEFS_NO_LOOKUP = 0
	FUSE_ROOT_INODE = 1

	FUSE_SET_ATTR_MODE      = (1 << 0)
	FUSE_SET_ATTR_UID       = (1 << 1)
	FUSE_SET_ATTR_GID       = (1 << 2)
	FUSE_SET_ATTR_SIZE      = (1 << 3)
	FUSE_SET_ATTR_ATIME     = (1 << 4)
	FUSE_SET_ATTR_MTIME     = (1 << 5)
	FUSE_SET_ATTR_ATIME_NOW = (1 << 7)
	FUSE_SET_ATTR_MTIME_NOW = (1 << 8)
	FUSE_SET_ATTR_FORCE     = (1 << 9)
	FUSE_SET_ATTR_CTIME     = (1 << 10)
	FUSE_SET_ATTR_KILL_SUID = (1 << 11)
	FUSE_SET_ATTR_KILL_SGID = (1 << 12)
	FUSE_SET_ATTR_FILE      = (1 << 13)
	FUSE_SET_ATTR_KILL_PRIV = (1 << 14)
	FUSE_SET_ATTR_OPEN      = (1 << 15)
	FUSE_SET_ATTR_TIMES_SET = (1 << 16)
	FUSE_SET_ATTR_TOUCH     = (1 << 17)
)

func FuseEntryParamBuilder(inode *IcefsInode, timeout float64) *pb.FuseEntryParam {
	return &pb.FuseEntryParam{
		Inode:        inode.fakeInode,
		Generation:   inode.generation,
		Attr:         StatStructBuilder(&inode.stat),
		AttrTimeout:  timeout,
		EntryTimeout: timeout,
	}
}

func UnixStatFillSyscallStat(dstStat *syscall.Stat_t, srcStat *unix.Stat_t) {
	dstStat.Dev = srcStat.Dev
	dstStat.Ino = srcStat.Ino
	dstStat.Nlink = srcStat.Nlink
	dstStat.Mode = srcStat.Mode
	dstStat.Uid = srcStat.Uid
	dstStat.Gid = srcStat.Gid
	dstStat.Rdev = srcStat.Rdev
	dstStat.Size = srcStat.Size
	dstStat.Blksize = srcStat.Blksize
	dstStat.Blocks = srcStat.Blocks
	dstStat.Atim = syscall.Timespec(srcStat.Atim)
	dstStat.Mtim = syscall.Timespec(srcStat.Mtim)
	dstStat.Ctim = syscall.Timespec(srcStat.Ctim)
}

// 有锁操作
func (s *IcefsServer) doLookUp(parentInodeNum uint64, name string) (*pb.FuseEntryParam, error) {
	var err error
	inode := new(IcefsInode)
	s.inodeCacheLock.RLock()
	parentInode := s.getIcefsInode(parentInodeNum)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		err = syscall.Errno(icefserror.ICEFS_BUG_ERR)
		return nil, err
	}

	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	newfd, err := syscall.Openat(parentInode.fd, name, unix.O_PATH|unix.O_NOFOLLOW, 0)
	parentInode.inodeLock.RUnlock()
	if err != nil {
		return nil, err
	}
	var stat unix.Stat_t
	err = unix.Fstatat(newfd, "", &stat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		return nil, err
	}

	UnixStatFillSyscallStat(&inode.stat, &stat)

	inode.nlookup++
	s.inodeCacheLock.Lock()
	if s.checkIcefsInode(&inode) {
		syscall.Close(newfd)
	} else {
		inode.fd = newfd
		IcefsSetFakeInode(inode)
		s.putIcefsInode(inode)
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.Unlock()
	entry := FuseEntryParamBuilder(inode, s.timeout)
	inode.inodeLock.RUnlock()

	return entry, err
}

func (s *IcefsServer) doForget(inodeNum uint64, nlookup uint64) {
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(inodeNum)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		return
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	if nlookup > inode.nlookup {
		inode.inodeLock.Unlock()
		log.Fatal("Icefs ERROR: Invalid lookup count")
	}
	inode.nlookup -= nlookup
	if inode.nlookup == ICEFS_NO_LOOKUP {
		inode.inodeLock.Unlock()
		s.delIcefsInode(inodeNum)
		return
	}
	inode.inodeLock.Unlock()
}

func (s *IcefsServer) doGetAttr(fd int, stat *syscall.Stat_t) error {
	var unixStat unix.Stat_t
	err := unix.Fstatat(fd, "", &unixStat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		return err
	}
	UnixStatFillSyscallStat(stat, &unixStat)
	return err
}

// 提升性能
func checkNameIsDotOrDotDot(name string) bool {
	return name == "." || name == ".."
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 14:47:23
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 14:55:16
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpHelper.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"icefs-server/icefserror"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"log"
	"syscall"
	"unsafe"

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

type FuseEntryParamBuilder func(inode *IcefsInode, timeout float64) any
type StatStructBuilder func(stat *syscall.Stat_t) any
type StatvfsStructBuilder func(statfs *syscall.Statfs_t) any

func GRpcFuseEntryParamBuilder(inode *IcefsInode, timeout float64) any {
	return &pb.FuseEntryParam{
		Inode:        inode.fakeInode,
		Generation:   inode.generation,
		Attr:         GRpcStatStructBuilder(&inode.stat).(*pb.StatStruct),
		AttrTimeout:  timeout,
		EntryTimeout: timeout,
	}
}

func ThriftFuseEntryParamBuilder(inode *IcefsInode, timeout float64) any {
	return &icefsthrift.FuseEntryParam{
		Inode:        icefsthrift.Ui64(inode.fakeInode),
		Generation:   icefsthrift.Ui64(inode.generation),
		Attr:         ThriftStatStructBuilder(&inode.stat).(*icefsthrift.StatStruct),
		AttrTimeout:  timeout,
		EntryTimeout: timeout,
	}
}

func GRpcStatvfsStructBuilder(statfs *syscall.Statfs_t) any {
	return &pb.StatvfsStruct{
		FBsize:   uint64(statfs.Bsize),
		FFrsize:  uint64(statfs.Frsize),
		FBlocks:  statfs.Blocks,
		FBfree:   statfs.Ffree,
		FBavail:  statfs.Bavail,
		FFiles:   statfs.Files,
		FFfree:   statfs.Ffree,
		FFlag:    uint64(statfs.Flags),
		FNamemax: uint64(statfs.Namelen),
	}
}
func ThriftStatvfsStructBuilder(statfs *syscall.Statfs_t) any {
	return &icefsthrift.StatvfsStruct{
		FBsize:   icefsthrift.Ui64(statfs.Bsize),
		FFrsize:  icefsthrift.Ui64(statfs.Frsize),
		FBlocks:  icefsthrift.Ui64(statfs.Blocks),
		FBfree:   icefsthrift.Ui64(statfs.Ffree),
		FBavail:  icefsthrift.Ui64(statfs.Bavail),
		FFiles:   icefsthrift.Ui64(statfs.Files),
		FFfree:   icefsthrift.Ui64(statfs.Ffree),
		FFlag:    icefsthrift.Ui64(statfs.Flags),
		FNamemax: icefsthrift.Ui64(statfs.Namelen),
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

func GRpcSyscallStatBuilder(stat *pb.StatStruct) *syscall.Stat_t {
	return &syscall.Stat_t{
		Dev:     stat.StDev,
		Ino:     stat.StIno,
		Nlink:   stat.StNlink,
		Mode:    stat.StMode,
		Uid:     stat.StUid,
		Gid:     stat.StGid,
		Rdev:    stat.StRdev,
		Size:    stat.StSize,
		Blksize: stat.StBlksize,
		Blocks:  stat.StBlocks,
		Atim: syscall.Timespec{
			Sec:  stat.StAtim.TimeSec,
			Nsec: stat.StAtim.TimeNSec,
		},
		Mtim: syscall.Timespec{
			Sec:  stat.StMtim.TimeSec,
			Nsec: stat.StMtim.TimeNSec,
		},
		Ctim: syscall.Timespec{
			Sec:  stat.StCtim.TimeSec,
			Nsec: stat.StCtim.TimeNSec,
		},
	}
}

func ThriftSyscallStatBuilder(stat *icefsthrift.StatStruct) *syscall.Stat_t {
	return &syscall.Stat_t{
		Dev:     uint64(stat.StDev),
		Ino:     uint64(stat.StIno),
		Nlink:   uint64(stat.StNlink),
		Mode:    uint32(stat.StMode),
		Uid:     uint32(stat.StUID),
		Gid:     uint32(stat.StGid),
		Rdev:    uint64(stat.StRdev),
		Size:    stat.StSize,
		Blksize: stat.StBlksize,
		Blocks:  stat.StBlocks,
		Atim: syscall.Timespec{
			Sec:  stat.StAtim.TimeSec,
			Nsec: stat.StAtim.TimeNSec,
		},
		Mtim: syscall.Timespec{
			Sec:  stat.StMtim.TimeSec,
			Nsec: stat.StMtim.TimeNSec,
		},
		Ctim: syscall.Timespec{
			Sec:  stat.StCtim.TimeSec,
			Nsec: stat.StCtim.TimeNSec,
		},
	}
}

// 有锁操作
func (s *IcefsServer) doIcefsLookUp(parentFakeInode uint64, name string, fuseEntryParamBuilder FuseEntryParamBuilder) (entry any, err error) {
	inode := new(IcefsInode)
	s.inodeCacheLock.RLock()
	parentInode := s.getIcefsInode(parentFakeInode)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		err = syscall.Errno(icefserror.ICEFS_BUG_ERR)
		return
	}

	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	newfd, err := syscall.Openat(parentInode.fd, name, unix.O_PATH|unix.O_NOFOLLOW, 0)
	parentInode.inodeLock.RUnlock()
	if err != nil {
		return
	}
	var stat unix.Stat_t
	err = unix.Fstatat(newfd, "", &stat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		return
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
	entry = fuseEntryParamBuilder(inode, s.timeout)
	inode.inodeLock.RUnlock()

	return
}

func (s *IcefsServer) doIcefsForget(inodeNum uint64, nlookup uint64) {
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

// 提升性能
func checkNameIsDotOrDotDot(name string) bool {
	return name == "." || name == ".."
}

func getAlignedMemNormal(size uint64, alignSize uint64) []byte {
	mem := make([]byte, size+alignSize)
	offset := (uint64(uintptr(unsafe.Pointer(&mem[0]))) & (alignSize - 1))
	return mem[offset : offset+size]
}

func (s *IcefsServer) getAlignedMem(size uint64) []byte {
	mem := make([]byte, size+s.logicalBlockSize)
	offset := (uint64(uintptr(unsafe.Pointer(&mem[0]))) & (s.logicalBlockSize - 1))
	return mem[offset : offset+size]
}

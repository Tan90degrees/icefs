/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 14:35:52
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsSetAttr.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"fmt"
	"icefs-server/icefserror"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"math"
	"syscall"

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) doIcefsSetAttr(fakeInode uint64, reqStat *syscall.Stat_t, toSet int32, fh uint64, hasFh bool, statStructBuilder StatStructBuilder) (status int32, resStat any) {
	var err error
	var valid int32
	var unixStat unix.Stat_t

	status = icefserror.ICEFS_EOK
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	valid = toSet
	if (valid & FUSE_SET_ATTR_MODE) != 0 {
		if hasFh {
			err = syscall.Fchmod(int(fh), reqStat.Mode)
		} else {
			procName := fmt.Sprintf("/proc/self/fd/%v", inode.fd)
			err = syscall.Chmod(procName, reqStat.Mode)
		}
		if err != nil {
			status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	if (valid & (FUSE_SET_ATTR_UID | FUSE_SET_ATTR_GID)) != 0 {
		var uid uint32
		var gid uint32
		if (valid & FUSE_SET_ATTR_UID) != 0 {
			uid = reqStat.Uid
		} else {
			uid = math.MaxUint32
		}
		if (valid & FUSE_SET_ATTR_GID) != 0 {
			gid = reqStat.Gid
		} else {
			gid = math.MaxUint32
		}

		err = syscall.Fchownat(inode.fd, "", int(uid), int(gid), unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
		if err != nil {
			status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	if (valid & FUSE_SET_ATTR_SIZE) != 0 {
		if hasFh {
			err = syscall.Ftruncate(int(fh), reqStat.Size)
		} else {
			procName := fmt.Sprintf("/proc/self/fd/%v", inode.fd)
			err = syscall.Truncate(procName, reqStat.Size)
		}
		if err != nil {
			status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	if (valid & (FUSE_SET_ATTR_ATIME | FUSE_SET_ATTR_MTIME)) != 0 {

		ts := []syscall.Timespec{
			{Sec: 0, Nsec: unix.UTIME_OMIT},
			{Sec: 0, Nsec: unix.UTIME_OMIT},
		}

		if (valid & FUSE_SET_ATTR_ATIME_NOW) != 0 {
			ts[0].Nsec = unix.UTIME_NOW
		} else if (valid & FUSE_SET_ATTR_ATIME) != 0 {
			ts[0] = reqStat.Atim
		}

		if (valid & FUSE_SET_ATTR_MTIME_NOW) != 0 {
			ts[1].Nsec = unix.UTIME_NOW
		} else if (valid & FUSE_SET_ATTR_MTIME) != 0 {
			ts[1] = reqStat.Mtim
		}

		if hasFh {
			procName := fmt.Sprintf("/proc/self/fd/%v", fh)
			err = syscall.UtimesNano(procName, ts)
		} else {
			procName := fmt.Sprintf("/proc/self/fd/%v", inode.fd)
			err = syscall.UtimesNano(procName, ts)
		}
		if err != nil {
			status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	err = unix.Fstatat(inode.fd, "", &unixStat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		goto errOut
	}
	UnixStatFillSyscallStat(&inode.stat, &unixStat)

	resStat = statStructBuilder(&inode.stat)
	status = icefserror.ICEFS_EOK

errOut:
	inode.inodeLock.Unlock()
	return
}

func (s *IcefsGRpcServer) DoIcefsSetAttr(ctx context.Context, req *pb.IcefsSetAttrReq) (*pb.IcefsSetAttrRes, error) {
	var res pb.IcefsSetAttrRes
	var stat any

	res.Status, stat = s.server.doIcefsSetAttr(req.Inode, GRpcSyscallStatBuilder(req.Stat), req.ToSet, req.Fh, req.HasFh, GRpcStatStructBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Stat = stat.(*pb.StatStruct)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsSetAttr(ctx context.Context, req *icefsthrift.IcefsSetAttrReq) (*icefsthrift.IcefsSetAttrRes, error) {
	var res icefsthrift.IcefsSetAttrRes
	var stat any

	res.Status, stat = s.server.doIcefsSetAttr(uint64(req.Inode), ThriftSyscallStatBuilder(req.Stat), req.ToSet, uint64(req.Fh), req.HasFh, ThriftStatStructBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Stat = stat.(*icefsthrift.StatStruct)
	}

	return &res, nil
}

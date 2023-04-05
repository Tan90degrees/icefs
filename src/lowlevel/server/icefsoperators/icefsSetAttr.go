/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:11:36
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
	pb "icefs-server/icefsrpc"
	"math"
	"syscall"

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) DoIcefsSetAttr(ctx context.Context, req *pb.IcefsSetAttrReq) (*pb.IcefsSetAttrRes, error) {
	var res pb.IcefsSetAttrRes
	var err error
	var valid int32
	res.Status = icefserror.ICEFS_EOK
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	valid = req.ToSet
	if (valid & FUSE_SET_ATTR_MODE) != 0 {
		if req.HasFh {
			err = syscall.Fchmod(int(req.Fh), req.Stat.StMode)
		} else {
			procName := fmt.Sprintf("/proc/self/fd/%v", inode.fd)
			err = syscall.Chmod(procName, req.Stat.StMode)
		}
		if err != nil {
			res.Status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	if (valid & (FUSE_SET_ATTR_UID | FUSE_SET_ATTR_GID)) != 0 {
		var uid uint32
		var gid uint32
		if (valid & FUSE_SET_ATTR_UID) != 0 {
			uid = req.Stat.StUid
		} else {
			uid = math.MaxUint32
		}
		if (valid & FUSE_SET_ATTR_GID) != 0 {
			gid = req.Stat.StGid
		} else {
			gid = math.MaxUint32
		}

		err = syscall.Fchownat(inode.fd, "", int(uid), int(gid), unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
		if err != nil {
			res.Status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	if (valid & FUSE_SET_ATTR_SIZE) != 0 {
		if req.HasFh {
			err = syscall.Ftruncate(int(req.Fh), req.Stat.StSize)
		} else {
			procName := fmt.Sprintf("/proc/self/fd/%v", inode.fd)
			err = syscall.Truncate(procName, req.Stat.StSize)
		}
		if err != nil {
			res.Status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	if (valid & (FUSE_SET_ATTR_ATIME | FUSE_SET_ATTR_MTIME)) != 0 {

		ts := []unix.Timespec{
			{Sec: 0, Nsec: unix.UTIME_OMIT},
			{Sec: 0, Nsec: unix.UTIME_OMIT},
		}

		if (valid & FUSE_SET_ATTR_ATIME_NOW) != 0 {
			ts[0].Nsec = unix.UTIME_NOW
		} else if (valid & FUSE_SET_ATTR_ATIME) != 0 {
			ts[0].Sec = req.Stat.StAtim.TimeSec
			ts[0].Nsec = req.Stat.StAtim.TimeNSec
		}

		if (valid & FUSE_SET_ATTR_MTIME_NOW) != 0 {
			ts[1].Nsec = unix.UTIME_NOW
		} else if (valid & FUSE_SET_ATTR_MTIME) != 0 {
			ts[1].Sec = req.Stat.StMtim.TimeSec
			ts[1].Nsec = req.Stat.StMtim.TimeNSec
		}

		if req.HasFh {
			procName := fmt.Sprintf("/proc/self/fd/%v", req.Fh)
			err = unix.UtimesNanoAt(unix.AT_FDCWD, procName, ts, 0)
		} else {
			procName := fmt.Sprintf("/proc/self/fd/%v", inode.fd)
			err = unix.UtimesNanoAt(unix.AT_FDCWD, procName, ts, 0)
		}
		if err != nil {
			res.Status = icefserror.IcefsStdErrno(err)
			goto errOut
		}
	}

	err = s.doGetAttr(inode.fd, &inode.stat)
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	res.Stat = StatStructBuilder(&inode.stat)
	res.Status = icefserror.ICEFS_EOK

errOut:
	inode.inodeLock.Unlock()
	return &res, nil
}

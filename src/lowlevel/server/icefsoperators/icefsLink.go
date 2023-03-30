/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-19 15:29:52
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:42
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsLink.go
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

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) DoIcefsLink(ctx context.Context, req *pb.IcefsLinkReq) (*pb.IcefsLinkRes, error) {
	var res pb.IcefsLinkRes
	var stat unix.Stat_t
	var inode *IcefsInode
	var parentInode *IcefsInode
	var procName string
	var err error
	s.inodeCacheLock.RLock()
	inode = s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	parentInode = s.getIcefsInode(req.NewParentInode)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	inode.inodeLock.Lock()
	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	err = unix.Linkat(unix.AT_FDCWD, procName, parentInode.fd, req.NewName, unix.AT_SYMLINK_FOLLOW)
	parentInode.inodeLock.RUnlock()
	if err != nil {
		inode.inodeLock.Unlock()
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	err = unix.Fstatat(inode.fd, "", &stat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		inode.inodeLock.Unlock()
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	UnixStatFillSyscallStat(&inode.stat, &stat)
	inode.nlookup++
	res.Entry = FuseEntryParamBuilder(inode, s.timeout)
	inode.inodeLock.Unlock()
	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

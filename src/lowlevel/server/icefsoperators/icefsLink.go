/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-19 15:29:52
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 07:16:27
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
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) doIcefsLink(fakeInode uint64, newParentFakeInode uint64, newName string, fuseEntryParamBuilder FuseEntryParamBuilder) (status int32, entry any) {
	var stat unix.Stat_t
	var inode *IcefsInode
	var newParentInode *IcefsInode
	var procName string
	var err error
	s.inodeCacheLock.RLock()
	inode = s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	newParentInode = s.getIcefsInode(newParentFakeInode)
	if newParentInode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	inode.inodeLock.Lock()
	newParentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	err = unix.Linkat(unix.AT_FDCWD, procName, newParentInode.fd, newName, unix.AT_SYMLINK_FOLLOW)
	newParentInode.inodeLock.RUnlock()
	if err != nil {
		inode.inodeLock.Unlock()
		status = icefserror.IcefsStdErrno(err)
		return
	}

	err = unix.Fstatat(inode.fd, "", &stat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		inode.inodeLock.Unlock()
		status = icefserror.IcefsStdErrno(err)
		return
	}

	UnixStatFillSyscallStat(&inode.stat, &stat)
	inode.nlookup++
	entry = fuseEntryParamBuilder(inode, s.timeout)
	inode.inodeLock.Unlock()
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsLink(ctx context.Context, req *pb.IcefsLinkReq) (*pb.IcefsLinkRes, error) {
	var res pb.IcefsLinkRes
	var entry any

	res.Status, entry = s.server.doIcefsLink(req.Inode, req.NewParentInode, req.NewName, GRpcFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*pb.FuseEntryParam)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsLink(ctx context.Context, req *icefsthrift.IcefsLinkReq) (*icefsthrift.IcefsLinkRes, error) {
	var res icefsthrift.IcefsLinkRes
	var entry any

	res.Status, entry = s.server.doIcefsLink(uint64(req.Inode), uint64(req.NewParentInode_), req.NewName_, ThriftFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*icefsthrift.FuseEntryParam)
	}

	return &res, nil
}

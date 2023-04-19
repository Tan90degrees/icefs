/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 17:31:45
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsCreate.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"syscall"
)

func (s *IcefsServer) doIcefsCreate(parentInode uint64, name string, mode uint32, flags int32, fuseEntryParamBuilder FuseEntryParamBuilder) (status int32, fh uint64, entry any) {
	var newFd int
	var err error

	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(parentInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	newFd, err = syscall.Openat(inode.fd, name, int((flags|syscall.O_CREAT)&(^syscall.O_NOFOLLOW)), mode)
	inode.inodeLock.RUnlock()
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}
	fh = uint64(newFd)
	entry, err = s.doIcefsLookUp(parentInode, name, fuseEntryParamBuilder)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}

	return
}

func (s *IcefsServer) doIcefsOpenAfterCreate(fakeInode uint64) (status int32) {
	var newInode *IcefsInode
	s.inodeCacheLock.RLock()
	newInode = s.getIcefsInode(fakeInode)
	if newInode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	newInode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	newInode.nopen++
	newInode.inodeLock.Unlock()
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsCreate(ctx context.Context, req *pb.IcefsCreateReq) (*pb.IcefsCreateRes, error) {
	var res pb.IcefsCreateRes
	var entry any

	res.Status, res.Fh, entry = s.server.doIcefsCreate(req.ParentInode, req.Name, req.Mode, req.Flags, GRpcFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*pb.FuseEntryParam)
		res.Status = s.server.doIcefsOpenAfterCreate(res.Entry.Inode)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsCreate(ctx context.Context, req *icefsthrift.IcefsCreateReq) (*icefsthrift.IcefsCreateRes, error) {
	var res icefsthrift.IcefsCreateRes
	var fh uint64
	var entry any

	res.Status, fh, entry = s.server.doIcefsCreate(uint64(req.ParentInode), req.Name, uint32(req.Mode), req.Flags, ThriftFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Fh = icefsthrift.Ui64(fh)
		res.Entry = entry.(*icefsthrift.FuseEntryParam)
		res.Status = s.server.doIcefsOpenAfterCreate(uint64(res.Entry.Inode))
	}

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 08:00:53
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpenDir.go
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

func (s *IcefsServer) doIcefsOpenDir(fakeInode uint64) (status int32, fh uint64) {
	var fd int
	var err error
	var dir IcefsDir
	var errno int
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	fd, err = syscall.Openat(inode.fd, ".", syscall.O_RDONLY, 0)
	inode.inodeLock.RUnlock()
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}

	dir.dirStream, errno = IcefsFdOpenDir(fd)
	if dir.dirStream == nil {
		status = int32(errno)
		return
	}

	dir.offset = 0
	fh = uint64(uintptr(dir.dirStream))
	s.dirCacheLock.Lock()
	s.putIcefsDir(fh, &dir)
	s.dirCacheLock.Unlock()
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsOpenDir(ctx context.Context, req *pb.IcefsOpenDirReq) (*pb.IcefsOpenDirRes, error) {
	var res pb.IcefsOpenDirRes

	res.Status, res.Fh = s.server.doIcefsOpenDir(req.Inode)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsOpenDir(ctx context.Context, req *icefsthrift.IcefsOpenDirReq) (*icefsthrift.IcefsOpenDirRes, error) {
	var res icefsthrift.IcefsOpenDirRes
	var fh uint64

	res.Status, fh = s.server.doIcefsOpenDir(uint64(req.Inode))
	res.Fh = icefsthrift.Ui64(fh)

	return &res, nil
}

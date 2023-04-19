/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 14:58:59
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsStatFS.go
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

func (s *IcefsServer) doIcefsStatFS(fakeInode uint64, statvfsStructBuilder StatvfsStructBuilder) (status int32, statvfs any) {
	var err error
	var statfs syscall.Statfs_t
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = syscall.Fstatfs(inode.fd, &statfs)
	inode.inodeLock.RUnlock()
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}
	status = icefserror.ICEFS_EOK
	statvfs = statvfsStructBuilder(&statfs)
	return
}

func (s *IcefsGRpcServer) DoIcefsStatFS(ctx context.Context, req *pb.IcefsStatFSReq) (*pb.IcefsStatFSRes, error) {
	var res pb.IcefsStatFSRes
	var statvfs any

	res.Status, statvfs = s.server.doIcefsStatFS(req.Inode, GRpcStatvfsStructBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Statvfs = statvfs.(*pb.StatvfsStruct)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsStatFS(ctx context.Context, req *icefsthrift.IcefsStatFSReq) (*icefsthrift.IcefsStatFSRes, error) {
	var res icefsthrift.IcefsStatFSRes
	var statvfs any

	res.Status, statvfs = s.server.doIcefsStatFS(uint64(req.Inode), ThriftStatvfsStructBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Statvfs = statvfs.(*icefsthrift.StatvfsStruct)
	}

	return &res, nil
}

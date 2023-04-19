/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 07:16:00
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsMkDir.go
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

func (s *IcefsServer) doIcefsMkDir(parentFakeInode uint64, name string, mode uint32, fuseEntryParamBuilder FuseEntryParamBuilder) (status int32, entry any) {
	var err error
	s.inodeCacheLock.RLock()
	parentInode := s.getIcefsInode(parentFakeInode)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = syscall.Mkdirat(parentInode.fd, name, mode)
	parentInode.inodeLock.RUnlock()
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}

	entry, err = s.doIcefsLookUp(parentFakeInode, name, fuseEntryParamBuilder)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}

	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsMkDir(ctx context.Context, req *pb.IcefsMkDirReq) (*pb.IcefsMkDirRes, error) {
	var res pb.IcefsMkDirRes
	var entry any

	res.Status, entry = s.server.doIcefsMkDir(req.ParentInode, req.Name, req.Mode, GRpcFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*pb.FuseEntryParam)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsMkDir(ctx context.Context, req *icefsthrift.IcefsMkDirReq) (*icefsthrift.IcefsMkDirRes, error) {
	var res icefsthrift.IcefsMkDirRes
	var entry any

	res.Status, entry = s.server.doIcefsMkDir(uint64(req.ParentInode), req.Name, uint32(req.Mode), ThriftFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*icefsthrift.FuseEntryParam)
	}

	return &res, nil
}

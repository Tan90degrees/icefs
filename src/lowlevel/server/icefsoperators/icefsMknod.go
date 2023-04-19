/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 07:50:13
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsMknod.go
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

func (s *IcefsServer) doIcefsMknod(parentFakeInode uint64, name string, mode uint32, rdev uint64, fuseEntryParamBuilder FuseEntryParamBuilder) (status int32, entry any) {
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
	err = syscall.Mknodat(parentInode.fd, name, mode, int(rdev))
	parentInode.inodeLock.Unlock()
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}
	entry, err = s.doIcefsLookUp(parentFakeInode, name, fuseEntryParamBuilder)
	status = icefserror.IcefsStdErrno(err)

	return
}

func (s *IcefsGRpcServer) DoIcefsMknod(ctx context.Context, req *pb.IcefsMknodReq) (*pb.IcefsMknodRes, error) {
	var res pb.IcefsMknodRes
	var entry any

	res.Status, entry = s.server.doIcefsMknod(req.ParentInode, req.Name, req.Mode, req.Rdev, GRpcFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*pb.FuseEntryParam)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsMknod(ctx context.Context, req *icefsthrift.IcefsMknodReq) (*icefsthrift.IcefsMknodRes, error) {
	var res icefsthrift.IcefsMknodRes
	var entry any

	res.Status, entry = s.server.doIcefsMknod(uint64(req.ParentInode), req.Name, uint32(req.Mode), uint64(req.Rdev), ThriftFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*icefsthrift.FuseEntryParam)
	}

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 13:27:02
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRelease.go
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

func (s *IcefsServer) doIcefsRelease(fakeInode uint64, fh uint64) (status int32) {
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	inode.nopen--
	inode.inodeLock.Unlock()
	syscall.Close(int(fh))
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsRelease(ctx context.Context, req *pb.IcefsReleaseReq) (*pb.IcefsReleaseRes, error) {
	var res pb.IcefsReleaseRes

	res.Status = s.server.doIcefsRelease(req.Inode, req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsRelease(ctx context.Context, req *icefsthrift.IcefsReleaseReq) (*icefsthrift.IcefsReleaseRes, error) {
	var res icefsthrift.IcefsReleaseRes

	res.Status = s.server.doIcefsRelease(uint64(req.Inode), uint64(req.Fh))

	return &res, nil
}

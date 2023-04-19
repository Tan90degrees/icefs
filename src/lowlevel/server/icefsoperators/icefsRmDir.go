/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 13:55:12
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRmDir.go
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

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) doIcefsRmDir(parentFakeInode uint64, name string) (status int32) {
	var err error

	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(parentFakeInode)
	if inode == nil {
		s.inodeCacheLock.Unlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	inode.inodeLock.RLock()
	err = unix.Unlinkat(inode.fd, name, unix.AT_REMOVEDIR)
	inode.inodeLock.RUnlock()
	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsRmDir(ctx context.Context, req *pb.IcefsRmDirReq) (*pb.IcefsRmDirRes, error) {
	var res pb.IcefsRmDirRes

	res.Status = s.server.doIcefsRmDir(req.ParentInode, req.Name)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsRmDir(ctx context.Context, req *icefsthrift.IcefsRmDirReq) (*icefsthrift.IcefsRmDirRes, error) {
	var res icefsthrift.IcefsRmDirRes

	res.Status = s.server.doIcefsRmDir(uint64(req.ParentInode), req.Name)

	return &res, nil
}

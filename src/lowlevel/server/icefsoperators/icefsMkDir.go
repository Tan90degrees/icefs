/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:52
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsMkDir.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsMkDir(ctx context.Context, req *pb.IcefsMkDirReq) (*pb.IcefsMkDirRes, error) {
	var res pb.IcefsMkDirRes
	var err error
	s.inodeCacheLock.RLock()
	parentInode := s.getIcefsInode(req.ParentInode)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = syscall.Mkdirat(parentInode.fd, req.Name, req.Mode)
	parentInode.inodeLock.RUnlock()
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	res.Entry, err = s.doLookUp(req.ParentInode, req.Name)
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

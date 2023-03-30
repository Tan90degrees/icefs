/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:38
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRmDir.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) DoIcefsRmDir(ctx context.Context, req *pb.IcefsRmDirReq) (*pb.IcefsRmDirRes, error) {
	var res pb.IcefsRmDirRes
	var err error

	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.ParentInode)
	if inode == nil {
		s.inodeCacheLock.Unlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.RLock()
	err = unix.Unlinkat(inode.fd, req.Name, unix.AT_REMOVEDIR)
	inode.inodeLock.RUnlock()
	res.Status = icefserror.IcefsStdErrno(err)

errOut:
	return &res, nil
}

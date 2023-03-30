/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:24
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRelease.go
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

func (s *IcefsServer) DoIcefsRelease(ctx context.Context, req *pb.IcefsReleaseReq) (*pb.IcefsReleaseRes, error) {
	var res pb.IcefsReleaseRes
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	inode.nopen--
	inode.inodeLock.Unlock()
	syscall.Close(int(req.FileInfo.Fh))
	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

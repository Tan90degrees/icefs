/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:57
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsSymLink.go
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

func (s *IcefsServer) DoIcefsSymLink(ctx context.Context, req *pb.IcefsSymLinkReq) (*pb.IcefsSymLinkRes, error) {
	var res pb.IcefsSymLinkRes
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.ParentInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	res.Status = icefserror.IcefsStdErrno(unix.Symlinkat(req.Link, inode.fd, req.Name))
	inode.inodeLock.RUnlock()

	res.Entry, err = s.doLookUp(req.ParentInode, req.Name)
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

errOut:
	return &res, nil
}

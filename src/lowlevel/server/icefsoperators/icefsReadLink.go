/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:21
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReadLink.go
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

func (s *IcefsServer) DoIcefsReadLink(ctx context.Context, req *pb.IcefsReadLinkReq) (*pb.IcefsReadLinkRes, error) {
	var res pb.IcefsReadLinkRes
	var buf []byte
	var readN int
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	buf = make([]byte, unix.PathMax+1)

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	readN, err = unix.Readlinkat(inode.fd, "", buf)
	inode.inodeLock.RUnlock()
	if err == nil && readN == len(buf) {
		res.Path = string(buf)
	}
	res.Status = icefserror.IcefsStdErrno(err)

errOut:
	return &res, nil
}

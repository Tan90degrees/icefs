/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 11:54:31
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReadLink.go
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

func (s *IcefsServer) doIcefsReadLink(fakeInode uint64) (status int32, path []byte, readN int) {
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	path = make([]byte, unix.PathMax+1)

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	readN, err = unix.Readlinkat(inode.fd, "", path)
	inode.inodeLock.RUnlock()
	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsReadLink(ctx context.Context, req *pb.IcefsReadLinkReq) (*pb.IcefsReadLinkRes, error) {
	var res pb.IcefsReadLinkRes
	var path []byte
	var readN int

	res.Status, path, readN = s.server.doIcefsReadLink(req.Inode)
	if (res.Status == icefserror.ICEFS_EOK) && readN == len(path) {
		res.Path = string(path)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsReadLink(ctx context.Context, req *icefsthrift.IcefsReadLinkReq) (*icefsthrift.IcefsReadLinkRes, error) {
	var res icefsthrift.IcefsReadLinkRes
	var path []byte
	var readN int

	res.Status, path, readN = s.server.doIcefsReadLink(uint64(req.Inode))
	if (res.Status == icefserror.ICEFS_EOK) && readN == len(path) {
		res.Path = string(path)
	}

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-05-10 09:20:35
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsSymLink.go
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

func (s *IcefsServer) doIcefsSymLink(parentFakeInode uint64, link string, name string, fuseEntryParamBuilder FuseEntryParamBuilder) (status int32, entry any) {
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(parentFakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	status = icefserror.IcefsStdErrno(unix.Symlinkat(link, inode.fd, name))
	inode.inodeLock.RUnlock()

	entry, err = s.doIcefsLookUp(parentFakeInode, name, fuseEntryParamBuilder)
	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsSymLink(ctx context.Context, req *pb.IcefsSymLinkReq) (*pb.IcefsSymLinkRes, error) {
	var res pb.IcefsSymLinkRes
	var entry any

	res.Status, entry = s.server.doIcefsSymLink(req.ParentInode, req.Link, req.Name, GRpcFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*pb.FuseEntryParam)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsSymLink(ctx context.Context, req *icefsthrift.IcefsSymLinkReq) (*icefsthrift.IcefsSymLinkRes, error) {
	var res icefsthrift.IcefsSymLinkRes
	var entry any

	res.Status, entry = s.server.doIcefsSymLink(uint64(req.ParentInode), req.Link, req.Name, ThriftFuseEntryParamBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Entry = entry.(*icefsthrift.FuseEntryParam)
	}

	return &res, nil
}

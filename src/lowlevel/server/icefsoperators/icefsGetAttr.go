/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-19 15:29:43
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 14:31:12
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsGetAttr.go
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

func (s *IcefsServer) doIcefsGetAttr(fakeInode uint64, statStructBuilder StatStructBuilder) (status int32, attr any) {
	var err error
	var unixStat unix.Stat_t

	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	err = unix.Fstatat(inode.fd, "", &unixStat, unix.AT_EMPTY_PATH|unix.AT_SYMLINK_NOFOLLOW)
	if err != nil {
		inode.inodeLock.Unlock()
		status = icefserror.IcefsStdErrno(err)
		return
	}
	UnixStatFillSyscallStat(&inode.stat, &unixStat)

	attr = statStructBuilder(&inode.stat)
	inode.inodeLock.Unlock()
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsGetAttr(ctx context.Context, req *pb.IcefsGetAttrReq) (*pb.IcefsGetAttrRes, error) {
	var res pb.IcefsGetAttrRes
	var attr any

	res.Status, attr = s.server.doIcefsGetAttr(req.Inode, GRpcStatStructBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Stat = attr.(*pb.StatStruct)
	}

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsGetAttr(ctx context.Context, req *icefsthrift.IcefsGetAttrReq) (*icefsthrift.IcefsGetAttrRes, error) {
	var res icefsthrift.IcefsGetAttrRes
	var attr any

	res.Status, attr = s.server.doIcefsGetAttr(uint64(req.Inode), ThriftStatStructBuilder)
	if res.Status == icefserror.ICEFS_EOK {
		res.Stat = attr.(*icefsthrift.StatStruct)
	}

	return &res, nil
}

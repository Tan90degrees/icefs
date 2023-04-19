/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 07:58:23
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpen.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"fmt"
	"icefs-server/icefserror"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"syscall"
)

func (s *IcefsServer) doIcefsOpen(fakeInode uint64, flags int32) (status int32, fh uint64) {
	var procName string
	var fd int
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	fd, err = syscall.Open(procName, int(flags&(^syscall.O_NOFOLLOW)), 0)
	if err != nil {
		inode.inodeLock.Unlock()
		status = icefserror.IcefsStdErrno(err)
		return
	}
	inode.nopen++
	inode.inodeLock.Unlock()
	status = icefserror.ICEFS_EOK
	fh = uint64(fd)
	return
}

func (s *IcefsGRpcServer) DoIcefsOpen(ctx context.Context, req *pb.IcefsOpenReq) (*pb.IcefsOpenRes, error) {
	var res pb.IcefsOpenRes

	res.Status, res.Fh = s.server.doIcefsOpen(req.Inode, req.Flags)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsOpen(ctx context.Context, req *icefsthrift.IcefsOpenReq) (*icefsthrift.IcefsOpenRes, error) {
	var res icefsthrift.IcefsOpenRes
	var fh uint64

	res.Status, fh = s.server.doIcefsOpen(uint64(req.Inode), req.Flags)
	res.Fh = icefsthrift.Ui64(fh)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 05:57:27
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsListXattr.go
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

func (s *IcefsServer) doIcefsListXattr(fakeInode uint64, reqSize uint64) (status int32, resSize int64, value []byte) {
	var procName string
	var size int
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	inode.inodeLock.RUnlock()
	value = make([]byte, reqSize)
	size, err = syscall.Listxattr(procName, value)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}
	status = icefserror.ICEFS_EOK
	resSize = int64(size)
	return
}

func (s *IcefsGRpcServer) DoIcefsListXattr(ctx context.Context, req *pb.IcefsListXattrReq) (*pb.IcefsListXattrRes, error) {
	var res pb.IcefsListXattrRes

	res.Status, res.Size, res.Value = s.server.doIcefsListXattr(req.Inode, req.Size)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsListXattr(ctx context.Context, req *icefsthrift.IcefsListXattrReq) (*icefsthrift.IcefsListXattrRes, error) {
	var res icefsthrift.IcefsListXattrRes

	res.Status, res.Size, res.Value = s.server.doIcefsListXattr(uint64(req.Inode), uint64(req.Size))

	return &res, nil
}

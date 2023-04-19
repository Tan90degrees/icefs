/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 03:50:05
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsGetXattr.go
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

func (s *IcefsServer) doIcefsGetXattr(fakeInode uint64, name string, size uint64) (status int32, valueSize int64, value []byte) {
	var procName string
	var err error
	var sz int
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
	value = make([]byte, size)
	sz, err = syscall.Getxattr(procName, name, value)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}
	status = icefserror.ICEFS_EOK
	valueSize = int64(sz)
	return
}

func (s *IcefsGRpcServer) DoIcefsGetXattr(ctx context.Context, req *pb.IcefsGetXattrReq) (*pb.IcefsGetXattrRes, error) {
	var res pb.IcefsGetXattrRes

	res.Status, res.Size, res.Value = s.server.doIcefsGetXattr(req.Inode, req.Name, req.Size)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsGetXattr(ctx context.Context, req *icefsthrift.IcefsGetXattrReq) (*icefsthrift.IcefsGetXattrRes, error) {
	var res icefsthrift.IcefsGetXattrRes

	res.Status, res.Size, res.Value = s.server.doIcefsGetXattr(uint64(req.Inode), req.Name, uint64(req.Size))

	return &res, nil
}

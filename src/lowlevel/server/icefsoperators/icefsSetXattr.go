/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 14:42:37
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsSetXattr.go
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

func (s *IcefsServer) doIcefsSetXattr(fakeInode uint64, name string, value string, flags int32) (status int32) {
	var procName string
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
	status = icefserror.IcefsStdErrno(syscall.Setxattr(procName, name, []byte(value), int(flags)))
	return
}

func (s *IcefsGRpcServer) DoIcefsSetXattr(ctx context.Context, req *pb.IcefsSetXattrReq) (*pb.IcefsSetXattrRes, error) {
	var res pb.IcefsSetXattrRes

	res.Status = s.server.doIcefsSetXattr(req.Inode, req.Name, req.Value, req.Flags)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsSetXattr(ctx context.Context, req *icefsthrift.IcefsSetXattrReq) (*icefsthrift.IcefsSetXattrRes, error) {
	var res icefsthrift.IcefsSetXattrRes

	res.Status = s.server.doIcefsSetXattr(uint64(req.Inode), req.Name, req.Value, req.Flags)

	return &res, nil
}

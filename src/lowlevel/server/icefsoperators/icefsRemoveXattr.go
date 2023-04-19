/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 13:31:45
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRemoveXattr.go
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

func (s *IcefsServer) doIcefsRemoveXattr(fakeInode uint64, name string) (status int32) {
	var procName string
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
	err = syscall.Removexattr(procName, name)
	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsRemoveXattr(ctx context.Context, req *pb.IcefsRemoveXattrReq) (*pb.IcefsRemoveXattrRes, error) {
	var res pb.IcefsRemoveXattrRes

	res.Status = s.server.doIcefsRemoveXattr(req.Inode, req.Name)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsRemoveXattr(ctx context.Context, req *icefsthrift.IcefsRemoveXattrReq) (*icefsthrift.IcefsRemoveXattrRes, error) {
	var res icefsthrift.IcefsRemoveXattrRes

	res.Status = s.server.doIcefsRemoveXattr(uint64(req.Inode), req.Name)

	return &res, nil
}

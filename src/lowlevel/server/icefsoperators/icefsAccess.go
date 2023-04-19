/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 14:22:56
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsAccess.go
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

func (s *IcefsServer) doIcefsAccess(fakeInode uint64, mask int32) (status int32) {
	var procName string
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(fakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		return icefserror.ICEFS_BUG_ERR
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	status = icefserror.IcefsStdErrno(syscall.Access(procName, uint32(mask)))
	inode.inodeLock.RUnlock()
	return
}

func (s *IcefsGRpcServer) DoIcefsAccess(ctx context.Context, req *pb.IcefsAccessReq) (*pb.IcefsAccessRes, error) {
	var res pb.IcefsAccessRes

	res.Status = s.server.doIcefsAccess(req.Inode, req.Mask)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsAccess(ctx context.Context, req *icefsthrift.IcefsAccessReq) (*icefsthrift.IcefsAccessRes, error) {
	var res icefsthrift.IcefsAccessRes

	res.Status = s.server.doIcefsAccess(uint64(req.Inode), req.Mask)

	return &res, nil
}

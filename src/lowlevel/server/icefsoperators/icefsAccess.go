/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:26:57
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
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsAccess(ctx context.Context, req *pb.IcefsAccessReq) (*pb.IcefsAccessRes, error) {
	var res pb.IcefsAccessRes
	var procName string
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	res.Status = icefserror.IcefsStdErrno(syscall.Access(procName, uint32(req.Mask)))
	inode.inodeLock.RUnlock()
	// res.Status = int32(syscall.ENOSYS)

errOut:
	return &res, nil
}

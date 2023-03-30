/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:30
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
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsRemoveXattr(ctx context.Context, req *pb.IcefsRemoveXattrReq) (*pb.IcefsRemoveXattrRes, error) {
	var res pb.IcefsRemoveXattrRes
	var procName string
	var err error

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
	inode.inodeLock.RUnlock()
	err = syscall.Removexattr(procName, req.Name)
	res.Status = icefserror.IcefsStdErrno(err)

errOut:
	return &res, nil
}

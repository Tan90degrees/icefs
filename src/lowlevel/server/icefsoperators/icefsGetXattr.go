/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:31
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
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsGetXattr(ctx context.Context, req *pb.IcefsGetXattrReq) (*pb.IcefsGetXattrRes, error) {
	var res pb.IcefsGetXattrRes
	var procName string
	var err error
	var size int
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
	res.Value = make([]byte, req.Size)
	size, err = syscall.Getxattr(procName, req.Name, res.Value)
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}
	res.Status = icefserror.ICEFS_EOK
	res.Size = int64(size)

errOut:
	return &res, nil
}

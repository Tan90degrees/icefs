/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:56:16
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFsyncDir.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"syscall"
)

// TODO: Maybe not work
func (s *IcefsServer) DoIcefsFsyncDir(ctx context.Context, req *pb.IcefsFsyncDirReq) (*pb.IcefsFsyncDirRes, error) {
	var res pb.IcefsFsyncDirRes
	var fd int
	var err error

	s.dirCacheLock.RLock()
	dir := s.getIcefsDir(req.Fh)
	if dir == nil {
		s.dirCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	dir.dirLock.RLock()
	s.dirCacheLock.RUnlock()
	fd = IcefsDirFd(dir.dirStream)
	dir.dirLock.RUnlock()

	if req.DataSync != 0 {
		err = syscall.Fdatasync(fd)
	} else {
		err = syscall.Fsync(fd)
	}
	res.Status = icefserror.IcefsStdErrno(err)

errOut:
	return &res, nil
}

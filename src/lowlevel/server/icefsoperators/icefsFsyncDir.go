/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 16:44:07
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFsyncDir.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"syscall"
)

// TODO: Maybe not work
func (s *IcefsServer) doIcefsFsyncDir(dataSync int32, fh uint64) (status int32) {
	var fd int
	var err error

	s.dirCacheLock.RLock()
	dir := s.getIcefsDir(fh)
	if dir == nil {
		s.dirCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	dir.dirLock.RLock()
	s.dirCacheLock.RUnlock()
	fd = IcefsDirFd(dir.dirStream)
	dir.dirLock.RUnlock()

	if dataSync != 0 {
		err = syscall.Fdatasync(fd)
	} else {
		err = syscall.Fsync(fd)
	}
	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsFsyncDir(ctx context.Context, req *pb.IcefsFsyncDirReq) (*pb.IcefsFsyncDirRes, error) {
	var res pb.IcefsFsyncDirRes

	res.Status = s.server.doIcefsFsyncDir(req.DataSync, req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsFsyncDir(ctx context.Context, req *icefsthrift.IcefsFsyncDirReq) (*icefsthrift.IcefsFsyncDirRes, error) {
	var res icefsthrift.IcefsFsyncDirRes

	res.Status = s.server.doIcefsFsyncDir(req.DataSync, uint64(req.Fh))

	return &res, nil
}

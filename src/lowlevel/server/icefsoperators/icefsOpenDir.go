/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:00
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpenDir.go
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

func (s *IcefsServer) DoIcefsOpenDir(ctx context.Context, req *pb.IcefsOpenDirReq) (*pb.IcefsOpenDirRes, error) {
	var res pb.IcefsOpenDirRes
	var fd int
	var err error
	var dir IcefsDir
	var errno int
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	fd, err = syscall.Openat(inode.fd, ".", syscall.O_RDONLY, 0)
	inode.inodeLock.RUnlock()
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	dir.dirStream, errno = IcefsFdOpenDir(fd)
	if dir.dirStream == nil {
		res.Status = int32(errno)
		goto errOut
	}

	dir.offset = 0
	res.Fh = uint64(uintptr(dir.dirStream))
	s.dirCacheLock.Lock()
	s.putIcefsDir(res.Fh, &dir)
	s.dirCacheLock.Unlock()
	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

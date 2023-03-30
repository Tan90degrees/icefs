/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:54
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsStatFS.go
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

func (s *IcefsServer) DoIcefsStatFS(ctx context.Context, req *pb.IcefsStatFSReq) (*pb.IcefsStatFSRes, error) {
	var res pb.IcefsStatFSRes
	var err error
	var statfs syscall.Statfs_t
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = syscall.Fstatfs(inode.fd, &statfs)
	inode.inodeLock.RUnlock()
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}
	res.Status = icefserror.ICEFS_EOK
	res.Statvfs = &pb.StatvfsStruct{
		FBsize:   uint64(statfs.Bsize),
		FFrsize:  uint64(statfs.Frsize),
		FBlocks:  statfs.Blocks,
		FBfree:   statfs.Ffree,
		FBavail:  statfs.Bavail,
		FFiles:   statfs.Files,
		FFfree:   statfs.Ffree,
		FFlag:    uint64(statfs.Flags),
		FNamemax: uint64(statfs.Namelen),
	}

errOut:
	return &res, nil
}

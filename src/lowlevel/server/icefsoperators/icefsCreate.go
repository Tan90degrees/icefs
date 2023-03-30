/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:27:30
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsCreate.go
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

func (s *IcefsServer) DoIcefsCreate(ctx context.Context, req *pb.IcefsCreateReq) (*pb.IcefsCreateRes, error) {
	var res pb.IcefsCreateRes
	var newFd int
	var err error
	var entry *pb.FuseEntryParam
	var newInode *IcefsInode

	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.ParentInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	inode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	newFd, err = syscall.Openat(inode.fd, req.Name, int((req.FileInfo.Flags|syscall.O_CREAT)&(^syscall.O_NOFOLLOW)), req.Mode)
	inode.inodeLock.RUnlock()
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}
	res.Fh = uint64(newFd)
	entry, err = s.doLookUp(req.ParentInode, req.Name)
	if err != nil {
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	s.inodeCacheLock.RLock()
	newInode = s.getIcefsInode(entry.Inode)
	if newInode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	newInode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	newInode.nopen++
	newInode.inodeLock.Unlock()
	res.Entry = entry

errOut:
	return &res, nil
}

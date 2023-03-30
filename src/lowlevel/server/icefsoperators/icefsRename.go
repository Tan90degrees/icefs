/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:33
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRename.go
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

func (s *IcefsServer) DoIcefsRename(ctx context.Context, req *pb.IcefsRenameReq) (*pb.IcefsRenameRes, error) {
	var res pb.IcefsRenameRes
	var inode *IcefsInode
	var newInode *IcefsInode
	var err error

	if req.Flags != 0 {
		res.Status = int32(syscall.EINVAL)
		goto errOut
	}

	s.inodeCacheLock.RLock()
	inode = s.getIcefsInode(req.ParentInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	newInode = s.getIcefsInode(req.NewParentInode)
	if newInode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	inode.inodeLock.RLock()
	newInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = syscall.Renameat(inode.fd, req.Name, newInode.fd, req.NewName)
	inode.inodeLock.RUnlock()
	newInode.inodeLock.RUnlock()

	res.Status = icefserror.IcefsStdErrno(err)

errOut:
	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 13:44:09
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRename.go
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

func (s *IcefsServer) doIcefsRename(parentFakeInode uint64, name string, newParentFakeInode uint64, newName string, flags uint32) (status int32) {
	var inode *IcefsInode
	var newInode *IcefsInode
	var err error

	if flags != 0 {
		status = int32(syscall.EINVAL)
		return
	}

	s.inodeCacheLock.RLock()
	inode = s.getIcefsInode(parentFakeInode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	newInode = s.getIcefsInode(newParentFakeInode)
	if newInode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	inode.inodeLock.RLock()
	newInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = syscall.Renameat(inode.fd, name, newInode.fd, newName)
	inode.inodeLock.RUnlock()
	newInode.inodeLock.RUnlock()

	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsRename(ctx context.Context, req *pb.IcefsRenameReq) (*pb.IcefsRenameRes, error) {
	var res pb.IcefsRenameRes

	res.Status = s.server.doIcefsRename(req.ParentInode, req.Name, req.NewParentInode, req.NewName, req.Flags)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsRename(ctx context.Context, req *icefsthrift.IcefsRenameReq) (*icefsthrift.IcefsRenameRes, error) {
	var res icefsthrift.IcefsRenameRes

	res.Status = s.server.doIcefsRename(uint64(req.ParentInode), req.Name, uint64(req.NewParentInode_), req.NewName_, uint32(req.Flags))

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:59
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsUnlink.go
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

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) DoIcefsUnlink(ctx context.Context, req *pb.IcefsUnlinkReq) (*pb.IcefsUnlinkRes, error) {
	var res pb.IcefsUnlinkRes
	var err error
	var parentInode *IcefsInode
	if s.timeout == 0 {
		entry, err := s.doLookUp(req.ParentInode, req.Name)
		if err != nil {
			res.Status = int32(syscall.EIO)
		}
		if entry.Attr.StNlink == 1 {
			s.inodeCacheLock.RLock()
			inode := s.getIcefsInode(entry.Inode)
			if inode == nil {
				s.inodeCacheLock.RUnlock()
				res.Status = icefserror.ICEFS_BUG_ERR
				goto errOut
			}
			inode.inodeLock.Lock()
			s.inodeCacheLock.RUnlock()
			if inode.fd > 0 && inode.nopen == 0 {
				syscall.Close(inode.fd)
				inode.fd = -int(syscall.ENOENT)
				inode.generation++
			}
			inode.inodeLock.Unlock()
		}
		s.doForget(entry.Inode, 1)
	}

	s.inodeCacheLock.RLock()
	parentInode = s.getIcefsInode(req.ParentInode)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}

	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = unix.Unlinkat(parentInode.fd, req.Name, 0)
	parentInode.inodeLock.RUnlock()

	res.Status = icefserror.IcefsStdErrno(err)

errOut:
	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 17:20:03
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsUnlink.go
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

	"golang.org/x/sys/unix"
)

type IcefsDoUnlinkGetNlinkInode func(entry any) (stNlink uint64, inode uint64)

func GRpcIcefsDoUnlinkGetNlinkInode(entry any) (stNlink uint64, inode uint64) {
	return entry.(*pb.FuseEntryParam).Attr.StNlink, entry.(*pb.FuseEntryParam).Inode
}

func ThriftIcefsDoUnlinkGetNlinkInode(entry any) (stNlink uint64, inode uint64) {
	return uint64(entry.(*icefsthrift.FuseEntryParam).Attr.StNlink), uint64(entry.(*icefsthrift.FuseEntryParam).Inode)
}

func (s *IcefsServer) doIcefsUnlink(parentFakeInode uint64, name string, fuseEntryParamBuilder FuseEntryParamBuilder, icefsDoUnlinkGetNlinkInode IcefsDoUnlinkGetNlinkInode) (status int32) {
	var err error
	var parentInode *IcefsInode
	if s.timeout == 0 {
		entry, err := s.doIcefsLookUp(parentFakeInode, name, fuseEntryParamBuilder)
		if err != nil {
			status = int32(syscall.EIO)
		}
		stNlink, fakeInode := icefsDoUnlinkGetNlinkInode(entry)
		if stNlink == 1 {
			s.inodeCacheLock.RLock()
			inode := s.getIcefsInode(fakeInode)
			if inode == nil {
				s.inodeCacheLock.RUnlock()
				status = icefserror.ICEFS_BUG_ERR
				return
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
		s.doIcefsForget(fakeInode, 1)
	}

	s.inodeCacheLock.RLock()
	parentInode = s.getIcefsInode(parentFakeInode)
	if parentInode == nil {
		s.inodeCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}

	parentInode.inodeLock.RLock()
	s.inodeCacheLock.RUnlock()
	err = unix.Unlinkat(parentInode.fd, name, 0)
	parentInode.inodeLock.RUnlock()

	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsUnlink(ctx context.Context, req *pb.IcefsUnlinkReq) (*pb.IcefsUnlinkRes, error) {
	var res pb.IcefsUnlinkRes

	res.Status = s.server.doIcefsUnlink(req.ParentInode, req.Name, GRpcFuseEntryParamBuilder, GRpcIcefsDoUnlinkGetNlinkInode)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsUnlink(ctx context.Context, req *icefsthrift.IcefsUnlinkReq) (*icefsthrift.IcefsUnlinkRes, error) {
	var res icefsthrift.IcefsUnlinkRes

	res.Status = s.server.doIcefsUnlink(uint64(req.ParentInode), req.Name, ThriftFuseEntryParamBuilder, ThriftIcefsDoUnlinkGetNlinkInode)

	return &res, nil
}

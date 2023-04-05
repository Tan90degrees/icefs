/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:56:28
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpen.go
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
	"log"
	"syscall"
)

func (s *IcefsServer) DoIcefsOpen(ctx context.Context, req *pb.IcefsOpenReq) (*pb.IcefsOpenRes, error) {
	var res pb.IcefsOpenRes
	var procName string
	var fd int
	var err error
	s.inodeCacheLock.RLock()
	inode := s.getIcefsInode(req.Inode)
	if inode == nil {
		s.inodeCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	inode.inodeLock.Lock()
	s.inodeCacheLock.RUnlock()
	procName = fmt.Sprintf("/proc/self/fd/%v", inode.fd)
	fd, err = syscall.Open(procName, int(req.Flags&(^syscall.O_NOFOLLOW)), 0)
	if err != nil {
		inode.inodeLock.Unlock()
		log.Println(req.Flags)
		log.Println(err)
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}
	inode.nopen++
	inode.inodeLock.Unlock()
	res.Status = icefserror.ICEFS_EOK
	res.Fh = uint64(fd)

errOut:
	return &res, nil
}

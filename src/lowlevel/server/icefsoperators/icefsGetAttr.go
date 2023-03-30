/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-19 15:29:43
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:27
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsGetAttr.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
)

func (s *IcefsServer) DoIcefsGetAttr(ctx context.Context, req *pb.IcefsGetAttrReq) (*pb.IcefsGetAttrRes, error) {
	var res pb.IcefsGetAttrRes
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
	err = s.doGetAttr(inode.fd, &inode.stat)
	if err != nil {
		inode.inodeLock.Unlock()
		res.Status = icefserror.IcefsStdErrno(err)
		goto errOut
	}

	res.Stat = StatStructBuilder(&inode.stat)
	inode.inodeLock.Unlock()
	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

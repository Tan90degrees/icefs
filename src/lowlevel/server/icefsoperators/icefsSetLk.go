/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 14:39:36
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsSetLk.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"syscall"
)

func (s *IcefsServer) doIcefsSetLk(fakeInode uint64, sleep int32) (status int32) {
	return int32(syscall.ENOTSUP)
}

func (s *IcefsGRpcServer) DoIcefsSetLk(ctx context.Context, req *pb.IcefsSetLkReq) (*pb.IcefsSetLkRes, error) {
	var res pb.IcefsSetLkRes

	res.Status = s.server.doIcefsSetLk(req.Inode, req.Sleep)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsSetLk(ctx context.Context, req *icefsthrift.IcefsSetLkReq) (*icefsthrift.IcefsSetLkRes, error) {
	var res icefsthrift.IcefsSetLkRes

	res.Status = s.server.doIcefsSetLk(uint64(req.Inode), req.Sleep)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 03:33:39
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsGetLk.go
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

func (s *IcefsServer) doIcefsGetLk(Inode uint64) (status int32) {
	return int32(syscall.ENOTSUP)
}

func (s *IcefsGRpcServer) DoIcefsGetLk(ctx context.Context, req *pb.IcefsGetLkReq) (*pb.IcefsGetLkRes, error) {
	var res pb.IcefsGetLkRes

	res.Status = s.server.doIcefsGetLk(req.Inode)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsGetLk(ctx context.Context, req *icefsthrift.IcefsGetLkReq) (*icefsthrift.IcefsGetLkRes, error) {
	var res icefsthrift.IcefsGetLkRes

	res.Status = s.server.doIcefsGetLk(uint64(req.Inode))

	return &res, nil
}

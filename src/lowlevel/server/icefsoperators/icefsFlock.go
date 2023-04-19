/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 15:55:43
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFlock.go
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

func (s *IcefsServer) doIcefsFlock(fh uint64, op int32) (status int32) {
	return icefserror.IcefsStdErrno(syscall.Flock(int(fh), int(op)))
}

func (s *IcefsGRpcServer) DoIcefsFlock(ctx context.Context, req *pb.IcefsFlockReq) (*pb.IcefsFlockRes, error) {
	var res pb.IcefsFlockRes

	res.Status = s.server.doIcefsFlock(req.Fh, req.Op)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsFlock(ctx context.Context, req *icefsthrift.IcefsFlockReq) (*icefsthrift.IcefsFlockRes, error) {
	var res icefsthrift.IcefsFlockRes

	res.Status = s.server.doIcefsFlock(uint64(req.Fh), req.Op)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 15:55:04
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFallocate.go
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

func (s *IcefsServer) doIcefsFallocate(mode int32, offset int64, length int64, fh uint64) (status int32) {
	return icefserror.IcefsStdErrno(syscall.Fallocate(int(fh), uint32(mode), offset, length))
}

func (s *IcefsGRpcServer) DoIcefsFallocate(ctx context.Context, req *pb.IcefsFallocateReq) (*pb.IcefsFallocateRes, error) {
	var res pb.IcefsFallocateRes

	res.Status = s.server.doIcefsFallocate(req.Mode, req.Offset, req.Length, req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsFallocate(ctx context.Context, req *icefsthrift.IcefsFallocateReq) (*icefsthrift.IcefsFallocateRes, error) {
	var res icefsthrift.IcefsFallocateRes

	res.Status = s.server.doIcefsFallocate(req.Mode, req.Offset, req.Length, uint64(req.Fh))

	return &res, nil
}

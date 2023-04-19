/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 16:42:07
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFsync.go
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

func (s *IcefsServer) doIcefsFsync(dataSync int32, fh uint64) (status int32) {
	var err error
	if dataSync != 0 {
		err = syscall.Fdatasync(int(fh))
	} else {
		err = syscall.Fsync(int(fh))
	}
	status = icefserror.IcefsStdErrno(err)
	return
}

func (s *IcefsGRpcServer) DoIcefsFsync(ctx context.Context, req *pb.IcefsFsyncReq) (*pb.IcefsFsyncRes, error) {
	var res pb.IcefsFsyncRes

	res.Status = s.server.doIcefsFsync(req.DataSync, req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsFsync(ctx context.Context, req *icefsthrift.IcefsFsyncReq) (*icefsthrift.IcefsFsyncRes, error) {
	var res icefsthrift.IcefsFsyncRes

	res.Status = s.server.doIcefsFsync(req.DataSync, uint64(req.Fh))

	return &res, nil
}

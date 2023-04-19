/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 16:24:55
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFlush.go
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

func (s *IcefsServer) doIcefsFlush(fh uint64) (status int32) {
	newFh, err := syscall.Dup(int(fh))
	status = icefserror.IcefsStdErrno(err)
	if err == nil {
		syscall.Close(newFh)
	}
	return
}

func (s *IcefsGRpcServer) DoIcefsFlush(ctx context.Context, req *pb.IcefsFlushReq) (*pb.IcefsFlushRes, error) {
	var res pb.IcefsFlushRes

	res.Status = s.server.doIcefsFlush(req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsFlush(ctx context.Context, req *icefsthrift.IcefsFlushReq) (*icefsthrift.IcefsFlushRes, error) {
	var res icefsthrift.IcefsFlushRes

	res.Status = s.server.doIcefsFlush(uint64(req.Fh))

	return &res, nil
}

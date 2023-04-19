/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 15:20:09
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsWriteBuf.go
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

func (s *IcefsGRpcServer) DoIcefsWriteBuf(ctx context.Context, req *pb.IcefsWriteBufReq) (*pb.IcefsWriteBufRes, error) {
	var res pb.IcefsWriteBufRes
	res.Status = int32(syscall.ENOTSUP)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsWriteBuf(ctx context.Context, req *icefsthrift.IcefsWriteBufReq) (*icefsthrift.IcefsWriteBufRes, error) {
	var res icefsthrift.IcefsWriteBufRes

	return &res, nil
}

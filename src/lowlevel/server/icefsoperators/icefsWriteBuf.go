/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:30:04
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsWriteBuf.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsWriteBuf(ctx context.Context, req *pb.IcefsWriteBufReq) (*pb.IcefsWriteBufRes, error) {
	var res pb.IcefsWriteBufRes
	res.Status = int32(syscall.ENOTSUP)

	return &res, nil
}

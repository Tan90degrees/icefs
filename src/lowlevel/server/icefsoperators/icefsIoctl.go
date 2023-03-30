/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:39
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsIoctl.go
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

func (s *IcefsServer) DoIcefsIoctl(ctx context.Context, req *pb.IcefsIoctlReq) (*pb.IcefsIoctlRes, error) {
	var res pb.IcefsIoctlRes
	res.Status = int32(syscall.ENOTSUP)

	return &res, nil
}
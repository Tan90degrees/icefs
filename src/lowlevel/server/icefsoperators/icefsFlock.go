/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:55:59
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFlock.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsFlock(ctx context.Context, req *pb.IcefsFlockReq) (*pb.IcefsFlockRes, error) {
	var res pb.IcefsFlockRes
	res.Status = icefserror.IcefsStdErrno(syscall.Flock(int(req.Fh), int(req.Op)))

	return &res, nil
}

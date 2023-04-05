/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:56:04
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFlush.go
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

func (s *IcefsServer) DoIcefsFlush(ctx context.Context, req *pb.IcefsFlushReq) (*pb.IcefsFlushRes, error) {
	var res pb.IcefsFlushRes
	newFh, err := syscall.Dup(int(req.Fh))
	if err != nil {
		goto errOut
	}
	syscall.Close(newFh)
errOut:
	res.Status = icefserror.IcefsStdErrno(err)
	return &res, nil
}

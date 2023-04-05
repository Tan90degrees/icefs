/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:56:35
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRead.go
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

func (s *IcefsServer) DoIcefsRead(ctx context.Context, req *pb.IcefsReadReq) (*pb.IcefsReadRes, error) {
	var res pb.IcefsReadRes
	res.Data = make([]byte, req.Size)
	size, err := syscall.Pread(int(req.Fh), res.Data, req.Offset)
	res.Status = icefserror.IcefsStdErrno(err)
	res.Size = uint64(size)

	return &res, nil
}

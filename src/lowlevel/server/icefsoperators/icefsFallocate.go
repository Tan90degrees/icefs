/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:27:35
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFallocate.go
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

func (s *IcefsServer) DoIcefsFallocate(ctx context.Context, req *pb.IcefsFallocateReq) (*pb.IcefsFallocateRes, error) {
	var res pb.IcefsFallocateRes
	res.Status = icefserror.IcefsStdErrno(syscall.Fallocate(int(req.FileInfo.Fh), uint32(req.Mode), req.Offset, req.Length))

	return &res, nil
}

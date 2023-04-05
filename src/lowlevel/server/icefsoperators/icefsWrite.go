/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 15:12:26
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsWrite.go
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

func (s *IcefsServer) DoIcefsWrite(ctx context.Context, req *pb.IcefsWriteReq) (*pb.IcefsWriteRes, error) {
	var res pb.IcefsWriteRes
	size, err := syscall.Pwrite(int(req.Fh), req.Buf, req.Offset)
	res.Status = icefserror.IcefsStdErrno(err)
	res.Size = uint64(size)

	return &res, nil
}

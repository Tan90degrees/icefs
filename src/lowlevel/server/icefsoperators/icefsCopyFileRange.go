/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:27:01
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsCopyFileRange.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) DoIcefsCopyFileRange(ctx context.Context, req *pb.IcefsCopyFileRangeReq) (*pb.IcefsCopyFileRangeRes, error) {
	var res pb.IcefsCopyFileRangeRes
	size, err := unix.CopyFileRange(int(req.FileInfoIn.Fh), &req.OffsetIn, int(req.FileInfoOut.Fh), &req.OffsetOut, int(req.Len), int(req.Flags))
	res.Status = icefserror.IcefsStdErrno(err)
	res.Size = uint64(size) // 保证先进行status校验则回环不会出现错误
	return &res, nil
}

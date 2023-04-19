/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 15:11:11
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsCopyFileRange.go
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

	"golang.org/x/sys/unix"
)

func (s *IcefsServer) doIcefsCopyFileRange(offsetIn int64, fhIn uint64, offsetOut int64, fhOut uint64, len uint64, flags int32) (status int32, size uint64) {
	num, err := unix.CopyFileRange(int(fhIn), &offsetIn, int(fhOut), &offsetOut, int(len), int(flags))
	status = icefserror.IcefsStdErrno(err)
	size = uint64(num) // 保证先进行status校验则回环不会出现错误
	return
}

func (s *IcefsGRpcServer) DoIcefsCopyFileRange(ctx context.Context, req *pb.IcefsCopyFileRangeReq) (*pb.IcefsCopyFileRangeRes, error) {
	var res pb.IcefsCopyFileRangeRes

	res.Status, res.Size = s.server.doIcefsCopyFileRange(req.OffsetIn, req.FhIn, req.OffsetOut, req.FhOut, req.Len, req.Flags)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsCopyFileRange(ctx context.Context, req *icefsthrift.IcefsCopyFileRangeReq) (*icefsthrift.IcefsCopyFileRangeRes, error) {
	var res icefsthrift.IcefsCopyFileRangeRes
	var size uint64

	res.Status, size = s.server.doIcefsCopyFileRange(req.OffsetIn, uint64(req.FhIn), req.OffsetOut, uint64(req.FhOut), uint64(req.Len), req.Flags)
	res.Size = icefsthrift.Ui64(size)

	return &res, nil
}

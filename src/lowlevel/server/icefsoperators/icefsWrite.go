/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 17:24:30
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsWrite.go
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
	"syscall"
)

func (s *IcefsServer) doIcefsWrite(reqSize uint64, reqBuf []byte, fh uint64, offset int64) (status int32, resSize uint64) {
	bufObj := s.getRWBuf(reqSize)
	buf, _ := bufObj.(*IcefsRWBuf)
	copy(buf.mem, reqBuf)
	size, err := syscall.Pwrite(int(fh), buf.mem, offset)
	s.putRWBuf(buf)
	status = icefserror.IcefsStdErrno(err)
	resSize = uint64(size)
	return
}

func (s *IcefsGRpcServer) DoIcefsWrite(ctx context.Context, req *pb.IcefsWriteReq) (*pb.IcefsWriteRes, error) {
	var res pb.IcefsWriteRes

	res.Status, res.Size = s.server.doIcefsWrite(req.Size, req.Buf, req.Fh, req.Offset)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsWrite(ctx context.Context, req *icefsthrift.IcefsWriteReq) (*icefsthrift.IcefsWriteRes, error) {
	var res icefsthrift.IcefsWriteRes
	var size uint64

	res.Status, size = s.server.doIcefsWrite(uint64(req.Size), req.Buf, uint64(req.Fh), req.Offset)
	res.Size = icefsthrift.Ui64(size)

	return &res, nil
}

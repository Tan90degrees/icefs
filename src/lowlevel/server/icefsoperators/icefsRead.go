/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 08:15:28
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRead.go
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

func (s *IcefsServer) doIcefsRead(reqSize uint64, offset int64, fh uint64) (status int32, resSize uint64, data []byte) {
	var size int
	var err error

	if ((reqSize & (s.logicalBlockSize - 1)) | (uint64(offset) & (s.logicalBlockSize - 1))) == 0 {
		data = s.getAlignedMem(reqSize)
		size, err = syscall.Pread(int(fh), data, offset)
		status = icefserror.IcefsStdErrno(err)
		resSize = uint64(size)
	} else {
		var readOffset int64 = 0
		var paddingHead uint64 = 0
		var padding uint64 = 0
		if offset >= int64(s.logicalBlockSize) {
			paddingHead = uint64(offset & (int64(s.logicalBlockSize) - 1))
			readOffset = offset - int64(paddingHead)
		}
		padding = paddingHead + s.logicalBlockSize - ((reqSize + paddingHead) & (s.logicalBlockSize - 1))

		data = s.getAlignedMem(reqSize + padding)
		size, err = syscall.Pread(int(fh), data, readOffset)
		status = icefserror.IcefsStdErrno(err)
		size -= int(padding)
		if size >= 0 {
			resSize = uint64(size)
		} else {
			resSize = 0
		}

		data = data[paddingHead : paddingHead+resSize]
	}
	return
}

func (s *IcefsGRpcServer) DoIcefsRead(ctx context.Context, req *pb.IcefsReadReq) (*pb.IcefsReadRes, error) {
	var res pb.IcefsReadRes

	res.Status, res.Size, res.Data = s.server.doIcefsRead(req.Size, req.Offset, req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsRead(ctx context.Context, req *icefsthrift.IcefsReadReq) (*icefsthrift.IcefsReadRes, error) {
	var res icefsthrift.IcefsReadRes
	var size uint64

	res.Status, size, res.Data = s.server.doIcefsRead(uint64(req.Size), req.Offset, uint64(req.Fh))
	res.Size = icefsthrift.Ui64(size)

	return &res, nil
}

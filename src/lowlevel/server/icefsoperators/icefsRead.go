/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-11 09:37:08
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
	var size int
	var err error

	if ((req.Size & (s.logicalBlockSize - 1)) | (uint64(req.Offset) & (s.logicalBlockSize - 1))) == 0 {
		res.Data = s.getAlignedMem(req.Size)
		size, err = syscall.Pread(int(req.Fh), res.Data, req.Offset)
		res.Status = icefserror.IcefsStdErrno(err)
		res.Size = uint64(size)
	} else {
		var offset int64 = 0
		var paddingHead uint64 = 0
		var padding uint64 = 0
		if req.Offset >= int64(s.logicalBlockSize) {
			paddingHead = uint64(req.Offset & (int64(s.logicalBlockSize) - 1))
			offset = req.Offset - int64(paddingHead)
		}
		padding = paddingHead + s.logicalBlockSize - ((req.Size + paddingHead) & (s.logicalBlockSize - 1))

		res.Data = s.getAlignedMem(req.Size + padding)
		size, err = syscall.Pread(int(req.Fh), res.Data, offset)
		res.Status = icefserror.IcefsStdErrno(err)
		size -= int(padding)
		if size >= 0 {
			res.Size = uint64(size)
		} else {
			res.Size = 0
		}

		res.Data = res.Data[paddingHead : paddingHead+res.Size]
	}

	return &res, nil
}

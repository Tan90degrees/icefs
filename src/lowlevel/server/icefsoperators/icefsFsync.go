/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:19
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsFsync.go
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

func (s *IcefsServer) DoIcefsFsync(ctx context.Context, req *pb.IcefsFsyncReq) (*pb.IcefsFsyncRes, error) {
	var res pb.IcefsFsyncRes
	var err error
	if req.DataSync != 0 {
		err = syscall.Fdatasync(int(req.FileInfo.Fh))
	} else {
		err = syscall.Fsync(int(req.FileInfo.Fh))
	}
	res.Status = icefserror.IcefsStdErrno(err)

	return &res, nil
}

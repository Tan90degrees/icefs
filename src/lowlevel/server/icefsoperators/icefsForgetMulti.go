/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 16:38:26
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsForgetMulti.go
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
)

func (s *IcefsGRpcServer) DoIcefsForgetMulti(ctx context.Context, req *pb.IcefsForgetMultiReq) (*pb.IcefsForgetMultiRes, error) {
	var res pb.IcefsForgetMultiRes
	var i uint64
	count := req.Count
	for i = 0; i < count; i++ {
		s.server.doIcefsForget(req.ToForget[i].Inode, req.ToForget[i].Nlookup)
	}
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsForgetMulti(ctx context.Context, req *icefsthrift.IcefsForgetMultiReq) (*icefsthrift.IcefsForgetMultiRes, error) {
	var res icefsthrift.IcefsForgetMultiRes

	var i uint64
	count := uint64(req.Count)
	for i = 0; i < count; i++ {
		s.server.doIcefsForget(uint64(req.ToForget[i].Inode), uint64(req.ToForget[i].Nlookup))
	}
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

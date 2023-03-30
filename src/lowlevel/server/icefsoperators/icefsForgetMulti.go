/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:15
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsForgetMulti.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
)

func (s *IcefsServer) DoIcefsForgetMulti(ctx context.Context, req *pb.IcefsForgetMultiReq) (*pb.IcefsForgetMultiRes, error) {
	var res pb.IcefsForgetMultiRes
	var i uint64
	count := req.Count
	for i = 0; i < count; i++ {
		s.doForget(req.ToForget[i].Inode, req.ToForget[i].Nlookup)
	}
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

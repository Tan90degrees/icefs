/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:13
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsForget.go
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

func (s *IcefsServer) DoIcefsForget(ctx context.Context, req *pb.IcefsForgetReq) (*pb.IcefsForgetRes, error) {
	var res pb.IcefsForgetRes
	s.doForget(req.Inode, req.Nlookup)
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

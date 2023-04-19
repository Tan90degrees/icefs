/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 16:34:26
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsForget.go
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

func (s *IcefsGRpcServer) DoIcefsForget(ctx context.Context, req *pb.IcefsForgetReq) (*pb.IcefsForgetRes, error) {
	var res pb.IcefsForgetRes

	s.server.doIcefsForget(req.Inode, req.Nlookup)
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsForget(ctx context.Context, req *icefsthrift.IcefsForgetReq) (*icefsthrift.IcefsForgetRes, error) {
	var res icefsthrift.IcefsForgetRes

	s.server.doIcefsForget(uint64(req.Inode), uint64(req.Nlookup))
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 07:07:04
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsLseek.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
)

func (s *IcefsServer) doIcefsLseek(fh uint64, reqOffset int64, whence int32) (status int32, resOffset int64) {
	resOffset, status = IcefsLseek(int32(fh), reqOffset, whence)
	return
}

func (s *IcefsGRpcServer) DoIcefsLseek(ctx context.Context, req *pb.IcefsLseekReq) (*pb.IcefsLseekRes, error) {
	var res pb.IcefsLseekRes

	res.Status, res.Offset = s.server.doIcefsLseek(req.Fh, req.Offset, req.Whence)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsLseek(ctx context.Context, req *icefsthrift.IcefsLseekReq) (*icefsthrift.IcefsLseekRes, error) {
	var res icefsthrift.IcefsLseekRes

	res.Status, res.Offset = s.server.doIcefsLseek(uint64(req.Fh), req.Offset, req.Whence)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 08:03:55
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsPoll.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"syscall"
)

func (s *IcefsServer) doIcefsPoll(fakeInode uint64) (status int32, revents uint32) {
	status = int32(syscall.ENOTSUP)
	return
}

func (s *IcefsGRpcServer) DoIcefsPoll(ctx context.Context, req *pb.IcefsPollReq) (*pb.IcefsPollRes, error) {
	var res pb.IcefsPollRes

	res.Status, res.Revents = s.server.doIcefsPoll(req.Inode)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsPoll(ctx context.Context, req *icefsthrift.IcefsPollReq) (*icefsthrift.IcefsPollRes, error) {
	var res icefsthrift.IcefsPollRes
	var revents uint32

	res.Status, revents = s.server.doIcefsPoll(uint64(req.Inode))
	res.Revents = icefsthrift.Ui32(revents)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 04:08:11
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsIoctl.go
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

func (s *IcefsServer) doIcefsIoctl(inode uint64, cmd uint32, arg []byte, flags uint32, inBufSize uint64, outBufSize uint64) (status int32) {
	return int32(syscall.ENOTSUP)
}

func (s *IcefsGRpcServer) DoIcefsIoctl(ctx context.Context, req *pb.IcefsIoctlReq) (*pb.IcefsIoctlRes, error) {
	var res pb.IcefsIoctlRes

	res.Status = s.server.doIcefsIoctl(req.Inode, req.Cmd, req.Arg, req.Flags, req.InBufSize, req.OutBufSize)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsIoctl(ctx context.Context, req *icefsthrift.IcefsIoctlReq) (*icefsthrift.IcefsIoctlRes, error) {
	var res icefsthrift.IcefsIoctlRes

	res.Status = s.server.doIcefsIoctl(uint64(req.Inode), uint32(req.Cmd), req.Arg, uint32(req.Flags), uint64(req.InBufSize), uint64(req.OutBufSize))

	return &res, nil
}

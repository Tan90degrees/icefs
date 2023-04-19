/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 13:48:41
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRetrieveReply.go
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

func (s *IcefsServer) doIcefsRetrieveReply(cookie []byte, fakeInode uint64, offset int64) (status int32) {
	return int32(syscall.ENOTSUP)
}

func (s *IcefsGRpcServer) DoIcefsRetrieveReply(ctx context.Context, req *pb.IcefsRetrieveReplyReq) (*pb.IcefsRetrieveReplyRes, error) {
	var res pb.IcefsRetrieveReplyRes

	res.Status = s.server.doIcefsRetrieveReply(req.Cookie, req.Inode, req.Offset)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsRetrieveReply(ctx context.Context, req *icefsthrift.IcefsRetrieveReplyReq) (*icefsthrift.IcefsRetrieveReplyRes, error) {
	var res icefsthrift.IcefsRetrieveReplyRes

	res.Status = s.server.doIcefsRetrieveReply(req.Cookie, uint64(req.Inode), req.Offset)

	return &res, nil
}

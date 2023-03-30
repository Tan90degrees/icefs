/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:36
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsRetrieveReply.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsRetrieveReply(ctx context.Context, req *pb.IcefsRetrieveReplyReq) (*pb.IcefsRetrieveReplyRes, error) {
	var res pb.IcefsRetrieveReplyRes
	res.Status = int32(syscall.ENOTSUP)

	return &res, nil
}

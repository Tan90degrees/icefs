/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:56:21
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsLseek.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	pb "icefs-server/icefsrpc"
)

func (s *IcefsServer) DoIcefsLseek(ctx context.Context, req *pb.IcefsLseekReq) (*pb.IcefsLseekRes, error) {
	var res pb.IcefsLseekRes
	res.Offset, res.Status = IcefsLseek(int32(req.Fh), req.Offset, req.Whence)

	return &res, nil
}

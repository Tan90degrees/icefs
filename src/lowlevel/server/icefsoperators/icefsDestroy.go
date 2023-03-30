/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:27:33
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsDestroy.go
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

func (s *IcefsServer) DoIcefsDestroy(ctx context.Context, req *pb.IcefsDestroyReq) (*pb.IcefsDestroyRes, error) {
	res := pb.IcefsDestroyRes{
		Status: icefserror.ICEFS_EOK,
		Info:   "Bye",
	}

	return &res, nil
}

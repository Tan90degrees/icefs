/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:38
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsInit.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"log"
)

func (s *IcefsServer) DoIcefsInit(ctx context.Context, req *pb.IcefsInitReq) (*pb.IcefsInitRes, error) {
	var res pb.IcefsInitRes
	log.Println(req.Uuid, req.Info, req.Want, req.Timeout)
	s.timeout = req.Timeout
	res = pb.IcefsInitRes{
		Status: icefserror.ICEFS_EOK,
		Info:   "Welcome",
		Can:    req.Want,
	}

	return &res, nil
}

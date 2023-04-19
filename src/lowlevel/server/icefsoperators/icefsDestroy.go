/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 15:51:58
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsDestroy.go
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

func (s *IcefsServer) doIcefsDestroy(hostName string, reqInfo string) (status int32, resInfo string) {
	status = icefserror.ICEFS_EOK
	resInfo = "Bye"
	return
}

func (s *IcefsGRpcServer) DoIcefsDestroy(ctx context.Context, req *pb.IcefsDestroyReq) (*pb.IcefsDestroyRes, error) {
	var res pb.IcefsDestroyRes

	res.Status, res.Info = s.server.doIcefsDestroy(req.HostName, req.Info)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsDestroy(ctx context.Context, req *icefsthrift.IcefsDestroyReq) (*icefsthrift.IcefsDestroyRes, error) {
	var res icefsthrift.IcefsDestroyRes

	res.Status, res.Info = s.server.doIcefsDestroy(req.HostName, req.Info)

	return &res, nil
}

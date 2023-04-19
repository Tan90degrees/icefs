/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 04:03:01
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsInit.go
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
	"log"
)

func (s *IcefsServer) doIcefsInit(uuid string, reqInfo string, want uint32, timeout float64) (status int32, resInfo string, can uint32) {
	log.Printf("uuid: %s\nreq info: %s\nwant: %d\ntimeout: %f", uuid, reqInfo, want, timeout)
	s.timeout = timeout
	return icefserror.ICEFS_EOK, "Welcome", want
}

func (s *IcefsGRpcServer) DoIcefsInit(ctx context.Context, req *pb.IcefsInitReq) (*pb.IcefsInitRes, error) {
	var res pb.IcefsInitRes

	res.Status, res.Info, res.Can = s.server.doIcefsInit(req.Uuid, req.Info, req.Want, req.Timeout)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsInit(ctx context.Context, req *icefsthrift.IcefsInitReq) (*icefsthrift.IcefsInitRes, error) {
	var res icefsthrift.IcefsInitRes
	var can uint32

	res.Status, res.Info, can = s.server.doIcefsInit(req.UUID, req.Info, uint32(req.Want), req.Timeout)
	if res.Status == icefserror.ICEFS_EOK {
		res.Can = icefsthrift.Ui32(can)
	}

	return &res, nil
}

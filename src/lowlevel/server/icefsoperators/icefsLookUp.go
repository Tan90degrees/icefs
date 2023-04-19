/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 06:11:38
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsLookUp.go
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

func (s *IcefsGRpcServer) DoIcefsLookUp(ctx context.Context, req *pb.IcefsLookUpReq) (*pb.IcefsLookUpRes, error) {
	var res pb.IcefsLookUpRes
	var entry any
	var err error

	entry, err = s.server.doIcefsLookUp(req.ParentInode, req.Name, GRpcFuseEntryParamBuilder)
	if err == nil {
		res.Entry = entry.(*pb.FuseEntryParam)
	}
	res.Status = icefserror.IcefsStdErrno(err)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsLookUp(ctx context.Context, req *icefsthrift.IcefsLookUpReq) (*icefsthrift.IcefsLookUpRes, error) {
	var res icefsthrift.IcefsLookUpRes
	var entry any
	var err error

	entry, err = s.server.doIcefsLookUp(uint64(req.ParentInode), req.Name, ThriftFuseEntryParamBuilder)
	if err == nil {
		res.Entry = entry.(*icefsthrift.FuseEntryParam)
	}
	res.Status = icefserror.IcefsStdErrno(err)

	return &res, nil
}

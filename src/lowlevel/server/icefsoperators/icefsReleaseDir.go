/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 13:45:24
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReleaseDir.go
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

func (s *IcefsServer) doIcefsReleaseDir(fh uint64) (status int32) {
	s.delIcefsDir(fh)
	return int32(icefserror.ICEFS_EOK)
}

func (s *IcefsGRpcServer) DoIcefsReleaseDir(ctx context.Context, req *pb.IcefsReleaseDirReq) (*pb.IcefsReleaseDirRes, error) {
	var res pb.IcefsReleaseDirRes

	res.Status = s.server.doIcefsReleaseDir(req.Fh)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsReleaseDir(ctx context.Context, req *icefsthrift.IcefsReleaseDirReq) (*icefsthrift.IcefsReleaseDirRes, error) {
	var res icefsthrift.IcefsReleaseDirRes

	res.Status = s.server.doIcefsReleaseDir(uint64(req.Fh))

	return &res, nil
}

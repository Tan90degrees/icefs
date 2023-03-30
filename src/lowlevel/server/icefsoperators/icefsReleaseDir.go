/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:29
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReleaseDir.go
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

func (s *IcefsServer) DoIcefsReleaseDir(ctx context.Context, req *pb.IcefsReleaseDirReq) (*pb.IcefsReleaseDirRes, error) {
	var res pb.IcefsReleaseDirRes
	s.delIcefsDir(req.FileInfo.Fh)
	res.Status = icefserror.ICEFS_EOK

	return &res, nil
}

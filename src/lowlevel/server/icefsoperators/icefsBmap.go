/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-04-03 06:55:36
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 14:32:22
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsBmap.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 14:23:40
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsBmap.go
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

func (s *IcefsServer) doIcefsBmap(fakeInode uint64, blockSize uint64, index uint64) (status int32) {
	return int32(syscall.ENOTSUP)
}

func (s *IcefsGRpcServer) DoIcefsBmap(ctx context.Context, req *pb.IcefsBmapReq) (*pb.IcefsBmapRes, error) {
	var res pb.IcefsBmapRes

	res.Status = s.server.doIcefsBmap(req.Inode, req.BlockSize, req.Index)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsBmap(ctx context.Context, req *icefsthrift.IcefsBmapReq) (*icefsthrift.IcefsBmapRes, error) {
	var res icefsthrift.IcefsBmapRes

	res.Status = s.server.doIcefsBmap(uint64(req.Inode), uint64(req.BlockSize), uint64(req.Index))

	return &res, nil
}

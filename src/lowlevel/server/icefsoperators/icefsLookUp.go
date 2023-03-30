/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:28:46
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsLookUp.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"context"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"syscall"
)

func (s *IcefsServer) DoIcefsLookUp(ctx context.Context, req *pb.IcefsLookUpReq) (*pb.IcefsLookUpRes, error) {
	var res pb.IcefsLookUpRes
	var err error
	res.Entry, err = s.doLookUp(req.ParentInode, req.Name)
	res.Status = icefserror.ICEFS_EOK
	if err == syscall.ENOENT {
		res.Entry = new(pb.FuseEntryParam)
		res.Entry.Attr = new(pb.StatStruct)
		res.Entry.AttrTimeout = s.timeout
		res.Entry.EntryTimeout = s.timeout
		res.Entry.Inode = 0
		res.Entry.Attr.StIno = 0
	} else {
		res.Status = icefserror.IcefsStdErrno(err)
	}

	return &res, nil
}

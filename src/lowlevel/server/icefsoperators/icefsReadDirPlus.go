/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 18:02:13
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReadDirPlus.go
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

type IcefsReadDirPlusAppendEntryData func(fakeInode uint64, dir *IcefsDir, dirEntry any, dataSlice any) (status int32)

func (s *IcefsServer) GRpcIcefsReadDirPlusAppendEntryData(fakeInode uint64, dir *IcefsDir, dirEntry any, dataSlice any) (status int32) {
	var err error
	var entry any
	data := new(pb.IcefsReadDirPlusResDataPlus)
	data.DirEntry = dirEntry.(*pb.DirentStruct)

	dir.offset = data.DirEntry.Off

	if checkNameIsDotOrDotDot(data.DirEntry.Name) {
		return
	}

	entry, err = s.doIcefsLookUp(fakeInode, data.DirEntry.Name, GRpcFuseEntryParamBuilder)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}

	data.Entry = entry.(*pb.FuseEntryParam)

	realDataSlice := dataSlice.(*[]*pb.IcefsReadDirPlusResDataPlus)

	*realDataSlice = append(*realDataSlice, data)
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsServer) ThriftIcefsReadDirPlusAppendEntryData(fakeInode uint64, dir *IcefsDir, dirEntry any, dataSlice any) (status int32) {
	var err error
	var entry any
	data := new(icefsthrift.IcefsReadDirPlusData)
	data.DirEntry = dirEntry.(*icefsthrift.DirentStruct)

	dir.offset = data.DirEntry.Off

	if checkNameIsDotOrDotDot(data.DirEntry.Name) {
		return
	}

	entry, err = s.doIcefsLookUp(fakeInode, data.DirEntry.Name, ThriftFuseEntryParamBuilder)
	if err != nil {
		status = icefserror.IcefsStdErrno(err)
		return
	}

	data.Entry = entry.(*icefsthrift.FuseEntryParam)

	realDataSlice := dataSlice.(*[]*icefsthrift.IcefsReadDirPlusData)

	*realDataSlice = append((*realDataSlice), data)
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsServer) doIcefsReadDirPlus(fakeInode uint64, offset int64, fh uint64, dataSlice any, dirStructBuilder DirStructBuilder, icefsReadDirPlusAppendEntryData IcefsReadDirPlusAppendEntryData) (status int32) {
	var dirEntry any
	var errno int32

	s.dirCacheLock.RLock()
	dir := s.getIcefsDir(fh)
	if dir == nil {
		s.dirCacheLock.RUnlock()
		status = icefserror.ICEFS_BUG_ERR
		return
	}
	dir.dirLock.Lock()
	s.dirCacheLock.RUnlock()
	if offset != dir.offset {
		IcefsSeekDir(dir.dirStream, offset)
		dir.offset = offset
	}

	for {
		dirEntry, errno = IcefsReadDir(dir.dirStream, dirStructBuilder)
		if dirEntry == nil {
			if errno != icefserror.ICEFS_EOK {
				status = errno
				dir.dirLock.Unlock()
				return
			}
			break
		}

		status = icefsReadDirPlusAppendEntryData(fakeInode, dir, dirEntry, dataSlice)
		if status != icefserror.ICEFS_EOK {
			dir.dirLock.Unlock()
			return
		}
	}

	dir.dirLock.Unlock()
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsReadDirPlus(ctx context.Context, req *pb.IcefsReadDirPlusReq) (*pb.IcefsReadDirPlusRes, error) {
	var res pb.IcefsReadDirPlusRes
	res.Data = make([]*pb.IcefsReadDirPlusResDataPlus, 0)

	res.Status = s.server.doIcefsReadDirPlus(req.Inode, req.Offset, req.Fh, &res.Data, GRpcDirStructBuilder, s.server.GRpcIcefsReadDirPlusAppendEntryData)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsReadDirPlus(ctx context.Context, req *icefsthrift.IcefsReadDirPlusReq) (*icefsthrift.IcefsReadDirPlusRes, error) {
	var res icefsthrift.IcefsReadDirPlusRes
	res.Data = make([]*icefsthrift.IcefsReadDirPlusData, 0)

	res.Status = s.server.doIcefsReadDirPlus(uint64(req.Inode), req.Offset, uint64(req.Fh), &res.Data, ThriftDirStructBuilder, s.server.ThriftIcefsReadDirPlusAppendEntryData)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-05-09 09:12:47
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReadDir.go
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

type IcefsReadDirAppendEntryData func(dir *IcefsDir, entry any, dataSlice any)

func GRpcIcefsReadDirAppendEntryData(dir *IcefsDir, entry any, dataSlice any) {
	realEntry := entry.(*pb.DirentStruct)
	dir.offset = realEntry.Off
	if checkNameIsDotOrDotDot(realEntry.Name) {
		return
	}
	realDataSlice := dataSlice.(*[]*pb.DirentStruct)
	*realDataSlice = append(*realDataSlice, realEntry)
}

func ThriftIcefsReadDirAppendEntryData(dir *IcefsDir, entry any, dataSlice any) {
	realEntry := entry.(*icefsthrift.DirentStruct)
	dir.offset = realEntry.Off
	if checkNameIsDotOrDotDot(realEntry.Name) {
		return
	}
	realDataSlice := dataSlice.(*[]*icefsthrift.DirentStruct)
	*realDataSlice = append(*realDataSlice, realEntry)
}

func (s *IcefsServer) doIcefsReadDir(offset int64, fh uint64, dataSlice any, dirStructBuilder DirStructBuilder, icefsReadDirAppendEntryData IcefsReadDirAppendEntryData) (status int32) {
	var entry any
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
		entry, errno = IcefsReadDir(dir.dirStream, dirStructBuilder)
		if entry == nil {
			if errno != icefserror.ICEFS_EOK {
				status = errno
				dir.dirLock.Unlock()
				return
			}
			break
		}
		icefsReadDirAppendEntryData(dir, entry, dataSlice)
	}

	dir.dirLock.Unlock()
	status = icefserror.ICEFS_EOK
	return
}

func (s *IcefsGRpcServer) DoIcefsReadDir(ctx context.Context, req *pb.IcefsReadDirReq) (*pb.IcefsReadDirRes, error) {
	var res pb.IcefsReadDirRes
	res.Data = make([]*pb.DirentStruct, 0)

	res.Status = s.server.doIcefsReadDir(req.Offset, req.Fh, &res.Data, GRpcDirStructBuilder, GRpcIcefsReadDirAppendEntryData)

	return &res, nil
}

func (s *IcefsThriftServer) DoIcefsReadDir(ctx context.Context, req *icefsthrift.IcefsReadDirReq) (*icefsthrift.IcefsReadDirRes, error) {
	var res icefsthrift.IcefsReadDirRes
	res.Data = make([]*icefsthrift.DirentStruct, 0)

	res.Status = s.server.doIcefsReadDir(req.Offset, uint64(req.Fh), res.Data, ThriftDirStructBuilder, ThriftIcefsReadDirAppendEntryData)

	return &res, nil
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:17
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReadDirPlus.go
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

func (s *IcefsServer) DoIcefsReadDirPlus(ctx context.Context, req *pb.IcefsReadDirPlusReq) (*pb.IcefsReadDirPlusRes, error) {
	var res pb.IcefsReadDirPlusRes
	// var entry *pb.DirentStruct
	var data *pb.IcefsReadDirPlusResDataPlus
	var errno int32
	var err error

	s.dirCacheLock.RLock()
	dir := s.getIcefsDir(req.FileInfo.Fh)
	if dir == nil {
		s.dirCacheLock.RUnlock()
		res.Status = icefserror.ICEFS_BUG_ERR
		goto errOut
	}
	dir.dirLock.Lock()
	s.dirCacheLock.RUnlock()
	if req.Offset != dir.offset {
		IcefsSeekDir(dir.dirStream, req.Offset)
		dir.offset = req.Offset
	}

	for {
		data = new(pb.IcefsReadDirPlusResDataPlus)
		data.DirEntry, errno = IcefsReadDir(dir.dirStream)
		if data.DirEntry == nil {
			if errno != icefserror.ICEFS_EOK {
				res.Status = errno
				dir.dirLock.Unlock()
				goto errOut
			}
			break
		}
		dir.offset = data.DirEntry.Off

		if checkNameIsDotOrDotDot(data.DirEntry.Name) {
			continue
		}

		data.Entry, err = s.doLookUp(req.Inode, data.DirEntry.Name)
		if err != nil {
			res.Status = icefserror.IcefsStdErrno(err)
			goto errOut
		}

		res.Data = append(res.Data, data)
	}

	dir.dirLock.Unlock()
	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

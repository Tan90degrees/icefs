/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 07:18:32
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-04 14:56:40
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsReadDir.go
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

func (s *IcefsServer) DoIcefsReadDir(ctx context.Context, req *pb.IcefsReadDirReq) (*pb.IcefsReadDirRes, error) {
	var res pb.IcefsReadDirRes
	var entry *pb.DirentStruct
	var errno int32
	s.dirCacheLock.RLock()
	dir := s.getIcefsDir(req.Fh)
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
		entry, errno = IcefsReadDir(dir.dirStream)
		if entry == nil {
			if errno != icefserror.ICEFS_EOK {
				res.Status = errno
				dir.dirLock.Unlock()
				goto errOut
			}
			break
		}
		dir.offset = entry.Off

		if checkNameIsDotOrDotDot(entry.Name) {
			continue
		}
		res.Data = append(res.Data, entry)
	}

	dir.dirLock.Unlock()
	res.Status = icefserror.ICEFS_EOK

errOut:
	return &res, nil
}

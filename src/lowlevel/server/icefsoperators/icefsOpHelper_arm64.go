/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-24 09:06:11
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:05
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpHelper_arm64.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	pb "icefs-server/icefsrpc"
	"syscall"
)

func StatStructBuilder(stat *syscall.Stat_t) *pb.StatStruct {
	return &pb.StatStruct{
		StDev:   stat.Dev,
		StIno:   stat.Ino,
		StMode:  stat.Mode,
		StNlink: uint64(stat.Nlink),
		StUid:   stat.Uid,
		StGid:   stat.Gid,
		StRdev:  stat.Rdev,
		StSize:  stat.Size,
		StAtim: &pb.TimeStruct{TimeSec: stat.Atim.Sec,
			TimeNSec: stat.Atim.Nsec},
		StMtim: &pb.TimeStruct{TimeSec: stat.Mtim.Sec,
			TimeNSec: stat.Mtim.Nsec},
		StCtim: &pb.TimeStruct{TimeSec: stat.Ctim.Sec,
			TimeNSec: stat.Ctim.Nsec},
		StBlksize: int64(stat.Blksize),
		StBlocks:  stat.Blocks,
	}
}

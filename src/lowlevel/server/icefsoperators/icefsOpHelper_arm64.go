/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-24 09:06:11
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-17 17:07:31
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsOpHelper_arm64.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"syscall"
)

func GRpcStatStructBuilder(stat *syscall.Stat_t) any {
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

func ThriftStatStructBuilder(stat *syscall.Stat_t) any {
	return &icefsthrift.StatStruct{
		StDev:   icefsthrift.Ui64(stat.Dev),
		StIno:   icefsthrift.Ui64(stat.Ino),
		StMode:  icefsthrift.Ui32(stat.Mode),
		StNlink: icefsthrift.Ui64(stat.Nlink),
		StUID:   icefsthrift.Ui32(stat.Uid),
		StGid:   icefsthrift.Ui32(stat.Gid),
		StRdev:  icefsthrift.Ui64(stat.Rdev),
		StSize:  stat.Size,
		StAtim: &icefsthrift.TimeStruct{TimeSec: stat.Atim.Sec,
			TimeNSec: stat.Atim.Nsec},
		StMtim: &icefsthrift.TimeStruct{TimeSec: stat.Mtim.Sec,
			TimeNSec: stat.Mtim.Nsec},
		StCtim: &icefsthrift.TimeStruct{TimeSec: stat.Ctim.Sec,
			TimeNSec: stat.Ctim.Nsec},
		StBlksize: int64(stat.Blksize),
		StBlocks:  stat.Blocks,
	}
}

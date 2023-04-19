/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-18 03:31:47
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-18 17:31:41
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsGlibc.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

/*

#include <stdio.h>
#include <unistd.h>
#include <dirent.h>
#include <stdint.h>
#include <errno.h>

int32_t getErrno(void) {
	return errno;
}

int32_t initErrno(void) {
	errno = 0;
	return errno;
}

*/
import "C"
import (
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsthrift"
	"unsafe"
)

const DIRENT_NAME_LEN = 256

type DirStructBuilder func(dirent *C.struct_dirent) any

func IcefsDirFd(dir unsafe.Pointer) int {
	return int(C.dirfd((*C.DIR)(dir)))
}

func IcefsFdOpenDir(fd int) (unsafe.Pointer, int) {
	C.initErrno()
	return unsafe.Pointer(C.fdopendir(C.int(fd))), int(C.getErrno())
}

func IcefsLseek(fd int32, offset int64, whence int32) (int64, int32) {
	C.initErrno()
	return int64(C.lseek(C.int(fd), C.long(offset), C.int(whence))), int32(C.getErrno())
}

func IcefsCloseDir(dirStream unsafe.Pointer) {
	if dirStream != nil {
		C.closedir((*C.DIR)(dirStream))
	}
}

func IcefsSeekDir(dirStream unsafe.Pointer, offset int64) {
	C.seekdir((*C.DIR)(dirStream), C.long(offset))
}

func GRpcDirStructBuilder(dirent *C.struct_dirent) any {
	return &pb.DirentStruct{
		Ino:    uint64(dirent.d_ino),
		Off:    int64(dirent.d_off),
		Reclen: uint32(dirent.d_reclen),
		Type:   uint32(dirent.d_type),
		Name:   C.GoString(&(dirent.d_name[0])),
	}
}

func ThriftDirStructBuilder(dirent *C.struct_dirent) any {
	return &icefsthrift.DirentStruct{
		Ino:    icefsthrift.Ui64(dirent.d_ino),
		Off:    int64(dirent.d_off),
		Reclen: icefsthrift.Ui32(dirent.d_reclen),
		Type:   icefsthrift.Ui32(dirent.d_type),
		Name:   C.GoString(&(dirent.d_name[0])),
	}
}

func IcefsReadDir(dirStream unsafe.Pointer, dirStructBuilder DirStructBuilder) (any, int32) {
	C.initErrno()
	dirent := C.readdir((*C.DIR)(dirStream))
	if dirent == nil {
		return nil, int32(C.getErrno())
	}
	dirStruct := dirStructBuilder(dirent)

	return dirStruct, int32(C.getErrno())
}

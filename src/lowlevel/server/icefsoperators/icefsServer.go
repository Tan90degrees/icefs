/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-19 15:30:06
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:29:42
 * @FilePath: /icefs/src/lowlevel/server/icefsoperators/icefsServer.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package icefsoperators

import (
	"errors"
	"icefs-server/icefserror"
	pb "icefs-server/icefsrpc"
	"os"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

type IcefsInode struct {
	fd         int
	fakeInode  uint64
	generation uint64
	nopen      uint64
	nlookup    uint64
	stat       syscall.Stat_t
	inodeLock  sync.RWMutex
}

type IcefsDir struct {
	dirStream unsafe.Pointer
	offset    int64
	dirLock   sync.RWMutex
}

type IcefsServer struct {
	RootPathAbs    string
	timeout        float64
	devId          uint64
	inodeCache     map[uint64]*IcefsInode // TODO: use sync.map
	inodeCacheLock sync.RWMutex
	dirCache       map[uint64]*IcefsDir
	dirCacheLock   sync.RWMutex
	pb.UnimplementedIcefsServer
}

func (s *IcefsServer) IcefsServerInit() error {
	var err error
	var root IcefsInode

	s.inodeCache = make(map[uint64]*IcefsInode)
	s.dirCache = make(map[uint64]*IcefsDir)

	root.fakeInode = FUSE_ROOT_INODE

	// Can use root.stat.Mode to check.
	if fi, err := os.Stat(s.RootPathAbs); err == nil {
		if !fi.IsDir() {
			return errors.New(s.RootPathAbs + "is not a directory.")
		}
	} else {
		return err
	}

	err = syscall.Lstat(s.RootPathAbs, &root.stat)
	if err != nil {
		return err
	}

	s.devId = root.stat.Dev

	root.fd, err = syscall.Open(s.RootPathAbs, unix.O_PATH, 0)
	if err != nil {
		return err
	}

	root.nlookup = icefserror.ICEFS_BUG_ERR
	s.inodeCacheLock.Lock()
	s.inodeCache[FUSE_ROOT_INODE] = &root
	s.inodeCacheLock.Unlock()

	return err
}

func IcefsSetFakeInode(inode *IcefsInode) {
	inode.fakeInode = (uint64)(uintptr(unsafe.Pointer(inode)))
}

// TODO: 增加一个map做映射
func (s *IcefsServer) checkIcefsInode(newInode **IcefsInode) bool {
	exist := false
	for _, inode := range s.inodeCache {
		inode.inodeLock.Lock()
		if inode.stat.Ino == (*newInode).stat.Ino {
			inode.nlookup++
			inode.generation = (*newInode).generation
			inode.stat = (*newInode).stat
			*newInode = inode
			exist = true
			inode.inodeLock.Unlock()
			break
		}
		inode.inodeLock.Unlock()
	}
	return exist
}

func (s *IcefsServer) putIcefsInode(inode *IcefsInode) {
	s.inodeCache[inode.fakeInode] = inode
}

func (s *IcefsServer) getIcefsInode(inode uint64) *IcefsInode {
	icefsInode := s.inodeCache[inode]
	return icefsInode
}

// 自带锁操作
func (s *IcefsServer) delIcefsInode(inode uint64) error {
	s.inodeCacheLock.Lock()
	icefsInode := s.getIcefsInode(inode)
	if icefsInode == nil {
		s.inodeCacheLock.Unlock()
		return nil
	}
	icefsInode.inodeLock.Lock()
	err := syscall.Close(icefsInode.fd)
	delete(s.inodeCache, inode)
	s.inodeCacheLock.Unlock()
	return err
}

func (s *IcefsServer) putIcefsDir(fh uint64, dir *IcefsDir) {
	s.dirCache[fh] = dir
}

func (s *IcefsServer) getIcefsDir(fh uint64) *IcefsDir {
	dir := s.dirCache[fh]
	return dir
}

// 自带锁操作
func (s *IcefsServer) delIcefsDir(fh uint64) {
	s.dirCacheLock.Lock()
	dir := s.dirCache[fh]
	if dir == nil {
		return
	}
	dir.dirLock.Lock()
	IcefsCloseDir(dir.dirStream)
	delete(s.dirCache, fh)
	s.dirCacheLock.Unlock()
}

/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 05:57:15
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-03-30 04:30:12
 * @FilePath: /icefs/src/lowlevel/server/main.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package main

import (
	"flag"
	"fmt"
	"icefs-server/icefsoperators"
	pb "icefs-server/icefsrpc"
	"log"
	"net"
	"path/filepath"

	"google.golang.org/grpc"
)

func main() {
	rpcPort := flag.Uint("port", 10086, "The port of icefs server")
	srvPath := flag.String("path", ".", "The path to serve at")
	flag.Parse()
	ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *rpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	var icefsServer icefsoperators.IcefsServer
	icefsServer.RootPathAbs, err = filepath.Abs(*srvPath)
	if err != nil {
		log.Fatalf("filepath.Abs err: %v", err)
	}
	err = icefsServer.IcefsServerInit()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterIcefsServer(s, &icefsServer)
	log.Printf("server listening at %v", ln.Addr())
	if err = s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

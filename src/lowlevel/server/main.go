/*
 * @Author: Tan90degrees tangentninetydegrees@gmail.com
 * @Date: 2023-03-11 05:57:15
 * @LastEditors: Tan90degrees tangentninetydegrees@gmail.com
 * @LastEditTime: 2023-04-19 08:12:26
 * @FilePath: /icefs/src/lowlevel/server/main.go
 * @Description:
 *
 * Copyright (C) 2023 Tan90degrees <tangentninetydegrees@gmail.com>.
 */
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"icefs-server/icefserror"
	pb "icefs-server/icefsgrpc"
	"icefs-server/icefsoperators"
	"icefs-server/icefsthrift"
	"log"
	"net"
	"path/filepath"

	"github.com/apache/thrift/lib/go/thrift"
	"google.golang.org/grpc"
)

const (
	ICEFS_GRPC_WAY   = "gRpc"
	ICEFS_THRIFT_WAY = "thrift"
)

func main() {
	var err error

	rpcPort := flag.Uint("port", 10086, "The port of icefs server")
	srvPath := flag.String("path", ".", "The path to serve at")
	linkWay := flag.String("way", "gRpc", "The way of data transmission")
	openTls := flag.Bool("tls", false, "Whether to enable the tls function")
	flag.Parse()

	addr := fmt.Sprintf("0.0.0.0:%d", *rpcPort)

	var icefsServer icefsoperators.IcefsServer
	icefsServer.RootPathAbs, err = filepath.Abs(*srvPath)
	if err != nil {
		log.Fatalf("filepath.Abs err: %v", err)
	}
	err = icefsServer.IcefsServerInit()
	if err != nil {
		log.Fatal(err)
	}

	switch *linkWay {
	case ICEFS_GRPC_WAY:
		var ln net.Listener
		ln, err = net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		server := grpc.NewServer()
		pb.RegisterIcefsGRpcServer(server, &icefsServer.GRpcServer)
		log.Printf("Server is running on %v", addr)
		if err = server.Serve(ln); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	case ICEFS_THRIFT_WAY:
		var cert tls.Certificate
		var socket thrift.TServerTransport
		if *openTls {
			cfg := new(tls.Config)
			if cert, err = tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
				cfg.Certificates = append(cfg.Certificates, cert)
			} else {
				log.Fatal(err)
			}
			socket = icefserror.Must(thrift.NewTSSLServerSocket(addr, cfg)).(*thrift.TSSLServerSocket)
		} else {
			socket = icefserror.Must(thrift.NewTServerSocket(addr)).(*thrift.TServerSocket)
		}
		fmt.Printf("%T\n", socket)
		processor := icefsthrift.NewIcefsThriftProcessor(&icefsServer.ThriftServer)
		server := thrift.NewTSimpleServer2(processor, socket)
		log.Printf("Server is running on %v", addr)
		if err = server.Serve(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	default:
		log.Fatalf("The parameter of \"--way\" should be %s or %s\n", ICEFS_GRPC_WAY, ICEFS_THRIFT_WAY)
	}
}

package main

import (
	"flag"
	"net"

	"google.golang.org/grpc"

	"github.com/thinxer/rlistener"
	pb "github.com/thinxer/rlistener/proto"
)

var (
	flagBind   = flag.String("bind", ":2222", "Bind address")
	flagAccept = flag.String("acceptor", ":2223", "Acceptor bind address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *flagBind)
	if err != nil {
		panic(err)
	}
	server := rlistener.NewServer(lis)

	acceptLis, err := net.Listen("tcp", *flagAccept)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterRemoteListenerServer(s, server)
	panic(s.Serve(acceptLis))
}

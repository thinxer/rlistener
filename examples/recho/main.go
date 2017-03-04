package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"

	"github.com/thinxer/rlistener"
)

var (
	flagRemote = flag.String("remote", "localhost:2223", "rlistener remote")
)

func main() {
	flag.Parse()

	lis, err := rlistener.Dial(*flagRemote, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go func(conn net.Conn) {
			fmt.Fprintf(conn, "Hi, you are from %v.\n", conn.RemoteAddr())
			io.Copy(conn, conn)
		}(conn)
	}
}

package rlistener

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	pb "github.com/thinxer/rlistener/proto"
)

type server struct {
	conns chan net.Conn
}

func (r *server) Accept(s pb.RemoteListener_AcceptServer) error {
	select {
	case <-s.Context().Done():
		return s.Context().Err()
	case upstream, ok := <-r.conns:
		if !ok {
			return grpc.Errorf(codes.Unavailable, "Listener closed.")
		}

		s.SendHeader(metadata.Pairs("X-Forwarded-For", upstream.RemoteAddr().String()))
		go func() {
			for {
				d, err := s.Recv()
				if err != nil {
					upstream.Close()
					return
				}
				_, err = upstream.Write(d.Data)
				if err != nil {
					s.SetTrailer(metadata.Pairs("X-Close-Write", err.Error()))
					return
				}
			}
		}()

		buf := make([]byte, 4096)
		for {
			n, err := upstream.Read(buf)
			if err != nil {
				s.SetTrailer(metadata.Pairs("X-Close-Read", err.Error()))
				return nil
			}
			err = s.Send(&pb.Buffer{Data: buf[:n]})
			if err != nil {
				upstream.Close()
				return nil
			}
		}
	}
}

func NewServer(lis net.Listener) pb.RemoteListenerServer {
	conns := make(chan net.Conn)
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				close(conns)
				return
			}
			conns <- conn
		}
	}()
	return &server{conns: conns}
}

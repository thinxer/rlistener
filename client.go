package rlistener

import (
	"context"
	"errors"
	"io"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/thinxer/rlistener/proto"
)

type addr string

func (a addr) Network() string { return "rlistener" }
func (a addr) String() string  { return "rlistener:" + string(a) }

type conn struct {
	io.ReadWriteCloser

	local, remote net.Addr
}

func (c *conn) LocalAddr() net.Addr  { return c.local }
func (c *conn) RemoteAddr() net.Addr { return c.remote }

func (c *conn) SetDeadline(time.Time) error      { return errors.New("unimplemneted") }
func (c *conn) SetReadDeadline(time.Time) error  { return errors.New("unimplemneted") }
func (c *conn) SetWriteDeadline(time.Time) error { return errors.New("unimplemneted") }

type listener struct {
	target string
	conn   *grpc.ClientConn
	client pb.RemoteListenerClient
}

func Dial(target string, options ...grpc.DialOption) (net.Listener, error) {
	conn, err := grpc.Dial(target, options...)
	if err != nil {
		return nil, err
	}
	return &listener{
		target: target,
		conn:   conn,
		client: pb.NewRemoteListenerClient(conn),
	}, nil
}

type stream struct {
	ctx    context.Context
	cancel context.CancelFunc

	c pb.RemoteListener_AcceptClient

	rbuf []byte
}

func (s *stream) Read(b []byte) (n int, err error) {
	if len(s.rbuf) > 0 {
		n = copy(b, s.rbuf)
		s.rbuf = s.rbuf[n:]
		return n, nil
	}

	msg, err := s.c.Recv()
	if err != nil {
		return 0, err
	}

	n = copy(b, msg.Data)
	s.rbuf = msg.Data[n:]
	return n, nil
}

func (s *stream) Write(b []byte) (int, error) {
	err := s.c.Send(&pb.Buffer{Data: b})
	if err != nil {
		return 0, err
	}
	return len(b), err
}

func (s *stream) Close() error {
	s.cancel()
	return nil
}

func (l *listener) Accept() (net.Conn, error) {
	ctx, cancel := context.WithCancel(context.Background())

	c, err := l.client.Accept(ctx)
	if err != nil {
		return nil, err
	}

	md, err := c.Header()
	if err != nil {
		cancel()
		return nil, err
	}
	var remote string
	for _, v := range md["x-forwarded-for"] {
		remote = v
	}
	return &conn{
		ReadWriteCloser: &stream{ctx, cancel, c, nil},
		local:           addr(l.target),
		remote:          addr(remote),
	}, nil
}

func (l *listener) Close() error {
	l.conn.Close()
	return nil
}

func (l *listener) Addr() net.Addr {
	return addr(l.target)
}

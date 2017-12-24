package goproxy


import (
	"bufio"
	//"encoding/base64"
	"net"
	//"strings"
	"github.com/golang/glog"
)

type Server struct {
	listener   net.Listener
	addr       string
}

func NewServer(Addr string) *Server {
	glog.V(4).Infoln("create server")
	return  &Server{
		addr: Addr, }

}

func (s *Server) Start() {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		glog.Fatal(err)
	}


	glog.Infof("proxy listen in %s, waiting for connection...\n", s.addr)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			glog.Error(err)
			continue
		}
		go s.newConn(conn).serve()
	}
}


func (s *Server) newConn(rwc net.Conn) *conn {
	return &conn{
		server: s,
		rwc:    rwc,
		brc:    bufio.NewReader(rwc),
	}
}
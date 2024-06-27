package p2p

import (
	"net"
	"sync"
)

type Peer struct {
	conn net.Conn
}

type Server struct {
	mu    sync.Mutex
	peers map[net.Addr]*Peer
}

func NewServer() *Server {
	return &Server{
		peers: make(map[net.Addr]*Peer),
	}
}

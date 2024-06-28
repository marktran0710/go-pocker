package p2p

import (
	"net"
	"sync"
)

type Peer struct {
	conn net.Conn
}

type ServerConfig struct {
	ListenAddr string
}

type Server struct {
	ServerConfig
	mu      sync.Mutex
	peers   map[net.Addr]*Peer
	addPeer chan *Peer
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
	for {
		select {}
	}
}

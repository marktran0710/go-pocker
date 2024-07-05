package p2p

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
)

type TCPTransport struct {
}

type Peer struct {
	conn net.Conn
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}

type ServerConfig struct {
	Version    string
	ListenAddr string
}

type Message struct {
	Payload io.Reader
	From    net.Addr
}

type Server struct {
	ServerConfig

	handler  Handler
	listener net.Listener
	mu       sync.Mutex
	peers    map[net.Addr]*Peer
	addPeer  chan *Peer
	delPeer  chan *Peer
	msgCh    chan *Message
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		handler:      &DefaultHandler{},
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		delPeer:      make(chan *Peer),
		msgCh:        make(chan *Message),
	}
}

func (s *Server) Start() {
	go s.loop()

	if err := s.listen(); err != nil {
		panic(err)
	}

	fmt.Printf("game server running on port %s\n", s.ListenAddr)

	s.acceptLoop()
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}

		peer := &Peer{
			conn: conn,
		}

		s.addPeer <- peer
		peer.Send([]byte(s.Version))

		go s.handleConn(peer)
	}
}

func (s *Server) handleConn(p *Peer) {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			break
		}

		s.msgCh <- &Message{
			From:    p.conn.RemoteAddr(),
			Payload: bytes.NewReader(buf[:n]),
		}

	}

	s.delPeer <- p
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return (err)
	}

	s.listener = ln

	return nil
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.delPeer:
			delete(s.peers, peer.conn.RemoteAddr())
			fmt.Printf("player disconnected %s\n", peer.conn.RemoteAddr())

		case peer := <-s.addPeer:
			fmt.Printf("new player connected %s\n", peer.conn.RemoteAddr())
			s.peers[peer.conn.RemoteAddr()] = peer
		case msg := <-s.msgCh:
			if err := s.handler.HandleMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}

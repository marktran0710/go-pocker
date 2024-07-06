package main

import (
	"github.com/marktran77/go-pocker/p2p"
)

func main() {

	// for i := 0; i < 10; i++ {
	// 	d := deck.New()
	// 	fmt.Println(d)
	// 	fmt.Println()
	//}

	cfg := p2p.ServerConfig{
		Version:    "go-pocker v0.1-alpha",
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	remoteCfg := p2p.ServerConfig{
		Version:    "go-pocker v0.1-alpha",
		ListenAddr: ":4000",
	}
	remoteServer := p2p.NewServer(remoteCfg)
	remoteServer.Start()
}

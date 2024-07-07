package main

import (
	"fmt"
	"time"

	"github.com/marktran77/go-pocker/deck"
	"github.com/marktran77/go-pocker/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:    "go-pocker v0.1-alpha",
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	time.Sleep(1 * time.Second)

	remoteCfg := p2p.ServerConfig{
		Version:    "go-pocker v0.1-alpha",
		ListenAddr: ":4000",
	}
	remoteServer := p2p.NewServer(remoteCfg)
	go remoteServer.Start()
	if err := remoteServer.Connect(":3000"); err != nil {
		fmt.Println(err)
	}

	fmt.Print(deck.New())
	select {}
}

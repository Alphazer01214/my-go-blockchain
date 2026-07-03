package main

import (
	"my-go-blockchain/networks"
	"time"
)

// server
// transport tcp udp
// block
// transaction
// keypair
func main() {
	local := networks.NewLocalTransport("local")
	remote := networks.NewLocalTransport("remote")

	if err := local.Connect(remote); err != nil {
		panic(err)
	}
	if err := remote.Connect(local); err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			err := remote.SendMessage(local.Addr(), []byte("hello world"))
			if err != nil {
				return
			}
			time.Sleep(time.Second * 2)
		}
	}()
	cfg := &networks.ServerConfig{
		Transports: []networks.Transport{local, remote},
	}
	server := networks.NewServer(cfg)
	server.Start()
}

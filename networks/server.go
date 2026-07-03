package networks

import (
	"fmt"
	"time"
)

type ServerConfig struct {
	Transports []Transport
}
type Server struct {
	ServerConfig
	rpcChan  chan RPC
	quitChan chan struct{}
}

func NewServer(cfg *ServerConfig) *Server {
	return &Server{
		ServerConfig: *cfg,
		rpcChan:      make(chan RPC),
		quitChan:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransport()
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case rpc := <-s.rpcChan:
			fmt.Println(rpc)

		case <-s.quitChan:
			break
		case <-ticker.C:
			fmt.Println("ticker")

		}
	}
}

func (s *Server) initTransport() {
	for _, t := range s.Transports {
		// read data
		go func(t Transport) {
			for rpc := range t.Consume() {
				fmt.Println(rpc)
				// handle
			}
		}(t)
	}
}

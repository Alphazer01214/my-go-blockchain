package networks

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	a := NewLocalTransport("a")
	b := NewLocalTransport("b")

	if err := a.Connect(b); err != nil {
		panic(err)
	}
	if err := b.Connect(a); err != nil {
		panic(err)
	}
	if a.peers[b.addr] != b || b.peers[a.addr] != a {
		panic("peer not connected")
	}
}

func TestSendMessage(t *testing.T) {
	a := NewLocalTransport("a")
	b := NewLocalTransport("b")

	if err := a.Connect(b); err != nil {
		panic(err)
	}
	if err := b.Connect(a); err != nil {
		panic(err)
	}

	msg := []byte("hello world")
	// a sends to b
	a.SendMessage(b.addr, msg)
	// if b consumes
	rpc := <-b.Consume()

	fmt.Println(rpc)
}

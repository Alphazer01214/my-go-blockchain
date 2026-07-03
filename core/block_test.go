package core

import (
	"bytes"
	"fmt"
	"my-go-blockchain/types"
	"testing"
	"time"
)

func TestHeader(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevHash:  types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     1145,
	}
	fmt.Println("[test header] header: ", h)

	buffer := &bytes.Buffer{}
	h.EncodeBinary(buffer)
	fmt.Println("[test header] encode binary: ", buffer.Bytes())
	d := &Header{}
	d.DecodeBinary(buffer)
	fmt.Println("[test header] decode binary: ", d)
}

func TestBlock(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevHash:  types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     1145,
		},
		Transactions: []Transaction{},
	}

	fmt.Println("[test block] block: ", b)
	buffer := &bytes.Buffer{}
	b.EncodeBinary(buffer)
	fmt.Println("[test block] encode binary: ", buffer.Bytes())
	d := &Block{}
	d.DecodeBinary(buffer)
	fmt.Println("[test block] decode binary: ", d)
}

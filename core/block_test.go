package core

import (
	"my-go-blockchain/types"
	"testing"
	"time"
)

func TestHeader(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevHash:  types.RandomHash(),
		Timestamp: time.Now(),
		Height:    1,
		Nonce:     1,
	}

}

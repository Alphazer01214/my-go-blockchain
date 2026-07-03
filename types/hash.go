package types

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

type Hash [32]uint8

func (h Hash) IsEmpty() bool {
	for _, b := range h {
		if b != 0 {
			return false
		}
	}
	return true
}

func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

func (h Hash) ToSlice() []byte {
	return h[:]
}
func GetHashFromByte(data []byte) Hash {
	if len(data) != 32 {
		msg := fmt.Sprintf("data len is %d, want 32", len(data))
		panic(msg)
	}

	var value [32]uint8
	for i := range data {
		value[i] = data[i]
	}
	return value
}

func RandomHash() Hash {
	var value [32]uint8
	for i := range value {
		value[i] = uint8(rand.Uint32())
	}
	return value
}

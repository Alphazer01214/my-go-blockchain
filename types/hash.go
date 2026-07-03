package types

import (
	"fmt"
	"math/rand"
)

type Hash [32]uint8

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

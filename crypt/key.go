package crypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func NewPrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func (pk PrivateKey) GetPublicKey() PublicKey {
	return PublicKey{
		key: &pk.key.PublicKey,
	}
}

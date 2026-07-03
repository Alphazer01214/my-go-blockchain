package crypt

import "testing"

func TestKey(t *testing.T) {
	pk := NewPrivateKey()
	pk.GetPublicKey()
}

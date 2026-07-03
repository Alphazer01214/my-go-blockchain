package core

import (
	"encoding/binary"
	"io"
	"my-go-blockchain/types"
	"time"
)

type Header struct {
	Version   uint32
	PrevHash  types.Hash
	Timestamp time.Time
	Height    uint32
	Nonce     uint64
}

// hash the header in binary

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.PrevHash); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.Height); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.Nonce); err != nil {
		return err
	}
	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevHash); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Nonce); err != nil {
		return err
	}
	return nil
}

type Block struct {
	Header
	Transactions []Transaction
}

func NewBlock(header *Header, transactions []Transaction) *Block {
	return &Block{
		Header:       *header,
		Transactions: transactions,
	}
}

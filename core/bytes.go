package core

import "encoding/binary"

func UInt8ToBytes(in uint8) []byte {
	b := make([] byte, 1)
	b[0] = in
	return b
}

func UInt16BeToBytes(in uint16) []byte {
	b := make([] byte, 2)
	b[1] = byte(in >> 8)
	b[0] = byte(in)
	return b
}

func UInt32BeToBytes(in uint32) []byte {
	b := make([] byte, 4)
	binary.BigEndian.PutUint32(b, in)
	return b
}

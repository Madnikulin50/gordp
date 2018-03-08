package core

import (
	"testing"
	"bytes"
)

type TestStruct struct {
	b byte
	u8 uint8
	u16le uint16
	u16be uint16
	u32le uint32
	u32be uint32
}
func (str *TestStruct) Write(writer Writer) error {
	WriteByte(str.b, writer)
	WriteUInt8(str.u8, writer)
	WriteUInt16LE(str.u16le, writer)
	WriteUInt16BE(str.u16be, writer)
	WriteUInt32LE(str.u32le, writer)
	WriteUInt32BE(str.u32be, writer)
	return nil
}
func (str *TestStruct) Read(reader Reader) error {
	str.b = ReadByte(reader)
	str.u8 = ReadUInt8(reader)
	str.u16le = ReadUInt16LE(reader)
	str.u16be = ReadUInt16BE(reader)
	str.u32le = ReadUInt32LE(reader)
	str.u32be = ReadUInt32BE(reader)
	return nil
}

func TestReadWriteStruct(t *testing.T) {
	d1 := TestStruct { 0x23, 45, 345, 789, 99997676, 48446969 }
	b := bytes.NewBuffer(nil)
	d1.Write(b)
	r := bytes.NewReader(b.Bytes())
	d2 := TestStruct{}
	d2.Read(r)

	if d1.b != d2.b {
		t.Error("b != b")
	}

	t.Log("all ok", d2.u8)
}

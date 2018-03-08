package core

import (
	"encoding/binary"
	"io"
	"errors"
)

type Reader interface {
	Read([]byte) (int, error)
}

type LimitedReader struct {
	reader Reader
	needRead int
}

func (reader* LimitedReader) Read(buffer []byte) (int, error) {
	if reader.needRead == 0 {
		return 0, io.EOF
	}
	need := len(buffer)
	if need > reader.needRead {
		res, err := reader.Read(buffer)
		reader.needRead -= res
		return res, err
	}
	need = reader.needRead
	reader.needRead = 0
	return reader.Read(buffer[:need])
}

func (reader* LimitedReader) GetNeedRead() int  {
	return reader.needRead
}

func NewLimitedReader(reader Reader, needRead int) *LimitedReader {
	return &LimitedReader{reader, needRead}
}
/*
func AvailableLength(reader *Reader) int {
	bytes, err := reader.Peek(0)
	if err != nil {
		return 0
	}
	return len(bytes)
}*/

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Readable interface{
	Read(r Reader) error
}

type Writable interface {
	Write(w Writer) error
}

func AllocAndReadBytes(len int, r Reader) []byte {
	b := make([] byte, len)
	io.ReadFull(r, b)
	return b
}

func ReadBytes(b []byte, r Reader) (int, error) {
	return io.ReadFull(r, b)
}

func WriteBytes(b []byte, w Writer) (int, error) {
	return w.Write(b)
}

type ReadBytesComplete func (result [] byte, err error)

func StartReadBytes (b []byte, r Reader, cb ReadBytesComplete) {
	go func () {
		_, err := io.ReadFull(r, b)
		cb(b, err)
	}()
}

type OnBytes func (r Reader, err error)

func WaitData(r Reader, cb OnBytes) {
	go func () {
		_, err := io.ReadFull(r, b)
		cb(b, err)
	}()
}

func StartAllocAndReadBytes (length int, r Reader, cb ReadBytesComplete) {
	b := make([] byte, length)

	go func () {
		_, err := io.ReadFull(r, b)
		cb(b, err)
	}()
}

func WriteUInt8(data uint8, w Writer) {
	b := make([] byte, 1)
	b[0] = byte(data)
	w.Write(b)
}



func ReadUInt8(r Reader) (uint8) {
	b := AllocAndReadBytes(1, r)
	return uint8(b[0])
}

func WriteByte(data byte, w Writer) {
	b := make([] byte, 1)
	b[0] = byte(data)
	w.Write(b)
}


func ReadByte(r Reader) (byte) {
	b := AllocAndReadBytes(1, r)
	return b[0]
}

func ReadPadding(length int, r Reader) {
	AllocAndReadBytes(length, r)
}

func WritePadding(length int, w Writer) {
	b := make([] byte, length)
	w.Write(b)
}


func WriteUInt16LE(data uint16, w Writer) {
	b := make([] byte, 2)
	b[0] = byte(data >> 8)
	b[1] = byte(data)
	w.Write(b)
}


func ReadUInt16BE(r Reader) (uint16) {
	b := AllocAndReadBytes(2, r)
	return uint16(b[1]) << 8 + uint16(b[0])
}

func WriteUInt16BE(data uint16, w Writer) {
	b := make([] byte, 2)
	b[1] = byte(data >> 8)
	b[0] = byte(data)
	w.Write(b)
}


func ReadUInt16LE(r Reader) (uint16) {
	b := AllocAndReadBytes(2, r)
	return uint16(b[0]) << 8 + uint16(b[1])
}

func WriteUInt32LE(data uint32, w Writer) {
	b := make([] byte, 4)
	binary.LittleEndian.PutUint32(b, data)
	w.Write(b)
}


func ReadUInt32LE(r Reader) (uint32) {
	b := AllocAndReadBytes(4, r)
	return binary.LittleEndian.Uint32(b)
}


func ReadUInt32BE(r Reader) (uint32) {
	b := AllocAndReadBytes(4, r)
	return binary.BigEndian.Uint32(b)
}


func WriteUInt32BE(data uint32, w Writer) {
	b := make([] byte, 4)
	binary.BigEndian.PutUint32(b, data)
	w.Write(b)
}



type Data interface {
	Write(writer Writer) error
	Read(reader Reader) error
}

type LengthCalculator struct {
	len int
}

func (calc *LengthCalculator) Write(p []byte) (n int, err error) {
	l := len(p)
	calc.len += l
	return l, nil
}


func CalcDataLength(obj Writable) int {
	calc := &LengthCalculator{0}
	obj.Write(calc)
	return calc.len
}

type Component struct {
	Opt interface{}
}

func NewComponent(opt interface{}) *Component{
	return &Component{opt}
}

func (c *Component) Write(writer Writer) error {
	return errors.New("not implemented")
}

func (c *Component) Read(reader Reader) error {
	return errors.New("not implemented")
}


type ComponentOption struct {
	readLength uint16
	constant bool
	optional bool
}

func NewComponentOption(readLength uint16, constant bool, optional bool) *ComponentOption {
	return &ComponentOption{readLength, constant, optional}
}

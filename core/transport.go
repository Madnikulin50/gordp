package core

import "github.com/chuckpreslar/emission"

type Transport interface {
	GetEvents() *emission.Emitter
	Expect(int)
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close()
	Error()
}
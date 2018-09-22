package core

type Transport interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
	Error() error
}
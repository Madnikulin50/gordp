package core

import (
	"net"
	"crypto/tls"
)

type Layer struct {
	conn net.Conn
	tlsConn *tls.Conn
}

func NewLayer(conn net.Conn) *Layer {
	bl := Layer{conn, nil}
	return &bl
}

func (layer Layer) Read(b []byte) (n int, err error) {
	if layer.tlsConn != nil {
		return layer.tlsConn.Read(b)
	}
	return layer.conn.Read(b)
}

/**
 * Call tcp socket to write stream
 * @param {type.Type} packet
 */
func (layer Layer)Write(b []byte) (n int, err error) {
	if layer.tlsConn != nil {
		return layer.tlsConn.Write(b)
	}
	return layer.conn.Write(b)
}


func (layer *Layer) StartTLS() error {
	config := &tls.Config{}
	layer.tlsConn = tls.Client(layer.conn, config)
	return nil
}


func (layer *Layer) Close () error {
	if layer.tlsConn != nil {
		err := layer.tlsConn.Close()
		if err != nil {
			return err
		}
	}
	return layer.conn.Close()
}



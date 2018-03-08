package gordp

import (
	"log"
	"net"
	"errors"
)

type RdpConnectionParams struct {
	Connection string
	User string
	Password string
}

type RdpBitmap struct {
	DestTop int
	DestLeft int
	DestBottom int
	DestRight int
	Width int
	Height int
	BitsPerPixel int
	IsCompress bool
	Data []byte
}


type RdpConnectionBase struct {
	params *RdpConnectionParams
	conn *net.Conn
}

func NewRdpConnection(params *RdpConnectionParams) *RdpConnectionBase {
	return &RdpConnectionBase{params, nil}
}

func (con *RdpConnectionBase) Connect(params *RdpConnectionParams) error {
	return errors.New("Not implemented")
}

func (con *RdpConnectionBase) Close() error {
	return errors.New("Not implemented")
}

func (con *RdpConnectionBase) SendPointerEvent(x int, y int, button int, pressed bool) error {
	return errors.New("Not implemented")
}

func (con *RdpConnectionBase) SendWheelEvent(x int, y int, step int, isNegative bool, isHorizontal bool) error {
	return errors.New("Not implemented")
}

func (con *RdpConnectionBase) SendKeyEventScancode(code int, isPressed bool) error {
	return errors.New("Not implemented")
}

func (con *RdpConnectionBase) SendKeyEventUnicode(code int, isPressed bool) error {
	return errors.New("Not implemented")
}

func (con *RdpConnectionBase) OnConnect() (bool, error) {
	log.Print("OnConnect")
	return true, nil
}
func (con *RdpConnectionBase) OnBitmap(bitmap *RdpBitmap) {
	log.Printf("OnBitmap %d %d %d %d", bitmap.DestLeft, bitmap.DestTop, bitmap.DestRight, bitmap.DestBottom)
}

func (con *RdpConnectionBase) OnClose() {
	log.Print("OnClose")
}

func (con *RdpConnectionBase) OnError(err error) {
	log.Print("OnError")
}

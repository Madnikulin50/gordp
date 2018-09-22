package x224

import (
	"github.com/chuckpreslar/emission"
	"../../core"
	"errors"
	"bytes"
)

/**
 * Message type present in X224 packet header
*/
type MessageType byte
const (
TPDU_CONNECTION_REQUEST MessageType = 0xE0
TPDU_CONNECTION_CONFIRM = 0xD0
TPDU_DISCONNECT_REQUEST = 0x80
TPDU_DATA = 0xF0
TPDU_ERROR = 0x70
)

/**
 * Type of negotiation present in negotiation packet
 */
type NegotiationType byte
const (
	TYPE_RDP_NEG_REQ NegotiationType = 0x01
	TYPE_RDP_NEG_RSP = 0x02
	TYPE_RDP_NEG_FAILURE = 0x03
)

/**
 * Protocols available for x224 layer
*/
type Protocol uint32
const (
	PROTOCOL_RDP Protocol = 0x00000000
	PROTOCOL_SSL = 0x00000001
	PROTOCOL_HYBRID = 0x00000002
	PROTOCOL_HYBRID_EX = 0x00000008
)

/**
 * Use to negotiate security layer of RDP stack
 * In node-rdpjs only ssl is available
 * @param opt {object} component type options
 * @see request -> http://msdn.microsoft.com/en-us/library/cc240500.aspx
 * @see response -> http://msdn.microsoft.com/en-us/library/cc240506.aspx
 * @see failure ->http://msdn.microsoft.com/en-us/library/cc240507.aspx
*/

type Negotiation struct {
	Type NegotiationType
	Flag uint8
	Length uint16
	Result uint32
}

func NewNegotiation() *Negotiation {
	return &Negotiation{0, 0,0x0008 /*constant*/, uint32(PROTOCOL_RDP) }
}

func (x *Negotiation) Write(w core.Writer) {
	core.WriteByte(byte(x.Type), w)
	core.WriteUInt8(x.Flag, w)
	core.WriteUInt16LE(x.Length, w)
	core.WriteUInt32LE(x.Result, w)
}

func (x *Negotiation) Read(r core.Reader) error {
	var err error
	b, err := core.ReadByte(r)
	x.Type = NegotiationType(b)
	x.Flag, err = core.ReadUInt8(r)
	x.Length, err = core.ReadUInt16LE(r)
	x.Result, err = core.ReadUInt32LE(r)
	if x.Length == 0x0008 {
		return errors.New("invalid x224 negoitiate")
	}
	return err
}

/**
 * X224 client connection request
 * @param opt {object} component type options
 * @see	http://msdn.microsoft.com/en-us/library/cc240470.aspx
*/
type ClientConnectionRequestPDU struct {
	Len uint8
	Code MessageType
	Padding1 uint16
	Padding2 uint16
	Padding3 uint8
	Cookie []byte
	ProtocolNeg Negotiation
	//CorrelationInfo [36]byte
}

func NewClientConnectionRequestPDU(coockie []byte) *ClientConnectionRequestPDU {
	x := ClientConnectionRequestPDU{ 0, TPDU_CONNECTION_REQUEST, 0,0,0,
		coockie, *NewNegotiation()/*, [36]byte{}*/ }
	x.Len = uint8(core.CalcDataLength(&x) - 1)
	return &x
}


/*
function clientConnectionRequestPDU(opt, cookie) {
var self = {
len : uint8(function() {
return new type.Component(self).size() - 1;
}),
code : uint8(MessageType.TPDU_CONNECTION_REQUEST, { constant : true }),
padding : new type.Component([new type.UInt16Le(), new type.UInt16Le(), uint8()]),
cookie : cookie || new type.Factory( function (s) {
var offset = 0;
while (true) {
var token = s.buffer.readUInt16LE(s.offset + offset);
if (token === 0x0a0d) {
self.cookie = new type.BinaryString(null, { readLength : new type.CallableValue(offset + 2) }).read(s);
return;
}
else {
offset += 1;
}
}
}, { conditional : function () {
return self.len.value > 14;
}}),
protocolNeg : negotiation({ optional : true })
};

return new type.Component(self, opt);
}*/

func (x *ClientConnectionRequestPDU) Write(w core.Writer) error {
	core.WriteUInt8(x.Len, w)
	core.WriteUInt8(uint8(x.Code), w)
	core.WriteUInt16LE(x.Padding1, w)
	core.WriteUInt16LE(x.Padding2, w)
	core.WriteUInt8(x.Padding3, w)
	w.Write(x.Cookie)
	core.WriteUInt16LE(0x0a0d, w)
	x.ProtocolNeg.Write(w)
	return nil
}

func (x *ClientConnectionRequestPDU) Read(r core.Reader) error {
	var err error
	x.Len, err = core.ReadUInt8(r)
	lr := core.NewLimitedReader(r, int(x.Len))
	t, err := core.ReadUInt8(lr)
	x.Code = MessageType(t)
	x.Padding1, err = core.ReadUInt16LE(lr)
	x.Padding2, err = core.ReadUInt16LE(lr)
	x.Padding3, err = core.ReadUInt8(lr)
	b := make([]byte, 1)
	var prev byte = 0
	for {
		l, err := lr.Read(b)
		if l != 1 || err != nil {
			break
		}
		if prev == 0xd && b[0] == 0x0a {
			x.Cookie = x.Cookie[:len(x.Cookie) - 1]
			break
		}
		x.Cookie = append(x.Cookie, b[0])
		prev = b[0]
	}
	if lr.GetNeedRead() == 0 {
		return nil
	}
	x.ProtocolNeg.Read(lr)
	/*if lr.GetNeedRead() == 0 {
		return
	}
	x.CorrelationInfo = [36]byte(core.ReadBytes(36, lr))*/
	return err
}

/**
 * X224 Server connection confirm
 * @param opt {object} component type options
 * @see	http://msdn.microsoft.com/en-us/library/cc240506.aspx
*/
type ServerConnectionConfirm struct {
	Len uint8
	Code MessageType
	Padding1 uint16
	Padding2 uint16
	Padding3 uint8
	ProtocolNeg Negotiation
}

func NewServerConnectionConfirm() *ServerConnectionConfirm {
	x := ServerConnectionConfirm{ 0, TPDU_CONNECTION_CONFIRM, 0,0,0,*NewNegotiation() }
	x.Len = uint8(core.CalcDataLength(&x) - 1)
	return &x
}

func (x *ServerConnectionConfirm) Write(w core.Writer) error {
	core.WriteUInt8(x.Len, w)
	core.WriteUInt8(uint8(x.Code), w)
	core.WriteUInt16LE(x.Padding1, w)
	core.WriteUInt16LE(x.Padding2, w)
	core.WriteUInt8(x.Padding3, w)
	return nil
	}

func (x *ServerConnectionConfirm) Read(r core.Reader) error {
	var err error
	x.Len, err = core.ReadUInt8(r)
	lr := core.NewLimitedReader(r, int(x.Len))
	b, err := core.ReadUInt8(lr)
	x.Code = MessageType(b)
	x.Padding1, err = core.ReadUInt16LE(lr)
	x.Padding2, err = core.ReadUInt16LE(lr)
	x.Padding3, err = core.ReadUInt8(lr)

	if lr.GetNeedRead() == 0 {
		return nil
	}
	x.ProtocolNeg.Read(lr)
	return err
}

/*
function serverConnectionConfirm(opt) {
var self = {
len = uint8(function() {
return new type.Component(self).size() - 1;
}),
code = uint8(MessageType.TPDU_CONNECTION_CONFIRM, { constant = true }),
padding = new type.Component([new type.UInt16Le(), new type.UInt16Le(), uint8()]),
protocolNeg = negotiation({ optional = true })
};

return new type.Component(self, opt);
}
*/
/**
 * Header of each data message from x224 layer
 * @returns {type.Component}
*/
type DataHeader struct {
	Header uint8
	MessageType MessageType
	Separator uint8
}

func NewDataHeader() *DataHeader {
	return &DataHeader{2, TPDU_DATA /* constant */, 0x80 /*constant*/}
}

func (xdh *DataHeader) Write(w core.Writer) error {
	core.WriteUInt8(xdh.Header, w)
	core.WriteByte(byte(xdh.MessageType), w)
	core.WriteUInt8(xdh.Separator, w)
	return nil
}

func (xdh *DataHeader) Read(r core.Reader) error {
	var err error
	xdh.Header, err = core.ReadUInt8(r)
	b, err := core.ReadByte(r)
	xdh.MessageType = MessageType(b)
	if xdh.MessageType == TPDU_DATA {
		return errors.New("invalid x224 dataheader")
	}
	xdh.Separator, err = core.ReadUInt8(r)
	if xdh.Separator == 0x80 {
		return errors.New("invalid x224 dataheader")
	}
	return err
}

/**
 * Common X224 Automata
 * @param presentation {Layer} presentation layer
*/
type X224 struct {
	emission.Emitter
	transport core.Transport
	requestedProtocol Protocol
	selectedProtocol  Protocol
}

func NewX224(transport core.Transport) *X224 {
	t := &X224{*emission.NewEmitter(), transport, PROTOCOL_SSL,
		PROTOCOL_SSL}
	transpEmitter, ok := interface{}(transport).(*emission.Emitter)
	if !ok {
		panic("Invalid transport type")
		}
	transpEmitter.On("close", func() {
		t.Emit("close")
	}).On("error",func() {
		t.Emit("error")
		})
	return t
}

//inherit from Layer
/**
 * Main data received function
 * after connection sequence
 * @param s {type.Stream} stream formated from transport layer
 */
func (x* X224) RecvData(r core.Reader) {
// check header
	NewDataHeader().Read(r)
	x.Emit("data", r)
}

/**
 * Format message from x224 layer to transport layer
 * @param message {type}
 * @returns {type.Component} x224 formated message
 */
func (x *X224) Send (message core.Writable) {
	NewDataHeader().Write(x.transport)
	message.Write(x.transport)

};

/**
 * Client x224 automata
 * @param transport {events.EventEmitter} (bind data events)
 */
type Client struct {
	X224
}

func NewClient(transport core.Transport) *Client {
	return &Client{*NewX224(transport)}
}


/**
 * Client automata connect event
*/
func (x *Client) Connect() {
	message := NewClientConnectionRequestPDU(make([] byte, 0))
	message.ProtocolNeg.Type = TYPE_RDP_NEG_REQ
	message.ProtocolNeg.Result = uint32(x.requestedProtocol)

	message.Write(x.transport)

// next state wait connection confirm packet
	go func() {
		x.recvConnectionConfirm(x.transport)
	} ()
}

/**
 * close stack
 */
func (x *Client) Close() error {
	return x.transport.Close()
}

/**
 * Receive connection from server
 * @param s {Stream}
*/
func (x *Client) recvConnectionConfirm(tr core.Transport) error {
	var message = NewServerConnectionConfirm()
	message.Read(tr)

	if message.ProtocolNeg.Type == TYPE_RDP_NEG_FAILURE {
		return errors.New("NODE_RDP_PROTOCOL_NEG_FAILURE")
	}

	if message.ProtocolNeg.Type == TYPE_RDP_NEG_RSP {
		x.selectedProtocol = Protocol(message.ProtocolNeg.Result)
	}

	if x.selectedProtocol == PROTOCOL_HYBRID || x.selectedProtocol == PROTOCOL_HYBRID_EX {
		return errors.New("NODE_RDP_PROTOCOL_NLA_NOT_SUPPORTED")
	}

	if x.selectedProtocol == PROTOCOL_RDP {
		core.Warn("RDP standard security selected")
		return nil
	}

	// finish connection sequence
	core.StartAllocAndReadBytes(3, tr, func(s []byte, err error) {
		rd := bytes.NewReader(s)
		x.RecvData(rd)
	})

	if x.selectedProtocol == PROTOCOL_SSL {
		core.Info("SSL standard security selected")
		/* TODO
		this.transport.startTLS(function() {
	self.emit('connect', self.selectedProtocol);
	});
	return; */
	}
	return nil
}

package protocol

import (
	"github.com/chuckpreslar/emission"
	"../core"
	"errors"
	"bytes"
)

/**
 * Message type present in X224 packet header
*/
type X224MessageType byte
const (
X224_TPDU_CONNECTION_REQUEST X224MessageType = 0xE0
X224_TPDU_CONNECTION_CONFIRM = 0xD0
X224_TPDU_DISCONNECT_REQUEST = 0x80
X224_TPDU_DATA = 0xF0
X224_TPDU_ERROR = 0x70
)

/**
 * Type of negotiation present in negotiation packet
 */
type X224NegotiationType byte
const (
	X224_TYPE_RDP_NEG_REQ X224NegotiationType = 0x01
	X224_TYPE_RDP_NEG_RSP = 0x02
	X224_TYPE_RDP_NEG_FAILURE = 0x03
)

/**
 * Protocols available for x224 layer
*/
type X224Protocol uint32
const (
	X224_PROTOCOL_RDP X224Protocol = 0x00000000
	X224_PROTOCOL_SSL = 0x00000001
	X224_PROTOCOL_HYBRID = 0x00000002
	X224_PROTOCOL_HYBRID_EX = 0x00000008
)

/**
 * Use to negotiate security layer of RDP stack
 * In node-rdpjs only ssl is available
 * @param opt {object} component type options
 * @see request -> http://msdn.microsoft.com/en-us/library/cc240500.aspx
 * @see response -> http://msdn.microsoft.com/en-us/library/cc240506.aspx
 * @see failure ->http://msdn.microsoft.com/en-us/library/cc240507.aspx
*/

type X224Negotiation struct {
	Type X224NegotiationType
	Flag uint8
	Length uint16
	Result uint32
}

func NewX224Negotiation() *X224Negotiation {
	return &X224Negotiation{0, 0,0x0008 /*constant*/, uint32(X224_PROTOCOL_RDP) }
}

func (x *X224Negotiation) Write(w core.Writer) {
	core.WriteByte(byte(x.Type), w)
	core.WriteUInt8(x.Flag, w)
	core.WriteUInt16LE(x.Length, w)
	core.WriteUInt32LE(x.Result, w)
}

func (x *X224Negotiation) Read(r core.Reader) {
	x.Type = X224NegotiationType(core.ReadByte(r))
	x.Flag = core.ReadUInt8(r)
	x.Length = core.ReadUInt16LE(r)
	x.Result = core.ReadUInt32LE(r)
	if x.Length == 0x0008 {
		panic("invalid x224 negoitiate")
	}
}

/**
 * X224 client connection request
 * @param opt {object} component type options
 * @see	http://msdn.microsoft.com/en-us/library/cc240470.aspx
*/
type X224ClientConnectionRequestPDU struct {
	Len uint8
	Code X224MessageType
	Padding1 uint16
	Padding2 uint16
	Padding3 uint8
	Cookie []byte
	ProtocolNeg X224Negotiation
	//CorrelationInfo [36]byte
}

func NewX224ClientConnectionRequestPDU(coockie []byte) *X224ClientConnectionRequestPDU {
	x := X224ClientConnectionRequestPDU{ 0, X224_TPDU_CONNECTION_REQUEST, 0,0,0,
		coockie, *NewX224Negotiation()/*, [36]byte{}*/ }
	x.Len = uint8(core.CalcDataLength(&x) - 1)
	return &x
}


/*
function clientConnectionRequestPDU(opt, cookie) {
var self = {
len : uint8(function() {
return new type.Component(self).size() - 1;
}),
code : uint8(MessageType.X224_TPDU_CONNECTION_REQUEST, { constant : true }),
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

func (x *X224ClientConnectionRequestPDU) Write(w core.Writer) error {
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

func (x *X224ClientConnectionRequestPDU) Read(r core.Reader) error {
	x.Len = core.ReadUInt8(r)
	lr := core.NewLimitedReader(r, int(x.Len))
	x.Code = X224MessageType(core.ReadUInt8(lr))
	x.Padding1 = core.ReadUInt16LE(lr)
	x.Padding2 = core.ReadUInt16LE(lr)
	x.Padding3 = core.ReadUInt8(lr)
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
	return nil
}

/**
 * X224 Server connection confirm
 * @param opt {object} component type options
 * @see	http://msdn.microsoft.com/en-us/library/cc240506.aspx
*/
type X224ServerConnectionConfirm struct {
	Len uint8
	Code X224MessageType
	Padding1 uint16
	Padding2 uint16
	Padding3 uint8
	ProtocolNeg X224Negotiation
}

func NewX224ServerConnectionConfirm() *X224ServerConnectionConfirm {
	x := X224ServerConnectionConfirm{ 0, X224_TPDU_CONNECTION_CONFIRM, 0,0,0,*NewX224Negotiation() }
	x.Len = uint8(core.CalcDataLength(&x) - 1)
	return &x
}

func (x *X224ServerConnectionConfirm) Write(w core.Writer) error {
	core.WriteUInt8(x.Len, w)
	core.WriteUInt8(uint8(x.Code), w)
	core.WriteUInt16LE(x.Padding1, w)
	core.WriteUInt16LE(x.Padding2, w)
	core.WriteUInt8(x.Padding3, w)
	return nil
	}

func (x *X224ServerConnectionConfirm) Read(r core.Reader) error {
	x.Len = core.ReadUInt8(r)
	lr := core.NewLimitedReader(r, int(x.Len))
	x.Code = X224MessageType(core.ReadUInt8(lr))
	x.Padding1 = core.ReadUInt16LE(lr)
	x.Padding2 = core.ReadUInt16LE(lr)
	x.Padding3 = core.ReadUInt8(lr)

	if lr.GetNeedRead() == 0 {
		return nil
	}
	x.ProtocolNeg.Read(lr)
	return nil
}

/*
function serverConnectionConfirm(opt) {
var self = {
len = uint8(function() {
return new type.Component(self).size() - 1;
}),
code = uint8(MessageType.X224_TPDU_CONNECTION_CONFIRM, { constant = true }),
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
type X224DataHeader struct {
	Header uint8
	MessageType X224MessageType
	Separator uint8
}

func NewX224DataHeader() *X224DataHeader {
	return &X224DataHeader{2, X224_TPDU_DATA /* constant */, 0x80 /*constant*/}
}

func (xdh *X224DataHeader) Write(w core.Writer) error {
	core.WriteUInt8(xdh.Header, w)
	core.WriteByte(byte(xdh.MessageType), w)
	core.WriteUInt8(xdh.Separator, w)
	return nil
}

func (xdh *X224DataHeader) Read(r core.Reader) error {
	xdh.Header = core.ReadUInt8(r)
	xdh.MessageType = X224MessageType(core.ReadByte(r))
	if xdh.MessageType == X224_TPDU_DATA {
		panic("invalid x224 dataheader")
	}
	xdh.Separator = core.ReadUInt8(r)
	if xdh.Separator == 0x80 {
		panic("invalid x224 dataheader")
	}
	return nil
}

/**
 * Common X224 Automata
 * @param presentation {Layer} presentation layer
*/
type X224 struct {
	emission.Emitter
	transport *core.Transport
	requestedProtocol X224Protocol
	selectedProtocol  X224Protocol
}

func NewX224(transport *core.Transport) *X224 {
	t := &X224{*emission.NewEmitter(), transport, X224_PROTOCOL_SSL,
		X224_PROTOCOL_SSL}
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
	X224DataHeader{}.Read(r)
	x.Emit("data", r)
}

/**
 * Format message from x224 layer to transport layer
 * @param message {type}
 * @returns {type.Component} x224 formated message
 */
func (x *X224) Send (message core.Writable) {
	NewX224DataHeader().Write(x.transport)
	message.Write(x.transport)

};

/**
 * Client x224 automata
 * @param transport {events.EventEmitter} (bind data events)
 */
type X224Client struct {
	X224
}

func NewX224Client(transport *core.Transport) *X224Client {
	return &X224Client{*NewX224(transport)}
}


/**
 * Client automata connect event
*/
func (x *X224Client) Connect() {
	message := NewX224ClientConnectionRequestPDU(make([] byte, 0))
	message.ProtocolNeg.Type = X224_TYPE_RDP_NEG_REQ
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
func (x *X224Client) Close() {
	x.transport.Close()
}

/**
 * Receive connection from server
 * @param s {Stream}
*/
func (x *X224Client) recvConnectionConfirm(tr *core.Transport) error {
	var message = NewX224ServerConnectionConfirm()
	message.Read(tr)

	if message.ProtocolNeg.Type == X224_TYPE_RDP_NEG_FAILURE {
		return errors.New("NODE_RDP_PROTOCOL_X224_NEG_FAILURE")
	}

	if message.ProtocolNeg.Type == X224_TYPE_RDP_NEG_RSP {
		x.selectedProtocol = X224Protocol(message.ProtocolNeg.Result)
	}

	if x.selectedProtocol == X224_PROTOCOL_HYBRID || x.selectedProtocol == X224_PROTOCOL_HYBRID_EX {
		return errors.New("NODE_RDP_PROTOCOL_X224_NLA_NOT_SUPPORTED")
	}

	if x.selectedProtocol == X224_PROTOCOL_RDP {
		core.Warn("RDP standard security selected")
		return nil
	}

	// finish connection sequence
	core.StartAllocAndReadBytes(3, tr, func(s []byte, err error) {
		rd := bytes.NewReader(s)
		x.RecvData(rd)
	})

	if x.selectedProtocol == X224_PROTOCOL_SSL {
		core.Info("SSL standard security selected")
		/* TODO
		this.transport.startTLS(function() {
	self.emit('connect', self.selectedProtocol);
	});
	return; */
	}
	return nil
}

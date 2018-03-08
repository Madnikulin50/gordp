package protocol

import (
	"github.com/chuckpreslar/emission"
	"../core"
	"bytes"
)

/**
 * Type of tpkt packet
 * Fastpath is use to shortcut RDP stack
 * @see http://msdn.microsoft.com/en-us/library/cc240621.aspx
 * @see http://msdn.microsoft.com/en-us/library/cc240589.aspx
 */
type TpktAction byte
const (
FASTPATH_ACTION_FASTPATH TpktAction = 0x0
FASTPATH_ACTION_X224 = 0x3
)

/**
 * TPKT layer of rdp stack
*/

type TPKT struct {
	emission.Emitter
	transport core.Transport
	secFlag byte
}

func NewTPKT(transport core.Transport) {
	t := &TPKT{*emission.NewEmitter(), transport, nil}
	core.StartAllocAndReadBytes(2, t.transport, t.recvHeader)
}

/**
 * inherit from a packet layer
*/

func (t *TPKT) GetEvents() *emission.Emitter {
	return &t.Emitter
}


/**
 * Receive correct packet as expected
 * @param s {type.Stream}
 */
func (t *TPKT) recvHeader(s []byte, err error) {
	if err != nil {
		t.Emit("error", err)
		return
	}
	version := s[0]
	if version == FASTPATH_ACTION_X224 {
		//new type.UInt8().read(s);
		core.StartAllocAndReadBytes(2, t.transport, t.recvExtendedHeader)
	} else {
		t.secFlag = (version >> 6) & 0x3
		length := int(s[1])
		if length &0x80 != 0 {
			core.StartAllocAndReadBytes(1, t.transport, func(s []byte, err error) {
				t.recvExtendedFastPathHeader(s, length, err)
			})
		} else {
			core.StartAllocAndReadBytes(length - 2, t.transport, func(s []byte, err error) {
				t.recvFastPath(s, err)
			})
		}
	}
};

/**
 * Receive second part of header packet
 * @param s {type.Stream}
 */
func (t *TPKT) recvExtendedHeader(s []byte, err error) {
	if err != nil {
		panic(err)
		return
	}

	rd := bytes.NewReader(s)

	size := int(core.ReadUInt16BE(rd))
	core.StartAllocAndReadBytes(size - 4, t.transport, func(b[] byte, err error) {
		t.recvData(b, err)
	})
}

/**
 * Receive data available for presentation layer
 * @param s {type.Stream}
*/
// TODO странно
func (t *TPKT) recvData(s []byte, err error) {
	core.StartAllocAndReadBytes(2, t.transport, t.recvHeader)
}

/**
 * Read extended fastpath header
 * @param s {type.Stream}
 */
func (t *TPKT) recvExtendedFastPathHeader(s []byte, length int,  err error) {
	rd := bytes.NewReader(s)

	rightPart := core.ReadUInt8(rd)
	leftPart := length & ^0x80 // TODO проверить length & ~0x80;
	packetSize := (leftPart << 8) + int(rightPart)
	core.StartAllocAndReadBytes(int(packetSize - 3), t.transport, t.recvFastPath)
}

/**
 * Read fast path data
 * @param s {type.Stream}
 */
func (t *TPKT) recvFastPath(s []byte,  err error) {
 // TODO this.emit('fastPathData', this.secFlag, s);
	core.StartAllocAndReadBytes(2, t.transport, t.recvHeader)
}

/**
 * Send message throught TPKT layer
 * @param message {type.*}
 */
func (t *TPKT) Write (data [] byte) {
	core.WriteUInt8(FASTPATH_ACTION_X224, t.transport)
	core.WriteUInt8(0, t.transport)
	core.WriteUInt16BE(uint16(len(data) + 4), t.transport)
	core.WriteBytes(data, t.transport)
}

/**
 * close stack
 */
func (t *TPKT) Close() {
	t.transport.Close()
}



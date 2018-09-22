package pdu

import (
	"../../core"
	"fmt"
	"github.com/chuckpreslar/emission"
)

/**
 * Global channel for all graphic updates
 * capabilities exchange and input handles
*/
type Global struct {
	emission.Emitter
	Transport core.Transport
	FastPathTransport core.Transport
	UserId uint16
	ShareId uint16
	ServerCapabilities map[CapsType]interface{}
	ClientCapabilities map[CapsType]interface{}
}
func NewGlobal(transport core.Transport, fastPathTransport core.Transport) *Global {
	return &Global{*emission.NewEmitter(), transport, fastPathTransport, 0, 0,
	make(map[CapsType]interface{}), make(map[CapsType]interface{}) }
}


func (global Global) SendPDU(message core.Component) error {
	return NewPDU(global.UserId, message).Write(global.Transport)
}

/**
 * Send formated Data PDU message
 * @param message {type.Component} PDU message
 */
func (global Global) SendDataPDU(message core.Component) error{
	return global.sendPDU(NewDataPDU(message, global.ShareId))
};

/**
 * Client side of Global channel automata
 * @param transport
*/
type Client struct {
	Global
	GccCore interface{}
	UserId uint16
	ChannelId uint16
}


func NewClient(transport core.Transport, fastPathTransport core.Transport) *Client {
	cl := &Client{*NewGlobal(transport, fastPathTransport), nil, 0, 0}

/* TODO connect
cl.transport.once('connect', function(core, UserId, ChannelId) {

self.connect(core, UserId, ChannelId);
}).on('close', function() {
self.emit('close');
}).on('error', function (err) {
self.emit('error', err);
});

if (cl.fastPathTransport) {
cl.fastPathTransport.on('fastPathData', function (secFlag, s) {
self.recvFastPath(secFlag, s);
});
}*/

	// init client capabilities
	cl.ClientCapabilities[CAPSTYPE_GENERAL] = NewGeneralCapability()
	cl.ClientCapabilities[CAPSTYPE_BITMAP] = NewBitmapCapability()
	cl.ClientCapabilities[CAPSTYPE_ORDER] = NewOrderCapability()
	cl.ClientCapabilities[CAPSTYPE_BITMAPCACHE] = NewBitmapCacheCapability()
	cl.ClientCapabilities[CAPSTYPE_POINTER] = NewPointerCapability()
	cl.ClientCapabilities[CAPSTYPE_INPUT] = NewInputCapability()
	cl.ClientCapabilities[CAPSTYPE_BRUSH] = NewBrushCapability()
	cl.ClientCapabilities[CAPSTYPE_GLYPHCACHE] = NewGlyphCapability()
	cl.ClientCapabilities[CAPSTYPE_OFFSCREENCACHE] = NewOffscreenBitmapCacheCapability()
	cl.ClientCapabilities[CAPSTYPE_VIRTUALCHANNEL] = NewVirtualChannelCapability()
	cl.ClientCapabilities[CAPSTYPE_SOUND] = NewSoundCapability()
	cl.ClientCapabilities[CAPSETTYPE_MULTIFRAGMENTUPDATE] = NewMultiFragmentUpdate()
}

func (cl *Client) Connect(GccCore interface{}, UserId uint16, ChannelId uint16) error {
	cl.GccCore = GccCore
	cl.UserId = UserId
	cl.ChannelId = ChannelId
	cl.RecvDemandActivePDU(cl.Transport)
	return nil
}

/**
 * close stack
 */
func (cl *Client) Close() error {
	return cl.Transport.Close()
};

/**
 * Receive capabilities from server
 * @param s {type.Stream}
 */
func (cl *Client) RecvDemandActivePDU (r core.Reader) error {
	pdu := NewEmptyPDU(0)
	err := pdu.Read(r)
	if err != nil {
		return err
	}

	if pdu.ShareControlHeader.PduType != PDUTYPE_DEMANDACTIVEPDU {
		fmt.Printf("ignore message type %d  during connection sequence",  pdu.ShareControlHeader.PduType)
		return cl.RecvDemandActivePDU(cl.Transport)
	}
	daPdu := pdu.Message.(DemandActivePDU)
	// store share id
	cl.ShareId = daPdu.ShareId;
	// store server capabilities
	for key, cap := range daPdu.CapabilitySets {
		cl.ServerCapabilities[key] = cap
	}
	generalCaps := cl.ServerCapabilities[CAPSTYPE_GENERAL].(GeneralCapability)

	cl.Transport.EnableSecureCheckSum = !!(generalCaps.extraFlags & ENC_SALTED_CHECKSUM)

	cl.SendConfirmActivePDU()
	cl.SendClientFinalizeSynchronizePDU()

	cl.RecvServerSynchronizePDU(cl.Transport)
}

/**
 * global channel automata state
 * @param s {type.Stream}
 */
func (cl *Client) RecvServerSynchronizePDU (r core.Reader) error {

	pdu := NewEmptyPDU(0, nil)
	err := pdu.Read(r)
	if err != nil {
		return err
	}
	if pdu.ShareControlHeader.PduType != PDUTYPE_DATAPDU ||
		pdu.Message.(DataPDU).ShareDataHeader.PduType2 != PDUTYPE2_SYNCHRONIZE {
		fmt.Printf("ignore message type %d during connection sequence" + int(pdu.ShareControlHeader.PduType))
	// loop on state
		return cl.RecvServerSynchronizePDU(cl.Transport)
	}

	return cl.RecvServerControlCooperatePDU(cl.Transport)
}

/**
 * global channel automata state
 * @param s {type.Stream}
 */
func (cl *Client) RecvServerControlCooperatePDU (r core.Reader) error {
	pdu := NewEmptyPDU()
	err := pdu.Read(r)
	if err != nil {
		return err
	}
	if pdu.ShareControlHeader.PduType != PDUTYPE_DATAPDU ||
		pdu.Message.(DataPDU).ShareDataHeader.PduType2 != PDUTYPE2_CONTROL ||
		pdu.Message.(DataPDU).Data.(ControlDataPDU).Action != CTRLACTION_COOPERATE {
		fmt.Printf("ignore message type %d during connection sequence", pdu.ShareControlHeader.PduType);
		// loop on state
		return cl.RecvServerControlCooperatePDU(cl.Transport)
	}

	return cl.RecvServerControlGrantedPDU(cl.Transport)
}

/**
 * global channel automata state
 * @param s {type.Stream}
 */
func (cl *Client) RecvServerControlGrantedPDU(r core.Reader) error {
	pdu := NewEmptyPDU()
	err := pdu.Read(r)
	if err != nil {
		return err
	}
	if pdu.ShareControlHeader.PduType != PDUTYPE_DATAPDU ||
		pdu.Message.(DataPDU).ShareDataHeader.PduType2 != PDUTYPE2_CONTROL ||
		pdu.Message.(DataPDU).Data.(ControlDataPDU).Action != CTRLACTION_GRANTED_CONTROL {
		fmt.Printf("ignore message type %d during connection sequence", pdu.ShareControlHeader.PduType);
		// loop on state
		return cl.RecvServerControlGrantedPDU(cl.Transport)
	}

	return cl.RecvServerFontMapPDU(cl.Transport)
}

/**
 * global channel automata state
 * @param s {type.Stream}
 */
func (cl *Client) RecvServerFontMapPDU(r core.Reader) error {
	pdu := NewEmptyPDU()
	err := pdu.Read(r)
	if err != nil {
		return err
	}
	if pdu.ShareControlHeader.PduType != PDUTYPE_DATAPDU ||
		pdu.Message.(DataPDU).ShareDataHeader.PduType2 != PDUTYPE2_CONTROL ||
		pdu.Message.(DataPDU).Data.(ControlDataPDU).Action != CTRLACTION_GRANTED_CONTROL {
		fmt.Printf("ignore message type %d during connection sequence", pdu.ShareControlHeader.PduType);
		// loop on state
		return cl.RecvServerFontMapPDU(cl.Transport)
	}

	//TODO cl.emit('connect');
	return cl.RecvPDU(cl.Transport)
}

/**
 * Main reveive fast path
 * @param secFlag {integer}
 * @param s {type.Stream}
 */
func (cl *Client) RecvFastPath(secFlag int, r core.Reader) {

	pdur.
	while (s.availableLength() > 0) {
var pdu = fastPathUpdatePDU().read(s);
switch (pdu.updateHeader & 0xf) {
case FastPathUpdateType.FASTPATH_UPDATETYPE_BITMAP:
cl.emit('bitmap', pdu.updateData.rectangles.obj);
break;
default:
}
}
};

/**
 * global channel automata state
 * @param s {type.Stream}
 */
func (cl *Client) RecvPDU(r core.Reader) error {
while (s.availableLength() > 0) {
var pdu = pdu().read(s);
switch(pdu.shareControlHeader.pduType) {
case PDUType.PDUTYPE_DEACTIVATEALLPDU:
var self = this;
cl.transport.removeAllListeners('data');
cl.transport.once('data', function(s) {
self.recvDemandActivePDU(s);
});
break;
case PDUType.PDUTYPE_DATAPDU:
cl.readDataPDU(pdu.pduMessage)
break;
default:
log.debug('ignore pdu type ' + pdu.shareControlHeader.pduType);
}
}
};

/**
 * main receive for data PDU packet
 * @param dataPDU {dataPDU}
 */
func (cl *Client) ReadDataPDU = function (dataPDU) {
switch(dataPDU.shareDataHeader.pduType2) {
case PDUTYPE2_SET_ERROR_INFO_PDU:
break;
case PDUTYPE2_SHUTDOWN_DENIED:
cl.transport.close();
break;
case PDUTYPE2_SAVE_SESSION_INFO:
cl.emit('session');
break;
case PDUTYPE2_UPDATE:
cl.readUpdateDataPDU(dataPDU.pduData)
break;
}
};

/**
 * Main upadate pdu receive function
 * @param updateDataPDU
 */
func (cl *Client) ReadUpdateDataPDU (updateDataPDU) {
switch(updateDataPDU.updateType) {
case UpdateType.UPDATETYPE_BITMAP:
cl.emit('bitmap', updateDataPDU.updateData.rectangles.obj)
break;
}
};

/**
 * send all client capabilities
 */
func (cl *Client) SendConfirmActivePDU () error {
	generalCapability := cl.ClientCapabilities[CAPSTYPE_GENERAL].(GeneralCapability)
	generalCapability.osMajorType = OSMAJORTYPE_WINDOWS;
	generalCapability.osMinorType = OSMINORTYPE_WINDOWS_NT;
	generalCapability.extraFlags = LONG_CREDENTIALS_SUPPORTED
	| 	caps.GeneralExtraFlag.NO_BITMAP_COMPRESSION_HDR
	| 	caps.GeneralExtraFlag.ENC_SALTED_CHECKSUM
	|	caps.GeneralExtraFlag.FASTPATH_OUTPUT_SUPPORTED;

var bitmapCapability = cl.ClientCapabilities[caps.CapsType.CAPSTYPE_BITMAP].obj;
bitmapCapability.preferredBitsPerPixel = cl.GccCore.highColorDepth;
bitmapCapability.desktopWidth = cl.GccCore.desktopWidth;
bitmapCapability.desktopHeight = cl.GccCore.desktopHeight;

var orderCapability = cl.ClientCapabilities[caps.CapsType.CAPSTYPE_ORDER].obj;
orderCapability.orderFlags |= caps.OrderFlag.ZEROBOUNDSDELTASSUPPORT;

var inputCapability = cl.ClientCapabilities[caps.CapsType.CAPSTYPE_INPUT].obj;
inputCapability.inputFlags = caps.InputFlags.INPUT_FLAG_SCANCODES | caps.InputFlags.INPUT_FLAG_MOUSEX | caps.InputFlags.INPUT_FLAG_UNICODE;
inputCapability.keyboardLayout = cl.GccCore.kbdLayout;
inputCapability.keyboardType = cl.GccCore.keyboardType;
inputCapability.keyboardSubType = cl.GccCore.keyboardSubType;
inputCapability.keyboardrFunctionKey = cl.GccCore.keyboardFnKeys;
inputCapability.imeFileName = cl.GccCore.imeFileName;

var capabilities = new type.Component([]);
for(var i in cl.ClientCapabilities) {
capabilities.push(caps.capability(cl.ClientCapabilities[i]));
}

var confirmActivePDU = confirmActivePDU(capabilities, cl.shareId);

cl.sendPDU(confirmActivePDU);
};

/**
 * send synchronize PDU
*/
func (cl *Client) SendClientFinalizeSynchronizePDU () error{
	err := cl.sendDataPDU(synchronizeDataPDU(cl.ChannelId))
	if err != nil {
		return err
	}
	err = cl.sendDataPDU(controlDataPDU(Action.CTRLACTION_COOPERATE))
	if err != nil {
		return err
	}
	err = cl.sendDataPDU(controlDataPDU(Action.CTRLACTION_REQUEST_CONTROL))
	if err != nil {
		return err
	}
	err =cl.sendDataPDU(fontListDataPDU())
	if err != nil {
		return err
	}
	return err
}

/**
 * Send input event as slow path input
 * @param inputEvents {array}
 */
func (cl *Client) SendInputEvents(inputEvents) {
	var pdu = NewClientInputEventPDU(new type.Component(inputEvents.map(function (e) {
return slowPathInputEvent(e);
})));

cl.sendDataPDU(pdu);
}

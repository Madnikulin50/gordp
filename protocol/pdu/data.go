package pdu

// import "log"

type PDUType byte
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240576.aspx
 */
const /*PDUType*/ (
	PDUTYPE_DEMANDACTIVEPDU  PDUType = 0x11
	PDUTYPE_CONFIRMACTIVEPDU = 0x13
	PDUTYPE_DEACTIVATEALLPDU = 0x16
	PDUTYPE_DATAPDU = 0x17
	PDUTYPE_SERVER_REDIR_PKT = 0x1A
)


type PDUType2 byte

/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240577.aspx
 */
const /*PDUType2*/ (
	PDUTYPE2_UPDATE PDUType2 = 0x02
	PDUTYPE2_CONTROL = 0x14
	PDUTYPE2_POINTER = 0x1B
	PDUTYPE2_INPUT = 0x1C
	PDUTYPE2_SYNCHRONIZE = 0x1F
	PDUTYPE2_REFRESH_RECT = 0x21
	PDUTYPE2_PLAY_SOUND = 0x22
	PDUTYPE2_SUPPRESS_OUTPUT = 0x23
	PDUTYPE2_SHUTDOWN_REQUEST = 0x24
	PDUTYPE2_SHUTDOWN_DENIED = 0x25
	PDUTYPE2_SAVE_SESSION_INFO = 0x26
	PDUTYPE2_FONTLIST = 0x27
	PDUTYPE2_FONTMAP = 0x28
	PDUTYPE2_SET_KEYBOARD_INDICATORS = 0x29
	PDUTYPE2_BITMAPCACHE_PERSISTENT_LIST = 0x2B
	PDUTYPE2_BITMAPCACHE_ERROR_PDU = 0x2C
	PDUTYPE2_SET_KEYBOARD_IME_STATUS = 0x2D
	PDUTYPE2_OFFSCRCACHE_ERROR_PDU = 0x2E
	PDUTYPE2_SET_ERROR_INFO_PDU = 0x2F
	PDUTYPE2_DRAWNINEGRID_ERROR_PDU = 0x30
	PDUTYPE2_DRAWGDIPLUS_ERROR_PDU = 0x31
	PDUTYPE2_ARC_STATUS_PDU = 0x32
	PDUTYPE2_STATUS_INFO_PDU = 0x36
	PDUTYPE2_MONITOR_LAYOUT_PDU = 0x37
)

type StreamId int8
/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240577.aspx
 */
const /*StreamId*/ (
STREAM_UNDEFINED StreamId = 0x00
STREAM_LOW = 0x01
STREAM_MED = 0x02
STREAM_HI = 0x04
)

type CompressionOrder byte

/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240577.aspx
 */
const /*CompressionOrder*/ (
CompressionTypeMask CompressionOrder = 0x0F
PACKET_COMPRESSED = 0x20
PACKET_AT_FRONT = 0x40
PACKET_FLUSHED = 0x80
)

type CompressionType byte

/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240577.aspx
 */
const /*CompressionType*/ (
PACKET_COMPR_TYPE_8K CompressionType/*?*/ = 0x0
PACKET_COMPR_TYPE_64K = 0x1
PACKET_COMPR_TYPE_RDP6 = 0x2
PACKET_COMPR_TYPE_RDP61 = 0x3
)

type Action uint16
/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240492.aspx
 */
const /*Action*/ (
CTRLACTION_REQUEST_CONTROL Action = 0x0001
CTRLACTION_GRANTED_CONTROL = 0x0002
CTRLACTION_DETACH = 0x0003
CTRLACTION_COOPERATE = 0x0004
)

type PersistentKeyListFlag byte
/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240495.aspx
 */
const /*PersistentKeyListFlag*/ (
PERSIST_FIRST_PDU PersistentKeyListFlag = 0x01
PERSIST_LAST_PDU = 0x02
)

type BitmapFlag uint16
/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240612.aspx
 */
const /*BitmapFlag*/ (
BF_BITMAP_COMPRESSION BitmapFlag = 0x0001
BF_NO_BITMAP_COMPRESSION_HDR = 0x0400
)

type UpdateType uint16

/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240608.aspx
 */
const /*UpdateType*/ (
UPDATETYPE_ORDERS UpdateType = 0x0000
UPDATETYPE_BITMAP = 0x0001
UPDATETYPE_PALETTE = 0x0002
UPDATETYPE_SYNCHRONIZE = 0x0003
)

type InputMessageType uint16
/**
 * @see http=//msdn.microsoft.com/en-us/library/cc240583.aspx
 */
const /*InputMessageType*/ (
INPUT_EVENT_SYNC InputMessageType = 0x0000
INPUT_EVENT_UNUSED = 0x0002
INPUT_EVENT_SCANCODE = 0x0004
INPUT_EVENT_UNICODE = 0x0005
INPUT_EVENT_MOUSE = 0x8001
INPUT_EVENT_MOUSEX = 0x8002
)

type PointerFlag uint16

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240586.aspx
 */
const /*PointerFlag*/ (
PTRFLAGS_HWHEEL PointerFlag = 0x0400
PTRFLAGS_WHEEL= 0x0200
PTRFLAGS_WHEEL_NEGATIVE= 0x0100
WheelRotationMask= 0x01FF
PTRFLAGS_MOVE= 0x0800
PTRFLAGS_DOWN= 0x8000
PTRFLAGS_BUTTON1= 0x1000
PTRFLAGS_BUTTON2= 0x2000
PTRFLAGS_BUTTON3= 0x4000
)

type KeyboardFlag uint16
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240584.aspx
 */
const /*KeyboardFlag*/ (
KBDFLAGS_EXTENDED KeyboardFlag = 0x0100
KBDFLAGS_DOWN= 0x4000
KBDFLAGS_RELEASE= 0x8000
)

type FastPathUpdateType byte

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240622.aspx
 */
const /*FastPathUpdateType*/ (
FASTPATH_UPDATETYPE_ORDERS FastPathUpdateType = 0x0
FASTPATH_UPDATETYPE_BITMAP= 0x1
FASTPATH_UPDATETYPE_PALETTE= 0x2
FASTPATH_UPDATETYPE_SYNCHRONIZE= 0x3
FASTPATH_UPDATETYPE_SURFCMDS= 0x4
FASTPATH_UPDATETYPE_PTR_NULL= 0x5
FASTPATH_UPDATETYPE_PTR_DEFAULT= 0x6
FASTPATH_UPDATETYPE_PTR_POSITION= 0x8
FASTPATH_UPDATETYPE_COLOR= 0x9
FASTPATH_UPDATETYPE_CACHED= 0xA
FASTPATH_UPDATETYPE_POINTER= 0xB
)

type FastPathOutputCompression uint8
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240622.aspx
 */
const /*FastPathOutputCompression*/ (
FASTPATH_OUTPUT_COMPRESSION_USED FastPathOutputCompression = 0x2
)

type Display uint8
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240648.aspx
 */
const /*Display*/ (
SUPPRESS_DISPLAY_UPDATES Display = 0x00
ALLOW_DISPLAY_UPDATES= 0x01
)

type ToogleFlag uint32
/**
 * @see https://msdn.microsoft.com/en-us/library/cc240588.aspx
 */
const /*ToogleFlag*/ (
TS_SYNC_SCROLL_LOCK ToogleFlag = 0x00000001
TS_SYNC_NUM_LOCK= 0x00000002
TS_SYNC_CAPS_LOCK= 0x00000004
TS_SYNC_KANA_LOCK= 0x00000008
)

type ErrorInfo uint32
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240544.aspx
 */
const /*ErrorInfo*/ (
ERRINFO_RPC_INITIATED_DISCONNECT ErrorInfo = 0x00000001
ERRINFO_RPC_INITIATED_LOGOFF= 0x00000002
ERRINFO_IDLE_TIMEOUT= 0x00000003
ERRINFO_LOGON_TIMEOUT= 0x00000004
ERRINFO_DISCONNECTED_BY_OTHERCONNECTION= 0x00000005
ERRINFO_OUT_OF_MEMORY= 0x00000006
ERRINFO_SERVER_DENIED_CONNECTION= 0x00000007
ERRINFO_SERVER_INSUFFICIENT_PRIVILEGES= 0x00000009
ERRINFO_SERVER_FRESH_CREDENTIALS_REQUIRED= 0x0000000A
ERRINFO_RPC_INITIATED_DISCONNECT_BYUSER= 0x0000000B
ERRINFO_LOGOFF_BY_USER= 0x0000000C
ERRINFO_LICENSE_INTERNAL= 0x00000100
ERRINFO_LICENSE_NO_LICENSE_SERVER= 0x00000101
ERRINFO_LICENSE_NO_LICENSE= 0x00000102
ERRINFO_LICENSE_BAD_CLIENT_MSG= 0x00000103
ERRINFO_LICENSE_HWID_DOESNT_MATCH_LICENSE= 0x00000104
ERRINFO_LICENSE_BAD_CLIENT_LICENSE= 0x00000105
ERRINFO_LICENSE_CANT_FINISH_PROTOCOL= 0x00000106
ERRINFO_LICENSE_CLIENT_ENDED_PROTOCOL= 0x00000107
ERRINFO_LICENSE_BAD_CLIENT_ENCRYPTION= 0x00000108
ERRINFO_LICENSE_CANT_UPGRADE_LICENSE= 0x00000109
ERRINFO_LICENSE_NO_REMOTE_CONNECTIONS= 0x0000010A
ERRINFO_CB_DESTINATION_NOT_FOUND= 0x0000400
ERRINFO_CB_LOADING_DESTINATION= 0x0000402
ERRINFO_CB_REDIRECTING_TO_DESTINATION= 0x0000404
ERRINFO_CB_SESSION_ONLINE_VM_WAKE= 0x0000405
ERRINFO_CB_SESSION_ONLINE_VM_BOOT= 0x0000406
ERRINFO_CB_SESSION_ONLINE_VM_NO_DNS= 0x0000407
ERRINFO_CB_DESTINATION_POOL_NOT_FREE= 0x0000408
ERRINFO_CB_CONNECTION_CANCELLED= 0x0000409
ERRINFO_CB_CONNECTION_ERROR_INVALID_SETTINGS= 0x0000410
ERRINFO_CB_SESSION_ONLINE_VM_BOOT_TIMEOUT= 0x0000411
ERRINFO_CB_SESSION_ONLINE_VM_SESSMON_FAILED= 0x0000412
ERRINFO_UNKNOWNPDUTYPE2= 0x000010C9
ERRINFO_UNKNOWNPDUTYPE= 0x000010CA
ERRINFO_DATAPDUSEQUENCE= 0x000010CB
ERRINFO_CONTROLPDUSEQUENCE= 0x000010CD
ERRINFO_INVALIDCONTROLPDUACTION= 0x000010CE
ERRINFO_INVALIDINPUTPDUTYPE= 0x000010CF
ERRINFO_INVALIDINPUTPDUMOUSE= 0x000010D0
ERRINFO_INVALIDREFRESHRECTPDU= 0x000010D1
ERRINFO_CREATEUSERDATAFAILED= 0x000010D2
ERRINFO_CONNECTFAILED= 0x000010D3
ERRINFO_CONFIRMACTIVEWRONGSHAREID= 0x000010D4
ERRINFO_CONFIRMACTIVEWRONGORIGINATOR= 0x000010D5
ERRINFO_PERSISTENTKEYPDUBADLENGTH= 0x000010DA
ERRINFO_PERSISTENTKEYPDUILLEGALFIRST= 0x000010DB
ERRINFO_PERSISTENTKEYPDUTOOMANYTOTALKEYS= 0x000010DC
ERRINFO_PERSISTENTKEYPDUTOOMANYCACHEKEYS= 0x000010DD
ERRINFO_INPUTPDUBADLENGTH= 0x000010DE
ERRINFO_BITMAPCACHEERRORPDUBADLENGTH= 0x000010DF
ERRINFO_SECURITYDATATOOSHORT= 0x000010E0
ERRINFO_VCHANNELDATATOOSHORT= 0x000010E1
ERRINFO_SHAREDATATOOSHORT= 0x000010E2
ERRINFO_BADSUPRESSOUTPUTPDU= 0x000010E3
ERRINFO_CONFIRMACTIVEPDUTOOSHORT= 0x000010E5
ERRINFO_CAPABILITYSETTOOSMALL= 0x000010E7
ERRINFO_CAPABILITYSETTOOLARGE= 0x000010E8
ERRINFO_NOCURSORCACHE= 0x000010E9
ERRINFO_BADCAPABILITIES= 0x000010EA
ERRINFO_VIRTUALCHANNELDECOMPRESSIONERR= 0x000010EC
ERRINFO_INVALIDVCCOMPRESSIONTYPE= 0x000010ED
ERRINFO_INVALIDCHANNELID= 0x000010EF
ERRINFO_VCHANNELSTOOMANY= 0x000010F0
ERRINFO_REMOTEAPPSNOTENABLED= 0x000010F3
ERRINFO_CACHECAPNOTSET= 0x000010F4
ERRINFO_BITMAPCACHEERRORPDUBADLENGTH2= 0x000010F5
ERRINFO_OFFSCRCACHEERRORPDUBADLENGTH= 0x000010F6
ERRINFO_DNGCACHEERRORPDUBADLENGTH= 0x000010F7
ERRINFO_GDIPLUSPDUBADLENGTH= 0x000010F8
ERRINFO_SECURITYDATATOOSHORT2= 0x00001111
ERRINFO_SECURITYDATATOOSHORT3= 0x00001112
ERRINFO_SECURITYDATATOOSHORT4= 0x00001113
ERRINFO_SECURITYDATATOOSHORT5= 0x00001114
ERRINFO_SECURITYDATATOOSHORT6= 0x00001115
ERRINFO_SECURITYDATATOOSHORT7= 0x00001116
ERRINFO_SECURITYDATATOOSHORT8= 0x00001117
ERRINFO_SECURITYDATATOOSHORT9= 0x00001118
ERRINFO_SECURITYDATATOOSHORT10= 0x00001119
ERRINFO_SECURITYDATATOOSHORT11= 0x0000111A
ERRINFO_SECURITYDATATOOSHORT12= 0x0000111B
ERRINFO_SECURITYDATATOOSHORT13= 0x0000111C
ERRINFO_SECURITYDATATOOSHORT14= 0x0000111D
ERRINFO_SECURITYDATATOOSHORT15= 0x0000111E
ERRINFO_SECURITYDATATOOSHORT16= 0x0000111F
ERRINFO_SECURITYDATATOOSHORT17= 0x00001120
ERRINFO_SECURITYDATATOOSHORT18= 0x00001121
ERRINFO_SECURITYDATATOOSHORT19= 0x00001122
ERRINFO_SECURITYDATATOOSHORT20= 0x00001123
ERRINFO_SECURITYDATATOOSHORT21= 0x00001124
ERRINFO_SECURITYDATATOOSHORT22= 0x00001125
ERRINFO_SECURITYDATATOOSHORT23= 0x00001126
ERRINFO_BADMONITORDATA= 0x00001129
ERRINFO_VCDECOMPRESSEDREASSEMBLEFAILED= 0x0000112A
ERRINFO_VCDATATOOLONG= 0x0000112B
ERRINFO_BAD_FRAME_ACK_DATA= 0x0000112C
ERRINFO_GRAPHICSMODENOTSUPPORTED= 0x0000112D
ERRINFO_GRAPHICSSUBSYSTEMRESETFAILED= 0x0000112E
ERRINFO_GRAPHICSSUBSYSTEMFAILED= 0x0000112F
ERRINFO_TIMEZONEKEYNAMELENGTHTOOSHORT= 0x00001130
ERRINFO_TIMEZONEKEYNAMELENGTHTOOLONG= 0x00001131
ERRINFO_DYNAMICDSTDISABLEDFIELDMISSING= 0x00001132
ERRINFO_VCDECODINGERROR= 0x00001133
ERRINFO_UPDATESESSIONKEYFAILED= 0x00001191
ERRINFO_DECRYPTFAILED= 0x00001192
ERRINFO_ENCRYPTFAILED= 0x00001193
ERRINFO_ENCPKGMISMATCH= 0x00001194
ERRINFO_DECRYPTFAILED2= 0x00001195
)


/**
 * @see http://msdn.microsoft.com/en-us/library/cc240576.aspx
 * @param length {integer} length of entire pdu packet
 * @param pduType {PDUType.*} type of pdu packet
 * @param userId {integer}
 * @param opt {object} type option
 * @returns {type.Component}
 */
 type  ShareControlHeader struct {
	 Component
	 totalLength uint16
	 pduType uint16
	 // for xp sp3 and deactiveallpdu PDUSource may not be present
	 PDUSource uint16 "optional"
 }
func NewShareControlHeader(length uint16, pduType uint16, userId int16, opt interface{}) * ShareControlHeader {
	return *ShareControlHeader{ *NewComponent(opt), length, pduType, userId}
}

func (c *ShareControlHeader) Write(writer *Writer) {
	panic("Нереализовано")
}

func (c *ShareControlHeader) Read(reader *Reader) {
	panic("Нереализовано")
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240577.aspx
 * @param length {integer} lezngth of entire packet
 * @param pduType2 {PDUType2.*} sub PDU type
 * @param shareId {integer} global layer id
 * @param opt {object} type option
 * @returns {type.Component}
 */

type  SharedDataHeader struct {
	Component
	shareId uint32
	pad1 uint8
	streamId uint8
	uncompressedLength uint16
	pduType2 uint8
	compressedType uint8
	compressedLength uint16
}

func NewSharedDataHeader(length uint16, pduType2 uint8, shareId uint32, opt interface{}) *SharedDataHeader {
	return &SharedDataHeader{ *NewComponent(opt), shareId, 0, StreamId.STREAM_LOW,length -8,
		pduType2, 0, 0}
}

func (c *SharedDataHeader) Write(writer *Writer) {
	panic("Нереализовано")
}

func (c *SharedDataHeader) Read(reader *Reader) {
	panic("Нереализовано")
}



/**
 * @see http://msdn.microsoft.com/en-us/library/cc240485.aspx
 * @param capabilities {type.Component} capabilities array
 * @param opt {object} type option
 * @returns {type.Component}
 */

type DemandActivePDU struct {
	Component
	__PDUTYPE__  PDUType
	shareId uint32
	lengthSourceDescriptor uint16
	lengthCombinedCapabilities uint16
	sourceDescriptor []byte
	numberCapabilities uint16
	pad2Octets uint16
	capabilitySets interface{}
	sessionId uint32
}

func NewDemandActivePDU(capabilities interface{}, opt  interface{}) *DemandActivePDU {
	/*
	var self {
__PDUTYPE__ : PDUType.PDUTYPE_DEMANDACTIVEPDU,
shareId uint32
lengthSourceDescriptor uint16func() {
return self.sourceDescriptor.size();
}),
lengthCombinedCapabilities uint16func() {
return self.numberCapabilities.size() + self.pad2Octets.size() + self.capabilitySets.size();
}),
sourceDescriptor : new type.BinaryString(new Buffer('node-rdpjs', 'binary'), { readLength : new type.CallableValue(func() {
return self.lengthSourceDescriptor.value
}) }),
numberCapabilities uint16func() {
return self.capabilitySets.obj.length;
}),
pad2Octets uint16
capabilitySets : capabilities || new type.Factory(func(s) {
self.capabilitySets = new type.Component([]);
for(var i = 0; i < self.numberCapabilities.value; i++) {
self.capabilitySets.obj.push(caps.capability().read(s))
}
}),
sessionId uint32
};

return new type.Component(self, opt);
	 */
	sd := []byte("node-rdpjs")
	return &DemandActivePDU{ *NewComponent(opt), PDUTYPE_DEMANDACTIVEPDU, 0,
	len(sd), len(uint16) * 2 /*+ self.capabilitySets.size()*/,
	sd, 0, 0, nil}
}

func (c *DemandActivePDU) Write(writer *Writer) {
	panic("Нереализовано")
}

func (c *DemandActivePDU) Read(reader *Reader) {
	panic("Нереализовано")
}

func demandActivePDU(capabilities, opt) {

}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240488.aspx
 * @param capabilities {type.Component} capabilities array
 * @param shareId {integer} session id
 * @param opt {object} type option
 * @returns {type.Component}
 */
func confirmActivePDU(capabilities, shareId, opt) {
var self {
__PDUTYPE__ : PDUType.PDUTYPE_CONFIRMACTIVEPDU,
shareId : new type.UInt32Le(shareId),
originatorId uint160x03EA, { constant : true }),
lengthSourceDescriptor uint16func() {
return self.sourceDescriptor.size();
}),
lengthCombinedCapabilities uint16func() {
return self.numberCapabilities.size() + self.pad2Octets.size() + self.capabilitySets.size();
}),
sourceDescriptor : new type.BinaryString(new Buffer('rdpy', 'binary'), { readLength : new type.CallableValue(func() {
return self.lengthSourceDescriptor.value
}) }),
numberCapabilities uint16func() {
return self.capabilitySets.obj.length;
}),
pad2Octets uint16
capabilitySets : capabilities || new type.Factory(func(s) {
self.capabilitySets = new type.Component([]);
for(var i = 0; i < self.numberCapabilities.value; i++) {
self.capabilitySets.obj.push(caps.capability().read(s))
}
})
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240536.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func deactiveAllPDU(opt) {
var self = {
__PDUTYPE__ : PDUType.PDUTYPE_DEACTIVATEALLPDU,
shareId uint32
lengthSourceDescriptor uint16func() {
return self.sourceDescriptor.size();
}),
sourceDescriptor : new type.BinaryString(new Buffer('rdpy', 'binary'), { readLength : new type.CallableValue(func() {
self.lengthSourceDescriptor
}) })
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240490.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func synchronizeDataPDU(targetUser, opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_SYNCHRONIZE,
messageType uint161, { constant : true }),
targetUser uint16targetUser)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240492.aspx
 * @param action {integer}
 * @param opt  {object} type option
 * @returns {type.Component}
 */
func controlDataPDU(action, opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_CONTROL,
action uint16action),
grantId uint16
controlId uint32
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240544.aspx
 * @param errorInfo {integer}
 * @param opt {object} type option
 * @returns {type.Component}
 */
func errorInfoDataPDU(errorInfo, opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_SET_ERROR_INFO_PDU,
errorInfo : new type.UInt32Le(errorInfo)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240498.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func fontListDataPDU(opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_FONTLIST,
numberFonts uint16
totalNumFonts uint16
listFlags uint160x0003),
entrySize uint160x0032)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240498.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func fontMapDataPDU(opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_FONTMAP,
numberEntries uint16
totalNumEntries uint16
mapFlags uint160x0003),
entrySize uint160x0004)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240496.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func persistentListEntry(opt) {
var self = {
key1 uint32
key2 uint32
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240495.aspx
 * @param entries {type.Component}
 * @param opt {object} type option
 * @returns {type.Component}
 */
func persistentListPDU(entries, opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_BITMAPCACHE_PERSISTENT_LIST,
numEntriesCache0 uint16
numEntriesCache1 uint16
numEntriesCache2 uint16
numEntriesCache3 uint16
numEntriesCache4 uint16
totalEntriesCache0 uint16
totalEntriesCache1 uint16
totalEntriesCache2 uint16
totalEntriesCache3 uint16
totalEntriesCache4 uint16
bitMask : new type.UInt8(),
pad2 : new type.UInt8(),
pad3 uint16
entries : entries || new type.Factory(func(s) {
var numEntries = self.numEntriesCache0.value + self.numEntriesCache1.value + self.numEntriesCache2.value + self.numEntriesCache3.value + self.numEntriesCache4.value;
self.entries = new type.Component([]);
for(var i = 0; i < numEntries; i++) {
self.entries.obj.push(persistentListEntry().read(s));
}
})
};

return new type.Component(self, opt);
}

/**
 * @see https://msdn.microsoft.com/en-us/library/cc240588.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func synchronizeEvent(opt) {
var self = {
__INPUT_MESSAGE_TYPE__ : InputMessageType.INPUT_EVENT_SYNC,
pad2Octets uint16
toggleFlags uint32
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240586.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func pointerEvent(opt) {
var self = {
__INPUT_MESSAGE_TYPE__ : InputMessageType.INPUT_EVENT_MOUSE,
pointerFlags uint16
xPos uint16
yPos uint16)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240584.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func scancodeKeyEvent(opt) {
var self = {
__INPUT_MESSAGE_TYPE__ : InputMessageType.INPUT_EVENT_SCANCODE,
keyboardFlags uint16
keyCode uint16
pad2Octets uint16)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240585.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func unicodeKeyEvent(opt) {
var self = {
__INPUT_MESSAGE_TYPE__ : InputMessageType.INPUT_EVENT_UNICODE,
keyboardFlags uint16
unicode uint16
pad2Octets uint16)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240583.aspx
 * @param slowPathInputData {type.Component} message generate for slow path input event
 * @param opt {object} type option
 * @returns {type.Component}
 */
func slowPathInputEvent(slowPathInputData, opt) {
var self = {
eventTime uint32
messageType uint16func() {
return self.slowPathInputData.obj.__INPUT_MESSAGE_TYPE__;
}),
slowPathInputData : slowPathInputData || new type.Factory(func(s) {
switch(self.messageType.value) {
case InputMessageType.INPUT_EVENT_SYNC:
self.slowPathInputData = synchronizeEvent().read(s);
break;
case InputMessageType.INPUT_EVENT_MOUSE:
self.slowPathInputData = pointerEvent().read(s);
break;
case InputMessageType.INPUT_EVENT_SCANCODE:
self.slowPathInputData = scancodeKeyEvent().read(s);
break;
case InputMessageType.INPUT_EVENT_UNICODE:
self.slowPathInputData = unicodeKeyEvent().read(s);
break;
default:
log.error('unknown slowPathInputEvent ' + self.messageType.value);
}
})
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc746160.aspx
 * @param inputs {type.Component} list of inputs
 * @param opt {object} type option
 * @returns {type.Component}
 */
func clientInputEventPDU(inputs, opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_INPUT,
numEvents uint16func() {
return self.slowPathInputEvents.obj.length;
}),
pad2Octets uint16
slowPathInputEvents : inputs || new type.Factory(func(s) {
self.slowPathInputEvents = new type.Component([]);
for(var i = 0; i < self.numEvents.value; i++) {
self.slowPathInputEvents.obj.push(slowPathInputEvent().read(s));
}
})
};

return new type.Component(self, opt);
}

/**
 * @param opt {object} type option
 * @returns {type.Component}
 */
func shutdownRequestPDU(opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_SHUTDOWN_REQUEST
};

return new type.Component(self, opt);
}

/**
 * @param opt {object} type option
 * @returns {type.Component}
 */
func shutdownDeniedPDU(opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_SHUTDOWN_DENIED
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240643.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func inclusiveRectangle(opt) {
var self = {
left uint16
top uint16
right uint16
bottom uint16)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240648.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func supressOutputDataPDU(opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_SUPPRESS_OUTPUT,
allowDisplayUpdates : new type.UInt8(),
pad3Octets : new type.Component([new type.UInt8(), new type.UInt8(), new type.UInt8()]),
desktopRect : inclusiveRectangle({ conditional : func() {
return self.allowDisplayUpdates.value === Display.ALLOW_DISPLAY_UPDATES;
} })
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240646.aspx
 * @param rectangles {type.Component} list of inclusive rectangles
 * @param opt {object} type option
 * @returns {type.Component}
 */
func refreshRectPDU(rectangles, opt) {
var self = {
__PDUTYPE2__ : PDUType2.PDUTYPE2_REFRESH_RECT,
numberOfAreas : UInt8(func() {
return self.areasToRefresh.obj.length;
}),
pad3Octets : new type.Component([new type.UInt8(), new type.UInt8(), new type.UInt8()]),
areasToRefresh : rectangles || new type.Factory(func(s) {
self.areasToRefresh = new type.Component([]);
for(var i = 0; i < self.numberOfAreas.value; i++) {
self.areasToRefresh.obj.push(inclusiveRectangle().read(s));
}
})
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240644.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func bitmapCompressedDataHeader(opt) {
var self = {
cbCompFirstRowSize uint160x0000, { constant : true }),
// compressed data size
cbCompMainBodySize uint16
cbScanWidth uint16
// uncompressed data size
cbUncompressedSize uint16)
};

return new type.Component(self, opt);
}

/**
 * @see
 * @param coord {object}
 * 		.destLeft {integer}
 * 		.destTop {integer}
 * 		.destRight {integer}
 * 		.destBottom {integer}
 * 		.width {integer}
 * 		.height {integer}
 * 		.bitsPerPixel {integer}
 * 		.data {Buffer}
 * @param opt {object} type option
 * @returns {type.Component}
 */
func bitmapData(coord, opt) {
coord = coord || {};
var self = {
destLeft uint16coord.destLeft),
destTop uint16coord.destTop),
destRight uint16coord.destRight),
destBottom uint16coord.destBottom),
width uint16coord.width),
height uint16coord.height),
bitsPerPixel uint16coord.bitsPerPixel),
flags uint16
bitmapLength uint16func() {
return self.bitmapComprHdr.size() + self.bitmapDataStream.size();
}),
bitmapComprHdr : bitmapCompressedDataHeader( { conditional : func() {
return (self.flags.value & BitmapFlag.BITMAP_COMPRESSION) && !(self.flags.value & BitmapFlag.NO_BITMAP_COMPRESSION_HDR);
} }),
bitmapDataStream : new type.BinaryString(coord.data, { readLength : new type.CallableValue(func() {
if(!self.flags.value & BitmapFlag.BITMAP_COMPRESSION || (self.flags.value & BitmapFlag.NO_BITMAP_COMPRESSION_HDR)) {
return self.bitmapLength.value;
}
else {
return self.bitmapComprHdr.cbCompMainBodySize.value;
}
}) })
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/dd306368.aspx
 * @param data {type.Component} list of bitmap data type
 * @param opt {object} type option
 * @returns {type.Component}
 */
func bitmapUpdateDataPDU(data, opt) {
var self = {
__UPDATE_TYPE__ : UpdateType.UPDATETYPE_BITMAP,
numberRectangles uint16func() {
return self.rectangles.obj.length;
}),
rectangles : data || new type.Factory(func(s) {
self.rectangles = new type.Component([]);
for(var i = 0; i < self.numberRectangles.value; i++) {
self.rectangles.obj.push(bitmapData().read(s));
}
})
};

return new type.Component(self, opt);
}

/**
 * @see https://msdn.microsoft.com/en-us/library/cc240613.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func synchronizeUpdateDataPDU(opt) {
var self = {
pad2Octets uint16)
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240608.aspx
 * @param updateData {type.Component} update data (ex: bitmapUpdateDataPDU)
 * @param opt {object} type option
 * @returns {type.Component}
 */
func updateDataPDU(updateData, opt) {
var self = {
__PDUTYPE2_ : PDUType2.PDUTYPE2_UPDATE__,
updateType uint16func() {
return self.updateData.obj.__UPDATE_TYPE__;
}),
updateData : updateData || new type.Factory(func(s) {
var options =  { readLength : new type.CallableValue(func() {
return opt.readLength.value - 2;
})};
switch(self.updateType.value) {
case UpdateType.UPDATETYPE_BITMAP:
self.updateData = bitmapUpdateDataPDU(null, options).read(s);
break;
case UpdateType.UPDATETYPE_SYNCHRONIZE:
// do nothing artefact of protocol
self.updateData = synchronizeUpdateDataPDU(null, options).read(s);
break;
default:
self.updateData = new type.BinaryString(null, options).read(s);
log.debug('unknown updateDataPDU ' + self.updateType.value);
}
})
};

return new type.Component(self, opt);
}

/**
 * @param pduData {type.Component}
 * @param shareId {integer}
 * @param opt {object} type option
 * @returns {type.Component}
 */
func dataPDU(pduData, shareId, opt) {
var self = {
__PDUTYPE__ : PDUType.PDUTYPE_DATAPDU,
shareDataHeader : shareDataHeader(new type.CallableValue(func() {
return new type.Component(self).size();
}), func() {
return self.pduData.obj.__PDUTYPE2__;
}, shareId),
pduData : pduData || new type.Factory(func(s) {

//compute local readLength
var options = { readLength : new type.CallableValue(func() {
return opt.readLength.value - self.shareDataHeader.size();
}) };

switch(self.shareDataHeader.obj.pduType2.value) {
case PDUType2.PDUTYPE2_SYNCHRONIZE:
self.pduData = synchronizeDataPDU(null, options).read(s)
break;
case PDUType2.PDUTYPE2_CONTROL:
self.pduData = controlDataPDU(null, options).read(s);
break;
case PDUType2.PDUTYPE2_SET_ERROR_INFO_PDU:
self.pduData = errorInfoDataPDU(null, options).read(s);
break;
case PDUType2.PDUTYPE2_FONTLIST:
self.pduData = fontListDataPDU(options).read(s);
break;
case PDUType2.PDUTYPE2_FONTMAP:
self.pduData = fontMapDataPDU(options).read(s);
break;
case PDUType2.PDUTYPE2_BITMAPCACHE_PERSISTENT_LIST:
self.pduData = persistentListPDU(null, options).read(s);
break;
case PDUType2.PDUTYPE2_INPUT:
self.pduData = clientInputEventPDU(null, options).read(s);
break;
case PDUType2.PDUTYPE2_SHUTDOWN_REQUEST:
self.pduData = shutdownRequestPDU(options).read(s);
break;
case PDUType2.PDUTYPE2_SHUTDOWN_DENIED:
self.pduData = shutdownDeniedPDU(options).read(s);
break;
case PDUType2.PDUTYPE2_SUPPRESS_OUTPUT:
self.pduData = supressOutputDataPDU(options).read(s);
break;
case PDUType2.PDUTYPE2_REFRESH_RECT:
self.pduData = refreshRectPDU(null, options).read(s);
break;
case PDUType2.PDUTYPE2_UPDATE:
self.pduData = updateDataPDU(null, options).read(s);
break;
default:
self.pduData = new type.BinaryString(null, options).read(s);
log.debug('unknown PDUType2 ' + self.shareDataHeader.obj.pduType2.value);
}
})
};

return new type.Component(self, opt);
}

/**
 * @param userId {integer}
 * @param pduMessage {type.Component} pdu message
 * @param opt {object} type option
 * @returns {type.Component}
 */
func pdu(userId, pduMessage, opt) {
var self = {
shareControlHeader : shareControlHeader(func() {
return new type.Component(self).size();
}, func() {
return pduMessage.obj.__PDUTYPE__;
}, userId),
pduMessage : pduMessage || new type.Factory(func(s) {

// compute local common options
var options = { readLength : new type.CallableValue(func() {
return self.shareControlHeader.obj.totalLength.value - self.shareControlHeader.size();
}) };

switch(self.shareControlHeader.obj.pduType.value) {
case PDUType.PDUTYPE_DEMANDACTIVEPDU:
self.pduMessage = demandActivePDU(null, options).read(s);
break;
case PDUType.PDUTYPE_CONFIRMACTIVEPDU:
self.pduMessage = confirmActivePDU(null, options).read(s);
break;
case PDUType.PDUTYPE_DEACTIVATEALLPDU:
self.pduMessage = deactiveAllPDU(options).read(s);
break;
case PDUType.PDUTYPE_DATAPDU:
self.pduMessage = dataPDU(null, null, options).read(s);
break;
default:
self.pduMessage = new type.BinaryString(null, options).read(s);
log.debug('unknown pdu type ' + self.shareControlHeader.obj.pduType.value);
}
})
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/dd306368.aspx
 * @param opt {object} type option
 * @returns {type.Component}
 */
func fastPathBitmapUpdateDataPDU (opt) {
var self = {
__FASTPATH_UPDATE_TYPE__ : FastPathUpdateType.FASTPATH_UPDATETYPE_BITMAP,
header uint16FastPathUpdateType.FASTPATH_UPDATETYPE_BITMAP, { constant : true }),
numberRectangles uint16 func () {
return self.rectangles.obj.length;
}),
rectangles : new type.Factory( func (s) {
self.rectangles = new type.Component([]);
for(var i = 0; i < self.numberRectangles.value; i++) {
self.rectangles.obj.push(bitmapData().read(s));
}
})
};

return new type.Component(self, opt);
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240622.aspx
 * @param updateData {type.Component}
 * @param opt {object} type option
 * @returns {type.Component}
 */
func fastPathUpdatePDU (updateData, opt) {
var self = {
updateHeader : new type.UInt8( func () {
return self.updateData.obj.__FASTPATH_UPDATE_TYPE__;
}),
compressionFlags : new type.UInt8(null, { conditional : func () {
return (self.updateHeader.value >> 4) & FastPathOutputCompression.FASTPATH_OUTPUT_COMPRESSION_USED;
}}),
size uint16 func () {
return self.updateData.size();
}),
updateData : updateData || new type.Factory( func (s) {
var options = { readLength : new type.CallableValue( func () {
return self.size.value;
}) };

switch (self.updateHeader.value & 0xf) {
case FastPathUpdateType.FASTPATH_UPDATETYPE_BITMAP:
self.updateData = fastPathBitmapUpdateDataPDU(options).read(s);
break;
default:
self.updateData = new type.BinaryString(null, options).read(s);
log.debug('unknown fast path pdu type ' + (self.updateHeader.value & 0xf));
}
})
};

return new type.Component(self, opt);
}

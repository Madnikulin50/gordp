package t125

import (
	"github.com/chuckpreslar/emission"
	"../../core"
)

type McsMessage uint8

const (
MCS_TYPE_CONNECT_INITIAL McsMessage = 0x65
MCS_TYPE_CONNECT_RESPONSE = 0x66
)

type McsDomainPDU uint16
const (
ERECT_DOMAIN_REQUEST McsDomainPDU = 1
DISCONNECT_PROVIDER_ULTIMATUM = 8
ATTACH_USER_REQUEST = 10
ATTACH_USER_CONFIRM = 11
CHANNEL_JOIN_REQUEST = 14
CHANNEL_JOIN_CONFIRM = 15
SEND_DATA_REQUEST = 25
SEND_DATA_INDICATION = 26
)

type McsChannel uint16
const (
MCS_GLOBAL_CHANNEL McsChannel = 1003
MCS_USERCHANNEL_BASE = 1001
)


type DomainParameters struct {
	MaxChannelIds int `asn1:"tag:2"`
	MaxUserIds int `asn1:"tag:2"`
	MaxTokenIds int `asn1:"tag:2"`
	NumPriorities int `asn1:"tag:2"`
	MinThoughput int `asn1:"tag:2"`
	MaxHeight int `asn1:"tag:2"`
	MaxMCSPDUsize int `asn1:"tag:2"`
	ProtocolVersion int `asn1:"tag:2"`
}

/**
 * @see http://www.itu.int/rec/T-REC-T.125-199802-I/en page 25
 * @returns {asn1.univ.Sequence}
 */
func NewDomainParameters(maxChannelIds int,
	maxUserIds int,
	maxTokenIds int,
	numPriorities int,
	minThoughput int,
	maxHeight int,
	maxMCSPDUsize int,
	protocolVersion int) *DomainParameters {
	return &DomainParameters{maxChannelIds, maxUserIds, maxTokenIds,
	numPriorities, minThoughput, maxHeight, maxMCSPDUsize, protocolVersion}
}

/**
 * @see http://www.itu.int/rec/T-REC-T.125-199802-I/en page 25
 * @param userData {Buffer}
 * @returns {asn1.univ.Sequence}
*/
type ConnectInitial struct {
	CallingDomainSelector []byte `asn1:"tag:4"`
	CalledDomainSelector []byte `asn1:"tag:4"`
	UpwardFlag bool
	TargetParameters DomainParameters
	MinimumParameters DomainParameters
	MaximumParameters DomainParameters
	UserData []byte `asn1:"application, tag:101"`
}

func NewConnectInitial (userData []byte) *ConnectInitial{
	return &ConnectInitial{[]byte{0x1},
	[]byte{0x1},
	false,
	*NewDomainParameters(34, 2, 0, 1, 0, 1, 0xffff, 2),
	*NewDomainParameters(1, 1, 1, 1, 0, 1, 0x420, 2),
	*NewDomainParameters(0xffff, 0xfc17, 0xffff, 1, 0, 1, 0xffff, 2),
	userData}
		/*userData : new asn1.univ.OctetString(userData)
}).implicitTag(new asn1.spec.Asn1Tag(asn1.spec.TagClass.Application, asn1.spec.TagFormat.Constructed, 101));*/
}

/**
 * @see http://www.itu.int/rec/T-REC-T.125-199802-I/en page 25
 * @returns {asn1.univ.Sequence}
*/

type ConnectResponse struct {
	result int `asn1:"tag:10"`
	calledConnectId int
	domainParameters DomainParameters
	userData []byte `asn1:"tag:10"`
	//.implicitTag(new asn1.spec.Asn1Tag(asn1.spec.TagClass.Application, asn1.spec.TagFormat.Constructed, 102));
}


func NewConnectResponse (userData [] byte) *ConnectResponse{
	return &ConnectResponse{0,
		0,
		*NewDomainParameters(22, 3, 0, 1, 0, 1, 0xfff8, 2),
		userData}
}

/**
 * Format MCS PDU header packet
 * @param mcsPdu {integer}
 * @param options {integer}
 * @returns {type.UInt8} headers
 */
func writeMCSPDUHeader(mcsPdu McsDomainPDU, options uint8, w core.Writer) {
	core.WriteUInt8((uint8(mcsPdu) << 2) | options, w)
}

/**
 * Read MCS PDU header
 * @param opcode
 * @param mcsPdu
 * @returns {Boolean}
 */
func readMCSPDUHeader(opcode, mcsPdu) {
	return (opcode >> 2) == mcsPdu;
}

/**
 * Multi-Channel Services
 * @param transport {events.EventEmitter} transport layer listen (connect, data) events
 * @param recvOpCode {DomainMCSPDU} opcode use in receive automata
 * @param sendOpCode {DomainMCSPDU} opcode use to send message
*/

type McsChannelInfo struct {
	id McsChannel
	name string
}

type Mcs struct {
	emission.Emitter
	transport core.Transport
	recvOpCode McsDomainPDU
	sendOpCode McsDomainPDU
	channels []McsChannelInfo
}

func NewMcs(transport core.Transport, recvOpCode McsDomainPDU, sendOpCode McsDomainPDU) *Mcs{
	mcs := Mcs{*emission.NewEmitter(), transport, recvOpCode, sendOpCode,
	[]McsChannelInfo{{MCS_GLOBAL_CHANNEL, "global"}}}
	mcs.transport.GetEvents().On("close", func() {
		mcs.Emitter.Emit("close")
	})
	mcs.transport.GetEvents().On("error", func(err interface{}) {
		mcs.Emitter.Emit("error", err)
	})
	return &mcs
}

/**
 * Send message to a specific channel
 * @param channelName {string} name of channel
 * @param data {type.*} message to send
*/
func (mcs *Mcs) Send(channelName string, data []byte) error {
	var channelId McsChannel = 0
	for _, channelInfo := range mcs.channels {
		if channelInfo.name == channelName {
			channelId = channelInfo.id
			break
		}
	}

	writeMCSPDUHeader(mcs.sendOpCode, 0, mcs.transport)
	WriteInteger16(mcs.userId, mcs.transport, MCS_USERCHANNEL_BASE)
	mcs.transport.Send()

this.transport.send(new type.Component([
writeMCSPDUHeader(this.sendOpCode),
per.writeInteger16(this.userId, Channel.MCS_USERCHANNEL_BASE),
per.writeInteger16(channelId),
new type.UInt8(0x70),
per.writeLength(data.size()),
data
]));
};

/**
 * Main receive function
 * @param s {type.Stream}
 */
MCS.prototype.recv = function(s) {
opcode = new type.UInt8().read(s).value;

if (readMCSPDUHeader(opcode, DomainMCSPDU.DISCONNECT_PROVIDER_ULTIMATUM)) {
log.info("MCS DISCONNECT_PROVIDER_ULTIMATUM");
this.transport.close();
return
}
else if(!readMCSPDUHeader(opcode, this.recvOpCode)) {
throw new error.UnexpectedFatalError('NODE_RDP_PROTOCOL_T125_MCS_BAD_RECEIVE_OPCODE');
}

per.readInteger16(s, Channel.MCS_USERCHANNEL_BASE);

var channelId = per.readInteger16(s);

per.readEnumerates(s);
per.readLength(s);

var channelName = this.channels.find(function(e) {
if (e.id === channelId) return true;
}).name;

this.emit(channelName, s);
};

/**
 * Only main channels handle actually
 * @param transport {event.EventEmitter} bind connect and data events
 * @returns
 */
type McsClient struct {
	Mcs
	channelsConnected int,
	clientCoreData,
	clientNetworkData,
	clientSecurityData,
	serverCoreData,
	serverSecurityData,
	serverNetworkData

}

func NewMcsClient(transport core.Transport) {
	client := McsClient{ NewMcs(transport,SEND_DATA_INDICATION, SEND_DATA_REQUEST), 0, gcc.clientCoreData();
this.clientNetworkData = gcc.clientNetworkData(new type.Component([]));
this.clientSecurityData = gcc.clientSecurityData();

// must be readed from protocol
this.serverCoreData = null;
this.serverSecurityData = null;
this.serverNetworkData = null;

var self = this;
this.transport.on('connect', function(s) {
self.connect(s);
});
}

inherits(Client, MCS);

/**
 * Connect event layer
 */
Client.prototype.connect = function(selectedProtocol) {
this.clientCoreData.obj.serverSelectedProtocol.value = selectedProtocol;
this.sendConnectInitial();
};

/**
 * close stack
 */
Client.prototype.close = function() {
this.transport.close();
};

/**
 * MCS connect response (server->client)
 * @param s {type.Stream}
 */
Client.prototype.recvConnectResponse = function(s) {
var userData = new type.Stream(ConnectResponse().decode(s, asn1.ber).value.userData.value);
var serverSettings = gcc.readConferenceCreateResponse(userData);
// record server gcc block
for(var i in serverSettings) {
if(!serverSettings[i].obj) {
continue;
}
switch(serverSettings[i].obj.__TYPE__) {
case gcc.MessageType.SC_CORE:
this.serverCoreData = serverSettings[i];
break;
case gcc.MessageType.SC_SECURITY:
this.serverSecurityData = serverSettings[i];
break;
case gcc.MessageType.SC_NET:
this.serverNetworkData = serverSettings[i];
break;
default:
log.warn('unhandle server gcc block = ' + serverSettings[i].obj.__TYPE__);
}
}

// send domain request
this.sendErectDomainRequest();
// send attach user request
this.sendAttachUserRequest();
// now wait user confirm from server
var self = this;
this.transport.once('data', function(s) {
self.recvAttachUserConfirm(s);
});
};

/**
 * MCS connection automata step
 * @param s {type.Stream}
 */
Client.prototype.recvAttachUserConfirm = function(s) {
if (!readMCSPDUHeader(new type.UInt8().read(s).value, DomainMCSPDU.ATTACH_USER_CONFIRM)) {
throw new error.UnexpectedFatalError('NODE_RDP_PROTOCOL_T125_MCS_BAD_HEADER');
}

if (per.readEnumerates(s) !== 0) {
throw new error.UnexpectedFatalError('NODE_RDP_PROTOCOL_T125_MCS_SERVER_REJECT_USER');
}

this.userId = per.readInteger16(s, Channel.MCS_USERCHANNEL_BASE);
//ask channel for specific user
this.channels.push({ id : this.userId, name : 'user' });
// channel connect automata
this.connectChannels();
};

/**
 * Last state in channel connection automata
 * @param s {type.Stream}
 */
Client.prototype.recvChannelJoinConfirm = function(s) {
var opcode = new type.UInt8().read(s).value;

if (!readMCSPDUHeader(opcode, DomainMCSPDU.CHANNEL_JOIN_CONFIRM)) {
throw new error.UnexpectedFatalError('NODE_RDP_PROTOCOL_T125_MCS_WAIT_CHANNEL_JOIN_CONFIRM');
}

var confirm = per.readEnumerates(s);

var userId = per.readInteger16(s, Channel.MCS_USERCHANNEL_BASE);
if (this.userId !== userId) {
throw new error.UnexpectedFatalError('NODE_RDP_PROTOCOL_T125_MCS_INVALID_USER_ID');
}

var channelId = per.readInteger16(s);

if ((confirm !== 0) && (channelId === Channel.MCS_GLOBAL_CHANNEL || channelId === this.userId)) {
throw new error.UnexpectedFatalError('NODE_RDP_PROTOCOL_T125_MCS_SERVER_MUST_CONFIRM_STATIC_CHANNEL');
}

this.connectChannels();
};

/**
 * First MCS message
 */
Client.prototype.sendConnectInitial = function() {

var ccReq = gcc.writeConferenceCreateRequest(new type.Component([
gcc.block(this.clientCoreData),
gcc.block(this.clientNetworkData),
gcc.block(this.clientSecurityData)
])).toStream().getValue();

this.transport.send(ConnectInitial(ccReq).encode(asn1.ber));

// next event is connect response
var self = this;
this.transport.once('data', function(s) {
self.recvConnectResponse(s);
});
};

/**
 * MCS connection automata step
 */
Client.prototype.sendErectDomainRequest = function() {
this.transport.send(new type.Component([
writeMCSPDUHeader(DomainMCSPDU.ERECT_DOMAIN_REQUEST),
per.writeInteger(0),
per.writeInteger(0)
]));
};

/**
 * MCS connection automata step
 */
Client.prototype.sendAttachUserRequest = function() {
this.transport.send(writeMCSPDUHeader(DomainMCSPDU.ATTACH_USER_REQUEST));
};

/**
 * Send a channel join request
 * @param channelId {integer} channel id
 */
Client.prototype.sendChannelJoinRequest = function(channelId) {
this.transport.send(new type.Component([
writeMCSPDUHeader(DomainMCSPDU.CHANNEL_JOIN_REQUEST),
per.writeInteger16(this.userId, Channel.MCS_USERCHANNEL_BASE),
per.writeInteger16(channelId)
]));
};

/**
 * Connect channels automata
 * @param s {type.Stream}
 */
Client.prototype.connectChannels = function(s) {
if(this.channelsConnected == this.channels.length) {
var self = this;
this.transport.on('data', function(s) {
self.recv(s);
});

// send client and sever gcc informations
this.emit('connect',
{
core : this.clientCoreData.obj,
security : this.clientSecurityData.obj,
net : this.clientNetworkData.obj
},
{
core : this.serverCoreData.obj,
security : this.serverSecurityData.obj
}, this.userId, this.channels);
return;
}

this.sendChannelJoinRequest(this.channels[this.channelsConnected++].id);

var self = this;
this.transport.once('data', function(s) {
self.recvChannelJoinConfirm(s);
});
};


package protocol

import (
	"net"
	"../core"
	"./x224"
)

/**
 * Main RDP module
 */
type RdpClient struct {
	config map[string]string
	connected bool
	bufferLayer core.BufferLayer
	tpkt *TPKT
	x224 *X224Client
	mcs * MCSClient



}

func NewRdpClient(config *map[string]string) {
	client := RdpClient{}
	client.config = config || make(map[string]string)
	client.connected = false
	client.bufferLayer = core.NewLayer(nil)
	client.tpkt = NewTPKT(client.bufferLayer)
	client.x224 = new x224.NewClient(client.tpkt)
	client.mcs = new t125.mcs.Client(client.x224);
	client.sec = new pdu.sec.Client(client.mcs, client.tpkt);
	client.global = new pdu.global.Client(this.sec, this.sec);

// config log level
log.level = log.Levels[config.logLevel || 'INFO'] || log.Levels.INFO;

// credentials
if (config.domain) {
this.sec.infos.obj.domain.value = new Buffer(config.domain + '\x00', 'ucs2');
}
if (config.userName) {
this.sec.infos.obj.userName.value = new Buffer(config.userName + '\x00', 'ucs2');
}
if (config.password) {
this.sec.infos.obj.password.value = new Buffer(config.password + '\x00', 'ucs2');
}

if (config.enablePerf) {
this.sec.infos.obj.extendedInfo.obj.performanceFlags.value =
pdu.sec.PerfFlag.PERF_DISABLE_WALLPAPER
| 	pdu.sec.PerfFlag.PERF_DISABLE_MENUANIMATIONS
| 	pdu.sec.PerfFlag.PERF_DISABLE_CURSOR_SHADOW
| 	pdu.sec.PerfFlag.PERF_DISABLE_THEMING
| 	pdu.sec.PerfFlag.PERF_DISABLE_FULLWINDOWDRAG;
}

if (config.autoLogin) {
this.sec.infos.obj.flag.value |= pdu.sec.InfoFlag.INFO_AUTOLOGON;
}

if (config.screen && config.screen.width && config.screen.height) {
this.mcs.clientCoreData.obj.desktopWidth.value = config.screen.width;
this.mcs.clientCoreData.obj.desktopHeight.value = config.screen.height;
}

log.info('screen ' + this.mcs.clientCoreData.obj.desktopWidth.value + 'x' + this.mcs.clientCoreData.obj.desktopHeight.value);

// config keyboard layout
switch (config.locale) {
case 'fr':
log.info('french keyboard layout');
this.mcs.clientCoreData.obj.kbdLayout.value = t125.gcc.KeyboardLayout.FRENCH;
break;
case 'en':
default:
log.info('english keyboard layout');
this.mcs.clientCoreData.obj.kbdLayout.value = t125.gcc.KeyboardLayout.US;
}


//bind all events
var self = this;
this.global.on('connect', function () {
self.connected = true;
self.emit('connect');
}).on('session', function () {
self.emit('session');
}).on('close', function () {
self.connected = false;
self.emit('close');
}).on('bitmap', function (bitmaps) {
for(var bitmap in bitmaps) {
var bitmapData = bitmaps[bitmap].obj.bitmapDataStream.value;
var isCompress = bitmaps[bitmap].obj.flags.value & pdu.data.BitmapFlag.BITMAP_COMPRESSION;

if (isCompress && config.decompress) {
bitmapData = decompress(bitmaps[bitmap].obj);
isCompress = false;
}

self.emit('bitmap', {
destTop : bitmaps[bitmap].obj.destTop.value,
destLeft : bitmaps[bitmap].obj.destLeft.value,
destBottom : bitmaps[bitmap].obj.destBottom.value,
destRight : bitmaps[bitmap].obj.destRight.value,
width : bitmaps[bitmap].obj.width.value,
height : bitmaps[bitmap].obj.height.value,
bitsPerPixel : bitmaps[bitmap].obj.bitsPerPixel.value,
isCompress : isCompress,
data : bitmapData
});
}
}).on('error', function (err) {
log.error(err.code + '(' + err.message + ')\n' + err.stack);
if (err instanceof error.FatalError) {
throw err;
}
else {
self.emit('error', err);
}
});
}

inherits(RdpClient, events.EventEmitter);

/**
 * Connect RDP client
 * @param host {string} destination host
 * @param port {integer} destination port
 */
RdpClient.prototype.connect = function (host, port) {
log.info('connect to ' + host + ':' + port);
var self = this;
this.bufferLayer.socket.connect(port, host, function () {
// in client mode connection start from x224 layer
self.x224.connect();
});
return this;
};

/**
 * Close RDP client
 */
RdpClient.prototype.close = function () {
if(this.connected) {
this.global.close();
}
this.connected = false;
return this;
};

/**
 * Send pointer event to server
 * @param x {integer} mouse x position
 * @param y {integer} mouse y position
 * @param button {integer} button number of mouse
 * @param isPressed {boolean} state of button
 */
RdpClient.prototype.sendPointerEvent = function (x, y, button, isPressed) {
if (!this.connected)
return;

var event = pdu.data.pointerEvent();
if (isPressed) {
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_DOWN;
}

switch(button) {
case 1:
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_BUTTON1;
break;
case 2:
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_BUTTON2;
break;
case 3:
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_BUTTON3
break;
default:
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_MOVE;
}

event.obj.xPos.value = x;
event.obj.yPos.value = y;

this.global.sendInputEvents([event]);
};

/**
 * send scancode event
 * @param code {integer}
 * @param isPressed {boolean}
 * @param extended {boolenan} extended keys
 */
RdpClient.prototype.sendKeyEventScancode = function (code, isPressed, extended) {
if (!this.connected)
return;
extended = extended || false;
var event = pdu.data.scancodeKeyEvent();
event.obj.keyCode.value = code;

if (!isPressed) {
event.obj.keyboardFlags.value |= pdu.data.KeyboardFlag.KBDFLAGS_RELEASE;
}

if (extended) {
event.obj.keyboardFlags.value |= pdu.data.KeyboardFlag.KBDFLAGS_EXTENDED;
}

this.global.sendInputEvents([event]);
};

/**
 * Send key event as unicode
 * @param code {integer}
 * @param isPressed {boolean}
 */
RdpClient.prototype.sendKeyEventUnicode = function (code, isPressed) {
if (!this.connected)
return;

var event = pdu.data.unicodeKeyEvent();
event.obj.unicode.value = code;

if (!isPressed) {
event.obj.keyboardFlags.value |= pdu.data.KeyboardFlag.KBDFLAGS_RELEASE;
}
this.global.sendInputEvents([event]);
}

/**
 * Wheel mouse event
 * @param x {integer} mouse x position
 * @param y {integer} mouse y position
 * @param step {integer} wheel step
 * @param isNegative {boolean}
 * @param isHorizontal {boolean}
 */
RdpClient.prototype.sendWheelEvent = function (x, y, step, isNegative, isHorizontal) {
if (!this.connected)
return;

var event = pdu.data.pointerEvent();
if (isHorizontal) {
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_HWHEEL;
}
else {
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_WHEEL;
}


if (isNegative) {
event.obj.pointerFlags.value |= pdu.data.PointerFlag.PTRFLAGS_WHEEL_NEGATIVE;
}

event.obj.pointerFlags.value |= (step & pdu.data.PointerFlag.WheelRotationMask)

event.obj.xPos.value = x;
event.obj.yPos.value = y;

this.global.sendInputEvents([event]);
}

package core/*


/**
 * Buffer data from socket to present
 * well formed packets
*/import (
	"net"
	"crypto/tls"
	"github.com/chuckpreslar/emission"
)

type BufferLayer struct {
	conn *net.Conn
 	tlsConn *tls.Conn
 	buffers [][]byte
	bufferLength int
	expectedSize int
}

func NewBufferLayer(conn *net.Conn) *BufferLayer {
	conn.
	bl := BufferLayer{}
	tlsConn.securePair = nil
	conn.socket = conn
	conn

var self = this;
// bind event
this.socket.on('data', function(data) {
try {
self.recv(data);
}
catch(e) {
self.socket.destroy();
self.emit('error', e);
}
}).on('close', function() {
self.emit('close');
}).on('error', function (err) {
self.emit('error', err);
});

//buffer data
	 return &conn
}

inherits(BufferLayer, events.EventEmitter);

/**
 * Call from tcp layer
 * @param data tcp stream
 */
BufferLayer.prototype.recv = function(data) {
this.buffers[this.buffers.length] = data;
this.bufferLength += data.length;

while(this.bufferLength >= this.expectedSize) {
//linear buffer
var expectedData = new type.Stream(this.expectedSize);

//create expected data
while(expectedData.availableLength() > 0) {

var rest = expectedData.availableLength();
var buffer = this.buffers.shift();

if(buffer.length > expectedData.availableLength()) {
this.buffers.unshift(buffer.slice(rest));
new type.BinaryString(buffer, { readLength : new type.CallableValue(expectedData.availableLength()) }).write(expectedData);
}
else {
new type.BinaryString(buffer).write(expectedData);
}
}

this.bufferLength -= this.expectedSize;
expectedData.offset = 0;
this.emit('data', expectedData);
}
};

/**
 * Call tcp socket to write stream
 * @param {type.Type} packet
 */
BufferLayer.prototype.send = function(data) {
var s = new type.Stream(data.size());
data.write(s);
if(this.securePair) {
this.securePair.cleartext.write(s.buffer);
}
else {
this.socket.write(s.buffer);
}
};

/**
 * Wait expected size data before call callback function
 * @param {number} expectSize	size expected
 */
BufferLayer.prototype.expect = function(expectedSize) {
this.expectedSize = expectedSize;
};

/**
 * Convert connection to TLS connection
 * Use nodejs starttls module
 * @param callback {func} when connection is done
 */
BufferLayer.prototype.startTLS = function(callback) {
var options = {
socket : this.socket,
pair : tls.createSecurePair(crypto.createCredentials(), false, false, false)
};
var self = this;
this.securePair = starttls(options, function(err) {
log.warn(err);
callback();
})

this.securePair.cleartext.on('data', function(data) {
try {
self.recv(data);
}
catch(e) {
self.socket.destroy();
self.emit('error', e);
}
}).on('error', function (err) {
self.emit('error', err);
});
};

/**
 * Convert connection to TLS server
 * @param keyFilePath	{string} key file path
 * @param crtFilePath	{string} certificat file path
 * @param callback	{function}
 */

 /*BufferLayer.prototype.listenTLS = function(keyFilePath, crtFilePath, callback) {
var options = {
socket : this.socket,
pair : tls.createSecurePair(crypto.createCredentials({
key: fs.readFileSync(keyFilePath),
cert: fs.readFileSync(crtFilePath),
}), true, false, false)
};
var self = this;
this.securePair = starttls(options, function(err) {
log.warn(err);
self.cleartext = this.cleartext;
callback();
});

this.securePair.cleartext.on('data', function(data) {
try {
self.recv(data);
}
catch(e) {
self.socket.destroy();
self.emit('error', e);
}
}).on('error', function (err) {
self.emit('error', err);
});
};*/

/**
 * close stack
 */
func (layer *BufferLayer) Close () {
	layer.conn.Close()
}


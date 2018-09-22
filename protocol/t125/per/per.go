package t125

import (
	"../../core"
	"errors"
)

/**
 * @param s {type.Stream} read value from stream
 * @returns read length from per format
 */
func ReadLength(r core.Reader) uint16 {
	byte := core.ReadUInt8(r)

	if byte & 0x80 != 0{
		var size uint16= 0
		byte = byte &^0x80
		size = uint16(byte) << 8
		size += uint16(core.ReadUInt8(r))
		return size
	} else {
		return uint16(byte)
	}
}

/**
 * @param value {raw} value to convert to per format
 * @returns type objects per encoding value
 */
func WriteLength(value uint16, w core.Writer) {
	if value > 0x7f {
		core.WriteUInt16BE(value | 0x8000, w)
	} else {
		core.WriteUInt8(uint8(value), w)
	}
}

/**
 * @param s {type.Stream}
 * @returns {integer} choice decoding from per encoding
 */
func ReadChoice(r core.Reader) uint8 {
	return core.ReadUInt8(r)
}

/**
 * @param choice {integer}
 * @returns {type.UInt8} choice per encoded
 */
func WriteChoice(choice uint8, w core.Writer) {
	core.WriteUInt8(choice, w)
}

/**
 * @param s {type.Stream}
 * @returns {integer} number represent selection
*/
func ReadSelection(r core.Reader) uint8 {
	return core.ReadUInt8(r)
}


/**
 * @param selection {integer}
 * @returns {type.UInt8} per encoded selection
 */
func WriteSelection(selection uint8, w core.Writer) {
	core.WriteUInt8(selection, w)
}
/**
 * @param s {type.Stream}
 * @returns {integer} number of sets
 */
func ReadNumberOfSet(r core.Reader) uint8 {
	return core.ReadUInt8(r)
}

/**
 * @param numberOfSet {integer}
 * @returns {type.UInt8} per encoded nuimber of sets
 */
func WriteNumberOfSet(numberOfSet uint8, w core.Writer) {
	core.WriteUInt8(numberOfSet, w)
}

/**
 * @param s {type.Stream}
 * @returns {integer} enumerates number
 */
func ReadEnumerates(r core.Reader) uint8 {
	return core.ReadUInt8(r)
}

/**
 * @param enumerate {integer}
 * @returns {type.UInt8} per encoded enumerate
 */
func WriteEnumerates(enumerate uint8, w core.Writer) {
	core.WriteUInt8(enumerate, w)
}

/**
 * @param s {type.Stream}
 * @returns {integer} integer per decoded
 */
func ReadInteger(r core.Reader) uint32 {
	size := ReadLength(r)
	switch size	{
	case 1:
		return uint32(core.ReadUInt8(r))
	case 2:
		return uint32(core.ReadUInt16BE(r))
	case 4:
		return uint32(core.ReadUInt32BE(r))
	default:
		panic(errors.New("NODE_RDP_PROTOCOL_T125_PER_BAD_INTEGER_LENGTH"))
	}
}

/**
 * @param value {integer}
 * @returns {type.Component} per encoded integer
 */
func WriteInteger(value uint32, w core.Writer) {
	if value <= 0xff {
		WriteLength(1, w)
		core.WriteUInt8(uint8(value), w)
	} else if value < 0xffff {
		WriteLength(2, w)
		core.WriteUInt16BE(uint16(value), w)
	} else {
		WriteLength(4, w)
		core.WriteUInt32BE(value, w)
	}
}

/**
 * @param s {type.Stream}
 * @param minimum {integer} increment (default 0)
 * @returns {integer} per decoded integer 16 bits
 */
func ReadInteger16(r core.Reader, minimum uint16) uint16 {
	return core.ReadUInt16BE(r) + minimum
}

/**
 * @param value {integer}
 * @param minimum {integer} decrement (default 0)
 * @returns {type.UInt16Be} per encoded integer 16 bits
 */
func WriteInteger16(value uint16, w core.Writer, minimum uint16) {
	core.WriteUInt16BE(value - minimum, w)
}

/**
 * Check object identifier
 * @param s {type.Stream}
 * @param oid {array} object identifier to check
 */
func ReadObjectIdentifier(r core.Reader, oid []byte) bool {
	size := ReadLength(r)
	if size != 5 {
		return false
	}

	a_oid := []byte{0, 0, 0, 0, 0, 0}
	t12 := core.ReadByte(r)
	a_oid[0] = t12 >> 4
	a_oid[1] = t12 & 0x0f
	a_oid[2] = core.ReadByte(r)
	a_oid[3] = core.ReadByte(r)
	a_oid[4] = core.ReadByte(r)
	a_oid[5] = core.ReadByte(r)

	for i, _ := range oid {
		if oid[i] != a_oid[i] {
			return false
		}
	}
	return true
}

/**
 * @param oid {array} oid to write
 * @returns {type.Component} per encoded object identifier
 */
func WriteObjectIdentifier(oid []byte, w core.Writer) {
	core.WriteUInt8(5, w)
	core.WriteByte((oid[0] << 4) & (oid[1] & 0x0f), w)
	core.WriteByte(oid[2], w)
	core.WriteByte(oid[3], w)
	core.WriteByte(oid[4], w)
	core.WriteByte(oid[5], w)
}

/**
 * Read as padding...
 * @param s {type.Stream}
 * @param minValue
 */
func ReadNumericString(r core.Reader, minValue uint16) {
	length := ReadLength(r)
	length = (length + minValue + 1) / 2
	core.ReadPadding(int(length), r)
}

/**
 * @param nStr {String}
 * @param minValue {integer}
 * @returns {type.Component} per encoded numeric string
 */
func WriteNumericString(nStr string, w core.Writer, minValue int) {
	panic(errors.New("not implemented"))
	/*length := nStr.
	mlength := minValue
	if length - minValue >= 0 {
		mlength = length - minValue
	}

	result := make([]byte, 0, 100)

	for i, _ := range nStr
var c1 = nStr.charCodeAt(i);
var c2 = 0;
if(i + 1 < length) {
c2 = nStr.charCodeAt(i + 1);
}
else {
c2 = 0x30;
}
c1 = (c1 - 0x30) % 10;
c2 = (c2 - 0x30) % 10;

result[result.length] = new type.UInt8((c1 << 4) | c2);
}

return new type.Component([writeLength(mlength), new type.Component(result)]);*/
}

/**
 * @param s {type.Stream}
 * @param length {integer} length of padding
 */
func readPadding(length int, r core.Reader) {
	core.ReadPadding(length, r)
}

/**
 * @param length {integer} length of padding
 * @returns {type.BinaryString} per encoded padding
 */
func WritePadding(length int, w core.Writer) {
	core.WritePadding(length + 1, w)
}

/**
 * @param s {type.Stream}
 * @param octetStream {String}
 * @param minValue {integer} default 0
 * @returns {Boolean} true if read octectStream is equal to octetStream
 */
func ReadOctetStream(octetStream []byte, minValue int, r core.Reader) bool {
	size := int(ReadLength(r)) + minValue
	if size != len(octetStream) {
		return false
	}
	for i := 0; i < size; i++ {
		if core.ReadByte(r) != octetStream[i] {
			return false
		}
	}
	return true
}

/**
 * @param oStr {String}
 * @param minValue {integer} default 0
 * @returns {type.Component} per encoded octet stream
 */
func WriteOctetStream(oStr string, minValue int, w core.Writer) {
	length := len(oStr)
	mlength := minValue

	if length - minValue >= 0 {
		mlength = length - minValue
	}


	w.Write([]byte(oStr)[:mlength])
}



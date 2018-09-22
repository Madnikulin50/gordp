package asn1

import (
	"../core"
	"errors"
)

/**
 * Parse tag(T) field of BER TLV
 * And check with expected tag
 * @param s {type.Stream}
 * @param tag {spec.tag}
 * @returns {Boolean} True for valid tag matching
 */
func decodeTag(tag Tag, r core.Reader) bool {
	nextTag := core.ReadUInt8(r)
	var nextTagNumber uint8
	if tag.TagNumber > 30 {
		nextTagNumber = core.ReadUInt8(r)
	} else {
		nextTagNumber = nextTag & 0x1F
	}

	return ((nextTag & 0xE0) == uint8(tag.TagClass) | uint8(tag.TagFormat)) && (nextTagNumber == tag.TagNumber)
}

/**
 * Parse length(L) field of BER TLV
 * @param s {type.Stream}
 * @returns {integer}
 */
func decodeLength(r core.Reader) uint16 {
	size := core.ReadUInt8(r)
	if size & 0x80 != 0 {
		size &= ^0x80
		if size == 1 {
			return uint16(core.ReadUInt8(r))
		} else if size == 2 {
			return core.ReadUInt16BE(r)
		} else {
			panic(errors.New("NODE_RDP_ASN1_BER_INVALID_LENGTH"))
		}
	}
	return uint16(size)
};

/**
 * Decode tuple TLV (Tag Length Value) of BER
 * @param s {type.Stream}
 * @param tag {spec.Asn1Tag} expected tag
 * @returns {type.BinaryString} Value of tuple
 */
func decode(tag Tag, r core.Reader) []byte {
	if decodeTag(tag, r) == false {
		panic(errors.New("NODE_RDP_ASN1_BER_INVALID_TAG"))
	}
	length := decodeLength(r)

	if length == 0 {
		return []byte{}
	}
	result := make([]byte, length)
	r.Read(result)
	return result
}

func encodeTag(tag Tag) []byte {
	if tag.TagNumber > 30 {
		return []byte{byte(tag.TagClass) | byte(tag.TagFormat) | 0x1F, tag.TagNumber}
	} else {
		return []byte{byte(tag.TagClass) | byte(tag.TagFormat) | (byte(tag.TagNumber) & 0x1F)}
	}
}

func encodeLength(length uint16) []byte {
	if length > 0x7f {
		return []byte{0x82, byte(length >> 8), byte(length)}
	} else {
		return []byte{byte(uint8(length))}
	}
}

func encode(tag Tag, buffer []byte, w core.Writer) {
	w.Write(encodeTag(tag))
	w.Write(encodeLength(uint16(len(buffer))))
	w.Write(buffer)
}

type BerClass struct {

}

func (b *BerClass) Decode(tag Tag, r core.Reader) []byte {
	return decode(tag, r)
}
func (b *BerClass) Encode(tag Tag, buffer []byte, w core.Writer) {
	encode(tag, buffer, w)
}

var Ber BerClass = BerClass{}
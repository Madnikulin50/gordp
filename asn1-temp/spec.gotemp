package asn1

import (
	"../core"
	"errors"
)

/**
 * Tag Class
 */
type TagClass uint8
const (
	Universal TagClass = 0x00
	Application = 0x40
	Context = 0x80
	Private = 0xC0
)

/**
 * Tag Format
 */
type TagFormat uint8
const (
	Primitive TagFormat = 0x00
	Constructed = 0x20
)

/**
 * ASN.1 tag
 * @param tagClass {TagClass}
 * @param tagFormat {TagFormat}
 * @param tagNumber {integer}
*/

type Tag struct {
	TagClass TagClass
	TagFormat TagFormat
	TagNumber uint8
}

func NewAsn1Tag(tagClass TagClass, tagFormat TagFormat, tagNumber uint8) *Tag {
	return &Tag{tagClass,  tagFormat,  tagNumber}
}

/**
 * ASN.1 Specification
 * @param tag {Asn1Tag}
 */

type Spec struct {
	Tag Tag
	Opt bool
}

func NewSpec(tag Tag) *Spec {
	return &Spec{tag, false}
}

/**
 * Add an implicit tag
 * override tag
 * @param tag {Asn1Tag}
 * @returns {Asn1Spec}
*/
func (spec *Spec) ImplicitTag(tag Tag) *Spec {
	spec.Tag = tag
	return spec
}

/**
 * Set optional to true
 * @returns {Asn1Spec}
 */
func (spec *Spec) Optional () *Spec {
	spec.Opt = true
	return spec
}

/**
 * Add explicit tag
 * Append new tag header to existing tag
 * @param tag {Asn1Tag}
 * @returns {Asn1SpecExplicitTag}
 */
func (spec *Spec) ExplicitTag (tag Tag) *SpecExplicitTag {
	return NewSpecExplicitTag(tag, spec)
}

type Decoder interface {
	Decode(tag Tag, r core.Reader) []byte
	Encode(tag Tag, buffer []byte, w core.Writer)
}

/**
 * Decode must be implemented by all sub type
 * @param s {type.Stream}
 * @param decoder
 */
func (spec *Spec) Decode(r core.Reader, decoder interface {}) {
	panic(errors.New("NODE_RDP_AS1_SPEC_DECODE_NOT_IMPLEMENTED"))
}
/*
Asn1Spec.prototype.decode = function(s, decoder) {
var specStream = new type.Stream(decoder.decode(s, this.tag).value);
this.spec.decode(specStream, decoder);
};*/

/**
 * Encode must be implemented by all sub type
 * @param decoder
*/
func (spec *Spec) Encode(w core.Writer, encoder interface {}) {
	panic(errors.New("NODE_RDP_AS1_SPEC_ENCODE_NOT_IMPLEMENTED"))
}


/**
 * Component Asn1Spec object
*/
type SpecExplicitTag struct {
	Spec
	spec *Spec
}

func NewSpecExplicitTag (tag Tag, spec *Spec) *SpecExplicitTag {
	return &SpecExplicitTag{*NewSpec(tag), spec}
}

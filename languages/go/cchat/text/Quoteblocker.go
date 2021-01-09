// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package text

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Quoteblocker struct {
	_tab flatbuffers.Table
}

func GetRootAsQuoteblocker(buf []byte, offset flatbuffers.UOffsetT) *Quoteblocker {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Quoteblocker{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsQuoteblocker(buf []byte, offset flatbuffers.UOffsetT) *Quoteblocker {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Quoteblocker{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Quoteblocker) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Quoteblocker) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Quoteblocker) QuotePrefix() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func QuoteblockerStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func QuoteblockerAddQuotePrefix(builder *flatbuffers.Builder, quotePrefix flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(quotePrefix), 0)
}
func QuoteblockerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
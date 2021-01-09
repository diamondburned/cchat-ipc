// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package core

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StringPair struct {
	_tab flatbuffers.Table
}

func GetRootAsStringPair(buf []byte, offset flatbuffers.UOffsetT) *StringPair {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StringPair{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStringPair(buf []byte, offset flatbuffers.UOffsetT) *StringPair {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StringPair{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StringPair) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StringPair) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StringPair) K() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *StringPair) V() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StringPairStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StringPairAddK(builder *flatbuffers.Builder, k flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(k), 0)
}
func StringPairAddV(builder *flatbuffers.Builder, v flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(v), 0)
}
func StringPairEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

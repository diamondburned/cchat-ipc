// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package core

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Error struct {
	_tab flatbuffers.Table
}

func GetRootAsError(buf []byte, offset flatbuffers.UOffsetT) *Error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Error{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsError(buf []byte, offset flatbuffers.UOffsetT) *Error {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Error{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Error) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Error) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Error) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ErrorStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ErrorAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(error), 0)
}
func ErrorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
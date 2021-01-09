// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type CompleteParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsCompleteParameters(buf []byte, offset flatbuffers.UOffsetT) *CompleteParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CompleteParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCompleteParameters(buf []byte, offset flatbuffers.UOffsetT) *CompleteParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CompleteParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CompleteParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CompleteParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CompleteParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__call.ID)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *CompleteParameters) Words(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *CompleteParameters) WordsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CompleteParameters) Current() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CompleteParameters) MutateCurrent(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func CompleteParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CompleteParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func CompleteParametersAddWords(builder *flatbuffers.Builder, words flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(words), 0)
}
func CompleteParametersStartWordsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CompleteParametersAddCurrent(builder *flatbuffers.Builder, current int64) {
	builder.PrependInt64Slot(2, current, 0)
}
func CompleteParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

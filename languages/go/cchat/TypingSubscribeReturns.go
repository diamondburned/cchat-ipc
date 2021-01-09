// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
	cchat__core "cchat/core"
)

type TypingSubscribeReturns struct {
	_tab flatbuffers.Table
}

func GetRootAsTypingSubscribeReturns(buf []byte, offset flatbuffers.UOffsetT) *TypingSubscribeReturns {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TypingSubscribeReturns{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTypingSubscribeReturns(buf []byte, offset flatbuffers.UOffsetT) *TypingSubscribeReturns {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TypingSubscribeReturns{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TypingSubscribeReturns) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TypingSubscribeReturns) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TypingSubscribeReturns) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *TypingSubscribeReturns) Err(obj *cchat__core.Error) *cchat__core.Error {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__core.Error)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *TypingSubscribeReturns) StopHandle(obj *cchat__call.StopHandle) *cchat__call.StopHandle {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__call.StopHandle)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func TypingSubscribeReturnsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TypingSubscribeReturnsAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func TypingSubscribeReturnsAddErr(builder *flatbuffers.Builder, err flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(err), 0)
}
func TypingSubscribeReturnsAddStopHandle(builder *flatbuffers.Builder, stopHandle flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(stopHandle), 0)
}
func TypingSubscribeReturnsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

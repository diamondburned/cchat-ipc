// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
	cchat__core "cchat/core"
)

type SendReturns struct {
	_tab flatbuffers.Table
}

func GetRootAsSendReturns(buf []byte, offset flatbuffers.UOffsetT) *SendReturns {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SendReturns{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSendReturns(buf []byte, offset flatbuffers.UOffsetT) *SendReturns {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SendReturns{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SendReturns) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SendReturns) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SendReturns) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *SendReturns) Err(obj *cchat__core.Error) *cchat__core.Error {
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

func SendReturnsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SendReturnsAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func SendReturnsAddErr(builder *flatbuffers.Builder, err flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(err), 0)
}
func SendReturnsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type LoadMoreReturns struct {
	_tab flatbuffers.Table
}

func GetRootAsLoadMoreReturns(buf []byte, offset flatbuffers.UOffsetT) *LoadMoreReturns {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LoadMoreReturns{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLoadMoreReturns(buf []byte, offset flatbuffers.UOffsetT) *LoadMoreReturns {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LoadMoreReturns{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LoadMoreReturns) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LoadMoreReturns) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LoadMoreReturns) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *LoadMoreReturns) Bool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *LoadMoreReturns) MutateBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func LoadMoreReturnsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func LoadMoreReturnsAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func LoadMoreReturnsAddBool(builder *flatbuffers.Builder, bool bool) {
	builder.PrependBoolSlot(1, bool, false)
}
func LoadMoreReturnsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

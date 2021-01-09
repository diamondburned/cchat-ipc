// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
	cchat__core "cchat/core"
)

type SetConfigurationReturns struct {
	_tab flatbuffers.Table
}

func GetRootAsSetConfigurationReturns(buf []byte, offset flatbuffers.UOffsetT) *SetConfigurationReturns {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SetConfigurationReturns{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSetConfigurationReturns(buf []byte, offset flatbuffers.UOffsetT) *SetConfigurationReturns {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SetConfigurationReturns{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SetConfigurationReturns) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SetConfigurationReturns) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SetConfigurationReturns) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *SetConfigurationReturns) Err(obj *cchat__core.Error) *cchat__core.Error {
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

func SetConfigurationReturnsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SetConfigurationReturnsAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func SetConfigurationReturnsAddErr(builder *flatbuffers.Builder, err flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(err), 0)
}
func SetConfigurationReturnsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

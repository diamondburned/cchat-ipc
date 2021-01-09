// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
	cchat__core "cchat/core"
)

type RestoreSessionParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsRestoreSessionParameters(buf []byte, offset flatbuffers.UOffsetT) *RestoreSessionParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RestoreSessionParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRestoreSessionParameters(buf []byte, offset flatbuffers.UOffsetT) *RestoreSessionParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RestoreSessionParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RestoreSessionParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RestoreSessionParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RestoreSessionParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *RestoreSessionParameters) MapStringString(obj *cchat__core.StringPair) *cchat__core.StringPair {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__core.StringPair)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func RestoreSessionParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RestoreSessionParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func RestoreSessionParametersAddMapStringString(builder *flatbuffers.Builder, mapStringString flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(mapStringString), 0)
}
func RestoreSessionParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
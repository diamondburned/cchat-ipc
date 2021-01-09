// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type UpdateServerParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsUpdateServerParameters(buf []byte, offset flatbuffers.UOffsetT) *UpdateServerParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UpdateServerParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUpdateServerParameters(buf []byte, offset flatbuffers.UOffsetT) *UpdateServerParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UpdateServerParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UpdateServerParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UpdateServerParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UpdateServerParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *UpdateServerParameters) ServerUpdate(obj *ServerUpdate) *ServerUpdate {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ServerUpdate)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func UpdateServerParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func UpdateServerParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func UpdateServerParametersAddServerUpdate(builder *flatbuffers.Builder, serverUpdate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(serverUpdate), 0)
}
func UpdateServerParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

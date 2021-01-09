// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type EditParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsEditParameters(buf []byte, offset flatbuffers.UOffsetT) *EditParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EditParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEditParameters(buf []byte, offset flatbuffers.UOffsetT) *EditParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EditParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EditParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EditParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *EditParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *EditParameters) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *EditParameters) Content() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func EditParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func EditParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func EditParametersAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(id), 0)
}
func EditParametersAddContent(builder *flatbuffers.Builder, content flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(content), 0)
}
func EditParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
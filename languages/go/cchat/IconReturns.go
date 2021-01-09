// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
	cchat__core "cchat/core"
)

type IconReturns struct {
	_tab flatbuffers.Table
}

func GetRootAsIconReturns(buf []byte, offset flatbuffers.UOffsetT) *IconReturns {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IconReturns{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIconReturns(buf []byte, offset flatbuffers.UOffsetT) *IconReturns {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IconReturns{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IconReturns) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IconReturns) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IconReturns) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *IconReturns) Err(obj *cchat__core.Error) *cchat__core.Error {
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

func (rcv *IconReturns) StopHandle(obj *cchat__call.StopHandle) *cchat__call.StopHandle {
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

func IconReturnsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func IconReturnsAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func IconReturnsAddErr(builder *flatbuffers.Builder, err flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(err), 0)
}
func IconReturnsAddStopHandle(builder *flatbuffers.Builder, stopHandle flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(stopHandle), 0)
}
func IconReturnsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

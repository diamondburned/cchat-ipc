// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SendableMessage struct {
	_tab flatbuffers.Table
}

func GetRootAsSendableMessage(buf []byte, offset flatbuffers.UOffsetT) *SendableMessage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SendableMessage{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSendableMessage(buf []byte, offset flatbuffers.UOffsetT) *SendableMessage {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SendableMessage{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SendableMessage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SendableMessage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SendableMessage) Content() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SendableMessage) Noncer(obj *Noncer) *Noncer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Noncer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SendableMessage) Replier(obj *Replier) *Replier {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Replier)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SendableMessage) Attacher(obj *Attacher) *Attacher {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Attacher)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SendableMessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func SendableMessageAddContent(builder *flatbuffers.Builder, content flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(content), 0)
}
func SendableMessageAddNoncer(builder *flatbuffers.Builder, noncer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(noncer), 0)
}
func SendableMessageAddReplier(builder *flatbuffers.Builder, replier flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(replier), 0)
}
func SendableMessageAddAttacher(builder *flatbuffers.Builder, attacher flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(attacher), 0)
}
func SendableMessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Sender struct {
	_tab flatbuffers.Table
}

func GetRootAsSender(buf []byte, offset flatbuffers.UOffsetT) *Sender {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Sender{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSender(buf []byte, offset flatbuffers.UOffsetT) *Sender {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Sender{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Sender) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Sender) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Sender) CanAttach() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Sender) MutateCanAttach(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *Sender) Completer(obj *Completer) *Completer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Completer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SenderStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SenderAddCanAttach(builder *flatbuffers.Builder, canAttach bool) {
	builder.PrependBoolSlot(0, canAttach, false)
}
func SenderAddCompleter(builder *flatbuffers.Builder, completer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(completer), 0)
}
func SenderEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

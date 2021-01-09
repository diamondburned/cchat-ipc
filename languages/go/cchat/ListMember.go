// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__text "cchat/text"
)

type ListMember struct {
	_tab flatbuffers.Table
}

func GetRootAsListMember(buf []byte, offset flatbuffers.UOffsetT) *ListMember {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ListMember{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsListMember(buf []byte, offset flatbuffers.UOffsetT) *ListMember {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ListMember{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ListMember) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ListMember) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ListMember) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ListMember) Name(obj *cchat__text.Rich) *cchat__text.Rich {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__text.Rich)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ListMember) Status() Status {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Status(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ListMember) MutateStatus(n Status) bool {
	return rcv._tab.MutateByteSlot(8, byte(n))
}

func (rcv *ListMember) Secondary(obj *cchat__text.Rich) *cchat__text.Rich {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__text.Rich)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ListMember) Iconer(obj *Iconer) *Iconer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Iconer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ListMemberStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ListMemberAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func ListMemberAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func ListMemberAddStatus(builder *flatbuffers.Builder, status Status) {
	builder.PrependByteSlot(2, byte(status), 0)
}
func ListMemberAddSecondary(builder *flatbuffers.Builder, secondary flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(secondary), 0)
}
func ListMemberAddIconer(builder *flatbuffers.Builder, iconer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(iconer), 0)
}
func ListMemberEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
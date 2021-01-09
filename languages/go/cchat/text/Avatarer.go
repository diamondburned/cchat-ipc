// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package text

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Avatarer struct {
	_tab flatbuffers.Table
}

func GetRootAsAvatarer(buf []byte, offset flatbuffers.UOffsetT) *Avatarer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Avatarer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAvatarer(buf []byte, offset flatbuffers.UOffsetT) *Avatarer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Avatarer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Avatarer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Avatarer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Avatarer) Avatar() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Avatarer) AvatarSize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Avatarer) MutateAvatarSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Avatarer) AvatarText() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AvatarerStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AvatarerAddAvatar(builder *flatbuffers.Builder, avatar flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(avatar), 0)
}
func AvatarerAddAvatarSize(builder *flatbuffers.Builder, avatarSize int32) {
	builder.PrependInt32Slot(1, avatarSize, 0)
}
func AvatarerAddAvatarText(builder *flatbuffers.Builder, avatarText flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(avatarText), 0)
}
func AvatarerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
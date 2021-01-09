// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Nicknamer struct {
	_tab flatbuffers.Table
}

func GetRootAsNicknamer(buf []byte, offset flatbuffers.UOffsetT) *Nicknamer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Nicknamer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNicknamer(buf []byte, offset flatbuffers.UOffsetT) *Nicknamer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Nicknamer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Nicknamer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Nicknamer) Table() flatbuffers.Table {
	return rcv._tab
}

func NicknamerStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func NicknamerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
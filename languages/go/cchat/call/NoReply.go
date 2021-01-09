// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package call

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NoReply struct {
	_tab flatbuffers.Table
}

func GetRootAsNoReply(buf []byte, offset flatbuffers.UOffsetT) *NoReply {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NoReply{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNoReply(buf []byte, offset flatbuffers.UOffsetT) *NoReply {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NoReply{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NoReply) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NoReply) Table() flatbuffers.Table {
	return rcv._tab
}

func NoReplyStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func NoReplyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

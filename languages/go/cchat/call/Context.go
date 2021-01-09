// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package call

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Context struct {
	_tab flatbuffers.Table
}

func GetRootAsContext(buf []byte, offset flatbuffers.UOffsetT) *Context {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Context{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsContext(buf []byte, offset flatbuffers.UOffsetT) *Context {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Context{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Context) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Context) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Context) CtxId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Context) MutateCtxId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func ContextStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ContextAddCtxId(builder *flatbuffers.Builder, ctxId uint64) {
	builder.PrependUint64Slot(0, ctxId, 0)
}
func ContextEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
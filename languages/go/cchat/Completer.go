// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Completer struct {
	_tab flatbuffers.Table
}

func GetRootAsCompleter(buf []byte, offset flatbuffers.UOffsetT) *Completer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Completer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCompleter(buf []byte, offset flatbuffers.UOffsetT) *Completer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Completer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Completer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Completer) Table() flatbuffers.Table {
	return rcv._tab
}

func CompleterStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func CompleterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
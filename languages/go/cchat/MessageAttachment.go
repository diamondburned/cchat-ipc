// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__core "cchat/core"
)

type MessageAttachment struct {
	_tab flatbuffers.Table
}

func GetRootAsMessageAttachment(buf []byte, offset flatbuffers.UOffsetT) *MessageAttachment {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MessageAttachment{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMessageAttachment(buf []byte, offset flatbuffers.UOffsetT) *MessageAttachment {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MessageAttachment{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MessageAttachment) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MessageAttachment) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MessageAttachment) ReaderType() cchat__core.Reader {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return cchat__core.Reader(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *MessageAttachment) MutateReaderType(n cchat__core.Reader) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *MessageAttachment) Reader(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *MessageAttachment) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MessageAttachmentStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func MessageAttachmentAddReaderType(builder *flatbuffers.Builder, readerType cchat__core.Reader) {
	builder.PrependByteSlot(0, byte(readerType), 0)
}
func MessageAttachmentAddReader(builder *flatbuffers.Builder, reader flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(reader), 0)
}
func MessageAttachmentAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0)
}
func MessageAttachmentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

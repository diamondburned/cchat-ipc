// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package text

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Imager struct {
	_tab flatbuffers.Table
}

func GetRootAsImager(buf []byte, offset flatbuffers.UOffsetT) *Imager {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Imager{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsImager(buf []byte, offset flatbuffers.UOffsetT) *Imager {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Imager{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Imager) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Imager) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Imager) Image() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Imager) ImageSize(obj *ImagerImageSize) *ImagerImageSize {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ImagerImageSize)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Imager) ImageText() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ImagerStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ImagerAddImage(builder *flatbuffers.Builder, image flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(image), 0)
}
func ImagerAddImageSize(builder *flatbuffers.Builder, imageSize flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(imageSize), 0)
}
func ImagerAddImageText(builder *flatbuffers.Builder, imageText flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(imageText), 0)
}
func ImagerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

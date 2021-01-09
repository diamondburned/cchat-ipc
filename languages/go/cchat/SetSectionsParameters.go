// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type SetSectionsParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsSetSectionsParameters(buf []byte, offset flatbuffers.UOffsetT) *SetSectionsParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SetSectionsParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSetSectionsParameters(buf []byte, offset flatbuffers.UOffsetT) *SetSectionsParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SetSectionsParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SetSectionsParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SetSectionsParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SetSectionsParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__call.ID)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SetSectionsParameters) Sections(obj *MemberSection, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SetSectionsParameters) SectionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SetSectionsParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SetSectionsParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func SetSectionsParametersAddSections(builder *flatbuffers.Builder, sections flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(sections), 0)
}
func SetSectionsParametersStartSectionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SetSectionsParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

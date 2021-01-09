// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type ListMembersParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsListMembersParameters(buf []byte, offset flatbuffers.UOffsetT) *ListMembersParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ListMembersParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsListMembersParameters(buf []byte, offset flatbuffers.UOffsetT) *ListMembersParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ListMembersParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ListMembersParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ListMembersParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ListMembersParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func (rcv *ListMembersParameters) Ctx(obj *cchat__call.Context) *cchat__call.Context {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(cchat__call.Context)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ListMembersParameters) MemberListContainer(obj *MemberListContainer) *MemberListContainer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MemberListContainer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ListMembersParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ListMembersParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func ListMembersParametersAddCtx(builder *flatbuffers.Builder, ctx flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(ctx), 0)
}
func ListMembersParametersAddMemberListContainer(builder *flatbuffers.Builder, memberListContainer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(memberListContainer), 0)
}
func ListMembersParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

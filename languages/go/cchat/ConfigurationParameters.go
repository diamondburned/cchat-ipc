// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cchat

import (
	flatbuffers "github.com/google/flatbuffers/go"

	cchat__call "cchat/call"
)

type ConfigurationParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsConfigurationParameters(buf []byte, offset flatbuffers.UOffsetT) *ConfigurationParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConfigurationParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConfigurationParameters(buf []byte, offset flatbuffers.UOffsetT) *ConfigurationParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ConfigurationParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ConfigurationParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConfigurationParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConfigurationParameters) Call(obj *cchat__call.ID) *cchat__call.ID {
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

func ConfigurationParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ConfigurationParametersAddCall(builder *flatbuffers.Builder, call flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(call), 0)
}
func ConfigurationParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

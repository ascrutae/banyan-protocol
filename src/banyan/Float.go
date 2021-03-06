// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Float struct {
	_tab flatbuffers.Table
}

func GetRootAsFloat(buf []byte, offset flatbuffers.UOffsetT) *Float {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Float{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFloat(buf []byte, offset flatbuffers.UOffsetT) *Float {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Float{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Float) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Float) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Float) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Float) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func FloatStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FloatAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(0, value, 0.0)
}
func FloatEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

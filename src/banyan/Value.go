// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Value struct {
	_tab flatbuffers.Table
}

func GetRootAsValue(buf []byte, offset flatbuffers.UOffsetT) *Value {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Value{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsValue(buf []byte, offset flatbuffers.UOffsetT) *Value {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Value{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Value) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Value) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Value) ValueType() DataType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return DataType(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Value) MutateValueType(n DataType) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *Value) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func ValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ValueAddValueType(builder *flatbuffers.Builder, valueType DataType) {
	builder.PrependByteSlot(0, byte(valueType), 0)
}
func ValueAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func ValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

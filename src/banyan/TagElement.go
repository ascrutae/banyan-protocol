// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TagElement struct {
	_tab flatbuffers.Table
}

func GetRootAsTagElement(buf []byte, offset flatbuffers.UOffsetT) *TagElement {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TagElement{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTagElement(buf []byte, offset flatbuffers.UOffsetT) *TagElement {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TagElement{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TagElement) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TagElement) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TagElement) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TagElement) ValueType() DataType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return DataType(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *TagElement) MutateValueType(n DataType) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *TagElement) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func TagElementStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TagElementAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func TagElementAddValueType(builder *flatbuffers.Builder, valueType DataType) {
	builder.PrependByteSlot(1, byte(valueType), 0)
}
func TagElementAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(value), 0)
}
func TagElementEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

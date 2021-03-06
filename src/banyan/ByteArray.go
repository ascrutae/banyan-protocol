// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ByteArray struct {
	_tab flatbuffers.Table
}

func GetRootAsByteArray(buf []byte, offset flatbuffers.UOffsetT) *ByteArray {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ByteArray{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsByteArray(buf []byte, offset flatbuffers.UOffsetT) *ByteArray {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ByteArray{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ByteArray) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ByteArray) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ByteArray) Value(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ByteArray) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ByteArray) MutateValue(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func ByteArrayStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ByteArrayAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ByteArrayStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ByteArrayEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

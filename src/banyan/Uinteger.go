// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Uinteger struct {
	_tab flatbuffers.Table
}

func GetRootAsUinteger(buf []byte, offset flatbuffers.UOffsetT) *Uinteger {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uinteger{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUinteger(buf []byte, offset flatbuffers.UOffsetT) *Uinteger {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Uinteger{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Uinteger) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uinteger) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Uinteger) Value() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uinteger) MutateValue(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func UintegerStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UintegerAddValue(builder *flatbuffers.Builder, value uint32) {
	builder.PrependUint32Slot(0, value, 0)
}
func UintegerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

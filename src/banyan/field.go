// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type field struct {
	_tab flatbuffers.Table
}

func GetRootAsfield(buf []byte, offset flatbuffers.UOffsetT) *field {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &field{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsfield(buf []byte, offset flatbuffers.UOffsetT) *field {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &field{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *field) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *field) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *field) Segmentid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *field) Statement() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *field) Serviceid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *field) Serviceinstanceid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *field) Endpointname() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *field) Endpointid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *field) Latency() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *field) MutateLatency(n int32) bool {
	return rcv._tab.MutateInt32Slot(16, n)
}

func (rcv *field) Error() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *field) MutateError(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *field) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func fieldStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func fieldAddSegmentid(builder *flatbuffers.Builder, segmentid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(segmentid), 0)
}
func fieldAddStatement(builder *flatbuffers.Builder, statement flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(statement), 0)
}
func fieldAddServiceid(builder *flatbuffers.Builder, serviceid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(serviceid), 0)
}
func fieldAddServiceinstanceid(builder *flatbuffers.Builder, serviceinstanceid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(serviceinstanceid), 0)
}
func fieldAddEndpointname(builder *flatbuffers.Builder, endpointname flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(endpointname), 0)
}
func fieldAddEndpointid(builder *flatbuffers.Builder, endpointid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(endpointid), 0)
}
func fieldAddLatency(builder *flatbuffers.Builder, latency int32) {
	builder.PrependInt32Slot(6, latency, 0)
}
func fieldAddError(builder *flatbuffers.Builder, error bool) {
	builder.PrependBoolSlot(7, error, false)
}
func fieldAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(version), 0)
}
func fieldEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
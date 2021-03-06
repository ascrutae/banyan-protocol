// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package banyan

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConditionExpression struct {
	_tab flatbuffers.Table
}

func GetRootAsConditionExpression(buf []byte, offset flatbuffers.UOffsetT) *ConditionExpression {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConditionExpression{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConditionExpression(buf []byte, offset flatbuffers.UOffsetT) *ConditionExpression {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ConditionExpression{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ConditionExpression) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConditionExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConditionExpression) Operator() Operator {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Operator(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ConditionExpression) MutateOperator(n Operator) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *ConditionExpression) Values(obj *Value, j int) bool {
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

func (rcv *ConditionExpression) ValuesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ConditionExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ConditionExpressionAddOperator(builder *flatbuffers.Builder, operator Operator) {
	builder.PrependInt8Slot(0, int8(operator), 0)
}
func ConditionExpressionAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(values), 0)
}
func ConditionExpressionStartValuesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ConditionExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

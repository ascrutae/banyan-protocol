package main

import flatbuffers "github.com/google/flatbuffers/go"
import banyanDB "banyan"

func main() {
	builder := flatbuffers.NewBuilder(0)
	traceIdString := builder.CreateString("xxx.xxx.xxx.1")
	segmentIdString := builder.CreateString("xxx.xxx.xxx.2")
	httpUrlTagKeyString := builder.CreateString("http.url")
	httpUrlTagValueString := builder.CreateString("http://localhost:8080/test")
	httpStatusTagKeyString := builder.CreateString("http.status")
	banyanDB.IntegerStart(builder)
	banyanDB.IntegerAddValue(builder, 200)
	status := banyanDB.IntegerEnd(builder)

	banyanDB.StringStart(builder)
	banyanDB.StringAddValue(builder, traceIdString)
	traceId := banyanDB.StringEnd(builder)

	banyanDB.LongStart(builder)
	banyanDB.LongAddValue(builder, 1613526708000)
	startTime := banyanDB.LongEnd(builder)

	// end time
	banyanDB.LongStart(builder)
	banyanDB.LongAddValue(builder, 1613526708001)
	endTime := banyanDB.LongEnd(builder)

	// segment id
	banyanDB.StringStart(builder)
	banyanDB.StringAddValue(builder, segmentIdString)
	segmentId := banyanDB.StringEnd(builder)
	// statement ...
	// service id ...
	// service instance id
	// endpoint name
	// endpoint id
	// latency
	// error
	// version

	// tag1
	banyanDB.TagElementStart(builder)
	banyanDB.TagElementAddKey(builder, httpUrlTagKeyString)
	banyanDB.TagElementAddValue(builder, httpUrlTagValueString)
	tag1 := banyanDB.TagElementEnd(builder)
	// tag2
	banyanDB.TagElementStart(builder)
	banyanDB.TagElementAddKey(builder, httpStatusTagKeyString)
	banyanDB.TagElementAddValueType(builder, banyanDB.DataTypeUinteger)
	banyanDB.TagElementAddValue(builder, status)
	tag2 := banyanDB.TagElementEnd(builder)

	// tags vector
	banyanDB.TagsStartTagVector(builder, 2)
	builder.PrependUOffsetT(tag1)
	builder.PrependUOffsetT(tag2)
	tagVectors := builder.EndVector(2)

	// tags
	banyanDB.ValueStart(builder)
	banyanDB.ValueAddValue(builder, tagVectors)
	tags := banyanDB.ValueEnd(builder)

	banyanDB.WriteRequestStartValuesVector(builder, 5)
	builder.PrependUOffsetT(traceId)
	builder.PrependUOffsetT(startTime)
	builder.PrependUOffsetT(endTime)
	builder.PrependUOffsetT(segmentId)
	builder.PrependUOffsetT(tags)
	values := builder.EndVector(5)
	//
	banyanDB.WriteRequestStart(builder)
	banyanDB.ValueAddValue(builder, values)
	writeRequest := banyanDB.WriteRequestEnd(builder)
	//
	builder.Finish(writeRequest)
	println(builder.FinishedBytes())
}

// Code generated by protoc-gen-gogo.
// source: combos/marshaler/types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	combos/marshaler/types.proto

It has these top-level messages:
	KnownTypes
	StandardLibrary
*/
package types

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import google_protobuf1 "github.com/maditya/protobuf/types"
import google_protobuf2 "github.com/maditya/protobuf/types"
import google_protobuf3 "github.com/maditya/protobuf/types"

import time "time"

import github_com_maditya_protobuf_types "github.com/maditya/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KnownTypes struct {
	Dur   *google_protobuf1.Duration    `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	Ts    *google_protobuf2.Timestamp   `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Dbl   *google_protobuf3.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt   *google_protobuf3.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64   *google_protobuf3.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64   *google_protobuf3.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32   *google_protobuf3.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32   *google_protobuf3.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool  *google_protobuf3.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str   *google_protobuf3.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes *google_protobuf3.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
}

func (m *KnownTypes) Reset()                    { *m = KnownTypes{} }
func (m *KnownTypes) String() string            { return proto.CompactTextString(m) }
func (*KnownTypes) ProtoMessage()               {}
func (*KnownTypes) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *KnownTypes) GetDur() *google_protobuf1.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *KnownTypes) GetTs() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *KnownTypes) GetDbl() *google_protobuf3.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *KnownTypes) GetFlt() *google_protobuf3.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *KnownTypes) GetI64() *google_protobuf3.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *KnownTypes) GetU64() *google_protobuf3.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *KnownTypes) GetI32() *google_protobuf3.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *KnownTypes) GetU32() *google_protobuf3.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *KnownTypes) GetBool() *google_protobuf3.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *KnownTypes) GetStr() *google_protobuf3.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *KnownTypes) GetBytes() *google_protobuf3.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type StandardLibrary struct {
	NullableTimestamp    *google_protobuf2.Timestamp   `protobuf:"bytes,1,opt,name=nullableTimestamp" json:"nullableTimestamp,omitempty"`
	NullableDuration     *google_protobuf1.Duration    `protobuf:"bytes,2,opt,name=nullableDuration" json:"nullableDuration,omitempty"`
	NullableStdTime      *time.Time                    `protobuf:"bytes,3,opt,name=nullableStdTime" json:"nullableStdTime,omitempty"`
	NullableStdDuration  *time.Duration                `protobuf:"bytes,4,opt,name=nullableStdDuration" json:"nullableStdDuration,omitempty"`
	StdTime              time.Time                     `protobuf:"bytes,5,opt,name=stdTime" json:"stdTime"`
	StdDuration          time.Duration                 `protobuf:"bytes,6,opt,name=stdDuration" json:"stdDuration"`
	NullableTimestamps   []*google_protobuf2.Timestamp `protobuf:"bytes,11,rep,name=nullableTimestamps" json:"nullableTimestamps,omitempty"`
	NullableDurations    []*google_protobuf1.Duration  `protobuf:"bytes,12,rep,name=nullableDurations" json:"nullableDurations,omitempty"`
	NullableStdTimes     []*time.Time                  `protobuf:"bytes,13,rep,name=nullableStdTimes" json:"nullableStdTimes,omitempty"`
	NullableStdDurations []*time.Duration              `protobuf:"bytes,14,rep,name=nullableStdDurations" json:"nullableStdDurations,omitempty"`
	StdTimes             []time.Time                   `protobuf:"bytes,15,rep,name=stdTimes" json:"stdTimes"`
	StdDurations         []time.Duration               `protobuf:"bytes,16,rep,name=stdDurations" json:"stdDurations"`
}

func (m *StandardLibrary) Reset()                    { *m = StandardLibrary{} }
func (m *StandardLibrary) String() string            { return proto.CompactTextString(m) }
func (*StandardLibrary) ProtoMessage()               {}
func (*StandardLibrary) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *StandardLibrary) GetNullableTimestamp() *google_protobuf2.Timestamp {
	if m != nil {
		return m.NullableTimestamp
	}
	return nil
}

func (m *StandardLibrary) GetNullableDuration() *google_protobuf1.Duration {
	if m != nil {
		return m.NullableDuration
	}
	return nil
}

func (m *StandardLibrary) GetNullableStdTime() *time.Time {
	if m != nil {
		return m.NullableStdTime
	}
	return nil
}

func (m *StandardLibrary) GetNullableStdDuration() *time.Duration {
	if m != nil {
		return m.NullableStdDuration
	}
	return nil
}

func (m *StandardLibrary) GetStdTime() time.Time {
	if m != nil {
		return m.StdTime
	}
	return time.Time{}
}

func (m *StandardLibrary) GetStdDuration() time.Duration {
	if m != nil {
		return m.StdDuration
	}
	return 0
}

func (m *StandardLibrary) GetNullableTimestamps() []*google_protobuf2.Timestamp {
	if m != nil {
		return m.NullableTimestamps
	}
	return nil
}

func (m *StandardLibrary) GetNullableDurations() []*google_protobuf1.Duration {
	if m != nil {
		return m.NullableDurations
	}
	return nil
}

func (m *StandardLibrary) GetNullableStdTimes() []*time.Time {
	if m != nil {
		return m.NullableStdTimes
	}
	return nil
}

func (m *StandardLibrary) GetNullableStdDurations() []*time.Duration {
	if m != nil {
		return m.NullableStdDurations
	}
	return nil
}

func (m *StandardLibrary) GetStdTimes() []time.Time {
	if m != nil {
		return m.StdTimes
	}
	return nil
}

func (m *StandardLibrary) GetStdDurations() []time.Duration {
	if m != nil {
		return m.StdDurations
	}
	return nil
}

func init() {
	proto.RegisterType((*KnownTypes)(nil), "types.KnownTypes")
	proto.RegisterType((*StandardLibrary)(nil), "types.StandardLibrary")
}
func (this *KnownTypes) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KnownTypes)
	if !ok {
		that2, ok := that.(KnownTypes)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *KnownTypes")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KnownTypes but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KnownTypes but is not nil && this == nil")
	}
	if !this.Dur.Equal(that1.Dur) {
		return fmt.Errorf("Dur this(%v) Not Equal that(%v)", this.Dur, that1.Dur)
	}
	if !this.Ts.Equal(that1.Ts) {
		return fmt.Errorf("Ts this(%v) Not Equal that(%v)", this.Ts, that1.Ts)
	}
	if !this.Dbl.Equal(that1.Dbl) {
		return fmt.Errorf("Dbl this(%v) Not Equal that(%v)", this.Dbl, that1.Dbl)
	}
	if !this.Flt.Equal(that1.Flt) {
		return fmt.Errorf("Flt this(%v) Not Equal that(%v)", this.Flt, that1.Flt)
	}
	if !this.I64.Equal(that1.I64) {
		return fmt.Errorf("I64 this(%v) Not Equal that(%v)", this.I64, that1.I64)
	}
	if !this.U64.Equal(that1.U64) {
		return fmt.Errorf("U64 this(%v) Not Equal that(%v)", this.U64, that1.U64)
	}
	if !this.I32.Equal(that1.I32) {
		return fmt.Errorf("I32 this(%v) Not Equal that(%v)", this.I32, that1.I32)
	}
	if !this.U32.Equal(that1.U32) {
		return fmt.Errorf("U32 this(%v) Not Equal that(%v)", this.U32, that1.U32)
	}
	if !this.Bool.Equal(that1.Bool) {
		return fmt.Errorf("Bool this(%v) Not Equal that(%v)", this.Bool, that1.Bool)
	}
	if !this.Str.Equal(that1.Str) {
		return fmt.Errorf("Str this(%v) Not Equal that(%v)", this.Str, that1.Str)
	}
	if !this.Bytes.Equal(that1.Bytes) {
		return fmt.Errorf("Bytes this(%v) Not Equal that(%v)", this.Bytes, that1.Bytes)
	}
	return nil
}
func (this *KnownTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KnownTypes)
	if !ok {
		that2, ok := that.(KnownTypes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Dur.Equal(that1.Dur) {
		return false
	}
	if !this.Ts.Equal(that1.Ts) {
		return false
	}
	if !this.Dbl.Equal(that1.Dbl) {
		return false
	}
	if !this.Flt.Equal(that1.Flt) {
		return false
	}
	if !this.I64.Equal(that1.I64) {
		return false
	}
	if !this.U64.Equal(that1.U64) {
		return false
	}
	if !this.I32.Equal(that1.I32) {
		return false
	}
	if !this.U32.Equal(that1.U32) {
		return false
	}
	if !this.Bool.Equal(that1.Bool) {
		return false
	}
	if !this.Str.Equal(that1.Str) {
		return false
	}
	if !this.Bytes.Equal(that1.Bytes) {
		return false
	}
	return true
}
func (this *StandardLibrary) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*StandardLibrary)
	if !ok {
		that2, ok := that.(StandardLibrary)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *StandardLibrary")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *StandardLibrary but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *StandardLibrary but is not nil && this == nil")
	}
	if !this.NullableTimestamp.Equal(that1.NullableTimestamp) {
		return fmt.Errorf("NullableTimestamp this(%v) Not Equal that(%v)", this.NullableTimestamp, that1.NullableTimestamp)
	}
	if !this.NullableDuration.Equal(that1.NullableDuration) {
		return fmt.Errorf("NullableDuration this(%v) Not Equal that(%v)", this.NullableDuration, that1.NullableDuration)
	}
	if that1.NullableStdTime == nil {
		if this.NullableStdTime != nil {
			return fmt.Errorf("this.NullableStdTime != nil && that1.NullableStdTime == nil")
		}
	} else if !this.NullableStdTime.Equal(*that1.NullableStdTime) {
		return fmt.Errorf("NullableStdTime this(%v) Not Equal that(%v)", this.NullableStdTime, that1.NullableStdTime)
	}
	if this.NullableStdDuration != nil && that1.NullableStdDuration != nil {
		if *this.NullableStdDuration != *that1.NullableStdDuration {
			return fmt.Errorf("NullableStdDuration this(%v) Not Equal that(%v)", *this.NullableStdDuration, *that1.NullableStdDuration)
		}
	} else if this.NullableStdDuration != nil {
		return fmt.Errorf("this.NullableStdDuration == nil && that.NullableStdDuration != nil")
	} else if that1.NullableStdDuration != nil {
		return fmt.Errorf("NullableStdDuration this(%v) Not Equal that(%v)", this.NullableStdDuration, that1.NullableStdDuration)
	}
	if !this.StdTime.Equal(that1.StdTime) {
		return fmt.Errorf("StdTime this(%v) Not Equal that(%v)", this.StdTime, that1.StdTime)
	}
	if this.StdDuration != that1.StdDuration {
		return fmt.Errorf("StdDuration this(%v) Not Equal that(%v)", this.StdDuration, that1.StdDuration)
	}
	if len(this.NullableTimestamps) != len(that1.NullableTimestamps) {
		return fmt.Errorf("NullableTimestamps this(%v) Not Equal that(%v)", len(this.NullableTimestamps), len(that1.NullableTimestamps))
	}
	for i := range this.NullableTimestamps {
		if !this.NullableTimestamps[i].Equal(that1.NullableTimestamps[i]) {
			return fmt.Errorf("NullableTimestamps this[%v](%v) Not Equal that[%v](%v)", i, this.NullableTimestamps[i], i, that1.NullableTimestamps[i])
		}
	}
	if len(this.NullableDurations) != len(that1.NullableDurations) {
		return fmt.Errorf("NullableDurations this(%v) Not Equal that(%v)", len(this.NullableDurations), len(that1.NullableDurations))
	}
	for i := range this.NullableDurations {
		if !this.NullableDurations[i].Equal(that1.NullableDurations[i]) {
			return fmt.Errorf("NullableDurations this[%v](%v) Not Equal that[%v](%v)", i, this.NullableDurations[i], i, that1.NullableDurations[i])
		}
	}
	if len(this.NullableStdTimes) != len(that1.NullableStdTimes) {
		return fmt.Errorf("NullableStdTimes this(%v) Not Equal that(%v)", len(this.NullableStdTimes), len(that1.NullableStdTimes))
	}
	for i := range this.NullableStdTimes {
		if !this.NullableStdTimes[i].Equal(*that1.NullableStdTimes[i]) {
			return fmt.Errorf("NullableStdTimes this[%v](%v) Not Equal that[%v](%v)", i, this.NullableStdTimes[i], i, that1.NullableStdTimes[i])
		}
	}
	if len(this.NullableStdDurations) != len(that1.NullableStdDurations) {
		return fmt.Errorf("NullableStdDurations this(%v) Not Equal that(%v)", len(this.NullableStdDurations), len(that1.NullableStdDurations))
	}
	for i := range this.NullableStdDurations {
		if *(this.NullableStdDurations[i]) != *(that1.NullableStdDurations[i]) {
			return fmt.Errorf("NullableStdDurations this[%v](%v) Not Equal that[%v](%v)", i, this.NullableStdDurations[i], i, that1.NullableStdDurations[i])
		}
	}
	if len(this.StdTimes) != len(that1.StdTimes) {
		return fmt.Errorf("StdTimes this(%v) Not Equal that(%v)", len(this.StdTimes), len(that1.StdTimes))
	}
	for i := range this.StdTimes {
		if !this.StdTimes[i].Equal(that1.StdTimes[i]) {
			return fmt.Errorf("StdTimes this[%v](%v) Not Equal that[%v](%v)", i, this.StdTimes[i], i, that1.StdTimes[i])
		}
	}
	if len(this.StdDurations) != len(that1.StdDurations) {
		return fmt.Errorf("StdDurations this(%v) Not Equal that(%v)", len(this.StdDurations), len(that1.StdDurations))
	}
	for i := range this.StdDurations {
		if this.StdDurations[i] != that1.StdDurations[i] {
			return fmt.Errorf("StdDurations this[%v](%v) Not Equal that[%v](%v)", i, this.StdDurations[i], i, that1.StdDurations[i])
		}
	}
	return nil
}
func (this *StandardLibrary) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*StandardLibrary)
	if !ok {
		that2, ok := that.(StandardLibrary)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.NullableTimestamp.Equal(that1.NullableTimestamp) {
		return false
	}
	if !this.NullableDuration.Equal(that1.NullableDuration) {
		return false
	}
	if that1.NullableStdTime == nil {
		if this.NullableStdTime != nil {
			return false
		}
	} else if !this.NullableStdTime.Equal(*that1.NullableStdTime) {
		return false
	}
	if this.NullableStdDuration != nil && that1.NullableStdDuration != nil {
		if *this.NullableStdDuration != *that1.NullableStdDuration {
			return false
		}
	} else if this.NullableStdDuration != nil {
		return false
	} else if that1.NullableStdDuration != nil {
		return false
	}
	if !this.StdTime.Equal(that1.StdTime) {
		return false
	}
	if this.StdDuration != that1.StdDuration {
		return false
	}
	if len(this.NullableTimestamps) != len(that1.NullableTimestamps) {
		return false
	}
	for i := range this.NullableTimestamps {
		if !this.NullableTimestamps[i].Equal(that1.NullableTimestamps[i]) {
			return false
		}
	}
	if len(this.NullableDurations) != len(that1.NullableDurations) {
		return false
	}
	for i := range this.NullableDurations {
		if !this.NullableDurations[i].Equal(that1.NullableDurations[i]) {
			return false
		}
	}
	if len(this.NullableStdTimes) != len(that1.NullableStdTimes) {
		return false
	}
	for i := range this.NullableStdTimes {
		if !this.NullableStdTimes[i].Equal(*that1.NullableStdTimes[i]) {
			return false
		}
	}
	if len(this.NullableStdDurations) != len(that1.NullableStdDurations) {
		return false
	}
	for i := range this.NullableStdDurations {
		if *(this.NullableStdDurations[i]) != *(that1.NullableStdDurations[i]) {
			return false
		}
	}
	if len(this.StdTimes) != len(that1.StdTimes) {
		return false
	}
	for i := range this.StdTimes {
		if !this.StdTimes[i].Equal(that1.StdTimes[i]) {
			return false
		}
	}
	if len(this.StdDurations) != len(that1.StdDurations) {
		return false
	}
	for i := range this.StdDurations {
		if this.StdDurations[i] != that1.StdDurations[i] {
			return false
		}
	}
	return true
}
func (m *KnownTypes) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KnownTypes) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dur != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTypes(data, i, uint64(m.Dur.Size()))
		n1, err := m.Dur.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Ts != nil {
		data[i] = 0x12
		i++
		i = encodeVarintTypes(data, i, uint64(m.Ts.Size()))
		n2, err := m.Ts.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Dbl != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintTypes(data, i, uint64(m.Dbl.Size()))
		n3, err := m.Dbl.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Flt != nil {
		data[i] = 0x22
		i++
		i = encodeVarintTypes(data, i, uint64(m.Flt.Size()))
		n4, err := m.Flt.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.I64 != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintTypes(data, i, uint64(m.I64.Size()))
		n5, err := m.I64.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.U64 != nil {
		data[i] = 0x32
		i++
		i = encodeVarintTypes(data, i, uint64(m.U64.Size()))
		n6, err := m.U64.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.I32 != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintTypes(data, i, uint64(m.I32.Size()))
		n7, err := m.I32.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.U32 != nil {
		data[i] = 0x42
		i++
		i = encodeVarintTypes(data, i, uint64(m.U32.Size()))
		n8, err := m.U32.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Bool != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintTypes(data, i, uint64(m.Bool.Size()))
		n9, err := m.Bool.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.Str != nil {
		data[i] = 0x52
		i++
		i = encodeVarintTypes(data, i, uint64(m.Str.Size()))
		n10, err := m.Str.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.Bytes != nil {
		data[i] = 0x5a
		i++
		i = encodeVarintTypes(data, i, uint64(m.Bytes.Size()))
		n11, err := m.Bytes.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	return i, nil
}

func (m *StandardLibrary) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StandardLibrary) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NullableTimestamp != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTypes(data, i, uint64(m.NullableTimestamp.Size()))
		n12, err := m.NullableTimestamp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	if m.NullableDuration != nil {
		data[i] = 0x12
		i++
		i = encodeVarintTypes(data, i, uint64(m.NullableDuration.Size()))
		n13, err := m.NullableDuration.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.NullableStdTime != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdTime(*m.NullableStdTime)))
		n14, err := github_com_maditya_protobuf_types.StdTimeMarshalTo(*m.NullableStdTime, data[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	if m.NullableStdDuration != nil {
		data[i] = 0x22
		i++
		i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdDuration(*m.NullableStdDuration)))
		n15, err := github_com_maditya_protobuf_types.StdDurationMarshalTo(*m.NullableStdDuration, data[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	data[i] = 0x2a
	i++
	i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdTime(m.StdTime)))
	n16, err := github_com_maditya_protobuf_types.StdTimeMarshalTo(m.StdTime, data[i:])
	if err != nil {
		return 0, err
	}
	i += n16
	data[i] = 0x32
	i++
	i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdDuration(m.StdDuration)))
	n17, err := github_com_maditya_protobuf_types.StdDurationMarshalTo(m.StdDuration, data[i:])
	if err != nil {
		return 0, err
	}
	i += n17
	if len(m.NullableTimestamps) > 0 {
		for _, msg := range m.NullableTimestamps {
			data[i] = 0x5a
			i++
			i = encodeVarintTypes(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.NullableDurations) > 0 {
		for _, msg := range m.NullableDurations {
			data[i] = 0x62
			i++
			i = encodeVarintTypes(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.NullableStdTimes) > 0 {
		for _, msg := range m.NullableStdTimes {
			data[i] = 0x6a
			i++
			i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdTime(*msg)))
			n, err := github_com_maditya_protobuf_types.StdTimeMarshalTo(*msg, data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.NullableStdDurations) > 0 {
		for _, msg := range m.NullableStdDurations {
			data[i] = 0x72
			i++
			i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdDuration(*msg)))
			n, err := github_com_maditya_protobuf_types.StdDurationMarshalTo(*msg, data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.StdTimes) > 0 {
		for _, msg := range m.StdTimes {
			data[i] = 0x7a
			i++
			i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdTime(msg)))
			n, err := github_com_maditya_protobuf_types.StdTimeMarshalTo(msg, data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.StdDurations) > 0 {
		for _, msg := range m.StdDurations {
			data[i] = 0x82
			i++
			data[i] = 0x1
			i++
			i = encodeVarintTypes(data, i, uint64(github_com_maditya_protobuf_types.SizeOfStdDuration(msg)))
			n, err := github_com_maditya_protobuf_types.StdDurationMarshalTo(msg, data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Types(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Types(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTypes(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedKnownTypes(r randyTypes, easy bool) *KnownTypes {
	this := &KnownTypes{}
	if r.Intn(10) != 0 {
		this.Dur = google_protobuf1.NewPopulatedDuration(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Ts = google_protobuf2.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Dbl = google_protobuf3.NewPopulatedDoubleValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Flt = google_protobuf3.NewPopulatedFloatValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.I64 = google_protobuf3.NewPopulatedInt64Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.U64 = google_protobuf3.NewPopulatedUInt64Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.I32 = google_protobuf3.NewPopulatedInt32Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.U32 = google_protobuf3.NewPopulatedUInt32Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Bool = google_protobuf3.NewPopulatedBoolValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Str = google_protobuf3.NewPopulatedStringValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Bytes = google_protobuf3.NewPopulatedBytesValue(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStandardLibrary(r randyTypes, easy bool) *StandardLibrary {
	this := &StandardLibrary{}
	if r.Intn(10) != 0 {
		this.NullableTimestamp = google_protobuf2.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.NullableDuration = google_protobuf1.NewPopulatedDuration(r, easy)
	}
	if r.Intn(10) != 0 {
		this.NullableStdTime = github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
	}
	if r.Intn(10) != 0 {
		this.NullableStdDuration = github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	v1 := github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
	this.StdTime = *v1
	v2 := github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.StdDuration = *v2
	if r.Intn(10) != 0 {
		v3 := r.Intn(5)
		this.NullableTimestamps = make([]*google_protobuf2.Timestamp, v3)
		for i := 0; i < v3; i++ {
			this.NullableTimestamps[i] = google_protobuf2.NewPopulatedTimestamp(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v4 := r.Intn(5)
		this.NullableDurations = make([]*google_protobuf1.Duration, v4)
		for i := 0; i < v4; i++ {
			this.NullableDurations[i] = google_protobuf1.NewPopulatedDuration(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v5 := r.Intn(5)
		this.NullableStdTimes = make([]*time.Time, v5)
		for i := 0; i < v5; i++ {
			this.NullableStdTimes[i] = github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v6 := r.Intn(5)
		this.NullableStdDurations = make([]*time.Duration, v6)
		for i := 0; i < v6; i++ {
			this.NullableStdDurations[i] = github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v7 := r.Intn(5)
		this.StdTimes = make([]time.Time, v7)
		for i := 0; i < v7; i++ {
			v8 := github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
			this.StdTimes[i] = *v8
		}
	}
	if r.Intn(10) != 0 {
		v9 := r.Intn(5)
		this.StdDurations = make([]time.Duration, v9)
		for i := 0; i < v9; i++ {
			v10 := github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
			this.StdDurations[i] = *v10
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypes(r randyTypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypes(r randyTypes) string {
	v11 := r.Intn(100)
	tmps := make([]rune, v11)
	for i := 0; i < v11; i++ {
		tmps[i] = randUTF8RuneTypes(r)
	}
	return string(tmps)
}
func randUnrecognizedTypes(r randyTypes, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldTypes(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldTypes(data []byte, r randyTypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateTypes(data, uint64(key))
		v12 := r.Int63()
		if r.Intn(2) == 0 {
			v12 *= -1
		}
		data = encodeVarintPopulateTypes(data, uint64(v12))
	case 1:
		data = encodeVarintPopulateTypes(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateTypes(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateTypes(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateTypes(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateTypes(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *KnownTypes) Size() (n int) {
	var l int
	_ = l
	if m.Dur != nil {
		l = m.Dur.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Ts != nil {
		l = m.Ts.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Dbl != nil {
		l = m.Dbl.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Flt != nil {
		l = m.Flt.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.I64 != nil {
		l = m.I64.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.U64 != nil {
		l = m.U64.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.I32 != nil {
		l = m.I32.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.U32 != nil {
		l = m.U32.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Bool != nil {
		l = m.Bool.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Str != nil {
		l = m.Str.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Bytes != nil {
		l = m.Bytes.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *StandardLibrary) Size() (n int) {
	var l int
	_ = l
	if m.NullableTimestamp != nil {
		l = m.NullableTimestamp.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.NullableDuration != nil {
		l = m.NullableDuration.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.NullableStdTime != nil {
		l = github_com_maditya_protobuf_types.SizeOfStdTime(*m.NullableStdTime)
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.NullableStdDuration != nil {
		l = github_com_maditya_protobuf_types.SizeOfStdDuration(*m.NullableStdDuration)
		n += 1 + l + sovTypes(uint64(l))
	}
	l = github_com_maditya_protobuf_types.SizeOfStdTime(m.StdTime)
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_maditya_protobuf_types.SizeOfStdDuration(m.StdDuration)
	n += 1 + l + sovTypes(uint64(l))
	if len(m.NullableTimestamps) > 0 {
		for _, e := range m.NullableTimestamps {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.NullableDurations) > 0 {
		for _, e := range m.NullableDurations {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.NullableStdTimes) > 0 {
		for _, e := range m.NullableStdTimes {
			l = github_com_maditya_protobuf_types.SizeOfStdTime(*e)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.NullableStdDurations) > 0 {
		for _, e := range m.NullableStdDurations {
			l = github_com_maditya_protobuf_types.SizeOfStdDuration(*e)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.StdTimes) > 0 {
		for _, e := range m.StdTimes {
			l = github_com_maditya_protobuf_types.SizeOfStdTime(e)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.StdDurations) > 0 {
		for _, e := range m.StdDurations {
			l = github_com_maditya_protobuf_types.SizeOfStdDuration(e)
			n += 2 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func init() { proto.RegisterFile("combos/marshaler/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x9b, 0x35, 0xfd, 0x83, 0x3b, 0xe8, 0x30, 0x1c, 0x4c, 0x99, 0x5c, 0xb4, 0xd3, 0x04,
	0x22, 0x85, 0xb6, 0xea, 0x11, 0xa1, 0x8a, 0x31, 0x18, 0xbb, 0xd0, 0x0e, 0xee, 0xce, 0x92, 0x66,
	0x91, 0x9c, 0xb8, 0xb2, 0x1d, 0x4d, 0xbd, 0x71, 0xe7, 0xc2, 0x05, 0x89, 0x47, 0xe0, 0x11, 0x38,
	0x72, 0xdc, 0x91, 0x27, 0x80, 0x2d, 0xbc, 0xc4, 0x8e, 0x28, 0xae, 0xd3, 0x96, 0xa5, 0x69, 0x6e,
	0xb5, 0xfd, 0xf9, 0x7e, 0xfb, 0x69, 0x7f, 0x3f, 0xb0, 0x7b, 0xca, 0x02, 0x9b, 0x89, 0x4e, 0x40,
	0xb8, 0x38, 0x23, 0xd4, 0xe5, 0x1d, 0x39, 0x9b, 0xba, 0xc2, 0x9a, 0x72, 0x26, 0x19, 0xac, 0xa8,
	0x43, 0xeb, 0x99, 0xe7, 0xcb, 0xb3, 0xc8, 0xb6, 0x4e, 0x59, 0xd0, 0x09, 0x88, 0xe3, 0xcb, 0x19,
	0xe9, 0x28, 0xc0, 0x8e, 0x26, 0x1d, 0x8f, 0x79, 0x4c, 0x1d, 0xd4, 0xa7, 0x79, 0xb0, 0x85, 0x3d,
	0xc6, 0x3c, 0xea, 0x2e, 0x29, 0x27, 0xe2, 0x44, 0xfa, 0x2c, 0xd4, 0xef, 0xed, 0x9b, 0xef, 0xd2,
	0x0f, 0x5c, 0x21, 0x49, 0x30, 0xcd, 0x2b, 0x38, 0xe7, 0x64, 0x3a, 0x75, 0xb9, 0x36, 0xdb, 0xfb,
	0x6a, 0x02, 0xf0, 0x2e, 0x64, 0xe7, 0xe1, 0x49, 0x62, 0x08, 0x9f, 0x80, 0xb2, 0x13, 0x71, 0x64,
	0x3c, 0x32, 0xf6, 0x1b, 0xdd, 0x07, 0xd6, 0x3c, 0x6c, 0xa5, 0x61, 0xeb, 0x95, 0xfe, 0xf6, 0x51,
	0x42, 0xc1, 0xc7, 0x60, 0x4b, 0x0a, 0xb4, 0xa5, 0xd8, 0x56, 0x86, 0x3d, 0x49, 0x4d, 0x46, 0x5b,
	0x52, 0x40, 0x0b, 0x94, 0x1d, 0x9b, 0xa2, 0xb2, 0x82, 0x77, 0xb3, 0xc5, 0x2c, 0xb2, 0xa9, 0xfb,
	0x91, 0xd0, 0xc8, 0x1d, 0x25, 0x20, 0x7c, 0x0a, 0xca, 0x13, 0x2a, 0x91, 0xa9, 0xf8, 0x87, 0x19,
	0xfe, 0x35, 0x65, 0x44, 0x6a, 0x7c, 0x42, 0x65, 0x82, 0xfb, 0x83, 0x3e, 0xaa, 0xe4, 0xe0, 0x6f,
	0x43, 0x39, 0xe8, 0x6b, 0xdc, 0x1f, 0xf4, 0x13, 0x9b, 0x68, 0xd0, 0x47, 0xd5, 0x1c, 0x9b, 0x0f,
	0xab, 0x7c, 0x34, 0xe8, 0xab, 0xfa, 0x5e, 0x17, 0xd5, 0xf2, 0xeb, 0x7b, 0xdd, 0xb4, 0xbe, 0xd7,
	0x55, 0xf5, 0xbd, 0x2e, 0xaa, 0x6f, 0xa8, 0x5f, 0xf0, 0x91, 0xe2, 0x4d, 0x9b, 0x31, 0x8a, 0x6e,
	0xe5, 0xfc, 0x95, 0x43, 0xc6, 0xe8, 0x1c, 0x57, 0x5c, 0xd2, 0x2f, 0x24, 0x47, 0x20, 0xa7, 0x7f,
	0x2c, 0xb9, 0x1f, 0x7a, 0xba, 0x5f, 0x48, 0x0e, 0x9f, 0x83, 0x8a, 0x3d, 0x93, 0xae, 0x40, 0x8d,
	0x9c, 0x1f, 0x30, 0x4c, 0x5e, 0xe7, 0x81, 0x39, 0xb9, 0xf7, 0xb9, 0x06, 0x9a, 0x63, 0x49, 0x42,
	0x87, 0x70, 0xe7, 0xd8, 0xb7, 0x39, 0xe1, 0x33, 0xf8, 0x06, 0xdc, 0x0d, 0x23, 0x4a, 0x89, 0x4d,
	0xdd, 0xc5, 0x70, 0xf5, 0xaa, 0x6c, 0x1a, 0x7f, 0x36, 0x04, 0x0f, 0xc0, 0x4e, 0x7a, 0x99, 0xae,
	0x94, 0xde, 0xa3, 0x0d, 0x3b, 0x97, 0x89, 0xc0, 0x23, 0xd0, 0x4c, 0xef, 0xc6, 0xd2, 0x49, 0xea,
	0xf5, 0x82, 0x6d, 0xd0, 0x19, 0x9a, 0x5f, 0xfe, 0xb4, 0x8d, 0xd1, 0xcd, 0x20, 0x7c, 0x0f, 0xee,
	0xad, 0x5c, 0x2d, 0xac, 0xcc, 0x02, 0xab, 0xa1, 0xf9, 0x2d, 0xa9, 0x5b, 0x97, 0x85, 0x2f, 0x40,
	0x4d, 0x68, 0xad, 0x4a, 0xa1, 0x56, 0xfd, 0xe2, 0x77, 0xbb, 0xa4, 0xd4, 0xd2, 0x10, 0x3c, 0x00,
	0x0d, 0xb1, 0xa2, 0x52, 0x2d, 0x52, 0x51, 0x15, 0x4a, 0x67, 0x35, 0x07, 0x8f, 0x00, 0xcc, 0x4c,
	0x20, 0x59, 0x85, 0x72, 0xc1, 0xdc, 0xd6, 0xa4, 0xe0, 0xe1, 0x72, 0x05, 0xd2, 0x7e, 0x81, 0xb6,
	0x55, 0xd5, 0x86, 0xc9, 0x65, 0x33, 0xf0, 0x78, 0xb9, 0x01, 0x7a, 0x02, 0x02, 0xdd, 0x2e, 0x52,
	0xd2, 0xb3, 0xcb, 0x24, 0xe1, 0x18, 0xdc, 0x5f, 0x33, 0x00, 0x81, 0xee, 0x14, 0x98, 0xe9, 0xe9,
	0xad, 0x0d, 0xc3, 0x97, 0xa0, 0x2e, 0x52, 0xb5, 0x66, 0xa1, 0xda, 0x72, 0x7e, 0x8b, 0x14, 0x3c,
	0x04, 0xdb, 0x62, 0x55, 0x67, 0xa7, 0x48, 0x67, 0x39, 0xc1, 0xff, 0x82, 0xc3, 0xfd, 0xcb, 0x2b,
	0x6c, 0x5c, 0x5f, 0x61, 0xe3, 0x7b, 0x8c, 0x8d, 0x1f, 0x31, 0x36, 0x7e, 0xc6, 0xd8, 0xb8, 0x88,
	0xb1, 0xf1, 0x2b, 0xc6, 0xa5, 0xcb, 0x18, 0x1b, 0xd7, 0x31, 0x2e, 0x7d, 0xfa, 0x8b, 0x4b, 0x76,
	0x55, 0x95, 0xf6, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x7f, 0xca, 0x98, 0x90, 0x06, 0x00,
	0x00,
}

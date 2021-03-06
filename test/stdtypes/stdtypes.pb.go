// Code generated by protoc-gen-gogo.
// source: stdtypes.proto
// DO NOT EDIT!

/*
Package stdtypes is a generated protocol buffer package.

It is generated from these files:
	stdtypes.proto

It has these top-level messages:
	StdTypes
	RepStdTypes
	MapStdTypes
	OneofStdTypes
*/
package stdtypes

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import _ "github.com/maditya/protobuf/types"
import _ "github.com/maditya/protobuf/types"

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

type StdTypes struct {
	NullableTimestamp *time.Time     `protobuf:"bytes,1,opt,name=nullableTimestamp,stdtime" json:"nullableTimestamp,omitempty"`
	NullableDuration  *time.Duration `protobuf:"bytes,2,opt,name=nullableDuration,stdduration" json:"nullableDuration,omitempty"`
	Timestamp         time.Time      `protobuf:"bytes,3,opt,name=timestamp,stdtime" json:"timestamp"`
	Duration          time.Duration  `protobuf:"bytes,4,opt,name=duration,stdduration" json:"duration"`
}

func (m *StdTypes) Reset()                    { *m = StdTypes{} }
func (m *StdTypes) String() string            { return proto.CompactTextString(m) }
func (*StdTypes) ProtoMessage()               {}
func (*StdTypes) Descriptor() ([]byte, []int) { return fileDescriptorStdtypes, []int{0} }

func (m *StdTypes) GetNullableTimestamp() *time.Time {
	if m != nil {
		return m.NullableTimestamp
	}
	return nil
}

func (m *StdTypes) GetNullableDuration() *time.Duration {
	if m != nil {
		return m.NullableDuration
	}
	return nil
}

func (m *StdTypes) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *StdTypes) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

type RepStdTypes struct {
	NullableTimestamps []*time.Time     `protobuf:"bytes,1,rep,name=nullableTimestamps,stdtime" json:"nullableTimestamps,omitempty"`
	NullableDurations  []*time.Duration `protobuf:"bytes,2,rep,name=nullableDurations,stdduration" json:"nullableDurations,omitempty"`
	Timestamps         []time.Time      `protobuf:"bytes,3,rep,name=timestamps,stdtime" json:"timestamps"`
	Durations          []time.Duration  `protobuf:"bytes,4,rep,name=durations,stdduration" json:"durations"`
}

func (m *RepStdTypes) Reset()                    { *m = RepStdTypes{} }
func (m *RepStdTypes) String() string            { return proto.CompactTextString(m) }
func (*RepStdTypes) ProtoMessage()               {}
func (*RepStdTypes) Descriptor() ([]byte, []int) { return fileDescriptorStdtypes, []int{1} }

func (m *RepStdTypes) GetNullableTimestamps() []*time.Time {
	if m != nil {
		return m.NullableTimestamps
	}
	return nil
}

func (m *RepStdTypes) GetNullableDurations() []*time.Duration {
	if m != nil {
		return m.NullableDurations
	}
	return nil
}

func (m *RepStdTypes) GetTimestamps() []time.Time {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *RepStdTypes) GetDurations() []time.Duration {
	if m != nil {
		return m.Durations
	}
	return nil
}

type MapStdTypes struct {
	NullableTimestamp map[int32]*time.Time     `protobuf:"bytes,1,rep,name=nullableTimestamp,stdtime" json:"nullableTimestamp,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Timestamp         map[int32]time.Time      `protobuf:"bytes,2,rep,name=timestamp,stdtime" json:"timestamp" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	NullableDuration  map[int32]*time.Duration `protobuf:"bytes,3,rep,name=nullableDuration,stdduration" json:"nullableDuration,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Duration          map[int32]time.Duration  `protobuf:"bytes,4,rep,name=duration,stdduration" json:"duration" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MapStdTypes) Reset()                    { *m = MapStdTypes{} }
func (m *MapStdTypes) String() string            { return proto.CompactTextString(m) }
func (*MapStdTypes) ProtoMessage()               {}
func (*MapStdTypes) Descriptor() ([]byte, []int) { return fileDescriptorStdtypes, []int{2} }

func (m *MapStdTypes) GetNullableTimestamp() map[int32]*time.Time {
	if m != nil {
		return m.NullableTimestamp
	}
	return nil
}

func (m *MapStdTypes) GetTimestamp() map[int32]time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *MapStdTypes) GetNullableDuration() map[int32]*time.Duration {
	if m != nil {
		return m.NullableDuration
	}
	return nil
}

func (m *MapStdTypes) GetDuration() map[int32]time.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type OneofStdTypes struct {
	// Types that are valid to be assigned to OneOfStdTimes:
	//	*OneofStdTypes_Timestamp
	//	*OneofStdTypes_Duration
	OneOfStdTimes isOneofStdTypes_OneOfStdTimes `protobuf_oneof:"OneOfStdTimes"`
}

func (m *OneofStdTypes) Reset()                    { *m = OneofStdTypes{} }
func (m *OneofStdTypes) String() string            { return proto.CompactTextString(m) }
func (*OneofStdTypes) ProtoMessage()               {}
func (*OneofStdTypes) Descriptor() ([]byte, []int) { return fileDescriptorStdtypes, []int{3} }

type isOneofStdTypes_OneOfStdTimes interface {
	isOneofStdTypes_OneOfStdTimes()
	Equal(interface{}) bool
	VerboseEqual(interface{}) error
	Size() int
}

type OneofStdTypes_Timestamp struct {
	Timestamp *time.Time `protobuf:"bytes,1,opt,name=timestamp,oneof,stdtime"`
}
type OneofStdTypes_Duration struct {
	Duration *time.Duration `protobuf:"bytes,2,opt,name=duration,oneof,stdduration"`
}

func (*OneofStdTypes_Timestamp) isOneofStdTypes_OneOfStdTimes() {}
func (*OneofStdTypes_Duration) isOneofStdTypes_OneOfStdTimes()  {}

func (m *OneofStdTypes) GetOneOfStdTimes() isOneofStdTypes_OneOfStdTimes {
	if m != nil {
		return m.OneOfStdTimes
	}
	return nil
}

func (m *OneofStdTypes) GetTimestamp() *time.Time {
	if x, ok := m.GetOneOfStdTimes().(*OneofStdTypes_Timestamp); ok {
		return x.Timestamp
	}
	return nil
}

func (m *OneofStdTypes) GetDuration() *time.Duration {
	if x, ok := m.GetOneOfStdTimes().(*OneofStdTypes_Duration); ok {
		return x.Duration
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*OneofStdTypes) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _OneofStdTypes_OneofMarshaler, _OneofStdTypes_OneofUnmarshaler, _OneofStdTypes_OneofSizer, []interface{}{
		(*OneofStdTypes_Timestamp)(nil),
		(*OneofStdTypes_Duration)(nil),
	}
}

func _OneofStdTypes_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*OneofStdTypes)
	// OneOfStdTimes
	switch x := m.OneOfStdTimes.(type) {
	case *OneofStdTypes_Timestamp:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		data, err := github_com_maditya_protobuf_types.StdTimeMarshal(*x.Timestamp)
		if err != nil {
			return err
		}
		if err := b.EncodeRawBytes(data); err != nil {
			return err
		}
	case *OneofStdTypes_Duration:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		data, err := github_com_maditya_protobuf_types.StdDurationMarshal(*x.Duration)
		if err != nil {
			return err
		}
		if err := b.EncodeRawBytes(data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("OneofStdTypes.OneOfStdTimes has unexpected type %T", x)
	}
	return nil
}

func _OneofStdTypes_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*OneofStdTypes)
	switch tag {
	case 1: // OneOfStdTimes.timestamp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		if err != nil {
			return true, err
		}
		c := new(time.Time)
		if err2 := github_com_maditya_protobuf_types.StdTimeUnmarshal(c, x); err2 != nil {
			return true, err
		}
		m.OneOfStdTimes = &OneofStdTypes_Timestamp{c}
		return true, err
	case 2: // OneOfStdTimes.duration
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		if err != nil {
			return true, err
		}
		c := new(time.Duration)
		if err2 := github_com_maditya_protobuf_types.StdDurationUnmarshal(c, x); err2 != nil {
			return true, err
		}
		m.OneOfStdTimes = &OneofStdTypes_Duration{c}
		return true, err
	default:
		return false, nil
	}
}

func _OneofStdTypes_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*OneofStdTypes)
	// OneOfStdTimes
	switch x := m.OneOfStdTimes.(type) {
	case *OneofStdTypes_Timestamp:
		s := github_com_maditya_protobuf_types.SizeOfStdTime(*x.Timestamp)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *OneofStdTypes_Duration:
		s := github_com_maditya_protobuf_types.SizeOfStdDuration(*x.Duration)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*StdTypes)(nil), "stdtypes.StdTypes")
	proto.RegisterType((*RepStdTypes)(nil), "stdtypes.RepStdTypes")
	proto.RegisterType((*MapStdTypes)(nil), "stdtypes.MapStdTypes")
	proto.RegisterType((*OneofStdTypes)(nil), "stdtypes.OneofStdTypes")
}
func (this *StdTypes) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*StdTypes)
	if !ok {
		that2, ok := that.(StdTypes)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *StdTypes")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *StdTypes but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *StdTypes but is not nil && this == nil")
	}
	if that1.NullableTimestamp == nil {
		if this.NullableTimestamp != nil {
			return fmt.Errorf("this.NullableTimestamp != nil && that1.NullableTimestamp == nil")
		}
	} else if !this.NullableTimestamp.Equal(*that1.NullableTimestamp) {
		return fmt.Errorf("NullableTimestamp this(%v) Not Equal that(%v)", this.NullableTimestamp, that1.NullableTimestamp)
	}
	if this.NullableDuration != nil && that1.NullableDuration != nil {
		if *this.NullableDuration != *that1.NullableDuration {
			return fmt.Errorf("NullableDuration this(%v) Not Equal that(%v)", *this.NullableDuration, *that1.NullableDuration)
		}
	} else if this.NullableDuration != nil {
		return fmt.Errorf("this.NullableDuration == nil && that.NullableDuration != nil")
	} else if that1.NullableDuration != nil {
		return fmt.Errorf("NullableDuration this(%v) Not Equal that(%v)", this.NullableDuration, that1.NullableDuration)
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", this.Timestamp, that1.Timestamp)
	}
	if this.Duration != that1.Duration {
		return fmt.Errorf("Duration this(%v) Not Equal that(%v)", this.Duration, that1.Duration)
	}
	return nil
}
func (this *StdTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*StdTypes)
	if !ok {
		that2, ok := that.(StdTypes)
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
	if that1.NullableTimestamp == nil {
		if this.NullableTimestamp != nil {
			return false
		}
	} else if !this.NullableTimestamp.Equal(*that1.NullableTimestamp) {
		return false
	}
	if this.NullableDuration != nil && that1.NullableDuration != nil {
		if *this.NullableDuration != *that1.NullableDuration {
			return false
		}
	} else if this.NullableDuration != nil {
		return false
	} else if that1.NullableDuration != nil {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	return true
}
func (this *RepStdTypes) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*RepStdTypes)
	if !ok {
		that2, ok := that.(RepStdTypes)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *RepStdTypes")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *RepStdTypes but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *RepStdTypes but is not nil && this == nil")
	}
	if len(this.NullableTimestamps) != len(that1.NullableTimestamps) {
		return fmt.Errorf("NullableTimestamps this(%v) Not Equal that(%v)", len(this.NullableTimestamps), len(that1.NullableTimestamps))
	}
	for i := range this.NullableTimestamps {
		if !this.NullableTimestamps[i].Equal(*that1.NullableTimestamps[i]) {
			return fmt.Errorf("NullableTimestamps this[%v](%v) Not Equal that[%v](%v)", i, this.NullableTimestamps[i], i, that1.NullableTimestamps[i])
		}
	}
	if len(this.NullableDurations) != len(that1.NullableDurations) {
		return fmt.Errorf("NullableDurations this(%v) Not Equal that(%v)", len(this.NullableDurations), len(that1.NullableDurations))
	}
	for i := range this.NullableDurations {
		if dthis, dthat := this.NullableDurations[i], that1.NullableDurations[i]; (dthis != nil && dthat != nil && *dthis != *dthat) || (dthis != nil && dthat == nil) || (dthis == nil && dthat != nil) {
			return fmt.Errorf("NullableDurations this[%v](%v) Not Equal that[%v](%v)", i, this.NullableDurations[i], i, that1.NullableDurations[i])
		}
	}
	if len(this.Timestamps) != len(that1.Timestamps) {
		return fmt.Errorf("Timestamps this(%v) Not Equal that(%v)", len(this.Timestamps), len(that1.Timestamps))
	}
	for i := range this.Timestamps {
		if !this.Timestamps[i].Equal(that1.Timestamps[i]) {
			return fmt.Errorf("Timestamps this[%v](%v) Not Equal that[%v](%v)", i, this.Timestamps[i], i, that1.Timestamps[i])
		}
	}
	if len(this.Durations) != len(that1.Durations) {
		return fmt.Errorf("Durations this(%v) Not Equal that(%v)", len(this.Durations), len(that1.Durations))
	}
	for i := range this.Durations {
		if this.Durations[i] != that1.Durations[i] {
			return fmt.Errorf("Durations this[%v](%v) Not Equal that[%v](%v)", i, this.Durations[i], i, that1.Durations[i])
		}
	}
	return nil
}
func (this *RepStdTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RepStdTypes)
	if !ok {
		that2, ok := that.(RepStdTypes)
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
	if len(this.NullableTimestamps) != len(that1.NullableTimestamps) {
		return false
	}
	for i := range this.NullableTimestamps {
		if !this.NullableTimestamps[i].Equal(*that1.NullableTimestamps[i]) {
			return false
		}
	}
	if len(this.NullableDurations) != len(that1.NullableDurations) {
		return false
	}
	for i := range this.NullableDurations {
		if dthis, dthat := this.NullableDurations[i], that1.NullableDurations[i]; (dthis != nil && dthat != nil && *dthis != *dthat) || (dthis != nil && dthat == nil) || (dthis == nil && dthat != nil) {
			return false
		}
	}
	if len(this.Timestamps) != len(that1.Timestamps) {
		return false
	}
	for i := range this.Timestamps {
		if !this.Timestamps[i].Equal(that1.Timestamps[i]) {
			return false
		}
	}
	if len(this.Durations) != len(that1.Durations) {
		return false
	}
	for i := range this.Durations {
		if this.Durations[i] != that1.Durations[i] {
			return false
		}
	}
	return true
}
func (this *MapStdTypes) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MapStdTypes)
	if !ok {
		that2, ok := that.(MapStdTypes)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MapStdTypes")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MapStdTypes but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MapStdTypes but is not nil && this == nil")
	}
	if len(this.NullableTimestamp) != len(that1.NullableTimestamp) {
		return fmt.Errorf("NullableTimestamp this(%v) Not Equal that(%v)", len(this.NullableTimestamp), len(that1.NullableTimestamp))
	}
	for i := range this.NullableTimestamp {
		if !this.NullableTimestamp[i].Equal(*that1.NullableTimestamp[i]) {
			return fmt.Errorf("NullableTimestamp this[%v](%v) Not Equal that[%v](%v)", i, this.NullableTimestamp[i], i, that1.NullableTimestamp[i])
		}
	}
	if len(this.Timestamp) != len(that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", len(this.Timestamp), len(that1.Timestamp))
	}
	for i := range this.Timestamp {
		if !this.Timestamp[i].Equal(that1.Timestamp[i]) {
			return fmt.Errorf("Timestamp this[%v](%v) Not Equal that[%v](%v)", i, this.Timestamp[i], i, that1.Timestamp[i])
		}
	}
	if len(this.NullableDuration) != len(that1.NullableDuration) {
		return fmt.Errorf("NullableDuration this(%v) Not Equal that(%v)", len(this.NullableDuration), len(that1.NullableDuration))
	}
	for i := range this.NullableDuration {
		if dthis, dthat := this.NullableDuration[i], that1.NullableDuration[i]; (dthis != nil && dthat != nil && *dthis != *dthat) || (dthis != nil && dthat == nil) || (dthis == nil && dthat != nil) {
			return fmt.Errorf("NullableDuration this[%v](%v) Not Equal that[%v](%v)", i, this.NullableDuration[i], i, that1.NullableDuration[i])
		}
	}
	if len(this.Duration) != len(that1.Duration) {
		return fmt.Errorf("Duration this(%v) Not Equal that(%v)", len(this.Duration), len(that1.Duration))
	}
	for i := range this.Duration {
		if this.Duration[i] != that1.Duration[i] {
			return fmt.Errorf("Duration this[%v](%v) Not Equal that[%v](%v)", i, this.Duration[i], i, that1.Duration[i])
		}
	}
	return nil
}
func (this *MapStdTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MapStdTypes)
	if !ok {
		that2, ok := that.(MapStdTypes)
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
	if len(this.NullableTimestamp) != len(that1.NullableTimestamp) {
		return false
	}
	for i := range this.NullableTimestamp {
		if !this.NullableTimestamp[i].Equal(*that1.NullableTimestamp[i]) {
			return false
		}
	}
	if len(this.Timestamp) != len(that1.Timestamp) {
		return false
	}
	for i := range this.Timestamp {
		if !this.Timestamp[i].Equal(that1.Timestamp[i]) {
			return false
		}
	}
	if len(this.NullableDuration) != len(that1.NullableDuration) {
		return false
	}
	for i := range this.NullableDuration {
		if dthis, dthat := this.NullableDuration[i], that1.NullableDuration[i]; (dthis != nil && dthat != nil && *dthis != *dthat) || (dthis != nil && dthat == nil) || (dthis == nil && dthat != nil) {
			return false
		}
	}
	if len(this.Duration) != len(that1.Duration) {
		return false
	}
	for i := range this.Duration {
		if this.Duration[i] != that1.Duration[i] {
			return false
		}
	}
	return true
}
func (this *OneofStdTypes) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*OneofStdTypes)
	if !ok {
		that2, ok := that.(OneofStdTypes)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *OneofStdTypes")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *OneofStdTypes but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *OneofStdTypes but is not nil && this == nil")
	}
	if that1.OneOfStdTimes == nil {
		if this.OneOfStdTimes != nil {
			return fmt.Errorf("this.OneOfStdTimes != nil && that1.OneOfStdTimes == nil")
		}
	} else if this.OneOfStdTimes == nil {
		return fmt.Errorf("this.OneOfStdTimes == nil && that1.OneOfStdTimes != nil")
	} else if err := this.OneOfStdTimes.VerboseEqual(that1.OneOfStdTimes); err != nil {
		return err
	}
	return nil
}
func (this *OneofStdTypes_Timestamp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*OneofStdTypes_Timestamp)
	if !ok {
		that2, ok := that.(OneofStdTypes_Timestamp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *OneofStdTypes_Timestamp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *OneofStdTypes_Timestamp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *OneofStdTypes_Timestamp but is not nil && this == nil")
	}
	if that1.Timestamp == nil {
		if this.Timestamp != nil {
			return fmt.Errorf("this.Timestamp != nil && that1.Timestamp == nil")
		}
	} else if !this.Timestamp.Equal(*that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", this.Timestamp, that1.Timestamp)
	}
	return nil
}
func (this *OneofStdTypes_Duration) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*OneofStdTypes_Duration)
	if !ok {
		that2, ok := that.(OneofStdTypes_Duration)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *OneofStdTypes_Duration")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *OneofStdTypes_Duration but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *OneofStdTypes_Duration but is not nil && this == nil")
	}
	if this.Duration != nil && that1.Duration != nil {
		if *this.Duration != *that1.Duration {
			return fmt.Errorf("Duration this(%v) Not Equal that(%v)", *this.Duration, *that1.Duration)
		}
	} else if this.Duration != nil {
		return fmt.Errorf("this.Duration == nil && that.Duration != nil")
	} else if that1.Duration != nil {
		return fmt.Errorf("Duration this(%v) Not Equal that(%v)", this.Duration, that1.Duration)
	}
	return nil
}
func (this *OneofStdTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*OneofStdTypes)
	if !ok {
		that2, ok := that.(OneofStdTypes)
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
	if that1.OneOfStdTimes == nil {
		if this.OneOfStdTimes != nil {
			return false
		}
	} else if this.OneOfStdTimes == nil {
		return false
	} else if !this.OneOfStdTimes.Equal(that1.OneOfStdTimes) {
		return false
	}
	return true
}
func (this *OneofStdTypes_Timestamp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*OneofStdTypes_Timestamp)
	if !ok {
		that2, ok := that.(OneofStdTypes_Timestamp)
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
	if that1.Timestamp == nil {
		if this.Timestamp != nil {
			return false
		}
	} else if !this.Timestamp.Equal(*that1.Timestamp) {
		return false
	}
	return true
}
func (this *OneofStdTypes_Duration) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*OneofStdTypes_Duration)
	if !ok {
		that2, ok := that.(OneofStdTypes_Duration)
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
	if this.Duration != nil && that1.Duration != nil {
		if *this.Duration != *that1.Duration {
			return false
		}
	} else if this.Duration != nil {
		return false
	} else if that1.Duration != nil {
		return false
	}
	return true
}
func NewPopulatedStdTypes(r randyStdtypes, easy bool) *StdTypes {
	this := &StdTypes{}
	if r.Intn(10) != 0 {
		this.NullableTimestamp = github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
	}
	if r.Intn(10) != 0 {
		this.NullableDuration = github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	v1 := github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
	this.Timestamp = *v1
	v2 := github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.Duration = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedRepStdTypes(r randyStdtypes, easy bool) *RepStdTypes {
	this := &RepStdTypes{}
	if r.Intn(10) != 0 {
		v3 := r.Intn(5)
		this.NullableTimestamps = make([]*time.Time, v3)
		for i := 0; i < v3; i++ {
			this.NullableTimestamps[i] = github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v4 := r.Intn(5)
		this.NullableDurations = make([]*time.Duration, v4)
		for i := 0; i < v4; i++ {
			this.NullableDurations[i] = github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v5 := r.Intn(5)
		this.Timestamps = make([]time.Time, v5)
		for i := 0; i < v5; i++ {
			v6 := github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
			this.Timestamps[i] = *v6
		}
	}
	if r.Intn(10) != 0 {
		v7 := r.Intn(5)
		this.Durations = make([]time.Duration, v7)
		for i := 0; i < v7; i++ {
			v8 := github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
			this.Durations[i] = *v8
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMapStdTypes(r randyStdtypes, easy bool) *MapStdTypes {
	this := &MapStdTypes{}
	if r.Intn(10) != 0 {
		v9 := r.Intn(10)
		this.NullableTimestamp = make(map[int32]*time.Time)
		for i := 0; i < v9; i++ {
			this.NullableTimestamp[int32(r.Int31())] = github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v10 := r.Intn(10)
		this.Timestamp = make(map[int32]time.Time)
		for i := 0; i < v10; i++ {
			this.Timestamp[int32(r.Int31())] = *github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v11 := r.Intn(10)
		this.NullableDuration = make(map[int32]*time.Duration)
		for i := 0; i < v11; i++ {
			this.NullableDuration[int32(r.Int31())] = github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v12 := r.Intn(10)
		this.Duration = make(map[int32]time.Duration)
		for i := 0; i < v12; i++ {
			this.Duration[int32(r.Int31())] = *github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedOneofStdTypes(r randyStdtypes, easy bool) *OneofStdTypes {
	this := &OneofStdTypes{}
	oneofNumber_OneOfStdTimes := []int32{1, 2}[r.Intn(2)]
	switch oneofNumber_OneOfStdTimes {
	case 1:
		this.OneOfStdTimes = NewPopulatedOneofStdTypes_Timestamp(r, easy)
	case 2:
		this.OneOfStdTimes = NewPopulatedOneofStdTypes_Duration(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedOneofStdTypes_Timestamp(r randyStdtypes, easy bool) *OneofStdTypes_Timestamp {
	this := &OneofStdTypes_Timestamp{}
	this.Timestamp = github_com_maditya_protobuf_types.NewPopulatedStdTime(r, easy)
	return this
}
func NewPopulatedOneofStdTypes_Duration(r randyStdtypes, easy bool) *OneofStdTypes_Duration {
	this := &OneofStdTypes_Duration{}
	this.Duration = github_com_maditya_protobuf_types.NewPopulatedStdDuration(r, easy)
	return this
}

type randyStdtypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneStdtypes(r randyStdtypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringStdtypes(r randyStdtypes) string {
	v13 := r.Intn(100)
	tmps := make([]rune, v13)
	for i := 0; i < v13; i++ {
		tmps[i] = randUTF8RuneStdtypes(r)
	}
	return string(tmps)
}
func randUnrecognizedStdtypes(r randyStdtypes, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldStdtypes(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldStdtypes(data []byte, r randyStdtypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateStdtypes(data, uint64(key))
		v14 := r.Int63()
		if r.Intn(2) == 0 {
			v14 *= -1
		}
		data = encodeVarintPopulateStdtypes(data, uint64(v14))
	case 1:
		data = encodeVarintPopulateStdtypes(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateStdtypes(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateStdtypes(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateStdtypes(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateStdtypes(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *StdTypes) Size() (n int) {
	var l int
	_ = l
	if m.NullableTimestamp != nil {
		l = github_com_maditya_protobuf_types.SizeOfStdTime(*m.NullableTimestamp)
		n += 1 + l + sovStdtypes(uint64(l))
	}
	if m.NullableDuration != nil {
		l = github_com_maditya_protobuf_types.SizeOfStdDuration(*m.NullableDuration)
		n += 1 + l + sovStdtypes(uint64(l))
	}
	l = github_com_maditya_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovStdtypes(uint64(l))
	l = github_com_maditya_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovStdtypes(uint64(l))
	return n
}

func (m *RepStdTypes) Size() (n int) {
	var l int
	_ = l
	if len(m.NullableTimestamps) > 0 {
		for _, e := range m.NullableTimestamps {
			l = github_com_maditya_protobuf_types.SizeOfStdTime(*e)
			n += 1 + l + sovStdtypes(uint64(l))
		}
	}
	if len(m.NullableDurations) > 0 {
		for _, e := range m.NullableDurations {
			l = github_com_maditya_protobuf_types.SizeOfStdDuration(*e)
			n += 1 + l + sovStdtypes(uint64(l))
		}
	}
	if len(m.Timestamps) > 0 {
		for _, e := range m.Timestamps {
			l = github_com_maditya_protobuf_types.SizeOfStdTime(e)
			n += 1 + l + sovStdtypes(uint64(l))
		}
	}
	if len(m.Durations) > 0 {
		for _, e := range m.Durations {
			l = github_com_maditya_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovStdtypes(uint64(l))
		}
	}
	return n
}

func (m *MapStdTypes) Size() (n int) {
	var l int
	_ = l
	if len(m.NullableTimestamp) > 0 {
		for k, v := range m.NullableTimestamp {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = github_com_maditya_protobuf_types.SizeOfStdTime(*v)
				l += 1 + sovStdtypes(uint64(l))
			}
			mapEntrySize := 1 + sovStdtypes(uint64(k)) + l
			n += mapEntrySize + 1 + sovStdtypes(uint64(mapEntrySize))
		}
	}
	if len(m.Timestamp) > 0 {
		for k, v := range m.Timestamp {
			_ = k
			_ = v
			l = github_com_maditya_protobuf_types.SizeOfStdTime(v)
			mapEntrySize := 1 + sovStdtypes(uint64(k)) + 1 + l + sovStdtypes(uint64(l))
			n += mapEntrySize + 1 + sovStdtypes(uint64(mapEntrySize))
		}
	}
	if len(m.NullableDuration) > 0 {
		for k, v := range m.NullableDuration {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = github_com_maditya_protobuf_types.SizeOfStdDuration(*v)
				l += 1 + sovStdtypes(uint64(l))
			}
			mapEntrySize := 1 + sovStdtypes(uint64(k)) + l
			n += mapEntrySize + 1 + sovStdtypes(uint64(mapEntrySize))
		}
	}
	if len(m.Duration) > 0 {
		for k, v := range m.Duration {
			_ = k
			_ = v
			l = github_com_maditya_protobuf_types.SizeOfStdDuration(v)
			mapEntrySize := 1 + sovStdtypes(uint64(k)) + 1 + l + sovStdtypes(uint64(l))
			n += mapEntrySize + 1 + sovStdtypes(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *OneofStdTypes) Size() (n int) {
	var l int
	_ = l
	if m.OneOfStdTimes != nil {
		n += m.OneOfStdTimes.Size()
	}
	return n
}

func (m *OneofStdTypes_Timestamp) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != nil {
		l = github_com_maditya_protobuf_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovStdtypes(uint64(l))
	}
	return n
}
func (m *OneofStdTypes_Duration) Size() (n int) {
	var l int
	_ = l
	if m.Duration != nil {
		l = github_com_maditya_protobuf_types.SizeOfStdDuration(*m.Duration)
		n += 1 + l + sovStdtypes(uint64(l))
	}
	return n
}

func sovStdtypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStdtypes(x uint64) (n int) {
	return sovStdtypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func init() { proto.RegisterFile("stdtypes.proto", fileDescriptorStdtypes) }

var fileDescriptorStdtypes = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x1c, 0xc5, 0x7d, 0x4e, 0x82, 0xd2, 0x7f, 0xd4, 0x52, 0x4e, 0x02, 0x19, 0x0f, 0x97, 0x2a, 0x30,
	0x54, 0x02, 0x39, 0x15, 0x2c, 0x08, 0x09, 0x01, 0x56, 0x91, 0x8a, 0xa0, 0x2d, 0x0a, 0x15, 0x62,
	0x01, 0xd5, 0xc1, 0xae, 0x89, 0x70, 0x72, 0x51, 0x7c, 0x46, 0xca, 0xc6, 0x47, 0x60, 0x64, 0x65,
	0x63, 0x60, 0x87, 0x91, 0xb1, 0x23, 0x3b, 0x12, 0xb4, 0xe6, 0x4b, 0x74, 0xac, 0x7c, 0xce, 0xf9,
	0x9c, 0xf8, 0x52, 0x67, 0xc9, 0x66, 0xe7, 0xfe, 0xef, 0x77, 0x2f, 0xcf, 0xef, 0x0e, 0xd6, 0x42,
	0xe6, 0xb2, 0xf1, 0xd0, 0x0b, 0xad, 0xe1, 0x88, 0x32, 0x8a, 0xeb, 0xe2, 0xdd, 0xdc, 0xf2, 0x7b,
	0xec, 0x7d, 0xd4, 0xb5, 0xde, 0xd1, 0x7e, 0xbb, 0xef, 0xb8, 0x3d, 0x36, 0x76, 0xda, 0x7c, 0xa6,
	0x1b, 0x1d, 0xb5, 0x7d, 0xea, 0x53, 0xfe, 0xc2, 0x9f, 0x52, 0xad, 0x49, 0x7c, 0x4a, 0xfd, 0xc0,
	0x93, 0x53, 0x6e, 0x34, 0x72, 0x58, 0x8f, 0x0e, 0x26, 0xeb, 0xcd, 0xd9, 0x75, 0xd6, 0xeb, 0x7b,
	0x21, 0x73, 0xfa, 0xc3, 0x74, 0xa0, 0xf5, 0x5d, 0x87, 0xfa, 0x4b, 0xe6, 0x1e, 0x24, 0xfb, 0xe3,
	0x3d, 0xb8, 0x32, 0x88, 0x82, 0xc0, 0xe9, 0x06, 0xde, 0x81, 0x98, 0x33, 0xd0, 0x06, 0xda, 0x6c,
	0xdc, 0x31, 0xad, 0x94, 0x64, 0x09, 0x92, 0x95, 0x4d, 0xd8, 0xd5, 0xcf, 0xff, 0x9a, 0xa8, 0x53,
	0x94, 0xe2, 0x67, 0xb0, 0x2e, 0x7e, 0xdc, 0x9e, 0xf8, 0x32, 0x74, 0x8e, 0xbb, 0x5e, 0xc0, 0x89,
	0x01, 0xbb, 0xfa, 0x25, 0xa1, 0x15, 0x84, 0xd8, 0x86, 0x95, 0xcc, 0xbc, 0x51, 0x29, 0x35, 0x55,
	0x3f, 0xfe, 0xdb, 0xd4, 0xb8, 0x31, 0x29, 0xc3, 0x0f, 0xa1, 0x2e, 0x02, 0x32, 0xaa, 0x65, 0x46,
	0x38, 0x81, 0x9b, 0xc9, 0x44, 0xad, 0x1f, 0x3a, 0x34, 0x3a, 0xde, 0x30, 0x4b, 0xec, 0x05, 0xe0,
	0xc2, 0xdf, 0x0e, 0x0d, 0xb4, 0x51, 0x59, 0x28, 0x32, 0x85, 0x16, 0xef, 0xca, 0x6f, 0x20, 0x9c,
	0x84, 0x86, 0xce, 0x81, 0xa5, 0xa1, 0x15, 0x95, 0x78, 0x1b, 0x80, 0x49, 0x63, 0x95, 0x52, 0x63,
	0x32, 0xb6, 0x9c, 0x0e, 0x3f, 0x86, 0x15, 0x37, 0x33, 0x53, 0x2d, 0x33, 0x23, 0x83, 0x93, 0xaa,
	0xd6, 0x9f, 0x1a, 0x34, 0x76, 0x1d, 0x99, 0xdc, 0xa1, 0xba, 0x6b, 0x09, 0xfa, 0xb6, 0x95, 0x9d,
	0x90, 0x9c, 0xc2, 0xda, 0x9b, 0x1d, 0x7f, 0x32, 0x60, 0xa3, 0xf1, 0xfc, 0xf6, 0x3d, 0xcf, 0x17,
	0x26, 0x4d, 0xf0, 0xa6, 0x9a, 0x3c, 0x43, 0x54, 0x56, 0xe7, 0x8d, 0xa2, 0xcb, 0x69, 0x9c, 0xb7,
	0x2e, 0xb6, 0x2b, 0xa6, 0x27, 0x6e, 0xe7, 0xb4, 0xfb, 0xe9, 0x54, 0x33, 0x13, 0xec, 0x0d, 0x35,
	0x76, 0x1a, 0xa7, 0xe8, 0xa8, 0x79, 0x08, 0xd7, 0xd4, 0x51, 0xe1, 0x75, 0xa8, 0x7c, 0xf0, 0xc6,
	0xfc, 0x44, 0xd7, 0x3a, 0xc9, 0x23, 0xde, 0x82, 0xda, 0x47, 0x27, 0x88, 0xbc, 0xc9, 0xb1, 0xbc,
	0xa0, 0x19, 0x9d, 0x74, 0xf0, 0xbe, 0x7e, 0x0f, 0x99, 0xaf, 0x61, 0x6d, 0x49, 0xe4, 0xb7, 0x70,
	0x55, 0x99, 0x9b, 0x62, 0x83, 0xf6, 0xf4, 0x06, 0xf3, 0xfb, 0x98, 0xe7, 0xbf, 0x82, 0xd5, 0x65,
	0x70, 0x5b, 0x5f, 0x11, 0xac, 0xee, 0x0f, 0x3c, 0x7a, 0x94, 0xf5, 0xfb, 0x51, 0xbe, 0x7d, 0x0b,
	0xde, 0xa1, 0x3b, 0x5a, 0xbe, 0x71, 0x0f, 0x72, 0x95, 0x58, 0xec, 0xd6, 0xdc, 0xd1, 0x64, 0x0d,
	0xec, 0xcb, 0xdc, 0xd1, 0x3e, 0x77, 0x94, 0x30, 0xed, 0xcd, 0x93, 0x53, 0x82, 0xce, 0x4e, 0x09,
	0xfa, 0x16, 0x13, 0xf4, 0x33, 0x26, 0xe8, 0x57, 0x4c, 0xd0, 0x71, 0x4c, 0xb4, 0xdf, 0x31, 0xd1,
	0x4e, 0x62, 0x82, 0xce, 0x62, 0xa2, 0x7d, 0xfa, 0x4f, 0xb4, 0xee, 0x25, 0xce, 0xbf, 0x7b, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0x64, 0x01, 0xea, 0x52, 0xaa, 0x06, 0x00, 0x00,
}

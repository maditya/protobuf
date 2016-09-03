// Code generated by protoc-gen-gogo.
// source: combos/marshaler/types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	combos/marshaler/types.proto

It has these top-level messages:
	KnownTypes
*/
package types

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import google_protobuf1 "github.com/maditya/protobuf/types"
import google_protobuf2 "github.com/maditya/protobuf/types"
import google_protobuf3 "github.com/maditya/protobuf/types"

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
	// google.protobuf.Any an = 14;
	Dur *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	// google.protobuf.Struct st = 12;
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

func init() {
	proto.RegisterType((*KnownTypes)(nil), "types.KnownTypes")
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
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
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
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		data = encodeVarintPopulateTypes(data, uint64(v2))
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
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0xd0, 0x31, 0xae, 0xd3, 0x30,
	0x18, 0xc0, 0xf1, 0xba, 0x69, 0x1f, 0xe0, 0xb7, 0x65, 0x32, 0xe5, 0xc9, 0x20, 0xc4, 0x80, 0x40,
	0x24, 0x90, 0x44, 0x39, 0x40, 0x85, 0x90, 0x10, 0x5b, 0x79, 0xb0, 0xdb, 0x2f, 0x6e, 0x1a, 0xc9,
	0xc9, 0x17, 0xd9, 0x9f, 0x55, 0x75, 0xe3, 0x12, 0xdc, 0x81, 0x23, 0x30, 0x32, 0x32, 0x72, 0x04,
	0x08, 0x97, 0xe8, 0x88, 0xe2, 0xa4, 0x80, 0xa8, 0xc2, 0x16, 0xcb, 0xbf, 0xef, 0x9f, 0x2f, 0xa1,
	0x57, 0x37, 0x50, 0x4b, 0xb0, 0x71, 0x2d, 0x8c, 0xdd, 0x09, 0xad, 0x4c, 0x8c, 0x87, 0x56, 0xd9,
	0xa8, 0x35, 0x80, 0x10, 0x2e, 0xfd, 0x61, 0xf5, 0xbc, 0xac, 0x70, 0xe7, 0x64, 0x74, 0x03, 0x75,
	0x5c, 0x8b, 0xa2, 0xc2, 0x83, 0x88, 0x3d, 0x90, 0x6e, 0x1b, 0x97, 0x50, 0x82, 0x3f, 0xf8, 0xa7,
	0x61, 0x70, 0xc5, 0x4b, 0x80, 0x52, 0xab, 0x3f, 0xaa, 0x70, 0x46, 0x60, 0x05, 0xcd, 0x78, 0x7f,
	0xff, 0xdf, 0x7b, 0xac, 0x6a, 0x65, 0x51, 0xd4, 0xed, 0x54, 0x60, 0x6f, 0x44, 0xdb, 0x2a, 0x33,
	0x6e, 0xf6, 0xf0, 0xe3, 0x82, 0xd2, 0x37, 0x0d, 0xec, 0x9b, 0xeb, 0x7e, 0xc3, 0xf0, 0x29, 0x0d,
	0x0a, 0x67, 0x18, 0x79, 0x40, 0x1e, 0x5f, 0x26, 0x77, 0xa3, 0x61, 0x38, 0x3a, 0x0d, 0x47, 0x2f,
	0xc7, 0xb7, 0x6f, 0x7a, 0x15, 0x3e, 0xa1, 0x73, 0xb4, 0x6c, 0xee, 0xed, 0xea, 0xcc, 0x5e, 0x9f,
	0x36, 0xd9, 0xcc, 0xd1, 0x86, 0x11, 0x0d, 0x0a, 0xa9, 0x59, 0xe0, 0xf1, 0xd5, 0x79, 0x18, 0x9c,
	0xd4, 0xea, 0xbd, 0xd0, 0x4e, 0x6d, 0x7a, 0x18, 0x3e, 0xa3, 0xc1, 0x56, 0x23, 0x5b, 0x78, 0x7f,
	0xef, 0xcc, 0xbf, 0xd2, 0x20, 0x70, 0xe4, 0x5b, 0x8d, 0x3d, 0xaf, 0xf2, 0x8c, 0x2d, 0x27, 0xf8,
	0xeb, 0x06, 0xf3, 0x6c, 0xe4, 0x55, 0x9e, 0xf5, 0xdb, 0xb8, 0x3c, 0x63, 0x17, 0x13, 0xdb, 0xbc,
	0xfb, 0xdb, 0xbb, 0x3c, 0xf3, 0xf9, 0x34, 0x61, 0xb7, 0xa6, 0xf3, 0x69, 0x72, 0xca, 0xa7, 0x89,
	0xcf, 0xa7, 0x09, 0xbb, 0xfd, 0x9f, 0xfc, 0x6f, 0xef, 0xbc, 0x5f, 0x48, 0x00, 0xcd, 0xee, 0x4c,
	0xfc, 0xca, 0x35, 0x80, 0x1e, 0xb8, 0x77, 0x7d, 0xdf, 0xa2, 0x61, 0x74, 0xa2, 0xff, 0x16, 0x4d,
	0xd5, 0x94, 0x63, 0xdf, 0xa2, 0x09, 0x5f, 0xd0, 0xa5, 0x3c, 0xa0, 0xb2, 0xec, 0x72, 0xe2, 0x03,
	0xd6, 0xfd, 0xed, 0x30, 0x30, 0xc8, 0xf5, 0xa3, 0xe3, 0x0f, 0x4e, 0x3e, 0x75, 0x9c, 0x7c, 0xee,
	0x38, 0xf9, 0xd2, 0x71, 0xf2, 0xb5, 0xe3, 0xe4, 0x5b, 0xc7, 0x67, 0xdf, 0x3b, 0x4e, 0x8e, 0x1d,
	0x9f, 0x7d, 0xf8, 0xc9, 0x67, 0xf2, 0xc2, 0x17, 0xd2, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x26, 0x77, 0x51, 0xfe, 0x02, 0x00, 0x00,
}

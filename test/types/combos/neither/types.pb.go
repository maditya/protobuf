// Code generated by protoc-gen-gogo.
// source: combos/neither/types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	combos/neither/types.proto

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
import google_protobuf4 "github.com/maditya/protobuf/types"

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
	Dur   *google_protobuf1.Duration    `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	St    *google_protobuf2.Struct      `protobuf:"bytes,12,opt,name=st" json:"st,omitempty"`
	Ts    *google_protobuf3.Timestamp   `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Dbl   *google_protobuf4.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt   *google_protobuf4.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64   *google_protobuf4.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64   *google_protobuf4.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32   *google_protobuf4.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32   *google_protobuf4.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool  *google_protobuf4.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str   *google_protobuf4.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes *google_protobuf4.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
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

func (m *KnownTypes) GetSt() *google_protobuf2.Struct {
	if m != nil {
		return m.St
	}
	return nil
}

func (m *KnownTypes) GetTs() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *KnownTypes) GetDbl() *google_protobuf4.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *KnownTypes) GetFlt() *google_protobuf4.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *KnownTypes) GetI64() *google_protobuf4.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *KnownTypes) GetU64() *google_protobuf4.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *KnownTypes) GetI32() *google_protobuf4.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *KnownTypes) GetU32() *google_protobuf4.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *KnownTypes) GetBool() *google_protobuf4.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *KnownTypes) GetStr() *google_protobuf4.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *KnownTypes) GetBytes() *google_protobuf4.BytesValue {
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
	if !this.St.Equal(that1.St) {
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
func NewPopulatedKnownTypes(r randyTypes, easy bool) *KnownTypes {
	this := &KnownTypes{}
	if r.Intn(10) != 0 {
		this.Dur = google_protobuf1.NewPopulatedDuration(r, easy)
	}
	if r.Intn(10) == 0 {
		this.St = google_protobuf2.NewPopulatedStruct(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Ts = google_protobuf3.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Dbl = google_protobuf4.NewPopulatedDoubleValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Flt = google_protobuf4.NewPopulatedFloatValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.I64 = google_protobuf4.NewPopulatedInt64Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.U64 = google_protobuf4.NewPopulatedUInt64Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.I32 = google_protobuf4.NewPopulatedInt32Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.U32 = google_protobuf4.NewPopulatedUInt32Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Bool = google_protobuf4.NewPopulatedBoolValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Str = google_protobuf4.NewPopulatedStringValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Bytes = google_protobuf4.NewPopulatedBytesValue(r, easy)
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
	if m.St != nil {
		l = m.St.Size()
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

func init() { proto.RegisterFile("combos/neither/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0xd0, 0x4f, 0x8a, 0xd4, 0x40,
	0x14, 0x06, 0xf0, 0xfc, 0xe9, 0x1e, 0xb5, 0xc6, 0x55, 0x36, 0x96, 0x71, 0x28, 0x45, 0x04, 0x45,
	0x99, 0x04, 0x93, 0x90, 0x03, 0x34, 0x22, 0x88, 0xbb, 0xcc, 0xe8, 0x3e, 0xd5, 0x5d, 0x9d, 0x09,
	0x54, 0xf2, 0x42, 0xd5, 0x2b, 0x86, 0xd9, 0x79, 0x1c, 0x8f, 0xe0, 0x4a, 0x5c, 0xba, 0xf4, 0x08,
	0x1a, 0x2f, 0x31, 0x4b, 0x49, 0x25, 0xad, 0x32, 0x21, 0xb3, 0xeb, 0xe2, 0xfb, 0xbd, 0x8f, 0xaf,
	0x43, 0xc2, 0x2d, 0x34, 0x1c, 0x74, 0xdc, 0x8a, 0x1a, 0x2f, 0x84, 0x8a, 0xf1, 0xaa, 0x13, 0x3a,
	0xea, 0x14, 0x20, 0x04, 0x6b, 0xfb, 0x08, 0x4f, 0xab, 0x1a, 0x2f, 0x0c, 0x8f, 0xb6, 0xd0, 0xc4,
	0x15, 0x54, 0x10, 0xdb, 0x94, 0x9b, 0xbd, 0x7d, 0xd9, 0x87, 0xfd, 0x35, 0x5e, 0x85, 0xac, 0x02,
	0xa8, 0xa4, 0xf8, 0xa7, 0x76, 0x46, 0x95, 0x58, 0x43, 0x3b, 0xe5, 0x27, 0x37, 0x73, 0x8d, 0xca,
	0x6c, 0x71, 0x4a, 0x1f, 0xdf, 0x4c, 0xb1, 0x6e, 0x84, 0xc6, 0xb2, 0xe9, 0x96, 0xea, 0x2f, 0x55,
	0xd9, 0x75, 0x42, 0x4d, 0xa3, 0x9f, 0x7e, 0x5d, 0x11, 0xf2, 0xbe, 0x85, 0xcb, 0xf6, 0x7c, 0x18,
	0x1f, 0xbc, 0x22, 0xfe, 0xce, 0x28, 0xea, 0x3e, 0x71, 0x5f, 0x1c, 0x27, 0x0f, 0xa3, 0xf1, 0x38,
	0x3a, 0x1c, 0x47, 0x6f, 0xa6, 0x6d, 0xc5, 0xa0, 0x82, 0xe7, 0xc4, 0xd3, 0x48, 0xef, 0x5b, 0xfb,
	0x60, 0x66, 0xcf, 0xec, 0xce, 0xc2, 0xd3, 0x18, 0xbc, 0x24, 0x1e, 0x6a, 0xea, 0x59, 0x18, 0xce,
	0xe0, 0xf9, 0x61, 0x72, 0xe1, 0xa1, 0x0e, 0x22, 0xe2, 0xef, 0xb8, 0xa4, 0xbe, 0xc5, 0x27, 0xf3,
	0x05, 0x60, 0xb8, 0x14, 0x1f, 0x4b, 0x69, 0x44, 0x31, 0xc0, 0xe0, 0x94, 0xf8, 0x7b, 0x89, 0x74,
	0x65, 0xfd, 0xa3, 0x99, 0x7f, 0x2b, 0xa1, 0xc4, 0x89, 0xef, 0x25, 0x0e, 0xbc, 0xce, 0x33, 0xba,
	0x5e, 0xe0, 0xef, 0x5a, 0xcc, 0xb3, 0x89, 0xd7, 0x79, 0x36, 0xac, 0x31, 0x79, 0x46, 0x8f, 0x16,
	0xd6, 0x7c, 0xf8, 0xdf, 0x9b, 0x3c, 0xb3, 0xf5, 0x69, 0x42, 0xef, 0x2c, 0xd7, 0xa7, 0xc9, 0xa1,
	0x3e, 0x4d, 0x6c, 0x7d, 0x9a, 0xd0, 0xbb, 0xb7, 0xd4, 0xff, 0xf5, 0xc6, 0xfa, 0x15, 0x07, 0x90,
	0xf4, 0xde, 0xc2, 0xa7, 0xdc, 0x00, 0xc8, 0x91, 0x5b, 0x37, 0xf4, 0x6b, 0x54, 0x94, 0x2c, 0xf4,
	0x9f, 0xa1, 0xaa, 0xdb, 0x6a, 0xea, 0xd7, 0xa8, 0x82, 0xd7, 0x64, 0xcd, 0xaf, 0x50, 0x68, 0x7a,
	0xbc, 0xf0, 0x07, 0x36, 0x43, 0x3a, 0x1e, 0x8c, 0x72, 0xf3, 0xec, 0xfa, 0x17, 0x73, 0x3f, 0xf7,
	0xcc, 0xfd, 0xd2, 0x33, 0xf7, 0x5b, 0xcf, 0xdc, 0xef, 0x3d, 0x73, 0x7e, 0xf4, 0xcc, 0xf9, 0xd9,
	0x33, 0xf7, 0xba, 0x67, 0xce, 0xa7, 0xdf, 0xcc, 0xe1, 0x47, 0xb6, 0x21, 0xfd, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x20, 0x2c, 0x60, 0x5f, 0x40, 0x03, 0x00, 0x00,
}

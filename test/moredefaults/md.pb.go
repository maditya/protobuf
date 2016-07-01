// Code generated by protoc-gen-gogo.
// source: md.proto
// DO NOT EDIT!

/*
	Package moredefaults is a generated protocol buffer package.

	It is generated from these files:
		md.proto

	It has these top-level messages:
		MoreDefaultsB
		MoreDefaultsA
*/
package moredefaults

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import test "github.com/maditya/protobuf/test/example"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type MoreDefaultsB struct {
	Field1           *string `protobuf:"bytes,1,opt,name=Field1,json=field1" json:"Field1,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MoreDefaultsB) Reset()                    { *m = MoreDefaultsB{} }
func (m *MoreDefaultsB) String() string            { return proto.CompactTextString(m) }
func (*MoreDefaultsB) ProtoMessage()               {}
func (*MoreDefaultsB) Descriptor() ([]byte, []int) { return fileDescriptorMd, []int{0} }

func (m *MoreDefaultsB) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

type MoreDefaultsA struct {
	Field1           *int64         `protobuf:"varint,1,opt,name=Field1,json=field1,def=1234" json:"Field1,omitempty"`
	Field2           int64          `protobuf:"varint,2,opt,name=Field2,json=field2" json:"Field2"`
	B1               *MoreDefaultsB `protobuf:"bytes,3,opt,name=B1,json=b1" json:"B1,omitempty"`
	B2               MoreDefaultsB  `protobuf:"bytes,4,opt,name=B2,json=b2" json:"B2"`
	A1               *test.A        `protobuf:"bytes,5,opt,name=A1,json=a1" json:"A1,omitempty"`
	A2               test.A         `protobuf:"bytes,6,opt,name=A2,json=a2" json:"A2"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MoreDefaultsA) Reset()                    { *m = MoreDefaultsA{} }
func (m *MoreDefaultsA) String() string            { return proto.CompactTextString(m) }
func (*MoreDefaultsA) ProtoMessage()               {}
func (*MoreDefaultsA) Descriptor() ([]byte, []int) { return fileDescriptorMd, []int{1} }

const Default_MoreDefaultsA_Field1 int64 = 1234

func (m *MoreDefaultsA) GetField1() int64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_MoreDefaultsA_Field1
}

func (m *MoreDefaultsA) GetField2() int64 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *MoreDefaultsA) GetB1() *MoreDefaultsB {
	if m != nil {
		return m.B1
	}
	return nil
}

func (m *MoreDefaultsA) GetB2() MoreDefaultsB {
	if m != nil {
		return m.B2
	}
	return MoreDefaultsB{}
}

func (m *MoreDefaultsA) GetA1() *test.A {
	if m != nil {
		return m.A1
	}
	return nil
}

func (m *MoreDefaultsA) GetA2() test.A {
	if m != nil {
		return m.A2
	}
	return test.A{}
}

func init() {
	proto.RegisterType((*MoreDefaultsB)(nil), "moredefaults.MoreDefaultsB")
	proto.RegisterType((*MoreDefaultsA)(nil), "moredefaults.MoreDefaultsA")
}
func (this *MoreDefaultsB) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MoreDefaultsB)
	if !ok {
		that2, ok := that.(MoreDefaultsB)
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
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MoreDefaultsA) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*MoreDefaultsA)
	if !ok {
		that2, ok := that.(MoreDefaultsA)
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
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !this.B1.Equal(that1.B1) {
		return false
	}
	if !this.B2.Equal(&that1.B2) {
		return false
	}
	if !this.A1.Equal(that1.A1) {
		return false
	}
	if !this.A2.Equal(&that1.A2) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func NewPopulatedMoreDefaultsB(r randyMd, easy bool) *MoreDefaultsB {
	this := &MoreDefaultsB{}
	if r.Intn(10) != 0 {
		v1 := randStringMd(r)
		this.Field1 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMd(r, 2)
	}
	return this
}

func NewPopulatedMoreDefaultsA(r randyMd, easy bool) *MoreDefaultsA {
	this := &MoreDefaultsA{}
	if r.Intn(10) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Field1 = &v2
	}
	this.Field2 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	if r.Intn(10) != 0 {
		this.B1 = NewPopulatedMoreDefaultsB(r, easy)
	}
	v3 := NewPopulatedMoreDefaultsB(r, easy)
	this.B2 = *v3
	if r.Intn(10) != 0 {
		this.A1 = test.NewPopulatedA(r, easy)
	}
	v4 := test.NewPopulatedA(r, easy)
	this.A2 = *v4
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMd(r, 7)
	}
	return this
}

type randyMd interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMd(r randyMd) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMd(r randyMd) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneMd(r)
	}
	return string(tmps)
}
func randUnrecognizedMd(r randyMd, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMd(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMd(data []byte, r randyMd, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMd(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateMd(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateMd(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMd(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMd(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMd(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMd(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}

var fileDescriptorMd = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x4d, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0xcd, 0x2f, 0x4a, 0x4d, 0x49, 0x4d, 0x4b, 0x2c, 0xcd,
	0x29, 0x29, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2,
	0x59, 0xca, 0x18, 0xa7, 0xf2, 0x92, 0xd4, 0xe2, 0x12, 0xfd, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c,
	0x54, 0x18, 0x0d, 0xd1, 0xa4, 0xa4, 0xce, 0xc5, 0xeb, 0x0b, 0xb4, 0xd3, 0x05, 0x6a, 0xa7, 0x93,
	0x90, 0x18, 0x17, 0x9b, 0x5b, 0x66, 0x6a, 0x4e, 0x8a, 0xa1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x5b, 0x1a, 0x98, 0xa7, 0xf4, 0x98, 0x11, 0x55, 0xa5, 0xa3, 0x90, 0x0c, 0x8a, 0x4a, 0x66,
	0x2b, 0x16, 0x43, 0x23, 0x63, 0x13, 0x98, 0x7a, 0xb8, 0xac, 0x91, 0x04, 0x13, 0x48, 0xd6, 0x89,
	0xe5, 0xc4, 0x3d, 0x79, 0x06, 0xa8, 0xac, 0x91, 0x90, 0x36, 0x17, 0x93, 0x93, 0xa1, 0x04, 0x33,
	0x50, 0x86, 0xdb, 0x48, 0x5a, 0x0f, 0xd9, 0xd7, 0x7a, 0x28, 0xce, 0x09, 0x62, 0x4a, 0x32, 0x14,
	0x32, 0x04, 0x2a, 0x36, 0x92, 0x60, 0x21, 0xa8, 0x18, 0x6a, 0x07, 0x53, 0x92, 0x91, 0x90, 0x38,
	0x17, 0x93, 0xa3, 0xa1, 0x04, 0x2b, 0x58, 0x0b, 0xbb, 0x1e, 0xc8, 0xff, 0x7a, 0x8e, 0x41, 0x4c,
	0x89, 0x86, 0x42, 0xb2, 0x40, 0x09, 0x23, 0x09, 0x36, 0x14, 0x09, 0x98, 0xbe, 0x44, 0x23, 0x27,
	0x81, 0x13, 0x0f, 0xe5, 0x18, 0x7f, 0x00, 0xf1, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x80, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x4e, 0xe8, 0x88, 0x9d, 0x01, 0x00, 0x00,
}

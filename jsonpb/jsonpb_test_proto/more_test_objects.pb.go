// Code generated by protoc-gen-gogo.
// source: more_test_objects.proto
// DO NOT EDIT!

/*
Package jsonpb is a generated protocol buffer package.

It is generated from these files:
	more_test_objects.proto
	test_objects.proto

It has these top-level messages:
	Simple3
	Mappy
	Simple
	Repeats
	Widget
	Maps
	MsgWithOneof
	Real
	Complex
*/
package jsonpb

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Numeral int32

const (
	Numeral_UNKNOWN Numeral = 0
	Numeral_ARABIC  Numeral = 1
	Numeral_ROMAN   Numeral = 2
)

var Numeral_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARABIC",
	2: "ROMAN",
}
var Numeral_value = map[string]int32{
	"UNKNOWN": 0,
	"ARABIC":  1,
	"ROMAN":   2,
}

func (x Numeral) String() string {
	return proto.EnumName(Numeral_name, int32(x))
}
func (Numeral) EnumDescriptor() ([]byte, []int) { return fileDescriptorMoreTestObjects, []int{0} }

type Simple3 struct {
	Dub float64 `protobuf:"fixed64,1,opt,name=dub,proto3" json:"dub,omitempty"`
}

func (m *Simple3) Reset()                    { *m = Simple3{} }
func (m *Simple3) String() string            { return proto.CompactTextString(m) }
func (*Simple3) ProtoMessage()               {}
func (*Simple3) Descriptor() ([]byte, []int) { return fileDescriptorMoreTestObjects, []int{0} }

type Mappy struct {
	Nummy map[int64]int32    `protobuf:"bytes,1,rep,name=nummy" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Strry map[string]string  `protobuf:"bytes,2,rep,name=strry" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Objjy map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Buggy map[int64]string   `protobuf:"bytes,4,rep,name=buggy" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Booly map[bool]bool      `protobuf:"bytes,5,rep,name=booly" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Enumy map[string]Numeral `protobuf:"bytes,6,rep,name=enumy" json:"enumy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=jsonpb.Numeral"`
}

func (m *Mappy) Reset()                    { *m = Mappy{} }
func (m *Mappy) String() string            { return proto.CompactTextString(m) }
func (*Mappy) ProtoMessage()               {}
func (*Mappy) Descriptor() ([]byte, []int) { return fileDescriptorMoreTestObjects, []int{1} }

func (m *Mappy) GetNummy() map[int64]int32 {
	if m != nil {
		return m.Nummy
	}
	return nil
}

func (m *Mappy) GetStrry() map[string]string {
	if m != nil {
		return m.Strry
	}
	return nil
}

func (m *Mappy) GetObjjy() map[int32]*Simple3 {
	if m != nil {
		return m.Objjy
	}
	return nil
}

func (m *Mappy) GetBuggy() map[int64]string {
	if m != nil {
		return m.Buggy
	}
	return nil
}

func (m *Mappy) GetBooly() map[bool]bool {
	if m != nil {
		return m.Booly
	}
	return nil
}

func (m *Mappy) GetEnumy() map[string]Numeral {
	if m != nil {
		return m.Enumy
	}
	return nil
}

func init() {
	proto.RegisterType((*Simple3)(nil), "jsonpb.Simple3")
	proto.RegisterType((*Mappy)(nil), "jsonpb.Mappy")
	proto.RegisterEnum("jsonpb.Numeral", Numeral_name, Numeral_value)
}

var fileDescriptorMoreTestObjects = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0xcd, 0x4a, 0xf3, 0x40,
	0x18, 0x85, 0xbf, 0x34, 0xdf, 0xa4, 0xcd, 0x5b, 0xd0, 0x30, 0x08, 0x0e, 0xba, 0x91, 0x82, 0x50,
	0x04, 0xb3, 0x68, 0x37, 0xc5, 0x5d, 0x2b, 0x5d, 0x14, 0x69, 0x0a, 0x29, 0xe2, 0xb2, 0x34, 0x3a,
	0x14, 0x6b, 0xd2, 0x09, 0xf9, 0x11, 0xe6, 0x82, 0xbc, 0x4f, 0xf3, 0x4e, 0x52, 0x33, 0x96, 0x01,
	0x77, 0xd3, 0x9e, 0xe7, 0x09, 0x67, 0x4e, 0x08, 0x5c, 0x26, 0x22, 0xe3, 0x9b, 0x82, 0xe7, 0xc5,
	0x46, 0x44, 0x7b, 0xfe, 0x5a, 0xe4, 0x7e, 0x9a, 0x89, 0x42, 0x50, 0x67, 0x9f, 0x8b, 0x43, 0x1a,
	0x0d, 0xae, 0xa1, 0xbb, 0x7e, 0x4f, 0xd2, 0x98, 0x8f, 0xa9, 0x07, 0xf6, 0x5b, 0x19, 0x31, 0xeb,
	0xc6, 0x1a, 0x5a, 0x21, 0x1e, 0x07, 0x5f, 0x04, 0xc8, 0x72, 0x9b, 0xa6, 0x92, 0xfa, 0x40, 0x0e,
	0x65, 0x92, 0xc8, 0x2a, 0xb5, 0x87, 0xfd, 0x11, 0xf3, 0x6b, 0xdd, 0x57, 0xa9, 0x1f, 0x60, 0x34,
	0x3f, 0x14, 0x99, 0x0c, 0x6b, 0x0c, 0xf9, 0xbc, 0xc8, 0x32, 0xc9, 0x3a, 0x26, 0x7e, 0x8d, 0x51,
	0xc3, 0x2b, 0x0c, 0xf9, 0xaa, 0xdf, 0x5e, 0x32, 0xdb, 0xc4, 0xaf, 0x30, 0x6a, 0x78, 0x85, 0x21,
	0x1f, 0x95, 0xbb, 0x9d, 0x64, 0xff, 0x4d, 0xfc, 0x0c, 0xa3, 0x86, 0x57, 0x98, 0xe2, 0x85, 0x88,
	0x25, 0x23, 0x46, 0x1e, 0xa3, 0x23, 0x8f, 0x67, 0xe4, 0x79, 0x75, 0x13, 0xc9, 0x1c, 0x13, 0x3f,
	0xc7, 0xa8, 0xe1, 0x15, 0x76, 0x35, 0x01, 0x68, 0x47, 0xc0, 0x25, 0x3f, 0xb8, 0x54, 0x4b, 0xda,
	0x21, 0x1e, 0xe9, 0x05, 0x90, 0xcf, 0x6d, 0x5c, 0xf2, 0x6a, 0x0f, 0x6b, 0x48, 0xc2, 0xfa, 0xc7,
	0x43, 0x67, 0x62, 0xa1, 0xd9, 0xce, 0xa1, 0x9b, 0xae, 0xc1, 0x74, 0x75, 0x73, 0x01, 0xd0, 0x0e,
	0xa3, 0x9b, 0xa4, 0x36, 0x6f, 0x75, 0xb3, 0x3f, 0x3a, 0x3f, 0xde, 0xa1, 0x79, 0xdf, 0x27, 0x25,
	0xda, 0xcd, 0xfe, 0xaa, 0xef, 0x9e, 0x9a, 0x3f, 0xeb, 0xe9, 0x66, 0xcf, 0x60, 0xf6, 0x4e, 0xea,
	0xb7, 0x3b, 0x1a, 0x2e, 0xfe, 0xab, 0xfe, 0x59, 0x5b, 0xbf, 0xda, 0x99, 0x67, 0xdb, 0x58, 0x7b,
	0xd4, 0xdd, 0x3d, 0x74, 0x9b, 0x7f, 0x69, 0x1f, 0xba, 0xcf, 0xc1, 0x53, 0xb0, 0x7a, 0x09, 0xbc,
	0x7f, 0x14, 0xc0, 0x99, 0x86, 0xd3, 0xd9, 0xe2, 0xd1, 0xb3, 0xa8, 0x0b, 0x24, 0x5c, 0x2d, 0xa7,
	0x81, 0xd7, 0x89, 0x1c, 0xf5, 0x09, 0x8c, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x04, 0xff,
	0x62, 0x1d, 0x03, 0x00, 0x00,
}

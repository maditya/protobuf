// Code generated by protoc-gen-gogo.
// source: enumprefix.proto
// DO NOT EDIT!

/*
	Package enumprefix is a generated protocol buffer package.

	It is generated from these files:
		enumprefix.proto

	It has these top-level messages:
		MyMessage
*/
package enumprefix

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import test "github.com/maditya/protobuf/test"
import _ "github.com/maditya/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyMessage struct {
	TheField         test.TheTestEnum `protobuf:"varint,1,opt,name=TheField,json=theField,enum=test.TheTestEnum" json:"TheField"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MyMessage) Reset()                    { *m = MyMessage{} }
func (m *MyMessage) String() string            { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()               {}
func (*MyMessage) Descriptor() ([]byte, []int) { return fileDescriptorEnumprefix, []int{0} }

func (m *MyMessage) GetTheField() test.TheTestEnum {
	if m != nil {
		return m.TheField
	}
	return test.A
}

func init() {
	proto.RegisterType((*MyMessage)(nil), "enumprefix.MyMessage")
}

func init() { proto.RegisterFile("enumprefix.proto", fileDescriptorEnumprefix) }

var fileDescriptorEnumprefix = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcd, 0x2b, 0xcd,
	0x2d, 0x28, 0x4a, 0x4d, 0xcb, 0xac, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x26, 0xa6,
	0x64, 0x96, 0x54, 0x26, 0xea, 0x83, 0x55, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0xa4, 0x16, 0x97, 0xe8,
	0x97, 0x64, 0xa4, 0x82, 0x68, 0x88, 0x5e, 0x29, 0x03, 0x7c, 0xea, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1,
	0x1c, 0x30, 0x0b, 0xa2, 0x43, 0xc9, 0x81, 0x8b, 0xd3, 0xb7, 0xd2, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x3d, 0x55, 0xc8, 0x98, 0x8b, 0x23, 0x24, 0x23, 0xd5, 0x2d, 0x33, 0x35, 0x27, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xcf, 0x48, 0x50, 0x0f, 0x6c, 0x7a, 0x48, 0x46, 0x6a, 0x48, 0x6a, 0x71, 0x89,
	0x6b, 0x5e, 0x69, 0xae, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x1c, 0x25, 0x50, 0x85, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x91, 0x5f, 0xfe, 0xc2, 0x00, 0x00, 0x00,
}

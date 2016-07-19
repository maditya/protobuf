// Code generated by protoc-gen-gogo.
// source: proto.proto
// DO NOT EDIT!

/*
Package importdedup is a generated protocol buffer package.

It is generated from these files:
	proto.proto

It has these top-level messages:
	Object
*/
package importdedup

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import subpkg "github.com/maditya/protobuf/test/importdedup/subpkg"

import github_com_maditya_protobuf_test_importdedup_subpkg "github.com/maditya/protobuf/test/importdedup/subpkg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Object struct {
	CustomField      *github_com_maditya_protobuf_test_importdedup_subpkg.CustomType `protobuf:"bytes,1,opt,name=CustomField,json=customField,customtype=github.com/maditya/protobuf/test/importdedup/subpkg.CustomType" json:"CustomField,omitempty"`
	SubObject        *subpkg.SubObject                                               `protobuf:"bytes,2,opt,name=SubObject,json=subObject" json:"SubObject,omitempty"`
	XXX_unrecognized []byte                                                          `json:"-"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorProto, []int{0} }

func (m *Object) GetSubObject() *subpkg.SubObject {
	if m != nil {
		return m.SubObject
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "importdedup.Object")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptorProto) }

var fileDescriptorProto = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xdc, 0x99, 0xb9, 0x05, 0xf9, 0x45, 0x25, 0x29, 0xa9, 0x29, 0xa5,
	0x05, 0x52, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb9, 0x89,
	0x29, 0x99, 0x25, 0x95, 0x89, 0xfa, 0x60, 0x65, 0x49, 0xa5, 0x69, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9,
	0x60, 0x0e, 0x98, 0x05, 0xd1, 0x2e, 0xe5, 0x84, 0x4f, 0x47, 0x49, 0x6a, 0x71, 0x89, 0x3e, 0x92,
	0xf9, 0xfa, 0xc5, 0xa5, 0x49, 0x05, 0xd9, 0xe9, 0x60, 0x0a, 0xe1, 0x04, 0xa5, 0xf9, 0x8c, 0x5c,
	0x6c, 0xfe, 0x49, 0x59, 0xa9, 0xc9, 0x25, 0x42, 0x29, 0x5c, 0xdc, 0xce, 0xa5, 0xc5, 0x25, 0xf9,
	0xb9, 0x6e, 0x99, 0xa9, 0x39, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x4e, 0x4e, 0xb7, 0xee,
	0xc9, 0xdb, 0x91, 0x61, 0x8f, 0x1e, 0xc4, 0xa8, 0x90, 0xca, 0x82, 0xd4, 0x20, 0xee, 0x64, 0x84,
	0xb1, 0x42, 0xfa, 0x5c, 0x9c, 0xc1, 0xa5, 0x49, 0x10, 0x2b, 0x25, 0x98, 0x80, 0x76, 0x70, 0x1b,
	0x09, 0xea, 0x41, 0xf5, 0xc0, 0x25, 0x82, 0x38, 0x8b, 0x61, 0x4c, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x08, 0x14, 0xd9, 0xf6, 0x32, 0x01, 0x00, 0x00,
}

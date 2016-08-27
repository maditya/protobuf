// Code generated by protoc-gen-gogo.
// source: file.dot.proto
// DO NOT EDIT!

/*
Package filedotname is a generated protocol buffer package.

It is generated from these files:
	file.dot.proto

It has these top-level messages:
	M
*/
package filedotname

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"

import github_com_maditya_protobuf_protoc_gen_gogo_descriptor "github.com/maditya/protobuf/protoc-gen-gogo/descriptor"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type M struct {
	A                *string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *M) Reset()                    { *m = M{} }
func (*M) ProtoMessage()               {}
func (*M) Descriptor() ([]byte, []int) { return fileDescriptorFileDot, []int{0} }

func init() {
	proto.RegisterType((*M)(nil), "filedotname.M")
}
func (this *M) Description() (desc *github_com_maditya_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return FileDotDescription()
}
func FileDotDescription() (desc *github_com_maditya_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_maditya_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3435 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x4d, 0x70, 0x1c, 0xd5,
		0xb5, 0x56, 0xcf, 0x8f, 0x34, 0x73, 0x66, 0x34, 0x6a, 0xb5, 0x84, 0x3c, 0x16, 0x30, 0x96, 0x05,
		0x3c, 0x04, 0xbc, 0x27, 0x53, 0xc6, 0x36, 0xf6, 0xf8, 0x81, 0x6b, 0x24, 0x8d, 0x85, 0x5c, 0xfa,
		0x7b, 0x2d, 0x09, 0x0c, 0x2c, 0xba, 0xae, 0x7a, 0xee, 0x8c, 0xda, 0xee, 0xe9, 0x9e, 0xd7, 0xdd,
		0x63, 0x5b, 0x5e, 0x99, 0xe2, 0xfd, 0x14, 0x45, 0xbd, 0x97, 0xdf, 0xaa, 0xf0, 0x1f, 0x42, 0x25,
		0x81, 0x90, 0x3f, 0xc8, 0x5f, 0xa5, 0xb2, 0x62, 0x43, 0xc2, 0x2a, 0x05, 0xbb, 0x2c, 0xb2, 0xc0,
		0x0a, 0x55, 0x21, 0x09, 0x49, 0x48, 0x95, 0xab, 0x42, 0x15, 0x9b, 0xd4, 0xfd, 0xeb, 0xe9, 0x9e,
		0x19, 0xb9, 0x47, 0x54, 0x01, 0x59, 0x69, 0xee, 0xb9, 0xe7, 0xfb, 0xfa, 0xdc, 0x73, 0xcf, 0x3d,
		0xe7, 0xf4, 0x6d, 0xc1, 0x63, 0x47, 0x60, 0xa2, 0x66, 0xdb, 0x35, 0x13, 0x1f, 0x6a, 0x38, 0xb6,
		0x67, 0x6f, 0x36, 0xab, 0x87, 0x2a, 0xd8, 0xd5, 0x1d, 0xa3, 0xe1, 0xd9, 0xce, 0x34, 0x95, 0x29,
		0x43, 0x4c, 0x63, 0x5a, 0x68, 0x4c, 0x2e, 0xc1, 0xf0, 0x69, 0xc3, 0xc4, 0x73, 0xbe, 0xe2, 0x1a,
		0xf6, 0x94, 0xe3, 0x90, 0xa8, 0x1a, 0x26, 0xce, 0x4b, 0x13, 0xf1, 0xa9, 0xcc, 0xe1, 0x5b, 0xa7,
		0xdb, 0x40, 0xd3, 0x61, 0xc4, 0x2a, 0x11, 0xab, 0x14, 0x31, 0xf9, 0x5e, 0x02, 0x46, 0xba, 0xcc,
		0x2a, 0x0a, 0x24, 0x2c, 0x54, 0x27, 0x8c, 0xd2, 0x54, 0x5a, 0xa5, 0xbf, 0x95, 0x3c, 0x0c, 0x34,
		0x90, 0x7e, 0x1e, 0xd5, 0x70, 0x3e, 0x46, 0xc5, 0x62, 0xa8, 0x14, 0x00, 0x2a, 0xb8, 0x81, 0xad,
		0x0a, 0xb6, 0xf4, 0xed, 0x7c, 0x7c, 0x22, 0x3e, 0x95, 0x56, 0x03, 0x12, 0xe5, 0x2e, 0x18, 0x6e,
		0x34, 0x37, 0x4d, 0x43, 0xd7, 0x02, 0x6a, 0x30, 0x11, 0x9f, 0x4a, 0xaa, 0x32, 0x9b, 0x98, 0x6b,
		0x29, 0xdf, 0x0e, 0x43, 0x17, 0x31, 0x3a, 0x1f, 0x54, 0xcd, 0x50, 0xd5, 0x1c, 0x11, 0x07, 0x14,
		0x67, 0x21, 0x5b, 0xc7, 0xae, 0x8b, 0x6a, 0x58, 0xf3, 0xb6, 0x1b, 0x38, 0x9f, 0xa0, 0xab, 0x9f,
		0xe8, 0x58, 0x7d, 0xfb, 0xca, 0x33, 0x1c, 0xb5, 0xbe, 0xdd, 0xc0, 0x4a, 0x09, 0xd2, 0xd8, 0x6a,
		0xd6, 0x19, 0x43, 0x72, 0x17, 0xff, 0x95, 0xad, 0x66, 0xbd, 0x9d, 0x25, 0x45, 0x60, 0x9c, 0x62,
		0xc0, 0xc5, 0xce, 0x05, 0x43, 0xc7, 0xf9, 0x7e, 0x4a, 0x70, 0x7b, 0x07, 0xc1, 0x1a, 0x9b, 0x6f,
		0xe7, 0x10, 0x38, 0x65, 0x16, 0xd2, 0xf8, 0x92, 0x87, 0x2d, 0xd7, 0xb0, 0xad, 0xfc, 0x00, 0x25,
		0xb9, 0xad, 0xcb, 0x2e, 0x62, 0xb3, 0xd2, 0x4e, 0xd1, 0xc2, 0x29, 0xc7, 0x60, 0xc0, 0x6e, 0x78,
		0x86, 0x6d, 0xb9, 0xf9, 0xd4, 0x84, 0x34, 0x95, 0x39, 0x7c, 0x53, 0xd7, 0x40, 0x58, 0x61, 0x3a,
		0xaa, 0x50, 0x56, 0x16, 0x40, 0x76, 0xed, 0xa6, 0xa3, 0x63, 0x4d, 0xb7, 0x2b, 0x58, 0x33, 0xac,
		0xaa, 0x9d, 0x4f, 0x53, 0x82, 0x03, 0x9d, 0x0b, 0xa1, 0x8a, 0xb3, 0x76, 0x05, 0x2f, 0x58, 0x55,
		0x5b, 0xcd, 0xb9, 0xa1, 0xb1, 0x32, 0x06, 0xfd, 0xee, 0xb6, 0xe5, 0xa1, 0x4b, 0xf9, 0x2c, 0x8d,
		0x10, 0x3e, 0x9a, 0xfc, 0x7b, 0x12, 0x86, 0x7a, 0x09, 0xb1, 0x93, 0x90, 0xac, 0x92, 0x55, 0xe6,
		0x63, 0x7b, 0xf1, 0x01, 0xc3, 0x84, 0x9d, 0xd8, 0xff, 0x09, 0x9d, 0x58, 0x82, 0x8c, 0x85, 0x5d,
		0x0f, 0x57, 0x58, 0x44, 0xc4, 0x7b, 0x8c, 0x29, 0x60, 0xa0, 0xce, 0x90, 0x4a, 0x7c, 0xa2, 0x90,
		0x3a, 0x0b, 0x43, 0xbe, 0x49, 0x9a, 0x83, 0xac, 0x9a, 0x88, 0xcd, 0x43, 0x51, 0x96, 0x4c, 0x97,
		0x05, 0x4e, 0x25, 0x30, 0x35, 0x87, 0x43, 0x63, 0x65, 0x0e, 0xc0, 0xb6, 0xb0, 0x5d, 0xd5, 0x2a,
		0x58, 0x37, 0xf3, 0xa9, 0x5d, 0xbc, 0xb4, 0x42, 0x54, 0x3a, 0xbc, 0x64, 0x33, 0xa9, 0x6e, 0x2a,
		0x27, 0x5a, 0xa1, 0x36, 0xb0, 0x4b, 0xa4, 0x2c, 0xb1, 0x43, 0xd6, 0x11, 0x6d, 0x1b, 0x90, 0x73,
		0x30, 0x89, 0x7b, 0x5c, 0xe1, 0x2b, 0x4b, 0x53, 0x23, 0xa6, 0x23, 0x57, 0xa6, 0x72, 0x18, 0x5b,
		0xd8, 0xa0, 0x13, 0x1c, 0x2a, 0xb7, 0x80, 0x2f, 0xd0, 0x68, 0x58, 0x01, 0xcd, 0x42, 0x59, 0x21,
		0x5c, 0x46, 0x75, 0x3c, 0x7e, 0x1c, 0x72, 0x61, 0xf7, 0x28, 0xa3, 0x90, 0x74, 0x3d, 0xe4, 0x78,
		0x34, 0x0a, 0x93, 0x2a, 0x1b, 0x28, 0x32, 0xc4, 0xb1, 0x55, 0xa1, 0x59, 0x2e, 0xa9, 0x92, 0x9f,
		0xe3, 0xf7, 0xc2, 0x60, 0xe8, 0xf1, 0xbd, 0x02, 0x27, 0x9f, 0xea, 0x87, 0xd1, 0x6e, 0x31, 0xd7,
		0x35, 0xfc, 0xc7, 0xa0, 0xdf, 0x6a, 0xd6, 0x37, 0xb1, 0x93, 0x8f, 0x53, 0x06, 0x3e, 0x52, 0x4a,
		0x90, 0x34, 0xd1, 0x26, 0x36, 0xf3, 0x89, 0x09, 0x69, 0x2a, 0x77, 0xf8, 0xae, 0x9e, 0xa2, 0x7a,
		0x7a, 0x91, 0x40, 0x54, 0x86, 0x54, 0xee, 0x87, 0x04, 0x4f, 0x71, 0x84, 0xe1, 0xce, 0xde, 0x18,
		0x48, 0x2c, 0xaa, 0x14, 0xa7, 0xdc, 0x08, 0x69, 0xf2, 0x97, 0xf9, 0xb6, 0x9f, 0xda, 0x9c, 0x22,
		0x02, 0xe2, 0x57, 0x65, 0x1c, 0x52, 0x34, 0xcc, 0x2a, 0x58, 0x94, 0x06, 0x7f, 0x4c, 0x36, 0xa6,
		0x82, 0xab, 0xa8, 0x69, 0x7a, 0xda, 0x05, 0x64, 0x36, 0x31, 0x0d, 0x98, 0xb4, 0x9a, 0xe5, 0xc2,
		0x07, 0x89, 0x4c, 0x39, 0x00, 0x19, 0x16, 0x95, 0x86, 0x55, 0xc1, 0x97, 0x68, 0xf6, 0x49, 0xaa,
		0x2c, 0x50, 0x17, 0x88, 0x84, 0x3c, 0xfe, 0x9c, 0x6b, 0x5b, 0x62, 0x6b, 0xe9, 0x23, 0x88, 0x80,
		0x3e, 0xfe, 0xde, 0xf6, 0xc4, 0x77, 0x73, 0xf7, 0xe5, 0xb5, 0xc7, 0xe2, 0xe4, 0xcf, 0x62, 0x90,
		0xa0, 0xe7, 0x6d, 0x08, 0x32, 0xeb, 0x0f, 0xaf, 0x96, 0xb5, 0xb9, 0x95, 0x8d, 0x99, 0xc5, 0xb2,
		0x2c, 0x29, 0x39, 0x00, 0x2a, 0x38, 0xbd, 0xb8, 0x52, 0x5a, 0x97, 0x63, 0xfe, 0x78, 0x61, 0x79,
		0xfd, 0xd8, 0x11, 0x39, 0xee, 0x03, 0x36, 0x98, 0x20, 0x11, 0x54, 0xb8, 0xe7, 0xb0, 0x9c, 0x54,
		0x64, 0xc8, 0x32, 0x82, 0x85, 0xb3, 0xe5, 0xb9, 0x63, 0x47, 0xe4, 0xfe, 0xb0, 0xe4, 0x9e, 0xc3,
		0xf2, 0x80, 0x32, 0x08, 0x69, 0x2a, 0x99, 0x59, 0x59, 0x59, 0x94, 0x53, 0x3e, 0xe7, 0xda, 0xba,
		0xba, 0xb0, 0x3c, 0x2f, 0xa7, 0x7d, 0xce, 0x79, 0x75, 0x65, 0x63, 0x55, 0x06, 0x9f, 0x61, 0xa9,
		0xbc, 0xb6, 0x56, 0x9a, 0x2f, 0xcb, 0x19, 0x5f, 0x63, 0xe6, 0xe1, 0xf5, 0xf2, 0x9a, 0x9c, 0x0d,
		0x99, 0x75, 0xcf, 0x61, 0x79, 0xd0, 0x7f, 0x44, 0x79, 0x79, 0x63, 0x49, 0xce, 0x29, 0xc3, 0x30,
		0xc8, 0x1e, 0x21, 0x8c, 0x18, 0x6a, 0x13, 0x1d, 0x3b, 0x22, 0xcb, 0x2d, 0x43, 0x18, 0xcb, 0x70,
		0x48, 0x70, 0xec, 0x88, 0xac, 0x4c, 0xce, 0x42, 0x92, 0x46, 0x97, 0xa2, 0x40, 0x6e, 0xb1, 0x34,
		0x53, 0x5e, 0xd4, 0x56, 0x56, 0xd7, 0x17, 0x56, 0x96, 0x4b, 0x8b, 0xb2, 0xd4, 0x92, 0xa9, 0xe5,
		0xff, 0xd8, 0x58, 0x50, 0xcb, 0x73, 0x72, 0x2c, 0x28, 0x5b, 0x2d, 0x97, 0xd6, 0xcb, 0x73, 0x72,
		0x7c, 0x52, 0x87, 0xd1, 0x6e, 0x79, 0xa6, 0xeb, 0xc9, 0x08, 0x6c, 0x71, 0x6c, 0x97, 0x2d, 0xa6,
		0x5c, 0x1d, 0x5b, 0xfc, 0x92, 0x04, 0x23, 0x5d, 0x72, 0x6d, 0xd7, 0x87, 0x9c, 0x82, 0x24, 0x0b,
		0x51, 0x56, 0x7d, 0xee, 0xe8, 0x9a, 0xb4, 0x69, 0xc0, 0x76, 0x54, 0x20, 0x8a, 0x0b, 0x56, 0xe0,
		0xf8, 0x2e, 0x15, 0x98, 0x50, 0x74, 0x18, 0xf9, 0xb8, 0x04, 0xf9, 0xdd, 0xb8, 0x23, 0x12, 0x45,
		0x2c, 0x94, 0x28, 0x4e, 0xb6, 0x1b, 0x70, 0x70, 0xf7, 0x35, 0x74, 0x58, 0xf1, 0xb2, 0x04, 0x63,
		0xdd, 0x1b, 0x95, 0xae, 0x36, 0xdc, 0x0f, 0xfd, 0x75, 0xec, 0x6d, 0xd9, 0xa2, 0x58, 0xff, 0x4b,
		0x97, 0x12, 0x40, 0xa6, 0xdb, 0x7d, 0xc5, 0x51, 0xc1, 0x1a, 0x12, 0xdf, 0xad, 0xdb, 0x60, 0xd6,
		0x74, 0x58, 0xfa, 0x44, 0x0c, 0x6e, 0xe8, 0x4a, 0xde, 0xd5, 0xd0, 0x9b, 0x01, 0x0c, 0xab, 0xd1,
		0xf4, 0x58, 0x41, 0x66, 0xf9, 0x29, 0x4d, 0x25, 0xf4, 0xec, 0x93, 0xdc, 0xd3, 0xf4, 0xfc, 0xf9,
		0x38, 0x9d, 0x07, 0x26, 0xa2, 0x0a, 0xc7, 0x5b, 0x86, 0x26, 0xa8, 0xa1, 0x85, 0x5d, 0x56, 0xda,
		0x51, 0xeb, 0xee, 0x06, 0x59, 0x37, 0x0d, 0x6c, 0x79, 0x9a, 0xeb, 0x39, 0x18, 0xd5, 0x0d, 0xab,
		0x46, 0x13, 0x70, 0xaa, 0x98, 0xac, 0x22, 0xd3, 0xc5, 0xea, 0x10, 0x9b, 0x5e, 0x13, 0xb3, 0x04,
		0x41, 0xab, 0x8c, 0x13, 0x40, 0xf4, 0x87, 0x10, 0x6c, 0xda, 0x47, 0x4c, 0x3e, 0x39, 0x00, 0x99,
		0x40, 0x5b, 0xa7, 0x1c, 0x84, 0xec, 0x39, 0x74, 0x01, 0x69, 0xa2, 0x55, 0x67, 0x9e, 0xc8, 0x10,
		0xd9, 0x2a, 0x6f, 0xd7, 0xef, 0x86, 0x51, 0xaa, 0x62, 0x37, 0x3d, 0xec, 0x68, 0xba, 0x89, 0x5c,
		0x97, 0x3a, 0x2d, 0x45, 0x55, 0x15, 0x32, 0xb7, 0x42, 0xa6, 0x66, 0xc5, 0x8c, 0x72, 0x14, 0x46,
		0x28, 0xa2, 0xde, 0x34, 0x3d, 0xa3, 0x61, 0x62, 0x8d, 0xbc, 0x3c, 0xb8, 0x34, 0x11, 0xfb, 0x96,
		0x0d, 0x13, 0x8d, 0x25, 0xae, 0x40, 0x2c, 0x72, 0x95, 0x79, 0xb8, 0x99, 0xc2, 0x6a, 0xd8, 0xc2,
		0x0e, 0xf2, 0xb0, 0x86, 0xff, 0xb3, 0x89, 0x4c, 0x57, 0x43, 0x56, 0x45, 0xdb, 0x42, 0xee, 0x56,
		0x7e, 0x34, 0x48, 0xb0, 0x9f, 0xe8, 0xce, 0x73, 0xd5, 0x32, 0xd5, 0x2c, 0x59, 0x95, 0x07, 0x90,
		0xbb, 0xa5, 0x14, 0x61, 0x8c, 0x12, 0xb9, 0x9e, 0x63, 0x58, 0x35, 0x4d, 0xdf, 0xc2, 0xfa, 0x79,
		0xad, 0xe9, 0x55, 0x8f, 0xe7, 0x6f, 0x0c, 0x32, 0x50, 0x23, 0xd7, 0xa8, 0xce, 0x2c, 0x51, 0xd9,
		0xf0, 0xaa, 0xc7, 0x95, 0x35, 0xc8, 0x92, 0xfd, 0xa8, 0x1b, 0x97, 0xb1, 0x56, 0xb5, 0x1d, 0x5a,
		0x5c, 0x72, 0x5d, 0x0e, 0x77, 0xc0, 0x89, 0xd3, 0x2b, 0x1c, 0xb0, 0x64, 0x57, 0x70, 0x31, 0xb9,
		0xb6, 0x5a, 0x2e, 0xcf, 0xa9, 0x19, 0xc1, 0x72, 0xda, 0x76, 0x48, 0x4c, 0xd5, 0x6c, 0xdf, 0xc7,
		0x19, 0x16, 0x53, 0x35, 0x5b, 0x78, 0xf8, 0x28, 0x8c, 0xe8, 0x3a, 0x5b, 0xb6, 0xa1, 0x6b, 0xbc,
		0xcb, 0x77, 0xf3, 0x72, 0xc8, 0x5f, 0xba, 0x3e, 0xcf, 0x14, 0x78, 0x98, 0xbb, 0xca, 0x09, 0xb8,
		0xa1, 0xe5, 0xaf, 0x20, 0x70, 0xb8, 0x63, 0x95, 0xed, 0xd0, 0xa3, 0x30, 0xd2, 0xd8, 0xee, 0x04,
		0x2a, 0xa1, 0x27, 0x36, 0xb6, 0xdb, 0x61, 0xb7, 0xd1, 0x37, 0x37, 0x07, 0xeb, 0xc8, 0xc3, 0x95,
		0xfc, 0xbe, 0xa0, 0x76, 0x60, 0x42, 0x39, 0x04, 0xb2, 0xae, 0x6b, 0xd8, 0x42, 0x9b, 0x26, 0xd6,
		0x90, 0x83, 0x2d, 0xe4, 0xe6, 0x0f, 0x04, 0x95, 0x73, 0xba, 0x5e, 0xa6, 0xb3, 0x25, 0x3a, 0xa9,
		0xdc, 0x09, 0xc3, 0xf6, 0xe6, 0x39, 0x9d, 0x05, 0x97, 0xd6, 0x70, 0x70, 0xd5, 0xb8, 0x94, 0xbf,
		0x95, 0xba, 0x69, 0x88, 0x4c, 0xd0, 0xd0, 0x5a, 0xa5, 0x62, 0xe5, 0x0e, 0x90, 0x75, 0x77, 0x0b,
		0x39, 0x0d, 0x5a, 0xdd, 0xdd, 0x06, 0xd2, 0x71, 0xfe, 0x36, 0xa6, 0xca, 0xe4, 0xcb, 0x42, 0xac,
		0x9c, 0x85, 0xd1, 0xa6, 0x65, 0x58, 0x1e, 0x76, 0x1a, 0x0e, 0x26, 0x4d, 0x3a, 0x3b, 0x69, 0xf9,
		0xdf, 0x0f, 0xec, 0xd2, 0x66, 0x6f, 0x04, 0xb5, 0xd9, 0xee, 0xaa, 0x23, 0xcd, 0x4e, 0xe1, 0x64,
		0x11, 0xb2, 0xc1, 0x4d, 0x57, 0xd2, 0xc0, 0xb6, 0x5d, 0x96, 0x48, 0x0d, 0x9d, 0x5d, 0x99, 0x23,
		0xd5, 0xef, 0x91, 0xb2, 0x1c, 0x23, 0x55, 0x78, 0x71, 0x61, 0xbd, 0xac, 0xa9, 0x1b, 0xcb, 0xeb,
		0x0b, 0x4b, 0x65, 0x39, 0x7e, 0x67, 0x3a, 0xf5, 0xfe, 0x80, 0x7c, 0xe5, 0xca, 0x95, 0x2b, 0xb1,
		0xc9, 0x37, 0x63, 0x90, 0x0b, 0x77, 0xbe, 0xca, 0xbf, 0xc3, 0x3e, 0xf1, 0x9a, 0xea, 0x62, 0x4f,
		0xbb, 0x68, 0x38, 0x34, 0x0e, 0xeb, 0x88, 0xf5, 0x8e, 0xbe, 0x0b, 0x47, 0xb9, 0xd6, 0x1a, 0xf6,
		0x1e, 0x32, 0x1c, 0x12, 0x65, 0x75, 0xe4, 0x29, 0x8b, 0x70, 0xc0, 0xb2, 0x35, 0xd7, 0x43, 0x56,
		0x05, 0x39, 0x15, 0xad, 0x75, 0x41, 0xa0, 0x21, 0x5d, 0xc7, 0xae, 0x6b, 0xb3, 0x12, 0xe0, 0xb3,
		0xdc, 0x64, 0xd9, 0x6b, 0x5c, 0xb9, 0x95, 0x1b, 0x4b, 0x5c, 0xb5, 0x6d, 0xbb, 0xe3, 0xbb, 0x6d,
		0xf7, 0x8d, 0x90, 0xae, 0xa3, 0x86, 0x86, 0x2d, 0xcf, 0xd9, 0xa6, 0xfd, 0x5a, 0x4a, 0x4d, 0xd5,
		0x51, 0xa3, 0x4c, 0xc6, 0x9f, 0xde, 0x1e, 0x04, 0xfd, 0xf8, 0xdb, 0x38, 0x64, 0x83, 0x3d, 0x1b,
		0x69, 0x81, 0x75, 0x9a, 0x9f, 0x25, 0x7a, 0x7c, 0x6f, 0xb9, 0x6e, 0x87, 0x37, 0x3d, 0x4b, 0x12,
		0x77, 0xb1, 0x9f, 0x75, 0x52, 0x2a, 0x43, 0x92, 0xa2, 0x49, 0x0e, 0x2c, 0x66, 0xfd, 0x79, 0x4a,
		0xe5, 0x23, 0x65, 0x1e, 0xfa, 0xcf, 0xb9, 0x94, 0xbb, 0x9f, 0x72, 0xdf, 0x7a, 0x7d, 0xee, 0x33,
		0x6b, 0x94, 0x3c, 0x7d, 0x66, 0x4d, 0x5b, 0x5e, 0x51, 0x97, 0x4a, 0x8b, 0x2a, 0x87, 0x2b, 0xfb,
		0x21, 0x61, 0xa2, 0xcb, 0xdb, 0xe1, 0x14, 0x4f, 0x45, 0xbd, 0x3a, 0x7e, 0x3f, 0x24, 0x2e, 0x62,
		0x74, 0x3e, 0x9c, 0x58, 0xa9, 0xe8, 0x53, 0x0c, 0xfd, 0x43, 0x90, 0xa4, 0xfe, 0x52, 0x00, 0xb8,
		0xc7, 0xe4, 0x3e, 0x25, 0x05, 0x89, 0xd9, 0x15, 0x95, 0x84, 0xbf, 0x0c, 0x59, 0x26, 0xd5, 0x56,
		0x17, 0xca, 0xb3, 0x65, 0x39, 0x36, 0x79, 0x14, 0xfa, 0x99, 0x13, 0xc8, 0xd1, 0xf0, 0xdd, 0x20,
		0xf7, 0xf1, 0x21, 0xe7, 0x90, 0xc4, 0xec, 0xc6, 0xd2, 0x4c, 0x59, 0x95, 0x63, 0xc1, 0xed, 0x75,
		0x21, 0x1b, 0x6c, 0xd7, 0x3e, 0x9b, 0x98, 0xfa, 0x85, 0x04, 0x99, 0x40, 0xfb, 0x45, 0x0a, 0x3f,
		0x32, 0x4d, 0xfb, 0xa2, 0x86, 0x4c, 0x03, 0xb9, 0x3c, 0x28, 0x80, 0x8a, 0x4a, 0x44, 0xd2, 0xeb,
		0xa6, 0x7d, 0x26, 0xc6, 0xbf, 0x20, 0x81, 0xdc, 0xde, 0xba, 0xb5, 0x19, 0x28, 0x7d, 0xae, 0x06,
		0x3e, 0x27, 0x41, 0x2e, 0xdc, 0xaf, 0xb5, 0x99, 0x77, 0xf0, 0x73, 0x35, 0xef, 0x59, 0x09, 0x06,
		0x43, 0x5d, 0xda, 0x3f, 0x95, 0x75, 0xcf, 0xc4, 0x61, 0xa4, 0x0b, 0x4e, 0x29, 0xf1, 0x76, 0x96,
		0x75, 0xd8, 0xff, 0xd6, 0xcb, 0xb3, 0xa6, 0x49, 0xb5, 0x5c, 0x45, 0x8e, 0xc7, 0xbb, 0xdf, 0x3b,
		0x40, 0x36, 0x2a, 0xd8, 0xf2, 0x8c, 0xaa, 0x81, 0x1d, 0xfe, 0x0a, 0xce, 0x7a, 0xdc, 0xa1, 0x96,
		0x9c, 0xbd, 0x85, 0xff, 0x2b, 0x28, 0x0d, 0xdb, 0x35, 0x3c, 0xe3, 0x02, 0xd6, 0x0c, 0x4b, 0xbc,
		0xaf, 0x93, 0x9e, 0x37, 0xa1, 0xca, 0x62, 0x66, 0xc1, 0xf2, 0x7c, 0x6d, 0x0b, 0xd7, 0x50, 0x9b,
		0x36, 0xc9, 0x7d, 0x71, 0x55, 0x16, 0x33, 0xbe, 0xf6, 0x41, 0xc8, 0x56, 0xec, 0x26, 0x69, 0x1f,
		0x98, 0x1e, 0x49, 0xb5, 0x92, 0x9a, 0x61, 0x32, 0x5f, 0x85, 0xf7, 0x77, 0xad, 0x8b, 0x82, 0xac,
		0x9a, 0x61, 0x32, 0xa6, 0x72, 0x3b, 0x0c, 0xa1, 0x5a, 0xcd, 0x21, 0xe4, 0x82, 0x88, 0x35, 0xad,
		0x39, 0x5f, 0x4c, 0x15, 0xc7, 0xcf, 0x40, 0x4a, 0xf8, 0x81, 0x54, 0x33, 0xe2, 0x09, 0xad, 0xc1,
		0xae, 0x6b, 0x62, 0x53, 0x69, 0x35, 0x65, 0x89, 0xc9, 0x83, 0x90, 0x35, 0x5c, 0xad, 0x75, 0x6f,
		0x18, 0x9b, 0x88, 0x4d, 0xa5, 0xd4, 0x8c, 0xe1, 0xfa, 0x17, 0x45, 0x93, 0x2f, 0xc7, 0x20, 0x17,
		0xbe, 0xf7, 0x54, 0xe6, 0x20, 0x65, 0xda, 0x3a, 0xa2, 0x81, 0xc0, 0x2e, 0xdd, 0xa7, 0x22, 0xae,
		0x4a, 0xa7, 0x17, 0xb9, 0xbe, 0xea, 0x23, 0xc7, 0x7f, 0x2d, 0x41, 0x4a, 0x88, 0x95, 0x31, 0x48,
		0x34, 0x90, 0xb7, 0x45, 0xe9, 0x92, 0x33, 0x31, 0x59, 0x52, 0xe9, 0x98, 0xc8, 0xdd, 0x06, 0xb2,
		0x68, 0x08, 0x70, 0x39, 0x19, 0x93, 0x7d, 0x35, 0x31, 0xaa, 0xd0, 0x76, 0xd8, 0xae, 0xd7, 0xb1,
		0xe5, 0xb9, 0x62, 0x5f, 0xb9, 0x7c, 0x96, 0x8b, 0x95, 0xbb, 0x60, 0xd8, 0x73, 0x90, 0x61, 0x86,
		0x74, 0x13, 0x54, 0x57, 0x16, 0x13, 0xbe, 0x72, 0x11, 0xf6, 0x0b, 0xde, 0x0a, 0xf6, 0x90, 0xbe,
		0x85, 0x2b, 0x2d, 0x50, 0x3f, 0xbd, 0x54, 0xdb, 0xc7, 0x15, 0xe6, 0xf8, 0xbc, 0xc0, 0x4e, 0xbe,
		0x23, 0xc1, 0xb0, 0x68, 0xe0, 0x2b, 0xbe, 0xb3, 0x96, 0x00, 0x90, 0x65, 0xd9, 0x5e, 0xd0, 0x5d,
		0x9d, 0xa1, 0xdc, 0x81, 0x9b, 0x2e, 0xf9, 0x20, 0x35, 0x40, 0x30, 0x5e, 0x07, 0x68, 0xcd, 0xec,
		0xea, 0xb6, 0x03, 0x90, 0xe1, 0x97, 0xda, 0xf4, 0xcb, 0x08, 0x7b, 0xeb, 0x03, 0x26, 0x22, 0x9d,
		0xbe, 0x32, 0x0a, 0xc9, 0x4d, 0x5c, 0x33, 0x2c, 0x7e, 0xd5, 0xc6, 0x06, 0xe2, 0x02, 0x2f, 0xe1,
		0x5f, 0xe0, 0xcd, 0x3c, 0x0a, 0x23, 0xba, 0x5d, 0x6f, 0x37, 0x77, 0x46, 0x6e, 0x7b, 0xf3, 0x74,
		0x1f, 0x90, 0x1e, 0x81, 0x56, 0x77, 0xf6, 0xa2, 0x24, 0xbd, 0x14, 0x8b, 0xcf, 0xaf, 0xce, 0xbc,
		0x1a, 0x1b, 0x9f, 0x67, 0xd0, 0x55, 0xb1, 0x52, 0x15, 0x57, 0x4d, 0xac, 0x13, 0xeb, 0xe1, 0x9b,
		0xb7, 0xc0, 0xdd, 0x35, 0xc3, 0xdb, 0x6a, 0x6e, 0x4e, 0xeb, 0x76, 0xfd, 0x50, 0x1d, 0x55, 0x0c,
		0x6f, 0x1b, 0xb5, 0xbe, 0x07, 0xd5, 0xec, 0x9a, 0x4d, 0x07, 0xf4, 0x17, 0xff, 0x26, 0x94, 0xf6,
		0xa5, 0xe3, 0x91, 0x1f, 0x90, 0x8a, 0xcb, 0x30, 0xc2, 0x95, 0x35, 0x7a, 0x29, 0xcd, 0x5a, 0x71,
		0xe5, 0xba, 0x17, 0x13, 0xf9, 0xd7, 0xdf, 0xa3, 0xc5, 0x4e, 0x1d, 0xe6, 0x50, 0x32, 0xc7, 0x9a,
		0xf5, 0xa2, 0x0a, 0x37, 0x84, 0xf8, 0xd8, 0xe9, 0xc4, 0x4e, 0x04, 0xe3, 0x9b, 0x9c, 0x71, 0x24,
		0xc0, 0xb8, 0xc6, 0xa1, 0xc5, 0x59, 0x18, 0xdc, 0x0b, 0xd7, 0x2f, 0x39, 0x57, 0x16, 0x07, 0x49,
		0xe6, 0x61, 0x88, 0x92, 0xe8, 0x4d, 0xd7, 0xb3, 0xeb, 0x34, 0xf5, 0x5d, 0x9f, 0xe6, 0x57, 0xef,
		0xb1, 0xe3, 0x92, 0x23, 0xb0, 0x59, 0x1f, 0x55, 0x7c, 0x10, 0x46, 0x89, 0x84, 0x66, 0x97, 0x20,
		0x5b, 0xf4, 0x55, 0x4a, 0xfe, 0x9d, 0xc7, 0xd9, 0xa9, 0x1a, 0xf1, 0x09, 0x02, 0xbc, 0x81, 0x9d,
		0xa8, 0x61, 0xcf, 0xc3, 0x8e, 0xab, 0x21, 0xd3, 0x54, 0xae, 0xfb, 0x91, 0x26, 0xff, 0xf4, 0x07,
		0xe1, 0x9d, 0x98, 0x67, 0xc8, 0x92, 0x69, 0x16, 0x37, 0x60, 0x5f, 0x97, 0x9d, 0xed, 0x81, 0xf3,
		0x19, 0xce, 0x39, 0xda, 0xb1, 0xbb, 0x84, 0x76, 0x15, 0x84, 0xdc, 0xdf, 0x8f, 0x1e, 0x38, 0x9f,
		0xe5, 0x9c, 0x0a, 0xc7, 0x8a, 0x6d, 0x21, 0x8c, 0x67, 0x60, 0xf8, 0x02, 0x76, 0x36, 0x6d, 0x97,
		0xbf, 0xff, 0xf7, 0x40, 0xf7, 0x1c, 0xa7, 0x1b, 0xe2, 0x40, 0x7a, 0x1b, 0x40, 0xb8, 0x4e, 0x40,
		0xaa, 0x8a, 0x74, 0xdc, 0x03, 0xc5, 0xf3, 0x9c, 0x62, 0x80, 0xe8, 0x13, 0x68, 0x09, 0xb2, 0x35,
		0x9b, 0x17, 0x98, 0x68, 0xf8, 0x0b, 0x1c, 0x9e, 0x11, 0x18, 0x4e, 0xd1, 0xb0, 0x1b, 0x4d, 0x93,
		0x54, 0x9f, 0x68, 0x8a, 0xaf, 0x0b, 0x0a, 0x81, 0xe1, 0x14, 0x7b, 0x70, 0xeb, 0x8b, 0x82, 0xc2,
		0x0d, 0xf8, 0xf3, 0x14, 0x64, 0x6c, 0xcb, 0xdc, 0xb6, 0xad, 0x5e, 0x8c, 0xf8, 0x06, 0x67, 0x00,
		0x0e, 0x21, 0x04, 0x27, 0x21, 0xdd, 0xeb, 0x46, 0x7c, 0x8b, 0xc3, 0x53, 0x58, 0xec, 0xc0, 0x3c,
		0x0c, 0x89, 0x24, 0x63, 0xd8, 0x56, 0x0f, 0x14, 0xdf, 0xe6, 0x14, 0xb9, 0x00, 0x8c, 0x2f, 0xc3,
		0xc3, 0xae, 0x57, 0xc3, 0xbd, 0x90, 0xbc, 0x2c, 0x96, 0xc1, 0x21, 0xdc, 0x95, 0x9b, 0xd8, 0xd2,
		0xb7, 0x7a, 0x63, 0x78, 0x45, 0xb8, 0x52, 0x60, 0x08, 0xc5, 0x2c, 0x0c, 0xd6, 0x91, 0xe3, 0x6e,
		0x21, 0xb3, 0xa7, 0xed, 0xf8, 0x0e, 0xe7, 0xc8, 0xfa, 0x20, 0xee, 0x91, 0xa6, 0xb5, 0x17, 0x9a,
		0x57, 0x85, 0x47, 0x02, 0x30, 0x7e, 0xf4, 0x5c, 0x8f, 0x5e, 0xb1, 0xec, 0x85, 0xed, 0xbb, 0xe2,
		0xe8, 0x31, 0xec, 0x52, 0x90, 0xf1, 0x24, 0xa4, 0x5d, 0xe3, 0x72, 0x4f, 0x34, 0xdf, 0x13, 0x3b,
		0x4d, 0x01, 0x04, 0xfc, 0x30, 0xec, 0xef, 0x9a, 0xea, 0x7b, 0x20, 0xfb, 0x3e, 0x27, 0x1b, 0xeb,
		0x92, 0xee, 0x79, 0x4a, 0xd8, 0x2b, 0xe5, 0x0f, 0x44, 0x4a, 0xc0, 0x6d, 0x5c, 0xab, 0xa4, 0x41,
		0x77, 0x51, 0x75, 0x6f, 0x5e, 0xfb, 0xa1, 0xf0, 0x1a, 0xc3, 0x86, 0xbc, 0xb6, 0x0e, 0x63, 0x9c,
		0x71, 0x6f, 0xfb, 0xfa, 0x9a, 0x48, 0xac, 0x0c, 0xbd, 0x11, 0xde, 0xdd, 0x47, 0x61, 0xdc, 0x77,
		0xa7, 0xe8, 0x2d, 0x5d, 0xad, 0x8e, 0x1a, 0x3d, 0x30, 0xbf, 0xce, 0x99, 0x45, 0xc6, 0xf7, 0x9b,
		0x53, 0x77, 0x09, 0x35, 0x08, 0xf9, 0x59, 0xc8, 0x0b, 0xf2, 0xa6, 0xe5, 0x60, 0xdd, 0xae, 0x59,
		0xc6, 0x65, 0x5c, 0xe9, 0x81, 0xfa, 0x47, 0x6d, 0x5b, 0xb5, 0x11, 0x80, 0x13, 0xe6, 0x05, 0x90,
		0xfd, 0x7e, 0x43, 0x33, 0xea, 0x0d, 0xdb, 0xf1, 0x22, 0x18, 0x7f, 0x2c, 0x76, 0xca, 0xc7, 0x2d,
		0x50, 0x58, 0xb1, 0x0c, 0x39, 0x3a, 0xec, 0x35, 0x24, 0x7f, 0xc2, 0x89, 0x06, 0x5b, 0x28, 0x9e,
		0x38, 0x74, 0xbb, 0xde, 0x40, 0x4e, 0x2f, 0xf9, 0xef, 0xa7, 0x22, 0x71, 0x70, 0x08, 0x8b, 0xbe,
		0xa1, 0xb6, 0x4a, 0xac, 0x44, 0x7d, 0xbf, 0xce, 0x3f, 0x76, 0x8d, 0x9f, 0xd9, 0x70, 0x21, 0x2e,
		0x2e, 0x12, 0xf7, 0x84, 0xcb, 0x65, 0x34, 0xd9, 0xe3, 0xd7, 0x7c, 0x0f, 0x85, 0xaa, 0x65, 0xf1,
		0x34, 0x0c, 0x86, 0x4a, 0x65, 0x34, 0xd5, 0x7f, 0x71, 0xaa, 0x6c, 0xb0, 0x52, 0x16, 0x8f, 0x42,
		0x82, 0x94, 0xbd, 0x68, 0xf8, 0x7f, 0x73, 0x38, 0x55, 0x2f, 0xde, 0x07, 0x29, 0x51, 0xee, 0xa2,
		0xa1, 0xff, 0xc3, 0xa1, 0x3e, 0x84, 0xc0, 0x45, 0xa9, 0x8b, 0x86, 0xff, 0xaf, 0x80, 0x0b, 0x08,
		0x81, 0xf7, 0xee, 0xc2, 0x37, 0x9e, 0x4c, 0xf0, 0x74, 0x25, 0x7c, 0x77, 0x12, 0x06, 0x78, 0x8d,
		0x8b, 0x46, 0x3f, 0xc1, 0x1f, 0x2e, 0x10, 0xc5, 0x7b, 0x21, 0xd9, 0xa3, 0xc3, 0xff, 0x8f, 0x43,
		0x99, 0x7e, 0x71, 0x16, 0x32, 0x81, 0xba, 0x16, 0x0d, 0xff, 0x7f, 0x0e, 0x0f, 0xa2, 0x88, 0xe9,
		0xbc, 0xae, 0x45, 0x13, 0x7c, 0x41, 0x98, 0xce, 0x11, 0xc4, 0x6d, 0xa2, 0xa4, 0x45, 0xa3, 0xbf,
		0x28, 0xbc, 0x2e, 0x20, 0xc5, 0x53, 0x90, 0xf6, 0xd3, 0x54, 0x34, 0xfe, 0x4b, 0x1c, 0xdf, 0xc2,
		0x10, 0x0f, 0x04, 0xd2, 0x64, 0x34, 0xc5, 0x97, 0x85, 0x07, 0x02, 0x28, 0x72, 0x8c, 0xda, 0x4b,
		0x5f, 0x34, 0xd3, 0x57, 0xc4, 0x31, 0x6a, 0xab, 0x7c, 0x64, 0x37, 0x69, 0xb6, 0x88, 0xa6, 0xf8,
		0xaa, 0xd8, 0x4d, 0xaa, 0x4f, 0xcc, 0x68, 0xaf, 0x25, 0xd1, 0x1c, 0x5f, 0x13, 0x66, 0xb4, 0x95,
		0x92, 0xe2, 0x2a, 0x28, 0x9d, 0x75, 0x24, 0x9a, 0xef, 0x29, 0xce, 0x37, 0xdc, 0x51, 0x46, 0x8a,
		0x0f, 0xc1, 0x58, 0xf7, 0x1a, 0x12, 0xcd, 0xfa, 0xf4, 0xb5, 0xb6, 0xae, 0x3f, 0x58, 0x42, 0x8a,
		0xeb, 0xad, 0xae, 0x3f, 0x58, 0x3f, 0xa2, 0x69, 0x9f, 0xb9, 0x16, 0x7e, 0xb1, 0x0b, 0x96, 0x8f,
		0x62, 0x09, 0xa0, 0x95, 0xba, 0xa3, 0xb9, 0x9e, 0xe3, 0x5c, 0x01, 0x10, 0x39, 0x1a, 0x3c, 0x73,
		0x47, 0xe3, 0x9f, 0x17, 0x47, 0x83, 0x23, 0x8a, 0x27, 0x21, 0x65, 0x35, 0x4d, 0x93, 0x04, 0x87,
		0x72, 0xfd, 0xff, 0x09, 0xc9, 0xff, 0xe1, 0x63, 0x7e, 0x30, 0x04, 0xa0, 0x78, 0x14, 0x92, 0xb8,
		0xbe, 0x89, 0x2b, 0x51, 0xc8, 0x3f, 0x7e, 0x2c, 0x12, 0x02, 0xd1, 0x2e, 0x9e, 0x02, 0x60, 0x2f,
		0x8d, 0xf4, 0x93, 0x40, 0x04, 0xf6, 0x4f, 0x1f, 0xf3, 0xcf, 0xcd, 0x2d, 0x48, 0x8b, 0x80, 0x7d,
		0xbc, 0xbe, 0x3e, 0xc1, 0x07, 0x61, 0x02, 0xfa, 0xa2, 0x79, 0x02, 0x06, 0xce, 0xb9, 0xb6, 0xe5,
		0xa1, 0x5a, 0x14, 0xfa, 0xcf, 0x1c, 0x2d, 0xf4, 0x89, 0xc3, 0xea, 0xb6, 0x83, 0x3d, 0x54, 0x73,
		0xa3, 0xb0, 0x7f, 0xe1, 0x58, 0x1f, 0x40, 0xc0, 0x3a, 0x72, 0xbd, 0x5e, 0xd6, 0xfd, 0x57, 0x01,
		0x16, 0x00, 0x62, 0x34, 0xf9, 0x7d, 0x1e, 0x6f, 0x47, 0x61, 0x3f, 0x14, 0x46, 0x73, 0xfd, 0xe2,
		0x7d, 0x90, 0x26, 0x3f, 0xd9, 0xbf, 0x60, 0x44, 0x80, 0xff, 0xc6, 0xc1, 0x2d, 0xc4, 0xcc, 0xc1,
		0xee, 0x17, 0x3c, 0x30, 0x6f, 0xcf, 0xdb, 0xec, 0x6a, 0x07, 0x5e, 0x93, 0x20, 0x57, 0x35, 0x4c,
		0x3c, 0x5d, 0xb1, 0x3d, 0x7e, 0x09, 0x93, 0x21, 0xe3, 0x8a, 0xed, 0x11, 0x8f, 0x8f, 0xef, 0xf9,
		0x0e, 0x67, 0x72, 0x18, 0xa4, 0x25, 0x25, 0x0b, 0x12, 0xe2, 0xdf, 0xe7, 0x25, 0x34, 0xb3, 0xf8,
		0xd6, 0xd5, 0x42, 0xdf, 0xdb, 0x57, 0x0b, 0x7d, 0xbf, 0xb9, 0x5a, 0xe8, 0x7b, 0xf7, 0x6a, 0x41,
		0x7a, 0xff, 0x6a, 0x41, 0xfa, 0xf0, 0x6a, 0x41, 0xfa, 0xe8, 0x6a, 0x41, 0xba, 0xb2, 0x53, 0x90,
		0x5e, 0xd9, 0x29, 0x48, 0xaf, 0xed, 0x14, 0xa4, 0x9f, 0xef, 0x14, 0xa4, 0x37, 0x76, 0x0a, 0xd2,
		0x5b, 0x3b, 0x85, 0xbe, 0xb7, 0x77, 0x0a, 0x7d, 0xef, 0xee, 0x14, 0xa4, 0xf7, 0x77, 0x0a, 0x7d,
		0x1f, 0xee, 0x14, 0xa4, 0x8f, 0x76, 0x0a, 0x7d, 0x57, 0x7e, 0x57, 0xe8, 0xfb, 0x47, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xc2, 0x25, 0x20, 0x83, 0x5e, 0x2c, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_maditya_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *M) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *M")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *M but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *M but is not nil && this == nil")
	}
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return fmt.Errorf("A this(%v) Not Equal that(%v)", *this.A, *that1.A)
		}
	} else if this.A != nil {
		return fmt.Errorf("this.A == nil && that.A != nil")
	} else if that1.A != nil {
		return fmt.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *M) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*M)
	if !ok {
		that2, ok := that.(M)
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
	if this.A != nil && that1.A != nil {
		if *this.A != *that1.A {
			return false
		}
	} else if this.A != nil {
		return false
	} else if that1.A != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type MFace interface {
	Proto() github_com_maditya_protobuf_proto.Message
	GetA() *string
}

func (this *M) Proto() github_com_maditya_protobuf_proto.Message {
	return this
}

func (this *M) TestProto() github_com_maditya_protobuf_proto.Message {
	return NewMFromFace(this)
}

func (this *M) GetA() *string {
	return this.A
}

func NewMFromFace(that MFace) *M {
	this := &M{}
	this.A = that.GetA()
	return this
}

func (this *M) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filedotname.M{")
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringFileDot(this.A, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFileDot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringFileDot(m github_com_maditya_protobuf_proto.Message) string {
	e := github_com_maditya_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func NewPopulatedM(r randyFileDot, easy bool) *M {
	this := &M{}
	if r.Intn(10) != 0 {
		v1 := randStringFileDot(r)
		this.A = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFileDot(r, 2)
	}
	return this
}

type randyFileDot interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFileDot(r randyFileDot) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFileDot(r randyFileDot) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneFileDot(r)
	}
	return string(tmps)
}
func randUnrecognizedFileDot(r randyFileDot, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldFileDot(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldFileDot(data []byte, r randyFileDot, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateFileDot(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateFileDot(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateFileDot(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateFileDot(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *M) Size() (n int) {
	var l int
	_ = l
	if m.A != nil {
		l = len(*m.A)
		n += 1 + l + sovFileDot(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileDot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFileDot(x uint64) (n int) {
	return sovFileDot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *M) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&M{`,
		`A:` + valueToStringFileDot(this.A) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFileDot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() { proto.RegisterFile("file.dot.proto", fileDescriptorFileDot) }

var fileDescriptorFileDot = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x24, 0xcb, 0xaf, 0x6e, 0xc2, 0x50,
	0x1c, 0xc5, 0xf1, 0xdf, 0x91, 0xeb, 0x96, 0x25, 0xab, 0x5a, 0x26, 0x4e, 0x96, 0xa9, 0xa9, 0x76,
	0xcf, 0x30, 0xbd, 0x19, 0xde, 0xe0, 0x96, 0xfe, 0xa1, 0x09, 0xe5, 0x12, 0x72, 0x2b, 0x70, 0x7d,
	0x1c, 0x24, 0x12, 0x89, 0xac, 0xac, 0x44, 0xf6, 0xfe, 0x30, 0x95, 0x95, 0x95, 0x84, 0xe2, 0xce,
	0x27, 0x39, 0xdf, 0xe0, 0x35, 0x2f, 0xd7, 0x59, 0x94, 0x5a, 0x17, 0x6d, 0x77, 0xd6, 0xd9, 0xf0,
	0xf9, 0xee, 0xd4, 0xba, 0x8d, 0xa9, 0xb2, 0x8f, 0x9f, 0xa2, 0x74, 0xab, 0x3a, 0x89, 0x96, 0xb6,
	0x8a, 0x2b, 0x93, 0x96, 0x6e, 0x6f, 0xe2, 0xf9, 0x96, 0xd4, 0x79, 0x5c, 0xd8, 0xc2, 0xce, 0x98,
	0xd7, 0x23, 0xff, 0x7a, 0x0b, 0xf0, 0x1f, 0xbe, 0x04, 0x30, 0xef, 0xf8, 0xc4, 0xf7, 0xd3, 0x02,
	0xe6, 0xf7, 0xaf, 0xf5, 0x94, 0xce, 0x53, 0x2e, 0x9e, 0xd2, 0x7b, 0x62, 0xf0, 0xc4, 0xe8, 0x89,
	0xc9, 0x13, 0x8d, 0x12, 0x07, 0x25, 0x8e, 0x4a, 0x9c, 0x94, 0x38, 0x2b, 0xd1, 0x2a, 0xa5, 0x53,
	0x4a, 0xaf, 0xc4, 0xa0, 0x94, 0x51, 0x89, 0x49, 0x29, 0xcd, 0x95, 0x72, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xd6, 0x96, 0xb4, 0x1a, 0xb0, 0x00, 0x00, 0x00,
}

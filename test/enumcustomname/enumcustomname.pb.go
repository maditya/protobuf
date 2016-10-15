// Code generated by protoc-gen-gogo.
// source: enumcustomname.proto
// DO NOT EDIT!

/*
	Package enumcustomname is a generated protocol buffer package.

	Package enumcustomname tests the behavior of enum_customname and
	enumvalue_customname extensions.

	It is generated from these files:
		enumcustomname.proto

	It has these top-level messages:
		OnlyEnums
*/
package enumcustomname

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"
import test "github.com/maditya/protobuf/test"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyCustomEnum int32

const (
	// The following field will take on the custom name and the prefix, joined
	// by an underscore.
	MyCustomEnum_MyBetterNameA MyCustomEnum = 0
	MyCustomEnum_B             MyCustomEnum = 1
)

var MyCustomEnum_name = map[int32]string{
	0: "A",
	1: "B",
}
var MyCustomEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
}

func (x MyCustomEnum) Enum() *MyCustomEnum {
	p := new(MyCustomEnum)
	*p = x
	return p
}
func (x MyCustomEnum) String() string {
	return proto.EnumName(MyCustomEnum_name, int32(x))
}
func (x *MyCustomEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomEnum_value, data, "MyCustomEnum")
	if err != nil {
		return err
	}
	*x = MyCustomEnum(value)
	return nil
}
func (MyCustomEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnumcustomname, []int{0} }

type MyCustomUnprefixedEnum int32

const (
	MyBetterNameUnprefixedA MyCustomUnprefixedEnum = 0
	UNPREFIXED_B            MyCustomUnprefixedEnum = 1
)

var MyCustomUnprefixedEnum_name = map[int32]string{
	0: "UNPREFIXED_A",
	1: "UNPREFIXED_B",
}
var MyCustomUnprefixedEnum_value = map[string]int32{
	"UNPREFIXED_A": 0,
	"UNPREFIXED_B": 1,
}

func (x MyCustomUnprefixedEnum) Enum() *MyCustomUnprefixedEnum {
	p := new(MyCustomUnprefixedEnum)
	*p = x
	return p
}
func (x MyCustomUnprefixedEnum) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyCustomUnprefixedEnum_name, int32(x))
}
func (x *MyCustomUnprefixedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomUnprefixedEnum_value, data, "MyCustomUnprefixedEnum")
	if err != nil {
		return err
	}
	*x = MyCustomUnprefixedEnum(value)
	return nil
}
func (MyCustomUnprefixedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnumcustomname, []int{1}
}

type MyEnumWithEnumStringer int32

const (
	MyEnumWithEnumStringer_EnumValueStringerA MyEnumWithEnumStringer = 0
	MyEnumWithEnumStringer_STRINGER_B         MyEnumWithEnumStringer = 1
)

var MyEnumWithEnumStringer_name = map[int32]string{
	0: "STRINGER_A",
	1: "STRINGER_B",
}
var MyEnumWithEnumStringer_value = map[string]int32{
	"STRINGER_A": 0,
	"STRINGER_B": 1,
}

func (x MyEnumWithEnumStringer) Enum() *MyEnumWithEnumStringer {
	p := new(MyEnumWithEnumStringer)
	*p = x
	return p
}
func (x MyEnumWithEnumStringer) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyEnumWithEnumStringer_name, int32(x))
}
func (x *MyEnumWithEnumStringer) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyEnumWithEnumStringer_value, data, "MyEnumWithEnumStringer")
	if err != nil {
		return err
	}
	*x = MyEnumWithEnumStringer(value)
	return nil
}
func (MyEnumWithEnumStringer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnumcustomname, []int{2}
}

type OnlyEnums struct {
	MyEnum                         *MyCustomEnum               `protobuf:"varint,1,opt,name=my_enum,json=myEnum,enum=enumcustomname.MyCustomEnum" json:"my_enum,omitempty"`
	MyEnumDefaultA                 *MyCustomEnum               `protobuf:"varint,2,opt,name=my_enum_default_a,json=myEnumDefaultA,enum=enumcustomname.MyCustomEnum,def=0" json:"my_enum_default_a,omitempty"`
	MyEnumDefaultB                 *MyCustomEnum               `protobuf:"varint,3,opt,name=my_enum_default_b,json=myEnumDefaultB,enum=enumcustomname.MyCustomEnum,def=1" json:"my_enum_default_b,omitempty"`
	MyUnprefixedEnum               *MyCustomUnprefixedEnum     `protobuf:"varint,4,opt,name=my_unprefixed_enum,json=myUnprefixedEnum,enum=enumcustomname.MyCustomUnprefixedEnum" json:"my_unprefixed_enum,omitempty"`
	MyUnprefixedEnumDefaultA       *MyCustomUnprefixedEnum     `protobuf:"varint,5,opt,name=my_unprefixed_enum_default_a,json=myUnprefixedEnumDefaultA,enum=enumcustomname.MyCustomUnprefixedEnum,def=0" json:"my_unprefixed_enum_default_a,omitempty"`
	MyUnprefixedEnumDefaultB       *MyCustomUnprefixedEnum     `protobuf:"varint,6,opt,name=my_unprefixed_enum_default_b,json=myUnprefixedEnumDefaultB,enum=enumcustomname.MyCustomUnprefixedEnum,def=1" json:"my_unprefixed_enum_default_b,omitempty"`
	YetAnotherTestEnum             *test.YetAnotherTestEnum    `protobuf:"varint,7,opt,name=yet_another_test_enum,json=yetAnotherTestEnum,enum=test.YetAnotherTestEnum" json:"yet_another_test_enum,omitempty"`
	YetAnotherTestEnumDefaultAa    *test.YetAnotherTestEnum    `protobuf:"varint,8,opt,name=yet_another_test_enum_default_aa,json=yetAnotherTestEnumDefaultAa,enum=test.YetAnotherTestEnum,def=0" json:"yet_another_test_enum_default_aa,omitempty"`
	YetAnotherTestEnumDefaultBb    *test.YetAnotherTestEnum    `protobuf:"varint,9,opt,name=yet_another_test_enum_default_bb,json=yetAnotherTestEnumDefaultBb,enum=test.YetAnotherTestEnum,def=1" json:"yet_another_test_enum_default_bb,omitempty"`
	YetYetAnotherTestEnum          *test.YetYetAnotherTestEnum `protobuf:"varint,10,opt,name=yet_yet_another_test_enum,json=yetYetAnotherTestEnum,enum=test.YetYetAnotherTestEnum" json:"yet_yet_another_test_enum,omitempty"`
	YetYetAnotherTestEnumDefaultCc *test.YetYetAnotherTestEnum `protobuf:"varint,11,opt,name=yet_yet_another_test_enum_default_cc,json=yetYetAnotherTestEnumDefaultCc,enum=test.YetYetAnotherTestEnum,def=0" json:"yet_yet_another_test_enum_default_cc,omitempty"`
	YetYetAnotherTestEnumDefaultDd *test.YetYetAnotherTestEnum `protobuf:"varint,12,opt,name=yet_yet_another_test_enum_default_dd,json=yetYetAnotherTestEnumDefaultDd,enum=test.YetYetAnotherTestEnum,def=1" json:"yet_yet_another_test_enum_default_dd,omitempty"`
	XXX_unrecognized               []byte                      `json:"-"`
}

func (m *OnlyEnums) Reset()                    { *m = OnlyEnums{} }
func (m *OnlyEnums) String() string            { return proto.CompactTextString(m) }
func (*OnlyEnums) ProtoMessage()               {}
func (*OnlyEnums) Descriptor() ([]byte, []int) { return fileDescriptorEnumcustomname, []int{0} }

const Default_OnlyEnums_MyEnumDefaultA MyCustomEnum = MyCustomEnum_MyBetterNameA
const Default_OnlyEnums_MyEnumDefaultB MyCustomEnum = MyCustomEnum_B
const Default_OnlyEnums_MyUnprefixedEnumDefaultA MyCustomUnprefixedEnum = MyBetterNameUnprefixedA
const Default_OnlyEnums_MyUnprefixedEnumDefaultB MyCustomUnprefixedEnum = UNPREFIXED_B
const Default_OnlyEnums_YetAnotherTestEnumDefaultAa test.YetAnotherTestEnum = test.AA
const Default_OnlyEnums_YetAnotherTestEnumDefaultBb test.YetAnotherTestEnum = test.BetterYetBB
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_CC
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_BetterYetDD

func (m *OnlyEnums) GetMyEnum() MyCustomEnum {
	if m != nil && m.MyEnum != nil {
		return *m.MyEnum
	}
	return MyCustomEnum_MyBetterNameA
}

func (m *OnlyEnums) GetMyEnumDefaultA() MyCustomEnum {
	if m != nil && m.MyEnumDefaultA != nil {
		return *m.MyEnumDefaultA
	}
	return Default_OnlyEnums_MyEnumDefaultA
}

func (m *OnlyEnums) GetMyEnumDefaultB() MyCustomEnum {
	if m != nil && m.MyEnumDefaultB != nil {
		return *m.MyEnumDefaultB
	}
	return Default_OnlyEnums_MyEnumDefaultB
}

func (m *OnlyEnums) GetMyUnprefixedEnum() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnum != nil {
		return *m.MyUnprefixedEnum
	}
	return MyBetterNameUnprefixedA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultA() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultA != nil {
		return *m.MyUnprefixedEnumDefaultA
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultB() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultB != nil {
		return *m.MyUnprefixedEnumDefaultB
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultB
}

func (m *OnlyEnums) GetYetAnotherTestEnum() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnum != nil {
		return *m.YetAnotherTestEnum
	}
	return test.AA
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultAa() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultAa != nil {
		return *m.YetAnotherTestEnumDefaultAa
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultAa
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultBb() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultBb != nil {
		return *m.YetAnotherTestEnumDefaultBb
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultBb
}

func (m *OnlyEnums) GetYetYetAnotherTestEnum() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnum != nil {
		return *m.YetYetAnotherTestEnum
	}
	return test.YetYetAnotherTestEnum_CC
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultCc() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultCc != nil {
		return *m.YetYetAnotherTestEnumDefaultCc
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultDd() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultDd != nil {
		return *m.YetYetAnotherTestEnumDefaultDd
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd
}

func init() {
	proto.RegisterType((*OnlyEnums)(nil), "enumcustomname.OnlyEnums")
	proto.RegisterEnum("enumcustomname.MyCustomEnum", MyCustomEnum_name, MyCustomEnum_value)
	proto.RegisterEnum("enumcustomname.MyCustomUnprefixedEnum", MyCustomUnprefixedEnum_name, MyCustomUnprefixedEnum_value)
	proto.RegisterEnum("enumcustomname.MyEnumWithEnumStringer", MyEnumWithEnumStringer_name, MyEnumWithEnumStringer_value)
}
func (x MyEnumWithEnumStringer) String() string {
	s, ok := MyEnumWithEnumStringer_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("enumcustomname.proto", fileDescriptorEnumcustomname) }

var fileDescriptorEnumcustomname = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x5d, 0x8f, 0xd2, 0x4e,
	0x14, 0xc6, 0x29, 0xff, 0xff, 0xb2, 0xec, 0x88, 0xa4, 0x3b, 0x51, 0x1c, 0xc1, 0x34, 0xcd, 0xc6,
	0x18, 0x43, 0x22, 0x18, 0x2f, 0xf1, 0x6a, 0x4a, 0xd1, 0x6c, 0x0c, 0x68, 0xba, 0x8b, 0x6f, 0x37,
	0xcd, 0xb4, 0x1d, 0x5e, 0x12, 0xa6, 0xdd, 0x94, 0x69, 0xb4, 0xdf, 0xc0, 0xf0, 0x1d, 0xb8, 0x92,
	0x0b, 0x2f, 0xbd, 0xde, 0x6b, 0x3f, 0x98, 0x99, 0xe9, 0xc2, 0xf2, 0xd2, 0x2d, 0xc4, 0xab, 0x81,
	0x93, 0xe7, 0x3c, 0xbf, 0x73, 0x9e, 0x9c, 0x82, 0x07, 0xd4, 0x8f, 0x98, 0x1b, 0x4d, 0x79, 0xc0,
	0x7c, 0xc2, 0x68, 0xe3, 0x2a, 0x0c, 0x78, 0x00, 0xcb, 0x9b, 0xd5, 0xea, 0xcb, 0xe1, 0x98, 0x8f,
	0x22, 0xa7, 0xe1, 0x06, 0xac, 0xc9, 0x88, 0x37, 0xe6, 0x31, 0x69, 0x4a, 0xa5, 0x13, 0x0d, 0x9a,
	0xc3, 0x60, 0x18, 0xc8, 0x3f, 0xf2, 0x57, 0xe2, 0x50, 0x6d, 0x64, 0x75, 0x70, 0x3a, 0xe5, 0x4d,
	0x3e, 0xa2, 0xe2, 0x4d, 0xf4, 0x67, 0x7f, 0x8a, 0xe0, 0xe4, 0xbd, 0x3f, 0x89, 0x3b, 0x7e, 0xc4,
	0xa6, 0xb0, 0x09, 0x8e, 0x59, 0x6c, 0x8b, 0x21, 0x90, 0xa2, 0x2b, 0xcf, 0xcb, 0xaf, 0x2a, 0x8d,
	0xad, 0x39, 0xbb, 0x52, 0x69, 0x15, 0x98, 0x7c, 0xa1, 0x09, 0x4e, 0x6f, 0x1a, 0x6c, 0x8f, 0x0e,
	0x48, 0x34, 0xe1, 0x36, 0x41, 0xf9, 0xac, 0xd6, 0x96, 0x82, 0xad, 0x72, 0xd2, 0x6d, 0x26, 0x1d,
	0x38, 0xcd, 0xc5, 0x41, 0xff, 0x65, 0xbb, 0x18, 0x5b, 0x2e, 0x06, 0xec, 0x01, 0xc8, 0x62, 0x3b,
	0xf2, 0xaf, 0x42, 0x3a, 0x18, 0x7f, 0xa7, 0x5e, 0xb2, 0xc7, 0xff, 0xd2, 0x46, 0xdf, 0xb5, 0xe9,
	0xaf, 0x84, 0x72, 0x23, 0x95, 0x6d, 0x55, 0xa0, 0x0f, 0x9e, 0xec, 0xfa, 0xad, 0xad, 0x79, 0x74,
	0x98, 0x73, 0xab, 0xd4, 0xef, 0x7d, 0xb0, 0x3a, 0x6f, 0xce, 0x3f, 0x77, 0x4c, 0x1b, 0x5b, 0x68,
	0x9b, 0xb3, 0x4a, 0x21, 0x9b, 0xe7, 0xa0, 0xc2, 0x3f, 0xf0, 0x8c, 0x3b, 0x79, 0x06, 0x7c, 0x07,
	0x1e, 0xc6, 0x94, 0xdb, 0xc4, 0x0f, 0xf8, 0x88, 0x86, 0xb6, 0x38, 0x8a, 0x24, 0xb2, 0x63, 0x09,
	0x42, 0x0d, 0x79, 0x26, 0x5f, 0x28, 0xc7, 0x89, 0xe2, 0x92, 0x4e, 0xb9, 0x8c, 0x0a, 0xc6, 0x3b,
	0x35, 0xe8, 0x02, 0x3d, 0xd5, 0xec, 0x36, 0x2f, 0x82, 0x8a, 0xd9, 0xbe, 0xad, 0x3c, 0xc6, 0x56,
	0x6d, 0xd7, 0x7b, 0x19, 0x10, 0xd9, 0x0f, 0x71, 0x1c, 0x74, 0xb2, 0x0f, 0x62, 0x18, 0x19, 0x10,
	0xc3, 0x81, 0x7d, 0xf0, 0x58, 0x40, 0xd2, 0xa3, 0x01, 0xd2, 0xbd, 0xb6, 0x72, 0x4f, 0x49, 0x47,
	0x84, 0xba, 0x5b, 0x86, 0x0c, 0x3c, 0xbd, 0xd3, 0x76, 0x35, 0xbf, 0xeb, 0xa2, 0x7b, 0x7b, 0x09,
	0xad, 0x7c, 0xbb, 0x6d, 0x69, 0xa9, 0x94, 0x9b, 0x2d, 0xda, 0xee, 0x61, 0x38, 0xcf, 0x43, 0xa5,
	0x03, 0x70, 0xa6, 0x99, 0x8d, 0x33, 0xbd, 0xfa, 0x6b, 0x50, 0x48, 0x3e, 0x4c, 0x88, 0x80, 0x82,
	0xd5, 0x5c, 0xf5, 0x74, 0x36, 0xd7, 0xef, 0x77, 0x63, 0x83, 0x72, 0x4e, 0xc3, 0x1e, 0x61, 0x14,
	0xc3, 0x23, 0xa0, 0x18, 0xaa, 0x52, 0x55, 0xaf, 0x17, 0x5a, 0xa9, 0x1b, 0xb7, 0xe5, 0x05, 0x8b,
	0x96, 0xfa, 0x37, 0xa0, 0x6e, 0x1f, 0x31, 0x7c, 0x01, 0x36, 0x3e, 0x1b, 0x35, 0x57, 0xad, 0xcd,
	0xe6, 0xfa, 0xa3, 0x75, 0xc7, 0xdb, 0x0e, 0x0c, 0xd5, 0x0d, 0xb9, 0xc0, 0x9c, 0xfd, 0xf8, 0xa9,
	0xe5, 0x7e, 0x2d, 0xb4, 0xdc, 0xf5, 0x42, 0xab, 0x2c, 0x71, 0x9b, 0x90, 0xfa, 0x57, 0x50, 0x49,
	0xa6, 0xfe, 0x34, 0xe6, 0x23, 0xf1, 0x5e, 0xf0, 0x70, 0xec, 0x0f, 0x69, 0x08, 0x9f, 0x01, 0x70,
	0x71, 0x69, 0x9d, 0xf7, 0xde, 0x76, 0x2c, 0x09, 0xaf, 0xcc, 0xe6, 0x3a, 0x14, 0x8a, 0x8f, 0x64,
	0x12, 0xd1, 0xa5, 0x0c, 0xc3, 0xf2, 0x9a, 0x4e, 0x50, 0x8b, 0x82, 0xf8, 0x7b, 0xa1, 0x29, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x28, 0x5b, 0x7d, 0x84, 0xe1, 0x05, 0x00, 0x00,
}

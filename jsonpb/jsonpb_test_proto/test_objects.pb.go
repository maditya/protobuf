// Code generated by protoc-gen-gogo.
// source: test_objects.proto
// DO NOT EDIT!

package jsonpb

import proto "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/maditya/protobuf/types"
import google_protobuf1 "github.com/maditya/protobuf/types"
import google_protobuf2 "github.com/maditya/protobuf/types"
import google_protobuf3 "github.com/maditya/protobuf/types"
import google_protobuf4 "github.com/maditya/protobuf/types"

// skipping weak import gogoproto "gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Widget_Color int32

const (
	Widget_RED   Widget_Color = 0
	Widget_GREEN Widget_Color = 1
	Widget_BLUE  Widget_Color = 2
)

var Widget_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Widget_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Widget_Color) Enum() *Widget_Color {
	p := new(Widget_Color)
	*p = x
	return p
}
func (x Widget_Color) String() string {
	return proto.EnumName(Widget_Color_name, int32(x))
}
func (x *Widget_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Widget_Color_value, data, "Widget_Color")
	if err != nil {
		return err
	}
	*x = Widget_Color(value)
	return nil
}
func (Widget_Color) EnumDescriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{2, 0} }

// Test message for holding primitive types.
type Simple struct {
	OBool            *bool    `protobuf:"varint,1,opt,name=o_bool,json=oBool" json:"o_bool,omitempty"`
	OInt32           *int32   `protobuf:"varint,2,opt,name=o_int32,json=oInt32" json:"o_int32,omitempty"`
	OInt64           *int64   `protobuf:"varint,3,opt,name=o_int64,json=oInt64" json:"o_int64,omitempty"`
	OUint32          *uint32  `protobuf:"varint,4,opt,name=o_uint32,json=oUint32" json:"o_uint32,omitempty"`
	OUint64          *uint64  `protobuf:"varint,5,opt,name=o_uint64,json=oUint64" json:"o_uint64,omitempty"`
	OSint32          *int32   `protobuf:"zigzag32,6,opt,name=o_sint32,json=oSint32" json:"o_sint32,omitempty"`
	OSint64          *int64   `protobuf:"zigzag64,7,opt,name=o_sint64,json=oSint64" json:"o_sint64,omitempty"`
	OFloat           *float32 `protobuf:"fixed32,8,opt,name=o_float,json=oFloat" json:"o_float,omitempty"`
	ODouble          *float64 `protobuf:"fixed64,9,opt,name=o_double,json=oDouble" json:"o_double,omitempty"`
	OString          *string  `protobuf:"bytes,10,opt,name=o_string,json=oString" json:"o_string,omitempty"`
	OBytes           []byte   `protobuf:"bytes,11,opt,name=o_bytes,json=oBytes" json:"o_bytes,omitempty"`
	OCastBytes       Bytes    `protobuf:"bytes,12,opt,name=o_cast_bytes,json=oCastBytes,casttype=Bytes" json:"o_cast_bytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Simple) Reset()                    { *m = Simple{} }
func (m *Simple) String() string            { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()               {}
func (*Simple) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{0} }

func (m *Simple) GetOBool() bool {
	if m != nil && m.OBool != nil {
		return *m.OBool
	}
	return false
}

func (m *Simple) GetOInt32() int32 {
	if m != nil && m.OInt32 != nil {
		return *m.OInt32
	}
	return 0
}

func (m *Simple) GetOInt64() int64 {
	if m != nil && m.OInt64 != nil {
		return *m.OInt64
	}
	return 0
}

func (m *Simple) GetOUint32() uint32 {
	if m != nil && m.OUint32 != nil {
		return *m.OUint32
	}
	return 0
}

func (m *Simple) GetOUint64() uint64 {
	if m != nil && m.OUint64 != nil {
		return *m.OUint64
	}
	return 0
}

func (m *Simple) GetOSint32() int32 {
	if m != nil && m.OSint32 != nil {
		return *m.OSint32
	}
	return 0
}

func (m *Simple) GetOSint64() int64 {
	if m != nil && m.OSint64 != nil {
		return *m.OSint64
	}
	return 0
}

func (m *Simple) GetOFloat() float32 {
	if m != nil && m.OFloat != nil {
		return *m.OFloat
	}
	return 0
}

func (m *Simple) GetODouble() float64 {
	if m != nil && m.ODouble != nil {
		return *m.ODouble
	}
	return 0
}

func (m *Simple) GetOString() string {
	if m != nil && m.OString != nil {
		return *m.OString
	}
	return ""
}

func (m *Simple) GetOBytes() []byte {
	if m != nil {
		return m.OBytes
	}
	return nil
}

func (m *Simple) GetOCastBytes() Bytes {
	if m != nil {
		return m.OCastBytes
	}
	return nil
}

// Test message for holding repeated primitives.
type Repeats struct {
	RBool            []bool    `protobuf:"varint,1,rep,name=r_bool,json=rBool" json:"r_bool,omitempty"`
	RInt32           []int32   `protobuf:"varint,2,rep,name=r_int32,json=rInt32" json:"r_int32,omitempty"`
	RInt64           []int64   `protobuf:"varint,3,rep,name=r_int64,json=rInt64" json:"r_int64,omitempty"`
	RUint32          []uint32  `protobuf:"varint,4,rep,name=r_uint32,json=rUint32" json:"r_uint32,omitempty"`
	RUint64          []uint64  `protobuf:"varint,5,rep,name=r_uint64,json=rUint64" json:"r_uint64,omitempty"`
	RSint32          []int32   `protobuf:"zigzag32,6,rep,name=r_sint32,json=rSint32" json:"r_sint32,omitempty"`
	RSint64          []int64   `protobuf:"zigzag64,7,rep,name=r_sint64,json=rSint64" json:"r_sint64,omitempty"`
	RFloat           []float32 `protobuf:"fixed32,8,rep,name=r_float,json=rFloat" json:"r_float,omitempty"`
	RDouble          []float64 `protobuf:"fixed64,9,rep,name=r_double,json=rDouble" json:"r_double,omitempty"`
	RString          []string  `protobuf:"bytes,10,rep,name=r_string,json=rString" json:"r_string,omitempty"`
	RBytes           [][]byte  `protobuf:"bytes,11,rep,name=r_bytes,json=rBytes" json:"r_bytes,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Repeats) Reset()                    { *m = Repeats{} }
func (m *Repeats) String() string            { return proto.CompactTextString(m) }
func (*Repeats) ProtoMessage()               {}
func (*Repeats) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{1} }

func (m *Repeats) GetRBool() []bool {
	if m != nil {
		return m.RBool
	}
	return nil
}

func (m *Repeats) GetRInt32() []int32 {
	if m != nil {
		return m.RInt32
	}
	return nil
}

func (m *Repeats) GetRInt64() []int64 {
	if m != nil {
		return m.RInt64
	}
	return nil
}

func (m *Repeats) GetRUint32() []uint32 {
	if m != nil {
		return m.RUint32
	}
	return nil
}

func (m *Repeats) GetRUint64() []uint64 {
	if m != nil {
		return m.RUint64
	}
	return nil
}

func (m *Repeats) GetRSint32() []int32 {
	if m != nil {
		return m.RSint32
	}
	return nil
}

func (m *Repeats) GetRSint64() []int64 {
	if m != nil {
		return m.RSint64
	}
	return nil
}

func (m *Repeats) GetRFloat() []float32 {
	if m != nil {
		return m.RFloat
	}
	return nil
}

func (m *Repeats) GetRDouble() []float64 {
	if m != nil {
		return m.RDouble
	}
	return nil
}

func (m *Repeats) GetRString() []string {
	if m != nil {
		return m.RString
	}
	return nil
}

func (m *Repeats) GetRBytes() [][]byte {
	if m != nil {
		return m.RBytes
	}
	return nil
}

// Test message for holding enums and nested messages.
type Widget struct {
	Color            *Widget_Color  `protobuf:"varint,1,opt,name=color,enum=jsonpb.Widget_Color" json:"color,omitempty"`
	RColor           []Widget_Color `protobuf:"varint,2,rep,name=r_color,json=rColor,enum=jsonpb.Widget_Color" json:"r_color,omitempty"`
	Simple           *Simple        `protobuf:"bytes,10,opt,name=simple" json:"simple,omitempty"`
	RSimple          []*Simple      `protobuf:"bytes,11,rep,name=r_simple,json=rSimple" json:"r_simple,omitempty"`
	Repeats          *Repeats       `protobuf:"bytes,20,opt,name=repeats" json:"repeats,omitempty"`
	RRepeats         []*Repeats     `protobuf:"bytes,21,rep,name=r_repeats,json=rRepeats" json:"r_repeats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Widget) Reset()                    { *m = Widget{} }
func (m *Widget) String() string            { return proto.CompactTextString(m) }
func (*Widget) ProtoMessage()               {}
func (*Widget) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{2} }

func (m *Widget) GetColor() Widget_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Widget_RED
}

func (m *Widget) GetRColor() []Widget_Color {
	if m != nil {
		return m.RColor
	}
	return nil
}

func (m *Widget) GetSimple() *Simple {
	if m != nil {
		return m.Simple
	}
	return nil
}

func (m *Widget) GetRSimple() []*Simple {
	if m != nil {
		return m.RSimple
	}
	return nil
}

func (m *Widget) GetRepeats() *Repeats {
	if m != nil {
		return m.Repeats
	}
	return nil
}

func (m *Widget) GetRRepeats() []*Repeats {
	if m != nil {
		return m.RRepeats
	}
	return nil
}

type Maps struct {
	MInt64Str        map[int64]string `protobuf:"bytes,1,rep,name=m_int64_str,json=mInt64Str" json:"m_int64_str,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MBoolSimple      map[bool]*Simple `protobuf:"bytes,2,rep,name=m_bool_simple,json=mBoolSimple" json:"m_bool_simple,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Maps) Reset()                    { *m = Maps{} }
func (m *Maps) String() string            { return proto.CompactTextString(m) }
func (*Maps) ProtoMessage()               {}
func (*Maps) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{3} }

func (m *Maps) GetMInt64Str() map[int64]string {
	if m != nil {
		return m.MInt64Str
	}
	return nil
}

func (m *Maps) GetMBoolSimple() map[bool]*Simple {
	if m != nil {
		return m.MBoolSimple
	}
	return nil
}

type MsgWithOneof struct {
	// Types that are valid to be assigned to Union:
	//	*MsgWithOneof_Title
	//	*MsgWithOneof_Salary
	//	*MsgWithOneof_Country
	//	*MsgWithOneof_HomeAddress
	Union            isMsgWithOneof_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *MsgWithOneof) Reset()                    { *m = MsgWithOneof{} }
func (m *MsgWithOneof) String() string            { return proto.CompactTextString(m) }
func (*MsgWithOneof) ProtoMessage()               {}
func (*MsgWithOneof) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{4} }

type isMsgWithOneof_Union interface {
	isMsgWithOneof_Union()
}

type MsgWithOneof_Title struct {
	Title string `protobuf:"bytes,1,opt,name=title,oneof"`
}
type MsgWithOneof_Salary struct {
	Salary int64 `protobuf:"varint,2,opt,name=salary,oneof"`
}
type MsgWithOneof_Country struct {
	Country string `protobuf:"bytes,3,opt,name=Country,json=country,oneof"`
}
type MsgWithOneof_HomeAddress struct {
	HomeAddress string `protobuf:"bytes,4,opt,name=home_address,json=homeAddress,oneof"`
}

func (*MsgWithOneof_Title) isMsgWithOneof_Union()       {}
func (*MsgWithOneof_Salary) isMsgWithOneof_Union()      {}
func (*MsgWithOneof_Country) isMsgWithOneof_Union()     {}
func (*MsgWithOneof_HomeAddress) isMsgWithOneof_Union() {}

func (m *MsgWithOneof) GetUnion() isMsgWithOneof_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *MsgWithOneof) GetTitle() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Title); ok {
		return x.Title
	}
	return ""
}

func (m *MsgWithOneof) GetSalary() int64 {
	if x, ok := m.GetUnion().(*MsgWithOneof_Salary); ok {
		return x.Salary
	}
	return 0
}

func (m *MsgWithOneof) GetCountry() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Country); ok {
		return x.Country
	}
	return ""
}

func (m *MsgWithOneof) GetHomeAddress() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_HomeAddress); ok {
		return x.HomeAddress
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MsgWithOneof) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MsgWithOneof_OneofMarshaler, _MsgWithOneof_OneofUnmarshaler, _MsgWithOneof_OneofSizer, []interface{}{
		(*MsgWithOneof_Title)(nil),
		(*MsgWithOneof_Salary)(nil),
		(*MsgWithOneof_Country)(nil),
		(*MsgWithOneof_HomeAddress)(nil),
	}
}

func _MsgWithOneof_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Title)
	case *MsgWithOneof_Salary:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Salary))
	case *MsgWithOneof_Country:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Country)
	case *MsgWithOneof_HomeAddress:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.HomeAddress)
	case nil:
	default:
		return fmt.Errorf("MsgWithOneof.Union has unexpected type %T", x)
	}
	return nil
}

func _MsgWithOneof_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MsgWithOneof)
	switch tag {
	case 1: // union.title
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Title{x}
		return true, err
	case 2: // union.salary
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &MsgWithOneof_Salary{int64(x)}
		return true, err
	case 3: // union.Country
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Country{x}
		return true, err
	case 4: // union.home_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_HomeAddress{x}
		return true, err
	default:
		return false, nil
	}
}

func _MsgWithOneof_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Title)))
		n += len(x.Title)
	case *MsgWithOneof_Salary:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Salary))
	case *MsgWithOneof_Country:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Country)))
		n += len(x.Country)
	case *MsgWithOneof_HomeAddress:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.HomeAddress)))
		n += len(x.HomeAddress)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Real struct {
	Value                        *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Real) Reset()                    { *m = Real{} }
func (m *Real) String() string            { return proto.CompactTextString(m) }
func (*Real) ProtoMessage()               {}
func (*Real) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{5} }

var extRange_Real = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Real) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Real
}

func (m *Real) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Complex struct {
	Imaginary                    *float64 `protobuf:"fixed64,1,opt,name=imaginary" json:"imaginary,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Complex) Reset()                    { *m = Complex{} }
func (m *Complex) String() string            { return proto.CompactTextString(m) }
func (*Complex) ProtoMessage()               {}
func (*Complex) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{6} }

var extRange_Complex = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Complex) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Complex
}

func (m *Complex) GetImaginary() float64 {
	if m != nil && m.Imaginary != nil {
		return *m.Imaginary
	}
	return 0
}

var E_Complex_RealExtension = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*Complex)(nil),
	Field:         123,
	Name:          "jsonpb.Complex.real_extension",
	Tag:           "bytes,123,opt,name=real_extension,json=realExtension",
}

type KnownTypes struct {
	An               *google_protobuf.Any          `protobuf:"bytes,14,opt,name=an" json:"an,omitempty"`
	Dur              *google_protobuf1.Duration    `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	St               *google_protobuf2.Struct      `protobuf:"bytes,12,opt,name=st" json:"st,omitempty"`
	Ts               *google_protobuf3.Timestamp   `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Dbl              *google_protobuf4.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt              *google_protobuf4.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64              *google_protobuf4.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64              *google_protobuf4.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32              *google_protobuf4.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32              *google_protobuf4.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool             *google_protobuf4.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str              *google_protobuf4.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes            *google_protobuf4.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *KnownTypes) Reset()                    { *m = KnownTypes{} }
func (m *KnownTypes) String() string            { return proto.CompactTextString(m) }
func (*KnownTypes) ProtoMessage()               {}
func (*KnownTypes) Descriptor() ([]byte, []int) { return fileDescriptorTestObjects, []int{7} }

func (m *KnownTypes) GetAn() *google_protobuf.Any {
	if m != nil {
		return m.An
	}
	return nil
}

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

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*Real)(nil),
	ExtensionType: (*string)(nil),
	Field:         124,
	Name:          "jsonpb.name",
	Tag:           "bytes,124,opt,name=name",
}

func init() {
	proto.RegisterType((*Simple)(nil), "jsonpb.Simple")
	proto.RegisterType((*Repeats)(nil), "jsonpb.Repeats")
	proto.RegisterType((*Widget)(nil), "jsonpb.Widget")
	proto.RegisterType((*Maps)(nil), "jsonpb.Maps")
	proto.RegisterType((*MsgWithOneof)(nil), "jsonpb.MsgWithOneof")
	proto.RegisterType((*Real)(nil), "jsonpb.Real")
	proto.RegisterType((*Complex)(nil), "jsonpb.Complex")
	proto.RegisterType((*KnownTypes)(nil), "jsonpb.KnownTypes")
	proto.RegisterEnum("jsonpb.Widget_Color", Widget_Color_name, Widget_Color_value)
	proto.RegisterExtension(E_Complex_RealExtension)
	proto.RegisterExtension(E_Name)
}

func init() { proto.RegisterFile("test_objects.proto", fileDescriptorTestObjects) }

var fileDescriptorTestObjects = []byte{
	// 1100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x95, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xc7, 0x2b, 0xc9, 0xf2, 0xc7, 0xda, 0x0d, 0x66, 0x27, 0xa5, 0x8a, 0x09, 0x54, 0x63, 0x4a,
	0x11, 0x2d, 0x75, 0x06, 0xc5, 0xe3, 0x61, 0x0a, 0x37, 0xf9, 0x30, 0x94, 0x81, 0x94, 0x99, 0x4d,
	0x43, 0xb9, 0xf3, 0xc8, 0xf1, 0xc6, 0x55, 0x91, 0xb5, 0x9e, 0xdd, 0x15, 0xa9, 0x07, 0x2e, 0x72,
	0xcd, 0x35, 0xcf, 0xc0, 0x23, 0xf0, 0x20, 0x3c, 0x00, 0x57, 0x3c, 0x05, 0x57, 0xcc, 0x39, 0x2b,
	0x59, 0x89, 0x1d, 0x5f, 0x59, 0xbb, 0xe7, 0x7f, 0xfe, 0x5e, 0xfd, 0xf6, 0xe8, 0x1c, 0x42, 0x35,
	0x57, 0x7a, 0x24, 0xc6, 0x6f, 0xf8, 0xb9, 0x56, 0xbd, 0xb9, 0x14, 0x5a, 0xd0, 0xea, 0x1b, 0x25,
	0xd2, 0xf9, 0xb8, 0xb3, 0x33, 0x15, 0x62, 0x9a, 0xf0, 0x3d, 0xdc, 0x1d, 0x67, 0x17, 0x7b, 0x51,
	0xba, 0x30, 0x92, 0xce, 0x87, 0xab, 0xa1, 0x49, 0x26, 0x23, 0x1d, 0x8b, 0x34, 0x8f, 0xef, 0xae,
	0xc6, 0x95, 0x96, 0xd9, 0xb9, 0xce, 0xa3, 0x0f, 0x56, 0xa3, 0x3a, 0x9e, 0x71, 0xa5, 0xa3, 0xd9,
	0x7c, 0x93, 0xfd, 0xa5, 0x8c, 0xe6, 0x73, 0x2e, 0xf3, 0x13, 0x76, 0xb6, 0xa7, 0x62, 0x2a, 0xf0,
	0x71, 0x0f, 0x9e, 0xcc, 0x6e, 0xf7, 0x6f, 0x9b, 0x54, 0x4f, 0xe3, 0xd9, 0x3c, 0xe1, 0xf4, 0x1e,
	0xa9, 0x8a, 0xd1, 0x58, 0x88, 0xc4, 0xb3, 0x7c, 0x2b, 0xa8, 0x33, 0x57, 0x1c, 0x0a, 0x91, 0xd0,
	0xfb, 0xa4, 0x26, 0x46, 0x71, 0xaa, 0xf7, 0x43, 0xcf, 0xf6, 0xad, 0xc0, 0x65, 0x55, 0xf1, 0x2d,
	0xac, 0x96, 0x81, 0x41, 0xdf, 0x73, 0x7c, 0x2b, 0x70, 0x4c, 0x60, 0xd0, 0xa7, 0x3b, 0xa4, 0x2e,
	0x46, 0x99, 0x49, 0xa9, 0xf8, 0x56, 0x70, 0x97, 0xd5, 0xc4, 0x19, 0x2e, 0xcb, 0xd0, 0xa0, 0xef,
	0xb9, 0xbe, 0x15, 0x54, 0xf2, 0x50, 0x91, 0xa5, 0x4c, 0x56, 0xd5, 0xb7, 0x82, 0x77, 0x59, 0x4d,
	0x9c, 0x5e, 0xcb, 0x52, 0x26, 0xab, 0xe6, 0x5b, 0x01, 0xcd, 0x43, 0x83, 0xbe, 0x39, 0xc4, 0x45,
	0x22, 0x22, 0xed, 0xd5, 0x7d, 0x2b, 0xb0, 0x59, 0x55, 0x7c, 0x0d, 0x2b, 0x93, 0x33, 0x11, 0xd9,
	0x38, 0xe1, 0x5e, 0xc3, 0xb7, 0x02, 0x8b, 0xd5, 0xc4, 0x31, 0x2e, 0x73, 0x3b, 0x2d, 0xe3, 0x74,
	0xea, 0x11, 0xdf, 0x0a, 0x1a, 0x60, 0x87, 0x4b, 0x63, 0x37, 0x5e, 0x68, 0xae, 0xbc, 0xa6, 0x6f,
	0x05, 0x2d, 0x56, 0x15, 0x87, 0xb0, 0xa2, 0x4f, 0x48, 0x4b, 0x8c, 0xce, 0x23, 0xa5, 0xf3, 0x68,
	0x0b, 0xa2, 0x87, 0x8d, 0xff, 0xfe, 0x79, 0xe0, 0xa2, 0x80, 0x11, 0x71, 0x14, 0x29, 0x8d, 0xcf,
	0xdd, 0x3f, 0x6d, 0x52, 0x63, 0x7c, 0xce, 0x23, 0xad, 0x80, 0xaa, 0x2c, 0xa8, 0x3a, 0x40, 0x55,
	0x16, 0x54, 0xe5, 0x92, 0xaa, 0x03, 0x54, 0xe5, 0x92, 0xaa, 0x5c, 0x52, 0x75, 0x80, 0xaa, 0x5c,
	0x52, 0x95, 0x25, 0x55, 0x07, 0xa8, 0xca, 0x92, 0xaa, 0x2c, 0xa9, 0x3a, 0x40, 0x55, 0x96, 0x54,
	0x65, 0x49, 0xd5, 0x01, 0xaa, 0xf2, 0xf4, 0x5a, 0xd6, 0x92, 0xaa, 0x03, 0x54, 0x65, 0x49, 0x55,
	0x2e, 0xa9, 0x3a, 0x40, 0x55, 0x2e, 0xa9, 0xca, 0x92, 0xaa, 0x03, 0x54, 0x65, 0x49, 0x55, 0x96,
	0x54, 0x1d, 0xa0, 0x2a, 0x4b, 0xaa, 0x72, 0x49, 0xd5, 0x01, 0xaa, 0xd2, 0x80, 0xfa, 0xcb, 0x26,
	0xd5, 0x57, 0xf1, 0x64, 0xca, 0x35, 0x7d, 0x4c, 0xdc, 0x73, 0x91, 0x08, 0x89, 0xc5, 0xb7, 0x15,
	0x6e, 0xf7, 0xcc, 0x07, 0xd5, 0x33, 0xe1, 0xde, 0x11, 0xc4, 0x98, 0x91, 0xd0, 0xa7, 0xe0, 0x67,
	0xd4, 0x00, 0x6f, 0x93, 0xba, 0x2a, 0xf1, 0x97, 0x3e, 0x22, 0x55, 0x85, 0x25, 0x8e, 0xb7, 0xdd,
	0x0c, 0xb7, 0x0a, 0xb5, 0x29, 0x7c, 0x96, 0x47, 0xe9, 0xa7, 0x06, 0x08, 0x2a, 0xe1, 0x9c, 0xeb,
	0x4a, 0x00, 0x94, 0x4b, 0x6b, 0xd2, 0x5c, 0xb0, 0xb7, 0x8d, 0x9e, 0xef, 0x14, 0xca, 0xfc, 0xde,
	0x59, 0x11, 0xa7, 0x9f, 0x91, 0x86, 0x1c, 0x15, 0xe2, 0x7b, 0x68, 0xbb, 0x26, 0xae, 0xcb, 0xfc,
	0xa9, 0xfb, 0x31, 0x71, 0xcd, 0xa1, 0x6b, 0xc4, 0x61, 0xc3, 0xe3, 0xf6, 0x1d, 0xda, 0x20, 0xee,
	0x37, 0x6c, 0x38, 0x7c, 0xd1, 0xb6, 0x68, 0x9d, 0x54, 0x0e, 0xbf, 0x3f, 0x1b, 0xb6, 0xed, 0xee,
	0x1f, 0x36, 0xa9, 0x9c, 0x44, 0x73, 0x45, 0xbf, 0x24, 0xcd, 0x99, 0x29, 0x17, 0x60, 0x8f, 0x35,
	0xd6, 0x0c, 0xdf, 0x2f, 0xfc, 0x41, 0xd2, 0x3b, 0xc1, 0xfa, 0x39, 0xd5, 0x72, 0x98, 0x6a, 0xb9,
	0x60, 0x8d, 0x59, 0xb1, 0xa6, 0x07, 0xe4, 0xee, 0x0c, 0x6b, 0xb3, 0x78, 0x6b, 0x1b, 0xd3, 0x3f,
	0xb8, 0x99, 0x0e, 0xf5, 0x6a, 0x5e, 0xdb, 0x18, 0x34, 0x67, 0xe5, 0x4e, 0xe7, 0x2b, 0xb2, 0x75,
	0xd3, 0x9f, 0xb6, 0x89, 0xf3, 0x33, 0x5f, 0xe0, 0x35, 0x3a, 0x0c, 0x1e, 0xe9, 0x36, 0x71, 0x7f,
	0x89, 0x92, 0x8c, 0x63, 0xff, 0x68, 0x30, 0xb3, 0x78, 0x66, 0x7f, 0x61, 0x75, 0x5e, 0x90, 0xf6,
	0xaa, 0xfd, 0xf5, 0xfc, 0xba, 0xc9, 0x7f, 0x78, 0x3d, 0x7f, 0xfd, 0x52, 0x4a, 0xbf, 0xee, 0xef,
	0x16, 0x69, 0x9d, 0xa8, 0xe9, 0xab, 0x58, 0xbf, 0xfe, 0x21, 0xe5, 0xe2, 0x82, 0xbe, 0x47, 0x5c,
	0x1d, 0xeb, 0x84, 0xa3, 0x5d, 0xe3, 0xf9, 0x1d, 0x66, 0x96, 0xd4, 0x23, 0x55, 0x15, 0x25, 0x91,
	0x5c, 0xa0, 0xa7, 0xf3, 0xfc, 0x0e, 0xcb, 0xd7, 0xb4, 0x43, 0x6a, 0x47, 0x22, 0x83, 0x93, 0x60,
	0x57, 0x83, 0x9c, 0xda, 0xb9, 0xd9, 0xa0, 0x1f, 0x91, 0xd6, 0x6b, 0x31, 0xe3, 0xa3, 0x68, 0x32,
	0x91, 0x5c, 0x29, 0x6c, 0x6e, 0x20, 0x68, 0xc2, 0xee, 0x81, 0xd9, 0x3c, 0xac, 0x11, 0x37, 0x4b,
	0x63, 0x91, 0x76, 0x1f, 0x91, 0x0a, 0xe3, 0x51, 0x52, 0xbe, 0xbe, 0x85, 0x6d, 0xc8, 0x2c, 0x1e,
	0xd7, 0xeb, 0x93, 0xf6, 0xd5, 0xd5, 0xd5, 0x95, 0xdd, 0xbd, 0x84, 0x7f, 0x84, 0x37, 0x79, 0x4b,
	0x77, 0x49, 0x23, 0x9e, 0x45, 0xd3, 0x38, 0x85, 0x93, 0x19, 0x79, 0xb9, 0x51, 0xa6, 0x84, 0xc7,
	0x64, 0x4b, 0xf2, 0x28, 0x19, 0xf1, 0xb7, 0x9a, 0xa7, 0x2a, 0x16, 0x29, 0x6d, 0x95, 0x25, 0x15,
	0x25, 0xde, 0xaf, 0x37, 0x6b, 0x32, 0xb7, 0x67, 0x77, 0x21, 0x69, 0x58, 0xe4, 0x74, 0xff, 0xad,
	0x10, 0xf2, 0x5d, 0x2a, 0x2e, 0xd3, 0x97, 0x8b, 0x39, 0x57, 0xf4, 0x21, 0xb1, 0xa3, 0xd4, 0xdb,
	0xc2, 0xd4, 0xed, 0x9e, 0x99, 0x26, 0xbd, 0x62, 0x9a, 0xf4, 0x0e, 0xd2, 0x05, 0xb3, 0xa3, 0x94,
	0x3e, 0x21, 0xce, 0x24, 0x33, 0x5f, 0x69, 0x33, 0xdc, 0x59, 0x93, 0x1d, 0xe7, 0x33, 0x8d, 0x81,
	0x8a, 0x7e, 0x42, 0x6c, 0xa5, 0xb1, 0x57, 0x36, 0xc3, 0xfb, 0x6b, 0xda, 0x53, 0x9c, 0x6f, 0xcc,
	0x56, 0xf0, 0xf5, 0xdb, 0x5a, 0xe5, 0xf7, 0xdb, 0x59, 0x13, 0xbe, 0x2c, 0x46, 0x1d, 0xb3, 0xb5,
	0xa2, 0x3d, 0xe2, 0x4c, 0xc6, 0x09, 0xde, 0x4e, 0x33, 0xdc, 0x5d, 0x3f, 0x01, 0xb6, 0xa3, 0x1f,
	0x01, 0x32, 0x03, 0x21, 0x7d, 0x4a, 0x9c, 0x8b, 0x44, 0xe3, 0x65, 0xc1, 0xa7, 0xb1, 0xaa, 0xc7,
	0xc6, 0x96, 0xcb, 0x2f, 0x12, 0x0d, 0xf2, 0x38, 0x9f, 0x4e, 0xb7, 0xc9, 0xb1, 0xd8, 0x73, 0x79,
	0x3c, 0xe8, 0xc3, 0x69, 0xb2, 0x41, 0x1f, 0x27, 0xd6, 0x6d, 0xa7, 0x39, 0xbb, 0xae, 0xcf, 0x06,
	0x7d, 0xb4, 0xdf, 0x0f, 0x71, 0x8c, 0x6d, 0xb0, 0xdf, 0x0f, 0x0b, 0xfb, 0xfd, 0x10, 0xed, 0xf7,
	0x43, 0x9c, 0x6d, 0x9b, 0xec, 0x97, 0xfa, 0x0c, 0xf5, 0x15, 0x1c, 0x36, 0x8d, 0x0d, 0x28, 0xe1,
	0x6b, 0x33, 0x72, 0xd4, 0x81, 0x3f, 0xf4, 0x0d, 0xb2, 0xc1, 0xdf, 0x34, 0xf0, 0xdc, 0x5f, 0x69,
	0x49, 0x3f, 0x27, 0x6e, 0x39, 0x1e, 0x6f, 0x7b, 0x01, 0x6c, 0xec, 0x26, 0xc1, 0x28, 0x9f, 0xf9,
	0xa4, 0x92, 0x46, 0x33, 0xbe, 0x52, 0xa2, 0xbf, 0x61, 0x2f, 0xc0, 0xc8, 0x4f, 0xee, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x84, 0x44, 0x55, 0xda, 0x53, 0x09, 0x00, 0x00,
}

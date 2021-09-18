// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator.proto

package validator

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,5,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of length strictly equal to this value.
	LengthEq *int64 `protobuf:"varint,16,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	// Requires that the value is in the enum.
	IsInEnum *bool `protobuf:"varint,17,opt,name=is_in_enum,json=isInEnum" json:"is_in_enum,omitempty"`
	// Ensures that a string value is in UUID format.
	// uuid_ver specifies the valid UUID versions. Valid values are: 0-5.
	// If uuid_ver is 0 all UUID versions are accepted.
	UuidVer              *int32   `protobuf:"varint,18,opt,name=uuid_ver,json=uuidVer" json:"uuid_ver,omitempty"`
	Alpha                *bool    `protobuf:"varint,19,opt,name=alpha" json:"alpha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValidator) Reset()         { *m = FieldValidator{} }
func (m *FieldValidator) String() string { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()    {}
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{0}
}
func (m *FieldValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValidator.Unmarshal(m, b)
}
func (m *FieldValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValidator.Marshal(b, m, deterministic)
}
func (m *FieldValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValidator.Merge(m, src)
}
func (m *FieldValidator) XXX_Size() int {
	return xxx_messageInfo_FieldValidator.Size(m)
}
func (m *FieldValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValidator.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValidator proto.InternalMessageInfo

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

func (m *FieldValidator) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidator) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil && m.StringNotEmpty != nil {
		return *m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() int64 {
	if m != nil && m.RepeatedCountMin != nil {
		return *m.RepeatedCountMin
	}
	return 0
}

func (m *FieldValidator) GetRepeatedCountMax() int64 {
	if m != nil && m.RepeatedCountMax != nil {
		return *m.RepeatedCountMax
	}
	return 0
}

func (m *FieldValidator) GetLengthGt() int64 {
	if m != nil && m.LengthGt != nil {
		return *m.LengthGt
	}
	return 0
}

func (m *FieldValidator) GetLengthLt() int64 {
	if m != nil && m.LengthLt != nil {
		return *m.LengthLt
	}
	return 0
}

func (m *FieldValidator) GetLengthEq() int64 {
	if m != nil && m.LengthEq != nil {
		return *m.LengthEq
	}
	return 0
}

func (m *FieldValidator) GetIsInEnum() bool {
	if m != nil && m.IsInEnum != nil {
		return *m.IsInEnum
	}
	return false
}

func (m *FieldValidator) GetUuidVer() int32 {
	if m != nil && m.UuidVer != nil {
		return *m.UuidVer
	}
	return 0
}

func (m *FieldValidator) GetAlpha() bool {
	if m != nil && m.Alpha != nil {
		return *m.Alpha
	}
	return false
}

type OneofValidator struct {
	// Require that one of the oneof fields is set.
	Required             *bool    `protobuf:"varint,1,opt,name=required" json:"required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneofValidator) Reset()         { *m = OneofValidator{} }
func (m *OneofValidator) String() string { return proto.CompactTextString(m) }
func (*OneofValidator) ProtoMessage()    {}
func (*OneofValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{1}
}
func (m *OneofValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneofValidator.Unmarshal(m, b)
}
func (m *OneofValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneofValidator.Marshal(b, m, deterministic)
}
func (m *OneofValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneofValidator.Merge(m, src)
}
func (m *OneofValidator) XXX_Size() int {
	return xxx_messageInfo_OneofValidator.Size(m)
}
func (m *OneofValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_OneofValidator.DiscardUnknown(m)
}

var xxx_messageInfo_OneofValidator proto.InternalMessageInfo

func (m *OneofValidator) GetRequired() bool {
	if m != nil && m.Required != nil {
		return *m.Required
	}
	return false
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "validator.proto",
}

var E_Oneof = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.OneofOptions)(nil),
	ExtensionType: (*OneofValidator)(nil),
	Field:         65021,
	Name:          "validator.oneof",
	Tag:           "bytes,65021,opt,name=oneof",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterType((*OneofValidator)(nil), "validator.OneofValidator")
	proto.RegisterExtension(E_Field)
	proto.RegisterExtension(E_Oneof)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptor_bf1c6ec7c0d80dd5) }

var fileDescriptor_bf1c6ec7c0d80dd5 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x15, 0xb6, 0xae, 0xa9, 0xbb, 0x75, 0xc5, 0x80, 0xe4, 0x0d, 0x26, 0xa2, 0xc2, 0x21,
	0x87, 0x29, 0x95, 0x38, 0xc2, 0x0d, 0x14, 0x2a, 0xa4, 0xc2, 0x50, 0x0e, 0x3b, 0x70, 0x89, 0xb2,
	0xe6, 0x35, 0xb5, 0xe4, 0xd8, 0xa9, 0xfd, 0x32, 0x75, 0x7f, 0x28, 0x7f, 0x0d, 0x20, 0x21, 0x3b,
	0x4b, 0x9a, 0x49, 0x3d, 0xbe, 0xef, 0xf7, 0xfc, 0xbd, 0xe7, 0xa7, 0x8f, 0x9c, 0xdf, 0x67, 0x82,
	0xe7, 0x19, 0x2a, 0x1d, 0x55, 0x5a, 0xa1, 0xa2, 0xa3, 0x4e, 0xb8, 0x0c, 0x0a, 0xa5, 0x0a, 0x01,
	0x73, 0x07, 0xee, 0xea, 0xf5, 0x3c, 0x07, 0xb3, 0xd2, 0xbc, 0xea, 0x9a, 0x67, 0xbf, 0x8f, 0xc9,
	0xe4, 0x2b, 0x07, 0x91, 0xdf, 0xb6, 0x8f, 0xe8, 0x4b, 0x32, 0xd0, 0x50, 0xc0, 0x8e, 0x79, 0x81,
	0x17, 0x8e, 0x92, 0xa6, 0xa0, 0xaf, 0xc8, 0x09, 0x97, 0x98, 0x16, 0xc8, 0x9e, 0x05, 0x5e, 0x78,
	0x94, 0x0c, 0xb8, 0xc4, 0x05, 0xb6, 0xb2, 0x40, 0x76, 0xd4, 0xc9, 0x4b, 0xa4, 0x57, 0x84, 0x94,
	0xa6, 0x48, 0x61, 0xc7, 0x0d, 0x1a, 0x76, 0x1c, 0x78, 0xa1, 0x9f, 0x8c, 0x4a, 0x53, 0xc4, 0x4e,
	0xa0, 0x6f, 0xc9, 0x78, 0x53, 0x97, 0x99, 0x4c, 0x41, 0x6b, 0xa5, 0xd9, 0xc0, 0x0d, 0x22, 0x4e,
	0x8a, 0xad, 0x42, 0x2f, 0x88, 0xbf, 0x16, 0x2a, 0x73, 0xf3, 0x4e, 0x02, 0x2f, 0xf4, 0x92, 0xa1,
	0xab, 0x17, 0xb8, 0x47, 0x02, 0xd9, 0xb0, 0x87, 0x96, 0x48, 0xdf, 0x91, 0xb3, 0x06, 0x41, 0x65,
	0xb8, 0x50, 0x92, 0xf9, 0x8e, 0x9f, 0x3a, 0x31, 0x6e, 0x34, 0xfa, 0x9a, 0x8c, 0x5a, 0x6b, 0x60,
	0x23, 0xd7, 0xe0, 0x3f, 0x7a, 0xc3, 0x1e, 0x0a, 0x04, 0x46, 0x7a, 0x70, 0x89, 0x40, 0x43, 0x32,
	0x35, 0xa8, 0xb9, 0x2c, 0x52, 0xa9, 0x30, 0x85, 0xb2, 0xc2, 0x07, 0x36, 0x76, 0x5f, 0x9b, 0x34,
	0xfa, 0x0f, 0x85, 0xb1, 0x55, 0xe9, 0x35, 0xa1, 0x1a, 0x2a, 0xc8, 0x10, 0xf2, 0x74, 0xa5, 0x6a,
	0x89, 0x69, 0xc9, 0x25, 0x3b, 0x75, 0x17, 0x9a, 0xb6, 0xe4, 0x8b, 0x05, 0xdf, 0xb9, 0x3c, 0xd4,
	0x9d, 0xed, 0xd8, 0xd9, 0xa1, 0xee, 0x6c, 0x67, 0x57, 0x14, 0x20, 0x0b, 0xdc, 0xd8, 0xdb, 0x4c,
	0x5c, 0x93, 0xdf, 0x08, 0x0b, 0xec, 0x41, 0x81, 0xec, 0xbc, 0x0f, 0x97, 0x7d, 0x08, 0x5b, 0x36,
	0xed, 0xc3, 0x78, 0x4b, 0xdf, 0x10, 0xc2, 0x4d, 0xca, 0x65, 0x0a, 0xb2, 0x2e, 0xd9, 0x73, 0xf7,
	0x2d, 0x9f, 0x9b, 0x6f, 0x32, 0x96, 0x75, 0x69, 0x8f, 0x5e, 0xd7, 0x3c, 0x4f, 0xef, 0x41, 0x33,
	0x1a, 0x78, 0xe1, 0x20, 0x19, 0xda, 0xfa, 0x16, 0x5c, 0x5c, 0x32, 0x51, 0x6d, 0x32, 0xf6, 0xc2,
	0xbd, 0x69, 0x8a, 0xd9, 0x35, 0x99, 0xdc, 0x48, 0x50, 0xeb, 0x7d, 0xac, 0x2e, 0x89, 0xaf, 0x61,
	0x5b, 0x73, 0x0d, 0xb9, 0x4b, 0x96, 0x9f, 0x74, 0xf5, 0xc7, 0x9f, 0x64, 0xb0, 0xb6, 0x21, 0xa4,
	0x57, 0x51, 0x93, 0xd8, 0xa8, 0x4d, 0x6c, 0xe4, 0xc2, 0x79, 0x53, 0x21, 0x57, 0xd2, 0xb0, 0xbf,
	0x7f, 0x6c, 0xca, 0xc6, 0x1f, 0x2e, 0xa2, 0x7d, 0xe8, 0x9f, 0xa6, 0x37, 0x69, 0x8c, 0xac, 0xa3,
	0xb2, 0xf3, 0x0f, 0x38, 0xba, 0xbd, 0x5a, 0xc7, 0x7f, 0x07, 0x1c, 0x9f, 0x2e, 0x9e, 0x34, 0x46,
	0x9f, 0xdf, 0xff, 0x9a, 0x15, 0x1c, 0x37, 0xf5, 0x5d, 0xb4, 0x52, 0xe5, 0xbc, 0x94, 0xa6, 0x50,
	0x0f, 0x99, 0x98, 0x23, 0x18, 0xfc, 0xd4, 0xbd, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xb5,
	0xd0, 0xe8, 0x8e, 0x03, 0x00, 0x00,
}

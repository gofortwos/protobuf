// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations/annotations.proto

package annotations

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AnnotationsTestEnum int32

const (
	AnnotationsTestEnum_ANNOTATIONS_TEST_ENUM_VALUE AnnotationsTestEnum = 0
)

func (e AnnotationsTestEnum) Type() protoreflect.EnumType {
	return xxx_Annotations_protoFile_EnumTypes[0]
}
func (e AnnotationsTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var AnnotationsTestEnum_name = map[int32]string{
	0: "ANNOTATIONS_TEST_ENUM_VALUE",
}

var AnnotationsTestEnum_value = map[string]int32{
	"ANNOTATIONS_TEST_ENUM_VALUE": 0,
}

func (x AnnotationsTestEnum) Enum() *AnnotationsTestEnum {
	p := new(AnnotationsTestEnum)
	*p = x
	return p
}

func (x AnnotationsTestEnum) String() string {
	return proto.EnumName(AnnotationsTestEnum_name, int32(x))
}

func (x *AnnotationsTestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AnnotationsTestEnum_value, data, "AnnotationsTestEnum")
	if err != nil {
		return err
	}
	*x = AnnotationsTestEnum(value)
	return nil
}

func (AnnotationsTestEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21dfaf6fd39fa3b7, []int{0}
}

type AnnotationsTestMessage struct {
	AnnotationsTestField *string  `protobuf:"bytes,1,opt,name=AnnotationsTestField" json:"AnnotationsTestField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_AnnotationsTestMessage struct{ m *AnnotationsTestMessage }

func (m *AnnotationsTestMessage) ProtoReflect() protoreflect.Message {
	return xxx_AnnotationsTestMessage{m}
}
func (m xxx_AnnotationsTestMessage) Type() protoreflect.MessageType {
	return xxx_Annotations_protoFile_MessageTypes[0].Type
}
func (m xxx_AnnotationsTestMessage) KnownFields() protoreflect.KnownFields {
	return xxx_Annotations_protoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_AnnotationsTestMessage) UnknownFields() protoreflect.UnknownFields {
	return xxx_Annotations_protoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_AnnotationsTestMessage) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *AnnotationsTestMessage) Reset()         { *m = AnnotationsTestMessage{} }
func (m *AnnotationsTestMessage) String() string { return proto.CompactTextString(m) }
func (*AnnotationsTestMessage) ProtoMessage()    {}
func (*AnnotationsTestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_21dfaf6fd39fa3b7, []int{0}
}

func (m *AnnotationsTestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationsTestMessage.Unmarshal(m, b)
}
func (m *AnnotationsTestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationsTestMessage.Marshal(b, m, deterministic)
}
func (m *AnnotationsTestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationsTestMessage.Merge(m, src)
}
func (m *AnnotationsTestMessage) XXX_Size() int {
	return xxx_messageInfo_AnnotationsTestMessage.Size(m)
}
func (m *AnnotationsTestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationsTestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationsTestMessage proto.InternalMessageInfo

func (m *AnnotationsTestMessage) GetAnnotationsTestField() string {
	if m != nil && m.AnnotationsTestField != nil {
		return *m.AnnotationsTestField
	}
	return ""
}

func init() {
	proto.RegisterFile("annotations/annotations.proto", fileDescriptor_21dfaf6fd39fa3b7)
	proto.RegisterEnum("goproto.protoc.annotations.AnnotationsTestEnum", AnnotationsTestEnum_name, AnnotationsTestEnum_value)
	proto.RegisterType((*AnnotationsTestMessage)(nil), "goproto.protoc.annotations.AnnotationsTestMessage")
}

var fileDescriptor_21dfaf6fd39fa3b7 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x47, 0x62, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0x49, 0xa5, 0xe7, 0x83, 0x19, 0x10, 0x6e, 0xb2, 0x1e, 0x92, 0x0a, 0x25, 0x1f, 0x2e, 0x31,
	0x47, 0x04, 0x37, 0x24, 0xb5, 0xb8, 0xc4, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0xc8, 0x88,
	0x4b, 0x04, 0x4d, 0xc6, 0x2d, 0x33, 0x35, 0x27, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xab, 0x9c, 0x96, 0x19, 0x97, 0x30, 0x9a, 0xb8, 0x6b, 0x5e, 0x69, 0xae, 0x90, 0x3c, 0x97, 0xb4,
	0xa3, 0x9f, 0x9f, 0x7f, 0x88, 0x63, 0x88, 0xa7, 0xbf, 0x5f, 0x70, 0x7c, 0x88, 0x6b, 0x70, 0x48,
	0xbc, 0xab, 0x5f, 0xa8, 0x6f, 0x7c, 0x98, 0xa3, 0x4f, 0xa8, 0xab, 0x00, 0x83, 0x93, 0x5b, 0x94,
	0x4b, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x4e, 0x62,
	0x5e, 0xba, 0x3e, 0xd8, 0xb5, 0x49, 0xa5, 0x69, 0xfa, 0x65, 0x46, 0xfa, 0xc9, 0xb9, 0x29, 0x10,
	0x7e, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0xbe, 0x7e, 0x49, 0x6a, 0x71, 0x49, 0x4a, 0x62,
	0x49, 0x22, 0xb2, 0x7f, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0xd6, 0xe5, 0x9d, 0x09, 0x01,
	0x00, 0x00,
}

func init() {
	xxx_Annotations_protoFile_FileDesc.Enums = xxx_Annotations_protoFile_EnumDescs[0:1]
	xxx_Annotations_protoFile_FileDesc.Messages = xxx_Annotations_protoFile_MessageDescs[0:1]
	var err error
	Annotations_protoFile, err = prototype.NewFile(&xxx_Annotations_protoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Annotations_protoFile protoreflect.FileDescriptor

var xxx_Annotations_protoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto2,
	Path:    "annotations/annotations.proto",
	Package: "goproto.protoc.annotations",
}
var xxx_Annotations_protoFile_EnumTypes = [1]protoreflect.EnumType{
	prototype.GoEnum(
		xxx_Annotations_protoFile_EnumDescs[0].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.Enum {
			return AnnotationsTestEnum(n)
		},
	),
}
var xxx_Annotations_protoFile_EnumDescs = [1]prototype.Enum{
	{
		Name: "AnnotationsTestEnum",
		Values: []prototype.EnumValue{
			{Name: "ANNOTATIONS_TEST_ENUM_VALUE", Number: 0},
		},
	},
}
var xxx_Annotations_protoFile_MessageTypes = [1]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_Annotations_protoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.Message {
			return xxx_AnnotationsTestMessage{new(AnnotationsTestMessage)}
		},
	)},
}
var xxx_Annotations_protoFile_MessageDescs = [1]prototype.Message{
	{
		Name: "AnnotationsTestMessage",
		Fields: []prototype.Field{
			{
				Name:        "AnnotationsTestField",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "AnnotationsTestField",
				IsPacked:    prototype.False,
			},
		},
	},
}

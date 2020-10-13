// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: semantic.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SemanticDescriptor_Labels int32

const (
	SemanticDescriptor_DEFAULT  SemanticDescriptor_Labels = 0
	SemanticDescriptor_ADVANCED SemanticDescriptor_Labels = 1 // Field should be offered for modification by advanced users
	// only.
	SemanticDescriptor_HIDDEN SemanticDescriptor_Labels = 2 // Field should be hidden from UIs - typically for internally
)

// Enum value maps for SemanticDescriptor_Labels.
var (
	SemanticDescriptor_Labels_name = map[int32]string{
		0: "DEFAULT",
		1: "ADVANCED",
		2: "HIDDEN",
	}
	SemanticDescriptor_Labels_value = map[string]int32{
		"DEFAULT":  0,
		"ADVANCED": 1,
		"HIDDEN":   2,
	}
)

func (x SemanticDescriptor_Labels) Enum() *SemanticDescriptor_Labels {
	p := new(SemanticDescriptor_Labels)
	*p = x
	return p
}

func (x SemanticDescriptor_Labels) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SemanticDescriptor_Labels) Descriptor() protoreflect.EnumDescriptor {
	return file_semantic_proto_enumTypes[0].Descriptor()
}

func (SemanticDescriptor_Labels) Type() protoreflect.EnumType {
	return &file_semantic_proto_enumTypes[0]
}

func (x SemanticDescriptor_Labels) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SemanticDescriptor_Labels.Descriptor instead.
func (SemanticDescriptor_Labels) EnumDescriptor() ([]byte, []int) {
	return file_semantic_proto_rawDescGZIP(), []int{0, 0}
}

type SemanticDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The semantic name of the SemanticValue contained in this field.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The type of this field can be specified dynamically. This is the name of
	// the attribute to use to retrieve the SemanticValue class to be used for
	// parsing this field.
	DynamicType string `protobuf:"bytes,5,opt,name=dynamic_type,json=dynamicType,proto3" json:"dynamic_type,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// A JSON encoded default. This should be the same type as the
	// field it describes.
	Default string                      `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	Label   []SemanticDescriptor_Labels `protobuf:"varint,3,rep,packed,name=label,proto3,enum=proto.SemanticDescriptor_Labels" json:"label,omitempty"`
	// A friendly name for this field - to be used in GUIs etc.
	FriendlyName string `protobuf:"bytes,4,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
}

func (x *SemanticDescriptor) Reset() {
	*x = SemanticDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_semantic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemanticDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemanticDescriptor) ProtoMessage() {}

func (x *SemanticDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_semantic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemanticDescriptor.ProtoReflect.Descriptor instead.
func (*SemanticDescriptor) Descriptor() ([]byte, []int) {
	return file_semantic_proto_rawDescGZIP(), []int{0}
}

func (x *SemanticDescriptor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SemanticDescriptor) GetDynamicType() string {
	if x != nil {
		return x.DynamicType
	}
	return ""
}

func (x *SemanticDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SemanticDescriptor) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *SemanticDescriptor) GetLabel() []SemanticDescriptor_Labels {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *SemanticDescriptor) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

type SemanticMessageDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describe the purpose of this protobuf.
	Description  string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	FriendlyName string `protobuf:"bytes,3,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
}

func (x *SemanticMessageDescriptor) Reset() {
	*x = SemanticMessageDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_semantic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemanticMessageDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemanticMessageDescriptor) ProtoMessage() {}

func (x *SemanticMessageDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_semantic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemanticMessageDescriptor.ProtoReflect.Descriptor instead.
func (*SemanticMessageDescriptor) Descriptor() ([]byte, []int) {
	return file_semantic_proto_rawDescGZIP(), []int{1}
}

func (x *SemanticMessageDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SemanticMessageDescriptor) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

var file_semantic_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*SemanticDescriptor)(nil),
		Field:         51584972,
		Name:          "proto.sem_type",
		Tag:           "bytes,51584972,opt,name=sem_type",
		Filename:      "semantic.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*SemanticMessageDescriptor)(nil),
		Field:         51584971,
		Name:          "proto.semantic",
		Tag:           "bytes,51584971,opt,name=semantic",
		Filename:      "semantic.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         48651165,
		Name:          "proto.description",
		Tag:           "bytes,48651165,opt,name=description",
		Filename:      "semantic.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*SemanticDescriptor_Labels)(nil),
		Field:         48651166,
		Name:          "proto.label",
		Tag:           "varint,48651166,opt,name=label,enum=proto.SemanticDescriptor_Labels",
		Filename:      "semantic.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional proto.SemanticDescriptor sem_type = 51584972;
	E_SemType = &file_semantic_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional proto.SemanticMessageDescriptor semantic = 51584971;
	E_Semantic = &file_semantic_proto_extTypes[1]
)

// Extension fields to descriptor.EnumValueOptions.
var (
	// optional string description = 48651165;
	E_Description = &file_semantic_proto_extTypes[2]
	// optional proto.SemanticDescriptor.Labels label = 48651166;
	E_Label = &file_semantic_proto_extTypes[3]
)

var File_semantic_proto protoreflect.FileDescriptor

var file_semantic_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x12, 0x53, 0x65,
	0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e,
	0x74, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x2f, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x44, 0x56, 0x41, 0x4e,
	0x43, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10,
	0x02, 0x22, 0x62, 0x0a, 0x19, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x56, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xcc, 0xbf, 0xcc, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x73, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x60, 0x0a,
	0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcb, 0xbf, 0xcc, 0x18, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x3a,
	0x46, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x9d, 0xb7, 0x99, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x5c, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x9e, 0xb7, 0x99, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x2d, 0x5a, 0x2b, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c,
	0x6f, 0x63, 0x69, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_semantic_proto_rawDescOnce sync.Once
	file_semantic_proto_rawDescData = file_semantic_proto_rawDesc
)

func file_semantic_proto_rawDescGZIP() []byte {
	file_semantic_proto_rawDescOnce.Do(func() {
		file_semantic_proto_rawDescData = protoimpl.X.CompressGZIP(file_semantic_proto_rawDescData)
	})
	return file_semantic_proto_rawDescData
}

var file_semantic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_semantic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_semantic_proto_goTypes = []interface{}{
	(SemanticDescriptor_Labels)(0),      // 0: proto.SemanticDescriptor.Labels
	(*SemanticDescriptor)(nil),          // 1: proto.SemanticDescriptor
	(*SemanticMessageDescriptor)(nil),   // 2: proto.SemanticMessageDescriptor
	(*descriptor.FieldOptions)(nil),     // 3: google.protobuf.FieldOptions
	(*descriptor.MessageOptions)(nil),   // 4: google.protobuf.MessageOptions
	(*descriptor.EnumValueOptions)(nil), // 5: google.protobuf.EnumValueOptions
}
var file_semantic_proto_depIdxs = []int32{
	0, // 0: proto.SemanticDescriptor.label:type_name -> proto.SemanticDescriptor.Labels
	3, // 1: proto.sem_type:extendee -> google.protobuf.FieldOptions
	4, // 2: proto.semantic:extendee -> google.protobuf.MessageOptions
	5, // 3: proto.description:extendee -> google.protobuf.EnumValueOptions
	5, // 4: proto.label:extendee -> google.protobuf.EnumValueOptions
	1, // 5: proto.sem_type:type_name -> proto.SemanticDescriptor
	2, // 6: proto.semantic:type_name -> proto.SemanticMessageDescriptor
	0, // 7: proto.label:type_name -> proto.SemanticDescriptor.Labels
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	5, // [5:8] is the sub-list for extension type_name
	1, // [1:5] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_semantic_proto_init() }
func file_semantic_proto_init() {
	if File_semantic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_semantic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemanticDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_semantic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemanticMessageDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_semantic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_semantic_proto_goTypes,
		DependencyIndexes: file_semantic_proto_depIdxs,
		EnumInfos:         file_semantic_proto_enumTypes,
		MessageInfos:      file_semantic_proto_msgTypes,
		ExtensionInfos:    file_semantic_proto_extTypes,
	}.Build()
	File_semantic_proto = out.File
	file_semantic_proto_rawDesc = nil
	file_semantic_proto_goTypes = nil
	file_semantic_proto_depIdxs = nil
}

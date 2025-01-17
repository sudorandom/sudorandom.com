// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: types.proto

package gen

import (
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

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Basic types
	IntValue    int32  `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	UintValue   uint64 `protobuf:"varint,2,opt,name=uint_value,json=uintValue,proto3" json:"uint_value,omitempty"`
	BoolValue   bool   `protobuf:"varint,3,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	// Repeated fields (varint encoding)
	RepeatedIntValue []int32                    `protobuf:"varint,10,rep,packed,name=repeated_int_value,json=repeatedIntValue,proto3" json:"repeated_int_value,omitempty"`
	NestedMessage    *TestMessage_NestedMessage `protobuf:"bytes,11,opt,name=nested_message,json=nestedMessage,proto3" json:"nested_message,omitempty"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetIntValue() int32 {
	if x != nil {
		return x.IntValue
	}
	return 0
}

func (x *TestMessage) GetUintValue() uint64 {
	if x != nil {
		return x.UintValue
	}
	return 0
}

func (x *TestMessage) GetBoolValue() bool {
	if x != nil {
		return x.BoolValue
	}
	return false
}

func (x *TestMessage) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

func (x *TestMessage) GetRepeatedIntValue() []int32 {
	if x != nil {
		return x.RepeatedIntValue
	}
	return nil
}

func (x *TestMessage) GetNestedMessage() *TestMessage_NestedMessage {
	if x != nil {
		return x.NestedMessage
	}
	return nil
}

// Nested message
type TestMessage_NestedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedString string `protobuf:"bytes,1,opt,name=nested_string,json=nestedString,proto3" json:"nested_string,omitempty"`
}

func (x *TestMessage_NestedMessage) Reset() {
	*x = TestMessage_NestedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage_NestedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage_NestedMessage) ProtoMessage() {}

func (x *TestMessage_NestedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage_NestedMessage.ProtoReflect.Descriptor instead.
func (*TestMessage_NestedMessage) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TestMessage_NestedMessage) GetNestedString() string {
	if x != nil {
		return x.NestedString
	}
	return ""
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xba, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x34, 0x0a,
	0x0d, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x42, 0x91, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75,
	0x64, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x64, 0x6f, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x66, 0x72, 0x6f,
	0x6d, 0x2d, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03,
	0x45, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xca, 0x02, 0x07,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xe2, 0x02, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_proto_goTypes = []interface{}{
	(*TestMessage)(nil),               // 0: example.TestMessage
	(*TestMessage_NestedMessage)(nil), // 1: example.TestMessage.NestedMessage
}
var file_types_proto_depIdxs = []int32{
	1, // 0: example.TestMessage.nested_message:type_name -> example.TestMessage.NestedMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage_NestedMessage); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}

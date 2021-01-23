// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: ElectionMessage.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type ElectionMessage_Type int32

const (
	ElectionMessage_EXPLORER ElectionMessage_Type = 0
	ElectionMessage_ECHO     ElectionMessage_Type = 1
)

// Enum value maps for ElectionMessage_Type.
var (
	ElectionMessage_Type_name = map[int32]string{
		0: "EXPLORER",
		1: "ECHO",
	}
	ElectionMessage_Type_value = map[string]int32{
		"EXPLORER": 0,
		"ECHO":     1,
	}
)

func (x ElectionMessage_Type) Enum() *ElectionMessage_Type {
	p := new(ElectionMessage_Type)
	*p = x
	return p
}

func (x ElectionMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ElectionMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_ElectionMessage_proto_enumTypes[0].Descriptor()
}

func (ElectionMessage_Type) Type() protoreflect.EnumType {
	return &file_ElectionMessage_proto_enumTypes[0]
}

func (x ElectionMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ElectionMessage_Type.Descriptor instead.
func (ElectionMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_ElectionMessage_proto_rawDescGZIP(), []int{0, 0}
}

type ElectionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      ElectionMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ueb03.ElectionMessage_Type" json:"type,omitempty"`
	Initiator uint64               `protobuf:"varint,2,opt,name=initiator,proto3" json:"initiator,omitempty"`
}

func (x *ElectionMessage) Reset() {
	*x = ElectionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ElectionMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionMessage) ProtoMessage() {}

func (x *ElectionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ElectionMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionMessage.ProtoReflect.Descriptor instead.
func (*ElectionMessage) Descriptor() ([]byte, []int) {
	return file_ElectionMessage_proto_rawDescGZIP(), []int{0}
}

func (x *ElectionMessage) GetType() ElectionMessage_Type {
	if x != nil {
		return x.Type
	}
	return ElectionMessage_EXPLORER
}

func (x *ElectionMessage) GetInitiator() uint64 {
	if x != nil {
		return x.Initiator
	}
	return 0
}

var File_ElectionMessage_proto protoreflect.FileDescriptor

var file_ElectionMessage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x65, 0x62, 0x30, 0x33, 0x22, 0x80,
	0x01, 0x0a, 0x0f, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x33, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x50,
	0x4c, 0x4f, 0x52, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x43, 0x48, 0x4f, 0x10,
	0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6b, 0x65, 0x6c, 0x73, 0x63, 0x68, 0x2f, 0x76, 0x61, 0x61, 0x2f, 0x75, 0x65, 0x62, 0x30,
	0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ElectionMessage_proto_rawDescOnce sync.Once
	file_ElectionMessage_proto_rawDescData = file_ElectionMessage_proto_rawDesc
)

func file_ElectionMessage_proto_rawDescGZIP() []byte {
	file_ElectionMessage_proto_rawDescOnce.Do(func() {
		file_ElectionMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_ElectionMessage_proto_rawDescData)
	})
	return file_ElectionMessage_proto_rawDescData
}

var file_ElectionMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ElectionMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ElectionMessage_proto_goTypes = []interface{}{
	(ElectionMessage_Type)(0), // 0: ueb03.ElectionMessage.Type
	(*ElectionMessage)(nil),   // 1: ueb03.ElectionMessage
}
var file_ElectionMessage_proto_depIdxs = []int32{
	0, // 0: ueb03.ElectionMessage.type:type_name -> ueb03.ElectionMessage.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ElectionMessage_proto_init() }
func file_ElectionMessage_proto_init() {
	if File_ElectionMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ElectionMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionMessage); i {
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
			RawDescriptor: file_ElectionMessage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ElectionMessage_proto_goTypes,
		DependencyIndexes: file_ElectionMessage_proto_depIdxs,
		EnumInfos:         file_ElectionMessage_proto_enumTypes,
		MessageInfos:      file_ElectionMessage_proto_msgTypes,
	}.Build()
	File_ElectionMessage_proto = out.File
	file_ElectionMessage_proto_rawDesc = nil
	file_ElectionMessage_proto_goTypes = nil
	file_ElectionMessage_proto_depIdxs = nil
}
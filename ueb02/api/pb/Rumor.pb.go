// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Rumor.proto

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

type Rumor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Rumor) Reset() {
	*x = Rumor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Rumor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rumor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rumor) ProtoMessage() {}

func (x *Rumor) ProtoReflect() protoreflect.Message {
	mi := &file_Rumor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rumor.ProtoReflect.Descriptor instead.
func (*Rumor) Descriptor() ([]byte, []int) {
	return file_Rumor_proto_rawDescGZIP(), []int{0}
}

func (x *Rumor) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_Rumor_proto protoreflect.FileDescriptor

var file_Rumor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x52, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75,
	0x65, 0x62, 0x30, 0x32, 0x22, 0x21, 0x0a, 0x05, 0x52, 0x75, 0x6d, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x65, 0x6c, 0x73, 0x63, 0x68, 0x2f, 0x76, 0x61,
	0x61, 0x2f, 0x75, 0x65, 0x62, 0x30, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Rumor_proto_rawDescOnce sync.Once
	file_Rumor_proto_rawDescData = file_Rumor_proto_rawDesc
)

func file_Rumor_proto_rawDescGZIP() []byte {
	file_Rumor_proto_rawDescOnce.Do(func() {
		file_Rumor_proto_rawDescData = protoimpl.X.CompressGZIP(file_Rumor_proto_rawDescData)
	})
	return file_Rumor_proto_rawDescData
}

var file_Rumor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Rumor_proto_goTypes = []interface{}{
	(*Rumor)(nil), // 0: ueb02.Rumor
}
var file_Rumor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Rumor_proto_init() }
func file_Rumor_proto_init() {
	if File_Rumor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Rumor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rumor); i {
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
			RawDescriptor: file_Rumor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Rumor_proto_goTypes,
		DependencyIndexes: file_Rumor_proto_depIdxs,
		MessageInfos:      file_Rumor_proto_msgTypes,
	}.Build()
	File_Rumor_proto = out.File
	file_Rumor_proto_rawDesc = nil
	file_Rumor_proto_goTypes = nil
	file_Rumor_proto_depIdxs = nil
}

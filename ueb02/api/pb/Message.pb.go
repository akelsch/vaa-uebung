// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Message.proto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// Types that are assignable to Msg:
	//	*Message_ControlMessage
	//	*Message_ApplicationMessage
	//	*Message_Rumor
	//	*Message_Election
	//	*Message_Status
	Msg isMessage_Msg `protobuf_oneof:"msg"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (m *Message) GetMsg() isMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *Message) GetControlMessage() *ControlMessage {
	if x, ok := x.GetMsg().(*Message_ControlMessage); ok {
		return x.ControlMessage
	}
	return nil
}

func (x *Message) GetApplicationMessage() *ApplicationMessage {
	if x, ok := x.GetMsg().(*Message_ApplicationMessage); ok {
		return x.ApplicationMessage
	}
	return nil
}

func (x *Message) GetRumor() *Rumor {
	if x, ok := x.GetMsg().(*Message_Rumor); ok {
		return x.Rumor
	}
	return nil
}

func (x *Message) GetElection() *Election {
	if x, ok := x.GetMsg().(*Message_Election); ok {
		return x.Election
	}
	return nil
}

func (x *Message) GetStatus() *Status {
	if x, ok := x.GetMsg().(*Message_Status); ok {
		return x.Status
	}
	return nil
}

type isMessage_Msg interface {
	isMessage_Msg()
}

type Message_ControlMessage struct {
	ControlMessage *ControlMessage `protobuf:"bytes,2,opt,name=control_message,json=controlMessage,proto3,oneof"`
}

type Message_ApplicationMessage struct {
	ApplicationMessage *ApplicationMessage `protobuf:"bytes,3,opt,name=application_message,json=applicationMessage,proto3,oneof"`
}

type Message_Rumor struct {
	Rumor *Rumor `protobuf:"bytes,4,opt,name=rumor,proto3,oneof"`
}

type Message_Election struct {
	Election *Election `protobuf:"bytes,5,opt,name=election,proto3,oneof"`
}

type Message_Status struct {
	Status *Status `protobuf:"bytes,6,opt,name=status,proto3,oneof"`
}

func (*Message_ControlMessage) isMessage_Msg() {}

func (*Message_ApplicationMessage) isMessage_Msg() {}

func (*Message_Rumor) isMessage_Msg() {}

func (*Message_Election) isMessage_Msg() {}

func (*Message_Status) isMessage_Msg() {}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x75, 0x65, 0x62, 0x30, 0x32, 0x1a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x52, 0x75, 0x6d, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x75, 0x65, 0x62, 0x30, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x75, 0x6d, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x32, 0x2e, 0x52, 0x75, 0x6d,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x72, 0x75, 0x6d, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x08, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x75, 0x65, 0x62, 0x30, 0x32, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x65, 0x62,
	0x30, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x65, 0x6c, 0x73, 0x63, 0x68,
	0x2f, 0x76, 0x61, 0x61, 0x2f, 0x75, 0x65, 0x62, 0x30, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Message_proto_rawDescOnce sync.Once
	file_Message_proto_rawDescData = file_Message_proto_rawDesc
)

func file_Message_proto_rawDescGZIP() []byte {
	file_Message_proto_rawDescOnce.Do(func() {
		file_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Message_proto_rawDescData)
	})
	return file_Message_proto_rawDescData
}

var file_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Message_proto_goTypes = []interface{}{
	(*Message)(nil),            // 0: ueb02.Message
	(*ControlMessage)(nil),     // 1: ueb02.ControlMessage
	(*ApplicationMessage)(nil), // 2: ueb02.ApplicationMessage
	(*Rumor)(nil),              // 3: ueb02.Rumor
	(*Election)(nil),           // 4: ueb02.Election
	(*Status)(nil),             // 5: ueb02.Status
}
var file_Message_proto_depIdxs = []int32{
	1, // 0: ueb02.Message.control_message:type_name -> ueb02.ControlMessage
	2, // 1: ueb02.Message.application_message:type_name -> ueb02.ApplicationMessage
	3, // 2: ueb02.Message.rumor:type_name -> ueb02.Rumor
	4, // 3: ueb02.Message.election:type_name -> ueb02.Election
	5, // 4: ueb02.Message.status:type_name -> ueb02.Status
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Message_proto_init() }
func file_Message_proto_init() {
	if File_Message_proto != nil {
		return
	}
	file_ControlMessage_proto_init()
	file_ApplicationMessage_proto_init()
	file_Rumor_proto_init()
	file_Election_proto_init()
	file_Status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_Message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_ControlMessage)(nil),
		(*Message_ApplicationMessage)(nil),
		(*Message_Rumor)(nil),
		(*Message_Election)(nil),
		(*Message_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Message_proto_goTypes,
		DependencyIndexes: file_Message_proto_depIdxs,
		MessageInfos:      file_Message_proto_msgTypes,
	}.Build()
	File_Message_proto = out.File
	file_Message_proto_rawDesc = nil
	file_Message_proto_goTypes = nil
	file_Message_proto_depIdxs = nil
}

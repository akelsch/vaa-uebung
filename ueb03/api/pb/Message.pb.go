// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: Message.proto

package pb

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Sender     uint64 `protobuf:"varint,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver   uint64 `protobuf:"varint,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Types that are assignable to Msg:
	//	*Message_ControlMessage
	//	*Message_ApplicationMessage
	//	*Message_MutexMessage
	//	*Message_ElectionMessage
	//	*Message_SnapshotMessage
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

func (x *Message) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Message) GetSender() uint64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *Message) GetReceiver() uint64 {
	if x != nil {
		return x.Receiver
	}
	return 0
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

func (x *Message) GetMutexMessage() *MutexMessage {
	if x, ok := x.GetMsg().(*Message_MutexMessage); ok {
		return x.MutexMessage
	}
	return nil
}

func (x *Message) GetElectionMessage() *ElectionMessage {
	if x, ok := x.GetMsg().(*Message_ElectionMessage); ok {
		return x.ElectionMessage
	}
	return nil
}

func (x *Message) GetSnapshotMessage() *SnapshotMessage {
	if x, ok := x.GetMsg().(*Message_SnapshotMessage); ok {
		return x.SnapshotMessage
	}
	return nil
}

type isMessage_Msg interface {
	isMessage_Msg()
}

type Message_ControlMessage struct {
	ControlMessage *ControlMessage `protobuf:"bytes,4,opt,name=control_message,json=controlMessage,proto3,oneof"`
}

type Message_ApplicationMessage struct {
	ApplicationMessage *ApplicationMessage `protobuf:"bytes,5,opt,name=application_message,json=applicationMessage,proto3,oneof"`
}

type Message_MutexMessage struct {
	MutexMessage *MutexMessage `protobuf:"bytes,6,opt,name=mutex_message,json=mutexMessage,proto3,oneof"`
}

type Message_ElectionMessage struct {
	ElectionMessage *ElectionMessage `protobuf:"bytes,7,opt,name=election_message,json=electionMessage,proto3,oneof"`
}

type Message_SnapshotMessage struct {
	SnapshotMessage *SnapshotMessage `protobuf:"bytes,8,opt,name=snapshot_message,json=snapshotMessage,proto3,oneof"`
}

func (*Message_ControlMessage) isMessage_Msg() {}

func (*Message_ApplicationMessage) isMessage_Msg() {}

func (*Message_MutexMessage) isMessage_Msg() {}

func (*Message_ElectionMessage) isMessage_Msg() {}

func (*Message_SnapshotMessage) isMessage_Msg() {}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x75, 0x65, 0x62, 0x30, 0x33, 0x1a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x13, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x33, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x6d, 0x75, 0x74, 0x65,
	0x78, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x33, 0x2e, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x75, 0x65, 0x62, 0x30, 0x33, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x65, 0x62, 0x30, 0x33, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x05,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x65, 0x6c, 0x73, 0x63, 0x68, 0x2f, 0x76, 0x61, 0x61, 0x2f,
	0x75, 0x65, 0x62, 0x30, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	(*Message)(nil),            // 0: ueb03.Message
	(*ControlMessage)(nil),     // 1: ueb03.ControlMessage
	(*ApplicationMessage)(nil), // 2: ueb03.ApplicationMessage
	(*MutexMessage)(nil),       // 3: ueb03.MutexMessage
	(*ElectionMessage)(nil),    // 4: ueb03.ElectionMessage
	(*SnapshotMessage)(nil),    // 5: ueb03.SnapshotMessage
}
var file_Message_proto_depIdxs = []int32{
	1, // 0: ueb03.Message.control_message:type_name -> ueb03.ControlMessage
	2, // 1: ueb03.Message.application_message:type_name -> ueb03.ApplicationMessage
	3, // 2: ueb03.Message.mutex_message:type_name -> ueb03.MutexMessage
	4, // 3: ueb03.Message.election_message:type_name -> ueb03.ElectionMessage
	5, // 4: ueb03.Message.snapshot_message:type_name -> ueb03.SnapshotMessage
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
	file_MutexMessage_proto_init()
	file_ElectionMessage_proto_init()
	file_SnapshotMessage_proto_init()
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
		(*Message_MutexMessage)(nil),
		(*Message_ElectionMessage)(nil),
		(*Message_SnapshotMessage)(nil),
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

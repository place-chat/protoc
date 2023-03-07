// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: proto/Message.proto

package message_service

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

	Id     int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Coords []float32 `protobuf:"fixed32,2,rep,packed,name=coords,proto3" json:"coords,omitempty"`
	Text   string    `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Message_proto_msgTypes[0]
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
	return file_proto_Message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetCoords() []float32 {
	if x != nil {
		return x.Coords
	}
	return nil
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MessageCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coords []float32 `protobuf:"fixed32,1,rep,packed,name=coords,proto3" json:"coords,omitempty"`
	Text   string    `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MessageCreateRequest) Reset() {
	*x = MessageCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCreateRequest) ProtoMessage() {}

func (x *MessageCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCreateRequest.ProtoReflect.Descriptor instead.
func (*MessageCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_Message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageCreateRequest) GetCoords() []float32 {
	if x != nil {
		return x.Coords
	}
	return nil
}

func (x *MessageCreateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MessageCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageCreateReply) Reset() {
	*x = MessageCreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCreateReply) ProtoMessage() {}

func (x *MessageCreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCreateReply.ProtoReflect.Descriptor instead.
func (*MessageCreateReply) Descriptor() ([]byte, []int) {
	return file_proto_Message_proto_rawDescGZIP(), []int{2}
}

func (x *MessageCreateReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_Message_proto protoreflect.FileDescriptor

var file_proto_Message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x42, 0x0a,
	0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x22, 0x24, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x6d, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_Message_proto_rawDescOnce sync.Once
	file_proto_Message_proto_rawDescData = file_proto_Message_proto_rawDesc
)

func file_proto_Message_proto_rawDescGZIP() []byte {
	file_proto_Message_proto_rawDescOnce.Do(func() {
		file_proto_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Message_proto_rawDescData)
	})
	return file_proto_Message_proto_rawDescData
}

var file_proto_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_Message_proto_goTypes = []interface{}{
	(*Message)(nil),              // 0: message_service.Message
	(*MessageCreateRequest)(nil), // 1: message_service.MessageCreateRequest
	(*MessageCreateReply)(nil),   // 2: message_service.MessageCreateReply
}
var file_proto_Message_proto_depIdxs = []int32{
	1, // 0: message_service.MessageService.MessageCreate:input_type -> message_service.MessageCreateRequest
	2, // 1: message_service.MessageService.MessageCreate:output_type -> message_service.MessageCreateReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_Message_proto_init() }
func file_proto_Message_proto_init() {
	if File_proto_Message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_Message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCreateRequest); i {
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
		file_proto_Message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCreateReply); i {
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
			RawDescriptor: file_proto_Message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Message_proto_goTypes,
		DependencyIndexes: file_proto_Message_proto_depIdxs,
		MessageInfos:      file_proto_Message_proto_msgTypes,
	}.Build()
	File_proto_Message_proto = out.File
	file_proto_Message_proto_rawDesc = nil
	file_proto_Message_proto_goTypes = nil
	file_proto_Message_proto_depIdxs = nil
}

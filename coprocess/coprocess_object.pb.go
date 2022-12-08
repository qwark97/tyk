// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: coprocess_object.proto

package coprocess

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HookType HookType           `protobuf:"varint,1,opt,name=hook_type,json=hookType,proto3,enum=coprocess.HookType" json:"hook_type,omitempty"`
	HookName string             `protobuf:"bytes,2,opt,name=hook_name,json=hookName,proto3" json:"hook_name,omitempty"`
	Request  *MiniRequestObject `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Session  *SessionState      `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	Metadata map[string]string  `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec     map[string]string  `protobuf:"bytes,6,rep,name=spec,proto3" json:"spec,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Response *ResponseObject    `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_coprocess_object_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetHookType() HookType {
	if x != nil {
		return x.HookType
	}
	return HookType_Unknown
}

func (x *Object) GetHookName() string {
	if x != nil {
		return x.HookName
	}
	return ""
}

func (x *Object) GetRequest() *MiniRequestObject {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Object) GetSession() *SessionState {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *Object) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Object) GetSpec() map[string]string {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Object) GetResponse() *ResponseObject {
	if x != nil {
		return x.Response
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_coprocess_object_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type EventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventReply) Reset() {
	*x = EventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocess_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReply) ProtoMessage() {}

func (x *EventReply) ProtoReflect() protoreflect.Message {
	mi := &file_coprocess_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventReply.ProtoReflect.Descriptor instead.
func (*EventReply) Descriptor() ([]byte, []int) {
	return file_coprocess_object_proto_rawDescGZIP(), []int{2}
}

var File_coprocess_object_proto protoreflect.FileDescriptor

var file_coprocess_object_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x1a, 0x23, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6d,
	0x69, 0x6e, 0x69, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdd, 0x03, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x68,
	0x6f, 0x6f, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x68, 0x6f, 0x6f, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x53, 0x70, 0x65, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x21, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0x7c, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x11,
	0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x79,
	0x6b, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x6b, 0x2f, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_coprocess_object_proto_rawDescOnce sync.Once
	file_coprocess_object_proto_rawDescData = file_coprocess_object_proto_rawDesc
)

func file_coprocess_object_proto_rawDescGZIP() []byte {
	file_coprocess_object_proto_rawDescOnce.Do(func() {
		file_coprocess_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_coprocess_object_proto_rawDescData)
	})
	return file_coprocess_object_proto_rawDescData
}

var file_coprocess_object_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_coprocess_object_proto_goTypes = []interface{}{
	(*Object)(nil),            // 0: coprocess.Object
	(*Event)(nil),             // 1: coprocess.Event
	(*EventReply)(nil),        // 2: coprocess.EventReply
	nil,                       // 3: coprocess.Object.MetadataEntry
	nil,                       // 4: coprocess.Object.SpecEntry
	(HookType)(0),             // 5: coprocess.HookType
	(*MiniRequestObject)(nil), // 6: coprocess.MiniRequestObject
	(*SessionState)(nil),      // 7: coprocess.SessionState
	(*ResponseObject)(nil),    // 8: coprocess.ResponseObject
}
var file_coprocess_object_proto_depIdxs = []int32{
	5, // 0: coprocess.Object.hook_type:type_name -> coprocess.HookType
	6, // 1: coprocess.Object.request:type_name -> coprocess.MiniRequestObject
	7, // 2: coprocess.Object.session:type_name -> coprocess.SessionState
	3, // 3: coprocess.Object.metadata:type_name -> coprocess.Object.MetadataEntry
	4, // 4: coprocess.Object.spec:type_name -> coprocess.Object.SpecEntry
	8, // 5: coprocess.Object.response:type_name -> coprocess.ResponseObject
	0, // 6: coprocess.Dispatcher.Dispatch:input_type -> coprocess.Object
	1, // 7: coprocess.Dispatcher.DispatchEvent:input_type -> coprocess.Event
	0, // 8: coprocess.Dispatcher.Dispatch:output_type -> coprocess.Object
	2, // 9: coprocess.Dispatcher.DispatchEvent:output_type -> coprocess.EventReply
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_coprocess_object_proto_init() }
func file_coprocess_object_proto_init() {
	if File_coprocess_object_proto != nil {
		return
	}
	file_coprocess_mini_request_object_proto_init()
	file_coprocess_response_object_proto_init()
	file_coprocess_session_state_proto_init()
	file_coprocess_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coprocess_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_coprocess_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_coprocess_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReply); i {
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
			RawDescriptor: file_coprocess_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coprocess_object_proto_goTypes,
		DependencyIndexes: file_coprocess_object_proto_depIdxs,
		MessageInfos:      file_coprocess_object_proto_msgTypes,
	}.Build()
	File_coprocess_object_proto = out.File
	file_coprocess_object_proto_rawDesc = nil
	file_coprocess_object_proto_goTypes = nil
	file_coprocess_object_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherClient interface {
	Dispatch(ctx context.Context, in *Object, opts ...grpc.CallOption) (*Object, error)
	DispatchEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventReply, error)
}

type dispatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherClient(cc grpc.ClientConnInterface) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) Dispatch(ctx context.Context, in *Object, opts ...grpc.CallOption) (*Object, error) {
	out := new(Object)
	err := c.cc.Invoke(ctx, "/coprocess.Dispatcher/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) DispatchEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventReply, error) {
	out := new(EventReply)
	err := c.cc.Invoke(ctx, "/coprocess.Dispatcher/DispatchEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServer is the server API for Dispatcher service.
type DispatcherServer interface {
	Dispatch(context.Context, *Object) (*Object, error)
	DispatchEvent(context.Context, *Event) (*EventReply, error)
}

// UnimplementedDispatcherServer can be embedded to have forward compatible implementations.
type UnimplementedDispatcherServer struct {
}

func (*UnimplementedDispatcherServer) Dispatch(context.Context, *Object) (*Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}
func (*UnimplementedDispatcherServer) DispatchEvent(context.Context, *Event) (*EventReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DispatchEvent not implemented")
}

func RegisterDispatcherServer(s *grpc.Server, srv DispatcherServer) {
	s.RegisterService(&_Dispatcher_serviceDesc, srv)
}

func _Dispatcher_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coprocess.Dispatcher/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Dispatch(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_DispatchEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).DispatchEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coprocess.Dispatcher/DispatchEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).DispatchEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coprocess.Dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _Dispatcher_Dispatch_Handler,
		},
		{
			MethodName: "DispatchEvent",
			Handler:    _Dispatcher_DispatchEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coprocess_object.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: participant.proto

package __

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

type ParticipantRequestType int32

const (
	ParticipantRequestType_PREPARE    ParticipantRequestType = 0
	ParticipantRequestType_COMMIT     ParticipantRequestType = 1
	ParticipantRequestType_PAUSE      ParticipantRequestType = 2
	ParticipantRequestType_UNPAUSE    ParticipantRequestType = 3
	ParticipantRequestType_READ       ParticipantRequestType = 4
	ParticipantRequestType_DELETE     ParticipantRequestType = 5
	ParticipantRequestType_CONNECT    ParticipantRequestType = 6
	ParticipantRequestType_DISCONNECT ParticipantRequestType = 7
)

// Enum value maps for ParticipantRequestType.
var (
	ParticipantRequestType_name = map[int32]string{
		0: "PREPARE",
		1: "COMMIT",
		2: "PAUSE",
		3: "UNPAUSE",
		4: "READ",
		5: "DELETE",
		6: "CONNECT",
		7: "DISCONNECT",
	}
	ParticipantRequestType_value = map[string]int32{
		"PREPARE":    0,
		"COMMIT":     1,
		"PAUSE":      2,
		"UNPAUSE":    3,
		"READ":       4,
		"DELETE":     5,
		"CONNECT":    6,
		"DISCONNECT": 7,
	}
)

func (x ParticipantRequestType) Enum() *ParticipantRequestType {
	p := new(ParticipantRequestType)
	*p = x
	return p
}

func (x ParticipantRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParticipantRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_participant_proto_enumTypes[0].Descriptor()
}

func (ParticipantRequestType) Type() protoreflect.EnumType {
	return &file_participant_proto_enumTypes[0]
}

func (x ParticipantRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParticipantRequestType.Descriptor instead.
func (ParticipantRequestType) EnumDescriptor() ([]byte, []int) {
	return file_participant_proto_rawDescGZIP(), []int{0}
}

type ParticipantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    ParticipantRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ParticipantRequestType" json:"type,omitempty"`
	IsAdmin bool                   `protobuf:"varint,2,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	Key     *string                `protobuf:"bytes,3,opt,name=key,proto3,oneof" json:"key,omitempty"`
	Value   *string                `protobuf:"bytes,4,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *ParticipantRequest) Reset() {
	*x = ParticipantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_participant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantRequest) ProtoMessage() {}

func (x *ParticipantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_participant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantRequest.ProtoReflect.Descriptor instead.
func (*ParticipantRequest) Descriptor() ([]byte, []int) {
	return file_participant_proto_rawDescGZIP(), []int{0}
}

func (x *ParticipantRequest) GetType() ParticipantRequestType {
	if x != nil {
		return x.Type
	}
	return ParticipantRequestType_PREPARE
}

func (x *ParticipantRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *ParticipantRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *ParticipantRequest) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type ParticipantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   ParticipantRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ParticipantRequestType" json:"type,omitempty"`
	Status bool                   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Value  *string                `protobuf:"bytes,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *ParticipantResponse) Reset() {
	*x = ParticipantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_participant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantResponse) ProtoMessage() {}

func (x *ParticipantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_participant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantResponse.ProtoReflect.Descriptor instead.
func (*ParticipantResponse) Descriptor() ([]byte, []int) {
	return file_participant_proto_rawDescGZIP(), []int{1}
}

func (x *ParticipantResponse) GetType() ParticipantRequestType {
	if x != nil {
		return x.Type
	}
	return ParticipantRequestType_PREPARE
}

func (x *ParticipantResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ParticipantResponse) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_participant_proto protoreflect.FileDescriptor

var file_participant_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x12, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x15,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x7c, 0x0a, 0x16, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x50, 0x41,
	0x55, 0x53, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x07, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_participant_proto_rawDescOnce sync.Once
	file_participant_proto_rawDescData = file_participant_proto_rawDesc
)

func file_participant_proto_rawDescGZIP() []byte {
	file_participant_proto_rawDescOnce.Do(func() {
		file_participant_proto_rawDescData = protoimpl.X.CompressGZIP(file_participant_proto_rawDescData)
	})
	return file_participant_proto_rawDescData
}

var file_participant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_participant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_participant_proto_goTypes = []any{
	(ParticipantRequestType)(0), // 0: proto.ParticipantRequestType
	(*ParticipantRequest)(nil),  // 1: proto.ParticipantRequest
	(*ParticipantResponse)(nil), // 2: proto.ParticipantResponse
}
var file_participant_proto_depIdxs = []int32{
	0, // 0: proto.ParticipantRequest.type:type_name -> proto.ParticipantRequestType
	0, // 1: proto.ParticipantResponse.type:type_name -> proto.ParticipantRequestType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_participant_proto_init() }
func file_participant_proto_init() {
	if File_participant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_participant_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ParticipantRequest); i {
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
		file_participant_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ParticipantResponse); i {
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
	file_participant_proto_msgTypes[0].OneofWrappers = []any{}
	file_participant_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_participant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_participant_proto_goTypes,
		DependencyIndexes: file_participant_proto_depIdxs,
		EnumInfos:         file_participant_proto_enumTypes,
		MessageInfos:      file_participant_proto_msgTypes,
	}.Build()
	File_participant_proto = out.File
	file_participant_proto_rawDesc = nil
	file_participant_proto_goTypes = nil
	file_participant_proto_depIdxs = nil
}

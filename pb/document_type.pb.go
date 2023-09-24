// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: document_type.proto

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

type DocumentTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt uint64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint64  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *uint64 `protobuf:"varint,5,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
}

func (x *DocumentTypeReply) Reset() {
	*x = DocumentTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentTypeReply) ProtoMessage() {}

func (x *DocumentTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_document_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentTypeReply.ProtoReflect.Descriptor instead.
func (*DocumentTypeReply) Descriptor() ([]byte, []int) {
	return file_document_type_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentTypeReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DocumentTypeReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DocumentTypeReply) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DocumentTypeReply) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DocumentTypeReply) GetDeletedAt() uint64 {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return 0
}

var File_document_type_proto protoreflect.FileDescriptor

var file_document_type_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x75, 0x61, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x11,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x32, 0x0a, 0x1f, 0x62, 0x72, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x75, 0x73, 0x2e, 0x6f, 0x73, 0x61, 0x73, 0x63, 0x6f, 0x2e, 0x64, 0x75, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x07, 0x64, 0x75, 0x61,
	0x74, 0x2f, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_document_type_proto_rawDescOnce sync.Once
	file_document_type_proto_rawDescData = file_document_type_proto_rawDesc
)

func file_document_type_proto_rawDescGZIP() []byte {
	file_document_type_proto_rawDescOnce.Do(func() {
		file_document_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_document_type_proto_rawDescData)
	})
	return file_document_type_proto_rawDescData
}

var file_document_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_document_type_proto_goTypes = []interface{}{
	(*DocumentTypeReply)(nil), // 0: duat.DocumentTypeReply
}
var file_document_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_document_type_proto_init() }
func file_document_type_proto_init() {
	if File_document_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_document_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentTypeReply); i {
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
	file_document_type_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_document_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_document_type_proto_goTypes,
		DependencyIndexes: file_document_type_proto_depIdxs,
		MessageInfos:      file_document_type_proto_msgTypes,
	}.Build()
	File_document_type_proto = out.File
	file_document_type_proto_rawDesc = nil
	file_document_type_proto_goTypes = nil
	file_document_type_proto_depIdxs = nil
}
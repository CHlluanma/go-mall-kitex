// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.4
// source: home.proto

package home

import (
	_ "github.com/CHlluanma/go-mall-kitex/app/frontend/hertz_gen/api"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{0}
}

var File_home_proto protoreflect.FileDescriptor

var file_home_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x1a, 0x09, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x48, 0x0a, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x05, 0xca, 0xc1, 0x18, 0x01, 0x2f, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x48, 0x6c, 0x6c, 0x75, 0x61, 0x6e, 0x6d,
	0x61, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x72,
	0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_proto_rawDescOnce sync.Once
	file_home_proto_rawDescData = file_home_proto_rawDesc
)

func file_home_proto_rawDescGZIP() []byte {
	file_home_proto_rawDescOnce.Do(func() {
		file_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_proto_rawDescData)
	})
	return file_home_proto_rawDescData
}

var file_home_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_home_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: frontend.home.Empty
}
var file_home_proto_depIdxs = []int32{
	0, // 0: frontend.home.HomeService.Home:input_type -> frontend.home.Empty
	0, // 1: frontend.home.HomeService.Home:output_type -> frontend.home.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_home_proto_init() }
func file_home_proto_init() {
	if File_home_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_home_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_home_proto_goTypes,
		DependencyIndexes: file_home_proto_depIdxs,
		MessageInfos:      file_home_proto_msgTypes,
	}.Build()
	File_home_proto = out.File
	file_home_proto_rawDesc = nil
	file_home_proto_goTypes = nil
	file_home_proto_depIdxs = nil
}

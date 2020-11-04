// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: test4_vendored.proto

package test4

import (
	proto "github.com/golang/protobuf/proto"
	date "google.golang.org/genproto/googleapis/type/date"
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

type M4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *date.Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *M4) Reset() {
	*x = M4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test4_vendored_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M4) ProtoMessage() {}

func (x *M4) ProtoReflect() protoreflect.Message {
	mi := &file_test4_vendored_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M4.ProtoReflect.Descriptor instead.
func (*M4) Descriptor() ([]byte, []int) {
	return file_test4_vendored_proto_rawDescGZIP(), []int{0}
}

func (x *M4) GetDate() *date.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_test4_vendored_proto protoreflect.FileDescriptor

var file_test4_vendored_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x65, 0x73, 0x74, 0x34, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x65, 0x73, 0x74, 0x34, 0x1a, 0x16, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x02, 0x4d, 0x34, 0x12, 0x25, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test4_vendored_proto_rawDescOnce sync.Once
	file_test4_vendored_proto_rawDescData = file_test4_vendored_proto_rawDesc
)

func file_test4_vendored_proto_rawDescGZIP() []byte {
	file_test4_vendored_proto_rawDescOnce.Do(func() {
		file_test4_vendored_proto_rawDescData = protoimpl.X.CompressGZIP(file_test4_vendored_proto_rawDescData)
	})
	return file_test4_vendored_proto_rawDescData
}

var file_test4_vendored_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test4_vendored_proto_goTypes = []interface{}{
	(*M4)(nil),        // 0: test4.M4
	(*date.Date)(nil), // 1: google.type.Date
}
var file_test4_vendored_proto_depIdxs = []int32{
	1, // 0: test4.M4.date:type_name -> google.type.Date
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test4_vendored_proto_init() }
func file_test4_vendored_proto_init() {
	if File_test4_vendored_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test4_vendored_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M4); i {
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
			RawDescriptor: file_test4_vendored_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test4_vendored_proto_goTypes,
		DependencyIndexes: file_test4_vendored_proto_depIdxs,
		MessageInfos:      file_test4_vendored_proto_msgTypes,
	}.Build()
	File_test4_vendored_proto = out.File
	file_test4_vendored_proto_rawDesc = nil
	file_test4_vendored_proto_goTypes = nil
	file_test4_vendored_proto_depIdxs = nil
}
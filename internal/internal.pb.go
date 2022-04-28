// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/internal.proto

package influxql

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

type Measurements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (x *Measurements) Reset() {
	*x = Measurements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurements) ProtoMessage() {}

func (x *Measurements) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurements.ProtoReflect.Descriptor instead.
func (*Measurements) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{0}
}

func (x *Measurements) GetItems() []*Measurement {
	if x != nil {
		return x.Items
	}
	return nil
}

type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database        *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name            *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex           *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget        *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_internal_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_internal_internal_proto_rawDescGZIP(), []int{1}
}

func (x *Measurement) GetDatabase() string {
	if x != nil && x.Database != nil {
		return *x.Database
	}
	return ""
}

func (x *Measurement) GetRetentionPolicy() string {
	if x != nil && x.RetentionPolicy != nil {
		return *x.RetentionPolicy
	}
	return ""
}

func (x *Measurement) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Measurement) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

func (x *Measurement) GetIsTarget() bool {
	if x != nil && x.IsTarget != nil {
		return *x.IsTarget
	}
	return false
}

var File_internal_internal_proto protoreflect.FileDescriptor

var file_internal_internal_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e, 0x66, 0x6c, 0x75,
	0x78, 0x71, 0x6c, 0x22, 0x3b, 0x0a, 0x0c, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x71, 0x6c, 0x2e, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x99, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65,
	0x67, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x71, 0x6c,
}

var (
	file_internal_internal_proto_rawDescOnce sync.Once
	file_internal_internal_proto_rawDescData = file_internal_internal_proto_rawDesc
)

func file_internal_internal_proto_rawDescGZIP() []byte {
	file_internal_internal_proto_rawDescOnce.Do(func() {
		file_internal_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_internal_proto_rawDescData)
	})
	return file_internal_internal_proto_rawDescData
}

var file_internal_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_internal_proto_goTypes = []interface{}{
	(*Measurements)(nil), // 0: influxql.Measurements
	(*Measurement)(nil),  // 1: influxql.Measurement
}
var file_internal_internal_proto_depIdxs = []int32{
	1, // 0: influxql.Measurements.Items:type_name -> influxql.Measurement
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_internal_proto_init() }
func file_internal_internal_proto_init() {
	if File_internal_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurements); i {
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
		file_internal_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
			RawDescriptor: file_internal_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_internal_proto_goTypes,
		DependencyIndexes: file_internal_internal_proto_depIdxs,
		MessageInfos:      file_internal_internal_proto_msgTypes,
	}.Build()
	File_internal_internal_proto = out.File
	file_internal_internal_proto_rawDesc = nil
	file_internal_internal_proto_goTypes = nil
	file_internal_internal_proto_depIdxs = nil
}

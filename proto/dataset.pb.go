//
// Copyright (C) HereweTech, Inc - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: dataset.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// {{{ [Measuring instrument]
type MeasuringInstrument struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ident         string                 `protobuf:"bytes,3,opt,name=ident,proto3" json:"ident,omitempty"`
	Category      string                 `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,127,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MeasuringInstrument) Reset() {
	*x = MeasuringInstrument{}
	mi := &file_dataset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeasuringInstrument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasuringInstrument) ProtoMessage() {}

func (x *MeasuringInstrument) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasuringInstrument.ProtoReflect.Descriptor instead.
func (*MeasuringInstrument) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *MeasuringInstrument) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MeasuringInstrument) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeasuringInstrument) GetIdent() string {
	if x != nil {
		return x.Ident
	}
	return ""
}

func (x *MeasuringInstrument) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *MeasuringInstrument) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type AppendMeasureReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Detail        string                 `protobuf:"bytes,127,opt,name=detail,proto3" json:"detail,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppendMeasureReq) Reset() {
	*x = AppendMeasureReq{}
	mi := &file_dataset_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppendMeasureReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendMeasureReq) ProtoMessage() {}

func (x *AppendMeasureReq) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendMeasureReq.ProtoReflect.Descriptor instead.
func (*AppendMeasureReq) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{1}
}

func (x *AppendMeasureReq) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *AppendMeasureReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AppendMeasureReq) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type AppendMeasureResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppendMeasureResp) Reset() {
	*x = AppendMeasureResp{}
	mi := &file_dataset_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppendMeasureResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendMeasureResp) ProtoMessage() {}

func (x *AppendMeasureResp) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendMeasureResp.ProtoReflect.Descriptor instead.
func (*AppendMeasureResp) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{2}
}

func (x *AppendMeasureResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *AppendMeasureResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListInstrumentsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInstrumentsReq) Reset() {
	*x = ListInstrumentsReq{}
	mi := &file_dataset_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstrumentsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstrumentsReq) ProtoMessage() {}

func (x *ListInstrumentsReq) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstrumentsReq.ProtoReflect.Descriptor instead.
func (*ListInstrumentsReq) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{3}
}

type ListInstrumentsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*MeasuringInstrument `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInstrumentsResp) Reset() {
	*x = ListInstrumentsResp{}
	mi := &file_dataset_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstrumentsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstrumentsResp) ProtoMessage() {}

func (x *ListInstrumentsResp) ProtoReflect() protoreflect.Message {
	mi := &file_dataset_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstrumentsResp.ProtoReflect.Descriptor instead.
func (*ListInstrumentsResp) Descriptor() ([]byte, []int) {
	return file_dataset_proto_rawDescGZIP(), []int{4}
}

func (x *ListInstrumentsResp) GetList() []*MeasuringInstrument {
	if x != nil {
		return x.List
	}
	return nil
}

var File_dataset_proto protoreflect.FileDescriptor

const file_dataset_proto_rawDesc = "" +
	"\n" +
	"\rdataset.proto\x12\asvc.api\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa6\x01\n" +
	"\x13MeasuringInstrument\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05ident\x18\x03 \x01(\tR\x05ident\x12\x1a\n" +
	"\bcategory\x18\x04 \x01(\tR\bcategory\x129\n" +
	"\n" +
	"created_at\x18\x7f \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"]\n" +
	"\x10AppendMeasureReq\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\x12\x16\n" +
	"\x06detail\x18\x7f \x01(\tR\x06detail\"E\n" +
	"\x11AppendMeasureResp\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\x14\n" +
	"\x12ListInstrumentsReq\"G\n" +
	"\x13ListInstrumentsResp\x120\n" +
	"\x04list\x18\x01 \x03(\v2\x1c.svc.api.MeasuringInstrumentR\x04list2\x9a\x01\n" +
	"\aDataset\x12F\n" +
	"\rAppendMeasure\x12\x19.svc.api.AppendMeasureReq\x1a\x1a.svc.api.AppendMeasureResp\x12G\n" +
	"\x0fListInstruments\x12\x16.google.protobuf.Empty\x1a\x1c.svc.api.ListInstrumentsRespB\x0fZ\r./proto;protob\x06proto3"

var (
	file_dataset_proto_rawDescOnce sync.Once
	file_dataset_proto_rawDescData []byte
)

func file_dataset_proto_rawDescGZIP() []byte {
	file_dataset_proto_rawDescOnce.Do(func() {
		file_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dataset_proto_rawDesc), len(file_dataset_proto_rawDesc)))
	})
	return file_dataset_proto_rawDescData
}

var file_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dataset_proto_goTypes = []any{
	(*MeasuringInstrument)(nil),   // 0: svc.api.MeasuringInstrument
	(*AppendMeasureReq)(nil),      // 1: svc.api.AppendMeasureReq
	(*AppendMeasureResp)(nil),     // 2: svc.api.AppendMeasureResp
	(*ListInstrumentsReq)(nil),    // 3: svc.api.ListInstrumentsReq
	(*ListInstrumentsResp)(nil),   // 4: svc.api.ListInstrumentsResp
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_dataset_proto_depIdxs = []int32{
	5, // 0: svc.api.MeasuringInstrument.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: svc.api.ListInstrumentsResp.list:type_name -> svc.api.MeasuringInstrument
	1, // 2: svc.api.Dataset.AppendMeasure:input_type -> svc.api.AppendMeasureReq
	6, // 3: svc.api.Dataset.ListInstruments:input_type -> google.protobuf.Empty
	2, // 4: svc.api.Dataset.AppendMeasure:output_type -> svc.api.AppendMeasureResp
	4, // 5: svc.api.Dataset.ListInstruments:output_type -> svc.api.ListInstrumentsResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dataset_proto_init() }
func file_dataset_proto_init() {
	if File_dataset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dataset_proto_rawDesc), len(file_dataset_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dataset_proto_goTypes,
		DependencyIndexes: file_dataset_proto_depIdxs,
		MessageInfos:      file_dataset_proto_msgTypes,
	}.Build()
	File_dataset_proto = out.File
	file_dataset_proto_goTypes = nil
	file_dataset_proto_depIdxs = nil
}

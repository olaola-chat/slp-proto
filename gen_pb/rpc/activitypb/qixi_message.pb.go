// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: rpc/activity/qixi_message.proto

package activitypb

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

type ReqQixiMyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReqQixiMyData) Reset() {
	*x = ReqQixiMyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_activity_qixi_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQixiMyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQixiMyData) ProtoMessage() {}

func (x *ReqQixiMyData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_activity_qixi_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQixiMyData.ProtoReflect.Descriptor instead.
func (*ReqQixiMyData) Descriptor() ([]byte, []int) {
	return file_rpc_activity_qixi_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReqQixiMyData) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type RepQixiMyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	SentSum uint32 `protobuf:"varint,3,opt,name=sent_sum,json=sentSum,proto3" json:"sent_sum,omitempty"`
	RecvSum uint32 `protobuf:"varint,4,opt,name=recv_sum,json=recvSum,proto3" json:"recv_sum,omitempty"`
	Rank    uint32 `protobuf:"varint,5,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *RepQixiMyData) Reset() {
	*x = RepQixiMyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_activity_qixi_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepQixiMyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepQixiMyData) ProtoMessage() {}

func (x *RepQixiMyData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_activity_qixi_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepQixiMyData.ProtoReflect.Descriptor instead.
func (*RepQixiMyData) Descriptor() ([]byte, []int) {
	return file_rpc_activity_qixi_message_proto_rawDescGZIP(), []int{1}
}

func (x *RepQixiMyData) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RepQixiMyData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RepQixiMyData) GetSentSum() uint32 {
	if x != nil {
		return x.SentSum
	}
	return 0
}

func (x *RepQixiMyData) GetRecvSum() uint32 {
	if x != nil {
		return x.RecvSum
	}
	return 0
}

func (x *RepQixiMyData) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type ReqQixiAwardPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReqQixiAwardPool) Reset() {
	*x = ReqQixiAwardPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_activity_qixi_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqQixiAwardPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqQixiAwardPool) ProtoMessage() {}

func (x *ReqQixiAwardPool) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_activity_qixi_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqQixiAwardPool.ProtoReflect.Descriptor instead.
func (*ReqQixiAwardPool) Descriptor() ([]byte, []int) {
	return file_rpc_activity_qixi_message_proto_rawDescGZIP(), []int{2}
}

func (x *ReqQixiAwardPool) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type RepQixiAwardPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg       string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	SumLv1    uint32   `protobuf:"varint,3,opt,name=sum_lv1,json=sumLv1,proto3" json:"sum_lv1,omitempty"`
	SumLv2    uint32   `protobuf:"varint,4,opt,name=sum_lv2,json=sumLv2,proto3" json:"sum_lv2,omitempty"`
	AwardCnts []uint32 `protobuf:"varint,5,rep,packed,name=award_cnts,json=awardCnts,proto3" json:"award_cnts,omitempty"`
}

func (x *RepQixiAwardPool) Reset() {
	*x = RepQixiAwardPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_activity_qixi_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepQixiAwardPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepQixiAwardPool) ProtoMessage() {}

func (x *RepQixiAwardPool) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_activity_qixi_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepQixiAwardPool.ProtoReflect.Descriptor instead.
func (*RepQixiAwardPool) Descriptor() ([]byte, []int) {
	return file_rpc_activity_qixi_message_proto_rawDescGZIP(), []int{3}
}

func (x *RepQixiAwardPool) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RepQixiAwardPool) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RepQixiAwardPool) GetSumLv1() uint32 {
	if x != nil {
		return x.SumLv1
	}
	return 0
}

func (x *RepQixiAwardPool) GetSumLv2() uint32 {
	if x != nil {
		return x.SumLv2
	}
	return 0
}

func (x *RepQixiAwardPool) GetAwardCnts() []uint32 {
	if x != nil {
		return x.AwardCnts
	}
	return nil
}

var File_rpc_activity_qixi_message_proto protoreflect.FileDescriptor

var file_rpc_activity_qixi_message_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x71,
	0x69, 0x78, 0x69, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x72, 0x62, 0x70, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x22, 0x21, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x51, 0x69, 0x78, 0x69, 0x4d, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x51, 0x69,
	0x78, 0x69, 0x4d, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x63, 0x76, 0x53, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x24,
	0x0a, 0x10, 0x52, 0x65, 0x71, 0x51, 0x69, 0x78, 0x69, 0x41, 0x77, 0x61, 0x72, 0x64, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x51, 0x69, 0x78, 0x69,
	0x41, 0x77, 0x61, 0x72, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x5f, 0x6c, 0x76, 0x31,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x75, 0x6d, 0x4c, 0x76, 0x31, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x75, 0x6d, 0x5f, 0x6c, 0x76, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x75, 0x6d, 0x4c, 0x76, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x63, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x77, 0x61,
	0x72, 0x64, 0x43, 0x6e, 0x74, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x61, 0x6f, 0x6c, 0x61, 0x2d, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x72, 0x62, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_activity_qixi_message_proto_rawDescOnce sync.Once
	file_rpc_activity_qixi_message_proto_rawDescData = file_rpc_activity_qixi_message_proto_rawDesc
)

func file_rpc_activity_qixi_message_proto_rawDescGZIP() []byte {
	file_rpc_activity_qixi_message_proto_rawDescOnce.Do(func() {
		file_rpc_activity_qixi_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_activity_qixi_message_proto_rawDescData)
	})
	return file_rpc_activity_qixi_message_proto_rawDescData
}

var file_rpc_activity_qixi_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_activity_qixi_message_proto_goTypes = []interface{}{
	(*ReqQixiMyData)(nil),    // 0: slp.rpc.activity.ReqQixiMyData
	(*RepQixiMyData)(nil),    // 1: slp.rpc.activity.RepQixiMyData
	(*ReqQixiAwardPool)(nil), // 2: slp.rpc.activity.ReqQixiAwardPool
	(*RepQixiAwardPool)(nil), // 3: slp.rpc.activity.RepQixiAwardPool
}
var file_rpc_activity_qixi_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_activity_qixi_message_proto_init() }
func file_rpc_activity_qixi_message_proto_init() {
	if File_rpc_activity_qixi_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_activity_qixi_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQixiMyData); i {
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
		file_rpc_activity_qixi_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepQixiMyData); i {
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
		file_rpc_activity_qixi_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqQixiAwardPool); i {
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
		file_rpc_activity_qixi_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepQixiAwardPool); i {
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
			RawDescriptor: file_rpc_activity_qixi_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_activity_qixi_message_proto_goTypes,
		DependencyIndexes: file_rpc_activity_qixi_message_proto_depIdxs,
		MessageInfos:      file_rpc_activity_qixi_message_proto_msgTypes,
	}.Build()
	File_rpc_activity_qixi_message_proto = out.File
	file_rpc_activity_qixi_message_proto_rawDesc = nil
	file_rpc_activity_qixi_message_proto_goTypes = nil
	file_rpc_activity_qixi_message_proto_depIdxs = nil
}

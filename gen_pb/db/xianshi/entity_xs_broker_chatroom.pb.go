// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// 重要! 增加数据库字段，只能向后增加，否则不能使用工具生成proto
// ==========================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: db/xianshi/entity_xs_broker_chatroom.proto

package xianshi

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

type EntityXsBrokerChatroom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: orm:"id,primary"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" orm:"id,primary"` //
	//@inject_tag: orm:"bid"
	Bid int32 `protobuf:"varint,2,opt,name=bid,proto3" json:"bid,omitempty" orm:"bid"` // 公会id
	//@inject_tag: orm:"rid"
	Rid int32 `protobuf:"varint,3,opt,name=rid,proto3" json:"rid,omitempty" orm:"rid"` // 房间id
	//@inject_tag: orm:"dateline"
	Dateline int32 `protobuf:"varint,4,opt,name=dateline,proto3" json:"dateline,omitempty" orm:"dateline"` // 加入时间
}

func (x *EntityXsBrokerChatroom) Reset() {
	*x = EntityXsBrokerChatroom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_xianshi_entity_xs_broker_chatroom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityXsBrokerChatroom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityXsBrokerChatroom) ProtoMessage() {}

func (x *EntityXsBrokerChatroom) ProtoReflect() protoreflect.Message {
	mi := &file_db_xianshi_entity_xs_broker_chatroom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityXsBrokerChatroom.ProtoReflect.Descriptor instead.
func (*EntityXsBrokerChatroom) Descriptor() ([]byte, []int) {
	return file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescGZIP(), []int{0}
}

func (x *EntityXsBrokerChatroom) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EntityXsBrokerChatroom) GetBid() int32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *EntityXsBrokerChatroom) GetRid() int32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *EntityXsBrokerChatroom) GetDateline() int32 {
	if x != nil {
		return x.Dateline
	}
	return 0
}

var File_db_xianshi_entity_xs_broker_chatroom_proto protoreflect.FileDescriptor

var file_db_xianshi_entity_xs_broker_chatroom_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x64, 0x62, 0x2f, 0x78, 0x69, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x78, 0x73, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x62,
	0x2e, 0x78, 0x69, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x22, 0x68, 0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x58, 0x73, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x62, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x6c, 0x61, 0x6f, 0x6c, 0x61, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x72, 0x62, 0x70,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x62, 0x2f, 0x64, 0x62,
	0x2f, 0x78, 0x69, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescOnce sync.Once
	file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescData = file_db_xianshi_entity_xs_broker_chatroom_proto_rawDesc
)

func file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescGZIP() []byte {
	file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescOnce.Do(func() {
		file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescData)
	})
	return file_db_xianshi_entity_xs_broker_chatroom_proto_rawDescData
}

var file_db_xianshi_entity_xs_broker_chatroom_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_xianshi_entity_xs_broker_chatroom_proto_goTypes = []interface{}{
	(*EntityXsBrokerChatroom)(nil), // 0: db.xianshi.EntityXsBrokerChatroom
}
var file_db_xianshi_entity_xs_broker_chatroom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_db_xianshi_entity_xs_broker_chatroom_proto_init() }
func file_db_xianshi_entity_xs_broker_chatroom_proto_init() {
	if File_db_xianshi_entity_xs_broker_chatroom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_xianshi_entity_xs_broker_chatroom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityXsBrokerChatroom); i {
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
			RawDescriptor: file_db_xianshi_entity_xs_broker_chatroom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_xianshi_entity_xs_broker_chatroom_proto_goTypes,
		DependencyIndexes: file_db_xianshi_entity_xs_broker_chatroom_proto_depIdxs,
		MessageInfos:      file_db_xianshi_entity_xs_broker_chatroom_proto_msgTypes,
	}.Build()
	File_db_xianshi_entity_xs_broker_chatroom_proto = out.File
	file_db_xianshi_entity_xs_broker_chatroom_proto_rawDesc = nil
	file_db_xianshi_entity_xs_broker_chatroom_proto_goTypes = nil
	file_db_xianshi_entity_xs_broker_chatroom_proto_depIdxs = nil
}

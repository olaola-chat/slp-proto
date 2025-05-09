// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// 重要! 增加数据库字段，只能向后增加，否则不能使用工具生成proto
// ==========================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: db/xianshi/entity_xs_chatroom_package.proto

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

type EntityXsChatroomPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: orm:"id,primary"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" orm:"id,primary"` //
	//@inject_tag: orm:"rid"
	Rid uint32 `protobuf:"varint,2,opt,name=rid,proto3" json:"rid,omitempty" orm:"rid"` //
	//@inject_tag: orm:"bid"
	Bid uint32 `protobuf:"varint,3,opt,name=bid,proto3" json:"bid,omitempty" orm:"bid"` // 老板ID
	//@inject_tag: orm:"sender"
	Sender uint32 `protobuf:"varint,4,opt,name=sender,proto3" json:"sender,omitempty" orm:"sender"` //
	//@inject_tag: orm:"uid"
	Uid uint32 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty" orm:"uid"` //
	//@inject_tag: orm:"money"
	Money uint32 `protobuf:"varint,6,opt,name=money,proto3" json:"money,omitempty" orm:"money"` //
	//@inject_tag: orm:"dateline,primary"
	Dateline uint32 `protobuf:"varint,7,opt,name=dateline,proto3" json:"dateline,omitempty" orm:"dateline,primary"` //
}

func (x *EntityXsChatroomPackage) Reset() {
	*x = EntityXsChatroomPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_xianshi_entity_xs_chatroom_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityXsChatroomPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityXsChatroomPackage) ProtoMessage() {}

func (x *EntityXsChatroomPackage) ProtoReflect() protoreflect.Message {
	mi := &file_db_xianshi_entity_xs_chatroom_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityXsChatroomPackage.ProtoReflect.Descriptor instead.
func (*EntityXsChatroomPackage) Descriptor() ([]byte, []int) {
	return file_db_xianshi_entity_xs_chatroom_package_proto_rawDescGZIP(), []int{0}
}

func (x *EntityXsChatroomPackage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EntityXsChatroomPackage) GetRid() uint32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *EntityXsChatroomPackage) GetBid() uint32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *EntityXsChatroomPackage) GetSender() uint32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *EntityXsChatroomPackage) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *EntityXsChatroomPackage) GetMoney() uint32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *EntityXsChatroomPackage) GetDateline() uint32 {
	if x != nil {
		return x.Dateline
	}
	return 0
}

var File_db_xianshi_entity_xs_chatroom_package_proto protoreflect.FileDescriptor

var file_db_xianshi_entity_xs_chatroom_package_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x64, 0x62, 0x2f, 0x78, 0x69, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x78, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64,
	0x62, 0x2e, 0x78, 0x69, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x22, 0xa9, 0x01, 0x0a, 0x17, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x58, 0x73, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x61, 0x6f, 0x6c, 0x61, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x72, 0x62, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x62,
	0x2f, 0x64, 0x62, 0x2f, 0x78, 0x69, 0x61, 0x6e, 0x73, 0x68, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_db_xianshi_entity_xs_chatroom_package_proto_rawDescOnce sync.Once
	file_db_xianshi_entity_xs_chatroom_package_proto_rawDescData = file_db_xianshi_entity_xs_chatroom_package_proto_rawDesc
)

func file_db_xianshi_entity_xs_chatroom_package_proto_rawDescGZIP() []byte {
	file_db_xianshi_entity_xs_chatroom_package_proto_rawDescOnce.Do(func() {
		file_db_xianshi_entity_xs_chatroom_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_xianshi_entity_xs_chatroom_package_proto_rawDescData)
	})
	return file_db_xianshi_entity_xs_chatroom_package_proto_rawDescData
}

var file_db_xianshi_entity_xs_chatroom_package_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_xianshi_entity_xs_chatroom_package_proto_goTypes = []interface{}{
	(*EntityXsChatroomPackage)(nil), // 0: db.xianshi.EntityXsChatroomPackage
}
var file_db_xianshi_entity_xs_chatroom_package_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_db_xianshi_entity_xs_chatroom_package_proto_init() }
func file_db_xianshi_entity_xs_chatroom_package_proto_init() {
	if File_db_xianshi_entity_xs_chatroom_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_xianshi_entity_xs_chatroom_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityXsChatroomPackage); i {
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
			RawDescriptor: file_db_xianshi_entity_xs_chatroom_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_xianshi_entity_xs_chatroom_package_proto_goTypes,
		DependencyIndexes: file_db_xianshi_entity_xs_chatroom_package_proto_depIdxs,
		MessageInfos:      file_db_xianshi_entity_xs_chatroom_package_proto_msgTypes,
	}.Build()
	File_db_xianshi_entity_xs_chatroom_package_proto = out.File
	file_db_xianshi_entity_xs_chatroom_package_proto_rawDesc = nil
	file_db_xianshi_entity_xs_chatroom_package_proto_goTypes = nil
	file_db_xianshi_entity_xs_chatroom_package_proto_depIdxs = nil
}

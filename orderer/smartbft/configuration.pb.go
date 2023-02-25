// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: orderer/smartbft/configuration.proto

package smartbft

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

type Options_Rotation int32

const (
	Options_ROTATION_UNSPECIFIED Options_Rotation = 0
	Options_ROTATION_OFF         Options_Rotation = 1
	Options_ROTATION_ON          Options_Rotation = 2
)

// Enum value maps for Options_Rotation.
var (
	Options_Rotation_name = map[int32]string{
		0: "ROTATION_UNSPECIFIED",
		1: "ROTATION_OFF",
		2: "ROTATION_ON",
	}
	Options_Rotation_value = map[string]int32{
		"ROTATION_UNSPECIFIED": 0,
		"ROTATION_OFF":         1,
		"ROTATION_ON":          2,
	}
)

func (x Options_Rotation) Enum() *Options_Rotation {
	p := new(Options_Rotation)
	*p = x
	return p
}

func (x Options_Rotation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Options_Rotation) Descriptor() protoreflect.EnumDescriptor {
	return file_orderer_smartbft_configuration_proto_enumTypes[0].Descriptor()
}

func (Options_Rotation) Type() protoreflect.EnumType {
	return &file_orderer_smartbft_configuration_proto_enumTypes[0]
}

func (x Options_Rotation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Options_Rotation.Descriptor instead.
func (Options_Rotation) EnumDescriptor() ([]byte, []int) {
	return file_orderer_smartbft_configuration_proto_rawDescGZIP(), []int{0, 0}
}

// Options to be specified for all the smartbft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestBatchMaxCount      uint64           `protobuf:"varint,1,opt,name=request_batch_max_count,json=requestBatchMaxCount,proto3" json:"request_batch_max_count,omitempty"`
	RequestBatchMaxBytes      uint64           `protobuf:"varint,2,opt,name=request_batch_max_bytes,json=requestBatchMaxBytes,proto3" json:"request_batch_max_bytes,omitempty"`
	RequestBatchMaxInterval   string           `protobuf:"bytes,3,opt,name=request_batch_max_interval,json=requestBatchMaxInterval,proto3" json:"request_batch_max_interval,omitempty"`
	IncomingMessageBufferSize uint64           `protobuf:"varint,4,opt,name=incoming_message_buffer_size,json=incomingMessageBufferSize,proto3" json:"incoming_message_buffer_size,omitempty"`
	RequestPoolSize           uint64           `protobuf:"varint,5,opt,name=request_pool_size,json=requestPoolSize,proto3" json:"request_pool_size,omitempty"`
	RequestForwardTimeout     string           `protobuf:"bytes,6,opt,name=request_forward_timeout,json=requestForwardTimeout,proto3" json:"request_forward_timeout,omitempty"`
	RequestComplainTimeout    string           `protobuf:"bytes,7,opt,name=request_complain_timeout,json=requestComplainTimeout,proto3" json:"request_complain_timeout,omitempty"`
	RequestAutoRemoveTimeout  string           `protobuf:"bytes,8,opt,name=request_auto_remove_timeout,json=requestAutoRemoveTimeout,proto3" json:"request_auto_remove_timeout,omitempty"`
	RequestMaxBytes           uint64           `protobuf:"varint,9,opt,name=request_max_bytes,json=requestMaxBytes,proto3" json:"request_max_bytes,omitempty"`
	ViewChangeResendInterval  string           `protobuf:"bytes,10,opt,name=view_change_resend_interval,json=viewChangeResendInterval,proto3" json:"view_change_resend_interval,omitempty"`
	ViewChangeTimeout         string           `protobuf:"bytes,11,opt,name=view_change_timeout,json=viewChangeTimeout,proto3" json:"view_change_timeout,omitempty"`
	LeaderHeartbeatTimeout    string           `protobuf:"bytes,12,opt,name=leader_heartbeat_timeout,json=leaderHeartbeatTimeout,proto3" json:"leader_heartbeat_timeout,omitempty"`
	LeaderHeartbeatCount      uint64           `protobuf:"varint,13,opt,name=leader_heartbeat_count,json=leaderHeartbeatCount,proto3" json:"leader_heartbeat_count,omitempty"`
	CollectTimeout            string           `protobuf:"bytes,14,opt,name=collect_timeout,json=collectTimeout,proto3" json:"collect_timeout,omitempty"`
	SyncOnStart               bool             `protobuf:"varint,15,opt,name=sync_on_start,json=syncOnStart,proto3" json:"sync_on_start,omitempty"`
	SpeedUpViewChange         bool             `protobuf:"varint,16,opt,name=speed_up_view_change,json=speedUpViewChange,proto3" json:"speed_up_view_change,omitempty"`
	LeaderRotation            Options_Rotation `protobuf:"varint,17,opt,name=leader_rotation,json=leaderRotation,proto3,enum=orderer.smartbft.Options_Rotation" json:"leader_rotation,omitempty"`
	DecisionsPerLeader        uint64           `protobuf:"varint,18,opt,name=decisions_per_leader,json=decisionsPerLeader,proto3" json:"decisions_per_leader,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_smartbft_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_smartbft_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_orderer_smartbft_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetRequestBatchMaxCount() uint64 {
	if x != nil {
		return x.RequestBatchMaxCount
	}
	return 0
}

func (x *Options) GetRequestBatchMaxBytes() uint64 {
	if x != nil {
		return x.RequestBatchMaxBytes
	}
	return 0
}

func (x *Options) GetRequestBatchMaxInterval() string {
	if x != nil {
		return x.RequestBatchMaxInterval
	}
	return ""
}

func (x *Options) GetIncomingMessageBufferSize() uint64 {
	if x != nil {
		return x.IncomingMessageBufferSize
	}
	return 0
}

func (x *Options) GetRequestPoolSize() uint64 {
	if x != nil {
		return x.RequestPoolSize
	}
	return 0
}

func (x *Options) GetRequestForwardTimeout() string {
	if x != nil {
		return x.RequestForwardTimeout
	}
	return ""
}

func (x *Options) GetRequestComplainTimeout() string {
	if x != nil {
		return x.RequestComplainTimeout
	}
	return ""
}

func (x *Options) GetRequestAutoRemoveTimeout() string {
	if x != nil {
		return x.RequestAutoRemoveTimeout
	}
	return ""
}

func (x *Options) GetRequestMaxBytes() uint64 {
	if x != nil {
		return x.RequestMaxBytes
	}
	return 0
}

func (x *Options) GetViewChangeResendInterval() string {
	if x != nil {
		return x.ViewChangeResendInterval
	}
	return ""
}

func (x *Options) GetViewChangeTimeout() string {
	if x != nil {
		return x.ViewChangeTimeout
	}
	return ""
}

func (x *Options) GetLeaderHeartbeatTimeout() string {
	if x != nil {
		return x.LeaderHeartbeatTimeout
	}
	return ""
}

func (x *Options) GetLeaderHeartbeatCount() uint64 {
	if x != nil {
		return x.LeaderHeartbeatCount
	}
	return 0
}

func (x *Options) GetCollectTimeout() string {
	if x != nil {
		return x.CollectTimeout
	}
	return ""
}

func (x *Options) GetSyncOnStart() bool {
	if x != nil {
		return x.SyncOnStart
	}
	return false
}

func (x *Options) GetSpeedUpViewChange() bool {
	if x != nil {
		return x.SpeedUpViewChange
	}
	return false
}

func (x *Options) GetLeaderRotation() Options_Rotation {
	if x != nil {
		return x.LeaderRotation
	}
	return Options_ROTATION_UNSPECIFIED
}

func (x *Options) GetDecisionsPerLeader() uint64 {
	if x != nil {
		return x.DecisionsPerLeader
	}
	return 0
}

var File_orderer_smartbft_configuration_proto protoreflect.FileDescriptor

var file_orderer_smartbft_configuration_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62,
	0x66, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x22, 0xa3, 0x08, 0x0a, 0x07, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x3f, 0x0a, 0x1c, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x19, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x17,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d,
	0x0a, 0x1b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x61, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18,
	0x76, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x14, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x4f, 0x6e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x75,
	0x70, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x70, 0x65, 0x65, 0x64, 0x55, 0x70, 0x56, 0x69, 0x65, 0x77,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62,
	0x66, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x12, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x08, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0xe5,
	0x01, 0x0a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66,
	0x74, 0x42, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f,
	0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f,
	0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0xa2, 0x02, 0x03, 0x4f, 0x53, 0x58, 0xaa, 0x02, 0x10,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x62, 0x66, 0x74,
	0xca, 0x02, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5c, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x62, 0x66, 0x74, 0xe2, 0x02, 0x1c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5c, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x62, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_smartbft_configuration_proto_rawDescOnce sync.Once
	file_orderer_smartbft_configuration_proto_rawDescData = file_orderer_smartbft_configuration_proto_rawDesc
)

func file_orderer_smartbft_configuration_proto_rawDescGZIP() []byte {
	file_orderer_smartbft_configuration_proto_rawDescOnce.Do(func() {
		file_orderer_smartbft_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_smartbft_configuration_proto_rawDescData)
	})
	return file_orderer_smartbft_configuration_proto_rawDescData
}

var file_orderer_smartbft_configuration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orderer_smartbft_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_orderer_smartbft_configuration_proto_goTypes = []interface{}{
	(Options_Rotation)(0), // 0: orderer.smartbft.Options.Rotation
	(*Options)(nil),       // 1: orderer.smartbft.Options
}
var file_orderer_smartbft_configuration_proto_depIdxs = []int32{
	0, // 0: orderer.smartbft.Options.leader_rotation:type_name -> orderer.smartbft.Options.Rotation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orderer_smartbft_configuration_proto_init() }
func file_orderer_smartbft_configuration_proto_init() {
	if File_orderer_smartbft_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_smartbft_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_orderer_smartbft_configuration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderer_smartbft_configuration_proto_goTypes,
		DependencyIndexes: file_orderer_smartbft_configuration_proto_depIdxs,
		EnumInfos:         file_orderer_smartbft_configuration_proto_enumTypes,
		MessageInfos:      file_orderer_smartbft_configuration_proto_msgTypes,
	}.Build()
	File_orderer_smartbft_configuration_proto = out.File
	file_orderer_smartbft_configuration_proto_rawDesc = nil
	file_orderer_smartbft_configuration_proto_goTypes = nil
	file_orderer_smartbft_configuration_proto_depIdxs = nil
}
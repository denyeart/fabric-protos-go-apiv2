// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: orderer/blockattestation.proto

package orderer

import (
	common "github.com/hyperledger/fabric-protos-go-apiv2/common"
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

type BlockAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *common.BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Metadata *common.BlockMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *BlockAttestation) Reset() {
	*x = BlockAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_blockattestation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockAttestation) ProtoMessage() {}

func (x *BlockAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_blockattestation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockAttestation.ProtoReflect.Descriptor instead.
func (*BlockAttestation) Descriptor() ([]byte, []int) {
	return file_orderer_blockattestation_proto_rawDescGZIP(), []int{0}
}

func (x *BlockAttestation) GetHeader() *common.BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BlockAttestation) GetMetadata() *common.BlockMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type BlockAttestationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*BlockAttestationResponse_Status
	//	*BlockAttestationResponse_BlockAttestation
	Type isBlockAttestationResponse_Type `protobuf_oneof:"Type"`
}

func (x *BlockAttestationResponse) Reset() {
	*x = BlockAttestationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_blockattestation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockAttestationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockAttestationResponse) ProtoMessage() {}

func (x *BlockAttestationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_blockattestation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockAttestationResponse.ProtoReflect.Descriptor instead.
func (*BlockAttestationResponse) Descriptor() ([]byte, []int) {
	return file_orderer_blockattestation_proto_rawDescGZIP(), []int{1}
}

func (m *BlockAttestationResponse) GetType() isBlockAttestationResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *BlockAttestationResponse) GetStatus() common.Status {
	if x, ok := x.GetType().(*BlockAttestationResponse_Status); ok {
		return x.Status
	}
	return common.Status(0)
}

func (x *BlockAttestationResponse) GetBlockAttestation() *BlockAttestation {
	if x, ok := x.GetType().(*BlockAttestationResponse_BlockAttestation); ok {
		return x.BlockAttestation
	}
	return nil
}

type isBlockAttestationResponse_Type interface {
	isBlockAttestationResponse_Type()
}

type BlockAttestationResponse_Status struct {
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status,oneof"`
}

type BlockAttestationResponse_BlockAttestation struct {
	BlockAttestation *BlockAttestation `protobuf:"bytes,2,opt,name=block_attestation,json=blockAttestation,proto3,oneof"`
}

func (*BlockAttestationResponse_Status) isBlockAttestationResponse_Type() {}

func (*BlockAttestationResponse_BlockAttestation) isBlockAttestationResponse_Type() {}

var File_orderer_blockattestation_proto protoreflect.FileDescriptor

var file_orderer_blockattestation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72,
	0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x96, 0x01, 0x0a, 0x18, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x11, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x32, 0x5f, 0x0a, 0x11, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4a, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0xb1, 0x01, 0x0a,
	0x25, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0x15, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xca, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0xe2, 0x02, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_blockattestation_proto_rawDescOnce sync.Once
	file_orderer_blockattestation_proto_rawDescData = file_orderer_blockattestation_proto_rawDesc
)

func file_orderer_blockattestation_proto_rawDescGZIP() []byte {
	file_orderer_blockattestation_proto_rawDescOnce.Do(func() {
		file_orderer_blockattestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_blockattestation_proto_rawDescData)
	})
	return file_orderer_blockattestation_proto_rawDescData
}

var file_orderer_blockattestation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orderer_blockattestation_proto_goTypes = []interface{}{
	(*BlockAttestation)(nil),         // 0: orderer.BlockAttestation
	(*BlockAttestationResponse)(nil), // 1: orderer.BlockAttestationResponse
	(*common.BlockHeader)(nil),       // 2: common.BlockHeader
	(*common.BlockMetadata)(nil),     // 3: common.BlockMetadata
	(common.Status)(0),               // 4: common.Status
	(*common.Envelope)(nil),          // 5: common.Envelope
}
var file_orderer_blockattestation_proto_depIdxs = []int32{
	2, // 0: orderer.BlockAttestation.header:type_name -> common.BlockHeader
	3, // 1: orderer.BlockAttestation.metadata:type_name -> common.BlockMetadata
	4, // 2: orderer.BlockAttestationResponse.status:type_name -> common.Status
	0, // 3: orderer.BlockAttestationResponse.block_attestation:type_name -> orderer.BlockAttestation
	5, // 4: orderer.BlockAttestations.BlockAttestations:input_type -> common.Envelope
	1, // 5: orderer.BlockAttestations.BlockAttestations:output_type -> orderer.BlockAttestationResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orderer_blockattestation_proto_init() }
func file_orderer_blockattestation_proto_init() {
	if File_orderer_blockattestation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_blockattestation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockAttestation); i {
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
		file_orderer_blockattestation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockAttestationResponse); i {
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
	file_orderer_blockattestation_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BlockAttestationResponse_Status)(nil),
		(*BlockAttestationResponse_BlockAttestation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderer_blockattestation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderer_blockattestation_proto_goTypes,
		DependencyIndexes: file_orderer_blockattestation_proto_depIdxs,
		MessageInfos:      file_orderer_blockattestation_proto_msgTypes,
	}.Build()
	File_orderer_blockattestation_proto = out.File
	file_orderer_blockattestation_proto_rawDesc = nil
	file_orderer_blockattestation_proto_goTypes = nil
	file_orderer_blockattestation_proto_depIdxs = nil
}
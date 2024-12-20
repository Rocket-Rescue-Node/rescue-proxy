// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: api.proto

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

type RocketPoolNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RocketPoolNodesRequest) Reset() {
	*x = RocketPoolNodesRequest{}
	mi := &file_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RocketPoolNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RocketPoolNodesRequest) ProtoMessage() {}

func (x *RocketPoolNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RocketPoolNodesRequest.ProtoReflect.Descriptor instead.
func (*RocketPoolNodesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type RocketPoolNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIds [][]byte `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
}

func (x *RocketPoolNodes) Reset() {
	*x = RocketPoolNodes{}
	mi := &file_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RocketPoolNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RocketPoolNodes) ProtoMessage() {}

func (x *RocketPoolNodes) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RocketPoolNodes.ProtoReflect.Descriptor instead.
func (*RocketPoolNodes) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RocketPoolNodes) GetNodeIds() [][]byte {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

type OdaoNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OdaoNodesRequest) Reset() {
	*x = OdaoNodesRequest{}
	mi := &file_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OdaoNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OdaoNodesRequest) ProtoMessage() {}

func (x *OdaoNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OdaoNodesRequest.ProtoReflect.Descriptor instead.
func (*OdaoNodesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type OdaoNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIds [][]byte `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
}

func (x *OdaoNodes) Reset() {
	*x = OdaoNodes{}
	mi := &file_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OdaoNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OdaoNodes) ProtoMessage() {}

func (x *OdaoNodes) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OdaoNodes.ProtoReflect.Descriptor instead.
func (*OdaoNodes) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *OdaoNodes) GetNodeIds() [][]byte {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

type SoloValidatorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SoloValidatorsRequest) Reset() {
	*x = SoloValidatorsRequest{}
	mi := &file_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SoloValidatorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloValidatorsRequest) ProtoMessage() {}

func (x *SoloValidatorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloValidatorsRequest.ProtoReflect.Descriptor instead.
func (*SoloValidatorsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

type SoloValidators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithdrawalAddresses [][]byte `protobuf:"bytes,1,rep,name=withdrawal_addresses,json=withdrawalAddresses,proto3" json:"withdrawal_addresses,omitempty"`
}

func (x *SoloValidators) Reset() {
	*x = SoloValidators{}
	mi := &file_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SoloValidators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloValidators) ProtoMessage() {}

func (x *SoloValidators) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloValidators.ProtoReflect.Descriptor instead.
func (*SoloValidators) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *SoloValidators) GetWithdrawalAddresses() [][]byte {
	if x != nil {
		return x.WithdrawalAddresses
	}
	return nil
}

type ValidateEIP1271Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataHash  []byte `protobuf:"bytes,1,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Address   []byte `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ValidateEIP1271Request) Reset() {
	*x = ValidateEIP1271Request{}
	mi := &file_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateEIP1271Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateEIP1271Request) ProtoMessage() {}

func (x *ValidateEIP1271Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateEIP1271Request.ProtoReflect.Descriptor instead.
func (*ValidateEIP1271Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateEIP1271Request) GetDataHash() []byte {
	if x != nil {
		return x.DataHash
	}
	return nil
}

func (x *ValidateEIP1271Request) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ValidateEIP1271Request) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type ValidateEIP1271Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ValidateEIP1271Response) Reset() {
	*x = ValidateEIP1271Response{}
	mi := &file_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateEIP1271Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateEIP1271Response) ProtoMessage() {}

func (x *ValidateEIP1271Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateEIP1271Response.ProtoReflect.Descriptor instead.
func (*ValidateEIP1271Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateEIP1271Response) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateEIP1271Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x18, 0x0a, 0x16, 0x52, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0f, 0x52, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x4f, 0x64, 0x61, 0x6f, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x09, 0x4f,
	0x64, 0x61, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x6f, 0x6c, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x0e,
	0x53, 0x6f, 0x6c, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x31,
	0x0a, 0x14, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x13, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x22, 0x6d, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x49, 0x50,
	0x31, 0x32, 0x37, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x45, 0x0a, 0x17, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x49, 0x50, 0x31,
	0x32, 0x37, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x99, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12,
	0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f,
	0x64, 0x61, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x64,
	0x61, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x64, 0x61, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x6c, 0x6f, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x6c, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x49, 0x50, 0x31, 0x32, 0x37, 0x31, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x49, 0x50, 0x31, 0x32, 0x37, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x49, 0x50, 0x31, 0x32, 0x37, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []any{
	(*RocketPoolNodesRequest)(nil),  // 0: pb.RocketPoolNodesRequest
	(*RocketPoolNodes)(nil),         // 1: pb.RocketPoolNodes
	(*OdaoNodesRequest)(nil),        // 2: pb.OdaoNodesRequest
	(*OdaoNodes)(nil),               // 3: pb.OdaoNodes
	(*SoloValidatorsRequest)(nil),   // 4: pb.SoloValidatorsRequest
	(*SoloValidators)(nil),          // 5: pb.SoloValidators
	(*ValidateEIP1271Request)(nil),  // 6: pb.ValidateEIP1271Request
	(*ValidateEIP1271Response)(nil), // 7: pb.ValidateEIP1271Response
}
var file_api_proto_depIdxs = []int32{
	0, // 0: pb.Api.GetRocketPoolNodes:input_type -> pb.RocketPoolNodesRequest
	2, // 1: pb.Api.GetOdaoNodes:input_type -> pb.OdaoNodesRequest
	4, // 2: pb.Api.GetSoloValidators:input_type -> pb.SoloValidatorsRequest
	6, // 3: pb.Api.ValidateEIP1271:input_type -> pb.ValidateEIP1271Request
	1, // 4: pb.Api.GetRocketPoolNodes:output_type -> pb.RocketPoolNodes
	3, // 5: pb.Api.GetOdaoNodes:output_type -> pb.OdaoNodes
	5, // 6: pb.Api.GetSoloValidators:output_type -> pb.SoloValidators
	7, // 7: pb.Api.ValidateEIP1271:output_type -> pb.ValidateEIP1271Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

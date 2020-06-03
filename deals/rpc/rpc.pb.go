// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: deals/rpc/rpc.proto

package rpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DealConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Miner      string `protobuf:"bytes,1,opt,name=miner,proto3" json:"miner,omitempty"`
	EpochPrice uint64 `protobuf:"varint,2,opt,name=epoch_price,json=epochPrice,proto3" json:"epoch_price,omitempty"`
}

func (x *DealConfig) Reset() {
	*x = DealConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealConfig) ProtoMessage() {}

func (x *DealConfig) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealConfig.ProtoReflect.Descriptor instead.
func (*DealConfig) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DealConfig) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *DealConfig) GetEpochPrice() uint64 {
	if x != nil {
		return x.EpochPrice
	}
	return 0
}

type DealInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalCid     string `protobuf:"bytes,1,opt,name=proposal_cid,json=proposalCid,proto3" json:"proposal_cid,omitempty"`
	StateId         uint64 `protobuf:"varint,2,opt,name=state_id,json=stateId,proto3" json:"state_id,omitempty"`
	StateName       string `protobuf:"bytes,3,opt,name=state_name,json=stateName,proto3" json:"state_name,omitempty"`
	Miner           string `protobuf:"bytes,4,opt,name=miner,proto3" json:"miner,omitempty"`
	PieceCid        []byte `protobuf:"bytes,5,opt,name=piece_cid,json=pieceCid,proto3" json:"piece_cid,omitempty"`
	Size            uint64 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	PricePerEpoch   uint64 `protobuf:"varint,7,opt,name=price_per_epoch,json=pricePerEpoch,proto3" json:"price_per_epoch,omitempty"`
	StartEpoch      uint64 `protobuf:"varint,8,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	Duration        uint64 `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	DealId          uint64 `protobuf:"varint,10,opt,name=deal_id,json=dealId,proto3" json:"deal_id,omitempty"`
	ActivationEpoch int64  `protobuf:"varint,11,opt,name=activation_epoch,json=activationEpoch,proto3" json:"activation_epoch,omitempty"`
	Msg             string `protobuf:"bytes,12,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DealInfo) Reset() {
	*x = DealInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealInfo) ProtoMessage() {}

func (x *DealInfo) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealInfo.ProtoReflect.Descriptor instead.
func (*DealInfo) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DealInfo) GetProposalCid() string {
	if x != nil {
		return x.ProposalCid
	}
	return ""
}

func (x *DealInfo) GetStateId() uint64 {
	if x != nil {
		return x.StateId
	}
	return 0
}

func (x *DealInfo) GetStateName() string {
	if x != nil {
		return x.StateName
	}
	return ""
}

func (x *DealInfo) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *DealInfo) GetPieceCid() []byte {
	if x != nil {
		return x.PieceCid
	}
	return nil
}

func (x *DealInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DealInfo) GetPricePerEpoch() uint64 {
	if x != nil {
		return x.PricePerEpoch
	}
	return 0
}

func (x *DealInfo) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *DealInfo) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *DealInfo) GetDealId() uint64 {
	if x != nil {
		return x.DealId
	}
	return 0
}

func (x *DealInfo) GetActivationEpoch() int64 {
	if x != nil {
		return x.ActivationEpoch
	}
	return 0
}

func (x *DealInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type StoreParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DealConfigs []*DealConfig `protobuf:"bytes,2,rep,name=deal_configs,json=dealConfigs,proto3" json:"deal_configs,omitempty"`
	MinDuration uint64        `protobuf:"varint,3,opt,name=min_duration,json=minDuration,proto3" json:"min_duration,omitempty"`
}

func (x *StoreParams) Reset() {
	*x = StoreParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreParams) ProtoMessage() {}

func (x *StoreParams) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreParams.ProtoReflect.Descriptor instead.
func (*StoreParams) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *StoreParams) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StoreParams) GetDealConfigs() []*DealConfig {
	if x != nil {
		return x.DealConfigs
	}
	return nil
}

func (x *StoreParams) GetMinDuration() uint64 {
	if x != nil {
		return x.MinDuration
	}
	return 0
}

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*StoreRequest_StoreParams
	//	*StoreRequest_Chunk
	Payload isStoreRequest_Payload `protobuf_oneof:"payload"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{3}
}

func (m *StoreRequest) GetPayload() isStoreRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *StoreRequest) GetStoreParams() *StoreParams {
	if x, ok := x.GetPayload().(*StoreRequest_StoreParams); ok {
		return x.StoreParams
	}
	return nil
}

func (x *StoreRequest) GetChunk() []byte {
	if x, ok := x.GetPayload().(*StoreRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isStoreRequest_Payload interface {
	isStoreRequest_Payload()
}

type StoreRequest_StoreParams struct {
	StoreParams *StoreParams `protobuf:"bytes,1,opt,name=store_params,json=storeParams,proto3,oneof"`
}

type StoreRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*StoreRequest_StoreParams) isStoreRequest_Payload() {}

func (*StoreRequest_Chunk) isStoreRequest_Payload() {}

type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataCid      string        `protobuf:"bytes,1,opt,name=data_cid,json=dataCid,proto3" json:"data_cid,omitempty"`
	ProposalCids []string      `protobuf:"bytes,2,rep,name=proposal_cids,json=proposalCids,proto3" json:"proposal_cids,omitempty"`
	FailedDeals  []*DealConfig `protobuf:"bytes,3,rep,name=failed_deals,json=failedDeals,proto3" json:"failed_deals,omitempty"`
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *StoreResponse) GetDataCid() string {
	if x != nil {
		return x.DataCid
	}
	return ""
}

func (x *StoreResponse) GetProposalCids() []string {
	if x != nil {
		return x.ProposalCids
	}
	return nil
}

func (x *StoreResponse) GetFailedDeals() []*DealConfig {
	if x != nil {
		return x.FailedDeals
	}
	return nil
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposals []string `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *WatchRequest) GetProposals() []string {
	if x != nil {
		return x.Proposals
	}
	return nil
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealInfo *DealInfo `protobuf:"bytes,1,opt,name=deal_info,json=dealInfo,proto3" json:"deal_info,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *WatchResponse) GetDealInfo() *DealInfo {
	if x != nil {
		return x.DealInfo
	}
	return nil
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Cid     string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *RetrieveRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RetrieveRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type RetrieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *RetrieveResponse) Reset() {
	*x = RetrieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deals_rpc_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveResponse) ProtoMessage() {}

func (x *RetrieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deals_rpc_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveResponse.ProtoReflect.Descriptor instead.
func (*RetrieveResponse) Descriptor() ([]byte, []int) {
	return file_deals_rpc_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *RetrieveResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_deals_rpc_rpc_proto protoreflect.FileDescriptor

var file_deals_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70, 0x63,
	0x22, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xe9, 0x02, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x63,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x43, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x69, 0x65, 0x63, 0x65, 0x5f, 0x63,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x69, 0x65, 0x63, 0x65, 0x43,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x65, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x64, 0x65,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x84, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x64,
	0x65, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x64, 0x65, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x69, 0x6e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x43, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x5f, 0x63, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x43, 0x69, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44,
	0x65, 0x61, 0x6c, 0x73, 0x22, 0x2c, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x65, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3d, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xd5,
	0x01, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a,
	0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a,
	0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x1a, 0x2e, 0x64, 0x65, 0x61, 0x6c,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deals_rpc_rpc_proto_rawDescOnce sync.Once
	file_deals_rpc_rpc_proto_rawDescData = file_deals_rpc_rpc_proto_rawDesc
)

func file_deals_rpc_rpc_proto_rawDescGZIP() []byte {
	file_deals_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_deals_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_deals_rpc_rpc_proto_rawDescData)
	})
	return file_deals_rpc_rpc_proto_rawDescData
}

var file_deals_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_deals_rpc_rpc_proto_goTypes = []interface{}{
	(*DealConfig)(nil),       // 0: deals.rpc.DealConfig
	(*DealInfo)(nil),         // 1: deals.rpc.DealInfo
	(*StoreParams)(nil),      // 2: deals.rpc.StoreParams
	(*StoreRequest)(nil),     // 3: deals.rpc.StoreRequest
	(*StoreResponse)(nil),    // 4: deals.rpc.StoreResponse
	(*WatchRequest)(nil),     // 5: deals.rpc.WatchRequest
	(*WatchResponse)(nil),    // 6: deals.rpc.WatchResponse
	(*RetrieveRequest)(nil),  // 7: deals.rpc.RetrieveRequest
	(*RetrieveResponse)(nil), // 8: deals.rpc.RetrieveResponse
}
var file_deals_rpc_rpc_proto_depIdxs = []int32{
	0, // 0: deals.rpc.StoreParams.deal_configs:type_name -> deals.rpc.DealConfig
	2, // 1: deals.rpc.StoreRequest.store_params:type_name -> deals.rpc.StoreParams
	0, // 2: deals.rpc.StoreResponse.failed_deals:type_name -> deals.rpc.DealConfig
	1, // 3: deals.rpc.WatchResponse.deal_info:type_name -> deals.rpc.DealInfo
	3, // 4: deals.rpc.RPCService.Store:input_type -> deals.rpc.StoreRequest
	5, // 5: deals.rpc.RPCService.Watch:input_type -> deals.rpc.WatchRequest
	7, // 6: deals.rpc.RPCService.Retrieve:input_type -> deals.rpc.RetrieveRequest
	4, // 7: deals.rpc.RPCService.Store:output_type -> deals.rpc.StoreResponse
	6, // 8: deals.rpc.RPCService.Watch:output_type -> deals.rpc.WatchResponse
	8, // 9: deals.rpc.RPCService.Retrieve:output_type -> deals.rpc.RetrieveResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_deals_rpc_rpc_proto_init() }
func file_deals_rpc_rpc_proto_init() {
	if File_deals_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deals_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealConfig); i {
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
		file_deals_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealInfo); i {
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
		file_deals_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreParams); i {
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
		file_deals_rpc_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_deals_rpc_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
		file_deals_rpc_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_deals_rpc_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
		file_deals_rpc_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_deals_rpc_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveResponse); i {
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
	file_deals_rpc_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StoreRequest_StoreParams)(nil),
		(*StoreRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deals_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deals_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_deals_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_deals_rpc_rpc_proto_msgTypes,
	}.Build()
	File_deals_rpc_rpc_proto = out.File
	file_deals_rpc_rpc_proto_rawDesc = nil
	file_deals_rpc_rpc_proto_goTypes = nil
	file_deals_rpc_rpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCServiceClient interface {
	Store(ctx context.Context, opts ...grpc.CallOption) (RPCService_StoreClient, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (RPCService_WatchClient, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (RPCService_RetrieveClient, error)
}

type rPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCServiceClient(cc grpc.ClientConnInterface) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) Store(ctx context.Context, opts ...grpc.CallOption) (RPCService_StoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RPCService_serviceDesc.Streams[0], "/deals.rpc.RPCService/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCServiceStoreClient{stream}
	return x, nil
}

type RPCService_StoreClient interface {
	Send(*StoreRequest) error
	CloseAndRecv() (*StoreResponse, error)
	grpc.ClientStream
}

type rPCServiceStoreClient struct {
	grpc.ClientStream
}

func (x *rPCServiceStoreClient) Send(m *StoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCServiceStoreClient) CloseAndRecv() (*StoreResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rPCServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (RPCService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RPCService_serviceDesc.Streams[1], "/deals.rpc.RPCService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type rPCServiceWatchClient struct {
	grpc.ClientStream
}

func (x *rPCServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rPCServiceClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (RPCService_RetrieveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RPCService_serviceDesc.Streams[2], "/deals.rpc.RPCService/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCServiceRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_RetrieveClient interface {
	Recv() (*RetrieveResponse, error)
	grpc.ClientStream
}

type rPCServiceRetrieveClient struct {
	grpc.ClientStream
}

func (x *rPCServiceRetrieveClient) Recv() (*RetrieveResponse, error) {
	m := new(RetrieveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RPCServiceServer is the server API for RPCService service.
type RPCServiceServer interface {
	Store(RPCService_StoreServer) error
	Watch(*WatchRequest, RPCService_WatchServer) error
	Retrieve(*RetrieveRequest, RPCService_RetrieveServer) error
}

// UnimplementedRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (*UnimplementedRPCServiceServer) Store(RPCService_StoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedRPCServiceServer) Watch(*WatchRequest, RPCService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (*UnimplementedRPCServiceServer) Retrieve(*RetrieveRequest, RPCService_RetrieveServer) error {
	return status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RPCServiceServer).Store(&rPCServiceStoreServer{stream})
}

type RPCService_StoreServer interface {
	SendAndClose(*StoreResponse) error
	Recv() (*StoreRequest, error)
	grpc.ServerStream
}

type rPCServiceStoreServer struct {
	grpc.ServerStream
}

func (x *rPCServiceStoreServer) SendAndClose(m *StoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rPCServiceStoreServer) Recv() (*StoreRequest, error) {
	m := new(StoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RPCService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RPCServiceServer).Watch(m, &rPCServiceWatchServer{stream})
}

type RPCService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type rPCServiceWatchServer struct {
	grpc.ServerStream
}

func (x *rPCServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RPCService_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RPCServiceServer).Retrieve(m, &rPCServiceRetrieveServer{stream})
}

type RPCService_RetrieveServer interface {
	Send(*RetrieveResponse) error
	grpc.ServerStream
}

type rPCServiceRetrieveServer struct {
	grpc.ServerStream
}

func (x *rPCServiceRetrieveServer) Send(m *RetrieveResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deals.rpc.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Store",
			Handler:       _RPCService_Store_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _RPCService_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Retrieve",
			Handler:       _RPCService_Retrieve_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "deals/rpc/rpc.proto",
}

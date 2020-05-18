// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DealConfig struct {
	Miner                string   `protobuf:"bytes,1,opt,name=miner,proto3" json:"miner,omitempty"`
	EpochPrice           uint64   `protobuf:"varint,2,opt,name=epochPrice,proto3" json:"epochPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealConfig) Reset()         { *m = DealConfig{} }
func (m *DealConfig) String() string { return proto.CompactTextString(m) }
func (*DealConfig) ProtoMessage()    {}
func (*DealConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *DealConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealConfig.Unmarshal(m, b)
}
func (m *DealConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealConfig.Marshal(b, m, deterministic)
}
func (m *DealConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealConfig.Merge(m, src)
}
func (m *DealConfig) XXX_Size() int {
	return xxx_messageInfo_DealConfig.Size(m)
}
func (m *DealConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DealConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DealConfig proto.InternalMessageInfo

func (m *DealConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *DealConfig) GetEpochPrice() uint64 {
	if m != nil {
		return m.EpochPrice
	}
	return 0
}

type DealInfo struct {
	ProposalCid          string   `protobuf:"bytes,1,opt,name=proposalCid,proto3" json:"proposalCid,omitempty"`
	StateID              uint64   `protobuf:"varint,2,opt,name=stateID,proto3" json:"stateID,omitempty"`
	StateName            string   `protobuf:"bytes,3,opt,name=stateName,proto3" json:"stateName,omitempty"`
	Miner                string   `protobuf:"bytes,4,opt,name=miner,proto3" json:"miner,omitempty"`
	PieceCID             []byte   `protobuf:"bytes,5,opt,name=pieceCID,proto3" json:"pieceCID,omitempty"`
	Size                 uint64   `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	PricePerEpoch        uint64   `protobuf:"varint,7,opt,name=pricePerEpoch,proto3" json:"pricePerEpoch,omitempty"`
	Duration             uint64   `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealInfo) Reset()         { *m = DealInfo{} }
func (m *DealInfo) String() string { return proto.CompactTextString(m) }
func (*DealInfo) ProtoMessage()    {}
func (*DealInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *DealInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealInfo.Unmarshal(m, b)
}
func (m *DealInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealInfo.Marshal(b, m, deterministic)
}
func (m *DealInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealInfo.Merge(m, src)
}
func (m *DealInfo) XXX_Size() int {
	return xxx_messageInfo_DealInfo.Size(m)
}
func (m *DealInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DealInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DealInfo proto.InternalMessageInfo

func (m *DealInfo) GetProposalCid() string {
	if m != nil {
		return m.ProposalCid
	}
	return ""
}

func (m *DealInfo) GetStateID() uint64 {
	if m != nil {
		return m.StateID
	}
	return 0
}

func (m *DealInfo) GetStateName() string {
	if m != nil {
		return m.StateName
	}
	return ""
}

func (m *DealInfo) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *DealInfo) GetPieceCID() []byte {
	if m != nil {
		return m.PieceCID
	}
	return nil
}

func (m *DealInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *DealInfo) GetPricePerEpoch() uint64 {
	if m != nil {
		return m.PricePerEpoch
	}
	return 0
}

func (m *DealInfo) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type StoreParams struct {
	Address              string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DealConfigs          []*DealConfig `protobuf:"bytes,2,rep,name=dealConfigs,proto3" json:"dealConfigs,omitempty"`
	Duration             uint64        `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreParams) Reset()         { *m = StoreParams{} }
func (m *StoreParams) String() string { return proto.CompactTextString(m) }
func (*StoreParams) ProtoMessage()    {}
func (*StoreParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *StoreParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreParams.Unmarshal(m, b)
}
func (m *StoreParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreParams.Marshal(b, m, deterministic)
}
func (m *StoreParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreParams.Merge(m, src)
}
func (m *StoreParams) XXX_Size() int {
	return xxx_messageInfo_StoreParams.Size(m)
}
func (m *StoreParams) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreParams.DiscardUnknown(m)
}

var xxx_messageInfo_StoreParams proto.InternalMessageInfo

func (m *StoreParams) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StoreParams) GetDealConfigs() []*DealConfig {
	if m != nil {
		return m.DealConfigs
	}
	return nil
}

func (m *StoreParams) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type StoreRequest struct {
	// Types that are valid to be assigned to Payload:
	//	*StoreRequest_StoreParams
	//	*StoreRequest_Chunk
	Payload              isStoreRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

type isStoreRequest_Payload interface {
	isStoreRequest_Payload()
}

type StoreRequest_StoreParams struct {
	StoreParams *StoreParams `protobuf:"bytes,1,opt,name=storeParams,proto3,oneof"`
}

type StoreRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*StoreRequest_StoreParams) isStoreRequest_Payload() {}

func (*StoreRequest_Chunk) isStoreRequest_Payload() {}

func (m *StoreRequest) GetPayload() isStoreRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StoreRequest) GetStoreParams() *StoreParams {
	if x, ok := m.GetPayload().(*StoreRequest_StoreParams); ok {
		return x.StoreParams
	}
	return nil
}

func (m *StoreRequest) GetChunk() []byte {
	if x, ok := m.GetPayload().(*StoreRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StoreRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StoreRequest_StoreParams)(nil),
		(*StoreRequest_Chunk)(nil),
	}
}

type StoreReply struct {
	DataCid              string        `protobuf:"bytes,1,opt,name=dataCid,proto3" json:"dataCid,omitempty"`
	ProposalCids         []string      `protobuf:"bytes,2,rep,name=proposalCids,proto3" json:"proposalCids,omitempty"`
	FailedDeals          []*DealConfig `protobuf:"bytes,3,rep,name=failedDeals,proto3" json:"failedDeals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreReply) Reset()         { *m = StoreReply{} }
func (m *StoreReply) String() string { return proto.CompactTextString(m) }
func (*StoreReply) ProtoMessage()    {}
func (*StoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *StoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreReply.Unmarshal(m, b)
}
func (m *StoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreReply.Marshal(b, m, deterministic)
}
func (m *StoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreReply.Merge(m, src)
}
func (m *StoreReply) XXX_Size() int {
	return xxx_messageInfo_StoreReply.Size(m)
}
func (m *StoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_StoreReply proto.InternalMessageInfo

func (m *StoreReply) GetDataCid() string {
	if m != nil {
		return m.DataCid
	}
	return ""
}

func (m *StoreReply) GetProposalCids() []string {
	if m != nil {
		return m.ProposalCids
	}
	return nil
}

func (m *StoreReply) GetFailedDeals() []*DealConfig {
	if m != nil {
		return m.FailedDeals
	}
	return nil
}

type WatchRequest struct {
	Proposals            []string `protobuf:"bytes,1,rep,name=proposals,proto3" json:"proposals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetProposals() []string {
	if m != nil {
		return m.Proposals
	}
	return nil
}

type WatchReply struct {
	DealInfo             *DealInfo `protobuf:"bytes,1,opt,name=dealInfo,proto3" json:"dealInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WatchReply) Reset()         { *m = WatchReply{} }
func (m *WatchReply) String() string { return proto.CompactTextString(m) }
func (*WatchReply) ProtoMessage()    {}
func (*WatchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *WatchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchReply.Unmarshal(m, b)
}
func (m *WatchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchReply.Marshal(b, m, deterministic)
}
func (m *WatchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchReply.Merge(m, src)
}
func (m *WatchReply) XXX_Size() int {
	return xxx_messageInfo_WatchReply.Size(m)
}
func (m *WatchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchReply.DiscardUnknown(m)
}

var xxx_messageInfo_WatchReply proto.InternalMessageInfo

func (m *WatchReply) GetDealInfo() *DealInfo {
	if m != nil {
		return m.DealInfo
	}
	return nil
}

type RetrieveRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Cid                  string   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveRequest) Reset()         { *m = RetrieveRequest{} }
func (m *RetrieveRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveRequest) ProtoMessage()    {}
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *RetrieveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveRequest.Unmarshal(m, b)
}
func (m *RetrieveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveRequest.Marshal(b, m, deterministic)
}
func (m *RetrieveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveRequest.Merge(m, src)
}
func (m *RetrieveRequest) XXX_Size() int {
	return xxx_messageInfo_RetrieveRequest.Size(m)
}
func (m *RetrieveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveRequest proto.InternalMessageInfo

func (m *RetrieveRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RetrieveRequest) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type RetrieveReply struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveReply) Reset()         { *m = RetrieveReply{} }
func (m *RetrieveReply) String() string { return proto.CompactTextString(m) }
func (*RetrieveReply) ProtoMessage()    {}
func (*RetrieveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *RetrieveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveReply.Unmarshal(m, b)
}
func (m *RetrieveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveReply.Marshal(b, m, deterministic)
}
func (m *RetrieveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveReply.Merge(m, src)
}
func (m *RetrieveReply) XXX_Size() int {
	return xxx_messageInfo_RetrieveReply.Size(m)
}
func (m *RetrieveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveReply.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveReply proto.InternalMessageInfo

func (m *RetrieveReply) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*DealConfig)(nil), "deals.rpc.DealConfig")
	proto.RegisterType((*DealInfo)(nil), "deals.rpc.DealInfo")
	proto.RegisterType((*StoreParams)(nil), "deals.rpc.StoreParams")
	proto.RegisterType((*StoreRequest)(nil), "deals.rpc.StoreRequest")
	proto.RegisterType((*StoreReply)(nil), "deals.rpc.StoreReply")
	proto.RegisterType((*WatchRequest)(nil), "deals.rpc.WatchRequest")
	proto.RegisterType((*WatchReply)(nil), "deals.rpc.WatchReply")
	proto.RegisterType((*RetrieveRequest)(nil), "deals.rpc.RetrieveRequest")
	proto.RegisterType((*RetrieveReply)(nil), "deals.rpc.RetrieveReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x6f, 0x9a, 0x76, 0xdb, 0x7c, 0x69, 0x51, 0xc6, 0xdd, 0x35, 0x94, 0x22, 0x25, 0x28, 0xf4,
	0x20, 0x51, 0xea, 0x41, 0x58, 0xd9, 0x83, 0x6d, 0x57, 0xb6, 0x17, 0x29, 0x71, 0x41, 0xf0, 0x36,
	0x9b, 0x4c, 0xed, 0x60, 0x9a, 0x19, 0x27, 0x53, 0xb1, 0xe2, 0xcd, 0x37, 0xf1, 0x5d, 0x7c, 0x1b,
	0x1f, 0x42, 0x66, 0x3a, 0x49, 0xa6, 0x55, 0xf7, 0x36, 0xdf, 0xdf, 0xdf, 0x9f, 0x2f, 0x04, 0x3c,
	0xcc, 0x69, 0xc4, 0x05, 0x93, 0x0c, 0x79, 0x29, 0xc1, 0x59, 0x11, 0x09, 0x9e, 0x84, 0x53, 0x80,
	0x39, 0xc1, 0xd9, 0x8c, 0xe5, 0x2b, 0xfa, 0x11, 0x9d, 0x42, 0x7b, 0x43, 0x73, 0x22, 0x02, 0x67,
	0xe4, 0x8c, 0xbd, 0x78, 0x1f, 0xa0, 0x47, 0x00, 0x84, 0xb3, 0x64, 0xbd, 0x14, 0x34, 0x21, 0x41,
	0x73, 0xe4, 0x8c, 0x5b, 0xb1, 0x95, 0x09, 0x7f, 0x3b, 0xd0, 0x55, 0x4b, 0x16, 0xf9, 0x8a, 0xa1,
	0x11, 0xf8, 0x5c, 0x30, 0xce, 0x0a, 0x9c, 0xcd, 0x68, 0x6a, 0x16, 0xd9, 0x29, 0x14, 0x40, 0xa7,
	0x90, 0x58, 0x92, 0xc5, 0xdc, 0xec, 0x2a, 0x43, 0x34, 0x04, 0x4f, 0x3f, 0xdf, 0xe2, 0x0d, 0x09,
	0x5c, 0x3d, 0x59, 0x27, 0x6a, 0x72, 0x2d, 0x9b, 0xdc, 0x00, 0xba, 0x9c, 0x92, 0x84, 0xcc, 0x16,
	0xf3, 0xa0, 0x3d, 0x72, 0xc6, 0xbd, 0xb8, 0x8a, 0x11, 0x82, 0x56, 0x41, 0xbf, 0x91, 0xe0, 0x44,
	0xc3, 0xe8, 0x37, 0x7a, 0x0c, 0x7d, 0xae, 0x58, 0x2f, 0x89, 0xb8, 0x52, 0x12, 0x82, 0x8e, 0x2e,
	0x1e, 0x26, 0xd5, 0xd6, 0x74, 0x2b, 0xb0, 0xa4, 0x2c, 0x0f, 0xba, 0xba, 0xa1, 0x8a, 0xc3, 0xef,
	0xe0, 0xbf, 0x93, 0x4c, 0x90, 0x25, 0x16, 0x78, 0x53, 0x28, 0x39, 0x38, 0x4d, 0x05, 0x29, 0x0a,
	0x23, 0xb6, 0x0c, 0xd1, 0x4b, 0xf0, 0xd3, 0xca, 0xdb, 0x22, 0x68, 0x8e, 0xdc, 0xb1, 0x3f, 0x39,
	0x8b, 0x2a, 0xf3, 0xa3, 0xda, 0xf9, 0xd8, 0xee, 0x3c, 0x40, 0x77, 0x8f, 0xd0, 0x37, 0xd0, 0xd3,
	0xe8, 0x31, 0xf9, 0xbc, 0x25, 0x85, 0x44, 0x17, 0xe0, 0x17, 0x35, 0x1b, 0x4d, 0xc1, 0x9f, 0x9c,
	0x5b, 0x20, 0x16, 0xd7, 0xeb, 0x46, 0x6c, 0x37, 0xa3, 0x73, 0x68, 0x27, 0xeb, 0x6d, 0xfe, 0x49,
	0xdf, 0xa1, 0x77, 0xdd, 0x88, 0xf7, 0xe1, 0xd4, 0x83, 0x0e, 0xc7, 0xbb, 0x8c, 0xe1, 0x34, 0xfc,
	0xe1, 0x00, 0x18, 0x3c, 0x9e, 0xed, 0x94, 0xd8, 0x14, 0x4b, 0x5c, 0x5f, 0xb6, 0x0c, 0x51, 0x08,
	0x3d, 0xeb, 0xc8, 0x7b, 0xb5, 0x5e, 0x7c, 0x90, 0x53, 0x86, 0xac, 0x30, 0xcd, 0x48, 0xaa, 0x84,
	0x17, 0x81, 0x7b, 0xa7, 0x21, 0x56, 0x67, 0xf8, 0x14, 0x7a, 0xef, 0xb1, 0x4c, 0xd6, 0xa5, 0xe8,
	0x21, 0x78, 0xe5, 0x62, 0x25, 0x59, 0x21, 0xd5, 0x89, 0xf0, 0x12, 0xc0, 0x74, 0x2b, 0xca, 0xcf,
	0xa0, 0x9b, 0x9a, 0x8f, 0xd3, 0xb8, 0xf3, 0xe0, 0x08, 0x51, 0x95, 0xe2, 0xaa, 0x29, 0xbc, 0x84,
	0x7b, 0x31, 0x91, 0x82, 0x92, 0x2f, 0x95, 0xc9, 0xff, 0xbf, 0xf1, 0x7d, 0x70, 0x13, 0x9a, 0x6a,
	0x03, 0xbd, 0x58, 0x3d, 0xc3, 0x27, 0xd0, 0xaf, 0xc7, 0x15, 0x81, 0xd3, 0xd2, 0x65, 0x47, 0x7f,
	0x9e, 0xfb, 0x60, 0xf2, 0xcb, 0x01, 0xf7, 0xf5, 0x72, 0x81, 0x5e, 0x41, 0x5b, 0xfb, 0x8b, 0x1e,
	0x1e, 0xdf, 0xcc, 0x80, 0x0f, 0xce, 0xfe, 0x2e, 0xf0, 0x6c, 0x17, 0x36, 0xc6, 0x8e, 0x1a, 0xd6,
	0x4a, 0x0f, 0x86, 0x6d, 0xa7, 0x0e, 0x86, 0x6b, 0x53, 0xc2, 0xc6, 0x73, 0x07, 0xcd, 0xa1, 0x5b,
	0x12, 0x45, 0x03, 0xab, 0xed, 0x48, 0xfc, 0x20, 0xf8, 0x67, 0xcd, 0x6c, 0x99, 0x5e, 0xc0, 0x90,
	0xb2, 0x48, 0x92, 0xaf, 0x92, 0x66, 0x24, 0x5a, 0xd1, 0x8c, 0x24, 0x8c, 0xe6, 0x66, 0x84, 0xdf,
	0x4e, 0xfb, 0x6f, 0x4c, 0x4a, 0x5f, 0x72, 0xe9, 0x7c, 0x70, 0x05, 0x4f, 0x7e, 0x36, 0xdd, 0x9b,
	0x9b, 0xab, 0xdb, 0x13, 0xfd, 0x3b, 0x7a, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xa6, 0x11,
	0x86, 0x9b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Store(ctx context.Context, opts ...grpc.CallOption) (API_StoreClient, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (API_WatchClient, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (API_RetrieveClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Store(ctx context.Context, opts ...grpc.CallOption) (API_StoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/deals.rpc.API/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIStoreClient{stream}
	return x, nil
}

type API_StoreClient interface {
	Send(*StoreRequest) error
	CloseAndRecv() (*StoreReply, error)
	grpc.ClientStream
}

type aPIStoreClient struct {
	grpc.ClientStream
}

func (x *aPIStoreClient) Send(m *StoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIStoreClient) CloseAndRecv() (*StoreReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (API_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[1], "/deals.rpc.API/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_WatchClient interface {
	Recv() (*WatchReply, error)
	grpc.ClientStream
}

type aPIWatchClient struct {
	grpc.ClientStream
}

func (x *aPIWatchClient) Recv() (*WatchReply, error) {
	m := new(WatchReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (API_RetrieveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[2], "/deals.rpc.API/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_RetrieveClient interface {
	Recv() (*RetrieveReply, error)
	grpc.ClientStream
}

type aPIRetrieveClient struct {
	grpc.ClientStream
}

func (x *aPIRetrieveClient) Recv() (*RetrieveReply, error) {
	m := new(RetrieveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Store(API_StoreServer) error
	Watch(*WatchRequest, API_WatchServer) error
	Retrieve(*RetrieveRequest, API_RetrieveServer) error
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) Store(srv API_StoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedAPIServer) Watch(req *WatchRequest, srv API_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (*UnimplementedAPIServer) Retrieve(req *RetrieveRequest, srv API_RetrieveServer) error {
	return status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).Store(&aPIStoreServer{stream})
}

type API_StoreServer interface {
	SendAndClose(*StoreReply) error
	Recv() (*StoreRequest, error)
	grpc.ServerStream
}

type aPIStoreServer struct {
	grpc.ServerStream
}

func (x *aPIStoreServer) SendAndClose(m *StoreReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIStoreServer) Recv() (*StoreRequest, error) {
	m := new(StoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _API_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Watch(m, &aPIWatchServer{stream})
}

type API_WatchServer interface {
	Send(*WatchReply) error
	grpc.ServerStream
}

type aPIWatchServer struct {
	grpc.ServerStream
}

func (x *aPIWatchServer) Send(m *WatchReply) error {
	return x.ServerStream.SendMsg(m)
}

func _API_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Retrieve(m, &aPIRetrieveServer{stream})
}

type API_RetrieveServer interface {
	Send(*RetrieveReply) error
	grpc.ServerStream
}

type aPIRetrieveServer struct {
	grpc.ServerStream
}

func (x *aPIRetrieveServer) Send(m *RetrieveReply) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deals.rpc.API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Store",
			Handler:       _API_Store_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _API_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Retrieve",
			Handler:       _API_Retrieve_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

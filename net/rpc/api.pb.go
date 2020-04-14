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

type Connectedness int32

const (
	Connectedness_NotConnected  Connectedness = 0
	Connectedness_Connected     Connectedness = 1
	Connectedness_CanConnect    Connectedness = 2
	Connectedness_CannotConnect Connectedness = 3
	Connectedness_Unknown       Connectedness = 4
	Connectedness_Error         Connectedness = 5
)

var Connectedness_name = map[int32]string{
	0: "NotConnected",
	1: "Connected",
	2: "CanConnect",
	3: "CannotConnect",
	4: "Unknown",
	5: "Error",
}

var Connectedness_value = map[string]int32{
	"NotConnected":  0,
	"Connected":     1,
	"CanConnect":    2,
	"CannotConnect": 3,
	"Unknown":       4,
	"Error":         5,
}

func (x Connectedness) String() string {
	return proto.EnumName(Connectedness_name, int32(x))
}

func (Connectedness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type AddrInfo struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Addrs                []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddrInfo) Reset()         { *m = AddrInfo{} }
func (m *AddrInfo) String() string { return proto.CompactTextString(m) }
func (*AddrInfo) ProtoMessage()    {}
func (*AddrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *AddrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddrInfo.Unmarshal(m, b)
}
func (m *AddrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddrInfo.Marshal(b, m, deterministic)
}
func (m *AddrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrInfo.Merge(m, src)
}
func (m *AddrInfo) XXX_Size() int {
	return xxx_messageInfo_AddrInfo.Size(m)
}
func (m *AddrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddrInfo proto.InternalMessageInfo

func (m *AddrInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AddrInfo) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type ListenAddrRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenAddrRequest) Reset()         { *m = ListenAddrRequest{} }
func (m *ListenAddrRequest) String() string { return proto.CompactTextString(m) }
func (*ListenAddrRequest) ProtoMessage()    {}
func (*ListenAddrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ListenAddrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenAddrRequest.Unmarshal(m, b)
}
func (m *ListenAddrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenAddrRequest.Marshal(b, m, deterministic)
}
func (m *ListenAddrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenAddrRequest.Merge(m, src)
}
func (m *ListenAddrRequest) XXX_Size() int {
	return xxx_messageInfo_ListenAddrRequest.Size(m)
}
func (m *ListenAddrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenAddrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenAddrRequest proto.InternalMessageInfo

type ListenAddrReply struct {
	AddrInfo             *AddrInfo `protobuf:"bytes,1,opt,name=addrInfo,proto3" json:"addrInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListenAddrReply) Reset()         { *m = ListenAddrReply{} }
func (m *ListenAddrReply) String() string { return proto.CompactTextString(m) }
func (*ListenAddrReply) ProtoMessage()    {}
func (*ListenAddrReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *ListenAddrReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenAddrReply.Unmarshal(m, b)
}
func (m *ListenAddrReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenAddrReply.Marshal(b, m, deterministic)
}
func (m *ListenAddrReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenAddrReply.Merge(m, src)
}
func (m *ListenAddrReply) XXX_Size() int {
	return xxx_messageInfo_ListenAddrReply.Size(m)
}
func (m *ListenAddrReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenAddrReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListenAddrReply proto.InternalMessageInfo

func (m *ListenAddrReply) GetAddrInfo() *AddrInfo {
	if m != nil {
		return m.AddrInfo
	}
	return nil
}

type PeersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersRequest) Reset()         { *m = PeersRequest{} }
func (m *PeersRequest) String() string { return proto.CompactTextString(m) }
func (*PeersRequest) ProtoMessage()    {}
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *PeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersRequest.Unmarshal(m, b)
}
func (m *PeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersRequest.Marshal(b, m, deterministic)
}
func (m *PeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersRequest.Merge(m, src)
}
func (m *PeersRequest) XXX_Size() int {
	return xxx_messageInfo_PeersRequest.Size(m)
}
func (m *PeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeersRequest proto.InternalMessageInfo

type PeersReply struct {
	Peers                []*AddrInfo `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PeersReply) Reset()         { *m = PeersReply{} }
func (m *PeersReply) String() string { return proto.CompactTextString(m) }
func (*PeersReply) ProtoMessage()    {}
func (*PeersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *PeersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersReply.Unmarshal(m, b)
}
func (m *PeersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersReply.Marshal(b, m, deterministic)
}
func (m *PeersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersReply.Merge(m, src)
}
func (m *PeersReply) XXX_Size() int {
	return xxx_messageInfo_PeersReply.Size(m)
}
func (m *PeersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersReply.DiscardUnknown(m)
}

var xxx_messageInfo_PeersReply proto.InternalMessageInfo

func (m *PeersReply) GetPeers() []*AddrInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

type FindPeerRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPeerRequest) Reset()         { *m = FindPeerRequest{} }
func (m *FindPeerRequest) String() string { return proto.CompactTextString(m) }
func (*FindPeerRequest) ProtoMessage()    {}
func (*FindPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *FindPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPeerRequest.Unmarshal(m, b)
}
func (m *FindPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPeerRequest.Marshal(b, m, deterministic)
}
func (m *FindPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPeerRequest.Merge(m, src)
}
func (m *FindPeerRequest) XXX_Size() int {
	return xxx_messageInfo_FindPeerRequest.Size(m)
}
func (m *FindPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPeerRequest proto.InternalMessageInfo

func (m *FindPeerRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type FindPeerReply struct {
	PeerInfo             *AddrInfo `protobuf:"bytes,1,opt,name=peerInfo,proto3" json:"peerInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindPeerReply) Reset()         { *m = FindPeerReply{} }
func (m *FindPeerReply) String() string { return proto.CompactTextString(m) }
func (*FindPeerReply) ProtoMessage()    {}
func (*FindPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *FindPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPeerReply.Unmarshal(m, b)
}
func (m *FindPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPeerReply.Marshal(b, m, deterministic)
}
func (m *FindPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPeerReply.Merge(m, src)
}
func (m *FindPeerReply) XXX_Size() int {
	return xxx_messageInfo_FindPeerReply.Size(m)
}
func (m *FindPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindPeerReply proto.InternalMessageInfo

func (m *FindPeerReply) GetPeerInfo() *AddrInfo {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type ConnectPeerRequest struct {
	PeerInfo             *AddrInfo `protobuf:"bytes,1,opt,name=peerInfo,proto3" json:"peerInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ConnectPeerRequest) Reset()         { *m = ConnectPeerRequest{} }
func (m *ConnectPeerRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectPeerRequest) ProtoMessage()    {}
func (*ConnectPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *ConnectPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectPeerRequest.Unmarshal(m, b)
}
func (m *ConnectPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectPeerRequest.Marshal(b, m, deterministic)
}
func (m *ConnectPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectPeerRequest.Merge(m, src)
}
func (m *ConnectPeerRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectPeerRequest.Size(m)
}
func (m *ConnectPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectPeerRequest proto.InternalMessageInfo

func (m *ConnectPeerRequest) GetPeerInfo() *AddrInfo {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type ConnectPeerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectPeerReply) Reset()         { *m = ConnectPeerReply{} }
func (m *ConnectPeerReply) String() string { return proto.CompactTextString(m) }
func (*ConnectPeerReply) ProtoMessage()    {}
func (*ConnectPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *ConnectPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectPeerReply.Unmarshal(m, b)
}
func (m *ConnectPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectPeerReply.Marshal(b, m, deterministic)
}
func (m *ConnectPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectPeerReply.Merge(m, src)
}
func (m *ConnectPeerReply) XXX_Size() int {
	return xxx_messageInfo_ConnectPeerReply.Size(m)
}
func (m *ConnectPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectPeerReply proto.InternalMessageInfo

type DisconnectPeerRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectPeerRequest) Reset()         { *m = DisconnectPeerRequest{} }
func (m *DisconnectPeerRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectPeerRequest) ProtoMessage()    {}
func (*DisconnectPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *DisconnectPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectPeerRequest.Unmarshal(m, b)
}
func (m *DisconnectPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectPeerRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectPeerRequest.Merge(m, src)
}
func (m *DisconnectPeerRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectPeerRequest.Size(m)
}
func (m *DisconnectPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectPeerRequest proto.InternalMessageInfo

func (m *DisconnectPeerRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type DisconnectPeerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectPeerReply) Reset()         { *m = DisconnectPeerReply{} }
func (m *DisconnectPeerReply) String() string { return proto.CompactTextString(m) }
func (*DisconnectPeerReply) ProtoMessage()    {}
func (*DisconnectPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *DisconnectPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectPeerReply.Unmarshal(m, b)
}
func (m *DisconnectPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectPeerReply.Marshal(b, m, deterministic)
}
func (m *DisconnectPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectPeerReply.Merge(m, src)
}
func (m *DisconnectPeerReply) XXX_Size() int {
	return xxx_messageInfo_DisconnectPeerReply.Size(m)
}
func (m *DisconnectPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectPeerReply proto.InternalMessageInfo

type ConnectednessRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectednessRequest) Reset()         { *m = ConnectednessRequest{} }
func (m *ConnectednessRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectednessRequest) ProtoMessage()    {}
func (*ConnectednessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *ConnectednessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectednessRequest.Unmarshal(m, b)
}
func (m *ConnectednessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectednessRequest.Marshal(b, m, deterministic)
}
func (m *ConnectednessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectednessRequest.Merge(m, src)
}
func (m *ConnectednessRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectednessRequest.Size(m)
}
func (m *ConnectednessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectednessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectednessRequest proto.InternalMessageInfo

func (m *ConnectednessRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type ConnectednessReply struct {
	Connectedness        Connectedness `protobuf:"varint,1,opt,name=connectedness,proto3,enum=rpc.Connectedness" json:"connectedness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConnectednessReply) Reset()         { *m = ConnectednessReply{} }
func (m *ConnectednessReply) String() string { return proto.CompactTextString(m) }
func (*ConnectednessReply) ProtoMessage()    {}
func (*ConnectednessReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *ConnectednessReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectednessReply.Unmarshal(m, b)
}
func (m *ConnectednessReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectednessReply.Marshal(b, m, deterministic)
}
func (m *ConnectednessReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectednessReply.Merge(m, src)
}
func (m *ConnectednessReply) XXX_Size() int {
	return xxx_messageInfo_ConnectednessReply.Size(m)
}
func (m *ConnectednessReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectednessReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectednessReply proto.InternalMessageInfo

func (m *ConnectednessReply) GetConnectedness() Connectedness {
	if m != nil {
		return m.Connectedness
	}
	return Connectedness_NotConnected
}

func init() {
	proto.RegisterEnum("rpc.Connectedness", Connectedness_name, Connectedness_value)
	proto.RegisterType((*AddrInfo)(nil), "rpc.AddrInfo")
	proto.RegisterType((*ListenAddrRequest)(nil), "rpc.ListenAddrRequest")
	proto.RegisterType((*ListenAddrReply)(nil), "rpc.ListenAddrReply")
	proto.RegisterType((*PeersRequest)(nil), "rpc.PeersRequest")
	proto.RegisterType((*PeersReply)(nil), "rpc.PeersReply")
	proto.RegisterType((*FindPeerRequest)(nil), "rpc.FindPeerRequest")
	proto.RegisterType((*FindPeerReply)(nil), "rpc.FindPeerReply")
	proto.RegisterType((*ConnectPeerRequest)(nil), "rpc.ConnectPeerRequest")
	proto.RegisterType((*ConnectPeerReply)(nil), "rpc.ConnectPeerReply")
	proto.RegisterType((*DisconnectPeerRequest)(nil), "rpc.DisconnectPeerRequest")
	proto.RegisterType((*DisconnectPeerReply)(nil), "rpc.DisconnectPeerReply")
	proto.RegisterType((*ConnectednessRequest)(nil), "rpc.ConnectednessRequest")
	proto.RegisterType((*ConnectednessReply)(nil), "rpc.ConnectednessReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0x6b, 0x1b, 0x97, 0x7a, 0x12, 0x3b, 0xce, 0x34, 0x69, 0x83, 0xc5, 0x45, 0x64, 0x6e,
	0x5a, 0x24, 0x0c, 0x04, 0x84, 0x10, 0xaa, 0x54, 0xb5, 0x49, 0x10, 0x91, 0x50, 0x14, 0x45, 0xe5,
	0x01, 0x8c, 0xbd, 0x20, 0xab, 0xd1, 0xae, 0x59, 0x2f, 0x2a, 0x7d, 0x1d, 0x9e, 0x93, 0x0b, 0xb4,
	0xeb, 0x9f, 0xd8, 0x71, 0x2a, 0xc4, 0xe5, 0x9c, 0x33, 0xe7, 0xf3, 0xee, 0xce, 0xc8, 0x60, 0x85,
	0x69, 0x12, 0xa4, 0x9c, 0x09, 0x86, 0x06, 0x4f, 0x23, 0xff, 0x15, 0x1c, 0x5d, 0xc5, 0x31, 0x5f,
	0xd0, 0x6f, 0x0c, 0x1d, 0xd0, 0x17, 0xb3, 0x91, 0x36, 0xd6, 0xce, 0xac, 0xb5, 0xbe, 0x98, 0xe1,
	0x00, 0xcc, 0x30, 0x8e, 0x79, 0x36, 0xd2, 0xc7, 0xc6, 0x99, 0xb5, 0xce, 0x0b, 0xff, 0x18, 0xfa,
	0x9f, 0x93, 0x4c, 0x10, 0x2a, 0x73, 0x6b, 0xf2, 0xe3, 0x27, 0xc9, 0x84, 0x7f, 0x01, 0xbd, 0xba,
	0x98, 0x6e, 0xee, 0xf1, 0x1c, 0x8e, 0xc2, 0x82, 0xac, 0x98, 0x9d, 0x89, 0x1d, 0xf0, 0x34, 0x0a,
	0xca, 0xcf, 0xad, 0x2b, 0xdb, 0x77, 0xa0, 0xbb, 0x22, 0x84, 0x67, 0x25, 0xed, 0x35, 0x40, 0x51,
	0x4b, 0xd0, 0x33, 0x30, 0x53, 0x59, 0x8d, 0xb4, 0xb1, 0xd1, 0xa6, 0xe4, 0x9e, 0x7f, 0x0e, 0xbd,
	0x8f, 0x09, 0x8d, 0x65, 0xac, 0xa0, 0xe0, 0x09, 0x1c, 0x4a, 0xaf, 0xba, 0x52, 0x51, 0xf9, 0x1f,
	0xc0, 0xde, 0xb6, 0x16, 0x27, 0x55, 0xd6, 0xc3, 0x27, 0x2d, 0x6d, 0xff, 0x12, 0x70, 0xca, 0x28,
	0x25, 0x91, 0xa8, 0x7f, 0xe9, 0x3f, 0x00, 0x08, 0x6e, 0x03, 0x90, 0x6e, 0xee, 0xfd, 0x97, 0x30,
	0x9c, 0x25, 0x59, 0xd4, 0xe6, 0x3e, 0x74, 0x83, 0x21, 0x1c, 0xef, 0x06, 0x24, 0x27, 0x80, 0x41,
	0xc1, 0x26, 0x31, 0x25, 0x59, 0xf6, 0x2f, 0xcc, 0xb2, 0xba, 0x4c, 0xd9, 0x2f, 0x5f, 0xe3, 0x3d,
	0xd8, 0x51, 0x5d, 0x55, 0x21, 0x67, 0x82, 0xea, 0x46, 0xcd, 0xfe, 0x66, 0xe3, 0xf3, 0x5b, 0xb0,
	0x1b, 0x3e, 0xba, 0xd0, 0x5d, 0x32, 0x51, 0x69, 0xee, 0x01, 0xda, 0x60, 0x6d, 0x4b, 0x0d, 0x1d,
	0x80, 0x69, 0x48, 0x0b, 0xc5, 0xd5, 0xb1, 0x0f, 0xf6, 0x34, 0xa4, 0xb4, 0xca, 0xb8, 0x06, 0x76,
	0xe0, 0xf1, 0x17, 0x7a, 0x4b, 0xd9, 0x1d, 0x75, 0x1f, 0xa1, 0x05, 0xe6, 0x9c, 0x73, 0xc6, 0x5d,
	0x73, 0xf2, 0x47, 0x07, 0xe3, 0x6a, 0xb5, 0xc0, 0x0b, 0x80, 0xed, 0xe6, 0xe1, 0x89, 0x3a, 0x65,
	0x6b, 0x3f, 0xbd, 0x41, 0x4b, 0x97, 0x0f, 0x76, 0x80, 0x2f, 0xc0, 0x54, 0x9b, 0x86, 0x7d, 0xd5,
	0x50, 0xdf, 0x42, 0xaf, 0x57, 0x97, 0xf2, 0xf6, 0x77, 0x70, 0x54, 0xae, 0x0e, 0xe6, 0xc8, 0x9d,
	0xa5, 0xf3, 0x70, 0x47, 0xcd, 0x73, 0x97, 0xd0, 0xa9, 0x4d, 0x1d, 0x4f, 0xeb, 0x6f, 0x59, 0x4f,
	0x0f, 0xdb, 0x46, 0x0e, 0xf8, 0x04, 0x4e, 0x73, 0xe2, 0xe8, 0xa9, 0xd6, 0xbd, 0x7b, 0xe3, 0x8d,
	0xf6, 0x7a, 0x39, 0x69, 0xbe, 0x3b, 0xa4, 0x27, 0x7b, 0x06, 0x5b, 0x70, 0x4e, 0xf7, 0x59, 0x0a,
	0x73, 0xfd, 0x16, 0x9e, 0x26, 0x2c, 0x10, 0xe4, 0x97, 0x48, 0x36, 0x24, 0x48, 0xd9, 0x1d, 0xe1,
	0xdf, 0x43, 0x41, 0x02, 0x4a, 0x84, 0x0c, 0x5d, 0x77, 0x57, 0xa5, 0xb4, 0x24, 0x62, 0xa5, 0xfd,
	0xd6, 0x8d, 0x9b, 0x9b, 0xf9, 0xd7, 0x43, 0xf5, 0xe7, 0x79, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x3f, 0xb0, 0xda, 0x86, 0x04, 0x00, 0x00,
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
	ListenAddr(ctx context.Context, in *ListenAddrRequest, opts ...grpc.CallOption) (*ListenAddrReply, error)
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersReply, error)
	FindPeer(ctx context.Context, in *FindPeerRequest, opts ...grpc.CallOption) (*FindPeerReply, error)
	ConnectPeer(ctx context.Context, in *ConnectPeerRequest, opts ...grpc.CallOption) (*ConnectPeerReply, error)
	DisconnectPeer(ctx context.Context, in *DisconnectPeerRequest, opts ...grpc.CallOption) (*DisconnectPeerReply, error)
	Connectedness(ctx context.Context, in *ConnectednessRequest, opts ...grpc.CallOption) (*ConnectednessReply, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) ListenAddr(ctx context.Context, in *ListenAddrRequest, opts ...grpc.CallOption) (*ListenAddrReply, error) {
	out := new(ListenAddrReply)
	err := c.cc.Invoke(ctx, "/rpc.API/ListenAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersReply, error) {
	out := new(PeersReply)
	err := c.cc.Invoke(ctx, "/rpc.API/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) FindPeer(ctx context.Context, in *FindPeerRequest, opts ...grpc.CallOption) (*FindPeerReply, error) {
	out := new(FindPeerReply)
	err := c.cc.Invoke(ctx, "/rpc.API/FindPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ConnectPeer(ctx context.Context, in *ConnectPeerRequest, opts ...grpc.CallOption) (*ConnectPeerReply, error) {
	out := new(ConnectPeerReply)
	err := c.cc.Invoke(ctx, "/rpc.API/ConnectPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DisconnectPeer(ctx context.Context, in *DisconnectPeerRequest, opts ...grpc.CallOption) (*DisconnectPeerReply, error) {
	out := new(DisconnectPeerReply)
	err := c.cc.Invoke(ctx, "/rpc.API/DisconnectPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Connectedness(ctx context.Context, in *ConnectednessRequest, opts ...grpc.CallOption) (*ConnectednessReply, error) {
	out := new(ConnectednessReply)
	err := c.cc.Invoke(ctx, "/rpc.API/Connectedness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	ListenAddr(context.Context, *ListenAddrRequest) (*ListenAddrReply, error)
	Peers(context.Context, *PeersRequest) (*PeersReply, error)
	FindPeer(context.Context, *FindPeerRequest) (*FindPeerReply, error)
	ConnectPeer(context.Context, *ConnectPeerRequest) (*ConnectPeerReply, error)
	DisconnectPeer(context.Context, *DisconnectPeerRequest) (*DisconnectPeerReply, error)
	Connectedness(context.Context, *ConnectednessRequest) (*ConnectednessReply, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) ListenAddr(ctx context.Context, req *ListenAddrRequest) (*ListenAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenAddr not implemented")
}
func (*UnimplementedAPIServer) Peers(ctx context.Context, req *PeersRequest) (*PeersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (*UnimplementedAPIServer) FindPeer(ctx context.Context, req *FindPeerRequest) (*FindPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPeer not implemented")
}
func (*UnimplementedAPIServer) ConnectPeer(ctx context.Context, req *ConnectPeerRequest) (*ConnectPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectPeer not implemented")
}
func (*UnimplementedAPIServer) DisconnectPeer(ctx context.Context, req *DisconnectPeerRequest) (*DisconnectPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectPeer not implemented")
}
func (*UnimplementedAPIServer) Connectedness(ctx context.Context, req *ConnectednessRequest) (*ConnectednessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connectedness not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_ListenAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListenAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.API/ListenAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListenAddr(ctx, req.(*ListenAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.API/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_FindPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).FindPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.API/FindPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).FindPeer(ctx, req.(*FindPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ConnectPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ConnectPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.API/ConnectPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ConnectPeer(ctx, req.(*ConnectPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DisconnectPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DisconnectPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.API/DisconnectPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DisconnectPeer(ctx, req.(*DisconnectPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Connectedness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectednessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Connectedness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.API/Connectedness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Connectedness(ctx, req.(*ConnectednessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListenAddr",
			Handler:    _API_ListenAddr_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _API_Peers_Handler,
		},
		{
			MethodName: "FindPeer",
			Handler:    _API_FindPeer_Handler,
		},
		{
			MethodName: "ConnectPeer",
			Handler:    _API_ConnectPeer_Handler,
		},
		{
			MethodName: "DisconnectPeer",
			Handler:    _API_DisconnectPeer_Handler,
		},
		{
			MethodName: "Connectedness",
			Handler:    _API_Connectedness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

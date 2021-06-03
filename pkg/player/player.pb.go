// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

package player

import (
	"context"
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type StateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateRequest) Reset()         { *m = StateRequest{} }
func (m *StateRequest) String() string { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()    {}
func (*StateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{0}
}

func (m *StateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRequest.Unmarshal(m, b)
}
func (m *StateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRequest.Marshal(b, m, deterministic)
}
func (m *StateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRequest.Merge(m, src)
}
func (m *StateRequest) XXX_Size() int {
	return xxx_messageInfo_StateRequest.Size(m)
}
func (m *StateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StateRequest proto.InternalMessageInfo

type StateResponse struct {
	Ports                []string `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	Storages             []string `protobuf:"bytes,2,rep,name=storages,proto3" json:"storages,omitempty"`
	Devices              []string `protobuf:"bytes,3,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{1}
}

func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateResponse.Unmarshal(m, b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
}
func (m *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(m, src)
}
func (m *StateResponse) XXX_Size() int {
	return xxx_messageInfo_StateResponse.Size(m)
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func (m *StateResponse) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *StateResponse) GetStorages() []string {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *StateResponse) GetDevices() []string {
	if m != nil {
		return m.Devices
	}
	return nil
}

type StartReceiveRequest struct {
	Port                 string                `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	StorageUUID          *wrappers.StringValue `protobuf:"bytes,2,opt,name=storageUUID,proto3" json:"storageUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StartReceiveRequest) Reset()         { *m = StartReceiveRequest{} }
func (m *StartReceiveRequest) String() string { return proto.CompactTextString(m) }
func (*StartReceiveRequest) ProtoMessage()    {}
func (*StartReceiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{2}
}

func (m *StartReceiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartReceiveRequest.Unmarshal(m, b)
}
func (m *StartReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartReceiveRequest.Marshal(b, m, deterministic)
}
func (m *StartReceiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartReceiveRequest.Merge(m, src)
}
func (m *StartReceiveRequest) XXX_Size() int {
	return xxx_messageInfo_StartReceiveRequest.Size(m)
}
func (m *StartReceiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartReceiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartReceiveRequest proto.InternalMessageInfo

func (m *StartReceiveRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *StartReceiveRequest) GetStorageUUID() *wrappers.StringValue {
	if m != nil {
		return m.StorageUUID
	}
	return nil
}

type StartReceiveResponse struct {
	StorageUUID          string   `protobuf:"bytes,1,opt,name=storageUUID,proto3" json:"storageUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartReceiveResponse) Reset()         { *m = StartReceiveResponse{} }
func (m *StartReceiveResponse) String() string { return proto.CompactTextString(m) }
func (*StartReceiveResponse) ProtoMessage()    {}
func (*StartReceiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{3}
}

func (m *StartReceiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartReceiveResponse.Unmarshal(m, b)
}
func (m *StartReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartReceiveResponse.Marshal(b, m, deterministic)
}
func (m *StartReceiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartReceiveResponse.Merge(m, src)
}
func (m *StartReceiveResponse) XXX_Size() int {
	return xxx_messageInfo_StartReceiveResponse.Size(m)
}
func (m *StartReceiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartReceiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartReceiveResponse proto.InternalMessageInfo

func (m *StartReceiveResponse) GetStorageUUID() string {
	if m != nil {
		return m.StorageUUID
	}
	return ""
}

type StopReceiveRequest struct {
	Port                 string   `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopReceiveRequest) Reset()         { *m = StopReceiveRequest{} }
func (m *StopReceiveRequest) String() string { return proto.CompactTextString(m) }
func (*StopReceiveRequest) ProtoMessage()    {}
func (*StopReceiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{4}
}

func (m *StopReceiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopReceiveRequest.Unmarshal(m, b)
}
func (m *StopReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopReceiveRequest.Marshal(b, m, deterministic)
}
func (m *StopReceiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopReceiveRequest.Merge(m, src)
}
func (m *StopReceiveRequest) XXX_Size() int {
	return xxx_messageInfo_StopReceiveRequest.Size(m)
}
func (m *StopReceiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopReceiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopReceiveRequest proto.InternalMessageInfo

func (m *StopReceiveRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type StopReceiveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopReceiveResponse) Reset()         { *m = StopReceiveResponse{} }
func (m *StopReceiveResponse) String() string { return proto.CompactTextString(m) }
func (*StopReceiveResponse) ProtoMessage()    {}
func (*StopReceiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{5}
}

func (m *StopReceiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopReceiveResponse.Unmarshal(m, b)
}
func (m *StopReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopReceiveResponse.Marshal(b, m, deterministic)
}
func (m *StopReceiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopReceiveResponse.Merge(m, src)
}
func (m *StopReceiveResponse) XXX_Size() int {
	return xxx_messageInfo_StopReceiveResponse.Size(m)
}
func (m *StopReceiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopReceiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopReceiveResponse proto.InternalMessageInfo

type StartPlayRequest struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	Channels             uint32   `protobuf:"varint,2,opt,name=channels,proto3" json:"channels,omitempty"`
	Rate                 uint32   `protobuf:"varint,3,opt,name=rate,proto3" json:"rate,omitempty"`
	BitsPerSample        uint32   `protobuf:"varint,4,opt,name=bitsPerSample,proto3" json:"bitsPerSample,omitempty"`
	StorageUUID          string   `protobuf:"bytes,5,opt,name=storageUUID,proto3" json:"storageUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartPlayRequest) Reset()         { *m = StartPlayRequest{} }
func (m *StartPlayRequest) String() string { return proto.CompactTextString(m) }
func (*StartPlayRequest) ProtoMessage()    {}
func (*StartPlayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{6}
}

func (m *StartPlayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartPlayRequest.Unmarshal(m, b)
}
func (m *StartPlayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartPlayRequest.Marshal(b, m, deterministic)
}
func (m *StartPlayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartPlayRequest.Merge(m, src)
}
func (m *StartPlayRequest) XXX_Size() int {
	return xxx_messageInfo_StartPlayRequest.Size(m)
}
func (m *StartPlayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartPlayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartPlayRequest proto.InternalMessageInfo

func (m *StartPlayRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *StartPlayRequest) GetChannels() uint32 {
	if m != nil {
		return m.Channels
	}
	return 0
}

func (m *StartPlayRequest) GetRate() uint32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *StartPlayRequest) GetBitsPerSample() uint32 {
	if m != nil {
		return m.BitsPerSample
	}
	return 0
}

func (m *StartPlayRequest) GetStorageUUID() string {
	if m != nil {
		return m.StorageUUID
	}
	return ""
}

type StartPlayResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartPlayResponse) Reset()         { *m = StartPlayResponse{} }
func (m *StartPlayResponse) String() string { return proto.CompactTextString(m) }
func (*StartPlayResponse) ProtoMessage()    {}
func (*StartPlayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{7}
}

func (m *StartPlayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartPlayResponse.Unmarshal(m, b)
}
func (m *StartPlayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartPlayResponse.Marshal(b, m, deterministic)
}
func (m *StartPlayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartPlayResponse.Merge(m, src)
}
func (m *StartPlayResponse) XXX_Size() int {
	return xxx_messageInfo_StartPlayResponse.Size(m)
}
func (m *StartPlayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartPlayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartPlayResponse proto.InternalMessageInfo

type StopPlayRequest struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopPlayRequest) Reset()         { *m = StopPlayRequest{} }
func (m *StopPlayRequest) String() string { return proto.CompactTextString(m) }
func (*StopPlayRequest) ProtoMessage()    {}
func (*StopPlayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{8}
}

func (m *StopPlayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopPlayRequest.Unmarshal(m, b)
}
func (m *StopPlayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopPlayRequest.Marshal(b, m, deterministic)
}
func (m *StopPlayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopPlayRequest.Merge(m, src)
}
func (m *StopPlayRequest) XXX_Size() int {
	return xxx_messageInfo_StopPlayRequest.Size(m)
}
func (m *StopPlayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopPlayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopPlayRequest proto.InternalMessageInfo

func (m *StopPlayRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

type StopPlayResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopPlayResponse) Reset()         { *m = StopPlayResponse{} }
func (m *StopPlayResponse) String() string { return proto.CompactTextString(m) }
func (*StopPlayResponse) ProtoMessage()    {}
func (*StopPlayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{9}
}

func (m *StopPlayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopPlayResponse.Unmarshal(m, b)
}
func (m *StopPlayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopPlayResponse.Marshal(b, m, deterministic)
}
func (m *StopPlayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopPlayResponse.Merge(m, src)
}
func (m *StopPlayResponse) XXX_Size() int {
	return xxx_messageInfo_StopPlayResponse.Size(m)
}
func (m *StopPlayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopPlayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopPlayResponse proto.InternalMessageInfo

type ClearStorageRequest struct {
	StorageUUID          string   `protobuf:"bytes,1,opt,name=storageUUID,proto3" json:"storageUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearStorageRequest) Reset()         { *m = ClearStorageRequest{} }
func (m *ClearStorageRequest) String() string { return proto.CompactTextString(m) }
func (*ClearStorageRequest) ProtoMessage()    {}
func (*ClearStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{10}
}

func (m *ClearStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearStorageRequest.Unmarshal(m, b)
}
func (m *ClearStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearStorageRequest.Marshal(b, m, deterministic)
}
func (m *ClearStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearStorageRequest.Merge(m, src)
}
func (m *ClearStorageRequest) XXX_Size() int {
	return xxx_messageInfo_ClearStorageRequest.Size(m)
}
func (m *ClearStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearStorageRequest proto.InternalMessageInfo

func (m *ClearStorageRequest) GetStorageUUID() string {
	if m != nil {
		return m.StorageUUID
	}
	return ""
}

type ClearStorageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearStorageResponse) Reset()         { *m = ClearStorageResponse{} }
func (m *ClearStorageResponse) String() string { return proto.CompactTextString(m) }
func (*ClearStorageResponse) ProtoMessage()    {}
func (*ClearStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{11}
}

func (m *ClearStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearStorageResponse.Unmarshal(m, b)
}
func (m *ClearStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearStorageResponse.Marshal(b, m, deterministic)
}
func (m *ClearStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearStorageResponse.Merge(m, src)
}
func (m *ClearStorageResponse) XXX_Size() int {
	return xxx_messageInfo_ClearStorageResponse.Size(m)
}
func (m *ClearStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearStorageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StateRequest)(nil), "player.StateRequest")
	proto.RegisterType((*StateResponse)(nil), "player.StateResponse")
	proto.RegisterType((*StartReceiveRequest)(nil), "player.StartReceiveRequest")
	proto.RegisterType((*StartReceiveResponse)(nil), "player.StartReceiveResponse")
	proto.RegisterType((*StopReceiveRequest)(nil), "player.StopReceiveRequest")
	proto.RegisterType((*StopReceiveResponse)(nil), "player.StopReceiveResponse")
	proto.RegisterType((*StartPlayRequest)(nil), "player.StartPlayRequest")
	proto.RegisterType((*StartPlayResponse)(nil), "player.StartPlayResponse")
	proto.RegisterType((*StopPlayRequest)(nil), "player.StopPlayRequest")
	proto.RegisterType((*StopPlayResponse)(nil), "player.StopPlayResponse")
	proto.RegisterType((*ClearStorageRequest)(nil), "player.ClearStorageRequest")
	proto.RegisterType((*ClearStorageResponse)(nil), "player.ClearStorageResponse")
}

func init() { proto.RegisterFile("player.proto", fileDescriptor_41d803d1b635d5c6) }

var fileDescriptor_41d803d1b635d5c6 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x9b, 0x0f, 0xe8, 0x24, 0xe1, 0x63, 0x93, 0xc2, 0xe2, 0x56, 0x55, 0xb4, 0xe2, 0x90,
	0x93, 0x2b, 0x8a, 0x04, 0x48, 0x08, 0x2e, 0x70, 0x00, 0x21, 0xa1, 0xc8, 0x56, 0xb9, 0x70, 0xda,
	0x84, 0x21, 0x58, 0x72, 0xb3, 0xcb, 0xee, 0xa6, 0xa8, 0xff, 0x86, 0x33, 0xbf, 0x12, 0x79, 0x77,
	0xed, 0xac, 0x53, 0x4b, 0x51, 0x6f, 0x9e, 0x99, 0x7d, 0xef, 0xcd, 0x9b, 0x19, 0xc3, 0x50, 0x16,
	0xfc, 0x06, 0x55, 0x22, 0x95, 0x30, 0x82, 0xf4, 0x5d, 0x14, 0x9f, 0xad, 0x84, 0x58, 0x15, 0x78,
	0x6e, 0xb3, 0x8b, 0xcd, 0xcf, 0xf3, 0x3f, 0x8a, 0x4b, 0x89, 0x4a, 0xbb, 0x77, 0xec, 0x01, 0x0c,
	0x33, 0xc3, 0x0d, 0xa6, 0xf8, 0x7b, 0x83, 0xda, 0xb0, 0xef, 0x30, 0xf2, 0xb1, 0x96, 0x62, 0xad,
	0x91, 0x4c, 0xa0, 0x27, 0x85, 0x32, 0x9a, 0x46, 0xd3, 0xce, 0xec, 0x28, 0x75, 0x01, 0x89, 0xe1,
	0xbe, 0x36, 0x42, 0xf1, 0x15, 0x6a, 0x7a, 0x68, 0x0b, 0x75, 0x4c, 0x28, 0xdc, 0xfb, 0x81, 0xd7,
	0xf9, 0x12, 0x35, 0xed, 0xd8, 0x52, 0x15, 0xb2, 0x1c, 0xc6, 0x99, 0xe1, 0xca, 0xa4, 0xb8, 0xc4,
	0xfc, 0xba, 0xd2, 0x24, 0x04, 0xba, 0x25, 0x2b, 0x8d, 0xa6, 0xd1, 0xec, 0x28, 0xb5, 0xdf, 0xe4,
	0x3d, 0x0c, 0x3c, 0xe1, 0xe5, 0xe5, 0xe7, 0x8f, 0xf4, 0x70, 0x1a, 0xcd, 0x06, 0x17, 0xa7, 0x89,
	0x73, 0x93, 0x54, 0x6e, 0x92, 0xcc, 0xa8, 0x7c, 0xbd, 0xfa, 0xc6, 0x8b, 0x0d, 0xa6, 0x21, 0x80,
	0xbd, 0x81, 0x49, 0x53, 0xca, 0xdb, 0x99, 0x36, 0x79, 0x9d, 0x64, 0x03, 0x39, 0x03, 0x92, 0x19,
	0x21, 0xf7, 0xf7, 0xc8, 0x8e, 0x4b, 0x3b, 0xc1, 0x4b, 0x27, 0xc1, 0xfe, 0x45, 0xf0, 0xc8, 0x6a,
	0xcf, 0x0b, 0x7e, 0x53, 0xe1, 0xcf, 0x00, 0xdc, 0x14, 0xbe, 0xf2, 0x2b, 0xf4, 0x2c, 0x41, 0xa6,
	0x1c, 0xe8, 0xf2, 0x17, 0x5f, 0xaf, 0xb1, 0xd0, 0xd6, 0xec, 0x28, 0xad, 0xe3, 0x52, 0x5b, 0x71,
	0x83, 0xb4, 0x63, 0xf3, 0xf6, 0x9b, 0x3c, 0x87, 0xd1, 0x22, 0x37, 0x7a, 0x8e, 0x2a, 0xe3, 0x57,
	0xb2, 0x40, 0xda, 0xb5, 0xc5, 0x66, 0x72, 0xd7, 0x6d, 0xef, 0xb6, 0xdb, 0x31, 0x3c, 0x0e, 0x7a,
	0xf5, 0x0e, 0x5e, 0xc0, 0xc3, 0xd2, 0xd8, 0x1d, 0xfa, 0x67, 0xa4, 0xf4, 0x5c, 0x41, 0x3c, 0xcd,
	0x6b, 0x18, 0x7f, 0x28, 0x90, 0xab, 0xcc, 0xe9, 0x55, 0x54, 0xfb, 0x57, 0xf0, 0x04, 0x26, 0x4d,
	0xa0, 0x23, 0xbc, 0xf8, 0xdb, 0x81, 0xfe, 0xdc, 0xde, 0x35, 0x79, 0x05, 0x3d, 0x7b, 0xa7, 0x64,
	0x92, 0xf8, 0xbb, 0x0f, 0xcf, 0x38, 0x3e, 0xde, 0xc9, 0xfa, 0x8e, 0x0e, 0xc8, 0x17, 0x18, 0xfa,
	0x7d, 0x59, 0xdb, 0xe4, 0x24, 0x78, 0xb8, 0x7b, 0x98, 0xf1, 0x69, 0x7b, 0xb1, 0x26, 0xfb, 0x04,
	0x83, 0x9a, 0x4c, 0x48, 0x12, 0x6f, 0x9f, 0xef, 0xde, 0x4f, 0x7c, 0xd2, 0x5a, 0xab, 0x99, 0xde,
	0x41, 0xb7, 0x34, 0x46, 0x68, 0x43, 0x31, 0x58, 0x40, 0xfc, 0xac, 0xa5, 0x52, 0xc3, 0xdf, 0x42,
	0xd7, 0x76, 0xf0, 0x34, 0x54, 0x09, 0xd1, 0xf4, 0x76, 0x21, 0x1c, 0x49, 0x38, 0xed, 0xed, 0x48,
	0x5a, 0x96, 0xb7, 0x1d, 0x49, 0xdb, 0x82, 0xd8, 0xc1, 0xa2, 0x6f, 0x7f, 0xcd, 0x97, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x34, 0xfd, 0x36, 0xc0, 0x8e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlayerClient is the client API for Player service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerClient interface {
	// State return receiving ports, storages and busy device
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	// Start receive data on port and save on storage
	ReceiveStart(ctx context.Context, in *StartReceiveRequest, opts ...grpc.CallOption) (*StartReceiveResponse, error)
	ReceiveStop(ctx context.Context, in *StopReceiveRequest, opts ...grpc.CallOption) (*StopReceiveResponse, error)
	// Play audio on deviceName drom storage
	Play(ctx context.Context, in *StartPlayRequest, opts ...grpc.CallOption) (*StartPlayResponse, error)
	// Stop audio on deviceName
	Stop(ctx context.Context, in *StopPlayRequest, opts ...grpc.CallOption) (*StopPlayResponse, error)
	ClearStorage(ctx context.Context, in *ClearStorageRequest, opts ...grpc.CallOption) (*ClearStorageResponse, error)
}

type playerClient struct {
	cc *grpc.ClientConn
}

func NewPlayerClient(cc *grpc.ClientConn) PlayerClient {
	return &playerClient{cc}
}

func (c *playerClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, "/player.Player/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) ReceiveStart(ctx context.Context, in *StartReceiveRequest, opts ...grpc.CallOption) (*StartReceiveResponse, error) {
	out := new(StartReceiveResponse)
	err := c.cc.Invoke(ctx, "/player.Player/ReceiveStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) ReceiveStop(ctx context.Context, in *StopReceiveRequest, opts ...grpc.CallOption) (*StopReceiveResponse, error) {
	out := new(StopReceiveResponse)
	err := c.cc.Invoke(ctx, "/player.Player/ReceiveStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) Play(ctx context.Context, in *StartPlayRequest, opts ...grpc.CallOption) (*StartPlayResponse, error) {
	out := new(StartPlayResponse)
	err := c.cc.Invoke(ctx, "/player.Player/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) Stop(ctx context.Context, in *StopPlayRequest, opts ...grpc.CallOption) (*StopPlayResponse, error) {
	out := new(StopPlayResponse)
	err := c.cc.Invoke(ctx, "/player.Player/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) ClearStorage(ctx context.Context, in *ClearStorageRequest, opts ...grpc.CallOption) (*ClearStorageResponse, error) {
	out := new(ClearStorageResponse)
	err := c.cc.Invoke(ctx, "/player.Player/ClearStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServer is the server API for Player service.
type PlayerServer interface {
	// State return receiving ports, storages and busy device
	State(context.Context, *StateRequest) (*StateResponse, error)
	// Start receive data on port and save on storage
	ReceiveStart(context.Context, *StartReceiveRequest) (*StartReceiveResponse, error)
	ReceiveStop(context.Context, *StopReceiveRequest) (*StopReceiveResponse, error)
	// Play audio on deviceName drom storage
	Play(context.Context, *StartPlayRequest) (*StartPlayResponse, error)
	// Stop audio on deviceName
	Stop(context.Context, *StopPlayRequest) (*StopPlayResponse, error)
	ClearStorage(context.Context, *ClearStorageRequest) (*ClearStorageResponse, error)
}

// UnimplementedPlayerServer can be embedded to have forward compatible implementations.
type UnimplementedPlayerServer struct {
}

func (*UnimplementedPlayerServer) State(ctx context.Context, req *StateRequest) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (*UnimplementedPlayerServer) ReceiveStart(ctx context.Context, req *StartReceiveRequest) (*StartReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveStart not implemented")
}
func (*UnimplementedPlayerServer) ReceiveStop(ctx context.Context, req *StopReceiveRequest) (*StopReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveStop not implemented")
}
func (*UnimplementedPlayerServer) Play(ctx context.Context, req *StartPlayRequest) (*StartPlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (*UnimplementedPlayerServer) Stop(ctx context.Context, req *StopPlayRequest) (*StopPlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedPlayerServer) ClearStorage(ctx context.Context, req *ClearStorageRequest) (*ClearStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearStorage not implemented")
}

func RegisterPlayerServer(s *grpc.Server, srv PlayerServer) {
	s.RegisterService(&_Player_serviceDesc, srv)
}

func _Player_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.Player/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).State(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Player_ReceiveStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).ReceiveStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.Player/ReceiveStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).ReceiveStart(ctx, req.(*StartReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Player_ReceiveStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).ReceiveStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.Player/ReceiveStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).ReceiveStop(ctx, req.(*StopReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Player_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.Player/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).Play(ctx, req.(*StartPlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Player_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.Player/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).Stop(ctx, req.(*StopPlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Player_ClearStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).ClearStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/player.Player/ClearStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).ClearStorage(ctx, req.(*ClearStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Player_serviceDesc = grpc.ServiceDesc{
	ServiceName: "player.Player",
	HandlerType: (*PlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "State",
			Handler:    _Player_State_Handler,
		},
		{
			MethodName: "ReceiveStart",
			Handler:    _Player_ReceiveStart_Handler,
		},
		{
			MethodName: "ReceiveStop",
			Handler:    _Player_ReceiveStop_Handler,
		},
		{
			MethodName: "Play",
			Handler:    _Player_Play_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Player_Stop_Handler,
		},
		{
			MethodName: "ClearStorage",
			Handler:    _Player_ClearStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player.proto",
}
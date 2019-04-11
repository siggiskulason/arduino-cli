// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package rpc // import "github.com/arduino/arduino-cli/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Instance struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b292d2ccd88efb9a, []int{0}
}
func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (dst *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(dst, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DownloadProgress struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	File                 string   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	TotalSize            int64    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	Downloaded           int64    `protobuf:"varint,4,opt,name=downloaded,proto3" json:"downloaded,omitempty"`
	Completed            bool     `protobuf:"varint,5,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadProgress) Reset()         { *m = DownloadProgress{} }
func (m *DownloadProgress) String() string { return proto.CompactTextString(m) }
func (*DownloadProgress) ProtoMessage()    {}
func (*DownloadProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b292d2ccd88efb9a, []int{1}
}
func (m *DownloadProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadProgress.Unmarshal(m, b)
}
func (m *DownloadProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadProgress.Marshal(b, m, deterministic)
}
func (dst *DownloadProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadProgress.Merge(dst, src)
}
func (m *DownloadProgress) XXX_Size() int {
	return xxx_messageInfo_DownloadProgress.Size(m)
}
func (m *DownloadProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadProgress.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadProgress proto.InternalMessageInfo

func (m *DownloadProgress) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DownloadProgress) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *DownloadProgress) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *DownloadProgress) GetDownloaded() int64 {
	if m != nil {
		return m.Downloaded
	}
	return 0
}

func (m *DownloadProgress) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type TaskProgress struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Completed            bool     `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskProgress) Reset()         { *m = TaskProgress{} }
func (m *TaskProgress) String() string { return proto.CompactTextString(m) }
func (*TaskProgress) ProtoMessage()    {}
func (*TaskProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_b292d2ccd88efb9a, []int{2}
}
func (m *TaskProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskProgress.Unmarshal(m, b)
}
func (m *TaskProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskProgress.Marshal(b, m, deterministic)
}
func (dst *TaskProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskProgress.Merge(dst, src)
}
func (m *TaskProgress) XXX_Size() int {
	return xxx_messageInfo_TaskProgress.Size(m)
}
func (m *TaskProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskProgress.DiscardUnknown(m)
}

var xxx_messageInfo_TaskProgress proto.InternalMessageInfo

func (m *TaskProgress) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskProgress) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TaskProgress) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func init() {
	proto.RegisterType((*Instance)(nil), "arduino.Instance")
	proto.RegisterType((*DownloadProgress)(nil), "arduino.DownloadProgress")
	proto.RegisterType((*TaskProgress)(nil), "arduino.TaskProgress")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_b292d2ccd88efb9a) }

var fileDescriptor_common_b292d2ccd88efb9a = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xe9, 0x74, 0xc6, 0x99, 0x5e, 0x06, 0x19, 0xee, 0x2a, 0x88, 0x8a, 0x14, 0x17, 0x6e,
	0x9c, 0x59, 0xf8, 0x06, 0xe2, 0xc6, 0x9d, 0x54, 0x57, 0xb3, 0x91, 0xb4, 0x89, 0x35, 0x98, 0xe6,
	0x96, 0x24, 0x45, 0xf0, 0x3d, 0x7c, 0x5f, 0xaf, 0xa1, 0xf5, 0x67, 0x56, 0x3d, 0xfd, 0x4e, 0x38,
	0x7c, 0x5c, 0x58, 0x37, 0xd4, 0x75, 0xe4, 0xb6, 0xbd, 0xa7, 0x48, 0xb8, 0x94, 0x5e, 0x0d, 0xc6,
	0x51, 0x79, 0x02, 0xab, 0x7b, 0x17, 0xa2, 0x74, 0x8d, 0xc6, 0x63, 0x98, 0x19, 0x25, 0xb2, 0x8b,
	0xec, 0x6a, 0x51, 0x71, 0x2a, 0x3f, 0x33, 0xd8, 0xdc, 0xd1, 0xbb, 0xb3, 0x24, 0xd5, 0x83, 0xa7,
	0xd6, 0xeb, 0x10, 0x70, 0x03, 0xf9, 0xe0, 0x6d, 0x7a, 0x55, 0x54, 0xdf, 0x11, 0x11, 0xe6, 0x2f,
	0xc6, 0x6a, 0x31, 0x4b, 0x28, 0x65, 0x3c, 0x03, 0x88, 0x14, 0xa5, 0x7d, 0x0e, 0xe6, 0x43, 0x8b,
	0x9c, 0x9b, 0xbc, 0x2a, 0x12, 0x79, 0x64, 0x80, 0xe7, 0x00, 0x6a, 0x1c, 0xd6, 0x4a, 0xcc, 0x53,
	0xfd, 0x87, 0xe0, 0x29, 0x14, 0xac, 0xdb, 0x5b, 0x1d, 0xb9, 0x5e, 0x70, 0xbd, 0xaa, 0x7e, 0x41,
	0xb9, 0x87, 0xf5, 0x93, 0x0c, 0x6f, 0x3f, 0x4a, 0x2c, 0xe0, 0x64, 0xa7, 0x47, 0xa7, 0x94, 0x51,
	0xc0, 0xb2, 0xe3, 0x4e, 0xb6, 0x93, 0xd7, 0xf4, 0xfb, 0x7f, 0x3b, 0x3f, 0xd8, 0xbe, 0xbd, 0xdc,
	0x97, 0xad, 0x89, 0xaf, 0x43, 0xbd, 0x65, 0xb6, 0x1b, 0xaf, 0x34, 0x7d, 0xaf, 0x1b, 0x6b, 0x76,
	0xbe, 0x6f, 0xea, 0xa3, 0x74, 0xc5, 0x9b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x77, 0xb1,
	0xae, 0x55, 0x01, 0x00, 0x00,
}

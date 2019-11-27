// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/kernel_def.proto

package framework

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type KernelDef struct {
	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint,proto3" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg,proto3" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelDef) Reset()         { *m = KernelDef{} }
func (m *KernelDef) String() string { return proto.CompactTextString(m) }
func (*KernelDef) ProtoMessage()    {}
func (*KernelDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_18794e085ea7671a, []int{0}
}

func (m *KernelDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef.Unmarshal(m, b)
}
func (m *KernelDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef.Marshal(b, m, deterministic)
}
func (m *KernelDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef.Merge(m, src)
}
func (m *KernelDef) XXX_Size() int {
	return xxx_messageInfo_KernelDef.Size(m)
}
func (m *KernelDef) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef proto.InternalMessageInfo

func (m *KernelDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KernelDef) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *KernelDef) GetHostMemoryArg() []string {
	if m != nil {
		return m.HostMemoryArg
	}
	return nil
}

func (m *KernelDef) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type KernelDef_AttrConstraint struct {
	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues        *AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KernelDef_AttrConstraint) Reset()         { *m = KernelDef_AttrConstraint{} }
func (m *KernelDef_AttrConstraint) String() string { return proto.CompactTextString(m) }
func (*KernelDef_AttrConstraint) ProtoMessage()    {}
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_18794e085ea7671a, []int{0, 0}
}

func (m *KernelDef_AttrConstraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef_AttrConstraint.Unmarshal(m, b)
}
func (m *KernelDef_AttrConstraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef_AttrConstraint.Marshal(b, m, deterministic)
}
func (m *KernelDef_AttrConstraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef_AttrConstraint.Merge(m, src)
}
func (m *KernelDef_AttrConstraint) XXX_Size() int {
	return xxx_messageInfo_KernelDef_AttrConstraint.Size(m)
}
func (m *KernelDef_AttrConstraint) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef_AttrConstraint.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef_AttrConstraint proto.InternalMessageInfo

func (m *KernelDef_AttrConstraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelDef_AttrConstraint) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

// A collection of KernelDefs
type KernelList struct {
	Kernel               []*KernelDef `protobuf:"bytes,1,rep,name=kernel,proto3" json:"kernel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KernelList) Reset()         { *m = KernelList{} }
func (m *KernelList) String() string { return proto.CompactTextString(m) }
func (*KernelList) ProtoMessage()    {}
func (*KernelList) Descriptor() ([]byte, []int) {
	return fileDescriptor_18794e085ea7671a, []int{1}
}

func (m *KernelList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelList.Unmarshal(m, b)
}
func (m *KernelList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelList.Marshal(b, m, deterministic)
}
func (m *KernelList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelList.Merge(m, src)
}
func (m *KernelList) XXX_Size() int {
	return xxx_messageInfo_KernelList.Size(m)
}
func (m *KernelList) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelList.DiscardUnknown(m)
}

var xxx_messageInfo_KernelList proto.InternalMessageInfo

func (m *KernelList) GetKernel() []*KernelDef {
	if m != nil {
		return m.Kernel
	}
	return nil
}

func init() {
	proto.RegisterType((*KernelDef)(nil), "tensorflow.KernelDef")
	proto.RegisterType((*KernelDef_AttrConstraint)(nil), "tensorflow.KernelDef.AttrConstraint")
	proto.RegisterType((*KernelList)(nil), "tensorflow.KernelList")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/kernel_def.proto", fileDescriptor_18794e085ea7671a)
}

var fileDescriptor_18794e085ea7671a = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0xc6, 0x49, 0xfc, 0x03, 0x8e, 0xa8, 0xb0, 0xbc, 0x42, 0xf0, 0xf2, 0x8a, 0x94, 0x22, 0x85,
	0x26, 0x60, 0x8f, 0x6d, 0x0f, 0x5a, 0x6f, 0x6d, 0x41, 0x42, 0xe9, 0xa1, 0x97, 0xb0, 0x89, 0x93,
	0x18, 0x4c, 0x32, 0x61, 0x77, 0x55, 0xfc, 0x3e, 0xfd, 0x90, 0x3d, 0x96, 0x6c, 0x24, 0xa6, 0x45,
	0x7a, 0xdb, 0x99, 0x79, 0xe6, 0xd9, 0x79, 0xf8, 0xc1, 0x8d, 0xc2, 0x4c, 0x92, 0x08, 0x13, 0x3a,
	0x38, 0x01, 0x09, 0x74, 0x42, 0xc1, 0x53, 0x3c, 0x90, 0xd8, 0x3a, 0x5b, 0x14, 0x19, 0x26, 0xde,
	0x1a, 0x43, 0x3b, 0x17, 0xa4, 0x88, 0xc1, 0x59, 0x3b, 0xfa, 0x63, 0x8f, 0x2b, 0x25, 0xbc, 0x3d,
	0x4f, 0x76, 0x58, 0xee, 0x4d, 0x3e, 0x4d, 0xe8, 0x3c, 0x6b, 0xb3, 0x25, 0x86, 0xac, 0x0f, 0x26,
	0xe5, 0x96, 0x31, 0x36, 0xa6, 0x1d, 0xd7, 0xa4, 0x9c, 0xfd, 0x87, 0xee, 0x1a, 0xf7, 0x71, 0x80,
	0x9e, 0x3a, 0xe6, 0x68, 0x99, 0x7a, 0x00, 0x65, 0xeb, 0xed, 0x98, 0x23, 0x5b, 0x02, 0x04, 0x94,
	0x49, 0x25, 0x78, 0x9c, 0x29, 0xab, 0x31, 0x6e, 0x4c, 0xbb, 0xb3, 0x2b, 0xfb, 0xfc, 0xbf, 0x5d,
	0x79, 0xdb, 0x73, 0xa5, 0xc4, 0x53, 0xa5, 0x75, 0x6b, 0x7b, 0xec, 0x1a, 0x06, 0x1b, 0x92, 0xca,
	0x4b, 0x31, 0x25, 0x71, 0xf4, 0xb8, 0x88, 0xac, 0xe6, 0xb8, 0x31, 0xed, 0xb8, 0xbd, 0xa2, 0xfd,
	0xaa, 0xbb, 0x73, 0x11, 0xb1, 0x7f, 0xd0, 0x4a, 0xb8, 0x8f, 0x89, 0xd5, 0xd2, 0x87, 0x94, 0xc5,
	0xc8, 0x87, 0xfe, 0x4f, 0x6f, 0xc6, 0xa0, 0x99, 0xf1, 0x14, 0x4f, 0x41, 0xf4, 0x9b, 0x3d, 0x40,
	0x9f, 0x27, 0x09, 0x1d, 0x70, 0x5d, 0xe6, 0x97, 0x3a, 0x4d, 0x77, 0x36, 0xac, 0x5f, 0x5b, 0xf8,
	0xbc, 0x17, 0x53, 0xb7, 0x77, 0x12, 0xeb, 0x4a, 0x4e, 0xee, 0x01, 0xca, 0x24, 0x2f, 0xb1, 0x54,
	0xec, 0x16, 0xda, 0x25, 0x00, 0xcb, 0xd0, 0x89, 0x87, 0x17, 0x13, 0xbb, 0x27, 0xd1, 0x82, 0xc0,
	0x22, 0x11, 0xd5, 0x35, 0x15, 0x90, 0xc5, 0xa0, 0x92, 0xaf, 0x0a, 0x1e, 0x72, 0x65, 0x7c, 0x3c,
	0x46, 0xb1, 0xda, 0xec, 0x7c, 0x3b, 0xa0, 0xd4, 0xa9, 0x91, 0xbc, 0xfc, 0x8c, 0xe8, 0x17, 0xe2,
	0x2f, 0xc3, 0xf0, 0xdb, 0x9a, 0xed, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xa9, 0xc6,
	0x72, 0x41, 0x02, 0x00, 0x00,
}

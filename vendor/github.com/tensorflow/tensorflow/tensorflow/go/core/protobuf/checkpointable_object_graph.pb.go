// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/checkpointable_object_graph.proto

package protobuf

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

type CheckpointableObjectGraph struct {
	Nodes                []*CheckpointableObjectGraph_CheckpointableObject `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *CheckpointableObjectGraph) Reset()         { *m = CheckpointableObjectGraph{} }
func (m *CheckpointableObjectGraph) String() string { return proto.CompactTextString(m) }
func (*CheckpointableObjectGraph) ProtoMessage()    {}
func (*CheckpointableObjectGraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_c72025cda6a024e9, []int{0}
}

func (m *CheckpointableObjectGraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointableObjectGraph.Unmarshal(m, b)
}
func (m *CheckpointableObjectGraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointableObjectGraph.Marshal(b, m, deterministic)
}
func (m *CheckpointableObjectGraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointableObjectGraph.Merge(m, src)
}
func (m *CheckpointableObjectGraph) XXX_Size() int {
	return xxx_messageInfo_CheckpointableObjectGraph.Size(m)
}
func (m *CheckpointableObjectGraph) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointableObjectGraph.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointableObjectGraph proto.InternalMessageInfo

func (m *CheckpointableObjectGraph) GetNodes() []*CheckpointableObjectGraph_CheckpointableObject {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type CheckpointableObjectGraph_CheckpointableObject struct {
	// Objects which this object depends on.
	Children []*CheckpointableObjectGraph_CheckpointableObject_ObjectReference `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
	// Serialized data specific to this object.
	Attributes []*CheckpointableObjectGraph_CheckpointableObject_SerializedTensor `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// Slot variables owned by this object.
	SlotVariables        []*CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference `protobuf:"bytes,3,rep,name=slot_variables,json=slotVariables,proto3" json:"slot_variables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                `json:"-"`
	XXX_unrecognized     []byte                                                                  `json:"-"`
	XXX_sizecache        int32                                                                   `json:"-"`
}

func (m *CheckpointableObjectGraph_CheckpointableObject) Reset() {
	*m = CheckpointableObjectGraph_CheckpointableObject{}
}
func (m *CheckpointableObjectGraph_CheckpointableObject) String() string {
	return proto.CompactTextString(m)
}
func (*CheckpointableObjectGraph_CheckpointableObject) ProtoMessage() {}
func (*CheckpointableObjectGraph_CheckpointableObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_c72025cda6a024e9, []int{0, 0}
}

func (m *CheckpointableObjectGraph_CheckpointableObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject.Unmarshal(m, b)
}
func (m *CheckpointableObjectGraph_CheckpointableObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject.Marshal(b, m, deterministic)
}
func (m *CheckpointableObjectGraph_CheckpointableObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject.Merge(m, src)
}
func (m *CheckpointableObjectGraph_CheckpointableObject) XXX_Size() int {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject.Size(m)
}
func (m *CheckpointableObjectGraph_CheckpointableObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject proto.InternalMessageInfo

func (m *CheckpointableObjectGraph_CheckpointableObject) GetChildren() []*CheckpointableObjectGraph_CheckpointableObject_ObjectReference {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *CheckpointableObjectGraph_CheckpointableObject) GetAttributes() []*CheckpointableObjectGraph_CheckpointableObject_SerializedTensor {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *CheckpointableObjectGraph_CheckpointableObject) GetSlotVariables() []*CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference {
	if m != nil {
		return m.SlotVariables
	}
	return nil
}

type CheckpointableObjectGraph_CheckpointableObject_ObjectReference struct {
	// An index into `CheckpointableObjectGraph.nodes`, indicating the object
	// being referenced.
	NodeId int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// A user-provided name for the edge.
	LocalName            string   `protobuf:"bytes,2,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) Reset() {
	*m = CheckpointableObjectGraph_CheckpointableObject_ObjectReference{}
}
func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) String() string {
	return proto.CompactTextString(m)
}
func (*CheckpointableObjectGraph_CheckpointableObject_ObjectReference) ProtoMessage() {}
func (*CheckpointableObjectGraph_CheckpointableObject_ObjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_c72025cda6a024e9, []int{0, 0, 0}
}

func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_ObjectReference.Unmarshal(m, b)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_ObjectReference.Marshal(b, m, deterministic)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_ObjectReference.Merge(m, src)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) XXX_Size() int {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_ObjectReference.Size(m)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_ObjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_ObjectReference proto.InternalMessageInfo

func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *CheckpointableObjectGraph_CheckpointableObject_ObjectReference) GetLocalName() string {
	if m != nil {
		return m.LocalName
	}
	return ""
}

type CheckpointableObjectGraph_CheckpointableObject_SerializedTensor struct {
	// A name for the Tensor. Simple variables have only one
	// `SerializedTensor` named "VARIABLE_VALUE" by convention. This value may
	// be restored on object creation as an optimization.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The full name of the variable/tensor, if applicable. Used to allow
	// name-based loading of checkpoints which were saved using an
	// object-based API. Should match the checkpoint key which would have been
	// assigned by tf.train.Saver.
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// The generated name of the Tensor in the checkpoint.
	CheckpointKey        string   `protobuf:"bytes,3,opt,name=checkpoint_key,json=checkpointKey,proto3" json:"checkpoint_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) Reset() {
	*m = CheckpointableObjectGraph_CheckpointableObject_SerializedTensor{}
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) String() string {
	return proto.CompactTextString(m)
}
func (*CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) ProtoMessage() {}
func (*CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c72025cda6a024e9, []int{0, 0, 1}
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SerializedTensor.Unmarshal(m, b)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SerializedTensor.Marshal(b, m, deterministic)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SerializedTensor.Merge(m, src)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) XXX_Size() int {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SerializedTensor.Size(m)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SerializedTensor.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SerializedTensor proto.InternalMessageInfo

func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SerializedTensor) GetCheckpointKey() string {
	if m != nil {
		return m.CheckpointKey
	}
	return ""
}

type CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference struct {
	// An index into `CheckpointableObjectGraph.nodes`, indicating the
	// variable object this slot was created for.
	OriginalVariableNodeId int32 `protobuf:"varint,1,opt,name=original_variable_node_id,json=originalVariableNodeId,proto3" json:"original_variable_node_id,omitempty"`
	// The name of the slot (e.g. "m"/"v").
	SlotName string `protobuf:"bytes,2,opt,name=slot_name,json=slotName,proto3" json:"slot_name,omitempty"`
	// An index into `CheckpointableObjectGraph.nodes`, indicating the
	// `Object` with the value of the slot variable.
	SlotVariableNodeId   int32    `protobuf:"varint,3,opt,name=slot_variable_node_id,json=slotVariableNodeId,proto3" json:"slot_variable_node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) Reset() {
	*m = CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference{}
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) String() string {
	return proto.CompactTextString(m)
}
func (*CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) ProtoMessage() {}
func (*CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_c72025cda6a024e9, []int{0, 0, 2}
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference.Unmarshal(m, b)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference.Marshal(b, m, deterministic)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference.Merge(m, src)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) XXX_Size() int {
	return xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference.Size(m)
}
func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference proto.InternalMessageInfo

func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) GetOriginalVariableNodeId() int32 {
	if m != nil {
		return m.OriginalVariableNodeId
	}
	return 0
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) GetSlotName() string {
	if m != nil {
		return m.SlotName
	}
	return ""
}

func (m *CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference) GetSlotVariableNodeId() int32 {
	if m != nil {
		return m.SlotVariableNodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckpointableObjectGraph)(nil), "tensorflow.CheckpointableObjectGraph")
	proto.RegisterType((*CheckpointableObjectGraph_CheckpointableObject)(nil), "tensorflow.CheckpointableObjectGraph.CheckpointableObject")
	proto.RegisterType((*CheckpointableObjectGraph_CheckpointableObject_ObjectReference)(nil), "tensorflow.CheckpointableObjectGraph.CheckpointableObject.ObjectReference")
	proto.RegisterType((*CheckpointableObjectGraph_CheckpointableObject_SerializedTensor)(nil), "tensorflow.CheckpointableObjectGraph.CheckpointableObject.SerializedTensor")
	proto.RegisterType((*CheckpointableObjectGraph_CheckpointableObject_SlotVariableReference)(nil), "tensorflow.CheckpointableObjectGraph.CheckpointableObject.SlotVariableReference")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/checkpointable_object_graph.proto", fileDescriptor_c72025cda6a024e9)
}

var fileDescriptor_c72025cda6a024e9 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x96, 0x1b, 0x76, 0x69, 0x06, 0xb5, 0x20, 0x8b, 0x42, 0x9a, 0x0a, 0x69, 0x85, 0x84, 0xb4,
	0xa7, 0x44, 0xc0, 0x89, 0x8a, 0x0b, 0x70, 0x40, 0xa5, 0x52, 0xa9, 0x02, 0xe2, 0xc0, 0x25, 0x72,
	0x9c, 0x49, 0xe2, 0xae, 0x37, 0x5e, 0x39, 0x0e, 0xd5, 0xf2, 0x38, 0x3c, 0x01, 0x8f, 0xc0, 0x63,
	0x71, 0x44, 0xf1, 0xfe, 0x24, 0x59, 0x85, 0x4b, 0xf7, 0x66, 0x8f, 0xbf, 0x1f, 0x7f, 0x33, 0x36,
	0x9c, 0x1b, 0x2c, 0x2b, 0xa5, 0x33, 0xa9, 0x6e, 0x43, 0xae, 0x34, 0x86, 0x0b, 0xad, 0x8c, 0x4a,
	0xea, 0x2c, 0xe4, 0x05, 0xf2, 0xd9, 0x42, 0x89, 0xd2, 0xb0, 0x44, 0x62, 0xac, 0x92, 0x1b, 0xe4,
	0x26, 0xce, 0x35, 0x5b, 0x14, 0x81, 0x05, 0x51, 0x68, 0xb9, 0xcf, 0x7f, 0x8f, 0xe1, 0xf4, 0x43,
	0x8f, 0xf1, 0xd9, 0x12, 0x3e, 0x36, 0x78, 0x7a, 0x0d, 0xa3, 0x52, 0xa5, 0x58, 0x79, 0x64, 0xe2,
	0x4c, 0x1f, 0xbc, 0x3a, 0x0f, 0x5a, 0x66, 0xf0, 0x5f, 0xd6, 0xe0, 0x49, 0xb4, 0x12, 0xf2, 0xff,
	0x8c, 0xe0, 0xf1, 0xd0, 0x39, 0xcd, 0xe0, 0x90, 0x17, 0x42, 0xa6, 0x1a, 0xcb, 0xb5, 0xdb, 0xa7,
	0xbb, 0xbb, 0x05, 0x6b, 0x53, 0xcc, 0x50, 0x63, 0xc9, 0x31, 0xda, 0x6a, 0xd3, 0x19, 0x00, 0x33,
	0x46, 0x8b, 0xa4, 0x36, 0x58, 0x79, 0x07, 0xd6, 0xe9, 0x72, 0x0f, 0xa7, 0x2f, 0xa8, 0x05, 0x93,
	0xe2, 0x27, 0xa6, 0x5f, 0xad, 0x46, 0xd4, 0x91, 0xa7, 0xb7, 0x70, 0x5c, 0x49, 0x65, 0xe2, 0x1f,
	0x4c, 0x8b, 0x86, 0x53, 0x79, 0x8e, 0x35, 0xbc, 0xde, 0xc7, 0x50, 0x2a, 0xf3, 0x6d, 0xad, 0xd7,
	0x06, 0x3c, 0xaa, 0x3a, 0xe5, 0xca, 0xbf, 0x80, 0x87, 0x3b, 0x2d, 0xa0, 0x4f, 0xe1, 0x7e, 0x33,
	0x82, 0x58, 0xa4, 0x1e, 0x99, 0x90, 0xe9, 0x28, 0x1a, 0x37, 0xdb, 0x8b, 0x94, 0x3e, 0x03, 0x90,
	0x8a, 0x33, 0x19, 0x97, 0x6c, 0x8e, 0xde, 0xc1, 0x84, 0x4c, 0xdd, 0xc8, 0xb5, 0x95, 0x2b, 0x36,
	0x47, 0xff, 0x06, 0x1e, 0xed, 0x66, 0xa4, 0x14, 0xee, 0x59, 0x30, 0xb1, 0x60, 0xbb, 0xa6, 0x67,
	0xe0, 0x66, 0xb5, 0xec, 0xa9, 0x1c, 0x36, 0x85, 0x46, 0x84, 0xbe, 0x80, 0xe3, 0xf6, 0x5d, 0xc6,
	0x33, 0x5c, 0x7a, 0x8e, 0x45, 0x1c, 0xb5, 0xd5, 0x4b, 0x5c, 0xfa, 0xbf, 0x08, 0x9c, 0x0c, 0xe6,
	0xa3, 0x6f, 0xe0, 0x54, 0x69, 0x91, 0x8b, 0x92, 0xc9, 0x6d, 0x37, 0xe3, 0x7e, 0x9e, 0x27, 0x1b,
	0xc0, 0x86, 0x7d, 0xb5, 0xca, 0x77, 0x06, 0xae, 0x1d, 0x42, 0xf7, 0x62, 0x4d, 0xc1, 0x5e, 0xec,
	0x25, 0x9c, 0xf4, 0x26, 0xb4, 0xd5, 0x74, 0xac, 0x26, 0xed, 0xb6, 0x75, 0xa5, 0xf7, 0xfe, 0xdd,
	0xf7, 0xb7, 0xb9, 0x30, 0x45, 0x9d, 0x04, 0x5c, 0xcd, 0xc3, 0xce, 0x3f, 0x1c, 0x5e, 0xe6, 0xaa,
	0xff, 0x41, 0xff, 0x12, 0x92, 0x8c, 0xed, 0xe6, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0xaf, 0x93, 0xa3, 0xc6, 0x03, 0x00, 0x00,
}

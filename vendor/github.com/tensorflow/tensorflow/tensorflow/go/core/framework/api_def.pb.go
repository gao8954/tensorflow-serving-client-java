// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/api_def.proto

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

type ApiDef_Visibility int32

const (
	// Normally this is "VISIBLE" unless you are inheriting a
	// different value from another ApiDef.
	ApiDef_DEFAULT_VISIBILITY ApiDef_Visibility = 0
	// Publicly visible in the API.
	ApiDef_VISIBLE ApiDef_Visibility = 1
	// Do not include this op in the generated API. If visibility is
	// set to 'SKIP', other fields are ignored for this op.
	ApiDef_SKIP ApiDef_Visibility = 2
	// Hide this op by putting it into an internal namespace (or whatever
	// is appropriate in the target language).
	ApiDef_HIDDEN ApiDef_Visibility = 3
)

var ApiDef_Visibility_name = map[int32]string{
	0: "DEFAULT_VISIBILITY",
	1: "VISIBLE",
	2: "SKIP",
	3: "HIDDEN",
}

var ApiDef_Visibility_value = map[string]int32{
	"DEFAULT_VISIBILITY": 0,
	"VISIBLE":            1,
	"SKIP":               2,
	"HIDDEN":             3,
}

func (x ApiDef_Visibility) String() string {
	return proto.EnumName(ApiDef_Visibility_name, int32(x))
}

func (ApiDef_Visibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a850add58b816a, []int{0, 0}
}

// Used to specify and override the default API & behavior in the
// generated code for client languages, from what you would get from
// the OpDef alone. There will be a set of ApiDefs that are common
// to all client languages, and another set per client language.
// The per-client-language ApiDefs will inherit values from the
// common ApiDefs which it can either replace or modify.
//
// We separate the API definition from the OpDef so we can evolve the
// API while remaining backwards compatible when interpretting old
// graphs.  Overrides go in an "api_def.pbtxt" file with a text-format
// ApiDefs message.
//
// WARNING: Be *very* careful changing the API for any existing op --
// you can change the semantics of existing code.  These changes may
// need to wait until a major release of TensorFlow to avoid breaking
// our compatibility promises.
type ApiDef struct {
	// Name of the op (in the OpDef) to specify the API for.
	GraphOpName string `protobuf:"bytes,1,opt,name=graph_op_name,json=graphOpName,proto3" json:"graph_op_name,omitempty"`
	// If this op is deprecated, set deprecation message to the message
	// that should be logged when this op is used.
	// The message should indicate alternative op to use, if any.
	DeprecationMessage string             `protobuf:"bytes,12,opt,name=deprecation_message,json=deprecationMessage,proto3" json:"deprecation_message,omitempty"`
	Visibility         ApiDef_Visibility  `protobuf:"varint,2,opt,name=visibility,proto3,enum=tensorflow.ApiDef_Visibility" json:"visibility,omitempty"`
	Endpoint           []*ApiDef_Endpoint `protobuf:"bytes,3,rep,name=endpoint,proto3" json:"endpoint,omitempty"`
	InArg              []*ApiDef_Arg      `protobuf:"bytes,4,rep,name=in_arg,json=inArg,proto3" json:"in_arg,omitempty"`
	OutArg             []*ApiDef_Arg      `protobuf:"bytes,5,rep,name=out_arg,json=outArg,proto3" json:"out_arg,omitempty"`
	// List of original in_arg names to specify new argument order.
	// Length of arg_order should be either empty to keep current order
	// or match size of in_arg.
	ArgOrder []string       `protobuf:"bytes,11,rep,name=arg_order,json=argOrder,proto3" json:"arg_order,omitempty"`
	Attr     []*ApiDef_Attr `protobuf:"bytes,6,rep,name=attr,proto3" json:"attr,omitempty"`
	// One-line human-readable description of what the Op does.
	Summary string `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	// Additional, longer human-readable description of what the Op does.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// Modify an existing/inherited description by adding text to the beginning
	// or end.
	DescriptionPrefix    string   `protobuf:"bytes,9,opt,name=description_prefix,json=descriptionPrefix,proto3" json:"description_prefix,omitempty"`
	DescriptionSuffix    string   `protobuf:"bytes,10,opt,name=description_suffix,json=descriptionSuffix,proto3" json:"description_suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef) Reset()         { *m = ApiDef{} }
func (m *ApiDef) String() string { return proto.CompactTextString(m) }
func (*ApiDef) ProtoMessage()    {}
func (*ApiDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a850add58b816a, []int{0}
}

func (m *ApiDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef.Unmarshal(m, b)
}
func (m *ApiDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef.Marshal(b, m, deterministic)
}
func (m *ApiDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef.Merge(m, src)
}
func (m *ApiDef) XXX_Size() int {
	return xxx_messageInfo_ApiDef.Size(m)
}
func (m *ApiDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef proto.InternalMessageInfo

func (m *ApiDef) GetGraphOpName() string {
	if m != nil {
		return m.GraphOpName
	}
	return ""
}

func (m *ApiDef) GetDeprecationMessage() string {
	if m != nil {
		return m.DeprecationMessage
	}
	return ""
}

func (m *ApiDef) GetVisibility() ApiDef_Visibility {
	if m != nil {
		return m.Visibility
	}
	return ApiDef_DEFAULT_VISIBILITY
}

func (m *ApiDef) GetEndpoint() []*ApiDef_Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *ApiDef) GetInArg() []*ApiDef_Arg {
	if m != nil {
		return m.InArg
	}
	return nil
}

func (m *ApiDef) GetOutArg() []*ApiDef_Arg {
	if m != nil {
		return m.OutArg
	}
	return nil
}

func (m *ApiDef) GetArgOrder() []string {
	if m != nil {
		return m.ArgOrder
	}
	return nil
}

func (m *ApiDef) GetAttr() []*ApiDef_Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *ApiDef) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ApiDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ApiDef) GetDescriptionPrefix() string {
	if m != nil {
		return m.DescriptionPrefix
	}
	return ""
}

func (m *ApiDef) GetDescriptionSuffix() string {
	if m != nil {
		return m.DescriptionSuffix
	}
	return ""
}

// If you specify any endpoint, this will replace all of the
// inherited endpoints.  The first endpoint should be the
// "canonical" endpoint, and should not be deprecated (unless all
// endpoints are deprecated).
type ApiDef_Endpoint struct {
	// Name should be either like "CamelCaseName" or
	// "Package.CamelCaseName". Client-language-specific ApiDefs may
	// use a snake_case convention instead of CamelCase.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set if this endpoint is deprecated. If set to true, a message suggesting
	// to use a non-deprecated endpoint instead will be printed. If all
	// endpoints are deprecated, set deprecation_message in ApiDef instead.
	Deprecated           bool     `protobuf:"varint,3,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef_Endpoint) Reset()         { *m = ApiDef_Endpoint{} }
func (m *ApiDef_Endpoint) String() string { return proto.CompactTextString(m) }
func (*ApiDef_Endpoint) ProtoMessage()    {}
func (*ApiDef_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a850add58b816a, []int{0, 0}
}

func (m *ApiDef_Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef_Endpoint.Unmarshal(m, b)
}
func (m *ApiDef_Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef_Endpoint.Marshal(b, m, deterministic)
}
func (m *ApiDef_Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef_Endpoint.Merge(m, src)
}
func (m *ApiDef_Endpoint) XXX_Size() int {
	return xxx_messageInfo_ApiDef_Endpoint.Size(m)
}
func (m *ApiDef_Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef_Endpoint proto.InternalMessageInfo

func (m *ApiDef_Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDef_Endpoint) GetDeprecated() bool {
	if m != nil {
		return m.Deprecated
	}
	return false
}

type ApiDef_Arg struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Change the name used to access this arg in the API from what
	// is used in the GraphDef.  Note that these names in `backticks`
	// will also be replaced in the summary & description fields.
	RenameTo string `protobuf:"bytes,2,opt,name=rename_to,json=renameTo,proto3" json:"rename_to,omitempty"`
	// Note: this will replace any inherited arg doc. There is no
	// current way of modifying arg descriptions (other than replacing
	// them entirely) as can be done with op descriptions.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef_Arg) Reset()         { *m = ApiDef_Arg{} }
func (m *ApiDef_Arg) String() string { return proto.CompactTextString(m) }
func (*ApiDef_Arg) ProtoMessage()    {}
func (*ApiDef_Arg) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a850add58b816a, []int{0, 1}
}

func (m *ApiDef_Arg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef_Arg.Unmarshal(m, b)
}
func (m *ApiDef_Arg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef_Arg.Marshal(b, m, deterministic)
}
func (m *ApiDef_Arg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef_Arg.Merge(m, src)
}
func (m *ApiDef_Arg) XXX_Size() int {
	return xxx_messageInfo_ApiDef_Arg.Size(m)
}
func (m *ApiDef_Arg) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef_Arg.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef_Arg proto.InternalMessageInfo

func (m *ApiDef_Arg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDef_Arg) GetRenameTo() string {
	if m != nil {
		return m.RenameTo
	}
	return ""
}

func (m *ApiDef_Arg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Description of the graph-construction-time configuration of this
// Op.  That is to say, this describes the attr fields that will
// be specified in the NodeDef.
type ApiDef_Attr struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Change the name used to access this attr in the API from what
	// is used in the GraphDef.  Note that these names in `backticks`
	// will also be replaced in the summary & description fields.
	RenameTo string `protobuf:"bytes,2,opt,name=rename_to,json=renameTo,proto3" json:"rename_to,omitempty"`
	// Specify a new default value to use for this attr.  This default
	// will be used when creating new graphs, as opposed to the
	// default in the OpDef, which will be used when interpreting old
	// GraphDefs.
	DefaultValue *AttrValue `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Note: this will replace any inherited attr doc, there is no current
	// way of modifying attr descriptions as can be done with op descriptions.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef_Attr) Reset()         { *m = ApiDef_Attr{} }
func (m *ApiDef_Attr) String() string { return proto.CompactTextString(m) }
func (*ApiDef_Attr) ProtoMessage()    {}
func (*ApiDef_Attr) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a850add58b816a, []int{0, 2}
}

func (m *ApiDef_Attr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef_Attr.Unmarshal(m, b)
}
func (m *ApiDef_Attr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef_Attr.Marshal(b, m, deterministic)
}
func (m *ApiDef_Attr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef_Attr.Merge(m, src)
}
func (m *ApiDef_Attr) XXX_Size() int {
	return xxx_messageInfo_ApiDef_Attr.Size(m)
}
func (m *ApiDef_Attr) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef_Attr.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef_Attr proto.InternalMessageInfo

func (m *ApiDef_Attr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDef_Attr) GetRenameTo() string {
	if m != nil {
		return m.RenameTo
	}
	return ""
}

func (m *ApiDef_Attr) GetDefaultValue() *AttrValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *ApiDef_Attr) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ApiDefs struct {
	Op                   []*ApiDef `protobuf:"bytes,1,rep,name=op,proto3" json:"op,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ApiDefs) Reset()         { *m = ApiDefs{} }
func (m *ApiDefs) String() string { return proto.CompactTextString(m) }
func (*ApiDefs) ProtoMessage()    {}
func (*ApiDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a850add58b816a, []int{1}
}

func (m *ApiDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDefs.Unmarshal(m, b)
}
func (m *ApiDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDefs.Marshal(b, m, deterministic)
}
func (m *ApiDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDefs.Merge(m, src)
}
func (m *ApiDefs) XXX_Size() int {
	return xxx_messageInfo_ApiDefs.Size(m)
}
func (m *ApiDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDefs proto.InternalMessageInfo

func (m *ApiDefs) GetOp() []*ApiDef {
	if m != nil {
		return m.Op
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.ApiDef_Visibility", ApiDef_Visibility_name, ApiDef_Visibility_value)
	proto.RegisterType((*ApiDef)(nil), "tensorflow.ApiDef")
	proto.RegisterType((*ApiDef_Endpoint)(nil), "tensorflow.ApiDef.Endpoint")
	proto.RegisterType((*ApiDef_Arg)(nil), "tensorflow.ApiDef.Arg")
	proto.RegisterType((*ApiDef_Attr)(nil), "tensorflow.ApiDef.Attr")
	proto.RegisterType((*ApiDefs)(nil), "tensorflow.ApiDefs")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/api_def.proto", fileDescriptor_00a850add58b816a)
}

var fileDescriptor_00a850add58b816a = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x49, 0xd3, 0x35, 0xc9, 0x65, 0x43, 0xc5, 0x88, 0x61, 0x75, 0x02, 0x45, 0x7d, 0xa1,
	0x02, 0xad, 0x91, 0xc6, 0x03, 0x12, 0xd2, 0x90, 0x3a, 0xb5, 0x40, 0xc4, 0xd8, 0xaa, 0x6c, 0x4c,
	0xc0, 0x4b, 0xe4, 0x35, 0x4e, 0x66, 0xd1, 0xc4, 0x96, 0xe3, 0x6c, 0xec, 0x7b, 0xc0, 0x77, 0xe5,
	0x11, 0xc5, 0xd9, 0xd6, 0xd0, 0x56, 0x93, 0x78, 0xf3, 0xdd, 0xff, 0xf7, 0xbf, 0x38, 0xe7, 0xb3,
	0xe1, 0x85, 0xa2, 0x79, 0xc1, 0x65, 0x32, 0xe7, 0x57, 0xfe, 0x8c, 0x4b, 0xea, 0x27, 0x92, 0x64,
	0xf4, 0x8a, 0xcb, 0x1f, 0x3e, 0x11, 0x2c, 0x8a, 0x69, 0x32, 0x14, 0x92, 0x2b, 0x8e, 0x60, 0x01,
	0xf6, 0x5e, 0xde, 0x63, 0x52, 0x4a, 0x46, 0x97, 0x64, 0x5e, 0xd2, 0xda, 0xd7, 0xff, 0x65, 0x41,
	0x67, 0x24, 0xd8, 0x98, 0x26, 0xa8, 0x0f, 0x5b, 0xa9, 0x24, 0xe2, 0x22, 0xe2, 0x22, 0xca, 0x49,
	0x46, 0xb1, 0xe1, 0x19, 0x03, 0x27, 0x74, 0x75, 0xf2, 0x58, 0x1c, 0x91, 0x8c, 0x22, 0x1f, 0x1e,
	0xc7, 0x54, 0x48, 0x3a, 0x23, 0x8a, 0xf1, 0x3c, 0xca, 0x68, 0x51, 0x90, 0x94, 0xe2, 0x4d, 0x4d,
	0xa2, 0x86, 0xf4, 0xb9, 0x56, 0xd0, 0x3e, 0xc0, 0x25, 0x2b, 0xd8, 0x39, 0x9b, 0x33, 0x75, 0x8d,
	0x5b, 0x9e, 0x31, 0x78, 0xb8, 0xf7, 0x6c, 0xb8, 0xd8, 0xe0, 0xb0, 0xfe, 0xf8, 0xf0, 0xec, 0x0e,
	0x0a, 0x1b, 0x06, 0xf4, 0x06, 0x6c, 0x9a, 0xc7, 0x82, 0xb3, 0x5c, 0x61, 0xd3, 0x33, 0x07, 0xee,
	0xde, 0xce, 0x1a, 0xf3, 0xe4, 0x06, 0x09, 0xef, 0x60, 0xb4, 0x0b, 0x1d, 0x96, 0x47, 0x44, 0xa6,
	0xb8, 0xad, 0x6d, 0xdb, 0x6b, 0x6c, 0x23, 0x99, 0x86, 0x1b, 0x2c, 0x1f, 0xc9, 0x14, 0xf9, 0x60,
	0xf1, 0x52, 0x69, 0x7e, 0xe3, 0x5e, 0xbe, 0xc3, 0x4b, 0x55, 0x19, 0x76, 0xc0, 0x21, 0x32, 0x8d,
	0xb8, 0x8c, 0xa9, 0xc4, 0xae, 0x67, 0x0e, 0x9c, 0xd0, 0x26, 0x32, 0x3d, 0xae, 0x62, 0xf4, 0x0a,
	0xda, 0x55, 0xa3, 0x71, 0x47, 0x97, 0x7a, 0xba, 0xae, 0x94, 0x52, 0x32, 0xd4, 0x10, 0xc2, 0x60,
	0x15, 0x65, 0x96, 0x11, 0x79, 0x8d, 0x2d, 0xdd, 0xc6, 0xdb, 0x10, 0x79, 0xe0, 0xc6, 0xb4, 0x98,
	0x49, 0x26, 0xaa, 0x8e, 0x62, 0xbb, 0x3e, 0x8e, 0x46, 0x0a, 0xed, 0x02, 0x6a, 0x84, 0x91, 0x90,
	0x34, 0x61, 0x3f, 0xb1, 0xa3, 0xc1, 0x47, 0x0d, 0x65, 0xaa, 0x85, 0x65, 0xbc, 0x28, 0x93, 0x0a,
	0x87, 0x15, 0xfc, 0x44, 0x0b, 0xbd, 0x77, 0x60, 0xdf, 0x76, 0x16, 0x21, 0x68, 0x37, 0x66, 0x42,
	0xaf, 0xd1, 0x73, 0x80, 0xdb, 0x13, 0xa7, 0x31, 0x36, 0x3d, 0x63, 0x60, 0x87, 0x8d, 0x4c, 0xef,
	0x2b, 0x98, 0x55, 0xab, 0xd6, 0x59, 0x77, 0xc0, 0x91, 0xb4, 0x5a, 0x45, 0x8a, 0xeb, 0xa9, 0x70,
	0x42, 0xbb, 0x4e, 0x9c, 0xf2, 0xe5, 0xff, 0x36, 0x57, 0xfe, 0xbb, 0xf7, 0xdb, 0x80, 0x76, 0xd5,
	0xc2, 0xff, 0xaf, 0xfd, 0x16, 0xb6, 0x62, 0x9a, 0x90, 0x72, 0xae, 0xea, 0x6b, 0xa0, 0xab, 0xbb,
	0x7b, 0x4f, 0xfe, 0x39, 0x23, 0xa5, 0xe4, 0x59, 0x25, 0x86, 0x9b, 0x37, 0xac, 0x8e, 0x96, 0xf7,
	0xd5, 0x5e, 0xd9, 0x57, 0xff, 0x03, 0xc0, 0x62, 0x90, 0xd1, 0x36, 0xa0, 0xf1, 0xe4, 0xfd, 0xe8,
	0xcb, 0xe1, 0x69, 0x74, 0x16, 0x9c, 0x04, 0x07, 0xc1, 0x61, 0x70, 0xfa, 0xad, 0xfb, 0x00, 0xb9,
	0x60, 0xe9, 0xf8, 0x70, 0xd2, 0x35, 0x90, 0x0d, 0xed, 0x93, 0x4f, 0xc1, 0xb4, 0xdb, 0x42, 0x00,
	0x9d, 0x8f, 0xc1, 0x78, 0x3c, 0x39, 0xea, 0x9a, 0xfd, 0x5d, 0xb0, 0xea, 0x49, 0x29, 0x50, 0x1f,
	0x5a, 0x5c, 0x60, 0x43, 0x8f, 0x12, 0x5a, 0x1d, 0xa5, 0xb0, 0xc5, 0xc5, 0xc1, 0x1c, 0x30, 0x97,
	0x69, 0x53, 0xbc, 0xbb, 0xf2, 0x07, 0x9b, 0x35, 0x37, 0xad, 0xae, 0x7b, 0x31, 0x35, 0xbe, 0xef,
	0xa7, 0x4c, 0x5d, 0x94, 0xe7, 0xc3, 0x19, 0xcf, 0xfc, 0xc6, 0x43, 0xb1, 0x7e, 0x99, 0xf2, 0xa5,
	0x17, 0xe4, 0x8f, 0x61, 0x9c, 0x77, 0xf4, 0xd3, 0xf1, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4b, 0x71, 0x42, 0xdb, 0x9d, 0x04, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/listener/original_dst/v3/original_dst.proto

package envoy_extensions_filters_listener_original_dst_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type OriginalDst struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalDst) Reset()         { *m = OriginalDst{} }
func (m *OriginalDst) String() string { return proto.CompactTextString(m) }
func (*OriginalDst) ProtoMessage()    {}
func (*OriginalDst) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f08b98c629d032, []int{0}
}

func (m *OriginalDst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalDst.Unmarshal(m, b)
}
func (m *OriginalDst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalDst.Marshal(b, m, deterministic)
}
func (m *OriginalDst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalDst.Merge(m, src)
}
func (m *OriginalDst) XXX_Size() int {
	return xxx_messageInfo_OriginalDst.Size(m)
}
func (m *OriginalDst) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalDst.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalDst proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OriginalDst)(nil), "envoy.extensions.filters.listener.original_dst.v3.OriginalDst")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/listener/original_dst/v3/original_dst.proto", fileDescriptor_a4f08b98c629d032)
}

var fileDescriptor_a4f08b98c629d032 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0xcf, 0x2f,
	0xca, 0x4c, 0xcf, 0xcc, 0x4b, 0xcc, 0x89, 0x4f, 0x29, 0x2e, 0xd1, 0x2f, 0x33, 0x46, 0xe1, 0xeb,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x19, 0x82, 0x4d, 0xd1, 0x43, 0x98, 0xa2, 0x07, 0x35, 0x45,
	0x0f, 0x66, 0x8a, 0x1e, 0x8a, 0xae, 0x32, 0x63, 0x29, 0xc5, 0xd2, 0x94, 0x82, 0x44, 0xfd, 0xc4,
	0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x12, 0xb0, 0xc5, 0x65, 0xa9, 0x45, 0x20, 0xbd, 0x99, 0x79, 0xe9,
	0x10, 0x53, 0x95, 0xfc, 0xb8, 0xb8, 0xfd, 0xa1, 0xba, 0x5c, 0x8a, 0x4b, 0xac, 0xec, 0x67, 0x1d,
	0xed, 0x90, 0xb3, 0xe2, 0xb2, 0x80, 0xd8, 0x95, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x0e, 0xb5, 0x07,
	0x97, 0x35, 0x46, 0x7a, 0x48, 0x06, 0x38, 0x85, 0x72, 0xd9, 0x67, 0xe6, 0xeb, 0x81, 0xb5, 0x17,
	0x14, 0xe5, 0x57, 0x54, 0xea, 0x91, 0xec, 0x6a, 0x27, 0x01, 0x24, 0xf3, 0x02, 0x40, 0x8e, 0x0c,
	0x60, 0x4c, 0x62, 0x03, 0xbb, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x57, 0xf9, 0x58, 0x0a,
	0x4b, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/element.proto

/*
Package crypto is a generated protocol buffer package.

It is generated from these files:
	crypto/element.proto

It has these top-level messages:
	Point
	Scalar
*/
package crypto

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

//
// GroupID is an enumeration holding all possible groups that can be marshalled
// / unmarshalled, supported by the kyber library
type GroupID int32

const (
	GroupID_EDWARDS25519 GroupID = 0
	GroupID_BN256_G1     GroupID = 11
	GroupID_BN256_G2     GroupID = 12
	GroupID_BN256_GT     GroupID = 13
)

var GroupID_name = map[int32]string{
	0:  "EDWARDS25519",
	11: "BN256_G1",
	12: "BN256_G2",
	13: "BN256_GT",
}
var GroupID_value = map[string]int32{
	"EDWARDS25519": 0,
	"BN256_G1":     11,
	"BN256_G2":     12,
	"BN256_GT":     13,
}

func (x GroupID) String() string {
	return proto.EnumName(GroupID_name, int32(x))
}
func (GroupID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

//
// Point represents a point on a curve,i.e. a public key, a commitment etc
// It is parametrized by its group.
type Point struct {
	Gid  GroupID `protobuf:"varint,1,opt,name=gid,enum=element.GroupID" json:"gid,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetGid() GroupID {
	if m != nil {
		return m.Gid
	}
	return GroupID_EDWARDS25519
}

func (m *Point) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

//
// Scalar represents a scalar on the field attached to the group. It is
// parametrized by the group using this field (1-1 mapping).
type Scalar struct {
	Gid  GroupID `protobuf:"varint,1,opt,name=gid,enum=element.GroupID" json:"gid,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Scalar) Reset()                    { *m = Scalar{} }
func (m *Scalar) String() string            { return proto.CompactTextString(m) }
func (*Scalar) ProtoMessage()               {}
func (*Scalar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Scalar) GetGid() GroupID {
	if m != nil {
		return m.Gid
	}
	return GroupID_EDWARDS25519
}

func (m *Scalar) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "element.Point")
	proto.RegisterType((*Scalar)(nil), "element.Scalar")
	proto.RegisterEnum("element.GroupID", GroupID_name, GroupID_value)
}

func init() { proto.RegisterFile("crypto/element.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0xd7, 0x4f, 0xcd, 0x49, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x72, 0x95, 0xec, 0xb9, 0x58, 0x03, 0xf2, 0x33, 0xf3, 0x4a, 0x84, 0x94, 0xb8,
	0x98, 0xd3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x04, 0xf4, 0x60, 0xca, 0xdd,
	0x8b, 0xf2, 0x4b, 0x0b, 0x3c, 0x5d, 0x82, 0x40, 0x92, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25,
	0x89, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x92, 0x03, 0x17, 0x5b, 0x70, 0x72,
	0x62, 0x4e, 0x62, 0x11, 0xb9, 0x26, 0x68, 0xb9, 0x72, 0xb1, 0x43, 0xd5, 0x08, 0x09, 0x70, 0xf1,
	0xb8, 0xba, 0x84, 0x3b, 0x06, 0xb9, 0x04, 0x1b, 0x99, 0x9a, 0x1a, 0x5a, 0x0a, 0x30, 0x08, 0xf1,
	0x70, 0x71, 0x38, 0xf9, 0x19, 0x99, 0x9a, 0xc5, 0xbb, 0x1b, 0x0a, 0x70, 0x23, 0xf1, 0x8c, 0x04,
	0x78, 0x90, 0x78, 0x21, 0x02, 0xbc, 0x4e, 0x1a, 0x51, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x29, 0xa9, 0x29, 0x99, 0xc5, 0xfa, 0x29, 0x45, 0x89, 0x79, 0x29,
	0xfa, 0x60, 0x2f, 0x27, 0x95, 0xa6, 0xe9, 0x43, 0x82, 0x22, 0x89, 0x0d, 0x2c, 0x60, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0x6b, 0x0f, 0x92, 0x1b, 0x01, 0x00, 0x00,
}
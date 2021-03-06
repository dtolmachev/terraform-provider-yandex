// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/key.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Key_Algorithm int32

const (
	Key_ALGORITHM_UNSPECIFIED Key_Algorithm = 0
	// RSA with a 2048-bit key size. Default value.
	Key_RSA_2048 Key_Algorithm = 1
	// RSA with a 4096-bit key size.
	Key_RSA_4096 Key_Algorithm = 2
)

var Key_Algorithm_name = map[int32]string{
	0: "ALGORITHM_UNSPECIFIED",
	1: "RSA_2048",
	2: "RSA_4096",
}

var Key_Algorithm_value = map[string]int32{
	"ALGORITHM_UNSPECIFIED": 0,
	"RSA_2048":              1,
	"RSA_4096":              2,
}

func (x Key_Algorithm) String() string {
	return proto.EnumName(Key_Algorithm_name, int32(x))
}

func (Key_Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3c4378eab1afe9e, []int{0, 0}
}

// A Key resource. For more information, see [Authorized keys](/docs/iam/concepts/authorization/key).
type Key struct {
	// ID of the Key resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Subject:
	//	*Key_UserAccountId
	//	*Key_ServiceAccountId
	Subject isKey_Subject `protobuf_oneof:"subject"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Description of the Key resource. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// An algorithm used to generate a key pair of the Key resource.
	KeyAlgorithm Key_Algorithm `protobuf:"varint,6,opt,name=key_algorithm,json=keyAlgorithm,proto3,enum=yandex.cloud.iam.v1.Key_Algorithm" json:"key_algorithm,omitempty"`
	// A public key of the Key resource.
	PublicKey            string   `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c4378eab1afe9e, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isKey_Subject interface {
	isKey_Subject()
}

type Key_UserAccountId struct {
	UserAccountId string `protobuf:"bytes,2,opt,name=user_account_id,json=userAccountId,proto3,oneof"`
}

type Key_ServiceAccountId struct {
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3,oneof"`
}

func (*Key_UserAccountId) isKey_Subject() {}

func (*Key_ServiceAccountId) isKey_Subject() {}

func (m *Key) GetSubject() isKey_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Key) GetUserAccountId() string {
	if x, ok := m.GetSubject().(*Key_UserAccountId); ok {
		return x.UserAccountId
	}
	return ""
}

func (m *Key) GetServiceAccountId() string {
	if x, ok := m.GetSubject().(*Key_ServiceAccountId); ok {
		return x.ServiceAccountId
	}
	return ""
}

func (m *Key) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Key) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Key) GetKeyAlgorithm() Key_Algorithm {
	if m != nil {
		return m.KeyAlgorithm
	}
	return Key_ALGORITHM_UNSPECIFIED
}

func (m *Key) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Key) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Key_UserAccountId)(nil),
		(*Key_ServiceAccountId)(nil),
	}
}

func init() {
	proto.RegisterEnum("yandex.cloud.iam.v1.Key_Algorithm", Key_Algorithm_name, Key_Algorithm_value)
	proto.RegisterType((*Key)(nil), "yandex.cloud.iam.v1.Key")
}

func init() { proto.RegisterFile("yandex/cloud/iam/v1/key.proto", fileDescriptor_d3c4378eab1afe9e) }

var fileDescriptor_d3c4378eab1afe9e = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x6f, 0xd4, 0x30,
	0x10, 0xc6, 0x9b, 0x5d, 0x68, 0xc9, 0xf4, 0xc1, 0xca, 0x08, 0x11, 0x2a, 0x55, 0xac, 0xf6, 0xb4,
	0x97, 0xda, 0xed, 0x52, 0x21, 0xaa, 0x9e, 0xb2, 0x50, 0xda, 0x68, 0x79, 0x29, 0x2d, 0x1c, 0xb8,
	0x44, 0x8e, 0x3d, 0xa4, 0x26, 0x0f, 0x47, 0x89, 0xb3, 0xc2, 0x67, 0xfe, 0x71, 0xd4, 0x3c, 0x68,
	0x91, 0x7a, 0x9c, 0xef, 0xfb, 0x59, 0xfe, 0x69, 0x34, 0x70, 0x60, 0x79, 0x21, 0xf1, 0x37, 0x13,
	0x99, 0x6e, 0x24, 0x53, 0x3c, 0x67, 0xeb, 0x63, 0x96, 0xa2, 0xa5, 0x65, 0xa5, 0x8d, 0x26, 0xcf,
	0xba, 0x9a, 0xb6, 0x35, 0x55, 0x3c, 0xa7, 0xeb, 0xe3, 0xfd, 0x57, 0x89, 0xd6, 0x49, 0x86, 0xac,
	0x45, 0xe2, 0xe6, 0x27, 0x33, 0x2a, 0xc7, 0xda, 0xf0, 0xbc, 0xec, 0x5e, 0xcd, 0xfe, 0x8c, 0x61,
	0xbc, 0x42, 0x4b, 0xf6, 0x60, 0xa4, 0xa4, 0xe7, 0x4c, 0x9d, 0xb9, 0x1b, 0x8e, 0x94, 0x24, 0x73,
	0x78, 0xda, 0xd4, 0x58, 0x45, 0x5c, 0x08, 0xdd, 0x14, 0x26, 0x52, 0xd2, 0x1b, 0xdd, 0x96, 0x97,
	0x1b, 0xe1, 0xee, 0x6d, 0xe1, 0x77, 0x79, 0x20, 0x09, 0x05, 0x52, 0x63, 0xb5, 0x56, 0x02, 0xef,
	0xc3, 0xe3, 0x1e, 0x9e, 0xf4, 0xdd, 0x1d, 0x7f, 0x0a, 0x20, 0x2a, 0xe4, 0x06, 0x65, 0xc4, 0x8d,
	0xf7, 0x68, 0xea, 0xcc, 0xb7, 0x17, 0xfb, 0xb4, 0xf3, 0xa4, 0x83, 0x27, 0xbd, 0x1e, 0x3c, 0x43,
	0xb7, 0xa7, 0x7d, 0x43, 0xa6, 0xb0, 0x2d, 0xb1, 0x16, 0x95, 0x2a, 0x8d, 0xd2, 0x85, 0xf7, 0xb8,
	0xb5, 0xbd, 0x1f, 0x91, 0x0b, 0xd8, 0x4d, 0xd1, 0x46, 0x3c, 0x4b, 0x74, 0xa5, 0xcc, 0x4d, 0xee,
	0x6d, 0x4e, 0x9d, 0xf9, 0xde, 0x62, 0x46, 0x1f, 0x58, 0x0e, 0x5d, 0xa1, 0xa5, 0xfe, 0x40, 0x86,
	0x3b, 0x29, 0xda, 0x7f, 0x13, 0x39, 0x00, 0x28, 0x9b, 0x38, 0x53, 0x22, 0x4a, 0xd1, 0x7a, 0x5b,
	0xed, 0x4f, 0x6e, 0x97, 0xac, 0xd0, 0xce, 0x96, 0xe0, 0xde, 0xb1, 0x2f, 0xe1, 0xb9, 0xff, 0xf1,
	0xe2, 0x4b, 0x18, 0x5c, 0x5f, 0x7e, 0x8a, 0xbe, 0x7d, 0xbe, 0xfa, 0x7a, 0xfe, 0x2e, 0xf8, 0x10,
	0x9c, 0xbf, 0x9f, 0x6c, 0x90, 0x1d, 0x78, 0x12, 0x5e, 0xf9, 0xd1, 0xe2, 0xe8, 0xe4, 0xed, 0xc4,
	0x19, 0xa6, 0x93, 0xa3, 0xd3, 0x37, 0x93, 0xd1, 0xd2, 0x85, 0xad, 0xba, 0x89, 0x7f, 0xa1, 0x30,
	0xcb, 0xef, 0xf0, 0xe2, 0x3f, 0x41, 0x5e, 0xaa, 0x5e, 0xf2, 0xc7, 0x59, 0xa2, 0xcc, 0x4d, 0x13,
	0x53, 0xa1, 0x73, 0xd6, 0x31, 0x87, 0xdd, 0x01, 0x24, 0xfa, 0x30, 0xc1, 0xa2, 0x5d, 0x18, 0x7b,
	0xe0, 0x32, 0xce, 0x14, 0xcf, 0xe3, 0xcd, 0xb6, 0x7e, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x35,
	0xf9, 0x50, 0xd4, 0x3b, 0x02, 0x00, 0x00,
}

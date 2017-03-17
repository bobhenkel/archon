// Code generated by protoc-gen-gogo.
// source: pb/money/money.proto
// DO NOT EDIT!

package pbmoney

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Currency struct {
	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code        CurrencyCode `protobuf:"varint,2,opt,name=code,proto3,enum=pb.money.CurrencyCode" json:"code,omitempty"`
	NumericCode uint32       `protobuf:"varint,3,opt,name=numeric_code,json=numericCode,proto3" json:"numeric_code,omitempty"`
	MinorUnit   uint32       `protobuf:"varint,4,opt,name=minor_unit,json=minorUnit,proto3" json:"minor_unit,omitempty"`
}

func (m *Currency) Reset()                    { *m = Currency{} }
func (m *Currency) String() string            { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()               {}
func (*Currency) Descriptor() ([]byte, []int) { return fileDescriptorMoney, []int{0} }

type Money struct {
	CurrencyCode CurrencyCode `protobuf:"varint,1,opt,name=currency_code,json=currencyCode,proto3,enum=pb.money.CurrencyCode" json:"currency_code,omitempty"`
	ValueMicros  int64        `protobuf:"varint,2,opt,name=value_micros,json=valueMicros,proto3" json:"value_micros,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptorMoney, []int{1} }

type ExchangeRate struct {
	From        CurrencyCode `protobuf:"varint,1,opt,name=from,proto3,enum=pb.money.CurrencyCode" json:"from,omitempty"`
	To          CurrencyCode `protobuf:"varint,2,opt,name=to,proto3,enum=pb.money.CurrencyCode" json:"to,omitempty"`
	ValueMicros int64        `protobuf:"varint,3,opt,name=value_micros,json=valueMicros,proto3" json:"value_micros,omitempty"`
}

func (m *ExchangeRate) Reset()                    { *m = ExchangeRate{} }
func (m *ExchangeRate) String() string            { return proto.CompactTextString(m) }
func (*ExchangeRate) ProtoMessage()               {}
func (*ExchangeRate) Descriptor() ([]byte, []int) { return fileDescriptorMoney, []int{2} }

func init() {
	proto.RegisterType((*Currency)(nil), "pb.money.Currency")
	proto.RegisterType((*Money)(nil), "pb.money.Money")
	proto.RegisterType((*ExchangeRate)(nil), "pb.money.ExchangeRate")
}

var fileDescriptorMoney = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xc9, 0xb6, 0xea, 0x76, 0xb6, 0xeb, 0x21, 0x88, 0x14, 0x41, 0xd0, 0x1e, 0x44, 0x3c,
	0x54, 0xd0, 0xa3, 0x37, 0x17, 0x8f, 0x7b, 0x09, 0x78, 0xf1, 0x52, 0xda, 0x18, 0x6b, 0xc1, 0xcc,
	0x94, 0xd8, 0x88, 0xfb, 0x03, 0xbc, 0xf9, 0xa3, 0x4d, 0xa7, 0xbb, 0x20, 0x2e, 0xb8, 0x97, 0xd2,
	0x79, 0xf3, 0xbd, 0x99, 0x97, 0x04, 0x8e, 0xba, 0xfa, 0xda, 0x12, 0x9a, 0xd5, 0xf8, 0x2d, 0x3a,
	0x47, 0x3d, 0xc9, 0x69, 0x57, 0x17, 0x5c, 0x9f, 0x64, 0x7f, 0xfa, 0x8d, 0xc1, 0x91, 0xc9, 0xbf,
	0x05, 0x4c, 0x17, 0xde, 0x39, 0x83, 0x7a, 0x25, 0x25, 0xc4, 0x58, 0x59, 0x93, 0x89, 0x33, 0x71,
	0x99, 0x28, 0xfe, 0x97, 0x57, 0x10, 0x6b, 0x7a, 0x36, 0xd9, 0x24, 0x68, 0x87, 0x37, 0xc7, 0xc5,
	0x66, 0x66, 0xb1, 0x71, 0x2d, 0x42, 0x57, 0x31, 0x23, 0xcf, 0x21, 0x45, 0x6f, 0x8d, 0x6b, 0x75,
	0xc9, 0x9e, 0x28, 0x78, 0xe6, 0x6a, 0xb6, 0xd6, 0x06, 0x50, 0x9e, 0x02, 0xd8, 0x16, 0xc9, 0x95,
	0x1e, 0xdb, 0x3e, 0x8b, 0x19, 0x48, 0x58, 0x79, 0x0c, 0x42, 0xde, 0xc0, 0xde, 0x72, 0x98, 0x2e,
	0xef, 0x60, 0xae, 0xd7, 0x0b, 0xc6, 0x59, 0xe2, 0xdf, 0xfd, 0xa9, 0xfe, 0x55, 0x0d, 0x39, 0x3e,
	0xaa, 0x37, 0x6f, 0x4a, 0xdb, 0x6a, 0x47, 0xef, 0x9c, 0x3d, 0x52, 0x33, 0xd6, 0x96, 0x2c, 0xe5,
	0x5f, 0x02, 0xd2, 0x87, 0x4f, 0xfd, 0x5a, 0x61, 0x63, 0x54, 0xd5, 0xf3, 0x39, 0x5f, 0x1c, 0xd9,
	0x1d, 0x7b, 0x98, 0x91, 0x17, 0x30, 0xe9, 0x69, 0xc7, 0x8d, 0x04, 0x62, 0x2b, 0x47, 0xb4, 0x95,
	0xe3, 0x3e, 0x79, 0x3a, 0xe8, 0x6a, 0xb6, 0xd7, 0xfb, 0xfc, 0x22, 0xb7, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0xdc, 0x6f, 0x56, 0xcd, 0x01, 0x00, 0x00,
}

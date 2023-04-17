// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type OpType int32

const (
	OpType_INSERT OpType = 0
	OpType_DELETE OpType = 1
)

var OpType_name = map[int32]string{
	0: "INSERT",
	1: "DELETE",
}

var OpType_value = map[string]int32{
	"INSERT": 0,
	"DELETE": 1,
}

func (x OpType) String() string {
	return proto.EnumName(OpType_name, int32(x))
}

func (OpType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type Operation struct {
	Type                 OpType   `protobuf:"varint,1,opt,name=type,proto3,enum=api_pb.OpType" json:"type,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Length               int32    `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Text                 []byte   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Revision             int32    `protobuf:"varint,5,opt,name=revision,proto3" json:"revision,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *Operation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return m.Size()
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetType() OpType {
	if m != nil {
		return m.Type
	}
	return OpType_INSERT
}

func (m *Operation) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Operation) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Operation) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Operation) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *Operation) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OperationAck struct {
	Revision             int32    `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationAck) Reset()         { *m = OperationAck{} }
func (m *OperationAck) String() string { return proto.CompactTextString(m) }
func (*OperationAck) ProtoMessage()    {}
func (*OperationAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *OperationAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OperationAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OperationAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OperationAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationAck.Merge(m, src)
}
func (m *OperationAck) XXX_Size() int {
	return m.Size()
}
func (m *OperationAck) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationAck.DiscardUnknown(m)
}

var xxx_messageInfo_OperationAck proto.InternalMessageInfo

func (m *OperationAck) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type DocumentUpdate struct {
	Text                 []byte   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	LastRevision         int32    `protobuf:"varint,2,opt,name=last_revision,json=lastRevision,proto3" json:"last_revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentUpdate) Reset()         { *m = DocumentUpdate{} }
func (m *DocumentUpdate) String() string { return proto.CompactTextString(m) }
func (*DocumentUpdate) ProtoMessage()    {}
func (*DocumentUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}
func (m *DocumentUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DocumentUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DocumentUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DocumentUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentUpdate.Merge(m, src)
}
func (m *DocumentUpdate) XXX_Size() int {
	return m.Size()
}
func (m *DocumentUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentUpdate proto.InternalMessageInfo

func (m *DocumentUpdate) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *DocumentUpdate) GetLastRevision() int32 {
	if m != nil {
		return m.LastRevision
	}
	return 0
}

func init() {
	proto.RegisterEnum("api_pb.OpType", OpType_name, OpType_value)
	proto.RegisterType((*Operation)(nil), "api_pb.Operation")
	proto.RegisterType((*OperationAck)(nil), "api_pb.OperationAck")
	proto.RegisterType((*DocumentUpdate)(nil), "api_pb.DocumentUpdate")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x3b, 0xb6, 0x0d, 0x76, 0xa8, 0xa5, 0x2c, 0x22, 0x41, 0x24, 0x84, 0x78, 0x09, 0x85,
	0xa6, 0xa0, 0x37, 0x6f, 0x4a, 0x73, 0x28, 0x88, 0x85, 0x58, 0x2f, 0x5e, 0xca, 0x36, 0x5d, 0xdb,
	0xc5, 0x26, 0xbb, 0x74, 0x27, 0xd2, 0xbe, 0x89, 0xef, 0xe0, 0x8b, 0x78, 0xf4, 0x11, 0x24, 0xbe,
	0x88, 0x34, 0x29, 0xd1, 0xde, 0xe6, 0xff, 0x76, 0xf8, 0xf9, 0xf7, 0x1f, 0x6c, 0x71, 0x2d, 0x03,
	0xbd, 0x56, 0xa4, 0x98, 0xc5, 0xb5, 0x9c, 0xea, 0x99, 0xf7, 0x01, 0xd8, 0x1a, 0x6b, 0xb1, 0xe6,
	0x24, 0x55, 0xca, 0x3c, 0x6c, 0xd0, 0x56, 0x0b, 0x1b, 0x5c, 0xf0, 0x3b, 0x57, 0x9d, 0xa0, 0x5c,
	0x0a, 0xc6, 0x7a, 0xb2, 0xd5, 0x22, 0x2a, 0xde, 0xd8, 0x29, 0x36, 0x65, 0x3a, 0x17, 0x1b, 0xfb,
	0xc8, 0x05, 0xbf, 0x19, 0x95, 0x82, 0x9d, 0xa1, 0xb5, 0x12, 0xe9, 0x82, 0x96, 0x76, 0xbd, 0xc0,
	0x7b, 0xc5, 0x18, 0x36, 0x48, 0x6c, 0xc8, 0x6e, 0xb8, 0xe0, 0xb7, 0xa3, 0x62, 0x66, 0xe7, 0x78,
	0xbc, 0x16, 0x6f, 0xd2, 0x48, 0x95, 0xda, 0xcd, 0x62, 0xbb, 0xd2, 0xec, 0x02, 0x5b, 0x24, 0x13,
	0x61, 0x88, 0x27, 0xda, 0xb6, 0x5c, 0xf0, 0xeb, 0xd1, 0x1f, 0xf0, 0x7a, 0xd8, 0xae, 0xc2, 0xde,
	0xc6, 0xaf, 0x07, 0x4e, 0x70, 0xe8, 0xe4, 0x8d, 0xb0, 0x33, 0x54, 0x71, 0x96, 0x88, 0x94, 0x9e,
	0xf4, 0x9c, 0x93, 0xa8, 0xb2, 0xc0, 0xbf, 0x2c, 0x97, 0x78, 0xb2, 0xe2, 0x86, 0xa6, 0x95, 0x4d,
	0xf9, 0xab, 0xf6, 0x0e, 0x46, 0x7b, 0xd6, 0x73, 0xd1, 0x2a, 0x2b, 0x60, 0x88, 0xd6, 0xe8, 0xe1,
	0x31, 0x8c, 0x26, 0xdd, 0xda, 0x6e, 0x1e, 0x86, 0xf7, 0xe1, 0x24, 0xec, 0xc2, 0xdd, 0xcd, 0x67,
	0xee, 0xc0, 0x57, 0xee, 0xc0, 0x77, 0xee, 0xc0, 0xfb, 0x8f, 0x53, 0x7b, 0xf6, 0x17, 0x92, 0x96,
	0xd9, 0x2c, 0x88, 0x55, 0x32, 0x30, 0x86, 0x67, 0xfd, 0x17, 0x29, 0x69, 0x10, 0xaf, 0x54, 0x36,
	0x57, 0xb1, 0xe9, 0x73, 0x2d, 0x07, 0x65, 0xbb, 0x33, 0xab, 0xb8, 0xc8, 0xf5, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x80, 0x6b, 0x56, 0x9e, 0x01, 0x00, 0x00,
}

func (m *Operation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Operation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Operation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.Revision != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x22
	}
	if m.Length != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x18
	}
	if m.Index != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OperationAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OperationAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OperationAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Revision != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DocumentUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DocumentUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DocumentUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LastRevision != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.LastRevision))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Operation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovApi(uint64(m.Type))
	}
	if m.Index != 0 {
		n += 1 + sovApi(uint64(m.Index))
	}
	if m.Length != 0 {
		n += 1 + sovApi(uint64(m.Length))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Revision != 0 {
		n += 1 + sovApi(uint64(m.Revision))
	}
	if m.Timestamp != 0 {
		n += 1 + sovApi(uint64(m.Timestamp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OperationAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Revision != 0 {
		n += 1 + sovApi(uint64(m.Revision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DocumentUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.LastRevision != 0 {
		n += 1 + sovApi(uint64(m.LastRevision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Operation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Operation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Operation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= OpType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = append(m.Text[:0], dAtA[iNdEx:postIndex]...)
			if m.Text == nil {
				m.Text = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OperationAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OperationAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OperationAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DocumentUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DocumentUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DocumentUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = append(m.Text[:0], dAtA[iNdEx:postIndex]...)
			if m.Text == nil {
				m.Text = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRevision", wireType)
			}
			m.LastRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastRevision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
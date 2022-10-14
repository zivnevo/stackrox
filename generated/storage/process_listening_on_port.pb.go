// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/process_listening_on_port.proto

package storage

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
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

type ProcessListeningOnPort struct {
	Port                 uint32                     `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Protocol             L4Protocol                 `protobuf:"varint,2,opt,name=protocol,proto3,enum=storage.L4Protocol" json:"protocol,omitempty"`
	Process              *ProcessIndicatorUniqueKey `protobuf:"bytes,3,opt,name=process,proto3" json:"process,omitempty"`
	CloseTimestamp       *types.Timestamp           `protobuf:"bytes,4,opt,name=close_timestamp,json=closeTimestamp,proto3" json:"close_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ProcessListeningOnPort) Reset()         { *m = ProcessListeningOnPort{} }
func (m *ProcessListeningOnPort) String() string { return proto.CompactTextString(m) }
func (*ProcessListeningOnPort) ProtoMessage()    {}
func (*ProcessListeningOnPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_44bd1925a567394f, []int{0}
}
func (m *ProcessListeningOnPort) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessListeningOnPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessListeningOnPort.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessListeningOnPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessListeningOnPort.Merge(m, src)
}
func (m *ProcessListeningOnPort) XXX_Size() int {
	return m.Size()
}
func (m *ProcessListeningOnPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessListeningOnPort.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessListeningOnPort proto.InternalMessageInfo

func (m *ProcessListeningOnPort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ProcessListeningOnPort) GetProtocol() L4Protocol {
	if m != nil {
		return m.Protocol
	}
	return L4Protocol_L4_PROTOCOL_UNKNOWN
}

func (m *ProcessListeningOnPort) GetProcess() *ProcessIndicatorUniqueKey {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *ProcessListeningOnPort) GetCloseTimestamp() *types.Timestamp {
	if m != nil {
		return m.CloseTimestamp
	}
	return nil
}

func (m *ProcessListeningOnPort) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ProcessListeningOnPort) Clone() *ProcessListeningOnPort {
	if m == nil {
		return nil
	}
	cloned := new(ProcessListeningOnPort)
	*cloned = *m

	cloned.Process = m.Process.Clone()
	cloned.CloseTimestamp = m.CloseTimestamp.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*ProcessListeningOnPort)(nil), "storage.ProcessListeningOnPort")
}

func init() {
	proto.RegisterFile("storage/process_listening_on_port.proto", fileDescriptor_44bd1925a567394f)
}

var fileDescriptor_44bd1925a567394f = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x7f, 0xf9, 0x39, 0x9c, 0x44, 0x9c, 0x10, 0x41, 0x4a, 0x91, 0xae, 0xec, 0x62, 0x4f,
	0x29, 0x4c, 0x8f, 0x9e, 0xf4, 0x24, 0x0e, 0x2c, 0x45, 0x2f, 0x5e, 0x4a, 0x57, 0xb3, 0x12, 0xec,
	0xf2, 0xd4, 0xe4, 0x19, 0xc3, 0x77, 0xe2, 0x4b, 0xf2, 0xe8, 0x4b, 0x18, 0xf5, 0x4d, 0x78, 0x14,
	0xd3, 0xa4, 0x07, 0x6f, 0x2d, 0xcf, 0xe7, 0xfb, 0x27, 0x5f, 0x7a, 0x6e, 0x10, 0x74, 0x59, 0x8b,
	0xb4, 0xd5, 0x50, 0x09, 0x63, 0x8a, 0x46, 0x1a, 0x14, 0x4a, 0xaa, 0xba, 0x00, 0x55, 0xb4, 0xa0,
	0x91, 0xb7, 0x1a, 0x10, 0xd8, 0xd8, 0x81, 0xe1, 0xb4, 0x06, 0xa8, 0x1b, 0x2b, 0x40, 0x58, 0x6e,
	0x56, 0x29, 0xca, 0xb5, 0x30, 0x58, 0xae, 0xdb, 0x9e, 0x0c, 0x43, 0x6f, 0xa9, 0x04, 0x6e, 0x41,
	0xbf, 0x14, 0xab, 0x06, 0xb6, 0xee, 0x36, 0xfd, 0x1b, 0x27, 0xd5, 0xb3, 0xac, 0x4a, 0x04, 0xdd,
	0x03, 0xb3, 0x1d, 0xa1, 0xa7, 0x59, 0x7f, 0x5b, 0xf8, 0x26, 0xf7, 0x2a, 0x03, 0x8d, 0x8c, 0xd1,
	0xd1, 0x6f, 0x9f, 0x80, 0xc4, 0x24, 0x39, 0xca, 0xed, 0x37, 0x4b, 0xe9, 0x81, 0xd5, 0x55, 0xd0,
	0x04, 0xff, 0x63, 0x92, 0x4c, 0xe6, 0x27, 0xdc, 0x45, 0xf0, 0xc5, 0x65, 0xe6, 0x4e, 0xf9, 0x00,
	0xb1, 0x2b, 0x3a, 0x76, 0xd1, 0xc1, 0x5e, 0x4c, 0x92, 0xc3, 0xf9, 0x6c, 0xe0, 0x5d, 0xec, 0xad,
	0x6f, 0xf4, 0xa8, 0xe4, 0xeb, 0x46, 0xdc, 0x89, 0xb7, 0xdc, 0x4b, 0xd8, 0x0d, 0x3d, 0xae, 0x1a,
	0x30, 0xa2, 0x18, 0xde, 0x1c, 0x8c, 0xac, 0x4b, 0xc8, 0xfb, 0x55, 0xb8, 0x5f, 0x85, 0x3f, 0x78,
	0x22, 0x9f, 0x58, 0xc9, 0xf0, 0x7f, 0x7d, 0xf6, 0xd1, 0x45, 0xe4, 0xb3, 0x8b, 0xc8, 0xae, 0x8b,
	0xc8, 0xfb, 0x57, 0xf4, 0xef, 0xc9, 0x6f, 0xfb, 0x4d, 0xc8, 0x72, 0xdf, 0x3a, 0x5c, 0xfc, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x28, 0x89, 0xd6, 0xaa, 0x99, 0x01, 0x00, 0x00,
}

func (m *ProcessListeningOnPort) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessListeningOnPort) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessListeningOnPort) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CloseTimestamp != nil {
		{
			size, err := m.CloseTimestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProcessListeningOnPort(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Process != nil {
		{
			size, err := m.Process.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProcessListeningOnPort(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Protocol != 0 {
		i = encodeVarintProcessListeningOnPort(dAtA, i, uint64(m.Protocol))
		i--
		dAtA[i] = 0x10
	}
	if m.Port != 0 {
		i = encodeVarintProcessListeningOnPort(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProcessListeningOnPort(dAtA []byte, offset int, v uint64) int {
	offset -= sovProcessListeningOnPort(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProcessListeningOnPort) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Port != 0 {
		n += 1 + sovProcessListeningOnPort(uint64(m.Port))
	}
	if m.Protocol != 0 {
		n += 1 + sovProcessListeningOnPort(uint64(m.Protocol))
	}
	if m.Process != nil {
		l = m.Process.Size()
		n += 1 + l + sovProcessListeningOnPort(uint64(l))
	}
	if m.CloseTimestamp != nil {
		l = m.CloseTimestamp.Size()
		n += 1 + l + sovProcessListeningOnPort(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProcessListeningOnPort(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProcessListeningOnPort(x uint64) (n int) {
	return sovProcessListeningOnPort(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProcessListeningOnPort) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessListeningOnPort
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
			return fmt.Errorf("proto: ProcessListeningOnPort: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessListeningOnPort: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessListeningOnPort
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessListeningOnPort
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Protocol |= L4Protocol(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Process", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessListeningOnPort
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProcessListeningOnPort
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessListeningOnPort
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Process == nil {
				m.Process = &ProcessIndicatorUniqueKey{}
			}
			if err := m.Process.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloseTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessListeningOnPort
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProcessListeningOnPort
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessListeningOnPort
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CloseTimestamp == nil {
				m.CloseTimestamp = &types.Timestamp{}
			}
			if err := m.CloseTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessListeningOnPort(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessListeningOnPort
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
func skipProcessListeningOnPort(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProcessListeningOnPort
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
					return 0, ErrIntOverflowProcessListeningOnPort
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
					return 0, ErrIntOverflowProcessListeningOnPort
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
				return 0, ErrInvalidLengthProcessListeningOnPort
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProcessListeningOnPort
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProcessListeningOnPort
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProcessListeningOnPort        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProcessListeningOnPort          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProcessListeningOnPort = fmt.Errorf("proto: unexpected end of group")
)
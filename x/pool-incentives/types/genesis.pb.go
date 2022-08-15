// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/pool-incentives/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the pool incentives module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	LockableDurations []time.Duration `protobuf:"bytes,2,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
	DistrInfo         *DistrInfo      `protobuf:"bytes,3,opt,name=distr_info,json=distrInfo,proto3" json:"distr_info,omitempty" yaml:"distr_info"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc1f078212600632, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func (m *GenesisState) GetDistrInfo() *DistrInfo {
	if m != nil {
		return m.DistrInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.poolincentives.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/pool-incentives/v1beta1/genesis.proto", fileDescriptor_cc1f078212600632)
}

var fileDescriptor_cc1f078212600632 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0x87, 0xbb, 0x60, 0x48, 0x2c, 0x5e, 0x68, 0x3c, 0x00, 0x87, 0x2d, 0x69, 0xa2, 0xc1, 0x03,
	0xbb, 0x16, 0x6f, 0x7a, 0x23, 0x24, 0xc6, 0x9b, 0xc1, 0x78, 0xf1, 0x42, 0xb6, 0x65, 0xa9, 0x1b,
	0xdb, 0x4e, 0xd3, 0x5d, 0x88, 0xbc, 0x85, 0x47, 0x9f, 0xc1, 0x27, 0xe1, 0xc8, 0xd1, 0x53, 0x35,
	0xf0, 0x06, 0x3c, 0x81, 0xe9, 0x3f, 0x31, 0x21, 0x91, 0x5b, 0x27, 0xf3, 0xfd, 0x66, 0xbe, 0xe9,
	0xea, 0x3d, 0x90, 0x01, 0x48, 0x21, 0x69, 0x04, 0xe0, 0xf7, 0x44, 0xe8, 0xf2, 0x50, 0x89, 0x39,
	0x97, 0x74, 0x6e, 0x3b, 0x5c, 0x31, 0x9b, 0x7a, 0x3c, 0xe4, 0x52, 0x48, 0x12, 0xc5, 0xa0, 0xc0,
	0xc0, 0x05, 0x4e, 0x52, 0x7c, 0x47, 0x93, 0x82, 0x6e, 0x9f, 0x7a, 0xe0, 0x41, 0x86, 0xd2, 0xf4,
	0x2b, 0x4f, 0xb5, 0xb1, 0x07, 0xe0, 0xf9, 0x9c, 0x66, 0x95, 0x33, 0x9b, 0xd2, 0xc9, 0x2c, 0x66,
	0x4a, 0x40, 0x58, 0xf4, 0x2f, 0x0f, 0x49, 0xfc, 0xd9, 0x94, 0x25, 0xac, 0x8f, 0x8a, 0x7e, 0x72,
	0x9b, 0x9b, 0x3d, 0x28, 0xa6, 0xb8, 0x31, 0xd4, 0x6b, 0x11, 0x8b, 0x59, 0x20, 0x9b, 0xa8, 0x83,
	0xba, 0xf5, 0xfe, 0x39, 0xf9, 0xdf, 0x94, 0xdc, 0x67, 0xf4, 0xe0, 0x68, 0x99, 0x98, 0xda, 0xa8,
	0xc8, 0x1a, 0xa0, 0x1b, 0x3e, 0xb8, 0x2f, 0xcc, 0xf1, 0xf9, 0xb8, 0x74, 0x94, 0xcd, 0x4a, 0xa7,
	0xda, 0xad, 0xf7, 0x5b, 0x24, 0xbf, 0x82, 0x94, 0x57, 0x90, 0x61, 0x41, 0x0c, 0xce, 0xd2, 0x21,
	0xdb, 0xc4, 0x6c, 0x2d, 0x58, 0xe0, 0x5f, 0x5b, 0xfb, 0x23, 0xac, 0xf7, 0x2f, 0x13, 0x8d, 0x1a,
	0x65, 0xa3, 0x0c, 0x4a, 0xc3, 0xd5, 0xf5, 0x89, 0x90, 0x2a, 0x1e, 0x8b, 0x70, 0x0a, 0xcd, 0x6a,
	0xa6, 0x7e, 0x71, 0x48, 0x7d, 0x98, 0x26, 0xee, 0xc2, 0x29, 0x0c, 0x5a, 0xcb, 0xc4, 0x44, 0xdb,
	0xc4, 0x6c, 0xe4, 0x8b, 0x77, 0xa3, 0xac, 0xd1, 0xf1, 0xe4, 0x97, 0x7a, 0x5c, 0xae, 0x31, 0x5a,
	0xad, 0x31, 0xfa, 0x5e, 0x63, 0xf4, 0xb6, 0xc1, 0xda, 0x6a, 0x83, 0xb5, 0xcf, 0x0d, 0xd6, 0x9e,
	0x6e, 0x3c, 0xa1, 0x9e, 0x67, 0x0e, 0x71, 0x21, 0xa0, 0xc5, 0xd2, 0x9e, 0xcf, 0x1c, 0x59, 0x16,
	0x74, 0x6e, 0xdb, 0xf4, 0x75, 0xef, 0x59, 0xd4, 0x22, 0xe2, 0xd2, 0xa9, 0x65, 0x3f, 0xe2, 0xea,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x80, 0x92, 0x08, 0x02, 0x43, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DistrInfo != nil {
		{
			size, err := m.DistrInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintGenesis(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DistrInfo != nil {
		l = m.DistrInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DistrInfo == nil {
				m.DistrInfo = &DistrInfo{}
			}
			if err := m.DistrInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

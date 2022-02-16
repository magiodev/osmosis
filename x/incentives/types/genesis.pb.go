// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/genesis.proto

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

// GenesisState defines the incentives module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Gauges            []Gauge         `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges"`
	LockableDurations []time.Duration `protobuf:"bytes,3,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
	LastGaugeId       uint64          `protobuf:"varint,4,opt,name=last_gauge_id,json=lastGaugeId,proto3" json:"last_gauge_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a288ccc95d977d2d, []int{0}
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

func (m *GenesisState) GetGauges() []Gauge {
	if m != nil {
		return m.Gauges
	}
	return nil
}

func (m *GenesisState) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func (m *GenesisState) GetLastGaugeId() uint64 {
	if m != nil {
		return m.LastGaugeId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.incentives.GenesisState")
}

func init() { proto.RegisterFile("osmosis/incentives/genesis.proto", fileDescriptor_a288ccc95d977d2d) }

var fileDescriptor_a288ccc95d977d2d = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x93, 0xb6, 0xea, 0x90, 0xc2, 0x80, 0xc5, 0x90, 0x76, 0x70, 0xa2, 0x48, 0x48, 0x5d,
	0xb0, 0xa5, 0x22, 0xfe, 0x88, 0xb1, 0x42, 0xaa, 0x98, 0x40, 0x65, 0x63, 0xa9, 0x9c, 0xd6, 0x18,
	0x8b, 0x24, 0xae, 0x7a, 0x4e, 0x45, 0xdf, 0x82, 0x09, 0xf1, 0x48, 0x1d, 0x3b, 0x32, 0x15, 0xd4,
	0xbe, 0x01, 0x4f, 0x80, 0xe2, 0xc4, 0x02, 0xa9, 0xdd, 0x7c, 0xfe, 0x7e, 0xf7, 0xdd, 0x77, 0xe7,
	0x85, 0x0a, 0x52, 0x05, 0x12, 0xa8, 0xcc, 0xc6, 0x3c, 0xd3, 0x72, 0xce, 0x81, 0x0a, 0x9e, 0x71,
	0x90, 0x40, 0xa6, 0x33, 0xa5, 0x15, 0x42, 0x15, 0x41, 0xfe, 0x88, 0xce, 0xb1, 0x50, 0x42, 0x19,
	0x99, 0x16, 0xaf, 0x92, 0xec, 0x60, 0xa1, 0x94, 0x48, 0x38, 0x35, 0x55, 0x9c, 0x3f, 0xd1, 0x49,
	0x3e, 0x63, 0x5a, 0xaa, 0xac, 0xd2, 0x83, 0x3d, 0xb3, 0xa6, 0x6c, 0xc6, 0x52, 0xb0, 0x06, 0xfb,
	0xc2, 0xb0, 0x5c, 0xf0, 0x52, 0x8f, 0xde, 0x6b, 0xde, 0xc1, 0xa0, 0x0c, 0xf7, 0xa0, 0x99, 0xe6,
	0xe8, 0xca, 0x6b, 0x96, 0x06, 0xbe, 0x1b, 0xba, 0xdd, 0x56, 0xaf, 0x43, 0x76, 0xc3, 0x92, 0x7b,
	0x43, 0xf4, 0x1b, 0xcb, 0x75, 0xe0, 0x0c, 0x2b, 0x1e, 0x5d, 0x7a, 0x4d, 0xe3, 0x0c, 0x7e, 0x2d,
	0xac, 0x77, 0x5b, 0xbd, 0xf6, 0xbe, 0xce, 0x41, 0x41, 0xd8, 0xc6, 0x12, 0x47, 0xca, 0x43, 0x89,
	0x1a, 0xbf, 0xb0, 0x38, 0xe1, 0x23, 0xbb, 0x1f, 0xf8, 0xf5, 0xca, 0xa4, 0xbc, 0x00, 0xb1, 0x17,
	0x20, 0x37, 0x15, 0xd1, 0x3f, 0x29, 0x4c, 0x7e, 0xd6, 0x41, 0x7b, 0xc1, 0xd2, 0xe4, 0x3a, 0xda,
	0xb5, 0x88, 0x3e, 0xbe, 0x02, 0x77, 0x78, 0x64, 0x05, 0xdb, 0x08, 0x28, 0xf2, 0x0e, 0x13, 0x06,
	0x7a, 0x64, 0xe6, 0x8f, 0xe4, 0xc4, 0x6f, 0x84, 0x6e, 0xb7, 0x31, 0x6c, 0x15, 0x9f, 0x26, 0xe0,
	0xed, 0xa4, 0x7f, 0xb7, 0xdc, 0x60, 0x77, 0xb5, 0xc1, 0xee, 0xf7, 0x06, 0xbb, 0x6f, 0x5b, 0xec,
	0xac, 0xb6, 0xd8, 0xf9, 0xdc, 0x62, 0xe7, 0xf1, 0x5c, 0x48, 0xfd, 0x9c, 0xc7, 0x64, 0xac, 0x52,
	0x5a, 0x6d, 0x78, 0x9a, 0xb0, 0x18, 0x6c, 0x41, 0xe7, 0x17, 0xf4, 0xf5, 0xff, 0xbd, 0xf5, 0x62,
	0xca, 0x21, 0x6e, 0x9a, 0x0d, 0xce, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x64, 0x20, 0x17,
	0x1f, 0x02, 0x00, 0x00,
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
	if m.LastGaugeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastGaugeId))
		i--
		dAtA[i] = 0x20
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
			dAtA[i] = 0x1a
		}
	}
	if len(m.Gauges) > 0 {
		for iNdEx := len(m.Gauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
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
	if len(m.Gauges) > 0 {
		for _, e := range m.Gauges {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastGaugeId != 0 {
		n += 1 + sovGenesis(uint64(m.LastGaugeId))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
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
			m.Gauges = append(m.Gauges, Gauge{})
			if err := m.Gauges[len(m.Gauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastGaugeId", wireType)
			}
			m.LastGaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastGaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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

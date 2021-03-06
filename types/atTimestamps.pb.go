// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atTimestamps.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import types1 "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// A set of common database timestamps
//
type CreatedAt struct {
	CreatedAt *time.Time `protobuf:"bytes,1,opt,name=created_at,json=createdAt,stdtime" json:"created_at,omitempty" db:"created_at"`
}

func (m *CreatedAt) Reset()                    { *m = CreatedAt{} }
func (m *CreatedAt) String() string            { return proto.CompactTextString(m) }
func (*CreatedAt) ProtoMessage()               {}
func (*CreatedAt) Descriptor() ([]byte, []int) { return fileDescriptorAtTimestamps, []int{0} }

func (m *CreatedAt) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type UpdatedAt struct {
	UpdatedAt *time.Time `protobuf:"bytes,1,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *UpdatedAt) Reset()                    { *m = UpdatedAt{} }
func (m *UpdatedAt) String() string            { return proto.CompactTextString(m) }
func (*UpdatedAt) ProtoMessage()               {}
func (*UpdatedAt) Descriptor() ([]byte, []int) { return fileDescriptorAtTimestamps, []int{1} }

func (m *UpdatedAt) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ArchivedAt struct {
	ArchivedAt *time.Time `protobuf:"bytes,1,opt,name=archived_at,json=archivedAt,stdtime" json:"archived_at,omitempty" db:"archived_at"`
}

func (m *ArchivedAt) Reset()                    { *m = ArchivedAt{} }
func (m *ArchivedAt) String() string            { return proto.CompactTextString(m) }
func (*ArchivedAt) ProtoMessage()               {}
func (*ArchivedAt) Descriptor() ([]byte, []int) { return fileDescriptorAtTimestamps, []int{2} }

func (m *ArchivedAt) GetArchivedAt() *time.Time {
	if m != nil {
		return m.ArchivedAt
	}
	return nil
}

type DeletedAt struct {
	DeletedAt *time.Time `protobuf:"bytes,1,opt,name=deleted_at,json=deletedAt,stdtime" json:"deleted_at,omitempty" db:"deleted_at"`
}

func (m *DeletedAt) Reset()                    { *m = DeletedAt{} }
func (m *DeletedAt) String() string            { return proto.CompactTextString(m) }
func (*DeletedAt) ProtoMessage()               {}
func (*DeletedAt) Descriptor() ([]byte, []int) { return fileDescriptorAtTimestamps, []int{3} }

func (m *DeletedAt) GetDeletedAt() *time.Time {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type Timestamps struct {
	CreatedAt  *time.Time `protobuf:"bytes,1,opt,name=created_at,json=createdAt,stdtime" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at,omitempty" db:"updated_at"`
	ArchivedAt *time.Time `protobuf:"bytes,3,opt,name=archived_at,json=archivedAt,stdtime" json:"archived_at,omitempty" db:"archived_at"`
	DeletedAt  *time.Time `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,stdtime" json:"deleted_at,omitempty" db:"deleted_at"`
}

func (m *Timestamps) Reset()                    { *m = Timestamps{} }
func (m *Timestamps) String() string            { return proto.CompactTextString(m) }
func (*Timestamps) ProtoMessage()               {}
func (*Timestamps) Descriptor() ([]byte, []int) { return fileDescriptorAtTimestamps, []int{4} }

func (m *Timestamps) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Timestamps) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Timestamps) GetArchivedAt() *time.Time {
	if m != nil {
		return m.ArchivedAt
	}
	return nil
}

func (m *Timestamps) GetDeletedAt() *time.Time {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatedAt)(nil), "canaryhealth.protobuf.CreatedAt")
	proto.RegisterType((*UpdatedAt)(nil), "canaryhealth.protobuf.UpdatedAt")
	proto.RegisterType((*ArchivedAt)(nil), "canaryhealth.protobuf.ArchivedAt")
	proto.RegisterType((*DeletedAt)(nil), "canaryhealth.protobuf.DeletedAt")
	proto.RegisterType((*Timestamps)(nil), "canaryhealth.protobuf.Timestamps")
}
func (m *CreatedAt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatedAt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CreatedAt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.CreatedAt)))
		n1, err := types1.StdTimeMarshalTo(*m.CreatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *UpdatedAt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatedAt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.UpdatedAt)))
		n2, err := types1.StdTimeMarshalTo(*m.UpdatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ArchivedAt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchivedAt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ArchivedAt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.ArchivedAt)))
		n3, err := types1.StdTimeMarshalTo(*m.ArchivedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *DeletedAt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeletedAt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DeletedAt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.DeletedAt)))
		n4, err := types1.StdTimeMarshalTo(*m.DeletedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *Timestamps) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamps) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CreatedAt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.CreatedAt)))
		n5, err := types1.StdTimeMarshalTo(*m.CreatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.UpdatedAt != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.UpdatedAt)))
		n6, err := types1.StdTimeMarshalTo(*m.UpdatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.ArchivedAt != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.ArchivedAt)))
		n7, err := types1.StdTimeMarshalTo(*m.ArchivedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.DeletedAt != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAtTimestamps(dAtA, i, uint64(types1.SizeOfStdTime(*m.DeletedAt)))
		n8, err := types1.StdTimeMarshalTo(*m.DeletedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func encodeVarintAtTimestamps(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreatedAt) Size() (n int) {
	var l int
	_ = l
	if m.CreatedAt != nil {
		l = types1.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	return n
}

func (m *UpdatedAt) Size() (n int) {
	var l int
	_ = l
	if m.UpdatedAt != nil {
		l = types1.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	return n
}

func (m *ArchivedAt) Size() (n int) {
	var l int
	_ = l
	if m.ArchivedAt != nil {
		l = types1.SizeOfStdTime(*m.ArchivedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	return n
}

func (m *DeletedAt) Size() (n int) {
	var l int
	_ = l
	if m.DeletedAt != nil {
		l = types1.SizeOfStdTime(*m.DeletedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	return n
}

func (m *Timestamps) Size() (n int) {
	var l int
	_ = l
	if m.CreatedAt != nil {
		l = types1.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	if m.UpdatedAt != nil {
		l = types1.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	if m.ArchivedAt != nil {
		l = types1.SizeOfStdTime(*m.ArchivedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	if m.DeletedAt != nil {
		l = types1.SizeOfStdTime(*m.DeletedAt)
		n += 1 + l + sovAtTimestamps(uint64(l))
	}
	return n
}

func sovAtTimestamps(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAtTimestamps(x uint64) (n int) {
	return sovAtTimestamps(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatedAt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtTimestamps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatedAt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatedAt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtTimestamps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtTimestamps
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
func (m *UpdatedAt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtTimestamps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdatedAt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatedAt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedAt == nil {
				m.UpdatedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtTimestamps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtTimestamps
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
func (m *ArchivedAt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtTimestamps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ArchivedAt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchivedAt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArchivedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ArchivedAt == nil {
				m.ArchivedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.ArchivedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtTimestamps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtTimestamps
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
func (m *DeletedAt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtTimestamps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeletedAt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeletedAt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeletedAt == nil {
				m.DeletedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.DeletedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtTimestamps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtTimestamps
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
func (m *Timestamps) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtTimestamps
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamps: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamps: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedAt == nil {
				m.UpdatedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArchivedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ArchivedAt == nil {
				m.ArchivedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.ArchivedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtTimestamps
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeletedAt == nil {
				m.DeletedAt = new(time.Time)
			}
			if err := types1.StdTimeUnmarshal(m.DeletedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtTimestamps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtTimestamps
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
func skipAtTimestamps(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAtTimestamps
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
					return 0, ErrIntOverflowAtTimestamps
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAtTimestamps
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAtTimestamps
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAtTimestamps
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAtTimestamps(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAtTimestamps = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAtTimestamps   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("atTimestamps.proto", fileDescriptorAtTimestamps) }

var fileDescriptorAtTimestamps = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2c, 0x09, 0xc9,
	0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x28, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x4d, 0x4e, 0xcc, 0x4b, 0x2c, 0xaa, 0xcc, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x80, 0x88, 0x25, 0x95,
	0xa6, 0x49, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7,
	0xa7, 0xe7, 0xeb, 0xc3, 0x64, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0xe8, 0x90, 0x92, 0x4f, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0x45, 0xa8, 0x2a, 0x81, 0xd9, 0x03, 0x51, 0xa0, 0x14, 0xcf, 0xc5, 0xe9,
	0x5c, 0x94, 0x9a, 0x58, 0x92, 0x9a, 0xe2, 0x58, 0x22, 0x14, 0xc4, 0xc5, 0x95, 0x0c, 0xe1, 0xc4,
	0x27, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xe9, 0x41, 0x8c, 0x80, 0x3b, 0x41,
	0x0f, 0xee, 0x54, 0x27, 0xf1, 0x4f, 0xf7, 0xe4, 0xf9, 0x53, 0x92, 0xac, 0x94, 0x10, 0xba, 0x94,
	0x26, 0xdc, 0x97, 0x67, 0x0c, 0xe2, 0x4c, 0x86, 0x99, 0x09, 0xb2, 0x20, 0xb4, 0x20, 0x05, 0x61,
	0x41, 0x29, 0x84, 0x43, 0xa2, 0x05, 0x08, 0x5d, 0x50, 0x0b, 0x4a, 0x61, 0x66, 0x2a, 0x25, 0x73,
	0x71, 0x39, 0x16, 0x25, 0x67, 0x64, 0x96, 0x81, 0x6d, 0x08, 0xe5, 0xe2, 0x4e, 0x84, 0xf2, 0x88,
	0xb3, 0x42, 0xe2, 0xd3, 0x3d, 0x79, 0x01, 0x90, 0x15, 0x48, 0xda, 0x20, 0x76, 0x70, 0x25, 0xc2,
	0x8d, 0x05, 0xf9, 0xc2, 0x25, 0x35, 0x27, 0x15, 0xee, 0x8b, 0x14, 0x08, 0x87, 0x44, 0x5f, 0x20,
	0x74, 0x41, 0x7d, 0x91, 0x02, 0x33, 0x53, 0xe9, 0x3e, 0x13, 0x17, 0x17, 0x22, 0x0d, 0xd0, 0x22,
	0x26, 0xd0, 0x02, 0x9f, 0x89, 0x1a, 0x81, 0x8f, 0x1e, 0xdc, 0xcc, 0xd4, 0x09, 0x6e, 0xb4, 0x10,
	0x66, 0xa1, 0x46, 0x08, 0x3b, 0x89, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1,
	0x81, 0x0d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xae, 0xff, 0x61, 0x86, 0x03, 0x00,
	0x00,
}

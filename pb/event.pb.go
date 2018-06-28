// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		event.proto
		service_analyzer.proto
		service_data.proto

	It has these top-level messages:
		CommitRevision
		ReferencePointer
		PushEvent
		PullRequestEvent
		EventResponse
		Comment
		File
		Change
		ChangesRequest
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import gopkg_in_src_d_go_git_v4_plumbing "gopkg.in/src-d/go-git.v4/plumbing"
import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CommitRevision defines a range of commits, from a base to a head.
type CommitRevision struct {
	// Base of the revision.
	Base ReferencePointer `protobuf:"bytes,1,opt,name=base" json:"base"`
	// Head of the revision.
	Head ReferencePointer `protobuf:"bytes,2,opt,name=head" json:"head"`
}

func (m *CommitRevision) Reset()                    { *m = CommitRevision{} }
func (m *CommitRevision) String() string            { return proto.CompactTextString(m) }
func (*CommitRevision) ProtoMessage()               {}
func (*CommitRevision) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{0} }

// ReferencePointer is the reference to a git refererence in a repository.
type ReferencePointer struct {
	// InternalRepositoryURL is the origina; clone URL not canonized.
	InternalRepositoryURL string `protobuf:"bytes,1,opt,name=internal_repository_url,json=internalRepositoryUrl,proto3" json:"internal_repository_url,omitempty"`
	// ReferenceName is the name of the reference pointing.
	ReferenceName gopkg_in_src_d_go_git_v4_plumbing.ReferenceName `protobuf:"bytes,2,opt,name=ReferenceName,proto3,casttype=gopkg.in/src-d/go-git.v4/plumbing.ReferenceName" json:"ReferenceName,omitempty"`
	// Hash is the hash of the reference pointing.
	Hash string `protobuf:"bytes,3,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (m *ReferencePointer) Reset()                    { *m = ReferencePointer{} }
func (m *ReferencePointer) String() string            { return proto.CompactTextString(m) }
func (*ReferencePointer) ProtoMessage()               {}
func (*ReferencePointer) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{1} }

// PushEvent represents a Push to a git repository.
type PushEvent struct {
	// Provider triggering this event.
	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	// InternalId is the internal id for this event at the provider.
	InternalID string `protobuf:"bytes,2,opt,name=internal_id,json=internalId,proto3" json:"internal_id,omitempty"`
	// CreateAt is the timestamp of the creation date of the push event.
	CreatedAt time.Time `protobuf:"bytes,3,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	// Commits is the number of commits in the push.
	Commits uint32 `protobuf:"varint,4,opt,name=commits,proto3" json:"commits,omitempty"`
	// Commits is the number of distinct commits in the push.
	DistinctCommits uint32 `protobuf:"varint,5,opt,name=distinct_commits,json=distinctCommits,proto3" json:"distinct_commits,omitempty"`
	CommitRevision  `protobuf:"bytes,7,opt,name=commit_revision,json=commitRevision,embedded=commit_revision" json:"commit_revision"`
}

func (m *PushEvent) Reset()                    { *m = PushEvent{} }
func (m *PushEvent) String() string            { return proto.CompactTextString(m) }
func (*PushEvent) ProtoMessage()               {}
func (*PushEvent) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{2} }

// PullRequestEvent represents a PullRequest being created or updated.
type PullRequestEvent struct {
	// Provider triggering this event.
	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	// InternalId is the internal id for this event at the provider.
	InternalID string `protobuf:"bytes,2,opt,name=internal_id,json=internalId,proto3" json:"internal_id,omitempty"`
	// CreateAt is the timestamp of the creation date of the push event.
	CreatedAt time.Time `protobuf:"bytes,3,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	// UpdatedAt is the timestamp of the last modification of the pull request.
	UpdatedAt time.Time `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
	// IsMergeable, if the pull request is mergeable.
	IsMergeable bool `protobuf:"varint,5,opt,name=is_mergeable,json=isMergeable,proto3" json:"is_mergeable,omitempty"`
	// Source reference to the original branch and repository were the changes
	// came from.
	Source ReferencePointer `protobuf:"bytes,8,opt,name=source" json:"source"`
	// Merge  test branch where the PullRequest was merged.
	Merge          ReferencePointer `protobuf:"bytes,9,opt,name=merge" json:"merge"`
	CommitRevision `protobuf:"bytes,7,opt,name=commit_revision,json=commitRevision,embedded=commit_revision" json:"commit_revision"`
}

func (m *PullRequestEvent) Reset()                    { *m = PullRequestEvent{} }
func (m *PullRequestEvent) String() string            { return proto.CompactTextString(m) }
func (*PullRequestEvent) ProtoMessage()               {}
func (*PullRequestEvent) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{3} }

func init() {
	proto.RegisterType((*CommitRevision)(nil), "pb.CommitRevision")
	proto.RegisterType((*ReferencePointer)(nil), "pb.ReferencePointer")
	proto.RegisterType((*PushEvent)(nil), "pb.PushEvent")
	proto.RegisterType((*PullRequestEvent)(nil), "pb.PullRequestEvent")
}
func (m *CommitRevision) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommitRevision) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEvent(dAtA, i, uint64(m.Base.ProtoSize()))
	n1, err := m.Base.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintEvent(dAtA, i, uint64(m.Head.ProtoSize()))
	n2, err := m.Head.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *ReferencePointer) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReferencePointer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.InternalRepositoryURL) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.InternalRepositoryURL)))
		i += copy(dAtA[i:], m.InternalRepositoryURL)
	}
	if len(m.ReferenceName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ReferenceName)))
		i += copy(dAtA[i:], m.ReferenceName)
	}
	if len(m.Hash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	return i, nil
}

func (m *PushEvent) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Provider)))
		i += copy(dAtA[i:], m.Provider)
	}
	if len(m.InternalID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.InternalID)))
		i += copy(dAtA[i:], m.InternalID)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintEvent(dAtA, i, uint64(types.SizeOfStdTime(m.CreatedAt)))
	n3, err := types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.Commits != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Commits))
	}
	if m.DistinctCommits != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.DistinctCommits))
	}
	dAtA[i] = 0x3a
	i++
	i = encodeVarintEvent(dAtA, i, uint64(m.CommitRevision.ProtoSize()))
	n4, err := m.CommitRevision.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *PullRequestEvent) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullRequestEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Provider)))
		i += copy(dAtA[i:], m.Provider)
	}
	if len(m.InternalID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvent(dAtA, i, uint64(len(m.InternalID)))
		i += copy(dAtA[i:], m.InternalID)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintEvent(dAtA, i, uint64(types.SizeOfStdTime(m.CreatedAt)))
	n5, err := types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x22
	i++
	i = encodeVarintEvent(dAtA, i, uint64(types.SizeOfStdTime(m.UpdatedAt)))
	n6, err := types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.IsMergeable {
		dAtA[i] = 0x28
		i++
		if m.IsMergeable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x3a
	i++
	i = encodeVarintEvent(dAtA, i, uint64(m.CommitRevision.ProtoSize()))
	n7, err := m.CommitRevision.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x42
	i++
	i = encodeVarintEvent(dAtA, i, uint64(m.Source.ProtoSize()))
	n8, err := m.Source.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	dAtA[i] = 0x4a
	i++
	i = encodeVarintEvent(dAtA, i, uint64(m.Merge.ProtoSize()))
	n9, err := m.Merge.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n9
	return i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CommitRevision) ProtoSize() (n int) {
	var l int
	_ = l
	l = m.Base.ProtoSize()
	n += 1 + l + sovEvent(uint64(l))
	l = m.Head.ProtoSize()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *ReferencePointer) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.InternalRepositoryURL)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ReferenceName)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *PushEvent) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.InternalID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovEvent(uint64(l))
	if m.Commits != 0 {
		n += 1 + sovEvent(uint64(m.Commits))
	}
	if m.DistinctCommits != 0 {
		n += 1 + sovEvent(uint64(m.DistinctCommits))
	}
	l = m.CommitRevision.ProtoSize()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *PullRequestEvent) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.InternalID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovEvent(uint64(l))
	l = types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovEvent(uint64(l))
	if m.IsMergeable {
		n += 2
	}
	l = m.CommitRevision.ProtoSize()
	n += 1 + l + sovEvent(uint64(l))
	l = m.Source.ProtoSize()
	n += 1 + l + sovEvent(uint64(l))
	l = m.Merge.ProtoSize()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommitRevision) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: CommitRevision: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitRevision: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func (m *ReferencePointer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: ReferencePointer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReferencePointer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalRepositoryURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalRepositoryURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReferenceName = gopkg_in_src_d_go_git_v4_plumbing.ReferenceName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func (m *PushEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: PushEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commits", wireType)
			}
			m.Commits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commits |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistinctCommits", wireType)
			}
			m.DistinctCommits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistinctCommits |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitRevision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommitRevision.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func (m *PullRequestEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: PullRequestEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullRequestEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InternalID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsMergeable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsMergeable = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitRevision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommitRevision.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Source.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Merge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Merge.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvent
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
				next, err := skipEvent(dAtA[start:])
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
	ErrInvalidLengthEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("event.proto", fileDescriptorEvent) }

var fileDescriptorEvent = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0x36, 0x6d, 0x93, 0x9b, 0x7f, 0x3f, 0x34, 0xfa, 0x57, 0x98, 0x2c, 0xec, 0x92,
	0x55, 0x59, 0xd4, 0x46, 0x2d, 0x2f, 0xd0, 0x84, 0x4a, 0x44, 0x02, 0x14, 0x46, 0xb0, 0x60, 0x65,
	0xf9, 0xe3, 0xd6, 0x19, 0x61, 0x7b, 0xcc, 0xcc, 0x38, 0x12, 0x6f, 0xd1, 0x47, 0x60, 0xcb, 0x23,
	0xf0, 0x06, 0x59, 0x76, 0xc1, 0x3a, 0x40, 0xfa, 0x16, 0xac, 0x90, 0xc7, 0x76, 0x50, 0x60, 0x11,
	0x21, 0xb1, 0x60, 0x37, 0x77, 0xfc, 0xbb, 0x77, 0xce, 0x9c, 0x33, 0x86, 0x1e, 0xce, 0x30, 0x53,
	0x4e, 0x2e, 0xb8, 0xe2, 0x64, 0x2b, 0x0f, 0xfa, 0x67, 0x31, 0x53, 0xd3, 0x22, 0x70, 0x42, 0x9e,
	0xba, 0x31, 0x8f, 0xb9, 0xab, 0x3f, 0x05, 0xc5, 0xb5, 0xae, 0x74, 0xa1, 0x57, 0x55, 0x4b, 0xdf,
	0x8e, 0x39, 0x8f, 0x13, 0xfc, 0x49, 0x29, 0x96, 0xa2, 0x54, 0x7e, 0x9a, 0x57, 0xc0, 0x20, 0x87,
	0x83, 0x11, 0x4f, 0x53, 0xa6, 0x28, 0xce, 0x98, 0x64, 0x3c, 0x23, 0x0e, 0xb4, 0x03, 0x5f, 0xa2,
	0x69, 0x9c, 0x18, 0xa7, 0xbd, 0xf3, 0xff, 0x9d, 0x3c, 0x70, 0x28, 0x5e, 0xa3, 0xc0, 0x2c, 0xc4,
	0x09, 0x67, 0x99, 0x42, 0x31, 0x6c, 0xcf, 0x17, 0x76, 0x8b, 0x6a, 0xae, 0xe4, 0xa7, 0xe8, 0x47,
	0xe6, 0xd6, 0x66, 0xbe, 0xe4, 0x06, 0x9f, 0x0d, 0x38, 0xfa, 0x15, 0x20, 0x2f, 0xe1, 0x9e, 0x5e,
	0x64, 0x7e, 0xe2, 0x09, 0xcc, 0xb9, 0x64, 0x8a, 0x8b, 0xf7, 0x5e, 0x21, 0x12, 0xad, 0xa3, 0x3b,
	0xbc, 0xbf, 0x5c, 0xd8, 0xc7, 0xe3, 0x1a, 0xa1, 0x2b, 0xe2, 0x35, 0x7d, 0x46, 0x8f, 0xd9, 0xef,
	0xdb, 0x22, 0x21, 0x6f, 0x60, 0x7f, 0x75, 0xcc, 0x0b, 0x3f, 0x45, 0x2d, 0xb0, 0x3b, 0xbc, 0xf8,
	0xbe, 0xb0, 0xdd, 0x98, 0xe7, 0x6f, 0x63, 0x87, 0x65, 0xae, 0x14, 0xe1, 0x59, 0xe4, 0xc6, 0xbc,
	0xb4, 0xd5, 0x99, 0x3d, 0x76, 0xf3, 0xa4, 0x48, 0x03, 0x96, 0xc5, 0xce, 0x5a, 0x2b, 0x5d, 0x9f,
	0x44, 0x08, 0xb4, 0x9f, 0xfa, 0x72, 0x6a, 0x6e, 0x97, 0x13, 0xa9, 0x5e, 0x0f, 0x3e, 0x6e, 0x41,
	0x77, 0x52, 0xc8, 0xe9, 0x55, 0x19, 0x18, 0xe9, 0x43, 0x27, 0x17, 0x7c, 0xc6, 0x22, 0x14, 0xd5,
	0x05, 0xe8, 0xaa, 0x26, 0x2e, 0xf4, 0x56, 0x77, 0x65, 0x51, 0x2d, 0xeb, 0x60, 0xb9, 0xb0, 0xa1,
	0xb9, 0xdf, 0xf8, 0x09, 0x85, 0x06, 0x19, 0x47, 0x64, 0x04, 0x10, 0x0a, 0xf4, 0x15, 0x46, 0x9e,
	0xaf, 0xf4, 0xa1, 0xbd, 0xf3, 0xbe, 0x53, 0x25, 0xeb, 0x34, 0xc9, 0x3a, 0xaf, 0x9a, 0x64, 0x87,
	0x9d, 0xd2, 0xed, 0x9b, 0x2f, 0xb6, 0x41, 0xbb, 0x75, 0xdf, 0xa5, 0x22, 0x26, 0xec, 0x85, 0x3a,
	0x68, 0x69, 0xb6, 0x4f, 0x8c, 0xd3, 0x7d, 0xda, 0x94, 0xe4, 0x21, 0x1c, 0x45, 0x4c, 0x2a, 0x96,
	0x85, 0xca, 0x6b, 0x90, 0x1d, 0x8d, 0x1c, 0x36, 0xfb, 0xa3, 0x1a, 0xbd, 0x82, 0xc3, 0x8a, 0xf0,
	0x44, 0xfd, 0x5c, 0xcc, 0x3d, 0x2d, 0x87, 0x94, 0xb1, 0xaf, 0x3f, 0xa4, 0x4a, 0xc6, 0xed, 0xc2,
	0x36, 0xe8, 0x41, 0xb8, 0xf6, 0x65, 0xf0, 0x69, 0x1b, 0x8e, 0x26, 0x45, 0x92, 0x50, 0x7c, 0x57,
	0xa0, 0x54, 0xff, 0xaa, 0x65, 0x23, 0x80, 0x22, 0x8f, 0x9a, 0x21, 0xed, 0x3f, 0x19, 0x52, 0xf7,
	0x5d, 0x2a, 0xf2, 0x00, 0xfe, 0x63, 0xd2, 0x4b, 0x51, 0xc4, 0xe8, 0x07, 0x09, 0x6a, 0x67, 0x3b,
	0xb4, 0xc7, 0xe4, 0xf3, 0x66, 0xeb, 0x2f, 0xb9, 0x4a, 0xce, 0x61, 0x57, 0xf2, 0x42, 0x84, 0x68,
	0x76, 0x36, 0xfe, 0x8a, 0x35, 0x49, 0x1e, 0xc1, 0x8e, 0x96, 0x66, 0x76, 0x37, 0xb6, 0x54, 0xe0,
	0xd0, 0x9c, 0x7f, 0xb3, 0x5a, 0xf3, 0xa5, 0x65, 0xdc, 0x2e, 0x2d, 0xe3, 0xeb, 0xd2, 0x6a, 0xdd,
	0xdc, 0x59, 0xad, 0x0f, 0x77, 0x96, 0x11, 0xec, 0x6a, 0x4b, 0x2e, 0x7e, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x56, 0x9f, 0x05, 0xcc, 0xb4, 0x04, 0x00, 0x00,
}

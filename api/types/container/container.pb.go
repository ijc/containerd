// Code generated by protoc-gen-gogo.
// source: github.com/docker/containerd/api/types/container/container.proto
// DO NOT EDIT!

/*
	Package container is a generated protocol buffer package.

	It is generated from these files:
		github.com/docker/containerd/api/types/container/container.proto

	It has these top-level messages:
		Container
		Process
		User
		Event
		ExtraFd
*/
package container

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_CREATED Status = 0
	Status_RUNNING Status = 1
	Status_STOPPED Status = 2
	Status_PAUSED  Status = 3
)

var Status_name = map[int32]string{
	0: "CREATED",
	1: "RUNNING",
	2: "STOPPED",
	3: "PAUSED",
}
var Status_value = map[string]int32{
	"CREATED": 0,
	"RUNNING": 1,
	"STOPPED": 2,
	"PAUSED":  3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorContainer, []int{0} }

type Event_EventType int32

const (
	Event_EXIT       Event_EventType = 0
	Event_OOM        Event_EventType = 1
	Event_CREATE     Event_EventType = 2
	Event_START      Event_EventType = 3
	Event_EXEC_ADDED Event_EventType = 4
	Event_PAUSED     Event_EventType = 5
)

var Event_EventType_name = map[int32]string{
	0: "EXIT",
	1: "OOM",
	2: "CREATE",
	3: "START",
	4: "EXEC_ADDED",
	5: "PAUSED",
}
var Event_EventType_value = map[string]int32{
	"EXIT":       0,
	"OOM":        1,
	"CREATE":     2,
	"START":      3,
	"EXEC_ADDED": 4,
	"PAUSED":     5,
}

func (x Event_EventType) String() string {
	return proto.EnumName(Event_EventType_name, int32(x))
}
func (Event_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorContainer, []int{3, 0} }

type ExtraFd_Kind int32

const (
	ExtraFd_NULL ExtraFd_Kind = 0
	ExtraFd_TAP  ExtraFd_Kind = 1
)

var ExtraFd_Kind_name = map[int32]string{
	0: "NULL",
	1: "TAP",
}
var ExtraFd_Kind_value = map[string]int32{
	"NULL": 0,
	"TAP":  1,
}

func (x ExtraFd_Kind) String() string {
	return proto.EnumName(ExtraFd_Kind_name, int32(x))
}
func (ExtraFd_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptorContainer, []int{4, 0} }

type Container struct {
	ID     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid    uint32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{0} }

type Process struct {
	Pid        uint32   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Args       []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	Env        []string `protobuf:"bytes,3,rep,name=env" json:"env,omitempty"`
	User       *User    `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Cwd        string   `protobuf:"bytes,5,opt,name=cwd,proto3" json:"cwd,omitempty"`
	Terminal   bool     `protobuf:"varint,6,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExitStatus uint32   `protobuf:"varint,7,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	Status     Status   `protobuf:"varint,8,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{1} }

type User struct {
	Uid            uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid            uint32   `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	AdditionalGids []uint32 `protobuf:"varint,3,rep,packed,name=additional_gids,json=additionalGids" json:"additional_gids,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{2} }

type Event struct {
	ID         string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type       Event_EventType `protobuf:"varint,2,opt,name=type,proto3,enum=containerd.v1.types.Event_EventType" json:"type,omitempty"`
	Pid        uint32          `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	ExitStatus uint32          `protobuf:"varint,4,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{3} }

type ExtraFd struct {
	Kind ExtraFd_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=containerd.v1.types.ExtraFd_Kind" json:"kind,omitempty"`
	Args []string     `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *ExtraFd) Reset()                    { *m = ExtraFd{} }
func (*ExtraFd) ProtoMessage()               {}
func (*ExtraFd) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{4} }

func init() {
	proto.RegisterType((*Container)(nil), "containerd.v1.types.Container")
	proto.RegisterType((*Process)(nil), "containerd.v1.types.Process")
	proto.RegisterType((*User)(nil), "containerd.v1.types.User")
	proto.RegisterType((*Event)(nil), "containerd.v1.types.Event")
	proto.RegisterType((*ExtraFd)(nil), "containerd.v1.types.ExtraFd")
	proto.RegisterEnum("containerd.v1.types.Status", Status_name, Status_value)
	proto.RegisterEnum("containerd.v1.types.Event_EventType", Event_EventType_name, Event_EventType_value)
	proto.RegisterEnum("containerd.v1.types.ExtraFd_Kind", ExtraFd_Kind_name, ExtraFd_Kind_value)
}
func (m *Container) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Container) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Pid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Pid))
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *Process) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Process) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Pid))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.User != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Cwd) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Cwd)))
		i += copy(dAtA[i:], m.Cwd)
	}
	if m.Terminal {
		dAtA[i] = 0x30
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ExitStatus != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.ExitStatus))
	}
	if m.Status != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Uid))
	}
	if m.Gid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Gid))
	}
	if len(m.AdditionalGids) > 0 {
		dAtA3 := make([]byte, len(m.AdditionalGids)*10)
		var j2 int
		for _, num := range m.AdditionalGids {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContainer(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	return i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Type))
	}
	if m.Pid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Pid))
	}
	if m.ExitStatus != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.ExitStatus))
	}
	return i, nil
}

func (m *ExtraFd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtraFd) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Kind))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Container(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Container(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintContainer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Container) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Pid != 0 {
		n += 1 + sovContainer(uint64(m.Pid))
	}
	if m.Status != 0 {
		n += 1 + sovContainer(uint64(m.Status))
	}
	return n
}

func (m *Process) Size() (n int) {
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovContainer(uint64(m.Pid))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovContainer(uint64(l))
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			l = len(s)
			n += 1 + l + sovContainer(uint64(l))
		}
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Cwd)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	if m.ExitStatus != 0 {
		n += 1 + sovContainer(uint64(m.ExitStatus))
	}
	if m.Status != 0 {
		n += 1 + sovContainer(uint64(m.Status))
	}
	return n
}

func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovContainer(uint64(m.Uid))
	}
	if m.Gid != 0 {
		n += 1 + sovContainer(uint64(m.Gid))
	}
	if len(m.AdditionalGids) > 0 {
		l = 0
		for _, e := range m.AdditionalGids {
			l += sovContainer(uint64(e))
		}
		n += 1 + sovContainer(uint64(l)) + l
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovContainer(uint64(m.Type))
	}
	if m.Pid != 0 {
		n += 1 + sovContainer(uint64(m.Pid))
	}
	if m.ExitStatus != 0 {
		n += 1 + sovContainer(uint64(m.ExitStatus))
	}
	return n
}

func (m *ExtraFd) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovContainer(uint64(m.Kind))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovContainer(uint64(l))
		}
	}
	return n
}

func sovContainer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContainer(x uint64) (n int) {
	return sovContainer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Container) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Container{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Process) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Process{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Args:` + fmt.Sprintf("%v", this.Args) + `,`,
		`Env:` + fmt.Sprintf("%v", this.Env) + `,`,
		`User:` + strings.Replace(fmt.Sprintf("%v", this.User), "User", "User", 1) + `,`,
		`Cwd:` + fmt.Sprintf("%v", this.Cwd) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func (this *User) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&User{`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`Gid:` + fmt.Sprintf("%v", this.Gid) + `,`,
		`AdditionalGids:` + fmt.Sprintf("%v", this.AdditionalGids) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExtraFd) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExtraFd{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Args:` + fmt.Sprintf("%v", this.Args) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContainer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Container) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: Container: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Container: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *Process) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: Process: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Process: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cwd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cwd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
			m.Terminal = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			m.Gid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AdditionalGids = append(m.AdditionalGids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthContainer
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowContainer
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AdditionalGids = append(m.AdditionalGids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalGids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Event_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *ExtraFd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ExtraFd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtraFd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (ExtraFd_Kind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func skipContainer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContainer
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
					return 0, ErrIntOverflowContainer
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
					return 0, ErrIntOverflowContainer
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
				return 0, ErrInvalidLengthContainer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContainer
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
				next, err := skipContainer(dAtA[start:])
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
	ErrInvalidLengthContainer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContainer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/docker/containerd/api/types/container/container.proto", fileDescriptorContainer)
}

var fileDescriptorContainer = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0x7b, 0x93, 0x34, 0x6d, 0x4f, 0x59, 0x0d, 0x57, 0x91, 0x6c, 0x42, 0x16, 0x83, 0x60,
	0x10, 0x4c, 0xb1, 0x43, 0x10, 0x7c, 0xb1, 0x6b, 0xe2, 0x28, 0xce, 0xb6, 0xde, 0xa6, 0xb0, 0xb7,
	0x92, 0xf5, 0x86, 0x78, 0xdd, 0x96, 0x94, 0x24, 0xad, 0xdb, 0x9b, 0x1f, 0x6f, 0x8f, 0x3e, 0xfa,
	0xa4, 0xae, 0x9f, 0xc1, 0x0f, 0x20, 0xf7, 0xa6, 0x5d, 0xc6, 0xa8, 0xe0, 0x4b, 0x39, 0xf7, 0xf4,
	0x97, 0xff, 0xfd, 0xff, 0xcf, 0x49, 0xe0, 0x5d, 0xc4, 0xf2, 0xcf, 0x8b, 0x53, 0x67, 0x96, 0x5c,
	0xb4, 0x69, 0x32, 0x3b, 0x0b, 0xd3, 0xf6, 0x2c, 0x89, 0xf3, 0x80, 0xc5, 0x61, 0x4a, 0xdb, 0xc1,
	0x9c, 0xb5, 0xf3, 0xab, 0x79, 0x98, 0x95, 0xcd, 0xb2, 0x72, 0xe6, 0x69, 0x92, 0x27, 0xf8, 0x61,
	0xc9, 0x3b, 0xcb, 0x57, 0x8e, 0xc0, 0xf7, 0x1e, 0x45, 0x49, 0x94, 0x88, 0xff, 0xdb, 0xbc, 0x2a,
	0x50, 0xeb, 0x0b, 0x34, 0x7a, 0x1b, 0x18, 0x3f, 0x06, 0x89, 0x51, 0x1d, 0x99, 0xc8, 0x6e, 0x1c,
	0xaa, 0xab, 0x9f, 0xfb, 0x52, 0xdf, 0x25, 0x12, 0xa3, 0x58, 0x03, 0x79, 0xce, 0xa8, 0x2e, 0x99,
	0xc8, 0xde, 0x21, 0xbc, 0xc4, 0x07, 0xa0, 0x66, 0x79, 0x90, 0x2f, 0x32, 0x5d, 0x36, 0x91, 0xdd,
	0xea, 0x3c, 0x71, 0xb6, 0x5c, 0xe9, 0x8c, 0x05, 0x42, 0xd6, 0xa8, 0xf5, 0x07, 0x41, 0x6d, 0x94,
	0x26, 0xb3, 0x30, 0xcb, 0x36, 0x92, 0xa8, 0x94, 0xc4, 0xa0, 0x04, 0x69, 0x94, 0xe9, 0x92, 0x29,
	0xdb, 0x0d, 0x22, 0x6a, 0x4e, 0x85, 0xf1, 0x52, 0x97, 0x45, 0x8b, 0x97, 0xf8, 0x25, 0x28, 0x8b,
	0x2c, 0x4c, 0x75, 0xc5, 0x44, 0x76, 0xb3, 0xb3, 0xbb, 0xf5, 0xda, 0x49, 0x16, 0xa6, 0x44, 0x60,
	0x5c, 0x60, 0xf6, 0x95, 0xea, 0x55, 0x1e, 0x89, 0xf0, 0x12, 0xef, 0x41, 0x3d, 0x0f, 0xd3, 0x0b,
	0x16, 0x07, 0xe7, 0xba, 0x6a, 0x22, 0xbb, 0x4e, 0x6e, 0xcf, 0x78, 0x1f, 0x9a, 0xe1, 0x25, 0xcb,
	0xa7, 0xeb, 0x68, 0x35, 0x61, 0x0e, 0x78, 0xab, 0x48, 0x72, 0x27, 0x76, 0xfd, 0xff, 0x63, 0x8f,
	0x41, 0x99, 0xac, 0xbd, 0x2c, 0xca, 0xc8, 0x8b, 0x62, 0xae, 0x51, 0x39, 0xd7, 0x88, 0x51, 0xfc,
	0x1c, 0x1e, 0x04, 0x94, 0xb2, 0x9c, 0x25, 0x71, 0x70, 0x3e, 0x8d, 0x18, 0xcd, 0x44, 0xf8, 0x1d,
	0xd2, 0x2a, 0xdb, 0x47, 0x8c, 0x66, 0xd6, 0x2f, 0x04, 0x55, 0x6f, 0x19, 0xc6, 0xf9, 0x3f, 0x97,
	0xf6, 0x06, 0x14, 0x6e, 0x47, 0xa8, 0xb7, 0x3a, 0xcf, 0xb6, 0x3a, 0x15, 0x0a, 0xc5, 0xaf, 0x7f,
	0x35, 0x0f, 0x89, 0x78, 0x62, 0xb3, 0x1b, 0xb9, 0xdc, 0xcd, 0xbd, 0xc1, 0x28, 0xf7, 0x07, 0x63,
	0x7d, 0x82, 0xc6, 0xad, 0x0a, 0xae, 0x83, 0xe2, 0x9d, 0xf4, 0x7d, 0xad, 0x82, 0x6b, 0x20, 0x0f,
	0x87, 0x1f, 0x35, 0x84, 0x01, 0xd4, 0x1e, 0xf1, 0xba, 0xbe, 0xa7, 0x49, 0xb8, 0x01, 0xd5, 0xb1,
	0xdf, 0x25, 0xbe, 0x26, 0xe3, 0x16, 0x80, 0x77, 0xe2, 0xf5, 0xa6, 0x5d, 0xd7, 0xf5, 0x5c, 0x4d,
	0xe1, 0xd8, 0xa8, 0x3b, 0x19, 0x7b, 0xae, 0x56, 0xb5, 0x12, 0xa8, 0x79, 0x97, 0x79, 0x1a, 0xbc,
	0xa7, 0xf8, 0x35, 0x28, 0x67, 0x2c, 0x2e, 0x42, 0xb6, 0x3a, 0x4f, 0xb7, 0x47, 0x29, 0x58, 0xe7,
	0x03, 0x8b, 0x29, 0x11, 0xf8, 0xb6, 0x37, 0xca, 0xda, 0x05, 0x85, 0x13, 0xdc, 0xe3, 0x60, 0x72,
	0x7c, 0x5c, 0x78, 0xf4, 0xbb, 0x23, 0x0d, 0xbd, 0x78, 0x0b, 0xea, 0x7a, 0xcd, 0x4d, 0xa8, 0x15,
	0x6e, 0x5d, 0xad, 0xc2, 0x0f, 0x64, 0x32, 0x18, 0xf4, 0x07, 0x47, 0x1a, 0xe2, 0x87, 0xb1, 0x3f,
	0x1c, 0x8d, 0x3c, 0x57, 0x93, 0xee, 0xb8, 0x95, 0x0f, 0xf5, 0xeb, 0x1b, 0xa3, 0xf2, 0xe3, 0xc6,
	0xa8, 0x7c, 0x5b, 0x19, 0xe8, 0x7a, 0x65, 0xa0, 0xef, 0x2b, 0x03, 0xfd, 0x5e, 0x19, 0xe8, 0x54,
	0x15, 0x1f, 0xda, 0xc1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x1d, 0x0a, 0xdb, 0xd7, 0x03,
	0x00, 0x00,
}

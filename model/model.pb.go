// Code generated by protoc-gen-go.
// source: model.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	Instance
	InstanceState
	Resource
	Template
	NoneTemplate
	LxcTemplate
	NullTemplate
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InstanceState_State int32

const (
	InstanceState_REGISTERED   InstanceState_State = 0
	InstanceState_QUEUED       InstanceState_State = 1
	InstanceState_STARTING     InstanceState_State = 2
	InstanceState_RUNNING      InstanceState_State = 3
	InstanceState_STOPPING     InstanceState_State = 4
	InstanceState_STOPPED      InstanceState_State = 5
	InstanceState_SHUTTINGDOWN InstanceState_State = 6
	InstanceState_TERMINATED   InstanceState_State = 7
)

var InstanceState_State_name = map[int32]string{
	0: "REGISTERED",
	1: "QUEUED",
	2: "STARTING",
	3: "RUNNING",
	4: "STOPPING",
	5: "STOPPED",
	6: "SHUTTINGDOWN",
	7: "TERMINATED",
}
var InstanceState_State_value = map[string]int32{
	"REGISTERED":   0,
	"QUEUED":       1,
	"STARTING":     2,
	"RUNNING":      3,
	"STOPPING":     4,
	"STOPPED":      5,
	"SHUTTINGDOWN": 6,
	"TERMINATED":   7,
}

func (x InstanceState_State) String() string {
	return proto.EnumName(InstanceState_State_name, int32(x))
}
func (InstanceState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Resource_State int32

const (
	Resource_REGISTERED   Resource_State = 0
	Resource_UNREGISTERED Resource_State = 1
)

var Resource_State_name = map[int32]string{
	0: "REGISTERED",
	1: "UNREGISTERED",
}
var Resource_State_value = map[string]int32{
	"REGISTERED":   0,
	"UNREGISTERED": 1,
}

func (x Resource_State) String() string {
	return proto.EnumName(Resource_State_name, int32(x))
}
func (Resource_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Instance struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	SlaveId    string                     `protobuf:"bytes,2,opt,name=slave_id,json=slaveId" json:"slave_id,omitempty"`
	ResourceId string                     `protobuf:"bytes,3,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	LastState  *InstanceState             `protobuf:"bytes,4,opt,name=last_state,json=lastState" json:"last_state,omitempty"`
	CreatedAt  *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetSlaveId() string {
	if m != nil {
		return m.SlaveId
	}
	return ""
}

func (m *Instance) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Instance) GetLastState() *InstanceState {
	if m != nil {
		return m.LastState
	}
	return nil
}

func (m *Instance) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type InstanceState struct {
	State     InstanceState_State        `protobuf:"varint,1,opt,name=state,enum=model.InstanceState_State" json:"state,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *InstanceState) Reset()                    { *m = InstanceState{} }
func (m *InstanceState) String() string            { return proto.CompactTextString(m) }
func (*InstanceState) ProtoMessage()               {}
func (*InstanceState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InstanceState) GetState() InstanceState_State {
	if m != nil {
		return m.State
	}
	return InstanceState_REGISTERED
}

func (m *InstanceState) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type Resource struct {
	Id       string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	State    Resource_State `protobuf:"varint,3,opt,name=state,enum=model.Resource_State" json:"state,omitempty"`
	Template *Template      `protobuf:"bytes,4,opt,name=template" json:"template,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetState() Resource_State {
	if m != nil {
		return m.State
	}
	return Resource_REGISTERED
}

func (m *Resource) GetTemplate() *Template {
	if m != nil {
		return m.Template
	}
	return nil
}

type Template struct {
	TemplateUri string `protobuf:"bytes,1,opt,name=template_uri,json=templateUri" json:"template_uri,omitempty"`
	// Types that are valid to be assigned to Item:
	//	*Template_None
	//	*Template_Lxc
	//	*Template_Null
	Item      isTemplate_Item            `protobuf_oneof:"Item"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isTemplate_Item interface {
	isTemplate_Item()
}

type Template_None struct {
	None *NoneTemplate `protobuf:"bytes,500,opt,name=none,oneof"`
}
type Template_Lxc struct {
	Lxc *LxcTemplate `protobuf:"bytes,501,opt,name=lxc,oneof"`
}
type Template_Null struct {
	Null *NullTemplate `protobuf:"bytes,502,opt,name=null,oneof"`
}

func (*Template_None) isTemplate_Item() {}
func (*Template_Lxc) isTemplate_Item()  {}
func (*Template_Null) isTemplate_Item() {}

func (m *Template) GetItem() isTemplate_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Template) GetTemplateUri() string {
	if m != nil {
		return m.TemplateUri
	}
	return ""
}

func (m *Template) GetNone() *NoneTemplate {
	if x, ok := m.GetItem().(*Template_None); ok {
		return x.None
	}
	return nil
}

func (m *Template) GetLxc() *LxcTemplate {
	if x, ok := m.GetItem().(*Template_Lxc); ok {
		return x.Lxc
	}
	return nil
}

func (m *Template) GetNull() *NullTemplate {
	if x, ok := m.GetItem().(*Template_Null); ok {
		return x.Null
	}
	return nil
}

func (m *Template) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Template) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Template_OneofMarshaler, _Template_OneofUnmarshaler, _Template_OneofSizer, []interface{}{
		(*Template_None)(nil),
		(*Template_Lxc)(nil),
		(*Template_Null)(nil),
	}
}

func _Template_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Template)
	// Item
	switch x := m.Item.(type) {
	case *Template_None:
		b.EncodeVarint(500<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *Template_Lxc:
		b.EncodeVarint(501<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lxc); err != nil {
			return err
		}
	case *Template_Null:
		b.EncodeVarint(502<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Null); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Template.Item has unexpected type %T", x)
	}
	return nil
}

func _Template_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Template)
	switch tag {
	case 500: // Item.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NoneTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_None{msg}
		return true, err
	case 501: // Item.lxc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LxcTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Lxc{msg}
		return true, err
	case 502: // Item.null
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NullTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Null{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Template_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Template)
	// Item
	switch x := m.Item.(type) {
	case *Template_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(500<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Lxc:
		s := proto.Size(x.Lxc)
		n += proto.SizeVarint(501<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Null:
		s := proto.Size(x.Null)
		n += proto.SizeVarint(502<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NoneTemplate struct {
}

func (m *NoneTemplate) Reset()                    { *m = NoneTemplate{} }
func (m *NoneTemplate) String() string            { return proto.CompactTextString(m) }
func (*NoneTemplate) ProtoMessage()               {}
func (*NoneTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type LxcTemplate struct {
	Vcpu        int32                    `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb    int32                    `protobuf:"varint,2,opt,name=memory_gb,json=memoryGb" json:"memory_gb,omitempty"`
	MinVcpu     int32                    `protobuf:"varint,3,opt,name=min_vcpu,json=minVcpu" json:"min_vcpu,omitempty"`
	MinMemoryGb int32                    `protobuf:"varint,4,opt,name=min_memory_gb,json=minMemoryGb" json:"min_memory_gb,omitempty"`
	LxcImage    *LxcTemplate_Image       `protobuf:"bytes,5,opt,name=lxc_image,json=lxcImage" json:"lxc_image,omitempty"`
	Interfaces  []*LxcTemplate_Interface `protobuf:"bytes,6,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *LxcTemplate) Reset()                    { *m = LxcTemplate{} }
func (m *LxcTemplate) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate) ProtoMessage()               {}
func (*LxcTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LxcTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *LxcTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *LxcTemplate) GetMinVcpu() int32 {
	if m != nil {
		return m.MinVcpu
	}
	return 0
}

func (m *LxcTemplate) GetMinMemoryGb() int32 {
	if m != nil {
		return m.MinMemoryGb
	}
	return 0
}

func (m *LxcTemplate) GetLxcImage() *LxcTemplate_Image {
	if m != nil {
		return m.LxcImage
	}
	return nil
}

func (m *LxcTemplate) GetInterfaces() []*LxcTemplate_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type LxcTemplate_Image struct {
	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url,json=downloadUrl" json:"download_url,omitempty"`
	ChksumType  string `protobuf:"bytes,2,opt,name=chksum_type,json=chksumType" json:"chksum_type,omitempty"`
	Chksum      string `protobuf:"bytes,3,opt,name=chksum" json:"chksum,omitempty"`
}

func (m *LxcTemplate_Image) Reset()                    { *m = LxcTemplate_Image{} }
func (m *LxcTemplate_Image) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Image) ProtoMessage()               {}
func (*LxcTemplate_Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *LxcTemplate_Image) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *LxcTemplate_Image) GetChksumType() string {
	if m != nil {
		return m.ChksumType
	}
	return ""
}

func (m *LxcTemplate_Image) GetChksum() string {
	if m != nil {
		return m.Chksum
	}
	return ""
}

type LxcTemplate_Interface struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Macaddr  string `protobuf:"bytes,2,opt,name=macaddr" json:"macaddr,omitempty"`
	Ipv4Addr string `protobuf:"bytes,3,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
	Bridge   string `protobuf:"bytes,4,opt,name=bridge" json:"bridge,omitempty"`
}

func (m *LxcTemplate_Interface) Reset()                    { *m = LxcTemplate_Interface{} }
func (m *LxcTemplate_Interface) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Interface) ProtoMessage()               {}
func (*LxcTemplate_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *LxcTemplate_Interface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LxcTemplate_Interface) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *LxcTemplate_Interface) GetIpv4Addr() string {
	if m != nil {
		return m.Ipv4Addr
	}
	return ""
}

func (m *LxcTemplate_Interface) GetBridge() string {
	if m != nil {
		return m.Bridge
	}
	return ""
}

type NullTemplate struct {
}

func (m *NullTemplate) Reset()                    { *m = NullTemplate{} }
func (m *NullTemplate) String() string            { return proto.CompactTextString(m) }
func (*NullTemplate) ProtoMessage()               {}
func (*NullTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Instance)(nil), "model.Instance")
	proto.RegisterType((*InstanceState)(nil), "model.InstanceState")
	proto.RegisterType((*Resource)(nil), "model.Resource")
	proto.RegisterType((*Template)(nil), "model.Template")
	proto.RegisterType((*NoneTemplate)(nil), "model.NoneTemplate")
	proto.RegisterType((*LxcTemplate)(nil), "model.LxcTemplate")
	proto.RegisterType((*LxcTemplate_Image)(nil), "model.LxcTemplate.Image")
	proto.RegisterType((*LxcTemplate_Interface)(nil), "model.LxcTemplate.Interface")
	proto.RegisterType((*NullTemplate)(nil), "model.NullTemplate")
	proto.RegisterEnum("model.InstanceState_State", InstanceState_State_name, InstanceState_State_value)
	proto.RegisterEnum("model.Resource_State", Resource_State_name, Resource_State_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x5b, 0x6f, 0xd3, 0x4a,
	0x10, 0xae, 0xe3, 0x38, 0x71, 0xc6, 0x69, 0x8e, 0xb5, 0xe7, 0x22, 0x9f, 0x9c, 0x83, 0xda, 0xfa,
	0x85, 0x42, 0x25, 0x07, 0xb5, 0xf0, 0x80, 0xe0, 0xa5, 0x55, 0xac, 0xd6, 0x12, 0x4d, 0xcb, 0xc6,
	0x01, 0x89, 0x17, 0xcb, 0xb1, 0xb7, 0xae, 0xc5, 0xfa, 0x22, 0x5f, 0x42, 0x2a, 0x7e, 0x0c, 0xff,
	0xa6, 0xff, 0x07, 0x09, 0x90, 0x78, 0x43, 0xbb, 0xb6, 0xdb, 0xf4, 0xc2, 0x03, 0xe2, 0x6d, 0x67,
	0xbe, 0x6f, 0xe6, 0x9b, 0x99, 0xdd, 0x1d, 0x50, 0xa2, 0xc4, 0x27, 0xd4, 0x48, 0xb3, 0xa4, 0x48,
	0x90, 0xc4, 0x8d, 0xe1, 0x8b, 0x20, 0x2c, 0xce, 0xcb, 0xb9, 0xe1, 0x25, 0xd1, 0x28, 0x48, 0xa8,
	0x1b, 0x07, 0x23, 0x8e, 0xcf, 0xcb, 0xb3, 0x51, 0x5a, 0x5c, 0xa4, 0x24, 0x1f, 0x15, 0x61, 0x44,
	0xf2, 0xc2, 0x8d, 0xd2, 0xeb, 0x53, 0x95, 0x43, 0xbf, 0x14, 0x40, 0xb6, 0xe2, 0xbc, 0x70, 0x63,
	0x8f, 0xa0, 0x01, 0xb4, 0x42, 0x5f, 0x13, 0x36, 0x85, 0xed, 0x1e, 0x6e, 0x85, 0x3e, 0xfa, 0x17,
	0xe4, 0x9c, 0xba, 0x0b, 0xe2, 0x84, 0xbe, 0xd6, 0xe2, 0xde, 0x2e, 0xb7, 0x2d, 0x1f, 0x6d, 0x80,
	0x92, 0x91, 0x3c, 0x29, 0x33, 0x8f, 0xa3, 0x22, 0x47, 0xa1, 0x71, 0x59, 0x3e, 0xda, 0x03, 0xa0,
	0x6e, 0x5e, 0x38, 0x79, 0xe1, 0x16, 0x44, 0x6b, 0x6f, 0x0a, 0xdb, 0xca, 0xee, 0x5f, 0x46, 0x55,
	0x7e, 0x23, 0x38, 0x65, 0x18, 0xee, 0x31, 0x1e, 0x3f, 0xa2, 0xe7, 0x00, 0x5e, 0x46, 0xdc, 0x82,
	0xf8, 0x8e, 0x5b, 0x68, 0x12, 0x0f, 0x1a, 0x1a, 0x41, 0x92, 0x04, 0x94, 0x18, 0x4d, 0x53, 0x86,
	0xdd, 0xf4, 0x80, 0x7b, 0x35, 0x7b, 0xbf, 0xd0, 0xbf, 0x0b, 0xb0, 0x7e, 0x23, 0x2f, 0x7a, 0x02,
	0x52, 0x25, 0xce, 0x1a, 0x1a, 0xec, 0x0e, 0xef, 0x13, 0x37, 0xaa, 0x12, 0x2a, 0xe2, 0x2d, 0xf9,
	0xd6, 0xaf, 0xc8, 0x7f, 0x04, 0xa9, 0x52, 0x1d, 0x00, 0x60, 0xf3, 0xd0, 0x9a, 0xda, 0x26, 0x36,
	0xc7, 0xea, 0x1a, 0x02, 0xe8, 0xbc, 0x9e, 0x99, 0x33, 0x73, 0xac, 0x0a, 0xa8, 0x0f, 0xf2, 0xd4,
	0xde, 0xc7, 0xb6, 0x35, 0x39, 0x54, 0x5b, 0x48, 0x81, 0x2e, 0x9e, 0x4d, 0x26, 0xcc, 0x10, 0x2b,
	0xe8, 0xe4, 0xf4, 0x94, 0x59, 0x6d, 0x06, 0x71, 0xcb, 0x1c, 0xab, 0x12, 0x52, 0xa1, 0x3f, 0x3d,
	0x9a, 0xd9, 0x2c, 0x6a, 0x7c, 0xf2, 0x76, 0xa2, 0x76, 0x98, 0x86, 0x6d, 0xe2, 0x63, 0x6b, 0xb2,
	0x6f, 0x9b, 0x63, 0xb5, 0xab, 0x7f, 0x12, 0x40, 0xc6, 0xf5, 0xe8, 0xef, 0x5c, 0xe2, 0x4e, 0x33,
	0x06, 0x91, 0x8f, 0xe1, 0xef, 0x7a, 0x0c, 0x0d, 0xff, 0xe6, 0x04, 0x76, 0x40, 0x2e, 0x48, 0x94,
	0xd2, 0xeb, 0x3b, 0xfb, 0xa3, 0xe6, 0xdb, 0xb5, 0x1b, 0x5f, 0x11, 0xf4, 0x47, 0x3f, 0xeb, 0x59,
	0x85, 0xfe, 0x6c, 0xb2, 0xe2, 0x11, 0xf4, 0xcf, 0x02, 0xc8, 0x4d, 0x06, 0xb4, 0x05, 0xfd, 0x26,
	0x87, 0x53, 0x66, 0x61, 0x5d, 0xab, 0xd2, 0xf8, 0x66, 0x59, 0x88, 0x1e, 0x43, 0x3b, 0x4e, 0x62,
	0xa2, 0x7d, 0x11, 0x79, 0x11, 0x7f, 0xd6, 0x45, 0x4c, 0x92, 0x98, 0x34, 0x69, 0x8e, 0xd6, 0x30,
	0xe7, 0xa0, 0x87, 0x20, 0xd2, 0xa5, 0xa7, 0x7d, 0xad, 0xa8, 0xa8, 0xa6, 0xbe, 0x5a, 0x7a, 0x2b,
	0x4c, 0xc6, 0xe0, 0x49, 0x4b, 0x4a, 0xb5, 0x6f, 0xb7, 0x92, 0x96, 0x94, 0xde, 0x48, 0x5a, 0x52,
	0xfa, 0x1b, 0x4f, 0xe1, 0xa0, 0x03, 0x6d, 0xab, 0x20, 0x91, 0x3e, 0x80, 0xfe, 0x6a, 0xbd, 0xfa,
	0xa5, 0x08, 0xca, 0x4a, 0x55, 0x08, 0x41, 0x7b, 0xe1, 0xa5, 0x25, 0x6f, 0x5f, 0xc2, 0xfc, 0x8c,
	0xfe, 0x83, 0x5e, 0x44, 0xa2, 0x24, 0xbb, 0x70, 0x82, 0x39, 0x57, 0x95, 0xb0, 0x5c, 0x39, 0x0e,
	0xe7, 0xec, 0x3b, 0x46, 0x61, 0xec, 0xf0, 0x20, 0x91, 0x63, 0xdd, 0x28, 0x8c, 0xdf, 0xb0, 0x38,
	0x1d, 0xd6, 0x19, 0x74, 0x1d, 0xdb, 0xe6, 0xb8, 0x12, 0x85, 0xf1, 0x71, 0x13, 0xfe, 0x0c, 0x7a,
	0x74, 0xe9, 0x39, 0x61, 0xe4, 0x06, 0xa4, 0xfe, 0x5b, 0xda, 0xdd, 0x61, 0x19, 0x16, 0xc3, 0xb1,
	0x4c, 0x97, 0x1e, 0x3f, 0xa1, 0x97, 0x00, 0x61, 0x5c, 0x90, 0xec, 0xcc, 0xf5, 0x48, 0xae, 0x75,
	0x36, 0xc5, 0x6d, 0x65, 0xf7, 0xff, 0xfb, 0xe2, 0x1a, 0x12, 0x5e, 0xe1, 0x0f, 0x3d, 0x90, 0xaa,
	0x34, 0x5b, 0xd0, 0xf7, 0x93, 0x0f, 0x31, 0x4d, 0x5c, 0xdf, 0x29, 0x33, 0xda, 0x5c, 0x7a, 0xe3,
	0x9b, 0x65, 0x94, 0xed, 0x14, 0xef, 0xfc, 0x7d, 0x5e, 0x46, 0x0e, 0xdb, 0x5a, 0xf5, 0xc6, 0x81,
	0xca, 0x65, 0x5f, 0xa4, 0x04, 0xfd, 0x03, 0x9d, 0xca, 0xaa, 0xf7, 0x4d, 0x6d, 0x0d, 0x23, 0xe8,
	0x5d, 0xa9, 0xb3, 0xb1, 0xf2, 0xf0, 0x4a, 0x80, 0x9f, 0x91, 0x06, 0xdd, 0xc8, 0xf5, 0x5c, 0xdf,
	0xcf, 0x9a, 0x3d, 0x56, 0x9b, 0x68, 0x08, 0x72, 0x98, 0x2e, 0x9e, 0x72, 0xa8, 0x4a, 0x7a, 0x65,
	0x33, 0xb9, 0x79, 0x16, 0xfa, 0x41, 0xf5, 0x15, 0x7a, 0xb8, 0xb6, 0xf8, 0xc5, 0xae, 0xbc, 0x99,
	0x83, 0x8d, 0x77, 0x0f, 0x56, 0x56, 0xb0, 0xbb, 0xcc, 0xcf, 0x47, 0x49, 0x4a, 0xe2, 0x85, 0xef,
	0x8d, 0xf8, 0x98, 0xe6, 0x1d, 0xfe, 0x60, 0xf6, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x32, 0x81,
	0x23, 0x64, 0xbe, 0x05, 0x00, 0x00,
}

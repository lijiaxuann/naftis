// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/udp_listener_config.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type UdpListenerConfig struct {
	UdpListenerName string `protobuf:"bytes,1,opt,name=udp_listener_name,json=udpListenerName,proto3" json:"udp_listener_name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*UdpListenerConfig_HiddenEnvoyDeprecatedConfig
	//	*UdpListenerConfig_TypedConfig
	ConfigType           isUdpListenerConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UdpListenerConfig) Reset()         { *m = UdpListenerConfig{} }
func (m *UdpListenerConfig) String() string { return proto.CompactTextString(m) }
func (*UdpListenerConfig) ProtoMessage()    {}
func (*UdpListenerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac914914a155255, []int{0}
}

func (m *UdpListenerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpListenerConfig.Unmarshal(m, b)
}
func (m *UdpListenerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpListenerConfig.Marshal(b, m, deterministic)
}
func (m *UdpListenerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpListenerConfig.Merge(m, src)
}
func (m *UdpListenerConfig) XXX_Size() int {
	return xxx_messageInfo_UdpListenerConfig.Size(m)
}
func (m *UdpListenerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpListenerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UdpListenerConfig proto.InternalMessageInfo

func (m *UdpListenerConfig) GetUdpListenerName() string {
	if m != nil {
		return m.UdpListenerName
	}
	return ""
}

type isUdpListenerConfig_ConfigType interface {
	isUdpListenerConfig_ConfigType()
}

type UdpListenerConfig_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type UdpListenerConfig_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*UdpListenerConfig_HiddenEnvoyDeprecatedConfig) isUdpListenerConfig_ConfigType() {}

func (*UdpListenerConfig_TypedConfig) isUdpListenerConfig_ConfigType() {}

func (m *UdpListenerConfig) GetConfigType() isUdpListenerConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *UdpListenerConfig) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*UdpListenerConfig_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *UdpListenerConfig) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*UdpListenerConfig_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UdpListenerConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UdpListenerConfig_HiddenEnvoyDeprecatedConfig)(nil),
		(*UdpListenerConfig_TypedConfig)(nil),
	}
}

type ActiveRawUdpListenerConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveRawUdpListenerConfig) Reset()         { *m = ActiveRawUdpListenerConfig{} }
func (m *ActiveRawUdpListenerConfig) String() string { return proto.CompactTextString(m) }
func (*ActiveRawUdpListenerConfig) ProtoMessage()    {}
func (*ActiveRawUdpListenerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac914914a155255, []int{1}
}

func (m *ActiveRawUdpListenerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveRawUdpListenerConfig.Unmarshal(m, b)
}
func (m *ActiveRawUdpListenerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveRawUdpListenerConfig.Marshal(b, m, deterministic)
}
func (m *ActiveRawUdpListenerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveRawUdpListenerConfig.Merge(m, src)
}
func (m *ActiveRawUdpListenerConfig) XXX_Size() int {
	return xxx_messageInfo_ActiveRawUdpListenerConfig.Size(m)
}
func (m *ActiveRawUdpListenerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveRawUdpListenerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveRawUdpListenerConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UdpListenerConfig)(nil), "envoy.config.listener.v3.UdpListenerConfig")
	proto.RegisterType((*ActiveRawUdpListenerConfig)(nil), "envoy.config.listener.v3.ActiveRawUdpListenerConfig")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/udp_listener_config.proto", fileDescriptor_4ac914914a155255)
}

var fileDescriptor_4ac914914a155255 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xd7, 0x0a, 0x03, 0x33, 0x45, 0x37, 0x44, 0xeb, 0xd4, 0x31, 0x77, 0xd0, 0xe9, 0x21,
	0x91, 0xee, 0x20, 0xee, 0xb6, 0xaa, 0x30, 0x44, 0x64, 0x54, 0x76, 0x2e, 0x59, 0x9b, 0xcd, 0xc0,
	0x96, 0x84, 0x36, 0xad, 0xf6, 0xe6, 0xd1, 0xdf, 0x20, 0xf8, 0x47, 0xbc, 0x0b, 0x5e, 0xfd, 0x47,
	0xd2, 0x34, 0x73, 0x68, 0xdd, 0xb1, 0xfd, 0xde, 0xf7, 0xe9, 0xf7, 0x7c, 0x14, 0xd8, 0x84, 0x25,
	0x3c, 0x45, 0x3e, 0x67, 0x63, 0x3a, 0x41, 0x53, 0x1a, 0x49, 0xc2, 0x48, 0x88, 0x92, 0x0e, 0x8a,
	0x03, 0xe1, 0xcd, 0x9f, 0xbd, 0x7c, 0x0e, 0x45, 0xc8, 0x25, 0xaf, 0x59, 0xaa, 0x03, 0xf5, 0xbb,
	0x79, 0x06, 0x26, 0x9d, 0xfa, 0xee, 0x84, 0xf3, 0xc9, 0x94, 0x20, 0x95, 0x1b, 0xc5, 0x63, 0x84,
	0x59, 0x9a, 0x97, 0xea, 0xfb, 0x7f, 0x47, 0x91, 0x0c, 0x63, 0x5f, 0xea, 0xe9, 0x41, 0x1c, 0x08,
	0x8c, 0x30, 0x63, 0x5c, 0x62, 0x49, 0x39, 0x8b, 0x50, 0x24, 0xb1, 0x8c, 0x23, 0x3d, 0x3e, 0x2c,
	0x8c, 0x13, 0x12, 0x46, 0x94, 0x33, 0xca, 0xf4, 0x52, 0xad, 0x37, 0x13, 0x54, 0x87, 0x81, 0xb8,
	0xd5, 0xdb, 0x5c, 0xaa, 0xe5, 0x6a, 0xa7, 0xa0, 0xfa, 0xcb, 0x83, 0xe1, 0x19, 0xb1, 0x8c, 0xa6,
	0xd1, 0x5e, 0x75, 0x37, 0xe2, 0x45, 0xfa, 0x0e, 0xcf, 0x48, 0x6d, 0x04, 0x1a, 0x0f, 0x34, 0x08,
	0x08, 0xf3, 0x94, 0x9f, 0x17, 0x10, 0x11, 0x12, 0x1f, 0x4b, 0x12, 0x68, 0x7d, 0xcb, 0x6c, 0x1a,
	0xed, 0x8a, 0xbd, 0x03, 0x73, 0x15, 0x38, 0x57, 0x81, 0xf7, 0x4a, 0xc5, 0x31, 0x2d, 0xa3, 0x5f,
	0x72, 0xf7, 0x72, 0xc8, 0x75, 0xc6, 0xb8, 0xfa, 0x41, 0xe8, 0x7d, 0x2e, 0xc0, 0x9a, 0x4c, 0xc5,
	0x82, 0xb8, 0xa2, 0x88, 0x5b, 0x05, 0x62, 0x8f, 0xa5, 0xfd, 0x92, 0x5b, 0x51, 0xd9, 0xbc, 0xda,
	0x85, 0xaf, 0x1f, 0x2f, 0x8d, 0x13, 0x70, 0x9c, 0x1f, 0x1f, 0x0b, 0x0a, 0x13, 0x7b, 0x71, 0xfc,
	0x82, 0xba, 0xb3, 0x0e, 0x2a, 0xf9, 0x47, 0xbc, 0x8c, 0xd2, 0x1a, 0x82, 0x7a, 0xcf, 0x97, 0x34,
	0x21, 0x2e, 0x7e, 0x2c, 0x84, 0xbb, 0xe7, 0x19, 0xdc, 0x06, 0x67, 0xff, 0xc3, 0x97, 0x17, 0x9d,
	0x9b, 0xf7, 0xe7, 0xcf, 0xaf, 0xb2, 0xb9, 0x69, 0x82, 0x23, 0xca, 0xa1, 0xaa, 0x8b, 0x90, 0x3f,
	0xa5, 0x70, 0xd9, 0x3f, 0xe2, 0x6c, 0x17, 0x20, 0x83, 0xcc, 0x7a, 0x60, 0x8c, 0xca, 0x4a, 0xbf,
	0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xee, 0xda, 0x84, 0x44, 0x94, 0x02, 0x00, 0x00,
}

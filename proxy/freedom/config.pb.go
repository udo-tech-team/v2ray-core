package freedom

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_protocol1 "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config_DomainStrategy int32

const (
	Config_AS_IS  Config_DomainStrategy = 0
	Config_USE_IP Config_DomainStrategy = 1
)

var Config_DomainStrategy_name = map[int32]string{
	0: "AS_IS",
	1: "USE_IP",
}
var Config_DomainStrategy_value = map[string]int32{
	"AS_IS":  0,
	"USE_IP": 1,
}

func (x Config_DomainStrategy) String() string {
	return proto.EnumName(Config_DomainStrategy_name, int32(x))
}
func (Config_DomainStrategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type DestinationOverride struct {
	Server *v2ray_core_common_protocol1.ServerEndpoint `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
}

func (m *DestinationOverride) Reset()                    { *m = DestinationOverride{} }
func (m *DestinationOverride) String() string            { return proto.CompactTextString(m) }
func (*DestinationOverride) ProtoMessage()               {}
func (*DestinationOverride) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DestinationOverride) GetServer() *v2ray_core_common_protocol1.ServerEndpoint {
	if m != nil {
		return m.Server
	}
	return nil
}

type Config struct {
	DomainStrategy      Config_DomainStrategy `protobuf:"varint,1,opt,name=domain_strategy,json=domainStrategy,enum=v2ray.core.proxy.freedom.Config_DomainStrategy" json:"domain_strategy,omitempty"`
	Timeout             uint32                `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	DestinationOverride *DestinationOverride  `protobuf:"bytes,3,opt,name=destination_override,json=destinationOverride" json:"destination_override,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetDomainStrategy() Config_DomainStrategy {
	if m != nil {
		return m.DomainStrategy
	}
	return Config_AS_IS
}

func (m *Config) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Config) GetDestinationOverride() *DestinationOverride {
	if m != nil {
		return m.DestinationOverride
	}
	return nil
}

func init() {
	proto.RegisterType((*DestinationOverride)(nil), "v2ray.core.proxy.freedom.DestinationOverride")
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.freedom.Config")
	proto.RegisterEnum("v2ray.core.proxy.freedom.Config_DomainStrategy", Config_DomainStrategy_name, Config_DomainStrategy_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/freedom/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0xc5, 0x14, 0x47, 0xac, 0x65, 0xeb, 0x21, 0x88, 0x87, 0xd2, 0x8b, 0x55, 0x70,
	0x23, 0xf1, 0x09, 0xec, 0x1f, 0xa1, 0x27, 0x4b, 0x82, 0xa2, 0x5e, 0x62, 0xcc, 0x4e, 0xcb, 0x82,
	0xd9, 0x09, 0x9b, 0xb5, 0x98, 0x27, 0xf0, 0x5d, 0x7c, 0x4a, 0x71, 0x93, 0xa2, 0x55, 0x73, 0xdb,
	0x99, 0xfd, 0x7d, 0xdf, 0xcc, 0x37, 0x70, 0xba, 0x0a, 0x74, 0x52, 0xf2, 0x94, 0x32, 0x3f, 0x25,
	0x8d, 0x7e, 0xae, 0xe9, 0xad, 0xf4, 0x17, 0x1a, 0x51, 0xd8, 0x96, 0x5a, 0xc8, 0x25, 0xcf, 0x35,
	0x19, 0x62, 0xde, 0x1a, 0xd5, 0xc8, 0x2d, 0xc6, 0x6b, 0xec, 0xe8, 0xe2, 0x97, 0x49, 0x4a, 0x59,
	0x46, 0xca, 0xb7, 0xb2, 0x94, 0x5e, 0xfc, 0x02, 0xf5, 0x0a, 0x75, 0x5c, 0xe4, 0x98, 0x56, 0x5e,
	0x83, 0x07, 0xe8, 0x4d, 0xb0, 0x30, 0x52, 0x25, 0x46, 0x92, 0xba, 0x59, 0xa1, 0xd6, 0x52, 0x20,
	0x1b, 0x81, 0x5b, 0xb1, 0x9e, 0xd3, 0x77, 0x86, 0x7b, 0xc1, 0x19, 0xff, 0x31, 0xb3, 0x72, 0xe5,
	0x6b, 0x57, 0x1e, 0x59, 0x72, 0xaa, 0x44, 0x4e, 0x52, 0x99, 0xb0, 0x56, 0x0e, 0xde, 0x5b, 0xe0,
	0x8e, 0xed, 0xde, 0xec, 0x1e, 0x0e, 0x04, 0x65, 0x89, 0x54, 0x71, 0x61, 0x74, 0x62, 0x70, 0x59,
	0x5a, 0xdf, 0x4e, 0xe0, 0xf3, 0xa6, 0x2c, 0xbc, 0x92, 0xf2, 0x89, 0xd5, 0x45, 0xb5, 0x2c, 0xec,
	0x88, 0x8d, 0x9a, 0x79, 0xd0, 0x36, 0x32, 0x43, 0x7a, 0x35, 0x5e, 0xab, 0xef, 0x0c, 0xf7, 0xc3,
	0x75, 0xc9, 0x9e, 0xe0, 0x50, 0x7c, 0x27, 0x8b, 0xa9, 0x8e, 0xe6, 0x6d, 0xdb, 0x40, 0xe7, 0xcd,
	0x83, 0xff, 0xb9, 0x47, 0xd8, 0x13, 0x7f, 0x9b, 0x83, 0x13, 0xe8, 0x6c, 0x6e, 0xc7, 0x76, 0x61,
	0xe7, 0x2a, 0x8a, 0x67, 0x51, 0x77, 0x8b, 0x01, 0xb8, 0xb7, 0xd1, 0x34, 0x9e, 0xcd, 0xbb, 0xce,
	0x68, 0x02, 0xc7, 0x29, 0x65, 0x8d, 0x13, 0xe7, 0xce, 0x63, 0xbb, 0x7e, 0x7e, 0xb4, 0xbc, 0xbb,
	0x20, 0x4c, 0x4a, 0x3e, 0xfe, 0xa2, 0xe6, 0x96, 0xba, 0xae, 0xbe, 0x9e, 0x5d, 0x7b, 0xf0, 0xcb,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xb3, 0x84, 0x2e, 0x2a, 0x02, 0x00, 0x00,
}

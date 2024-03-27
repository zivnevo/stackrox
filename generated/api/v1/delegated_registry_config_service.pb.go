// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/delegated_registry_config_service.proto

package v1

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

type DelegatedRegistryConfig_EnabledFor int32

const (
	// Scan all images via central services except for images from the OCP integrated registry
	DelegatedRegistryConfig_NONE DelegatedRegistryConfig_EnabledFor = 0
	// Scan all images via the secured clusters
	DelegatedRegistryConfig_ALL DelegatedRegistryConfig_EnabledFor = 1
	// Scan images that match `registries` or are from the OCP integrated registry via the secured clusters
	// otherwise scan via central services
	DelegatedRegistryConfig_SPECIFIC DelegatedRegistryConfig_EnabledFor = 2
)

var DelegatedRegistryConfig_EnabledFor_name = map[int32]string{
	0: "NONE",
	1: "ALL",
	2: "SPECIFIC",
}

var DelegatedRegistryConfig_EnabledFor_value = map[string]int32{
	"NONE":     0,
	"ALL":      1,
	"SPECIFIC": 2,
}

func (x DelegatedRegistryConfig_EnabledFor) String() string {
	return proto.EnumName(DelegatedRegistryConfig_EnabledFor_name, int32(x))
}

func (DelegatedRegistryConfig_EnabledFor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_70904b65c1aaa73b, []int{2, 0}
}

type DelegatedRegistryClustersResponse struct {
	Clusters             []*DelegatedRegistryCluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DelegatedRegistryClustersResponse) Reset()         { *m = DelegatedRegistryClustersResponse{} }
func (m *DelegatedRegistryClustersResponse) String() string { return proto.CompactTextString(m) }
func (*DelegatedRegistryClustersResponse) ProtoMessage()    {}
func (*DelegatedRegistryClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70904b65c1aaa73b, []int{0}
}
func (m *DelegatedRegistryClustersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatedRegistryClustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatedRegistryClustersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatedRegistryClustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatedRegistryClustersResponse.Merge(m, src)
}
func (m *DelegatedRegistryClustersResponse) XXX_Size() int {
	return m.Size()
}
func (m *DelegatedRegistryClustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatedRegistryClustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatedRegistryClustersResponse proto.InternalMessageInfo

func (m *DelegatedRegistryClustersResponse) GetClusters() []*DelegatedRegistryCluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *DelegatedRegistryClustersResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *DelegatedRegistryClustersResponse) Clone() *DelegatedRegistryClustersResponse {
	if m == nil {
		return nil
	}
	cloned := new(DelegatedRegistryClustersResponse)
	*cloned = *m

	if m.Clusters != nil {
		cloned.Clusters = make([]*DelegatedRegistryCluster, len(m.Clusters))
		for idx, v := range m.Clusters {
			cloned.Clusters[idx] = v.Clone()
		}
	}
	return cloned
}

type DelegatedRegistryCluster struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsValid              bool     `protobuf:"varint,3,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegatedRegistryCluster) Reset()         { *m = DelegatedRegistryCluster{} }
func (m *DelegatedRegistryCluster) String() string { return proto.CompactTextString(m) }
func (*DelegatedRegistryCluster) ProtoMessage()    {}
func (*DelegatedRegistryCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_70904b65c1aaa73b, []int{1}
}
func (m *DelegatedRegistryCluster) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatedRegistryCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatedRegistryCluster.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatedRegistryCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatedRegistryCluster.Merge(m, src)
}
func (m *DelegatedRegistryCluster) XXX_Size() int {
	return m.Size()
}
func (m *DelegatedRegistryCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatedRegistryCluster.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatedRegistryCluster proto.InternalMessageInfo

func (m *DelegatedRegistryCluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DelegatedRegistryCluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DelegatedRegistryCluster) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *DelegatedRegistryCluster) MessageClone() proto.Message {
	return m.Clone()
}
func (m *DelegatedRegistryCluster) Clone() *DelegatedRegistryCluster {
	if m == nil {
		return nil
	}
	cloned := new(DelegatedRegistryCluster)
	*cloned = *m

	return cloned
}

// DelegatedRegistryConfig determines if and where scan requests are delegated to, such as kept in
// central services or sent to particular secured clusters.
type DelegatedRegistryConfig struct {
	// Determines if delegation is enabled for no registries, all registries, or specific registries
	EnabledFor DelegatedRegistryConfig_EnabledFor `protobuf:"varint,1,opt,name=enabled_for,json=enabledFor,proto3,enum=v1.DelegatedRegistryConfig_EnabledFor" json:"enabled_for,omitempty"`
	// The default cluster to delegate ad-hoc requests to if no match found in registries, not used
	// if `enabled for` is NONE
	DefaultClusterId string `protobuf:"bytes,2,opt,name=default_cluster_id,json=defaultClusterId,proto3" json:"default_cluster_id,omitempty"`
	// If `enabled for` is NONE registries has no effect.
	//
	// If `ALL` registries directs ad-hoc requests to the specified secured clusters if the path matches.
	//
	// If `SPECIFIC` registries directs ad-hoc requests to the specified secured clusters just like with `ALL`,
	// but in addition images that match the specified paths will be scanned locally by the secured clusters
	// (images from the OCP integrated registry are always scanned locally). Images that do not match a path
	// will be scanned via central services
	Registries           []*DelegatedRegistryConfig_DelegatedRegistry `protobuf:"bytes,3,rep,name=registries,proto3" json:"registries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *DelegatedRegistryConfig) Reset()         { *m = DelegatedRegistryConfig{} }
func (m *DelegatedRegistryConfig) String() string { return proto.CompactTextString(m) }
func (*DelegatedRegistryConfig) ProtoMessage()    {}
func (*DelegatedRegistryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_70904b65c1aaa73b, []int{2}
}
func (m *DelegatedRegistryConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatedRegistryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatedRegistryConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatedRegistryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatedRegistryConfig.Merge(m, src)
}
func (m *DelegatedRegistryConfig) XXX_Size() int {
	return m.Size()
}
func (m *DelegatedRegistryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatedRegistryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatedRegistryConfig proto.InternalMessageInfo

func (m *DelegatedRegistryConfig) GetEnabledFor() DelegatedRegistryConfig_EnabledFor {
	if m != nil {
		return m.EnabledFor
	}
	return DelegatedRegistryConfig_NONE
}

func (m *DelegatedRegistryConfig) GetDefaultClusterId() string {
	if m != nil {
		return m.DefaultClusterId
	}
	return ""
}

func (m *DelegatedRegistryConfig) GetRegistries() []*DelegatedRegistryConfig_DelegatedRegistry {
	if m != nil {
		return m.Registries
	}
	return nil
}

func (m *DelegatedRegistryConfig) MessageClone() proto.Message {
	return m.Clone()
}
func (m *DelegatedRegistryConfig) Clone() *DelegatedRegistryConfig {
	if m == nil {
		return nil
	}
	cloned := new(DelegatedRegistryConfig)
	*cloned = *m

	if m.Registries != nil {
		cloned.Registries = make([]*DelegatedRegistryConfig_DelegatedRegistry, len(m.Registries))
		for idx, v := range m.Registries {
			cloned.Registries[idx] = v.Clone()
		}
	}
	return cloned
}

type DelegatedRegistryConfig_DelegatedRegistry struct {
	// Registry + optional path, ie: quay.example.com/prod
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ID of the cluster to delegate ad-hoc requests to
	ClusterId            string   `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) Reset() {
	*m = DelegatedRegistryConfig_DelegatedRegistry{}
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) String() string {
	return proto.CompactTextString(m)
}
func (*DelegatedRegistryConfig_DelegatedRegistry) ProtoMessage() {}
func (*DelegatedRegistryConfig_DelegatedRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_70904b65c1aaa73b, []int{2, 0}
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatedRegistryConfig_DelegatedRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatedRegistryConfig_DelegatedRegistry.Merge(m, src)
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) XXX_Size() int {
	return m.Size()
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatedRegistryConfig_DelegatedRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatedRegistryConfig_DelegatedRegistry proto.InternalMessageInfo

func (m *DelegatedRegistryConfig_DelegatedRegistry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) MessageClone() proto.Message {
	return m.Clone()
}
func (m *DelegatedRegistryConfig_DelegatedRegistry) Clone() *DelegatedRegistryConfig_DelegatedRegistry {
	if m == nil {
		return nil
	}
	cloned := new(DelegatedRegistryConfig_DelegatedRegistry)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("v1.DelegatedRegistryConfig_EnabledFor", DelegatedRegistryConfig_EnabledFor_name, DelegatedRegistryConfig_EnabledFor_value)
	proto.RegisterType((*DelegatedRegistryClustersResponse)(nil), "v1.DelegatedRegistryClustersResponse")
	proto.RegisterType((*DelegatedRegistryCluster)(nil), "v1.DelegatedRegistryCluster")
	proto.RegisterType((*DelegatedRegistryConfig)(nil), "v1.DelegatedRegistryConfig")
	proto.RegisterType((*DelegatedRegistryConfig_DelegatedRegistry)(nil), "v1.DelegatedRegistryConfig.DelegatedRegistry")
}

func init() {
	proto.RegisterFile("api/v1/delegated_registry_config_service.proto", fileDescriptor_70904b65c1aaa73b)
}

var fileDescriptor_70904b65c1aaa73b = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0x8a, 0x26, 0x93, 0xaa, 0x0a, 0x73, 0xc1, 0xa4, 0x6d, 0x14, 0x0c, 0x44, 0x11,
	0x6a, 0x1d, 0x25, 0x5c, 0x10, 0x37, 0x08, 0x49, 0x15, 0xa9, 0x14, 0xe4, 0x0a, 0x54, 0x90, 0x90,
	0xb5, 0xcd, 0x6e, 0xc2, 0x0a, 0xc7, 0x6b, 0x79, 0xb7, 0x16, 0xbd, 0xf2, 0x17, 0xb8, 0xf0, 0x93,
	0x38, 0x22, 0x21, 0x71, 0xe1, 0x82, 0x02, 0x3f, 0x04, 0x79, 0xbd, 0x09, 0x45, 0x69, 0x7c, 0x1b,
	0xcf, 0xd7, 0x9b, 0xf7, 0xfc, 0x16, 0x3c, 0x12, 0xf3, 0x4e, 0xda, 0xed, 0x50, 0x16, 0xb2, 0x29,
	0x51, 0x8c, 0x06, 0x09, 0x9b, 0x72, 0xa9, 0x92, 0xcb, 0x60, 0x2c, 0xa2, 0x09, 0x9f, 0x06, 0x92,
	0x25, 0x29, 0x1f, 0x33, 0x2f, 0x4e, 0x84, 0x12, 0x68, 0xa7, 0xdd, 0x3a, 0x9a, 0x19, 0x36, 0x8b,
	0xd5, 0x65, 0x9e, 0xaf, 0xef, 0x4d, 0x85, 0x98, 0x86, 0xac, 0x93, 0x95, 0x48, 0x14, 0x09, 0x45,
	0x14, 0x17, 0x91, 0xcc, 0xab, 0xee, 0x3b, 0xb8, 0xf3, 0x6c, 0x01, 0xe0, 0x9b, 0xfd, 0xfd, 0xf0,
	0x42, 0x2a, 0x96, 0x48, 0x9f, 0xc9, 0x58, 0x44, 0x92, 0xe1, 0x23, 0x28, 0x8f, 0x4d, 0xce, 0xb1,
	0x9a, 0xa5, 0x76, 0xb5, 0xb7, 0xe7, 0xa5, 0x5d, 0x6f, 0xdd, 0xa0, 0xbf, 0xec, 0x76, 0xdf, 0x80,
	0xb3, 0xae, 0x0b, 0x77, 0xc0, 0xe6, 0xd4, 0xb1, 0x9a, 0x56, 0xbb, 0xe2, 0xdb, 0x9c, 0x22, 0xc2,
	0x66, 0x44, 0x66, 0xcc, 0xb1, 0x75, 0x46, 0xc7, 0x78, 0x1b, 0xca, 0x5c, 0x06, 0x29, 0x09, 0x39,
	0x75, 0x4a, 0x4d, 0xab, 0x5d, 0xf6, 0xb7, 0xb8, 0x7c, 0x9d, 0x7d, 0xba, 0x3f, 0x6c, 0xb8, 0xb5,
	0xba, 0x5b, 0x2b, 0x83, 0x47, 0x50, 0x65, 0x11, 0x39, 0x0f, 0x19, 0x0d, 0x26, 0x22, 0xd1, 0x18,
	0x3b, 0xbd, 0xd6, 0xf5, 0x37, 0xeb, 0x09, 0x6f, 0x90, 0xb7, 0x0f, 0x45, 0xe2, 0x03, 0x5b, 0xc6,
	0x78, 0x00, 0x48, 0xd9, 0x84, 0x5c, 0x84, 0x2a, 0x30, 0x9c, 0x02, 0x4e, 0xcd, 0x85, 0x35, 0x53,
	0x31, 0x7c, 0x46, 0x14, 0x9f, 0x03, 0x98, 0x7f, 0xc4, 0x99, 0x74, 0x4a, 0x5a, 0xa9, 0xc3, 0x22,
	0xd4, 0x95, 0xbc, 0x7f, 0x65, 0x41, 0x7d, 0x08, 0x37, 0x57, 0x1a, 0x32, 0x95, 0x62, 0xa2, 0xde,
	0x1b, 0xdd, 0x74, 0x8c, 0xfb, 0x00, 0x2b, 0xd7, 0x55, 0xc6, 0x8b, 0xb3, 0xdc, 0x43, 0x80, 0x7f,
	0xf4, 0xb0, 0x0c, 0x9b, 0x27, 0x2f, 0x4e, 0x06, 0xb5, 0x0d, 0xdc, 0x82, 0xd2, 0x93, 0xe3, 0xe3,
	0x9a, 0x85, 0xdb, 0x50, 0x3e, 0x7d, 0x39, 0xe8, 0x8f, 0x86, 0xa3, 0x7e, 0xcd, 0xee, 0xfd, 0xb4,
	0xa1, 0xb1, 0xe6, 0xe0, 0xd3, 0xdc, 0x71, 0x78, 0x06, 0x95, 0x23, 0xa6, 0x8c, 0xd8, 0x95, 0x8c,
	0xe1, 0x20, 0x73, 0x5c, 0x7d, 0xb7, 0x80, 0xac, 0x7b, 0xf7, 0xd3, 0xf7, 0x3f, 0x9f, 0xed, 0x7d,
	0xdc, 0xfd, 0xcf, 0xd5, 0x0b, 0x53, 0xe7, 0x9e, 0xc6, 0x19, 0x54, 0xb3, 0xcd, 0xc6, 0x3f, 0x57,
	0x77, 0xdf, 0x2f, 0xb2, 0xdc, 0xd2, 0xab, 0xee, 0x81, 0x46, 0x69, 0xe1, 0xbd, 0x02, 0x94, 0xce,
	0xc2, 0x9f, 0x18, 0xc3, 0xf6, 0xab, 0x98, 0x12, 0xc5, 0x0c, 0x97, 0x22, 0x02, 0xc5, 0xec, 0x5a,
	0x1a, 0xb7, 0x59, 0x2f, 0x62, 0xf7, 0xd8, 0x7a, 0xf0, 0xd4, 0xfb, 0x3a, 0x6f, 0x58, 0xdf, 0xe6,
	0x0d, 0xeb, 0xd7, 0xbc, 0x61, 0x7d, 0xf9, 0xdd, 0xd8, 0x00, 0x87, 0x0b, 0x4f, 0x2a, 0x32, 0xfe,
	0x90, 0x88, 0x8f, 0xf9, 0xab, 0xcc, 0x9e, 0xbe, 0x97, 0x76, 0xdf, 0xda, 0x69, 0xf7, 0xcc, 0x3a,
	0xbf, 0xa1, 0x73, 0x0f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xac, 0x24, 0x7f, 0x36, 0x11, 0x04,
	0x00, 0x00,
}

func (m *DelegatedRegistryClustersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedRegistryClustersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatedRegistryClustersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Clusters) > 0 {
		for iNdEx := len(m.Clusters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Clusters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DelegatedRegistryCluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedRegistryCluster) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatedRegistryCluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsValid {
		i--
		if m.IsValid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegatedRegistryConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedRegistryConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatedRegistryConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Registries) > 0 {
		for iNdEx := len(m.Registries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DefaultClusterId) > 0 {
		i -= len(m.DefaultClusterId)
		copy(dAtA[i:], m.DefaultClusterId)
		i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(len(m.DefaultClusterId)))
		i--
		dAtA[i] = 0x12
	}
	if m.EnabledFor != 0 {
		i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(m.EnabledFor))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintDelegatedRegistryConfigService(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegatedRegistryConfigService(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegatedRegistryConfigService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelegatedRegistryClustersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Clusters) > 0 {
		for _, e := range m.Clusters {
			l = e.Size()
			n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DelegatedRegistryCluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
	}
	if m.IsValid {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DelegatedRegistryConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnabledFor != 0 {
		n += 1 + sovDelegatedRegistryConfigService(uint64(m.EnabledFor))
	}
	l = len(m.DefaultClusterId)
	if l > 0 {
		n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
	}
	if len(m.Registries) > 0 {
		for _, e := range m.Registries {
			l = e.Size()
			n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DelegatedRegistryConfig_DelegatedRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovDelegatedRegistryConfigService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDelegatedRegistryConfigService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegatedRegistryConfigService(x uint64) (n int) {
	return sovDelegatedRegistryConfigService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelegatedRegistryClustersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegatedRegistryConfigService
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
			return fmt.Errorf("proto: DelegatedRegistryClustersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedRegistryClustersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
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
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, &DelegatedRegistryCluster{})
			if err := m.Clusters[len(m.Clusters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegatedRegistryConfigService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
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
func (m *DelegatedRegistryCluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegatedRegistryConfigService
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
			return fmt.Errorf("proto: DelegatedRegistryCluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedRegistryCluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsValid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsValid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDelegatedRegistryConfigService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
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
func (m *DelegatedRegistryConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegatedRegistryConfigService
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
			return fmt.Errorf("proto: DelegatedRegistryConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedRegistryConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnabledFor", wireType)
			}
			m.EnabledFor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnabledFor |= DelegatedRegistryConfig_EnabledFor(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
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
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registries = append(m.Registries, &DelegatedRegistryConfig_DelegatedRegistry{})
			if err := m.Registries[len(m.Registries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegatedRegistryConfigService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
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
func (m *DelegatedRegistryConfig_DelegatedRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegatedRegistryConfigService
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
			return fmt.Errorf("proto: DelegatedRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegatedRegistryConfigService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegatedRegistryConfigService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegatedRegistryConfigService
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
func skipDelegatedRegistryConfigService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegatedRegistryConfigService
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
					return 0, ErrIntOverflowDelegatedRegistryConfigService
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
					return 0, ErrIntOverflowDelegatedRegistryConfigService
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
				return 0, ErrInvalidLengthDelegatedRegistryConfigService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegatedRegistryConfigService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegatedRegistryConfigService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegatedRegistryConfigService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegatedRegistryConfigService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegatedRegistryConfigService = fmt.Errorf("proto: unexpected end of group")
)

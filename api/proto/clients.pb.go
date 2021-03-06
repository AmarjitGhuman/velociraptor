// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clients.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type ApiClient_IPAddressClass int32

const (
	ApiClient_UNKNOWN  ApiClient_IPAddressClass = 0
	ApiClient_INTERNAL ApiClient_IPAddressClass = 1
	ApiClient_EXTERNAL ApiClient_IPAddressClass = 2
	ApiClient_VPN      ApiClient_IPAddressClass = 3
)

var ApiClient_IPAddressClass_name = map[int32]string{
	0: "UNKNOWN",
	1: "INTERNAL",
	2: "EXTERNAL",
	3: "VPN",
}

var ApiClient_IPAddressClass_value = map[string]int32{
	"UNKNOWN":  0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
	"VPN":      3,
}

func (x ApiClient_IPAddressClass) String() string {
	return proto.EnumName(ApiClient_IPAddressClass_name, int32(x))
}

func (ApiClient_IPAddressClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2, 0}
}

type SearchClientsRequest_QueryType int32

const (
	SearchClientsRequest_VALUE SearchClientsRequest_QueryType = 0
	SearchClientsRequest_KEY   SearchClientsRequest_QueryType = 1
)

var SearchClientsRequest_QueryType_name = map[int32]string{
	0: "VALUE",
	1: "KEY",
}

var SearchClientsRequest_QueryType_value = map[string]int32{
	"VALUE": 0,
	"KEY":   1,
}

func (x SearchClientsRequest_QueryType) String() string {
	return proto.EnumName(SearchClientsRequest_QueryType_name, int32(x))
}

func (SearchClientsRequest_QueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3, 0}
}

// GRR uses an int for client_version which is difficult to use
// semantic versioning. We use a string instead which represents the
// commit number.
type AgentInformation struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BuildTime            string   `protobuf:"bytes,3,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentInformation) Reset()         { *m = AgentInformation{} }
func (m *AgentInformation) String() string { return proto.CompactTextString(m) }
func (*AgentInformation) ProtoMessage()    {}
func (*AgentInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{0}
}

func (m *AgentInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentInformation.Unmarshal(m, b)
}
func (m *AgentInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentInformation.Marshal(b, m, deterministic)
}
func (m *AgentInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentInformation.Merge(m, src)
}
func (m *AgentInformation) XXX_Size() int {
	return xxx_messageInfo_AgentInformation.Size(m)
}
func (m *AgentInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentInformation.DiscardUnknown(m)
}

var xxx_messageInfo_AgentInformation proto.InternalMessageInfo

func (m *AgentInformation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AgentInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentInformation) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

// Message to carry uname information.
type Uname struct {
	System               string   `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Release              string   `protobuf:"bytes,3,opt,name=release,proto3" json:"release,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Machine              string   `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
	Kernel               string   `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Fqdn                 string   `protobuf:"bytes,7,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	InstallDate          uint64   `protobuf:"varint,8,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`
	LibcVer              string   `protobuf:"bytes,9,opt,name=libc_ver,json=libcVer,proto3" json:"libc_ver,omitempty"`
	Architecture         string   `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uname) Reset()         { *m = Uname{} }
func (m *Uname) String() string { return proto.CompactTextString(m) }
func (*Uname) ProtoMessage()    {}
func (*Uname) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{1}
}

func (m *Uname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uname.Unmarshal(m, b)
}
func (m *Uname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uname.Marshal(b, m, deterministic)
}
func (m *Uname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uname.Merge(m, src)
}
func (m *Uname) XXX_Size() int {
	return xxx_messageInfo_Uname.Size(m)
}
func (m *Uname) XXX_DiscardUnknown() {
	xxx_messageInfo_Uname.DiscardUnknown(m)
}

var xxx_messageInfo_Uname proto.InternalMessageInfo

func (m *Uname) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Uname) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Uname) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Uname) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Uname) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *Uname) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *Uname) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Uname) GetInstallDate() uint64 {
	if m != nil {
		return m.InstallDate
	}
	return 0
}

func (m *Uname) GetLibcVer() string {
	if m != nil {
		return m.LibcVer
	}
	return ""
}

func (m *Uname) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

// Describe a client. We fill in some metadata about the client but
// this is by no means exhaustive.
type ApiClient struct {
	ClientId              string                   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AgentInformation      *AgentInformation        `protobuf:"bytes,2,opt,name=agent_information,json=agentInformation,proto3" json:"agent_information,omitempty"`
	OsInfo                *Uname                   `protobuf:"bytes,3,opt,name=os_info,json=osInfo,proto3" json:"os_info,omitempty"`
	FirstSeenAt           uint64                   `protobuf:"varint,6,opt,name=first_seen_at,json=firstSeenAt,proto3" json:"first_seen_at,omitempty"`
	LastSeenAt            uint64                   `protobuf:"varint,7,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	LastBootedAt          uint64                   `protobuf:"varint,8,opt,name=last_booted_at,json=lastBootedAt,proto3" json:"last_booted_at,omitempty"`
	LastClock             uint64                   `protobuf:"varint,9,opt,name=last_clock,json=lastClock,proto3" json:"last_clock,omitempty"`
	LastCrashAt           uint64                   `protobuf:"varint,10,opt,name=last_crash_at,json=lastCrashAt,proto3" json:"last_crash_at,omitempty"`
	LastIp                string                   `protobuf:"bytes,16,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	LastInterrogateFlowId string                   `protobuf:"bytes,19,opt,name=last_interrogate_flow_id,json=lastInterrogateFlowId,proto3" json:"last_interrogate_flow_id,omitempty"`
	LastIpClass           ApiClient_IPAddressClass `protobuf:"varint,17,opt,name=last_ip_class,json=lastIpClass,proto3,enum=proto.ApiClient_IPAddressClass" json:"last_ip_class,omitempty"`
	Labels                []string                 `protobuf:"bytes,18,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *ApiClient) Reset()         { *m = ApiClient{} }
func (m *ApiClient) String() string { return proto.CompactTextString(m) }
func (*ApiClient) ProtoMessage()    {}
func (*ApiClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2}
}

func (m *ApiClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiClient.Unmarshal(m, b)
}
func (m *ApiClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiClient.Marshal(b, m, deterministic)
}
func (m *ApiClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiClient.Merge(m, src)
}
func (m *ApiClient) XXX_Size() int {
	return xxx_messageInfo_ApiClient.Size(m)
}
func (m *ApiClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiClient.DiscardUnknown(m)
}

var xxx_messageInfo_ApiClient proto.InternalMessageInfo

func (m *ApiClient) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ApiClient) GetAgentInformation() *AgentInformation {
	if m != nil {
		return m.AgentInformation
	}
	return nil
}

func (m *ApiClient) GetOsInfo() *Uname {
	if m != nil {
		return m.OsInfo
	}
	return nil
}

func (m *ApiClient) GetFirstSeenAt() uint64 {
	if m != nil {
		return m.FirstSeenAt
	}
	return 0
}

func (m *ApiClient) GetLastSeenAt() uint64 {
	if m != nil {
		return m.LastSeenAt
	}
	return 0
}

func (m *ApiClient) GetLastBootedAt() uint64 {
	if m != nil {
		return m.LastBootedAt
	}
	return 0
}

func (m *ApiClient) GetLastClock() uint64 {
	if m != nil {
		return m.LastClock
	}
	return 0
}

func (m *ApiClient) GetLastCrashAt() uint64 {
	if m != nil {
		return m.LastCrashAt
	}
	return 0
}

func (m *ApiClient) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

func (m *ApiClient) GetLastInterrogateFlowId() string {
	if m != nil {
		return m.LastInterrogateFlowId
	}
	return ""
}

func (m *ApiClient) GetLastIpClass() ApiClient_IPAddressClass {
	if m != nil {
		return m.LastIpClass
	}
	return ApiClient_UNKNOWN
}

func (m *ApiClient) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SearchClientsRequest struct {
	Offset               uint64                         `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint64                         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Query                string                         `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	NameOnly             bool                           `protobuf:"varint,4,opt,name=name_only,json=nameOnly,proto3" json:"name_only,omitempty"`
	Type                 SearchClientsRequest_QueryType `protobuf:"varint,5,opt,name=type,proto3,enum=proto.SearchClientsRequest_QueryType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SearchClientsRequest) Reset()         { *m = SearchClientsRequest{} }
func (m *SearchClientsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchClientsRequest) ProtoMessage()    {}
func (*SearchClientsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3}
}

func (m *SearchClientsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClientsRequest.Unmarshal(m, b)
}
func (m *SearchClientsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClientsRequest.Marshal(b, m, deterministic)
}
func (m *SearchClientsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClientsRequest.Merge(m, src)
}
func (m *SearchClientsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchClientsRequest.Size(m)
}
func (m *SearchClientsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClientsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClientsRequest proto.InternalMessageInfo

func (m *SearchClientsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchClientsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchClientsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchClientsRequest) GetNameOnly() bool {
	if m != nil {
		return m.NameOnly
	}
	return false
}

func (m *SearchClientsRequest) GetType() SearchClientsRequest_QueryType {
	if m != nil {
		return m.Type
	}
	return SearchClientsRequest_VALUE
}

type SearchClientsResponse struct {
	Items                []*ApiClient `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Names                []string     `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchClientsResponse) Reset()         { *m = SearchClientsResponse{} }
func (m *SearchClientsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchClientsResponse) ProtoMessage()    {}
func (*SearchClientsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{4}
}

func (m *SearchClientsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClientsResponse.Unmarshal(m, b)
}
func (m *SearchClientsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClientsResponse.Marshal(b, m, deterministic)
}
func (m *SearchClientsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClientsResponse.Merge(m, src)
}
func (m *SearchClientsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchClientsResponse.Size(m)
}
func (m *SearchClientsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClientsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClientsResponse proto.InternalMessageInfo

func (m *SearchClientsResponse) GetItems() []*ApiClient {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SearchClientsResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type GetClientRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Lightweight          bool     `protobuf:"varint,2,opt,name=lightweight,proto3" json:"lightweight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClientRequest) Reset()         { *m = GetClientRequest{} }
func (m *GetClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetClientRequest) ProtoMessage()    {}
func (*GetClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{5}
}

func (m *GetClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClientRequest.Unmarshal(m, b)
}
func (m *GetClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClientRequest.Marshal(b, m, deterministic)
}
func (m *GetClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClientRequest.Merge(m, src)
}
func (m *GetClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetClientRequest.Size(m)
}
func (m *GetClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClientRequest proto.InternalMessageInfo

func (m *GetClientRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetClientRequest) GetLightweight() bool {
	if m != nil {
		return m.Lightweight
	}
	return false
}

type LabelClientsRequest struct {
	ClientIds            []string `protobuf:"bytes,1,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Labels               []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Operation            string   `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelClientsRequest) Reset()         { *m = LabelClientsRequest{} }
func (m *LabelClientsRequest) String() string { return proto.CompactTextString(m) }
func (*LabelClientsRequest) ProtoMessage()    {}
func (*LabelClientsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{6}
}

func (m *LabelClientsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelClientsRequest.Unmarshal(m, b)
}
func (m *LabelClientsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelClientsRequest.Marshal(b, m, deterministic)
}
func (m *LabelClientsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelClientsRequest.Merge(m, src)
}
func (m *LabelClientsRequest) XXX_Size() int {
	return xxx_messageInfo_LabelClientsRequest.Size(m)
}
func (m *LabelClientsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelClientsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LabelClientsRequest proto.InternalMessageInfo

func (m *LabelClientsRequest) GetClientIds() []string {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *LabelClientsRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LabelClientsRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

type ClientLabels struct {
	// When was the labeling record last updated.
	Timestamp            uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Label                []string `protobuf:"bytes,2,rep,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientLabels) Reset()         { *m = ClientLabels{} }
func (m *ClientLabels) String() string { return proto.CompactTextString(m) }
func (*ClientLabels) ProtoMessage()    {}
func (*ClientLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{7}
}

func (m *ClientLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientLabels.Unmarshal(m, b)
}
func (m *ClientLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientLabels.Marshal(b, m, deterministic)
}
func (m *ClientLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientLabels.Merge(m, src)
}
func (m *ClientLabels) XXX_Size() int {
	return xxx_messageInfo_ClientLabels.Size(m)
}
func (m *ClientLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientLabels.DiscardUnknown(m)
}

var xxx_messageInfo_ClientLabels proto.InternalMessageInfo

func (m *ClientLabels) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ClientLabels) GetLabel() []string {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ApiClient_IPAddressClass", ApiClient_IPAddressClass_name, ApiClient_IPAddressClass_value)
	proto.RegisterEnum("proto.SearchClientsRequest_QueryType", SearchClientsRequest_QueryType_name, SearchClientsRequest_QueryType_value)
	proto.RegisterType((*AgentInformation)(nil), "proto.AgentInformation")
	proto.RegisterType((*Uname)(nil), "proto.Uname")
	proto.RegisterType((*ApiClient)(nil), "proto.ApiClient")
	proto.RegisterType((*SearchClientsRequest)(nil), "proto.SearchClientsRequest")
	proto.RegisterType((*SearchClientsResponse)(nil), "proto.SearchClientsResponse")
	proto.RegisterType((*GetClientRequest)(nil), "proto.GetClientRequest")
	proto.RegisterType((*LabelClientsRequest)(nil), "proto.LabelClientsRequest")
	proto.RegisterType((*ClientLabels)(nil), "proto.ClientLabels")
}

func init() {
	proto.RegisterFile("clients.proto", fileDescriptor_6c7b36ecb5ad4a28)
}

var fileDescriptor_6c7b36ecb5ad4a28 = []byte{
	// 1472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x96, 0xdd, 0x72, 0x5b, 0xb7,
	0x11, 0xc7, 0x43, 0x8b, 0x12, 0x45, 0x50, 0x76, 0x19, 0xc4, 0x6e, 0xce, 0x24, 0x76, 0x8d, 0x72,
	0xea, 0x86, 0x6a, 0xe4, 0x23, 0x91, 0x52, 0x65, 0x3b, 0xed, 0x4c, 0x86, 0x94, 0x94, 0x0c, 0x63,
	0x45, 0x52, 0x8f, 0xe4, 0x8f, 0xe6, 0xa2, 0x1c, 0xf0, 0x9c, 0x25, 0x89, 0xfa, 0x10, 0xa0, 0x01,
	0xd0, 0x34, 0x67, 0x72, 0xd1, 0x97, 0xe8, 0x4c, 0x6f, 0x3a, 0xbd, 0xeb, 0x93, 0xe4, 0x0d, 0xfa,
	0x06, 0xed, 0x6b, 0xb4, 0x33, 0x1d, 0x2c, 0x40, 0x52, 0x54, 0x9b, 0x1b, 0xf1, 0x00, 0xbb, 0xfb,
	0xdb, 0x3f, 0x80, 0xc5, 0x42, 0xe4, 0x76, 0x9a, 0x0b, 0x90, 0xd6, 0xc4, 0x63, 0xad, 0xac, 0xa2,
	0xeb, 0xf8, 0xf3, 0xc9, 0x5d, 0xfc, 0xd9, 0x35, 0x30, 0xe2, 0xd2, 0x8a, 0xd4, 0x1b, 0x6b, 0x5d,
	0x52, 0x6d, 0x0d, 0x40, 0xda, 0x8e, 0xec, 0x2b, 0x3d, 0xe2, 0x56, 0x28, 0x49, 0x23, 0x52, 0x7a,
	0x07, 0xda, 0x08, 0x25, 0xa3, 0x02, 0x2b, 0xd4, 0xcb, 0xc9, 0x7c, 0x48, 0x29, 0x29, 0x4a, 0x3e,
	0x82, 0xe8, 0x16, 0x4e, 0xe3, 0x37, 0x7d, 0x40, 0x48, 0x6f, 0x22, 0xf2, 0xac, 0x6b, 0xc5, 0x08,
	0xa2, 0x35, 0xb4, 0x94, 0x71, 0xe6, 0x4a, 0x8c, 0xa0, 0xf6, 0x9f, 0x0d, 0xb2, 0xfe, 0x02, 0x1d,
	0x9f, 0x93, 0x0d, 0x33, 0x33, 0x16, 0x46, 0x9e, 0xda, 0xde, 0xff, 0xe7, 0xbf, 0xff, 0xf5, 0x43,
	0xe1, 0x31, 0xfd, 0xfc, 0x6a, 0x08, 0xcc, 0x5b, 0xd8, 0x38, 0xe7, 0xd6, 0x09, 0x61, 0xf5, 0x57,
	0x42, 0x66, 0x6a, 0x6a, 0xbe, 0x3f, 0xe6, 0x7a, 0x2a, 0xe4, 0xf7, 0xa7, 0x42, 0x4e, 0xde, 0x6f,
	0xc7, 0x49, 0x40, 0xd0, 0xa7, 0xa4, 0x28, 0x55, 0x16, 0x94, 0xb4, 0x7f, 0x81, 0xa8, 0x9f, 0xd1,
	0xfb, 0x0e, 0x35, 0x54, 0xc6, 0xba, 0x84, 0x4c, 0xf5, 0x99, 0x1d, 0x0a, 0x13, 0xd8, 0x71, 0x82,
	0x11, 0xf4, 0x82, 0x94, 0x34, 0xe4, 0xc0, 0x4d, 0x10, 0xdb, 0x3e, 0xc4, 0xe0, 0x3d, 0x1a, 0xbb,
	0xe0, 0xf3, 0x4b, 0x16, 0xac, 0x4c, 0x64, 0x20, 0xad, 0xe8, 0x0b, 0xd0, 0x0c, 0xe2, 0x41, 0xcc,
	0x9e, 0xec, 0xb0, 0xf3, 0xcb, 0xd7, 0x3b, 0x2c, 0x83, 0x9e, 0xe0, 0x32, 0x4e, 0xe6, 0x18, 0x7a,
	0xb5, 0xdc, 0xaf, 0x22, 0x12, 0xbf, 0x40, 0xe2, 0x01, 0x6d, 0x06, 0x62, 0xb0, 0xb2, 0xce, 0xb1,
	0x27, 0x1d, 0xc6, 0x8d, 0xf8, 0xc9, 0xe1, 0x5e, 0xe3, 0xf2, 0xa2, 0xb1, 0xc3, 0x1a, 0x7b, 0xf1,
	0xb3, 0xb8, 0xb9, 0xc3, 0x1a, 0x07, 0xf1, 0xde, 0x41, 0xbc, 0xdc, 0xeb, 0x6f, 0x49, 0x69, 0xc4,
	0xd3, 0xa1, 0x90, 0x10, 0xad, 0xff, 0xe8, 0x7e, 0x71, 0x9d, 0x0e, 0x85, 0x85, 0xd4, 0x4e, 0x34,
	0x78, 0x76, 0xeb, 0xdb, 0xe3, 0xc3, 0x83, 0x1d, 0xf6, 0xfe, 0xe9, 0x61, 0xf7, 0xd0, 0xe1, 0x02,
	0x83, 0x7e, 0x47, 0x36, 0xde, 0x80, 0x96, 0x90, 0x47, 0x1b, 0x48, 0x6b, 0x23, 0xed, 0xb7, 0xf4,
	0x0b, 0x47, 0xf3, 0x96, 0x85, 0x4e, 0x63, 0xb5, 0x90, 0x83, 0x55, 0xad, 0x3b, 0xac, 0xb1, 0x1f,
	0x37, 0xe2, 0xbd, 0x1d, 0xb6, 0x1f, 0x37, 0x7e, 0xfd, 0x58, 0xa7, 0xcd, 0x38, 0x09, 0x44, 0x7a,
	0x42, 0x8a, 0xfd, 0xb7, 0x99, 0x8c, 0x4a, 0x48, 0x6e, 0x20, 0xf9, 0x73, 0xba, 0xbd, 0xd4, 0xf9,
	0x99, 0x61, 0xfd, 0x49, 0x9e, 0xcf, 0xd8, 0xdb, 0x09, 0xcf, 0xdd, 0x96, 0x66, 0x2c, 0x53, 0x23,
	0x2e, 0x24, 0x73, 0x07, 0x15, 0x27, 0x18, 0x4e, 0x13, 0xb2, 0x25, 0xa4, 0xb1, 0x3c, 0xcf, 0xbb,
	0x19, 0xb7, 0x10, 0x6d, 0xb2, 0x42, 0xbd, 0xd8, 0xde, 0x45, 0xdc, 0x36, 0xa9, 0x24, 0xc7, 0x5f,
	0x1d, 0x73, 0x0b, 0xae, 0xcc, 0xe8, 0x27, 0xaf, 0x86, 0x20, 0xe7, 0x9b, 0x30, 0xe5, 0x86, 0x85,
	0x40, 0xc8, 0xe2, 0xa4, 0x12, 0xbe, 0x9d, 0x33, 0x7d, 0x4a, 0x36, 0x73, 0xd1, 0x4b, 0xbb, 0xef,
	0x40, 0x47, 0x65, 0x94, 0xf7, 0x00, 0x79, 0x1f, 0xd3, 0x7b, 0x4e, 0xde, 0x11, 0xcb, 0x45, 0x4f,
	0x73, 0x3d, 0x9b, 0xaf, 0x3d, 0x29, 0x39, 0xf7, 0x97, 0xa0, 0xe9, 0x0f, 0x05, 0xb2, 0x75, 0x7d,
	0x7b, 0x23, 0x82, 0xe1, 0x7f, 0x2b, 0x60, 0xfc, 0x5f, 0x0a, 0xf4, 0xcf, 0x05, 0x47, 0x58, 0x39,
	0x81, 0x79, 0xc5, 0xf5, 0x84, 0xe4, 0x7a, 0x16, 0xb3, 0xfa, 0x99, 0xb2, 0xe0, 0xa7, 0x52, 0x2e,
	0x59, 0x0f, 0x58, 0x26, 0xfa, 0x7d, 0xd0, 0x20, 0x2d, 0xeb, 0x6b, 0x35, 0x62, 0x76, 0x08, 0x2c,
	0x9c, 0xd0, 0x2a, 0x49, 0x48, 0xb4, 0xa5, 0xae, 0x10, 0x55, 0x9f, 0x71, 0xb6, 0xdf, 0x64, 0x3d,
	0x61, 0x03, 0x99, 0xe9, 0x89, 0x94, 0xee, 0x88, 0x94, 0x64, 0x9c, 0x1d, 0x1e, 0xa0, 0xc9, 0xef,
	0xc6, 0x76, 0xb2, 0xa2, 0xba, 0xf6, 0xd7, 0x12, 0x29, 0xb7, 0xc6, 0xe2, 0x08, 0x5b, 0x02, 0xfd,
	0x92, 0x94, 0x7d, 0x73, 0xe8, 0x8a, 0x2c, 0x5c, 0xc3, 0x1a, 0xae, 0xe7, 0x3e, 0xa9, 0x2c, 0xbc,
	0x3a, 0x19, 0xbd, 0xed, 0x96, 0xe6, 0x3d, 0x99, 0xc8, 0x92, 0xcd, 0x74, 0x6e, 0x38, 0x26, 0x1f,
	0xf2, 0x01, 0xc6, 0x2f, 0x1b, 0x06, 0x5e, 0xc2, 0x4a, 0xf3, 0x63, 0xdf, 0x52, 0xe2, 0x9b, 0xfd,
	0x24, 0xa9, 0xf2, 0x9b, 0x1d, 0xe6, 0x11, 0x29, 0x29, 0x83, 0x08, 0xbc, 0x83, 0x95, 0xe6, 0x56,
	0x88, 0xc5, 0x4e, 0x91, 0x6c, 0x28, 0xe3, 0xbc, 0xa9, 0x25, 0xb7, 0xfb, 0x42, 0x1b, 0xdb, 0x35,
	0x00, 0xb2, 0xcb, 0x2d, 0x96, 0x6e, 0xb1, 0x7d, 0x81, 0x8a, 0xbf, 0x59, 0xad, 0x88, 0xdf, 0x60,
	0x45, 0xd8, 0xa5, 0x6c, 0x57, 0x15, 0x18, 0xcd, 0x5c, 0x34, 0xab, 0x8b, 0x18, 0x62, 0x36, 0x75,
	0x4e, 0xc2, 0x1b, 0x41, 0x6a, 0xe5, 0x2a, 0x66, 0x3b, 0x4e, 0x2a, 0xe8, 0x78, 0x09, 0x20, 0x5b,
	0x96, 0xbe, 0x26, 0x5b, 0x39, 0xbf, 0x96, 0xb4, 0x84, 0x49, 0x43, 0x97, 0x58, 0x4d, 0xfa, 0xf3,
	0x53, 0x6e, 0x2c, 0x73, 0x9f, 0x9e, 0x1c, 0x52, 0xa7, 0x43, 0x48, 0xdf, 0x40, 0xc6, 0x84, 0x8c,
	0x13, 0xe2, 0x58, 0x81, 0xfc, 0x0d, 0xb9, 0x83, 0xe4, 0x9e, 0x52, 0x16, 0x32, 0xc7, 0xf6, 0x25,
	0x1e, 0xda, 0xd7, 0x2a, 0xfb, 0x27, 0xc8, 0x76, 0xae, 0x98, 0x20, 0x4e, 0x50, 0x55, 0x1b, 0x43,
	0x5b, 0x96, 0xfe, 0x81, 0x20, 0xb9, 0x9b, 0xe6, 0x2a, 0x7d, 0x83, 0xa5, 0x5d, 0x6c, 0x7f, 0x89,
	0x9c, 0x67, 0xab, 0x9c, 0x5f, 0x1d, 0x05, 0x51, 0xce, 0xd1, 0xb0, 0x77, 0x3c, 0x9f, 0x00, 0xcb,
	0x26, 0x78, 0xc3, 0x73, 0x6e, 0xc1, 0x04, 0xbd, 0x4e, 0x6c, 0xd9, 0x21, 0x8f, 0x9c, 0x23, 0xed,
	0x90, 0xdb, 0x9e, 0xaf, 0xb9, 0x19, 0x3a, 0xa9, 0x04, 0x53, 0x3c, 0xc2, 0x14, 0x0f, 0x57, 0x53,
	0x54, 0x51, 0x2a, 0x7a, 0x06, 0xad, 0x15, 0x04, 0xb9, 0x89, 0x96, 0xa5, 0x2d, 0x52, 0x42, 0x94,
	0x18, 0x47, 0x55, 0x2c, 0xb9, 0x3a, 0x42, 0x6a, 0x94, 0xb9, 0x2a, 0x73, 0x26, 0x7f, 0x44, 0x1a,
	0x46, 0xee, 0xba, 0xb4, 0x2e, 0x3a, 0x8c, 0x67, 0x99, 0x06, 0x63, 0x92, 0x0d, 0x67, 0xed, 0x8c,
	0xe9, 0x13, 0x12, 0x79, 0x84, 0xb4, 0xa0, 0xb5, 0x1a, 0x70, 0x0b, 0xdd, 0x7e, 0xae, 0xa6, 0xae,
	0x8c, 0x3f, 0xc2, 0x27, 0xe7, 0x1e, 0x7a, 0x2e, 0xcd, 0x5f, 0xe5, 0x6a, 0xda, 0xc9, 0xe8, 0x51,
	0x58, 0x86, 0x18, 0x77, 0xd3, 0x9c, 0x1b, 0x13, 0x7d, 0xc8, 0x0a, 0xf5, 0x3b, 0xcd, 0x87, 0xf3,
	0x5a, 0x9d, 0xd7, 0x7c, 0xdc, 0xb9, 0x68, 0xf9, 0x9c, 0x47, 0xce, 0xcd, 0x2f, 0xa0, 0x33, 0xc6,
	0x01, 0xfd, 0x29, 0xd9, 0xc8, 0x79, 0x0f, 0x72, 0x13, 0x51, 0xb6, 0x56, 0x2f, 0x27, 0x61, 0x54,
	0x6b, 0x93, 0x3b, 0xab, 0x61, 0xb4, 0x42, 0x4a, 0x2f, 0xce, 0x9e, 0x9f, 0x9d, 0xbf, 0x3a, 0xab,
	0x7e, 0x40, 0xb7, 0xc8, 0x66, 0xe7, 0xec, 0xea, 0x24, 0x39, 0x6b, 0x9d, 0x56, 0x0b, 0x6e, 0x74,
	0xf2, 0x3a, 0x8c, 0x6e, 0xd1, 0x12, 0x59, 0x7b, 0x79, 0x71, 0x56, 0x5d, 0xab, 0xfd, 0xa3, 0x40,
	0xee, 0x5e, 0x82, 0xbb, 0xb2, 0x5e, 0x88, 0x49, 0xe0, 0xed, 0x04, 0x8c, 0x75, 0x49, 0x55, 0xbf,
	0x6f, 0xc0, 0xe2, 0x3d, 0x2d, 0x26, 0x61, 0x44, 0xef, 0x92, 0xf5, 0x5c, 0x8c, 0x84, 0xc5, 0x5b,
	0x57, 0x4c, 0xfc, 0xc0, 0xcd, 0xbe, 0x9d, 0x80, 0x9e, 0x85, 0x07, 0xd8, 0x0f, 0xe8, 0xa7, 0xa4,
	0xec, 0x2e, 0x54, 0x57, 0xc9, 0x7c, 0x86, 0x6f, 0xd3, 0x66, 0xb2, 0xe9, 0x26, 0xce, 0x65, 0x3e,
	0xa3, 0xcf, 0x48, 0xd1, 0xce, 0xc6, 0xfe, 0x75, 0xb9, 0xd3, 0x7c, 0x14, 0x76, 0xe4, 0xff, 0x69,
	0x89, 0x7f, 0xe7, 0x68, 0x57, 0xb3, 0x31, 0x24, 0x18, 0x52, 0x7b, 0x48, 0xca, 0x8b, 0x29, 0x5a,
	0x26, 0xeb, 0x2f, 0x5b, 0xa7, 0x2f, 0x4e, 0xaa, 0x1f, 0xb8, 0x55, 0x3d, 0x3f, 0xf9, 0x7d, 0xb5,
	0x50, 0xfb, 0x7b, 0x81, 0xdc, 0xbb, 0x41, 0x32, 0x63, 0x25, 0x0d, 0xd0, 0x5f, 0x92, 0x75, 0x61,
	0x61, 0x64, 0xa2, 0x02, 0x5b, 0xab, 0x57, 0x9a, 0xd5, 0x9b, 0x07, 0x91, 0x78, 0x33, 0x05, 0xb2,
	0xee, 0x94, 0x9a, 0xe8, 0x96, 0xdb, 0xf2, 0xf6, 0x39, 0x96, 0x4c, 0x87, 0x7e, 0xdd, 0xe9, 0xb3,
	0xc5, 0x92, 0x98, 0x7b, 0xdc, 0xc7, 0x90, 0xfa, 0x17, 0x25, 0xb4, 0x4c, 0xed, 0x35, 0xb3, 0x29,
	0x30, 0xf4, 0xd1, 0x60, 0x27, 0xda, 0x9b, 0x10, 0xc8, 0x86, 0xa0, 0x21, 0x4e, 0x3c, 0xbd, 0xf6,
	0xa7, 0x02, 0xa9, 0x7e, 0x0d, 0x36, 0xe4, 0x0e, 0x5b, 0xff, 0xe9, 0xff, 0x74, 0xc9, 0x6b, 0x1d,
	0xf0, 0x9c, 0x54, 0x72, 0x31, 0x18, 0xda, 0x29, 0xb8, 0xbf, 0x78, 0x0a, 0x9b, 0xed, 0xc7, 0x28,
	0xef, 0x33, 0xfa, 0xa8, 0xd3, 0x67, 0x06, 0xac, 0xcf, 0x9c, 0xba, 0xf6, 0x92, 0x5a, 0x66, 0xd4,
	0xc8, 0x35, 0xf4, 0x45, 0xfb, 0x73, 0xd7, 0x63, 0x49, 0xa8, 0xfd, 0x91, 0x7c, 0x74, 0xea, 0xea,
	0xe9, 0xc6, 0xf9, 0x3f, 0x20, 0x64, 0x21, 0xc2, 0xef, 0x56, 0x39, 0x29, 0xcf, 0x55, 0x5c, 0xaf,
	0xc9, 0x5b, 0xd7, 0x6b, 0x92, 0xde, 0x27, 0x65, 0x35, 0x06, 0xed, 0x1b, 0x73, 0xf8, 0x6f, 0x6c,
	0x31, 0x51, 0x6b, 0x93, 0x2d, 0x9f, 0xe6, 0x74, 0xe1, 0xed, 0x2e, 0xac, 0xb1, 0x7c, 0x34, 0x0e,
	0x75, 0xb6, 0x9c, 0xc0, 0x52, 0x73, 0x7e, 0x21, 0x85, 0x1f, 0xb4, 0x1b, 0xdf, 0xed, 0x4e, 0xa7,
	0xd3, 0xf8, 0x1d, 0xe4, 0x2a, 0x15, 0x19, 0xbc, 0x8f, 0x53, 0x35, 0xda, 0x1d, 0xa8, 0x9c, 0xcb,
	0xc1, 0xae, 0x9f, 0xd4, 0x7c, 0x6c, 0x95, 0xde, 0xe5, 0x63, 0xb1, 0x8b, 0xa7, 0xdb, 0xdb, 0xc0,
	0x9f, 0xfd, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x6a, 0xf6, 0x18, 0x9a, 0x0a, 0x00, 0x00,
}

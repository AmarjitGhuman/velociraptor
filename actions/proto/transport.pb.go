// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport.proto

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

type Range struct {
	FileOffset           int64    `protobuf:"varint,1,opt,name=file_offset,json=fileOffset,proto3" json:"file_offset,omitempty"`
	OriginalOffset       int64    `protobuf:"varint,2,opt,name=original_offset,json=originalOffset,proto3" json:"original_offset,omitempty"`
	FileLength           int64    `protobuf:"varint,3,opt,name=file_length,json=fileLength,proto3" json:"file_length,omitempty"`
	Length               int64    `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97e32c760ec1b28, []int{0}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetFileOffset() int64 {
	if m != nil {
		return m.FileOffset
	}
	return 0
}

func (m *Range) GetOriginalOffset() int64 {
	if m != nil {
		return m.OriginalOffset
	}
	return 0
}

func (m *Range) GetFileLength() int64 {
	if m != nil {
		return m.FileLength
	}
	return 0
}

func (m *Range) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Index struct {
	Ranges               []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97e32c760ec1b28, []int{1}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetRanges() []*Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

// A message to encode a filesystem path (maybe for raw access)
// Next field: 15
type PathSpec struct {
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Accessor             string   `protobuf:"bytes,3,opt,name=accessor,proto3" json:"accessor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathSpec) Reset()         { *m = PathSpec{} }
func (m *PathSpec) String() string { return proto.CompactTextString(m) }
func (*PathSpec) ProtoMessage()    {}
func (*PathSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97e32c760ec1b28, []int{2}
}

func (m *PathSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathSpec.Unmarshal(m, b)
}
func (m *PathSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathSpec.Marshal(b, m, deterministic)
}
func (m *PathSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathSpec.Merge(m, src)
}
func (m *PathSpec) XXX_Size() int {
	return xxx_messageInfo_PathSpec.Size(m)
}
func (m *PathSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PathSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PathSpec proto.InternalMessageInfo

func (m *PathSpec) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathSpec) GetAccessor() string {
	if m != nil {
		return m.Accessor
	}
	return ""
}

// The Velociraptor client sends back the buffer and the filename and
// the server saves the entire file directly in the file storage
// filesystem. This allows easy recovery as well as data expiration
// policies (since the filestore is just a directory on disk with
// regular files and timestamps).
type FileBuffer struct {
	Pathspec *PathSpec `protobuf:"bytes,1,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	Offset   uint64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Expected size of file.
	Size       uint64 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	StoredSize uint64 `protobuf:"varint,8,opt,name=stored_size,json=storedSize,proto3" json:"stored_size,omitempty"`
	IsSparse   bool   `protobuf:"varint,9,opt,name=is_sparse,json=isSparse,proto3" json:"is_sparse,omitempty"`
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	FlowId     string `protobuf:"bytes,4,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Eof        bool   `protobuf:"varint,5,opt,name=eof,proto3" json:"eof,omitempty"`
	// Set when the file is sparse.
	Index                *Index   `protobuf:"bytes,6,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileBuffer) Reset()         { *m = FileBuffer{} }
func (m *FileBuffer) String() string { return proto.CompactTextString(m) }
func (*FileBuffer) ProtoMessage()    {}
func (*FileBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97e32c760ec1b28, []int{3}
}

func (m *FileBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBuffer.Unmarshal(m, b)
}
func (m *FileBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBuffer.Marshal(b, m, deterministic)
}
func (m *FileBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBuffer.Merge(m, src)
}
func (m *FileBuffer) XXX_Size() int {
	return xxx_messageInfo_FileBuffer.Size(m)
}
func (m *FileBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_FileBuffer proto.InternalMessageInfo

func (m *FileBuffer) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

func (m *FileBuffer) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FileBuffer) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileBuffer) GetStoredSize() uint64 {
	if m != nil {
		return m.StoredSize
	}
	return 0
}

func (m *FileBuffer) GetIsSparse() bool {
	if m != nil {
		return m.IsSparse
	}
	return false
}

func (m *FileBuffer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FileBuffer) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *FileBuffer) GetEof() bool {
	if m != nil {
		return m.Eof
	}
	return false
}

func (m *FileBuffer) GetIndex() *Index {
	if m != nil {
		return m.Index
	}
	return nil
}

type ForemanCheckin struct {
	LastHuntTimestamp     uint64   `protobuf:"varint,1,opt,name=last_hunt_timestamp,json=lastHuntTimestamp,proto3" json:"last_hunt_timestamp,omitempty"`
	LastEventTableVersion uint64   `protobuf:"varint,2,opt,name=last_event_table_version,json=lastEventTableVersion,proto3" json:"last_event_table_version,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ForemanCheckin) Reset()         { *m = ForemanCheckin{} }
func (m *ForemanCheckin) String() string { return proto.CompactTextString(m) }
func (*ForemanCheckin) ProtoMessage()    {}
func (*ForemanCheckin) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97e32c760ec1b28, []int{4}
}

func (m *ForemanCheckin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForemanCheckin.Unmarshal(m, b)
}
func (m *ForemanCheckin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForemanCheckin.Marshal(b, m, deterministic)
}
func (m *ForemanCheckin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForemanCheckin.Merge(m, src)
}
func (m *ForemanCheckin) XXX_Size() int {
	return xxx_messageInfo_ForemanCheckin.Size(m)
}
func (m *ForemanCheckin) XXX_DiscardUnknown() {
	xxx_messageInfo_ForemanCheckin.DiscardUnknown(m)
}

var xxx_messageInfo_ForemanCheckin proto.InternalMessageInfo

func (m *ForemanCheckin) GetLastHuntTimestamp() uint64 {
	if m != nil {
		return m.LastHuntTimestamp
	}
	return 0
}

func (m *ForemanCheckin) GetLastEventTableVersion() uint64 {
	if m != nil {
		return m.LastEventTableVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*Range)(nil), "proto.Range")
	proto.RegisterType((*Index)(nil), "proto.Index")
	proto.RegisterType((*PathSpec)(nil), "proto.PathSpec")
	proto.RegisterType((*FileBuffer)(nil), "proto.FileBuffer")
	proto.RegisterType((*ForemanCheckin)(nil), "proto.ForemanCheckin")
}

func init() {
	proto.RegisterFile("transport.proto", fileDescriptor_a97e32c760ec1b28)
}

var fileDescriptor_a97e32c760ec1b28 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x55, 0x3a, 0x8f, 0xce, 0xb8, 0x55, 0x0b, 0xe6, 0x15, 0xc1, 0x02, 0x33, 0x20, 0xb5, 0x15,
	0x22, 0x23, 0xb5, 0x42, 0xac, 0x19, 0x44, 0x45, 0x05, 0x12, 0xc8, 0x33, 0x62, 0xc1, 0x26, 0xf2,
	0x24, 0x37, 0x13, 0x8b, 0xc4, 0x8e, 0x6c, 0xcf, 0xa4, 0xc3, 0x8e, 0x2d, 0x7f, 0xc3, 0x3f, 0xf0,
	0x17, 0xec, 0xe0, 0x37, 0x58, 0x20, 0xdf, 0x24, 0x6d, 0x57, 0xb6, 0xcf, 0x39, 0x3e, 0x3e, 0xbe,
	0xf7, 0x92, 0x43, 0x67, 0x84, 0xb2, 0x95, 0x36, 0x2e, 0xaa, 0x8c, 0x76, 0x9a, 0x0e, 0x70, 0x79,
	0x78, 0x17, 0x97, 0xa9, 0x85, 0x52, 0x28, 0x27, 0x93, 0x86, 0x9c, 0xfc, 0x08, 0xc8, 0x80, 0x0b,
	0xb5, 0x02, 0xfa, 0x98, 0xec, 0x65, 0xb2, 0x80, 0x58, 0x67, 0x99, 0x05, 0x17, 0x06, 0x2c, 0x38,
	0xee, 0x71, 0xe2, 0xa1, 0x8f, 0x88, 0xd0, 0x23, 0x72, 0xa8, 0x8d, 0x5c, 0x49, 0x25, 0x8a, 0x4e,
	0xb4, 0x83, 0xa2, 0x83, 0x0e, 0x6e, 0x85, 0x9d, 0x53, 0x01, 0x6a, 0xe5, 0xf2, 0xb0, 0x77, 0xed,
	0xf4, 0x01, 0x11, 0x7a, 0x9f, 0x0c, 0x5b, 0xae, 0x8f, 0x5c, 0x7b, 0x9a, 0xbc, 0x20, 0x83, 0x0b,
	0x95, 0xc2, 0x25, 0x7d, 0x46, 0x86, 0xc6, 0x87, 0xb2, 0x61, 0xc0, 0x7a, 0xc7, 0x7b, 0xa7, 0xfb,
	0x4d, 0xda, 0x08, 0x93, 0xf2, 0x96, 0x9b, 0xfc, 0x0e, 0xc8, 0xe8, 0x93, 0x70, 0xf9, 0xbc, 0x82,
	0x84, 0x7e, 0x0f, 0x48, 0xbf, 0x12, 0x2e, 0xc7, 0x4c, 0xe3, 0x59, 0xf9, 0xe7, 0xdf, 0xdf, 0x5f,
	0xc1, 0x8a, 0xc2, 0x22, 0x07, 0xe6, 0x71, 0x56, 0x09, 0x6b, 0x21, 0x65, 0x4e, 0x33, 0x97, 0x03,
	0x3b, 0x97, 0x05, 0xd8, 0xad, 0x75, 0x50, 0x32, 0x91, 0x24, 0x60, 0xad, 0x36, 0x11, 0x5b, 0xe4,
	0xd2, 0xb2, 0x8d, 0x28, 0xd6, 0xc0, 0xa4, 0x65, 0x52, 0x39, 0x30, 0x95, 0x01, 0x07, 0x29, 0x5b,
	0x6e, 0xf1, 0x56, 0x27, 0x65, 0x52, 0x31, 0xe9, 0x2c, 0xd3, 0xb5, 0x62, 0xb5, 0xd8, 0x46, 0x1c,
	0x9f, 0xa6, 0xef, 0xc9, 0xa8, 0x53, 0xe0, 0xaf, 0xc7, 0xb3, 0x29, 0xc6, 0x38, 0xa1, 0x47, 0x8b,
	0x9b, 0xb7, 0xd7, 0x6d, 0x10, 0x03, 0xce, 0x48, 0xd8, 0x00, 0x7a, 0xfb, 0xfa, 0x44, 0xfc, 0xca,
	0x60, 0xf2, 0x73, 0x87, 0x10, 0x9f, 0x72, 0xb6, 0xce, 0x32, 0x30, 0xf4, 0x39, 0x19, 0xf9, 0x37,
	0x6c, 0x05, 0x09, 0xf6, 0x66, 0xef, 0xf4, 0xb0, 0x2d, 0x4a, 0x57, 0x02, 0x7e, 0x25, 0xa0, 0xaf,
	0xc9, 0xf0, 0x46, 0x87, 0xfa, 0xb3, 0x13, 0x8c, 0xf1, 0x94, 0x3e, 0x69, 0x3a, 0xc4, 0x74, 0x86,
	0x4f, 0x2e, 0xd1, 0xda, 0x7f, 0xe6, 0x3a, 0x40, 0x7b, 0x91, 0x52, 0xd2, 0xb7, 0xf2, 0x1b, 0x84,
	0xbb, 0xde, 0x80, 0xe3, 0xde, 0x37, 0xd6, 0x3a, 0x6d, 0x20, 0x8d, 0x91, 0x1a, 0x21, 0x45, 0x1a,
	0x68, 0xee, 0x05, 0x8f, 0xc8, 0x58, 0xda, 0xd8, 0x56, 0xc2, 0x58, 0x08, 0xc7, 0x2c, 0x38, 0x1e,
	0xf1, 0x91, 0xb4, 0x73, 0x3c, 0x7b, 0xc7, 0x54, 0x38, 0x81, 0x95, 0xd9, 0xe7, 0xb8, 0xa7, 0x0f,
	0xc8, 0x6e, 0x56, 0xe8, 0x3a, 0x96, 0x29, 0x8e, 0xc2, 0x98, 0x0f, 0xfd, 0xf1, 0x22, 0xa5, 0xb7,
	0x48, 0x0f, 0x74, 0x16, 0x0e, 0xd0, 0xc3, 0x6f, 0xe9, 0x84, 0x0c, 0xa4, 0x1f, 0x8e, 0x70, 0x88,
	0xbf, 0xef, 0x46, 0x02, 0x07, 0x86, 0x37, 0xd4, 0x64, 0x4b, 0x0e, 0xce, 0xb5, 0xf1, 0x13, 0xfe,
	0x26, 0x87, 0xe4, 0xab, 0x54, 0x34, 0x22, 0x77, 0x0a, 0x61, 0x5d, 0x9c, 0xaf, 0x95, 0x8b, 0x9d,
	0x2c, 0xc1, 0x3a, 0x51, 0x56, 0x58, 0xc1, 0x3e, 0xbf, 0xed, 0xa9, 0x77, 0x6b, 0xe5, 0x16, 0x1d,
	0x41, 0x5f, 0x91, 0x10, 0xf5, 0xb0, 0x01, 0x7f, 0x41, 0x2c, 0x0b, 0x88, 0x37, 0x60, 0xac, 0xd4,
	0xaa, 0xa9, 0x25, 0xbf, 0xe7, 0xf9, 0xb7, 0x9e, 0x5e, 0x78, 0xf6, 0x73, 0x43, 0xce, 0x5e, 0x7e,
	0x39, 0xab, 0xeb, 0x3a, 0xda, 0x40, 0xa1, 0x13, 0x99, 0xc2, 0x65, 0x94, 0xe8, 0x72, 0xba, 0xd2,
	0x85, 0x50, 0xab, 0x69, 0x03, 0x1a, 0x51, 0x39, 0x6d, 0xa6, 0x22, 0x71, 0x52, 0x2b, 0x3b, 0xc5,
	0xf8, 0xcb, 0x21, 0x2e, 0x67, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0x97, 0xac, 0x1b, 0xb6,
	0x03, 0x00, 0x00,
}

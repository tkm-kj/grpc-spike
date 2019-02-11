// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playlist.proto

package com_exapmle_spike

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

type Playlist struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Order                int64                `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Comments             *Comments            `protobuf:"bytes,5,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Playlist) Reset()         { *m = Playlist{} }
func (m *Playlist) String() string { return proto.CompactTextString(m) }
func (*Playlist) ProtoMessage()    {}
func (*Playlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_073b59cb25aac063, []int{0}
}

func (m *Playlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Playlist.Unmarshal(m, b)
}
func (m *Playlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Playlist.Marshal(b, m, deterministic)
}
func (m *Playlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Playlist.Merge(m, src)
}
func (m *Playlist) XXX_Size() int {
	return xxx_messageInfo_Playlist.Size(m)
}
func (m *Playlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Playlist.DiscardUnknown(m)
}

var xxx_messageInfo_Playlist proto.InternalMessageInfo

func (m *Playlist) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Playlist) GetOrder() int64 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Playlist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Playlist) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Playlist) GetComments() *Comments {
	if m != nil {
		return m.Comments
	}
	return nil
}

type PlaylistRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlaylistRequest) Reset()         { *m = PlaylistRequest{} }
func (m *PlaylistRequest) String() string { return proto.CompactTextString(m) }
func (*PlaylistRequest) ProtoMessage()    {}
func (*PlaylistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_073b59cb25aac063, []int{1}
}

func (m *PlaylistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaylistRequest.Unmarshal(m, b)
}
func (m *PlaylistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaylistRequest.Marshal(b, m, deterministic)
}
func (m *PlaylistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaylistRequest.Merge(m, src)
}
func (m *PlaylistRequest) XXX_Size() int {
	return xxx_messageInfo_PlaylistRequest.Size(m)
}
func (m *PlaylistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaylistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlaylistRequest proto.InternalMessageInfo

func (m *PlaylistRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Playlist)(nil), "com.exapmle.spike.Playlist")
	proto.RegisterType((*PlaylistRequest)(nil), "com.exapmle.spike.PlaylistRequest")
}

func init() { proto.RegisterFile("playlist.proto", fileDescriptor_073b59cb25aac063) }

var fileDescriptor_073b59cb25aac063 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xc9, 0xb5, 0x95, 0xf6, 0x2b, 0x56, 0x0c, 0x0e, 0xe1, 0x3a, 0x78, 0xde, 0x74, 0x53,
	0x0a, 0x75, 0x10, 0x47, 0x71, 0x70, 0x93, 0x12, 0xdd, 0x25, 0xbd, 0xfb, 0x2c, 0xc1, 0x4b, 0x13,
	0x93, 0x54, 0xf4, 0x9f, 0xf9, 0xf3, 0x84, 0x5c, 0xae, 0x82, 0x95, 0x6e, 0x49, 0x78, 0xde, 0x37,
	0xcf, 0x0b, 0x33, 0xdb, 0xca, 0xaf, 0x56, 0xf9, 0xc0, 0xad, 0x33, 0xc1, 0xd0, 0xf3, 0xda, 0x68,
	0x8e, 0x9f, 0xd2, 0xea, 0x16, 0xb9, 0xb7, 0xea, 0x0d, 0xf3, 0xcb, 0x8d, 0x31, 0x9b, 0x16, 0x17,
	0x11, 0x58, 0xef, 0x5e, 0x17, 0x41, 0x69, 0xf4, 0x41, 0x6a, 0xdb, 0x65, 0xf2, 0xd3, 0xda, 0x68,
	0x8d, 0xdb, 0x54, 0x51, 0x7e, 0x13, 0x18, 0xaf, 0x52, 0x2b, 0x9d, 0x41, 0xa6, 0x1a, 0x46, 0x0a,
	0x52, 0x0d, 0x44, 0xa6, 0x1a, 0x7a, 0x01, 0x23, 0xe3, 0x1a, 0x74, 0x2c, 0x8b, 0x4f, 0xdd, 0x85,
	0x52, 0x18, 0x6e, 0xa5, 0x46, 0x36, 0x28, 0x48, 0x35, 0x11, 0xf1, 0x4c, 0x6f, 0x01, 0x6a, 0x87,
	0x32, 0x60, 0xf3, 0x22, 0x03, 0x1b, 0x16, 0xa4, 0x9a, 0x2e, 0x73, 0xde, 0xb9, 0xf0, 0xde, 0x85,
	0x3f, 0xf7, 0x2e, 0x62, 0x92, 0xe8, 0xbb, 0x40, 0x6f, 0x60, 0x9c, 0x94, 0x3c, 0x1b, 0xc5, 0xe0,
	0x9c, 0x1f, 0xec, 0xe2, 0xf7, 0x09, 0x11, 0x7b, 0xb8, 0xbc, 0x82, 0xb3, 0xde, 0x5c, 0xe0, 0xfb,
	0x0e, 0x0f, 0x07, 0x2c, 0xe5, 0x2f, 0xf2, 0x84, 0xee, 0x43, 0xd5, 0x48, 0x1f, 0x61, 0xfa, 0x80,
	0x61, 0x3f, 0xb9, 0xfc, 0xe7, 0xaf, 0x3f, 0xad, 0xf9, 0xfc, 0x08, 0xb3, 0x22, 0xeb, 0x93, 0xb8,
	0xef, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x79, 0x24, 0x4d, 0xa7, 0x9e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlaylistServiceClient is the client API for PlaylistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaylistServiceClient interface {
	GetPlaylist(ctx context.Context, in *PlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
}

type playlistServiceClient struct {
	cc *grpc.ClientConn
}

func NewPlaylistServiceClient(cc *grpc.ClientConn) PlaylistServiceClient {
	return &playlistServiceClient{cc}
}

func (c *playlistServiceClient) GetPlaylist(ctx context.Context, in *PlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/com.exapmle.spike.PlaylistService/GetPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServiceServer is the server API for PlaylistService service.
type PlaylistServiceServer interface {
	GetPlaylist(context.Context, *PlaylistRequest) (*Playlist, error)
}

func RegisterPlaylistServiceServer(s *grpc.Server, srv PlaylistServiceServer) {
	s.RegisterService(&_PlaylistService_serviceDesc, srv)
}

func _PlaylistService_GetPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.exapmle.spike.PlaylistService/GetPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylist(ctx, req.(*PlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlaylistService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.exapmle.spike.PlaylistService",
	HandlerType: (*PlaylistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlaylist",
			Handler:    _PlaylistService_GetPlaylist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playlist.proto",
}

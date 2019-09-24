// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todolist.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetTodoResp struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResp) Reset()         { *m = GetTodoResp{} }
func (m *GetTodoResp) String() string { return proto.CompactTextString(m) }
func (*GetTodoResp) ProtoMessage()    {}
func (*GetTodoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_7d6f602f88ab0b11, []int{0}
}
func (m *GetTodoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResp.Unmarshal(m, b)
}
func (m *GetTodoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResp.Marshal(b, m, deterministic)
}
func (dst *GetTodoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResp.Merge(dst, src)
}
func (m *GetTodoResp) XXX_Size() int {
	return xxx_messageInfo_GetTodoResp.Size(m)
}
func (m *GetTodoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResp proto.InternalMessageInfo

func (m *GetTodoResp) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type Todo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Author               string   `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Done                 bool     `protobuf:"varint,4,opt,name=done" json:"done,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_7d6f602f88ab0b11, []int{1}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Todo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type GetTodoReq struct {
	Int                  int64    `protobuf:"varint,1,opt,name=int" json:"int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoReq) Reset()         { *m = GetTodoReq{} }
func (m *GetTodoReq) String() string { return proto.CompactTextString(m) }
func (*GetTodoReq) ProtoMessage()    {}
func (*GetTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_7d6f602f88ab0b11, []int{2}
}
func (m *GetTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoReq.Unmarshal(m, b)
}
func (m *GetTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoReq.Marshal(b, m, deterministic)
}
func (dst *GetTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoReq.Merge(dst, src)
}
func (m *GetTodoReq) XXX_Size() int {
	return xxx_messageInfo_GetTodoReq.Size(m)
}
func (m *GetTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoReq proto.InternalMessageInfo

func (m *GetTodoReq) GetInt() int64 {
	if m != nil {
		return m.Int
	}
	return 0
}

type ListTodosResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosResponse) Reset()         { *m = ListTodosResponse{} }
func (m *ListTodosResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodosResponse) ProtoMessage()    {}
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todolist_7d6f602f88ab0b11, []int{3}
}
func (m *ListTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosResponse.Unmarshal(m, b)
}
func (m *ListTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosResponse.Marshal(b, m, deterministic)
}
func (dst *ListTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosResponse.Merge(dst, src)
}
func (m *ListTodosResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodosResponse.Size(m)
}
func (m *ListTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosResponse proto.InternalMessageInfo

func (m *ListTodosResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTodoResp)(nil), "todolistpb.GetTodoResp")
	proto.RegisterType((*Todo)(nil), "todolistpb.Todo")
	proto.RegisterType((*GetTodoReq)(nil), "todolistpb.GetTodoReq")
	proto.RegisterType((*ListTodosResponse)(nil), "todolistpb.ListTodosResponse")
}

func init() { proto.RegisterFile("todolist.proto", fileDescriptor_todolist_7d6f602f88ab0b11) }

var fileDescriptor_todolist_7d6f602f88ab0b11 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x66, 0x26, 0x69, 0x6f, 0x7b, 0x7a, 0x29, 0xbd, 0xe7, 0x72, 0x7b, 0x43, 0xb4, 0x12, 0x82,
	0x48, 0x57, 0x09, 0xb4, 0x4b, 0x57, 0x0a, 0xe2, 0x46, 0x10, 0x82, 0x1b, 0xdd, 0x48, 0xd2, 0x19,
	0xeb, 0x40, 0xcd, 0x8c, 0x99, 0xe9, 0xc2, 0xad, 0xbe, 0x40, 0xc1, 0x77, 0xf1, 0x45, 0x7c, 0x05,
	0x1f, 0x44, 0x66, 0x92, 0xfe, 0x80, 0xdd, 0xcd, 0xf7, 0x33, 0xdf, 0x39, 0xdf, 0x81, 0xbe, 0x91,
	0x4c, 0x2e, 0x84, 0x36, 0x89, 0xaa, 0xa4, 0x91, 0x08, 0x6b, 0xac, 0x8a, 0xf0, 0x60, 0x2e, 0xe5,
	0x7c, 0xc1, 0x53, 0xa7, 0x14, 0xcb, 0x87, 0x94, 0x3f, 0x29, 0xf3, 0x52, 0x1b, 0xc3, 0xc3, 0x46,
	0xcc, 0x95, 0x48, 0xf3, 0xb2, 0x94, 0x26, 0x37, 0x42, 0x96, 0xba, 0x56, 0xe3, 0x29, 0xf4, 0x2e,
	0xb9, 0xb9, 0x91, 0x4c, 0x66, 0x5c, 0x2b, 0x3c, 0x06, 0xdf, 0xe6, 0x06, 0x24, 0x22, 0xe3, 0xde,
	0x64, 0x90, 0x6c, 0x87, 0x24, 0xce, 0xe3, 0xd4, 0xf8, 0x8d, 0x80, 0x6f, 0x21, 0xf6, 0x81, 0x0a,
	0xe6, 0xcc, 0x5e, 0x46, 0x05, 0xc3, 0x21, 0xb4, 0xf3, 0xa5, 0x79, 0x94, 0x55, 0x40, 0x23, 0x32,
	0xee, 0x66, 0x0d, 0xc2, 0x08, 0x7a, 0x8c, 0xeb, 0x59, 0x25, 0x94, 0x9d, 0x1d, 0x78, 0x4e, 0xdc,
	0xa5, 0x10, 0xc1, 0x67, 0xb2, 0xe4, 0x81, 0x1f, 0x91, 0x71, 0x27, 0x73, 0x6f, 0x1c, 0x01, 0xcc,
	0x2a, 0x9e, 0x1b, 0xce, 0xee, 0x73, 0x13, 0xb4, 0xdc, 0xa7, 0x6e, 0xc3, 0x9c, 0x99, 0xf8, 0x08,
	0x60, 0xb3, 0xfa, 0x33, 0x0e, 0xc0, 0x13, 0xa5, 0x69, 0x76, 0xb1, 0xcf, 0xf8, 0x14, 0xfe, 0x5c,
	0x09, 0xed, 0x0c, 0xda, 0x96, 0x93, 0xa5, 0xe6, 0x78, 0x02, 0x2d, 0x5b, 0x41, 0x07, 0x24, 0xf2,
	0xf6, 0x36, 0xac, 0xe5, 0xc9, 0x07, 0x81, 0x8e, 0xc5, 0x36, 0x01, 0x6f, 0xa1, 0xbb, 0x49, 0xc2,
	0x61, 0x52, 0x1f, 0x34, 0x59, 0x5f, 0x3b, 0xb9, 0xb0, 0xd7, 0x0e, 0x47, 0xbb, 0x51, 0x3f, 0x06,
	0xc7, 0xff, 0x5e, 0x3f, 0xbf, 0xde, 0xe9, 0x6f, 0x84, 0x74, 0xb1, 0xd6, 0x56, 0x94, 0xe0, 0x35,
	0xfc, 0x6a, 0x4a, 0xe0, 0x70, 0x37, 0x60, 0xdb, 0x2c, 0xfc, 0xbf, 0x97, 0xd7, 0x2a, 0xfe, 0xeb,
	0x22, 0x01, 0x3b, 0xe9, 0xbc, 0x66, 0x57, 0x94, 0x9c, 0xfb, 0x77, 0x54, 0x15, 0x45, 0xdb, 0x2d,
	0x37, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x48, 0x0a, 0xf3, 0x37, 0x36, 0x02, 0x00, 0x00,
}

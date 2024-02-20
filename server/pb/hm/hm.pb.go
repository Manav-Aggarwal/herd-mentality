// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hm.proto

package hm

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QuestionRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QuestionRequest) Reset()         { *m = QuestionRequest{} }
func (m *QuestionRequest) String() string { return proto.CompactTextString(m) }
func (*QuestionRequest) ProtoMessage()    {}
func (*QuestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5435a984cba607c, []int{0}
}
func (m *QuestionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuestionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionRequest.Merge(m, src)
}
func (m *QuestionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionRequest proto.InternalMessageInfo

func (m *QuestionRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Question struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Question string `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5435a984cba607c, []int{1}
}
func (m *Question) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Question.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return m.Size()
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Question) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

type Answer struct {
	QuestionId uint64 `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	UserName   string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Answer     string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5435a984cba607c, []int{2}
}
func (m *Answer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return m.Size()
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetQuestionId() uint64 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

func (m *Answer) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Answer) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type Confirmation struct {
}

func (m *Confirmation) Reset()         { *m = Confirmation{} }
func (m *Confirmation) String() string { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()    {}
func (*Confirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5435a984cba607c, []int{3}
}
func (m *Confirmation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Confirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Confirmation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Confirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Confirmation.Merge(m, src)
}
func (m *Confirmation) XXX_Size() int {
	return m.Size()
}
func (m *Confirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_Confirmation.DiscardUnknown(m)
}

var xxx_messageInfo_Confirmation proto.InternalMessageInfo

type ResultsRequest struct {
	QuestionId uint64 `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
}

func (m *ResultsRequest) Reset()         { *m = ResultsRequest{} }
func (m *ResultsRequest) String() string { return proto.CompactTextString(m) }
func (*ResultsRequest) ProtoMessage()    {}
func (*ResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5435a984cba607c, []int{4}
}
func (m *ResultsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResultsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultsRequest.Merge(m, src)
}
func (m *ResultsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResultsRequest proto.InternalMessageInfo

func (m *ResultsRequest) GetQuestionId() uint64 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

type Results struct {
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5435a984cba607c, []int{5}
}
func (m *Results) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Results.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(m, src)
}
func (m *Results) XXX_Size() int {
	return m.Size()
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QuestionRequest)(nil), "hm.QuestionRequest")
	proto.RegisterType((*Question)(nil), "hm.Question")
	proto.RegisterType((*Answer)(nil), "hm.Answer")
	proto.RegisterType((*Confirmation)(nil), "hm.Confirmation")
	proto.RegisterType((*ResultsRequest)(nil), "hm.ResultsRequest")
	proto.RegisterType((*Results)(nil), "hm.Results")
}

func init() { proto.RegisterFile("hm.proto", fileDescriptor_d5435a984cba607c) }

var fileDescriptor_d5435a984cba607c = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xdb, 0xde, 0x1b, 0x84, 0x03, 0x82, 0x19, 0x13, 0x43, 0x6a, 0x52, 0xb1, 0x2b, 0x36,
	0xb4, 0x51, 0x23, 0x7b, 0x64, 0xa1, 0x2e, 0xd0, 0x58, 0x77, 0x2e, 0x24, 0x53, 0x7b, 0xa4, 0x93,
	0x30, 0x53, 0x98, 0x4e, 0x21, 0xbe, 0x85, 0x8f, 0xe1, 0xa3, 0xb8, 0x64, 0xe9, 0xd2, 0xc0, 0x8b,
	0x98, 0xd6, 0x96, 0x00, 0x1b, 0x77, 0xe7, 0x7c, 0xdf, 0xf9, 0xf7, 0x9b, 0x0c, 0x94, 0x43, 0xee,
	0x4c, 0x64, 0xa4, 0x22, 0x62, 0x84, 0xdc, 0x3e, 0x85, 0xc6, 0x43, 0x82, 0xb1, 0x62, 0x91, 0xf0,
	0x70, 0x9a, 0x46, 0xa4, 0x0e, 0x06, 0x0b, 0x9a, 0x7a, 0x4b, 0x6f, 0xff, 0xf7, 0x0c, 0x16, 0xd8,
	0x5d, 0x28, 0x17, 0x25, 0xbb, 0x1e, 0x31, 0xa1, 0x3c, 0xcd, 0xbd, 0xa6, 0xd1, 0xd2, 0xdb, 0x15,
	0x6f, 0x9d, 0xdb, 0xcf, 0x50, 0xea, 0x89, 0x78, 0x8e, 0x92, 0x9c, 0x40, 0xb5, 0x50, 0x87, 0xeb,
	0x76, 0x28, 0xa4, 0xdb, 0x80, 0x1c, 0x43, 0x25, 0x89, 0x51, 0x0e, 0x05, 0xe5, 0x58, 0xcc, 0x49,
	0x85, 0x3b, 0xca, 0x91, 0x1c, 0x41, 0x89, 0x66, 0x73, 0x9a, 0xff, 0x32, 0x27, 0xcf, 0xec, 0x3a,
	0xd4, 0xfa, 0x91, 0x78, 0x65, 0x92, 0xd3, 0x6c, 0xdf, 0x19, 0xd4, 0x3d, 0x8c, 0x93, 0xb1, 0x8a,
	0x0b, 0x92, 0xbf, 0xf6, 0xda, 0x15, 0xd8, 0xcb, 0x5b, 0xce, 0x3f, 0x74, 0xd8, 0xbf, 0x41, 0x19,
	0x0c, 0x50, 0x28, 0x3a, 0x66, 0xea, 0x8d, 0x74, 0xa1, 0xd1, 0x4f, 0xa4, 0x44, 0xa1, 0xd6, 0xf8,
	0x87, 0x4e, 0xc8, 0x9d, 0x9d, 0xf7, 0x32, 0x6b, 0x9b, 0xa2, 0xad, 0x11, 0x07, 0x6a, 0x8f, 0x89,
	0xcf, 0x99, 0xca, 0xe9, 0x21, 0xf5, 0x7f, 0x63, 0xf3, 0x20, 0x8d, 0xb7, 0xae, 0xd6, 0x88, 0x0b,
	0x70, 0x8d, 0x2a, 0xbf, 0x83, 0x90, 0xb4, 0x62, 0x9b, 0xc3, 0xac, 0x6e, 0x68, 0xb6, 0x76, 0x75,
	0xff, 0xb9, 0xb4, 0xf4, 0xc5, 0xd2, 0xd2, 0xbf, 0x97, 0x96, 0xfe, 0xbe, 0xb2, 0xb4, 0xc5, 0xca,
	0xd2, 0xbe, 0x56, 0x96, 0xf6, 0x74, 0x39, 0x62, 0x2a, 0x4c, 0x7c, 0xe7, 0x25, 0xe2, 0xee, 0x80,
	0x0a, 0x3a, 0xeb, 0xf4, 0x46, 0x23, 0x2a, 0xe7, 0x74, 0xec, 0x86, 0x28, 0x83, 0x0e, 0x2f, 0xe0,
	0xdc, 0x18, 0xe5, 0x0c, 0xa5, 0x3b, 0xf1, 0xdd, 0x90, 0xfb, 0xa5, 0xec, 0x3f, 0x5c, 0xfc, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xa7, 0xe6, 0x82, 0x46, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HerdMentalityClient is the client API for HerdMentality service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HerdMentalityClient interface {
	CurrentQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*Question, error)
	SubmitAnswer(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*Confirmation, error)
	GetResults(ctx context.Context, in *ResultsRequest, opts ...grpc.CallOption) (*Results, error)
}

type herdMentalityClient struct {
	cc *grpc.ClientConn
}

func NewHerdMentalityClient(cc *grpc.ClientConn) HerdMentalityClient {
	return &herdMentalityClient{cc}
}

func (c *herdMentalityClient) CurrentQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := c.cc.Invoke(ctx, "/hm.HerdMentality/CurrentQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *herdMentalityClient) SubmitAnswer(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/hm.HerdMentality/SubmitAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *herdMentalityClient) GetResults(ctx context.Context, in *ResultsRequest, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/hm.HerdMentality/GetResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HerdMentalityServer is the server API for HerdMentality service.
type HerdMentalityServer interface {
	CurrentQuestion(context.Context, *QuestionRequest) (*Question, error)
	SubmitAnswer(context.Context, *Answer) (*Confirmation, error)
	GetResults(context.Context, *ResultsRequest) (*Results, error)
}

// UnimplementedHerdMentalityServer can be embedded to have forward compatible implementations.
type UnimplementedHerdMentalityServer struct {
}

func (*UnimplementedHerdMentalityServer) CurrentQuestion(ctx context.Context, req *QuestionRequest) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentQuestion not implemented")
}
func (*UnimplementedHerdMentalityServer) SubmitAnswer(ctx context.Context, req *Answer) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswer not implemented")
}
func (*UnimplementedHerdMentalityServer) GetResults(ctx context.Context, req *ResultsRequest) (*Results, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResults not implemented")
}

func RegisterHerdMentalityServer(s *grpc.Server, srv HerdMentalityServer) {
	s.RegisterService(&_HerdMentality_serviceDesc, srv)
}

func _HerdMentality_CurrentQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HerdMentalityServer).CurrentQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hm.HerdMentality/CurrentQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HerdMentalityServer).CurrentQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HerdMentality_SubmitAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Answer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HerdMentalityServer).SubmitAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hm.HerdMentality/SubmitAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HerdMentalityServer).SubmitAnswer(ctx, req.(*Answer))
	}
	return interceptor(ctx, in, info, handler)
}

func _HerdMentality_GetResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HerdMentalityServer).GetResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hm.HerdMentality/GetResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HerdMentalityServer).GetResults(ctx, req.(*ResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HerdMentality_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hm.HerdMentality",
	HandlerType: (*HerdMentalityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentQuestion",
			Handler:    _HerdMentality_CurrentQuestion_Handler,
		},
		{
			MethodName: "SubmitAnswer",
			Handler:    _HerdMentality_SubmitAnswer_Handler,
		},
		{
			MethodName: "GetResults",
			Handler:    _HerdMentality_GetResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hm.proto",
}

func (m *QuestionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuestionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintHm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Question) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Question) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Question) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintHm(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintHm(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Answer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Answer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Answer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Answer) > 0 {
		i -= len(m.Answer)
		copy(dAtA[i:], m.Answer)
		i = encodeVarintHm(dAtA, i, uint64(len(m.Answer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UserName) > 0 {
		i -= len(m.UserName)
		copy(dAtA[i:], m.UserName)
		i = encodeVarintHm(dAtA, i, uint64(len(m.UserName)))
		i--
		dAtA[i] = 0x12
	}
	if m.QuestionId != 0 {
		i = encodeVarintHm(dAtA, i, uint64(m.QuestionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Confirmation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Confirmation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Confirmation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ResultsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResultsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResultsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QuestionId != 0 {
		i = encodeVarintHm(dAtA, i, uint64(m.QuestionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Results) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Results) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Results) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintHm(dAtA []byte, offset int, v uint64) int {
	offset -= sovHm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuestionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHm(uint64(m.Id))
	}
	return n
}

func (m *Question) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHm(uint64(m.Id))
	}
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovHm(uint64(l))
	}
	return n
}

func (m *Answer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuestionId != 0 {
		n += 1 + sovHm(uint64(m.QuestionId))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovHm(uint64(l))
	}
	l = len(m.Answer)
	if l > 0 {
		n += 1 + l + sovHm(uint64(l))
	}
	return n
}

func (m *Confirmation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ResultsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuestionId != 0 {
		n += 1 + sovHm(uint64(m.QuestionId))
	}
	return n
}

func (m *Results) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovHm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHm(x uint64) (n int) {
	return sovHm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuestionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHm
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
			return fmt.Errorf("proto: QuestionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHm
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
func (m *Question) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHm
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
			return fmt.Errorf("proto: Question: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Question: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
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
				return ErrInvalidLengthHm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHm
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
func (m *Answer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHm
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
			return fmt.Errorf("proto: Answer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Answer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuestionId", wireType)
			}
			m.QuestionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuestionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
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
				return ErrInvalidLengthHm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
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
				return ErrInvalidLengthHm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHm
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
func (m *Confirmation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHm
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
			return fmt.Errorf("proto: Confirmation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Confirmation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHm
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
func (m *ResultsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHm
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
			return fmt.Errorf("proto: ResultsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuestionId", wireType)
			}
			m.QuestionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuestionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHm
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
func (m *Results) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHm
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
			return fmt.Errorf("proto: Results: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Results: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHm
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
func skipHm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHm
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
					return 0, ErrIntOverflowHm
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
					return 0, ErrIntOverflowHm
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
				return 0, ErrInvalidLengthHm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHm = fmt.Errorf("proto: unexpected end of group")
)
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: board.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BoardClient is the client API for Board service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardClient interface {
	ListSubjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SubjectList, error)
	GetSubject(ctx context.Context, in *SubjectId, opts ...grpc.CallOption) (*Subject, error)
	ListQuestions(ctx context.Context, in *SubjectId, opts ...grpc.CallOption) (*QuestionList, error)
	CreateQuestion(ctx context.Context, in *NewQuestion, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Like(ctx context.Context, in *QuestionId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unlike(ctx context.Context, in *QuestionId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type boardClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardClient(cc grpc.ClientConnInterface) BoardClient {
	return &boardClient{cc}
}

func (c *boardClient) ListSubjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SubjectList, error) {
	out := new(SubjectList)
	err := c.cc.Invoke(ctx, "/board.Board/ListSubjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) GetSubject(ctx context.Context, in *SubjectId, opts ...grpc.CallOption) (*Subject, error) {
	out := new(Subject)
	err := c.cc.Invoke(ctx, "/board.Board/GetSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) ListQuestions(ctx context.Context, in *SubjectId, opts ...grpc.CallOption) (*QuestionList, error) {
	out := new(QuestionList)
	err := c.cc.Invoke(ctx, "/board.Board/ListQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) CreateQuestion(ctx context.Context, in *NewQuestion, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/board.Board/CreateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Like(ctx context.Context, in *QuestionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/board.Board/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Unlike(ctx context.Context, in *QuestionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/board.Board/Unlike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServer is the server API for Board service.
// All implementations must embed UnimplementedBoardServer
// for forward compatibility
type BoardServer interface {
	ListSubjects(context.Context, *emptypb.Empty) (*SubjectList, error)
	GetSubject(context.Context, *SubjectId) (*Subject, error)
	ListQuestions(context.Context, *SubjectId) (*QuestionList, error)
	CreateQuestion(context.Context, *NewQuestion) (*emptypb.Empty, error)
	Like(context.Context, *QuestionId) (*emptypb.Empty, error)
	Unlike(context.Context, *QuestionId) (*emptypb.Empty, error)
	mustEmbedUnimplementedBoardServer()
}

// UnimplementedBoardServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServer struct {
}

func (UnimplementedBoardServer) ListSubjects(context.Context, *emptypb.Empty) (*SubjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubjects not implemented")
}
func (UnimplementedBoardServer) GetSubject(context.Context, *SubjectId) (*Subject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubject not implemented")
}
func (UnimplementedBoardServer) ListQuestions(context.Context, *SubjectId) (*QuestionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestions not implemented")
}
func (UnimplementedBoardServer) CreateQuestion(context.Context, *NewQuestion) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedBoardServer) Like(context.Context, *QuestionId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedBoardServer) Unlike(context.Context, *QuestionId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlike not implemented")
}
func (UnimplementedBoardServer) mustEmbedUnimplementedBoardServer() {}

// UnsafeBoardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServer will
// result in compilation errors.
type UnsafeBoardServer interface {
	mustEmbedUnimplementedBoardServer()
}

func RegisterBoardServer(s grpc.ServiceRegistrar, srv BoardServer) {
	s.RegisterService(&Board_ServiceDesc, srv)
}

func _Board_ListSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).ListSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.Board/ListSubjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).ListSubjects(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_GetSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.Board/GetSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetSubject(ctx, req.(*SubjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_ListQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).ListQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.Board/ListQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).ListQuestions(ctx, req.(*SubjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewQuestion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.Board/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).CreateQuestion(ctx, req.(*NewQuestion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.Board/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Like(ctx, req.(*QuestionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Unlike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Unlike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/board.Board/Unlike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Unlike(ctx, req.(*QuestionId))
	}
	return interceptor(ctx, in, info, handler)
}

// Board_ServiceDesc is the grpc.ServiceDesc for Board service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Board_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "board.Board",
	HandlerType: (*BoardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSubjects",
			Handler:    _Board_ListSubjects_Handler,
		},
		{
			MethodName: "GetSubject",
			Handler:    _Board_GetSubject_Handler,
		},
		{
			MethodName: "ListQuestions",
			Handler:    _Board_ListQuestions_Handler,
		},
		{
			MethodName: "CreateQuestion",
			Handler:    _Board_CreateQuestion_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _Board_Like_Handler,
		},
		{
			MethodName: "Unlike",
			Handler:    _Board_Unlike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "board.proto",
}
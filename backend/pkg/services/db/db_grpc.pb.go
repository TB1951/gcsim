// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: protos/backend/db.proto

package db

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DBStore_Get_FullMethodName                    = "/db.DBStore/Get"
	DBStore_GetAll_FullMethodName                 = "/db.DBStore/GetAll"
	DBStore_GetOne_FullMethodName                 = "/db.DBStore/GetOne"
	DBStore_GetPending_FullMethodName             = "/db.DBStore/GetPending"
	DBStore_GetBySubmitter_FullMethodName         = "/db.DBStore/GetBySubmitter"
	DBStore_ApproveTag_FullMethodName             = "/db.DBStore/ApproveTag"
	DBStore_RejectTag_FullMethodName              = "/db.DBStore/RejectTag"
	DBStore_RejectTagAllUnapproved_FullMethodName = "/db.DBStore/RejectTagAllUnapproved"
	DBStore_Submit_FullMethodName                 = "/db.DBStore/Submit"
	DBStore_DeletePending_FullMethodName          = "/db.DBStore/DeletePending"
	DBStore_GetWork_FullMethodName                = "/db.DBStore/GetWork"
	DBStore_CompleteWork_FullMethodName           = "/db.DBStore/CompleteWork"
	DBStore_RejectWork_FullMethodName             = "/db.DBStore/RejectWork"
	DBStore_WorkStatus_FullMethodName             = "/db.DBStore/WorkStatus"
	DBStore_ReplaceConfig_FullMethodName          = "/db.DBStore/ReplaceConfig"
	DBStore_ReplaceDesc_FullMethodName            = "/db.DBStore/ReplaceDesc"
)

// DBStoreClient is the client API for DBStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBStoreClient interface {
	// generic get for pulling from approved db
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetPending(ctx context.Context, in *GetPendingRequest, opts ...grpc.CallOption) (*GetPendingResponse, error)
	GetBySubmitter(ctx context.Context, in *GetBySubmitterRequest, opts ...grpc.CallOption) (*GetBySubmitterResponse, error)
	// tagging
	ApproveTag(ctx context.Context, in *ApproveTagRequest, opts ...grpc.CallOption) (*ApproveTagResponse, error)
	RejectTag(ctx context.Context, in *RejectTagRequest, opts ...grpc.CallOption) (*RejectTagResponse, error)
	RejectTagAllUnapproved(ctx context.Context, in *RejectTagAllUnapprovedRequest, opts ...grpc.CallOption) (*RejectTagAllUnapprovedResponse, error)
	// submissions
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	DeletePending(ctx context.Context, in *DeletePendingRequest, opts ...grpc.CallOption) (*DeletePendingResponse, error)
	// work related
	GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkResponse, error)
	CompleteWork(ctx context.Context, in *CompleteWorkRequest, opts ...grpc.CallOption) (*CompleteWorkResponse, error)
	RejectWork(ctx context.Context, in *RejectWorkRequest, opts ...grpc.CallOption) (*RejectWorkResponse, error)
	WorkStatus(ctx context.Context, in *WorkStatusRequest, opts ...grpc.CallOption) (*WorkStatusResponse, error)
	// admin endpoint
	ReplaceConfig(ctx context.Context, in *ReplaceConfigRequest, opts ...grpc.CallOption) (*ReplaceConfigResponse, error)
	ReplaceDesc(ctx context.Context, in *ReplaceDescRequest, opts ...grpc.CallOption) (*ReplaceDescResponse, error)
}

type dBStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewDBStoreClient(cc grpc.ClientConnInterface) DBStoreClient {
	return &dBStoreClient{cc}
}

func (c *dBStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, DBStore_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, DBStore_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, DBStore_GetOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetPending(ctx context.Context, in *GetPendingRequest, opts ...grpc.CallOption) (*GetPendingResponse, error) {
	out := new(GetPendingResponse)
	err := c.cc.Invoke(ctx, DBStore_GetPending_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetBySubmitter(ctx context.Context, in *GetBySubmitterRequest, opts ...grpc.CallOption) (*GetBySubmitterResponse, error) {
	out := new(GetBySubmitterResponse)
	err := c.cc.Invoke(ctx, DBStore_GetBySubmitter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) ApproveTag(ctx context.Context, in *ApproveTagRequest, opts ...grpc.CallOption) (*ApproveTagResponse, error) {
	out := new(ApproveTagResponse)
	err := c.cc.Invoke(ctx, DBStore_ApproveTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) RejectTag(ctx context.Context, in *RejectTagRequest, opts ...grpc.CallOption) (*RejectTagResponse, error) {
	out := new(RejectTagResponse)
	err := c.cc.Invoke(ctx, DBStore_RejectTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) RejectTagAllUnapproved(ctx context.Context, in *RejectTagAllUnapprovedRequest, opts ...grpc.CallOption) (*RejectTagAllUnapprovedResponse, error) {
	out := new(RejectTagAllUnapprovedResponse)
	err := c.cc.Invoke(ctx, DBStore_RejectTagAllUnapproved_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, DBStore_Submit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) DeletePending(ctx context.Context, in *DeletePendingRequest, opts ...grpc.CallOption) (*DeletePendingResponse, error) {
	out := new(DeletePendingResponse)
	err := c.cc.Invoke(ctx, DBStore_DeletePending_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*GetWorkResponse, error) {
	out := new(GetWorkResponse)
	err := c.cc.Invoke(ctx, DBStore_GetWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) CompleteWork(ctx context.Context, in *CompleteWorkRequest, opts ...grpc.CallOption) (*CompleteWorkResponse, error) {
	out := new(CompleteWorkResponse)
	err := c.cc.Invoke(ctx, DBStore_CompleteWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) RejectWork(ctx context.Context, in *RejectWorkRequest, opts ...grpc.CallOption) (*RejectWorkResponse, error) {
	out := new(RejectWorkResponse)
	err := c.cc.Invoke(ctx, DBStore_RejectWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) WorkStatus(ctx context.Context, in *WorkStatusRequest, opts ...grpc.CallOption) (*WorkStatusResponse, error) {
	out := new(WorkStatusResponse)
	err := c.cc.Invoke(ctx, DBStore_WorkStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) ReplaceConfig(ctx context.Context, in *ReplaceConfigRequest, opts ...grpc.CallOption) (*ReplaceConfigResponse, error) {
	out := new(ReplaceConfigResponse)
	err := c.cc.Invoke(ctx, DBStore_ReplaceConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBStoreClient) ReplaceDesc(ctx context.Context, in *ReplaceDescRequest, opts ...grpc.CallOption) (*ReplaceDescResponse, error) {
	out := new(ReplaceDescResponse)
	err := c.cc.Invoke(ctx, DBStore_ReplaceDesc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBStoreServer is the server API for DBStore service.
// All implementations must embed UnimplementedDBStoreServer
// for forward compatibility
type DBStoreServer interface {
	// generic get for pulling from approved db
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetPending(context.Context, *GetPendingRequest) (*GetPendingResponse, error)
	GetBySubmitter(context.Context, *GetBySubmitterRequest) (*GetBySubmitterResponse, error)
	// tagging
	ApproveTag(context.Context, *ApproveTagRequest) (*ApproveTagResponse, error)
	RejectTag(context.Context, *RejectTagRequest) (*RejectTagResponse, error)
	RejectTagAllUnapproved(context.Context, *RejectTagAllUnapprovedRequest) (*RejectTagAllUnapprovedResponse, error)
	// submissions
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	DeletePending(context.Context, *DeletePendingRequest) (*DeletePendingResponse, error)
	// work related
	GetWork(context.Context, *GetWorkRequest) (*GetWorkResponse, error)
	CompleteWork(context.Context, *CompleteWorkRequest) (*CompleteWorkResponse, error)
	RejectWork(context.Context, *RejectWorkRequest) (*RejectWorkResponse, error)
	WorkStatus(context.Context, *WorkStatusRequest) (*WorkStatusResponse, error)
	// admin endpoint
	ReplaceConfig(context.Context, *ReplaceConfigRequest) (*ReplaceConfigResponse, error)
	ReplaceDesc(context.Context, *ReplaceDescRequest) (*ReplaceDescResponse, error)
	mustEmbedUnimplementedDBStoreServer()
}

// UnimplementedDBStoreServer must be embedded to have forward compatible implementations.
type UnimplementedDBStoreServer struct {
}

func (UnimplementedDBStoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDBStoreServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedDBStoreServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedDBStoreServer) GetPending(context.Context, *GetPendingRequest) (*GetPendingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPending not implemented")
}
func (UnimplementedDBStoreServer) GetBySubmitter(context.Context, *GetBySubmitterRequest) (*GetBySubmitterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySubmitter not implemented")
}
func (UnimplementedDBStoreServer) ApproveTag(context.Context, *ApproveTagRequest) (*ApproveTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveTag not implemented")
}
func (UnimplementedDBStoreServer) RejectTag(context.Context, *RejectTagRequest) (*RejectTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectTag not implemented")
}
func (UnimplementedDBStoreServer) RejectTagAllUnapproved(context.Context, *RejectTagAllUnapprovedRequest) (*RejectTagAllUnapprovedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectTagAllUnapproved not implemented")
}
func (UnimplementedDBStoreServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedDBStoreServer) DeletePending(context.Context, *DeletePendingRequest) (*DeletePendingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePending not implemented")
}
func (UnimplementedDBStoreServer) GetWork(context.Context, *GetWorkRequest) (*GetWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWork not implemented")
}
func (UnimplementedDBStoreServer) CompleteWork(context.Context, *CompleteWorkRequest) (*CompleteWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteWork not implemented")
}
func (UnimplementedDBStoreServer) RejectWork(context.Context, *RejectWorkRequest) (*RejectWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectWork not implemented")
}
func (UnimplementedDBStoreServer) WorkStatus(context.Context, *WorkStatusRequest) (*WorkStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkStatus not implemented")
}
func (UnimplementedDBStoreServer) ReplaceConfig(context.Context, *ReplaceConfigRequest) (*ReplaceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceConfig not implemented")
}
func (UnimplementedDBStoreServer) ReplaceDesc(context.Context, *ReplaceDescRequest) (*ReplaceDescResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceDesc not implemented")
}
func (UnimplementedDBStoreServer) mustEmbedUnimplementedDBStoreServer() {}

// UnsafeDBStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBStoreServer will
// result in compilation errors.
type UnsafeDBStoreServer interface {
	mustEmbedUnimplementedDBStoreServer()
}

func RegisterDBStoreServer(s grpc.ServiceRegistrar, srv DBStoreServer) {
	s.RegisterService(&DBStore_ServiceDesc, srv)
}

func _DBStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetPending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetPending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_GetPending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetPending(ctx, req.(*GetPendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetBySubmitter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBySubmitterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetBySubmitter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_GetBySubmitter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetBySubmitter(ctx, req.(*GetBySubmitterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_ApproveTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).ApproveTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_ApproveTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).ApproveTag(ctx, req.(*ApproveTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_RejectTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).RejectTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_RejectTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).RejectTag(ctx, req.(*RejectTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_RejectTagAllUnapproved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectTagAllUnapprovedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).RejectTagAllUnapproved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_RejectTagAllUnapproved_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).RejectTagAllUnapproved(ctx, req.(*RejectTagAllUnapprovedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_DeletePending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).DeletePending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_DeletePending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).DeletePending(ctx, req.(*DeletePendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_GetWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).GetWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_GetWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).GetWork(ctx, req.(*GetWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_CompleteWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).CompleteWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_CompleteWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).CompleteWork(ctx, req.(*CompleteWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_RejectWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).RejectWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_RejectWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).RejectWork(ctx, req.(*RejectWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_WorkStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).WorkStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_WorkStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).WorkStatus(ctx, req.(*WorkStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_ReplaceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).ReplaceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_ReplaceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).ReplaceConfig(ctx, req.(*ReplaceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBStore_ReplaceDesc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceDescRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBStoreServer).ReplaceDesc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBStore_ReplaceDesc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBStoreServer).ReplaceDesc(ctx, req.(*ReplaceDescRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBStore_ServiceDesc is the grpc.ServiceDesc for DBStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DBStore",
	HandlerType: (*DBStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DBStore_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _DBStore_GetAll_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _DBStore_GetOne_Handler,
		},
		{
			MethodName: "GetPending",
			Handler:    _DBStore_GetPending_Handler,
		},
		{
			MethodName: "GetBySubmitter",
			Handler:    _DBStore_GetBySubmitter_Handler,
		},
		{
			MethodName: "ApproveTag",
			Handler:    _DBStore_ApproveTag_Handler,
		},
		{
			MethodName: "RejectTag",
			Handler:    _DBStore_RejectTag_Handler,
		},
		{
			MethodName: "RejectTagAllUnapproved",
			Handler:    _DBStore_RejectTagAllUnapproved_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _DBStore_Submit_Handler,
		},
		{
			MethodName: "DeletePending",
			Handler:    _DBStore_DeletePending_Handler,
		},
		{
			MethodName: "GetWork",
			Handler:    _DBStore_GetWork_Handler,
		},
		{
			MethodName: "CompleteWork",
			Handler:    _DBStore_CompleteWork_Handler,
		},
		{
			MethodName: "RejectWork",
			Handler:    _DBStore_RejectWork_Handler,
		},
		{
			MethodName: "WorkStatus",
			Handler:    _DBStore_WorkStatus_Handler,
		},
		{
			MethodName: "ReplaceConfig",
			Handler:    _DBStore_ReplaceConfig_Handler,
		},
		{
			MethodName: "ReplaceDesc",
			Handler:    _DBStore_ReplaceDesc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/backend/db.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: message.proto

package proto

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
	Continents_InformePersona_FullMethodName = "/grpc.Continents/InformePersona"
)

// ContinentsClient is the client API for Continents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContinentsClient interface {
	InformePersona(ctx context.Context, in *Informe, opts ...grpc.CallOption) (*RespuestaVacia, error)
}

type continentsClient struct {
	cc grpc.ClientConnInterface
}

func NewContinentsClient(cc grpc.ClientConnInterface) ContinentsClient {
	return &continentsClient{cc}
}

func (c *continentsClient) InformePersona(ctx context.Context, in *Informe, opts ...grpc.CallOption) (*RespuestaVacia, error) {
	out := new(RespuestaVacia)
	err := c.cc.Invoke(ctx, Continents_InformePersona_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContinentsServer is the server API for Continents service.
// All implementations must embed UnimplementedContinentsServer
// for forward compatibility
type ContinentsServer interface {
	InformePersona(context.Context, *Informe) (*RespuestaVacia, error)
	mustEmbedUnimplementedContinentsServer()
}

// UnimplementedContinentsServer must be embedded to have forward compatible implementations.
type UnimplementedContinentsServer struct {
}

func (UnimplementedContinentsServer) InformePersona(context.Context, *Informe) (*RespuestaVacia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InformePersona not implemented")
}
func (UnimplementedContinentsServer) mustEmbedUnimplementedContinentsServer() {}

// UnsafeContinentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContinentsServer will
// result in compilation errors.
type UnsafeContinentsServer interface {
	mustEmbedUnimplementedContinentsServer()
}

func RegisterContinentsServer(s grpc.ServiceRegistrar, srv ContinentsServer) {
	s.RegisterService(&Continents_ServiceDesc, srv)
}

func _Continents_InformePersona_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Informe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContinentsServer).InformePersona(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Continents_InformePersona_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContinentsServer).InformePersona(ctx, req.(*Informe))
	}
	return interceptor(ctx, in, info, handler)
}

// Continents_ServiceDesc is the grpc.ServiceDesc for Continents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Continents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Continents",
	HandlerType: (*ContinentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InformePersona",
			Handler:    _Continents_InformePersona_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

const (
	DataNode_RegistrarNombre_FullMethodName = "/grpc.DataNode/RegistrarNombre"
	DataNode_PedirNombre_FullMethodName     = "/grpc.DataNode/PedirNombre"
)

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeClient interface {
	RegistrarNombre(ctx context.Context, in *Registro, opts ...grpc.CallOption) (*RespuestaVacia, error)
	PedirNombre(ctx context.Context, in *Peticion, opts ...grpc.CallOption) (*Nombre, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) RegistrarNombre(ctx context.Context, in *Registro, opts ...grpc.CallOption) (*RespuestaVacia, error) {
	out := new(RespuestaVacia)
	err := c.cc.Invoke(ctx, DataNode_RegistrarNombre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) PedirNombre(ctx context.Context, in *Peticion, opts ...grpc.CallOption) (*Nombre, error) {
	out := new(Nombre)
	err := c.cc.Invoke(ctx, DataNode_PedirNombre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
// All implementations must embed UnimplementedDataNodeServer
// for forward compatibility
type DataNodeServer interface {
	RegistrarNombre(context.Context, *Registro) (*RespuestaVacia, error)
	PedirNombre(context.Context, *Peticion) (*Nombre, error)
	mustEmbedUnimplementedDataNodeServer()
}

// UnimplementedDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (UnimplementedDataNodeServer) RegistrarNombre(context.Context, *Registro) (*RespuestaVacia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrarNombre not implemented")
}
func (UnimplementedDataNodeServer) PedirNombre(context.Context, *Peticion) (*Nombre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirNombre not implemented")
}
func (UnimplementedDataNodeServer) mustEmbedUnimplementedDataNodeServer() {}

// UnsafeDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServer will
// result in compilation errors.
type UnsafeDataNodeServer interface {
	mustEmbedUnimplementedDataNodeServer()
}

func RegisterDataNodeServer(s grpc.ServiceRegistrar, srv DataNodeServer) {
	s.RegisterService(&DataNode_ServiceDesc, srv)
}

func _DataNode_RegistrarNombre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).RegistrarNombre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNode_RegistrarNombre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).RegistrarNombre(ctx, req.(*Registro))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNode_PedirNombre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Peticion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).PedirNombre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNode_PedirNombre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).PedirNombre(ctx, req.(*Peticion))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNode_ServiceDesc is the grpc.ServiceDesc for DataNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistrarNombre",
			Handler:    _DataNode_RegistrarNombre_Handler,
		},
		{
			MethodName: "PedirNombre",
			Handler:    _DataNode_PedirNombre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

const (
	ONU_ConsultarNombres_FullMethodName = "/grpc.ONU/ConsultarNombres"
)

// ONUClient is the client API for ONU service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ONUClient interface {
	ConsultarNombres(ctx context.Context, in *Consulta, opts ...grpc.CallOption) (*RespuestaNombres, error)
}

type oNUClient struct {
	cc grpc.ClientConnInterface
}

func NewONUClient(cc grpc.ClientConnInterface) ONUClient {
	return &oNUClient{cc}
}

func (c *oNUClient) ConsultarNombres(ctx context.Context, in *Consulta, opts ...grpc.CallOption) (*RespuestaNombres, error) {
	out := new(RespuestaNombres)
	err := c.cc.Invoke(ctx, ONU_ConsultarNombres_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ONUServer is the server API for ONU service.
// All implementations must embed UnimplementedONUServer
// for forward compatibility
type ONUServer interface {
	ConsultarNombres(context.Context, *Consulta) (*RespuestaNombres, error)
	mustEmbedUnimplementedONUServer()
}

// UnimplementedONUServer must be embedded to have forward compatible implementations.
type UnimplementedONUServer struct {
}

func (UnimplementedONUServer) ConsultarNombres(context.Context, *Consulta) (*RespuestaNombres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarNombres not implemented")
}
func (UnimplementedONUServer) mustEmbedUnimplementedONUServer() {}

// UnsafeONUServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ONUServer will
// result in compilation errors.
type UnsafeONUServer interface {
	mustEmbedUnimplementedONUServer()
}

func RegisterONUServer(s grpc.ServiceRegistrar, srv ONUServer) {
	s.RegisterService(&ONU_ServiceDesc, srv)
}

func _ONU_ConsultarNombres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consulta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ONUServer).ConsultarNombres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ONU_ConsultarNombres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ONUServer).ConsultarNombres(ctx, req.(*Consulta))
	}
	return interceptor(ctx, in, info, handler)
}

// ONU_ServiceDesc is the grpc.ServiceDesc for ONU service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ONU_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ONU",
	HandlerType: (*ONUServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConsultarNombres",
			Handler:    _ONU_ConsultarNombres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

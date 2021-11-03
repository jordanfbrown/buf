// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.19.1
// source: buf/alpha/registry/v1alpha1/plugin.proto

package registryv1alpha1

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

// PluginServiceClient is the client API for PluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginServiceClient interface {
	// ListPlugins returns all the plugins available to the user. This includes
	// public plugins, those uploaded to organizations the user is part of,
	// and any plugins uploaded directly by the user.
	ListPlugins(ctx context.Context, in *ListPluginsRequest, opts ...grpc.CallOption) (*ListPluginsResponse, error)
	// ListUserPlugins lists all plugins belonging to a user.
	ListUserPlugins(ctx context.Context, in *ListUserPluginsRequest, opts ...grpc.CallOption) (*ListUserPluginsResponse, error)
	// ListOrganizationPlugins lists all plugins for an organization.
	ListOrganizationPlugins(ctx context.Context, in *ListOrganizationPluginsRequest, opts ...grpc.CallOption) (*ListOrganizationPluginsResponse, error)
	// ListPluginVersions lists all the versions available for the specified plugin.
	ListPluginVersions(ctx context.Context, in *ListPluginVersionsRequest, opts ...grpc.CallOption) (*ListPluginVersionsResponse, error)
	// CreatePlugin creates a new plugin.
	CreatePlugin(ctx context.Context, in *CreatePluginRequest, opts ...grpc.CallOption) (*CreatePluginResponse, error)
	// GetPlugin returns the plugin, if found.
	GetPlugin(ctx context.Context, in *GetPluginRequest, opts ...grpc.CallOption) (*GetPluginResponse, error)
	// DeletePlugin deletes the plugin, if it exists. Note that deleting
	// a plugin may cause breaking changes for templates using that plugin,
	// and should be done with extreme care.
	DeletePlugin(ctx context.Context, in *DeletePluginRequest, opts ...grpc.CallOption) (*DeletePluginResponse, error)
	// GetTemplate returns the template, if found.
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error)
	// ListTemplates returns all the templates available to the user. This includes
	// public templates, those owned by organizations the user is part of,
	// and any created directly by the user.
	ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error)
	// ListUserPlugins lists all templates belonging to a user.
	ListUserTemplates(ctx context.Context, in *ListUserTemplatesRequest, opts ...grpc.CallOption) (*ListUserTemplatesResponse, error)
	// ListOrganizationTemplates lists all templates for an organization.
	ListOrganizationTemplates(ctx context.Context, in *ListOrganizationTemplatesRequest, opts ...grpc.CallOption) (*ListOrganizationTemplatesResponse, error)
	// GetTemplateVersion returns the template version, if found.
	GetTemplateVersion(ctx context.Context, in *GetTemplateVersionRequest, opts ...grpc.CallOption) (*GetTemplateVersionResponse, error)
	// ListTemplateVersions lists all the template versions available for the specified template.
	ListTemplateVersions(ctx context.Context, in *ListTemplateVersionsRequest, opts ...grpc.CallOption) (*ListTemplateVersionsResponse, error)
	// CreateTemplate creates a new template.
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	// DeleteTemplate deletes the template, if it exists.
	DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*DeleteTemplateResponse, error)
	// CreateTemplateVersion creates a new template version.
	CreateTemplateVersion(ctx context.Context, in *CreateTemplateVersionRequest, opts ...grpc.CallOption) (*CreateTemplateVersionResponse, error)
}

type pluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginServiceClient(cc grpc.ClientConnInterface) PluginServiceClient {
	return &pluginServiceClient{cc}
}

func (c *pluginServiceClient) ListPlugins(ctx context.Context, in *ListPluginsRequest, opts ...grpc.CallOption) (*ListPluginsResponse, error) {
	out := new(ListPluginsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListUserPlugins(ctx context.Context, in *ListUserPluginsRequest, opts ...grpc.CallOption) (*ListUserPluginsResponse, error) {
	out := new(ListUserPluginsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListUserPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListOrganizationPlugins(ctx context.Context, in *ListOrganizationPluginsRequest, opts ...grpc.CallOption) (*ListOrganizationPluginsResponse, error) {
	out := new(ListOrganizationPluginsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListOrganizationPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListPluginVersions(ctx context.Context, in *ListPluginVersionsRequest, opts ...grpc.CallOption) (*ListPluginVersionsResponse, error) {
	out := new(ListPluginVersionsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListPluginVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) CreatePlugin(ctx context.Context, in *CreatePluginRequest, opts ...grpc.CallOption) (*CreatePluginResponse, error) {
	out := new(CreatePluginResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/CreatePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) GetPlugin(ctx context.Context, in *GetPluginRequest, opts ...grpc.CallOption) (*GetPluginResponse, error) {
	out := new(GetPluginResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/GetPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) DeletePlugin(ctx context.Context, in *DeletePluginRequest, opts ...grpc.CallOption) (*DeletePluginResponse, error) {
	out := new(DeletePluginResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/DeletePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error) {
	out := new(ListTemplatesResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListUserTemplates(ctx context.Context, in *ListUserTemplatesRequest, opts ...grpc.CallOption) (*ListUserTemplatesResponse, error) {
	out := new(ListUserTemplatesResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListUserTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListOrganizationTemplates(ctx context.Context, in *ListOrganizationTemplatesRequest, opts ...grpc.CallOption) (*ListOrganizationTemplatesResponse, error) {
	out := new(ListOrganizationTemplatesResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListOrganizationTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) GetTemplateVersion(ctx context.Context, in *GetTemplateVersionRequest, opts ...grpc.CallOption) (*GetTemplateVersionResponse, error) {
	out := new(GetTemplateVersionResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/GetTemplateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ListTemplateVersions(ctx context.Context, in *ListTemplateVersionsRequest, opts ...grpc.CallOption) (*ListTemplateVersionsResponse, error) {
	out := new(ListTemplateVersionsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/ListTemplateVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*DeleteTemplateResponse, error) {
	out := new(DeleteTemplateResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) CreateTemplateVersion(ctx context.Context, in *CreateTemplateVersionRequest, opts ...grpc.CallOption) (*CreateTemplateVersionResponse, error) {
	out := new(CreateTemplateVersionResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.PluginService/CreateTemplateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServiceServer is the server API for PluginService service.
// All implementations should embed UnimplementedPluginServiceServer
// for forward compatibility
type PluginServiceServer interface {
	// ListPlugins returns all the plugins available to the user. This includes
	// public plugins, those uploaded to organizations the user is part of,
	// and any plugins uploaded directly by the user.
	ListPlugins(context.Context, *ListPluginsRequest) (*ListPluginsResponse, error)
	// ListUserPlugins lists all plugins belonging to a user.
	ListUserPlugins(context.Context, *ListUserPluginsRequest) (*ListUserPluginsResponse, error)
	// ListOrganizationPlugins lists all plugins for an organization.
	ListOrganizationPlugins(context.Context, *ListOrganizationPluginsRequest) (*ListOrganizationPluginsResponse, error)
	// ListPluginVersions lists all the versions available for the specified plugin.
	ListPluginVersions(context.Context, *ListPluginVersionsRequest) (*ListPluginVersionsResponse, error)
	// CreatePlugin creates a new plugin.
	CreatePlugin(context.Context, *CreatePluginRequest) (*CreatePluginResponse, error)
	// GetPlugin returns the plugin, if found.
	GetPlugin(context.Context, *GetPluginRequest) (*GetPluginResponse, error)
	// DeletePlugin deletes the plugin, if it exists. Note that deleting
	// a plugin may cause breaking changes for templates using that plugin,
	// and should be done with extreme care.
	DeletePlugin(context.Context, *DeletePluginRequest) (*DeletePluginResponse, error)
	// GetTemplate returns the template, if found.
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)
	// ListTemplates returns all the templates available to the user. This includes
	// public templates, those owned by organizations the user is part of,
	// and any created directly by the user.
	ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error)
	// ListUserPlugins lists all templates belonging to a user.
	ListUserTemplates(context.Context, *ListUserTemplatesRequest) (*ListUserTemplatesResponse, error)
	// ListOrganizationTemplates lists all templates for an organization.
	ListOrganizationTemplates(context.Context, *ListOrganizationTemplatesRequest) (*ListOrganizationTemplatesResponse, error)
	// GetTemplateVersion returns the template version, if found.
	GetTemplateVersion(context.Context, *GetTemplateVersionRequest) (*GetTemplateVersionResponse, error)
	// ListTemplateVersions lists all the template versions available for the specified template.
	ListTemplateVersions(context.Context, *ListTemplateVersionsRequest) (*ListTemplateVersionsResponse, error)
	// CreateTemplate creates a new template.
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	// DeleteTemplate deletes the template, if it exists.
	DeleteTemplate(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error)
	// CreateTemplateVersion creates a new template version.
	CreateTemplateVersion(context.Context, *CreateTemplateVersionRequest) (*CreateTemplateVersionResponse, error)
}

// UnimplementedPluginServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPluginServiceServer struct {
}

func (UnimplementedPluginServiceServer) ListPlugins(context.Context, *ListPluginsRequest) (*ListPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlugins not implemented")
}
func (UnimplementedPluginServiceServer) ListUserPlugins(context.Context, *ListUserPluginsRequest) (*ListUserPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserPlugins not implemented")
}
func (UnimplementedPluginServiceServer) ListOrganizationPlugins(context.Context, *ListOrganizationPluginsRequest) (*ListOrganizationPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizationPlugins not implemented")
}
func (UnimplementedPluginServiceServer) ListPluginVersions(context.Context, *ListPluginVersionsRequest) (*ListPluginVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPluginVersions not implemented")
}
func (UnimplementedPluginServiceServer) CreatePlugin(context.Context, *CreatePluginRequest) (*CreatePluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlugin not implemented")
}
func (UnimplementedPluginServiceServer) GetPlugin(context.Context, *GetPluginRequest) (*GetPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugin not implemented")
}
func (UnimplementedPluginServiceServer) DeletePlugin(context.Context, *DeletePluginRequest) (*DeletePluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlugin not implemented")
}
func (UnimplementedPluginServiceServer) GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedPluginServiceServer) ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplates not implemented")
}
func (UnimplementedPluginServiceServer) ListUserTemplates(context.Context, *ListUserTemplatesRequest) (*ListUserTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTemplates not implemented")
}
func (UnimplementedPluginServiceServer) ListOrganizationTemplates(context.Context, *ListOrganizationTemplatesRequest) (*ListOrganizationTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizationTemplates not implemented")
}
func (UnimplementedPluginServiceServer) GetTemplateVersion(context.Context, *GetTemplateVersionRequest) (*GetTemplateVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplateVersion not implemented")
}
func (UnimplementedPluginServiceServer) ListTemplateVersions(context.Context, *ListTemplateVersionsRequest) (*ListTemplateVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplateVersions not implemented")
}
func (UnimplementedPluginServiceServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedPluginServiceServer) DeleteTemplate(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (UnimplementedPluginServiceServer) CreateTemplateVersion(context.Context, *CreateTemplateVersionRequest) (*CreateTemplateVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplateVersion not implemented")
}

// UnsafePluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServiceServer will
// result in compilation errors.
type UnsafePluginServiceServer interface {
	mustEmbedUnimplementedPluginServiceServer()
}

func RegisterPluginServiceServer(s grpc.ServiceRegistrar, srv PluginServiceServer) {
	s.RegisterService(&PluginService_ServiceDesc, srv)
}

func _PluginService_ListPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListPlugins(ctx, req.(*ListPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListUserPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListUserPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListUserPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListUserPlugins(ctx, req.(*ListUserPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListOrganizationPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListOrganizationPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListOrganizationPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListOrganizationPlugins(ctx, req.(*ListOrganizationPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListPluginVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPluginVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListPluginVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListPluginVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListPluginVersions(ctx, req.(*ListPluginVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_CreatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).CreatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/CreatePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).CreatePlugin(ctx, req.(*CreatePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_GetPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).GetPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/GetPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).GetPlugin(ctx, req.(*GetPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_DeletePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).DeletePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/DeletePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).DeletePlugin(ctx, req.(*DeletePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListTemplates(ctx, req.(*ListTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListUserTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListUserTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListUserTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListUserTemplates(ctx, req.(*ListUserTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListOrganizationTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListOrganizationTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListOrganizationTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListOrganizationTemplates(ctx, req.(*ListOrganizationTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_GetTemplateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).GetTemplateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/GetTemplateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).GetTemplateVersion(ctx, req.(*GetTemplateVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ListTemplateVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplateVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListTemplateVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/ListTemplateVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListTemplateVersions(ctx, req.(*ListTemplateVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).DeleteTemplate(ctx, req.(*DeleteTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_CreateTemplateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).CreateTemplateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.PluginService/CreateTemplateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).CreateTemplateVersion(ctx, req.(*CreateTemplateVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginService_ServiceDesc is the grpc.ServiceDesc for PluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.PluginService",
	HandlerType: (*PluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPlugins",
			Handler:    _PluginService_ListPlugins_Handler,
		},
		{
			MethodName: "ListUserPlugins",
			Handler:    _PluginService_ListUserPlugins_Handler,
		},
		{
			MethodName: "ListOrganizationPlugins",
			Handler:    _PluginService_ListOrganizationPlugins_Handler,
		},
		{
			MethodName: "ListPluginVersions",
			Handler:    _PluginService_ListPluginVersions_Handler,
		},
		{
			MethodName: "CreatePlugin",
			Handler:    _PluginService_CreatePlugin_Handler,
		},
		{
			MethodName: "GetPlugin",
			Handler:    _PluginService_GetPlugin_Handler,
		},
		{
			MethodName: "DeletePlugin",
			Handler:    _PluginService_DeletePlugin_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _PluginService_GetTemplate_Handler,
		},
		{
			MethodName: "ListTemplates",
			Handler:    _PluginService_ListTemplates_Handler,
		},
		{
			MethodName: "ListUserTemplates",
			Handler:    _PluginService_ListUserTemplates_Handler,
		},
		{
			MethodName: "ListOrganizationTemplates",
			Handler:    _PluginService_ListOrganizationTemplates_Handler,
		},
		{
			MethodName: "GetTemplateVersion",
			Handler:    _PluginService_GetTemplateVersion_Handler,
		},
		{
			MethodName: "ListTemplateVersions",
			Handler:    _PluginService_ListTemplateVersions_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _PluginService_CreateTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _PluginService_DeleteTemplate_Handler,
		},
		{
			MethodName: "CreateTemplateVersion",
			Handler:    _PluginService_CreateTemplateVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/plugin.proto",
}

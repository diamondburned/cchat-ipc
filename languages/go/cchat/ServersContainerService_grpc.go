//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: cchat

package cchat

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for ServersContainerService service
type ServersContainerServiceClient interface{
  SetServers(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* NoReply, error)  
  UpdateServer(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* NoReply, error)  
}

type serversContainerServiceClient struct {
  cc *grpc.ClientConn
}

func NewServersContainerServiceClient(cc *grpc.ClientConn) ServersContainerServiceClient {
  return &serversContainerServiceClient{cc}
}

func (c *serversContainerServiceClient) SetServers(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* NoReply, error) {
  out := new(NoReply)
  err := grpc.Invoke(ctx, "/cchat.ServersContainerService/SetServers", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *serversContainerServiceClient) UpdateServer(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* NoReply, error) {
  out := new(NoReply)
  err := grpc.Invoke(ctx, "/cchat.ServersContainerService/UpdateServer", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for ServersContainerService service
type ServersContainerServiceServer interface {
  SetServers(context.Context, *SetServersParameters) (*flatbuffers.Builder, error)  
  UpdateServer(context.Context, *UpdateServerParameters) (*flatbuffers.Builder, error)  
}

func RegisterServersContainerServiceServer(s *grpc.Server, srv ServersContainerServiceServer) {
  s.RegisterService(&_ServersContainerService_serviceDesc, srv)
}

func _ServersContainerService_SetServers_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(SetServersParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(ServersContainerServiceServer).SetServers(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.ServersContainerService/SetServers",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(ServersContainerServiceServer).SetServers(ctx, req.(* SetServersParameters))
  }
  return interceptor(ctx, in, info, handler)
}


func _ServersContainerService_UpdateServer_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(UpdateServerParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(ServersContainerServiceServer).UpdateServer(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.ServersContainerService/UpdateServer",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(ServersContainerServiceServer).UpdateServer(ctx, req.(* UpdateServerParameters))
  }
  return interceptor(ctx, in, info, handler)
}


var _ServersContainerService_serviceDesc = grpc.ServiceDesc{
  ServiceName: "cchat.ServersContainerService",
  HandlerType: (*ServersContainerServiceServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "SetServers",
      Handler: _ServersContainerService_SetServers_Handler, 
    },
    {
      MethodName: "UpdateServer",
      Handler: _ServersContainerService_UpdateServer_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}


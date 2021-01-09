//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: cchat

package cchat

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for ConfiguratorService service
type ConfiguratorServiceClient interface{
  Configuration(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* ConfigurationReturns, error)  
  SetConfiguration(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* SetConfigurationReturns, error)  
}

type configuratorServiceClient struct {
  cc *grpc.ClientConn
}

func NewConfiguratorServiceClient(cc *grpc.ClientConn) ConfiguratorServiceClient {
  return &configuratorServiceClient{cc}
}

func (c *configuratorServiceClient) Configuration(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* ConfigurationReturns, error) {
  out := new(ConfigurationReturns)
  err := grpc.Invoke(ctx, "/cchat.ConfiguratorService/Configuration", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *configuratorServiceClient) SetConfiguration(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* SetConfigurationReturns, error) {
  out := new(SetConfigurationReturns)
  err := grpc.Invoke(ctx, "/cchat.ConfiguratorService/SetConfiguration", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for ConfiguratorService service
type ConfiguratorServiceServer interface {
  Configuration(context.Context, *ConfigurationParameters) (*flatbuffers.Builder, error)  
  SetConfiguration(context.Context, *SetConfigurationParameters) (*flatbuffers.Builder, error)  
}

func RegisterConfiguratorServiceServer(s *grpc.Server, srv ConfiguratorServiceServer) {
  s.RegisterService(&_ConfiguratorService_serviceDesc, srv)
}

func _ConfiguratorService_Configuration_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(ConfigurationParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(ConfiguratorServiceServer).Configuration(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.ConfiguratorService/Configuration",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(ConfiguratorServiceServer).Configuration(ctx, req.(* ConfigurationParameters))
  }
  return interceptor(ctx, in, info, handler)
}


func _ConfiguratorService_SetConfiguration_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(SetConfigurationParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(ConfiguratorServiceServer).SetConfiguration(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.ConfiguratorService/SetConfiguration",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(ConfiguratorServiceServer).SetConfiguration(ctx, req.(* SetConfigurationParameters))
  }
  return interceptor(ctx, in, info, handler)
}


var _ConfiguratorService_serviceDesc = grpc.ServiceDesc{
  ServiceName: "cchat.ConfiguratorService",
  HandlerType: (*ConfiguratorServiceServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Configuration",
      Handler: _ConfiguratorService_Configuration_Handler, 
    },
    {
      MethodName: "SetConfiguration",
      Handler: _ConfiguratorService_SetConfiguration_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}

//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: cchat

package call

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for Stopper service
type StopperClient interface{
  Stop(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* Error, error)  
}

type stopperClient struct {
  cc *grpc.ClientConn
}

func NewStopperClient(cc *grpc.ClientConn) StopperClient {
  return &stopperClient{cc}
}

func (c *stopperClient) Stop(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* Error, error) {
  out := new(Error)
  err := grpc.Invoke(ctx, "/cchat.call.Stopper/Stop", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for Stopper service
type StopperServer interface {
  Stop(context.Context, *StopHandle) (*flatbuffers.Builder, error)  
}

func RegisterStopperServer(s *grpc.Server, srv StopperServer) {
  s.RegisterService(&_Stopper_serviceDesc, srv)
}

func _Stopper_Stop_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(StopHandle)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(StopperServer).Stop(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.call.Stopper/Stop",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(StopperServer).Stop(ctx, req.(* StopHandle))
  }
  return interceptor(ctx, in, info, handler)
}


var _Stopper_serviceDesc = grpc.ServiceDesc{
  ServiceName: "cchat.call.Stopper",
  HandlerType: (*StopperServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Stop",
      Handler: _Stopper_Stop_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}


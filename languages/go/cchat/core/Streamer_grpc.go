//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: core

package core

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for Streamer service
type StreamerClient interface{
  Begin(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* Error, error)  
  Stop(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* Error, error)  
}

type streamerClient struct {
  cc *grpc.ClientConn
}

func NewStreamerClient(cc *grpc.ClientConn) StreamerClient {
  return &streamerClient{cc}
}

func (c *streamerClient) Begin(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* Error, error) {
  out := new(Error)
  err := grpc.Invoke(ctx, "/cchat.core.Streamer/Begin", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *streamerClient) Stop(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* Error, error) {
  out := new(Error)
  err := grpc.Invoke(ctx, "/cchat.core.Streamer/Stop", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for Streamer service
type StreamerServer interface {
  Begin(context.Context, *StreamTicket) (*flatbuffers.Builder, error)  
  Stop(context.Context, *StreamTicket) (*flatbuffers.Builder, error)  
}

func RegisterStreamerServer(s *grpc.Server, srv StreamerServer) {
  s.RegisterService(&_Streamer_serviceDesc, srv)
}

func _Streamer_Begin_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(StreamTicket)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(StreamerServer).Begin(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.core.Streamer/Begin",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(StreamerServer).Begin(ctx, req.(* StreamTicket))
  }
  return interceptor(ctx, in, info, handler)
}


func _Streamer_Stop_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(StreamTicket)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(StreamerServer).Stop(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.core.Streamer/Stop",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(StreamerServer).Stop(ctx, req.(* StreamTicket))
  }
  return interceptor(ctx, in, info, handler)
}


var _Streamer_serviceDesc = grpc.ServiceDesc{
  ServiceName: "cchat.core.Streamer",
  HandlerType: (*StreamerServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Begin",
      Handler: _Streamer_Begin_Handler, 
    },
    {
      MethodName: "Stop",
      Handler: _Streamer_Stop_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}

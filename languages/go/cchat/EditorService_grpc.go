//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: cchat

package cchat

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for EditorService service
type EditorServiceClient interface{
  IsEditable(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* IsEditableReturns, error)  
  RawContent(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* RawContentReturns, error)  
  Edit(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* EditReturns, error)  
}

type editorServiceClient struct {
  cc *grpc.ClientConn
}

func NewEditorServiceClient(cc *grpc.ClientConn) EditorServiceClient {
  return &editorServiceClient{cc}
}

func (c *editorServiceClient) IsEditable(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* IsEditableReturns, error) {
  out := new(IsEditableReturns)
  err := grpc.Invoke(ctx, "/cchat.EditorService/IsEditable", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *editorServiceClient) RawContent(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* RawContentReturns, error) {
  out := new(RawContentReturns)
  err := grpc.Invoke(ctx, "/cchat.EditorService/RawContent", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *editorServiceClient) Edit(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* EditReturns, error) {
  out := new(EditReturns)
  err := grpc.Invoke(ctx, "/cchat.EditorService/Edit", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for EditorService service
type EditorServiceServer interface {
  IsEditable(context.Context, *IsEditableParameters) (*flatbuffers.Builder, error)  
  RawContent(context.Context, *RawContentParameters) (*flatbuffers.Builder, error)  
  Edit(context.Context, *EditParameters) (*flatbuffers.Builder, error)  
}

func RegisterEditorServiceServer(s *grpc.Server, srv EditorServiceServer) {
  s.RegisterService(&_EditorService_serviceDesc, srv)
}

func _EditorService_IsEditable_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(IsEditableParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(EditorServiceServer).IsEditable(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.EditorService/IsEditable",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(EditorServiceServer).IsEditable(ctx, req.(* IsEditableParameters))
  }
  return interceptor(ctx, in, info, handler)
}


func _EditorService_RawContent_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(RawContentParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(EditorServiceServer).RawContent(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.EditorService/RawContent",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(EditorServiceServer).RawContent(ctx, req.(* RawContentParameters))
  }
  return interceptor(ctx, in, info, handler)
}


func _EditorService_Edit_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(EditParameters)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(EditorServiceServer).Edit(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/cchat.EditorService/Edit",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(EditorServiceServer).Edit(ctx, req.(* EditParameters))
  }
  return interceptor(ctx, in, info, handler)
}


var _EditorService_serviceDesc = grpc.ServiceDesc{
  ServiceName: "cchat.EditorService",
  HandlerType: (*EditorServiceServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "IsEditable",
      Handler: _EditorService_IsEditable_Handler, 
    },
    {
      MethodName: "RawContent",
      Handler: _EditorService_RawContent_Handler, 
    },
    {
      MethodName: "Edit",
      Handler: _EditorService_Edit_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}

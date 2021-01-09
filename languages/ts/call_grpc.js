// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as Stopper_fbs from './call_generated';

import * as grpc from 'grpc';

interface IStopperService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  Stop: IStopperService_IStop;
}
interface IStopperService_IStop extends grpc.MethodDefinition<Stopper_fbs.StopHandle, Stopper_fbs.Error> {
  path: string; // /cchat.call.Stopper/Stop
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<Stopper_fbs.StopHandle>;
  requestDeserialize: grpc.deserialize<Stopper_fbs.StopHandle>;
  responseSerialize: grpc.serialize<Stopper_fbs.Error>;
  responseDeserialize: grpc.deserialize<Stopper_fbs.Error>;
}


export const StopperService: IStopperService;

export interface IStopperServer {
  Stop: grpc.handleUnaryCall<Stopper_fbs.StopHandle, Stopper_fbs.Error>;
}

export interface IStopperClient {
  Stop(request: Stopper_fbs.StopHandle, callback: (error: grpc.ServiceError | null, response: Stopper_fbs.Error) => void): grpc.ClientUnaryCall;
  Stop(request: Stopper_fbs.StopHandle, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Stopper_fbs.Error) => void): grpc.ClientUnaryCall;
  Stop(request: Stopper_fbs.StopHandle, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Stopper_fbs.Error) => void): grpc.ClientUnaryCall;
}

export class StopperClient extends grpc.Client implements IStopperClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);  public Stop(request: Stopper_fbs.StopHandle, callback: (error: grpc.ServiceError | null, response: Stopper_fbs.Error) => void): grpc.ClientUnaryCall;
  public Stop(request: Stopper_fbs.StopHandle, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Stopper_fbs.Error) => void): grpc.ClientUnaryCall;
  public Stop(request: Stopper_fbs.StopHandle, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Stopper_fbs.Error) => void): grpc.ClientUnaryCall;
}

// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as Canceller_fbs from './call_generated';

var grpc = require('grpc');

function serialize_Error(buffer_args) {
  if (!(buffer_args instanceof Canceller_fbs.Error)) {
    throw new Error('Expected argument of type Canceller_fbs.Error');
  }
  return buffer_args.serialize();
}

function deserialize_Error(buffer) {
  return Canceller_fbs.Error.getRootAsError(new flatbuffers.ByteBuffer(buffer))
}


function serialize_Context(buffer_args) {
  if (!(buffer_args instanceof Canceller_fbs.Context)) {
    throw new Error('Expected argument of type Canceller_fbs.Context');
  }
  return buffer_args.serialize();
}

function deserialize_Context(buffer) {
  return Canceller_fbs.Context.getRootAsContext(new flatbuffers.ByteBuffer(buffer))
}

var CancellerService = exports.CancellerService = {
  Cancel: {
    path: '/cchat.call.Canceller/Cancel',
    requestStream: false,
    responseStream: false,
    requestType: flatbuffers.ByteBuffer,
    responseType: Canceller_fbs.Error,
    requestSerialize: serialize_Context,
    requestDeserialize: deserialize_Context,
    responseSerialize: serialize_Error,
    responseDeserialize: deserialize_Error,
  },
};
exports.CancellerClient = grpc.makeGenericClientConstructor(CancellerService);

// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as Canceller_fbs from './call_generated';

import * as grpc from 'grpc';

interface ICancellerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  Cancel: ICancellerService_ICancel;
}
interface ICancellerService_ICancel extends grpc.MethodDefinition<Canceller_fbs.Context, Canceller_fbs.Error> {
  path: string; // /cchat.call.Canceller/Cancel
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<Canceller_fbs.Context>;
  requestDeserialize: grpc.deserialize<Canceller_fbs.Context>;
  responseSerialize: grpc.serialize<Canceller_fbs.Error>;
  responseDeserialize: grpc.deserialize<Canceller_fbs.Error>;
}


export const CancellerService: ICancellerService;

export interface ICancellerServer {
  Cancel: grpc.handleUnaryCall<Canceller_fbs.Context, Canceller_fbs.Error>;
}

export interface ICancellerClient {
  Cancel(request: Canceller_fbs.Context, callback: (error: grpc.ServiceError | null, response: Canceller_fbs.Error) => void): grpc.ClientUnaryCall;
  Cancel(request: Canceller_fbs.Context, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Canceller_fbs.Error) => void): grpc.ClientUnaryCall;
  Cancel(request: Canceller_fbs.Context, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Canceller_fbs.Error) => void): grpc.ClientUnaryCall;
}

export class CancellerClient extends grpc.Client implements ICancellerClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);  public Cancel(request: Canceller_fbs.Context, callback: (error: grpc.ServiceError | null, response: Canceller_fbs.Error) => void): grpc.ClientUnaryCall;
  public Cancel(request: Canceller_fbs.Context, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Canceller_fbs.Error) => void): grpc.ClientUnaryCall;
  public Cancel(request: Canceller_fbs.Context, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Canceller_fbs.Error) => void): grpc.ClientUnaryCall;
}


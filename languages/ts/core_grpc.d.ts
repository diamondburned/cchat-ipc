// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as Streamer_fbs from './core_generated';

import * as grpc from 'grpc';

interface IStreamerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  Begin: IStreamerService_IBegin;
  Stop: IStreamerService_IStop;
}
interface IStreamerService_IBegin extends grpc.MethodDefinition<Streamer_fbs.StreamTicket, Streamer_fbs.Error> {
  path: string; // /cchat.core.Streamer/Begin
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<Streamer_fbs.StreamTicket>;
  requestDeserialize: grpc.deserialize<Streamer_fbs.StreamTicket>;
  responseSerialize: grpc.serialize<Streamer_fbs.Error>;
  responseDeserialize: grpc.deserialize<Streamer_fbs.Error>;
}

interface IStreamerService_IStop extends grpc.MethodDefinition<Streamer_fbs.StreamTicket, Streamer_fbs.Error> {
  path: string; // /cchat.core.Streamer/Stop
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<Streamer_fbs.StreamTicket>;
  requestDeserialize: grpc.deserialize<Streamer_fbs.StreamTicket>;
  responseSerialize: grpc.serialize<Streamer_fbs.Error>;
  responseDeserialize: grpc.deserialize<Streamer_fbs.Error>;
}


export const StreamerService: IStreamerService;

export interface IStreamerServer {
  Begin: grpc.handleUnaryCall<Streamer_fbs.StreamTicket, Streamer_fbs.Error>;
  Stop: grpc.handleUnaryCall<Streamer_fbs.StreamTicket, Streamer_fbs.Error>;
}

export interface IStreamerClient {
  Begin(request: Streamer_fbs.StreamTicket, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  Begin(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  Begin(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  Stop(request: Streamer_fbs.StreamTicket, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  Stop(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  Stop(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
}

export class StreamerClient extends grpc.Client implements IStreamerClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);  public Begin(request: Streamer_fbs.StreamTicket, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  public Begin(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  public Begin(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  public Stop(request: Streamer_fbs.StreamTicket, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  public Stop(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
  public Stop(request: Streamer_fbs.StreamTicket, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: Streamer_fbs.Error) => void): grpc.ClientUnaryCall;
}


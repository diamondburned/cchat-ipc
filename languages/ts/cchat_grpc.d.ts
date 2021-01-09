// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as MemberDynamicSectionService_fbs from './cchat_generated';

import * as grpc from 'grpc';

interface IMemberDynamicSectionServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  LoadMore: IMemberDynamicSectionServiceService_ILoadMore;
  LoadLess: IMemberDynamicSectionServiceService_ILoadLess;
}
interface IMemberDynamicSectionServiceService_ILoadMore extends grpc.MethodDefinition<MemberDynamicSectionService_fbs.LoadMoreParameters, MemberDynamicSectionService_fbs.LoadMoreReturns> {
  path: string; // /cchat.MemberDynamicSectionService/LoadMore
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<MemberDynamicSectionService_fbs.LoadMoreParameters>;
  requestDeserialize: grpc.deserialize<MemberDynamicSectionService_fbs.LoadMoreParameters>;
  responseSerialize: grpc.serialize<MemberDynamicSectionService_fbs.LoadMoreReturns>;
  responseDeserialize: grpc.deserialize<MemberDynamicSectionService_fbs.LoadMoreReturns>;
}

interface IMemberDynamicSectionServiceService_ILoadLess extends grpc.MethodDefinition<MemberDynamicSectionService_fbs.LoadLessParameters, MemberDynamicSectionService_fbs.LoadLessReturns> {
  path: string; // /cchat.MemberDynamicSectionService/LoadLess
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<MemberDynamicSectionService_fbs.LoadLessParameters>;
  requestDeserialize: grpc.deserialize<MemberDynamicSectionService_fbs.LoadLessParameters>;
  responseSerialize: grpc.serialize<MemberDynamicSectionService_fbs.LoadLessReturns>;
  responseDeserialize: grpc.deserialize<MemberDynamicSectionService_fbs.LoadLessReturns>;
}


export const MemberDynamicSectionServiceService: IMemberDynamicSectionServiceService;

export interface IMemberDynamicSectionServiceServer {
  LoadMore: grpc.handleUnaryCall<MemberDynamicSectionService_fbs.LoadMoreParameters, MemberDynamicSectionService_fbs.LoadMoreReturns>;
  LoadLess: grpc.handleUnaryCall<MemberDynamicSectionService_fbs.LoadLessParameters, MemberDynamicSectionService_fbs.LoadLessReturns>;
}

export interface IMemberDynamicSectionServiceClient {
  LoadMore(request: MemberDynamicSectionService_fbs.LoadMoreParameters, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadMoreReturns) => void): grpc.ClientUnaryCall;
  LoadMore(request: MemberDynamicSectionService_fbs.LoadMoreParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadMoreReturns) => void): grpc.ClientUnaryCall;
  LoadMore(request: MemberDynamicSectionService_fbs.LoadMoreParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadMoreReturns) => void): grpc.ClientUnaryCall;
  LoadLess(request: MemberDynamicSectionService_fbs.LoadLessParameters, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadLessReturns) => void): grpc.ClientUnaryCall;
  LoadLess(request: MemberDynamicSectionService_fbs.LoadLessParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadLessReturns) => void): grpc.ClientUnaryCall;
  LoadLess(request: MemberDynamicSectionService_fbs.LoadLessParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadLessReturns) => void): grpc.ClientUnaryCall;
}

export class MemberDynamicSectionServiceClient extends grpc.Client implements IMemberDynamicSectionServiceClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);  public LoadMore(request: MemberDynamicSectionService_fbs.LoadMoreParameters, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadMoreReturns) => void): grpc.ClientUnaryCall;
  public LoadMore(request: MemberDynamicSectionService_fbs.LoadMoreParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadMoreReturns) => void): grpc.ClientUnaryCall;
  public LoadMore(request: MemberDynamicSectionService_fbs.LoadMoreParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadMoreReturns) => void): grpc.ClientUnaryCall;
  public LoadLess(request: MemberDynamicSectionService_fbs.LoadLessParameters, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadLessReturns) => void): grpc.ClientUnaryCall;
  public LoadLess(request: MemberDynamicSectionService_fbs.LoadLessParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadLessReturns) => void): grpc.ClientUnaryCall;
  public LoadLess(request: MemberDynamicSectionService_fbs.LoadLessParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberDynamicSectionService_fbs.LoadLessReturns) => void): grpc.ClientUnaryCall;
}


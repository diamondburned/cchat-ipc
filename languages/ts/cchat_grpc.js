// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as MemberListContainerService_fbs from './cchat_generated';

import * as grpc from 'grpc';

interface IMemberListContainerServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  SetSections: IMemberListContainerServiceService_ISetSections;
  SetMember: IMemberListContainerServiceService_ISetMember;
  RemoveMember: IMemberListContainerServiceService_IRemoveMember;
}
interface IMemberListContainerServiceService_ISetSections extends grpc.MethodDefinition<MemberListContainerService_fbs.SetSectionsParameters, MemberListContainerService_fbs.NoReply> {
  path: string; // /cchat.MemberListContainerService/SetSections
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<MemberListContainerService_fbs.SetSectionsParameters>;
  requestDeserialize: grpc.deserialize<MemberListContainerService_fbs.SetSectionsParameters>;
  responseSerialize: grpc.serialize<MemberListContainerService_fbs.NoReply>;
  responseDeserialize: grpc.deserialize<MemberListContainerService_fbs.NoReply>;
}

interface IMemberListContainerServiceService_ISetMember extends grpc.MethodDefinition<MemberListContainerService_fbs.SetMemberParameters, MemberListContainerService_fbs.NoReply> {
  path: string; // /cchat.MemberListContainerService/SetMember
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<MemberListContainerService_fbs.SetMemberParameters>;
  requestDeserialize: grpc.deserialize<MemberListContainerService_fbs.SetMemberParameters>;
  responseSerialize: grpc.serialize<MemberListContainerService_fbs.NoReply>;
  responseDeserialize: grpc.deserialize<MemberListContainerService_fbs.NoReply>;
}

interface IMemberListContainerServiceService_IRemoveMember extends grpc.MethodDefinition<MemberListContainerService_fbs.RemoveMemberParameters, MemberListContainerService_fbs.NoReply> {
  path: string; // /cchat.MemberListContainerService/RemoveMember
  requestStream: boolean; // false
  responseStream: boolean; // false
  requestSerialize: grpc.serialize<MemberListContainerService_fbs.RemoveMemberParameters>;
  requestDeserialize: grpc.deserialize<MemberListContainerService_fbs.RemoveMemberParameters>;
  responseSerialize: grpc.serialize<MemberListContainerService_fbs.NoReply>;
  responseDeserialize: grpc.deserialize<MemberListContainerService_fbs.NoReply>;
}


export const MemberListContainerServiceService: IMemberListContainerServiceService;

export interface IMemberListContainerServiceServer {
  SetSections: grpc.handleUnaryCall<MemberListContainerService_fbs.SetSectionsParameters, MemberListContainerService_fbs.NoReply>;
  SetMember: grpc.handleUnaryCall<MemberListContainerService_fbs.SetMemberParameters, MemberListContainerService_fbs.NoReply>;
  RemoveMember: grpc.handleUnaryCall<MemberListContainerService_fbs.RemoveMemberParameters, MemberListContainerService_fbs.NoReply>;
}

export interface IMemberListContainerServiceClient {
  SetSections(request: MemberListContainerService_fbs.SetSectionsParameters, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  SetSections(request: MemberListContainerService_fbs.SetSectionsParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  SetSections(request: MemberListContainerService_fbs.SetSectionsParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  SetMember(request: MemberListContainerService_fbs.SetMemberParameters, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  SetMember(request: MemberListContainerService_fbs.SetMemberParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  SetMember(request: MemberListContainerService_fbs.SetMemberParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  RemoveMember(request: MemberListContainerService_fbs.RemoveMemberParameters, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  RemoveMember(request: MemberListContainerService_fbs.RemoveMemberParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  RemoveMember(request: MemberListContainerService_fbs.RemoveMemberParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
}

export class MemberListContainerServiceClient extends grpc.Client implements IMemberListContainerServiceClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);  public SetSections(request: MemberListContainerService_fbs.SetSectionsParameters, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public SetSections(request: MemberListContainerService_fbs.SetSectionsParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public SetSections(request: MemberListContainerService_fbs.SetSectionsParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public SetMember(request: MemberListContainerService_fbs.SetMemberParameters, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public SetMember(request: MemberListContainerService_fbs.SetMemberParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public SetMember(request: MemberListContainerService_fbs.SetMemberParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public RemoveMember(request: MemberListContainerService_fbs.RemoveMemberParameters, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public RemoveMember(request: MemberListContainerService_fbs.RemoveMemberParameters, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
  public RemoveMember(request: MemberListContainerService_fbs.RemoveMemberParameters, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: MemberListContainerService_fbs.NoReply) => void): grpc.ClientUnaryCall;
}

// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as MemberDynamicSectionService_fbs from './cchat_generated';

var grpc = require('grpc');

function serialize_LoadMoreReturns(buffer_args) {
  if (!(buffer_args instanceof MemberDynamicSectionService_fbs.LoadMoreReturns)) {
    throw new Error('Expected argument of type MemberDynamicSectionService_fbs.LoadMoreReturns');
  }
  return buffer_args.serialize();
}

function deserialize_LoadMoreReturns(buffer) {
  return MemberDynamicSectionService_fbs.LoadMoreReturns.getRootAsLoadMoreReturns(new flatbuffers.ByteBuffer(buffer))
}


function serialize_LoadMoreParameters(buffer_args) {
  if (!(buffer_args instanceof MemberDynamicSectionService_fbs.LoadMoreParameters)) {
    throw new Error('Expected argument of type MemberDynamicSectionService_fbs.LoadMoreParameters');
  }
  return buffer_args.serialize();
}

function deserialize_LoadMoreParameters(buffer) {
  return MemberDynamicSectionService_fbs.LoadMoreParameters.getRootAsLoadMoreParameters(new flatbuffers.ByteBuffer(buffer))
}

function serialize_LoadLessReturns(buffer_args) {
  if (!(buffer_args instanceof MemberDynamicSectionService_fbs.LoadLessReturns)) {
    throw new Error('Expected argument of type MemberDynamicSectionService_fbs.LoadLessReturns');
  }
  return buffer_args.serialize();
}

function deserialize_LoadLessReturns(buffer) {
  return MemberDynamicSectionService_fbs.LoadLessReturns.getRootAsLoadLessReturns(new flatbuffers.ByteBuffer(buffer))
}


function serialize_LoadLessParameters(buffer_args) {
  if (!(buffer_args instanceof MemberDynamicSectionService_fbs.LoadLessParameters)) {
    throw new Error('Expected argument of type MemberDynamicSectionService_fbs.LoadLessParameters');
  }
  return buffer_args.serialize();
}

function deserialize_LoadLessParameters(buffer) {
  return MemberDynamicSectionService_fbs.LoadLessParameters.getRootAsLoadLessParameters(new flatbuffers.ByteBuffer(buffer))
}

var MemberDynamicSectionServiceService = exports.MemberDynamicSectionServiceService = {
  LoadMore: {
    path: '/cchat.MemberDynamicSectionService/LoadMore',
    requestStream: false,
    responseStream: false,
    requestType: flatbuffers.ByteBuffer,
    responseType: MemberDynamicSectionService_fbs.LoadMoreReturns,
    requestSerialize: serialize_LoadMoreParameters,
    requestDeserialize: deserialize_LoadMoreParameters,
    responseSerialize: serialize_LoadMoreReturns,
    responseDeserialize: deserialize_LoadMoreReturns,
  },
  LoadLess: {
    path: '/cchat.MemberDynamicSectionService/LoadLess',
    requestStream: false,
    responseStream: false,
    requestType: flatbuffers.ByteBuffer,
    responseType: MemberDynamicSectionService_fbs.LoadLessReturns,
    requestSerialize: serialize_LoadLessParameters,
    requestDeserialize: deserialize_LoadLessParameters,
    responseSerialize: serialize_LoadLessReturns,
    responseDeserialize: deserialize_LoadLessReturns,
  },
};
exports.MemberDynamicSectionServiceClient = grpc.makeGenericClientConstructor(MemberDynamicSectionServiceService);

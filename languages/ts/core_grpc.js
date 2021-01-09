// Generated GRPC code for FlatBuffers TS *** DO NOT EDIT ***
import { flatbuffers } from 'flatbuffers';
import *  as Streamer_fbs from './core_generated';

var grpc = require('grpc');

function serialize_Error(buffer_args) {
  if (!(buffer_args instanceof Streamer_fbs.Error)) {
    throw new Error('Expected argument of type Streamer_fbs.Error');
  }
  return buffer_args.serialize();
}

function deserialize_Error(buffer) {
  return Streamer_fbs.Error.getRootAsError(new flatbuffers.ByteBuffer(buffer))
}


function serialize_StreamTicket(buffer_args) {
  if (!(buffer_args instanceof Streamer_fbs.StreamTicket)) {
    throw new Error('Expected argument of type Streamer_fbs.StreamTicket');
  }
  return buffer_args.serialize();
}

function deserialize_StreamTicket(buffer) {
  return Streamer_fbs.StreamTicket.getRootAsStreamTicket(new flatbuffers.ByteBuffer(buffer))
}


var StreamerService = exports.StreamerService = {
  Begin: {
    path: '/cchat.core.Streamer/Begin',
    requestStream: false,
    responseStream: false,
    requestType: flatbuffers.ByteBuffer,
    responseType: Streamer_fbs.Error,
    requestSerialize: serialize_StreamTicket,
    requestDeserialize: deserialize_StreamTicket,
    responseSerialize: serialize_Error,
    responseDeserialize: deserialize_Error,
  },
  Stop: {
    path: '/cchat.core.Streamer/Stop',
    requestStream: false,
    responseStream: false,
    requestType: flatbuffers.ByteBuffer,
    responseType: Streamer_fbs.Error,
    requestSerialize: serialize_StreamTicket,
    requestDeserialize: deserialize_StreamTicket,
    responseSerialize: serialize_Error,
    responseDeserialize: deserialize_Error,
  },
};
exports.StreamerClient = grpc.makeGenericClientConstructor(StreamerService);

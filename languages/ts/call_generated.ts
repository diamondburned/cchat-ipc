// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';
import * as NS16630687411378572714 from "./core_generated";
/**
 * @constructor
 */
export namespace cchat.call{
export class NoReply {
  bb: flatbuffers.ByteBuffer|null = null;

  bb_pos:number = 0;
/**
 * @param number i
 * @param flatbuffers.ByteBuffer bb
 * @returns NoReply
 */
__init(i:number, bb:flatbuffers.ByteBuffer):NoReply {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param NoReply= obj
 * @returns NoReply
 */
static getRootAsNoReply(bb:flatbuffers.ByteBuffer, obj?:NoReply):NoReply {
  return (obj || new NoReply()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param NoReply= obj
 * @returns NoReply
 */
static getSizePrefixedRootAsNoReply(bb:flatbuffers.ByteBuffer, obj?:NoReply):NoReply {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new NoReply()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param flatbuffers.Builder builder
 */
static startNoReply(builder:flatbuffers.Builder) {
  builder.startObject(0);
};

/**
 * @param flatbuffers.Builder builder
 * @returns flatbuffers.Offset
 */
static endNoReply(builder:flatbuffers.Builder):flatbuffers.Offset {
  var offset = builder.endObject();
  return offset;
};

static createNoReply(builder:flatbuffers.Builder):flatbuffers.Offset {
  NoReply.startNoReply(builder);
  return NoReply.endNoReply(builder);
}

serialize():Uint8Array {
  return this.bb!.bytes();
}

static deserialize(buffer: Uint8Array):NoReply {
  return NoReply.getRootAsNoReply(new flatbuffers.ByteBuffer(buffer))
}
}
}
/**
 * @constructor
 */
export namespace cchat.call{
export class ID {
  bb: flatbuffers.ByteBuffer|null = null;

  bb_pos:number = 0;
/**
 * @param number i
 * @param flatbuffers.ByteBuffer bb
 * @returns ID
 */
__init(i:number, bb:flatbuffers.ByteBuffer):ID {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param ID= obj
 * @returns ID
 */
static getRootAsID(bb:flatbuffers.ByteBuffer, obj?:ID):ID {
  return (obj || new ID()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param ID= obj
 * @returns ID
 */
static getSizePrefixedRootAsID(bb:flatbuffers.ByteBuffer, obj?:ID):ID {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new ID()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param cchat.core.Object= obj
 * @returns cchat.core.Object|null
 */
object(obj?:NS16630687411378572714.cchat.core.Object):NS16630687411378572714.cchat.core.Object|null {
  var offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? (obj || new NS16630687411378572714.cchat.core.Object()).__init(this.bb_pos + offset, this.bb!) : null;
};

/**
 * @returns flatbuffers.Long
 */
callId():flatbuffers.Long {
  var offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.readUint64(this.bb_pos + offset) : this.bb!.createLong(0, 0);
};

/**
 * @param flatbuffers.Builder builder
 */
static startID(builder:flatbuffers.Builder) {
  builder.startObject(2);
};

/**
 * @param flatbuffers.Builder builder
 * @param flatbuffers.Offset objectOffset
 */
static addObject(builder:flatbuffers.Builder, objectOffset:flatbuffers.Offset) {
  builder.addFieldStruct(0, objectOffset, 0);
};

/**
 * @param flatbuffers.Builder builder
 * @param flatbuffers.Long callId
 */
static addCallId(builder:flatbuffers.Builder, callId:flatbuffers.Long) {
  builder.addFieldInt64(1, callId, builder.createLong(0, 0));
};

/**
 * @param flatbuffers.Builder builder
 * @returns flatbuffers.Offset
 */
static endID(builder:flatbuffers.Builder):flatbuffers.Offset {
  var offset = builder.endObject();
  builder.requiredField(offset, 4); // object
  return offset;
};

static createID(builder:flatbuffers.Builder, objectOffset:flatbuffers.Offset, callId:flatbuffers.Long):flatbuffers.Offset {
  ID.startID(builder);
  ID.addObject(builder, objectOffset);
  ID.addCallId(builder, callId);
  return ID.endID(builder);
}

serialize():Uint8Array {
  return this.bb!.bytes();
}

static deserialize(buffer: Uint8Array):ID {
  return ID.getRootAsID(new flatbuffers.ByteBuffer(buffer))
}
}
}
/**
 * @constructor
 */
export namespace cchat.call{
export class StopHandle {
  bb: flatbuffers.ByteBuffer|null = null;

  bb_pos:number = 0;
/**
 * @param number i
 * @param flatbuffers.ByteBuffer bb
 * @returns StopHandle
 */
__init(i:number, bb:flatbuffers.ByteBuffer):StopHandle {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param StopHandle= obj
 * @returns StopHandle
 */
static getRootAsStopHandle(bb:flatbuffers.ByteBuffer, obj?:StopHandle):StopHandle {
  return (obj || new StopHandle()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param StopHandle= obj
 * @returns StopHandle
 */
static getSizePrefixedRootAsStopHandle(bb:flatbuffers.ByteBuffer, obj?:StopHandle):StopHandle {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new StopHandle()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns flatbuffers.Long
 */
stopId():flatbuffers.Long {
  var offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.readUint64(this.bb_pos + offset) : this.bb!.createLong(0, 0);
};

/**
 * @param flatbuffers.Builder builder
 */
static startStopHandle(builder:flatbuffers.Builder) {
  builder.startObject(1);
};

/**
 * @param flatbuffers.Builder builder
 * @param flatbuffers.Long stopId
 */
static addStopId(builder:flatbuffers.Builder, stopId:flatbuffers.Long) {
  builder.addFieldInt64(0, stopId, builder.createLong(0, 0));
};

/**
 * @param flatbuffers.Builder builder
 * @returns flatbuffers.Offset
 */
static endStopHandle(builder:flatbuffers.Builder):flatbuffers.Offset {
  var offset = builder.endObject();
  return offset;
};

static createStopHandle(builder:flatbuffers.Builder, stopId:flatbuffers.Long):flatbuffers.Offset {
  StopHandle.startStopHandle(builder);
  StopHandle.addStopId(builder, stopId);
  return StopHandle.endStopHandle(builder);
}

serialize():Uint8Array {
  return this.bb!.bytes();
}

static deserialize(buffer: Uint8Array):StopHandle {
  return StopHandle.getRootAsStopHandle(new flatbuffers.ByteBuffer(buffer))
}
}
}
/**
 * @constructor
 */
export namespace cchat.call{
export class Context {
  bb: flatbuffers.ByteBuffer|null = null;

  bb_pos:number = 0;
/**
 * @param number i
 * @param flatbuffers.ByteBuffer bb
 * @returns Context
 */
__init(i:number, bb:flatbuffers.ByteBuffer):Context {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param Context= obj
 * @returns Context
 */
static getRootAsContext(bb:flatbuffers.ByteBuffer, obj?:Context):Context {
  return (obj || new Context()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param flatbuffers.ByteBuffer bb
 * @param Context= obj
 * @returns Context
 */
static getSizePrefixedRootAsContext(bb:flatbuffers.ByteBuffer, obj?:Context):Context {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new Context()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns flatbuffers.Long
 */
ctxId():flatbuffers.Long {
  var offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.readUint64(this.bb_pos + offset) : this.bb!.createLong(0, 0);
};

/**
 * @param flatbuffers.Builder builder
 */
static startContext(builder:flatbuffers.Builder) {
  builder.startObject(1);
};

/**
 * @param flatbuffers.Builder builder
 * @param flatbuffers.Long ctxId
 */
static addCtxId(builder:flatbuffers.Builder, ctxId:flatbuffers.Long) {
  builder.addFieldInt64(0, ctxId, builder.createLong(0, 0));
};

/**
 * @param flatbuffers.Builder builder
 * @returns flatbuffers.Offset
 */
static endContext(builder:flatbuffers.Builder):flatbuffers.Offset {
  var offset = builder.endObject();
  return offset;
};

static createContext(builder:flatbuffers.Builder, ctxId:flatbuffers.Long):flatbuffers.Offset {
  Context.startContext(builder);
  Context.addCtxId(builder, ctxId);
  return Context.endContext(builder);
}

serialize():Uint8Array {
  return this.bb!.bytes();
}

static deserialize(buffer: Uint8Array):Context {
  return Context.getRootAsContext(new flatbuffers.ByteBuffer(buffer))
}
}
}
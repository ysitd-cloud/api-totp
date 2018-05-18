// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var service_pb = require('./service_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var github_com_gogo_protobuf_gogoproto_gogo_pb = require('./github.com/gogo/protobuf/gogoproto/gogo_pb.js');

function serialize_totp_IssueKeyReply(arg) {
  if (!(arg instanceof service_pb.IssueKeyReply)) {
    throw new Error('Expected argument of type totp.IssueKeyReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_IssueKeyReply(buffer_arg) {
  return service_pb.IssueKeyReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_IssueKeyRequest(arg) {
  if (!(arg instanceof service_pb.IssueKeyRequest)) {
    throw new Error('Expected argument of type totp.IssueKeyRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_IssueKeyRequest(buffer_arg) {
  return service_pb.IssueKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_RecoverReply(arg) {
  if (!(arg instanceof service_pb.RecoverReply)) {
    throw new Error('Expected argument of type totp.RecoverReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_RecoverReply(buffer_arg) {
  return service_pb.RecoverReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_RecoverRequest(arg) {
  if (!(arg instanceof service_pb.RecoverRequest)) {
    throw new Error('Expected argument of type totp.RecoverRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_RecoverRequest(buffer_arg) {
  return service_pb.RecoverRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_RemoveKeyReply(arg) {
  if (!(arg instanceof service_pb.RemoveKeyReply)) {
    throw new Error('Expected argument of type totp.RemoveKeyReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_RemoveKeyReply(buffer_arg) {
  return service_pb.RemoveKeyReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_RemoveKeyRequest(arg) {
  if (!(arg instanceof service_pb.RemoveKeyRequest)) {
    throw new Error('Expected argument of type totp.RemoveKeyRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_RemoveKeyRequest(buffer_arg) {
  return service_pb.RemoveKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_ValidateReply(arg) {
  if (!(arg instanceof service_pb.ValidateReply)) {
    throw new Error('Expected argument of type totp.ValidateReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_ValidateReply(buffer_arg) {
  return service_pb.ValidateReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_totp_ValidateRequest(arg) {
  if (!(arg instanceof service_pb.ValidateRequest)) {
    throw new Error('Expected argument of type totp.ValidateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_totp_ValidateRequest(buffer_arg) {
  return service_pb.ValidateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var TotpService = exports.TotpService = {
  issueKey: {
    path: '/totp.Totp/issueKey',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.IssueKeyRequest,
    responseType: service_pb.IssueKeyReply,
    requestSerialize: serialize_totp_IssueKeyRequest,
    requestDeserialize: deserialize_totp_IssueKeyRequest,
    responseSerialize: serialize_totp_IssueKeyReply,
    responseDeserialize: deserialize_totp_IssueKeyReply,
  },
  validate: {
    path: '/totp.Totp/validate',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.ValidateRequest,
    responseType: service_pb.ValidateReply,
    requestSerialize: serialize_totp_ValidateRequest,
    requestDeserialize: deserialize_totp_ValidateRequest,
    responseSerialize: serialize_totp_ValidateReply,
    responseDeserialize: deserialize_totp_ValidateReply,
  },
  recoverKey: {
    path: '/totp.Totp/recoverKey',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.RecoverRequest,
    responseType: service_pb.RecoverReply,
    requestSerialize: serialize_totp_RecoverRequest,
    requestDeserialize: deserialize_totp_RecoverRequest,
    responseSerialize: serialize_totp_RecoverReply,
    responseDeserialize: deserialize_totp_RecoverReply,
  },
  removeKey: {
    path: '/totp.Totp/removeKey',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.RemoveKeyRequest,
    responseType: service_pb.RemoveKeyReply,
    requestSerialize: serialize_totp_RemoveKeyRequest,
    requestDeserialize: deserialize_totp_RemoveKeyRequest,
    responseSerialize: serialize_totp_RemoveKeyReply,
    responseDeserialize: deserialize_totp_RemoveKeyReply,
  },
};

exports.TotpClient = grpc.makeGenericClientConstructor(TotpService);

/**
 * @fileoverview gRPC-Web generated client stub for digimon
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.digimon = require('./schema_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.digimon.DigimonClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.digimon.DigimonPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.digimon.CreateRequest,
 *   !proto.digimon.CreateResponse>}
 */
const methodDescriptor_Digimon_Create = new grpc.web.MethodDescriptor(
  '/digimon.Digimon/Create',
  grpc.web.MethodType.UNARY,
  proto.digimon.CreateRequest,
  proto.digimon.CreateResponse,
  /**
   * @param {!proto.digimon.CreateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.digimon.CreateResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.digimon.CreateRequest,
 *   !proto.digimon.CreateResponse>}
 */
const methodInfo_Digimon_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.digimon.CreateResponse,
  /**
   * @param {!proto.digimon.CreateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.digimon.CreateResponse.deserializeBinary
);


/**
 * @param {!proto.digimon.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.digimon.CreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.digimon.CreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.digimon.DigimonClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/digimon.Digimon/Create',
      request,
      metadata || {},
      methodDescriptor_Digimon_Create,
      callback);
};


/**
 * @param {!proto.digimon.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.digimon.CreateResponse>}
 *     A native promise that resolves to the response
 */
proto.digimon.DigimonPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/digimon.Digimon/Create',
      request,
      metadata || {},
      methodDescriptor_Digimon_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.digimon.QueryRequest,
 *   !proto.digimon.QueryResponse>}
 */
const methodDescriptor_Digimon_QueryStream = new grpc.web.MethodDescriptor(
  '/digimon.Digimon/QueryStream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.digimon.QueryRequest,
  proto.digimon.QueryResponse,
  /**
   * @param {!proto.digimon.QueryRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.digimon.QueryResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.digimon.QueryRequest,
 *   !proto.digimon.QueryResponse>}
 */
const methodInfo_Digimon_QueryStream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.digimon.QueryResponse,
  /**
   * @param {!proto.digimon.QueryRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.digimon.QueryResponse.deserializeBinary
);


/**
 * @param {!proto.digimon.QueryRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.digimon.QueryResponse>}
 *     The XHR Node Readable Stream
 */
proto.digimon.DigimonClient.prototype.queryStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/digimon.Digimon/QueryStream',
      request,
      metadata || {},
      methodDescriptor_Digimon_QueryStream);
};


/**
 * @param {!proto.digimon.QueryRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.digimon.QueryResponse>}
 *     The XHR Node Readable Stream
 */
proto.digimon.DigimonPromiseClient.prototype.queryStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/digimon.Digimon/QueryStream',
      request,
      metadata || {},
      methodDescriptor_Digimon_QueryStream);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.digimon.FosterRequest,
 *   !proto.digimon.FosterResponse>}
 */
const methodDescriptor_Digimon_Foster = new grpc.web.MethodDescriptor(
  '/digimon.Digimon/Foster',
  grpc.web.MethodType.UNARY,
  proto.digimon.FosterRequest,
  proto.digimon.FosterResponse,
  /**
   * @param {!proto.digimon.FosterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.digimon.FosterResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.digimon.FosterRequest,
 *   !proto.digimon.FosterResponse>}
 */
const methodInfo_Digimon_Foster = new grpc.web.AbstractClientBase.MethodInfo(
  proto.digimon.FosterResponse,
  /**
   * @param {!proto.digimon.FosterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.digimon.FosterResponse.deserializeBinary
);


/**
 * @param {!proto.digimon.FosterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.digimon.FosterResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.digimon.FosterResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.digimon.DigimonClient.prototype.foster =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/digimon.Digimon/Foster',
      request,
      metadata || {},
      methodDescriptor_Digimon_Foster,
      callback);
};


/**
 * @param {!proto.digimon.FosterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.digimon.FosterResponse>}
 *     A native promise that resolves to the response
 */
proto.digimon.DigimonPromiseClient.prototype.foster =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/digimon.Digimon/Foster',
      request,
      metadata || {},
      methodDescriptor_Digimon_Foster);
};


module.exports = proto.digimon;


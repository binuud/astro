/**
 * @fileoverview gRPC-Web generated client stub for astro
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as astro_astro_pb from '../astro/astro_pb';


export class AstroClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreate = new grpcWeb.MethodDescriptor(
    '/astro.Astro/Create',
    grpcWeb.MethodType.UNARY,
    astro_astro_pb.UserCreateRequest,
    astro_astro_pb.UserCreateResponse,
    (request: astro_astro_pb.UserCreateRequest) => {
      return request.serializeBinary();
    },
    astro_astro_pb.UserCreateResponse.deserializeBinary
  );

  create(
    request: astro_astro_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_astro_pb.UserCreateResponse>;

  create(
    request: astro_astro_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserCreateResponse) => void): grpcWeb.ClientReadableStream<astro_astro_pb.UserCreateResponse>;

  create(
    request: astro_astro_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserCreateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Astro/Create',
        request,
        metadata || {},
        this.methodDescriptorCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Astro/Create',
    request,
    metadata || {},
    this.methodDescriptorCreate);
  }

  methodDescriptorUpdate = new grpcWeb.MethodDescriptor(
    '/astro.Astro/Update',
    grpcWeb.MethodType.UNARY,
    astro_astro_pb.UserUpdateRequest,
    astro_astro_pb.UserUpdateResponse,
    (request: astro_astro_pb.UserUpdateRequest) => {
      return request.serializeBinary();
    },
    astro_astro_pb.UserUpdateResponse.deserializeBinary
  );

  update(
    request: astro_astro_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_astro_pb.UserUpdateResponse>;

  update(
    request: astro_astro_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserUpdateResponse) => void): grpcWeb.ClientReadableStream<astro_astro_pb.UserUpdateResponse>;

  update(
    request: astro_astro_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserUpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Astro/Update',
        request,
        metadata || {},
        this.methodDescriptorUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Astro/Update',
    request,
    metadata || {},
    this.methodDescriptorUpdate);
  }

  methodDescriptorGet = new grpcWeb.MethodDescriptor(
    '/astro.Astro/Get',
    grpcWeb.MethodType.UNARY,
    astro_astro_pb.UserGetRequest,
    astro_astro_pb.UserGetResponse,
    (request: astro_astro_pb.UserGetRequest) => {
      return request.serializeBinary();
    },
    astro_astro_pb.UserGetResponse.deserializeBinary
  );

  get(
    request: astro_astro_pb.UserGetRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_astro_pb.UserGetResponse>;

  get(
    request: astro_astro_pb.UserGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserGetResponse) => void): grpcWeb.ClientReadableStream<astro_astro_pb.UserGetResponse>;

  get(
    request: astro_astro_pb.UserGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Astro/Get',
        request,
        metadata || {},
        this.methodDescriptorGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Astro/Get',
    request,
    metadata || {},
    this.methodDescriptorGet);
  }

  methodDescriptorList = new grpcWeb.MethodDescriptor(
    '/astro.Astro/List',
    grpcWeb.MethodType.UNARY,
    astro_astro_pb.UserListRequest,
    astro_astro_pb.UserListResponse,
    (request: astro_astro_pb.UserListRequest) => {
      return request.serializeBinary();
    },
    astro_astro_pb.UserListResponse.deserializeBinary
  );

  list(
    request: astro_astro_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_astro_pb.UserListResponse>;

  list(
    request: astro_astro_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserListResponse) => void): grpcWeb.ClientReadableStream<astro_astro_pb.UserListResponse>;

  list(
    request: astro_astro_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_astro_pb.UserListResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Astro/List',
        request,
        metadata || {},
        this.methodDescriptorList,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Astro/List',
    request,
    metadata || {},
    this.methodDescriptorList);
  }

}


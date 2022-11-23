/**
 * @fileoverview gRPC-Web generated client stub for astro
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as astro_dasha_pb from '../astro/dasha_pb';


export class DashaClient {
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

  methodDescriptorGetDashas = new grpcWeb.MethodDescriptor(
    '/astro.Dasha/GetDashas',
    grpcWeb.MethodType.UNARY,
    astro_dasha_pb.DashaGetReqeust,
    astro_dasha_pb.DashaGetResponse,
    (request: astro_dasha_pb.DashaGetReqeust) => {
      return request.serializeBinary();
    },
    astro_dasha_pb.DashaGetResponse.deserializeBinary
  );

  getDashas(
    request: astro_dasha_pb.DashaGetReqeust,
    metadata: grpcWeb.Metadata | null): Promise<astro_dasha_pb.DashaGetResponse>;

  getDashas(
    request: astro_dasha_pb.DashaGetReqeust,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_dasha_pb.DashaGetResponse) => void): grpcWeb.ClientReadableStream<astro_dasha_pb.DashaGetResponse>;

  getDashas(
    request: astro_dasha_pb.DashaGetReqeust,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_dasha_pb.DashaGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Dasha/GetDashas',
        request,
        metadata || {},
        this.methodDescriptorGetDashas,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Dasha/GetDashas',
    request,
    metadata || {},
    this.methodDescriptorGetDashas);
  }

  methodDescriptorUpdate = new grpcWeb.MethodDescriptor(
    '/astro.Dasha/Update',
    grpcWeb.MethodType.UNARY,
    astro_dasha_pb.DashaUpdateReqeust,
    astro_dasha_pb.DashaUpdateResponse,
    (request: astro_dasha_pb.DashaUpdateReqeust) => {
      return request.serializeBinary();
    },
    astro_dasha_pb.DashaUpdateResponse.deserializeBinary
  );

  update(
    request: astro_dasha_pb.DashaUpdateReqeust,
    metadata: grpcWeb.Metadata | null): Promise<astro_dasha_pb.DashaUpdateResponse>;

  update(
    request: astro_dasha_pb.DashaUpdateReqeust,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_dasha_pb.DashaUpdateResponse) => void): grpcWeb.ClientReadableStream<astro_dasha_pb.DashaUpdateResponse>;

  update(
    request: astro_dasha_pb.DashaUpdateReqeust,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_dasha_pb.DashaUpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Dasha/Update',
        request,
        metadata || {},
        this.methodDescriptorUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Dasha/Update',
    request,
    metadata || {},
    this.methodDescriptorUpdate);
  }

}


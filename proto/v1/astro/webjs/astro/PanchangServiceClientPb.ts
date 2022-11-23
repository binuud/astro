/**
 * @fileoverview gRPC-Web generated client stub for astro
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as astro_panchang_pb from '../astro/panchang_pb';


export class PanchangClient {
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

  methodDescriptorGetMonth = new grpcWeb.MethodDescriptor(
    '/astro.Panchang/GetMonth',
    grpcWeb.MethodType.UNARY,
    astro_panchang_pb.PanchangMonthGetRequest,
    astro_panchang_pb.PanchangMonthGetResponse,
    (request: astro_panchang_pb.PanchangMonthGetRequest) => {
      return request.serializeBinary();
    },
    astro_panchang_pb.PanchangMonthGetResponse.deserializeBinary
  );

  getMonth(
    request: astro_panchang_pb.PanchangMonthGetRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_panchang_pb.PanchangMonthGetResponse>;

  getMonth(
    request: astro_panchang_pb.PanchangMonthGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_panchang_pb.PanchangMonthGetResponse) => void): grpcWeb.ClientReadableStream<astro_panchang_pb.PanchangMonthGetResponse>;

  getMonth(
    request: astro_panchang_pb.PanchangMonthGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_panchang_pb.PanchangMonthGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Panchang/GetMonth',
        request,
        metadata || {},
        this.methodDescriptorGetMonth,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Panchang/GetMonth',
    request,
    metadata || {},
    this.methodDescriptorGetMonth);
  }

  methodDescriptorGetDay = new grpcWeb.MethodDescriptor(
    '/astro.Panchang/GetDay',
    grpcWeb.MethodType.UNARY,
    astro_panchang_pb.PanchangDayGetRequest,
    astro_panchang_pb.PanchangDayGetResponse,
    (request: astro_panchang_pb.PanchangDayGetRequest) => {
      return request.serializeBinary();
    },
    astro_panchang_pb.PanchangDayGetResponse.deserializeBinary
  );

  getDay(
    request: astro_panchang_pb.PanchangDayGetRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_panchang_pb.PanchangDayGetResponse>;

  getDay(
    request: astro_panchang_pb.PanchangDayGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_panchang_pb.PanchangDayGetResponse) => void): grpcWeb.ClientReadableStream<astro_panchang_pb.PanchangDayGetResponse>;

  getDay(
    request: astro_panchang_pb.PanchangDayGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_panchang_pb.PanchangDayGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Panchang/GetDay',
        request,
        metadata || {},
        this.methodDescriptorGetDay,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Panchang/GetDay',
    request,
    metadata || {},
    this.methodDescriptorGetDay);
  }

}


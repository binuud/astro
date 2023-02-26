/**
 * @fileoverview gRPC-Web generated client stub for astro
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as astro_cities_pb from '../astro/cities_pb';


export class CitiesClient {
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

  methodDescriptorGet = new grpcWeb.MethodDescriptor(
    '/astro.Cities/Get',
    grpcWeb.MethodType.UNARY,
    astro_cities_pb.CitiesGetRequest,
    astro_cities_pb.CitiesGetResponse,
    (request: astro_cities_pb.CitiesGetRequest) => {
      return request.serializeBinary();
    },
    astro_cities_pb.CitiesGetResponse.deserializeBinary
  );

  get(
    request: astro_cities_pb.CitiesGetRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_cities_pb.CitiesGetResponse>;

  get(
    request: astro_cities_pb.CitiesGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_cities_pb.CitiesGetResponse) => void): grpcWeb.ClientReadableStream<astro_cities_pb.CitiesGetResponse>;

  get(
    request: astro_cities_pb.CitiesGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_cities_pb.CitiesGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Cities/Get',
        request,
        metadata || {},
        this.methodDescriptorGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Cities/Get',
    request,
    metadata || {},
    this.methodDescriptorGet);
  }

  methodDescriptorSearch = new grpcWeb.MethodDescriptor(
    '/astro.Cities/Search',
    grpcWeb.MethodType.UNARY,
    astro_cities_pb.CitiesSearchRequest,
    astro_cities_pb.CitiesSearchResponse,
    (request: astro_cities_pb.CitiesSearchRequest) => {
      return request.serializeBinary();
    },
    astro_cities_pb.CitiesSearchResponse.deserializeBinary
  );

  search(
    request: astro_cities_pb.CitiesSearchRequest,
    metadata: grpcWeb.Metadata | null): Promise<astro_cities_pb.CitiesSearchResponse>;

  search(
    request: astro_cities_pb.CitiesSearchRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: astro_cities_pb.CitiesSearchResponse) => void): grpcWeb.ClientReadableStream<astro_cities_pb.CitiesSearchResponse>;

  search(
    request: astro_cities_pb.CitiesSearchRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: astro_cities_pb.CitiesSearchResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/astro.Cities/Search',
        request,
        metadata || {},
        this.methodDescriptorSearch,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/astro.Cities/Search',
    request,
    metadata || {},
    this.methodDescriptorSearch);
  }

}


/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"
import * as AstroDasha from "./dasha.pb"
import * as AstroMetas from "./metas.pb"

export enum MetadataChartType {
  IgnoreChartType = "IgnoreChartType",
  Self = "Self",
  Global = "Global",
  Client = "Client",
  Test = "Test",
}

export type Metadata = {
  uid?: string
  name?: string
  owner?: string
  tenancy?: string
  parentUid?: string
  dataEncryption?: boolean
  labels?: {[key: string]: string}
  annotations?: {[key: string]: string}
  createdAt?: GoogleProtobufTimestamp.Timestamp
  updatedAt?: GoogleProtobufTimestamp.Timestamp
  chartType?: MetadataChartType
}

export type Pagination = {
  total?: string
  current?: string
  perPage?: string
}

export type PlanetDetails = {
  name?: string
  longitude?: number
  rasiLongitude?: number
  latitude?: number
  declination?: number
  speed?: number
  distance?: number
  star?: string
  paada?: string
  rasi?: string
  bhava?: string
  bhavaH?: string
  exalted?: boolean
  moolTrikona?: boolean
  deblited?: boolean
  ownHouse?: boolean
  friendHouse?: boolean
  enemyHouse?: boolean
  neutralHouse?: boolean
  amsa2?: string
  amsa9?: string
  amsa10?: string
  amsa27?: string
  amsa60?: string
}

export type Planets = {
  ephimeris?: string
  planetsMap?: {[key: string]: PlanetDetails}
}

export type AstroObject = {
  metadata?: Metadata
  birthData?: AstroMetas.BirthData
  planets?: Planets
  dashas?: AstroDasha.DashaItem[]
}

export type UserCreateRequest = {
  metadata?: Metadata
  birthData?: AstroMetas.BirthData
}

export type UserCreateResponse = {
  astro?: AstroObject
}

export type UserUpdateRequest = {
  uuid?: string
  birthData?: AstroMetas.BirthData
}

export type UserUpdateResponse = {
  astro?: AstroObject
}

export type UserGetRequest = {
  uuid?: string
  name?: string
}

export type UserGetResponse = {
  astro?: AstroObject
}

export type UserListRequest = {
  uuid?: string
  pagination?: Pagination
}

export type UserListResponse = {
  astro?: AstroObject[]
  pagination?: Pagination
}

export class Astro {
  static Create(req: UserCreateRequest, initReq?: fm.InitReq): Promise<UserCreateResponse> {
    return fm.fetchReq<UserCreateRequest, UserCreateResponse>(`/astro.Astro/Create`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Update(req: UserUpdateRequest, initReq?: fm.InitReq): Promise<UserUpdateResponse> {
    return fm.fetchReq<UserUpdateRequest, UserUpdateResponse>(`/astro.Astro/Update`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Get(req: UserGetRequest, initReq?: fm.InitReq): Promise<UserGetResponse> {
    return fm.fetchReq<UserGetRequest, UserGetResponse>(`/astro.Astro/Get`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static List(req: UserListRequest, initReq?: fm.InitReq): Promise<UserListResponse> {
    return fm.fetchReq<UserListRequest, UserListResponse>(`/astro.Astro/List`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}
/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"
import * as AstroMetas from "./metas.pb"
export type SunMoonEphimeris = {
  moon?: number
  sun?: number
}

export type Transit = {
  name?: number
  transit?: GoogleProtobufTimestamp.Timestamp
  yoga?: AstroMetas.YogaIdentifier
}

export type Kaalam = {
  startTime?: GoogleProtobufTimestamp.Timestamp
  endTime?: GoogleProtobufTimestamp.Timestamp
}

export type Transits = {
  year?: number
  month?: number
  startTime?: GoogleProtobufTimestamp.Timestamp
  ephimeris?: SunMoonEphimeris[]
  stars?: Transit[]
  yogas?: Transit[]
  karanas?: Transit[]
  tithis?: Transit[]
  paadas?: Transit[]
  rasis?: Transit[]
  rahuKalam?: Kaalam
  guliKalam?: Kaalam
  yamaGandam?: Kaalam
}

export type PanchangDay = {
  year?: number
  month?: number
  day?: number
  zone?: string
  geo?: AstroMetas.GeoLocation
  weekday?: number
  sunRiseSet?: Kaalam
  moonRiseSet?: Kaalam
  stars?: Transit[]
  yogas?: Transit[]
  karanas?: Transit[]
  tithis?: Transit[]
  paadas?: Transit[]
  rasis?: Transit[]
  lagnas?: Transit[]
  rahuKalam?: Kaalam
  guliKalam?: Kaalam
  yamaGandam?: Kaalam
  marriage?: Transit[]
  good?: Transit[]
  bad?: Transit[]
  tags?: AstroMetas.AstroMetasPanchangTags[]
}

export type PanchangMonth = {
  year?: number
  month?: number
  zone?: string
  geo?: AstroMetas.GeoLocation
  days?: PanchangDay[]
}

export type PanchangMonthGetRequest = {
  birthData?: AstroMetas.BirthData
}

export type PanchangMonthGetResponse = {
  month?: PanchangMonth
}

export type PanchangDayGetRequest = {
  birthData?: AstroMetas.BirthData
}

export type PanchangDayGetResponse = {
  day?: PanchangDay
}

export class Panchang {
  static GetMonth(req: PanchangMonthGetRequest, initReq?: fm.InitReq): Promise<PanchangMonthGetResponse> {
    return fm.fetchReq<PanchangMonthGetRequest, PanchangMonthGetResponse>(`/v1/panchang/getMonth?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetDay(req: PanchangDayGetRequest, initReq?: fm.InitReq): Promise<PanchangDayGetResponse> {
    return fm.fetchReq<PanchangDayGetRequest, PanchangDayGetResponse>(`/v1/panchang/getDay?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}
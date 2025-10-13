/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"

export enum DashaMetaTypes {
  Ignore = "Ignore",
  Vimshottari = "Vimshottari",
  Ashtottari = "Ashtottari",
}

export type SubDashaMeta = {
  name?: string
  duration?: number
}

export type DashaMeta = {
  name?: string
  yearInSeconds?: string
  durationYears?: string
  startDasa?: number
  dasaDegrees?: number
  startTime?: GoogleProtobufTimestamp.Timestamp
  subDashaMetas?: SubDashaMeta[]
  type?: DashaMetaTypes
}

export type DashaItem = {
  name?: string
  level?: number
  startTime?: GoogleProtobufTimestamp.Timestamp
  endTime?: GoogleProtobufTimestamp.Timestamp
  dashas?: DashaItem[]
}

export type DashaSystem = {
  meta?: DashaMeta
  dashas?: DashaItem[]
}

export type DashaGetReqeust = {
  uuid?: string
  level?: number
}

export type DashaGetResponse = {
  userUuid?: string
  dashaSystem?: DashaSystem
}

export type DashaUpdateReqeust = {
  uuid?: string
}

export type DashaUpdateResponse = {
  uuid?: string
}

export class Dasha {
  static GetDashas(req: DashaGetReqeust, initReq?: fm.InitReq): Promise<DashaGetResponse> {
    return fm.fetchReq<DashaGetReqeust, DashaGetResponse>(`/astro.Dasha/GetDashas`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Update(req: DashaUpdateReqeust, initReq?: fm.InitReq): Promise<DashaUpdateResponse> {
    return fm.fetchReq<DashaUpdateReqeust, DashaUpdateResponse>(`/astro.Dasha/Update`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}
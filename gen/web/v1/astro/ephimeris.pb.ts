/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as AstroMetas from "./metas.pb"
export type EphimerisRequest = {
  birthData?: AstroMetas.BirthData
}

export type EphimerisResponse = {
  output?: string
}

export class Ephimeris {
  static GetChart(req: EphimerisRequest, initReq?: fm.InitReq): Promise<EphimerisResponse> {
    return fm.fetchReq<EphimerisRequest, EphimerisResponse>(`/astro.Ephimeris/GetChart`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetMoonEphimeris(req: EphimerisRequest, initReq?: fm.InitReq): Promise<EphimerisResponse> {
    return fm.fetchReq<EphimerisRequest, EphimerisResponse>(`/astro.Ephimeris/GetMoonEphimeris`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetPanchangMonth(req: EphimerisRequest, initReq?: fm.InitReq): Promise<EphimerisResponse> {
    return fm.fetchReq<EphimerisRequest, EphimerisResponse>(`/astro.Ephimeris/GetPanchangMonth`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}
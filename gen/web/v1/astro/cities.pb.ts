/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type CityInfo = {
  id?: string
  name?: string
  stateId?: string
  stateCode?: string
  stateName?: string
  countryId?: string
  countryCode?: string
  countryName?: string
  latitude?: number
  longitude?: number
  wikiDataId?: string
  index?: string
}

export type CitiesGetRequest = {
  id?: string
  name?: string
  stateId?: string
  stateCode?: string
  stateName?: string
  countryId?: string
  countryCode?: string
  countryName?: string
  index?: string
}

export type CitiesGetResponse = {
  details?: CityInfo
}

export type CitiesSearchRequest = {
  id?: string
  name?: string
  stateId?: string
  stateCode?: string
  stateName?: string
  countryId?: string
  countryCode?: string
  countryName?: string
}

export type CitiesSearchResponse = {
  cities?: CityInfo[]
}

export class Cities {
  static Get(req: CitiesGetRequest, initReq?: fm.InitReq): Promise<CitiesGetResponse> {
    return fm.fetchReq<CitiesGetRequest, CitiesGetResponse>(`/astro.Cities/Get`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Search(req: CitiesSearchRequest, initReq?: fm.InitReq): Promise<CitiesSearchResponse> {
    return fm.fetchReq<CitiesSearchRequest, CitiesSearchResponse>(`/astro.Cities/Search`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}
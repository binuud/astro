import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as astro_metas_pb from '../astro/metas_pb';


export class SunMoonEphimeris extends jspb.Message {
  getMoon(): number;
  setMoon(value: number): SunMoonEphimeris;

  getSun(): number;
  setSun(value: number): SunMoonEphimeris;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SunMoonEphimeris.AsObject;
  static toObject(includeInstance: boolean, msg: SunMoonEphimeris): SunMoonEphimeris.AsObject;
  static serializeBinaryToWriter(message: SunMoonEphimeris, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SunMoonEphimeris;
  static deserializeBinaryFromReader(message: SunMoonEphimeris, reader: jspb.BinaryReader): SunMoonEphimeris;
}

export namespace SunMoonEphimeris {
  export type AsObject = {
    moon: number,
    sun: number,
  }
}

export class Transit extends jspb.Message {
  getName(): number;
  setName(value: number): Transit;

  getTransit(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTransit(value?: google_protobuf_timestamp_pb.Timestamp): Transit;
  hasTransit(): boolean;
  clearTransit(): Transit;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Transit.AsObject;
  static toObject(includeInstance: boolean, msg: Transit): Transit.AsObject;
  static serializeBinaryToWriter(message: Transit, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Transit;
  static deserializeBinaryFromReader(message: Transit, reader: jspb.BinaryReader): Transit;
}

export namespace Transit {
  export type AsObject = {
    name: number,
    transit?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Kaalam extends jspb.Message {
  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): Kaalam;
  hasStarttime(): boolean;
  clearStarttime(): Kaalam;

  getEndtime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEndtime(value?: google_protobuf_timestamp_pb.Timestamp): Kaalam;
  hasEndtime(): boolean;
  clearEndtime(): Kaalam;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Kaalam.AsObject;
  static toObject(includeInstance: boolean, msg: Kaalam): Kaalam.AsObject;
  static serializeBinaryToWriter(message: Kaalam, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Kaalam;
  static deserializeBinaryFromReader(message: Kaalam, reader: jspb.BinaryReader): Kaalam;
}

export namespace Kaalam {
  export type AsObject = {
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    endtime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Transits extends jspb.Message {
  getYear(): number;
  setYear(value: number): Transits;

  getMonth(): number;
  setMonth(value: number): Transits;

  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): Transits;
  hasStarttime(): boolean;
  clearStarttime(): Transits;

  getEphimerisList(): Array<SunMoonEphimeris>;
  setEphimerisList(value: Array<SunMoonEphimeris>): Transits;
  clearEphimerisList(): Transits;
  addEphimeris(value?: SunMoonEphimeris, index?: number): SunMoonEphimeris;

  getStarsList(): Array<Transit>;
  setStarsList(value: Array<Transit>): Transits;
  clearStarsList(): Transits;
  addStars(value?: Transit, index?: number): Transit;

  getYogasList(): Array<Transit>;
  setYogasList(value: Array<Transit>): Transits;
  clearYogasList(): Transits;
  addYogas(value?: Transit, index?: number): Transit;

  getKaranasList(): Array<Transit>;
  setKaranasList(value: Array<Transit>): Transits;
  clearKaranasList(): Transits;
  addKaranas(value?: Transit, index?: number): Transit;

  getTithisList(): Array<Transit>;
  setTithisList(value: Array<Transit>): Transits;
  clearTithisList(): Transits;
  addTithis(value?: Transit, index?: number): Transit;

  getPaadasList(): Array<Transit>;
  setPaadasList(value: Array<Transit>): Transits;
  clearPaadasList(): Transits;
  addPaadas(value?: Transit, index?: number): Transit;

  getRahukalam(): Kaalam | undefined;
  setRahukalam(value?: Kaalam): Transits;
  hasRahukalam(): boolean;
  clearRahukalam(): Transits;

  getGulikalam(): Kaalam | undefined;
  setGulikalam(value?: Kaalam): Transits;
  hasGulikalam(): boolean;
  clearGulikalam(): Transits;

  getYamagandam(): Kaalam | undefined;
  setYamagandam(value?: Kaalam): Transits;
  hasYamagandam(): boolean;
  clearYamagandam(): Transits;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Transits.AsObject;
  static toObject(includeInstance: boolean, msg: Transits): Transits.AsObject;
  static serializeBinaryToWriter(message: Transits, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Transits;
  static deserializeBinaryFromReader(message: Transits, reader: jspb.BinaryReader): Transits;
}

export namespace Transits {
  export type AsObject = {
    year: number,
    month: number,
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    ephimerisList: Array<SunMoonEphimeris.AsObject>,
    starsList: Array<Transit.AsObject>,
    yogasList: Array<Transit.AsObject>,
    karanasList: Array<Transit.AsObject>,
    tithisList: Array<Transit.AsObject>,
    paadasList: Array<Transit.AsObject>,
    rahukalam?: Kaalam.AsObject,
    gulikalam?: Kaalam.AsObject,
    yamagandam?: Kaalam.AsObject,
  }
}

export class PanchangDay extends jspb.Message {
  getYear(): number;
  setYear(value: number): PanchangDay;

  getMonth(): number;
  setMonth(value: number): PanchangDay;

  getDay(): number;
  setDay(value: number): PanchangDay;

  getZone(): string;
  setZone(value: string): PanchangDay;

  getGeo(): astro_metas_pb.GeoLocation | undefined;
  setGeo(value?: astro_metas_pb.GeoLocation): PanchangDay;
  hasGeo(): boolean;
  clearGeo(): PanchangDay;

  getSunriseset(): Kaalam | undefined;
  setSunriseset(value?: Kaalam): PanchangDay;
  hasSunriseset(): boolean;
  clearSunriseset(): PanchangDay;

  getMoonriseset(): Kaalam | undefined;
  setMoonriseset(value?: Kaalam): PanchangDay;
  hasMoonriseset(): boolean;
  clearMoonriseset(): PanchangDay;

  getStarsList(): Array<Transit>;
  setStarsList(value: Array<Transit>): PanchangDay;
  clearStarsList(): PanchangDay;
  addStars(value?: Transit, index?: number): Transit;

  getYogasList(): Array<Transit>;
  setYogasList(value: Array<Transit>): PanchangDay;
  clearYogasList(): PanchangDay;
  addYogas(value?: Transit, index?: number): Transit;

  getKaranasList(): Array<Transit>;
  setKaranasList(value: Array<Transit>): PanchangDay;
  clearKaranasList(): PanchangDay;
  addKaranas(value?: Transit, index?: number): Transit;

  getTithisList(): Array<Transit>;
  setTithisList(value: Array<Transit>): PanchangDay;
  clearTithisList(): PanchangDay;
  addTithis(value?: Transit, index?: number): Transit;

  getPaadasList(): Array<Transit>;
  setPaadasList(value: Array<Transit>): PanchangDay;
  clearPaadasList(): PanchangDay;
  addPaadas(value?: Transit, index?: number): Transit;

  getLagnasList(): Array<Transit>;
  setLagnasList(value: Array<Transit>): PanchangDay;
  clearLagnasList(): PanchangDay;
  addLagnas(value?: Transit, index?: number): Transit;

  getRahukalam(): Kaalam | undefined;
  setRahukalam(value?: Kaalam): PanchangDay;
  hasRahukalam(): boolean;
  clearRahukalam(): PanchangDay;

  getGulikalam(): Kaalam | undefined;
  setGulikalam(value?: Kaalam): PanchangDay;
  hasGulikalam(): boolean;
  clearGulikalam(): PanchangDay;

  getYamagandam(): Kaalam | undefined;
  setYamagandam(value?: Kaalam): PanchangDay;
  hasYamagandam(): boolean;
  clearYamagandam(): PanchangDay;

  getMarriageList(): Array<Transit>;
  setMarriageList(value: Array<Transit>): PanchangDay;
  clearMarriageList(): PanchangDay;
  addMarriage(value?: Transit, index?: number): Transit;

  getGoodList(): Array<Transit>;
  setGoodList(value: Array<Transit>): PanchangDay;
  clearGoodList(): PanchangDay;
  addGood(value?: Transit, index?: number): Transit;

  getBadList(): Array<Transit>;
  setBadList(value: Array<Transit>): PanchangDay;
  clearBadList(): PanchangDay;
  addBad(value?: Transit, index?: number): Transit;

  getTagsList(): Array<astro_metas_pb.AstroMetas.PanchangTags>;
  setTagsList(value: Array<astro_metas_pb.AstroMetas.PanchangTags>): PanchangDay;
  clearTagsList(): PanchangDay;
  addTags(value: astro_metas_pb.AstroMetas.PanchangTags, index?: number): PanchangDay;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangDay.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangDay): PanchangDay.AsObject;
  static serializeBinaryToWriter(message: PanchangDay, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangDay;
  static deserializeBinaryFromReader(message: PanchangDay, reader: jspb.BinaryReader): PanchangDay;
}

export namespace PanchangDay {
  export type AsObject = {
    year: number,
    month: number,
    day: number,
    zone: string,
    geo?: astro_metas_pb.GeoLocation.AsObject,
    sunriseset?: Kaalam.AsObject,
    moonriseset?: Kaalam.AsObject,
    starsList: Array<Transit.AsObject>,
    yogasList: Array<Transit.AsObject>,
    karanasList: Array<Transit.AsObject>,
    tithisList: Array<Transit.AsObject>,
    paadasList: Array<Transit.AsObject>,
    lagnasList: Array<Transit.AsObject>,
    rahukalam?: Kaalam.AsObject,
    gulikalam?: Kaalam.AsObject,
    yamagandam?: Kaalam.AsObject,
    marriageList: Array<Transit.AsObject>,
    goodList: Array<Transit.AsObject>,
    badList: Array<Transit.AsObject>,
    tagsList: Array<astro_metas_pb.AstroMetas.PanchangTags>,
  }
}

export class PanchangMonth extends jspb.Message {
  getYear(): number;
  setYear(value: number): PanchangMonth;

  getMonth(): number;
  setMonth(value: number): PanchangMonth;

  getZone(): string;
  setZone(value: string): PanchangMonth;

  getGeo(): astro_metas_pb.GeoLocation | undefined;
  setGeo(value?: astro_metas_pb.GeoLocation): PanchangMonth;
  hasGeo(): boolean;
  clearGeo(): PanchangMonth;

  getDaysList(): Array<PanchangDay>;
  setDaysList(value: Array<PanchangDay>): PanchangMonth;
  clearDaysList(): PanchangMonth;
  addDays(value?: PanchangDay, index?: number): PanchangDay;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangMonth.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangMonth): PanchangMonth.AsObject;
  static serializeBinaryToWriter(message: PanchangMonth, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangMonth;
  static deserializeBinaryFromReader(message: PanchangMonth, reader: jspb.BinaryReader): PanchangMonth;
}

export namespace PanchangMonth {
  export type AsObject = {
    year: number,
    month: number,
    zone: string,
    geo?: astro_metas_pb.GeoLocation.AsObject,
    daysList: Array<PanchangDay.AsObject>,
  }
}

export class PanchangMonthGetRequest extends jspb.Message {
  getBirthdata(): astro_metas_pb.BirthData | undefined;
  setBirthdata(value?: astro_metas_pb.BirthData): PanchangMonthGetRequest;
  hasBirthdata(): boolean;
  clearBirthdata(): PanchangMonthGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangMonthGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangMonthGetRequest): PanchangMonthGetRequest.AsObject;
  static serializeBinaryToWriter(message: PanchangMonthGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangMonthGetRequest;
  static deserializeBinaryFromReader(message: PanchangMonthGetRequest, reader: jspb.BinaryReader): PanchangMonthGetRequest;
}

export namespace PanchangMonthGetRequest {
  export type AsObject = {
    birthdata?: astro_metas_pb.BirthData.AsObject,
  }
}

export class PanchangMonthGetResponse extends jspb.Message {
  getMonth(): PanchangMonth | undefined;
  setMonth(value?: PanchangMonth): PanchangMonthGetResponse;
  hasMonth(): boolean;
  clearMonth(): PanchangMonthGetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangMonthGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangMonthGetResponse): PanchangMonthGetResponse.AsObject;
  static serializeBinaryToWriter(message: PanchangMonthGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangMonthGetResponse;
  static deserializeBinaryFromReader(message: PanchangMonthGetResponse, reader: jspb.BinaryReader): PanchangMonthGetResponse;
}

export namespace PanchangMonthGetResponse {
  export type AsObject = {
    month?: PanchangMonth.AsObject,
  }
}

export class PanchangDayGetRequest extends jspb.Message {
  getBirthdata(): astro_metas_pb.BirthData | undefined;
  setBirthdata(value?: astro_metas_pb.BirthData): PanchangDayGetRequest;
  hasBirthdata(): boolean;
  clearBirthdata(): PanchangDayGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangDayGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangDayGetRequest): PanchangDayGetRequest.AsObject;
  static serializeBinaryToWriter(message: PanchangDayGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangDayGetRequest;
  static deserializeBinaryFromReader(message: PanchangDayGetRequest, reader: jspb.BinaryReader): PanchangDayGetRequest;
}

export namespace PanchangDayGetRequest {
  export type AsObject = {
    birthdata?: astro_metas_pb.BirthData.AsObject,
  }
}

export class PanchangDayGetResponse extends jspb.Message {
  getDay(): PanchangDay | undefined;
  setDay(value?: PanchangDay): PanchangDayGetResponse;
  hasDay(): boolean;
  clearDay(): PanchangDayGetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangDayGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangDayGetResponse): PanchangDayGetResponse.AsObject;
  static serializeBinaryToWriter(message: PanchangDayGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangDayGetResponse;
  static deserializeBinaryFromReader(message: PanchangDayGetResponse, reader: jspb.BinaryReader): PanchangDayGetResponse;
}

export namespace PanchangDayGetResponse {
  export type AsObject = {
    day?: PanchangDay.AsObject,
  }
}


import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class AstroConstants extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AstroConstants.AsObject;
  static toObject(includeInstance: boolean, msg: AstroConstants): AstroConstants.AsObject;
  static serializeBinaryToWriter(message: AstroConstants, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AstroConstants;
  static deserializeBinaryFromReader(message: AstroConstants, reader: jspb.BinaryReader): AstroConstants;
}

export namespace AstroConstants {
  export type AsObject = {
  }

  export enum Planets { 
    LAGNA = 0,
    SUN = 1,
    MOON = 2,
    MARS = 3,
    MERCURY = 4,
    JUPITER = 5,
    VENUS = 6,
    SATURDAY = 7,
    RAHU = 8,
    KETU = 9,
    URANUS = 10,
    NEPTUNE = 11,
    PLUTO = 12,
    MAANDI = 13,
    GULIKA = 14,
    BHAVALAGNA = 15,
    HORALAGNA = 16,
    GHATILAGNA = 17,
    VIGHATILAGNA = 18,
    VAMADALAGNA = 19,
    SREELAGNA = 20,
    PRANAPADALAGNA = 21,
    INDULAGNA = 22,
    BRIGHUBINDU = 23,
    DHOOMA = 24,
    VYATIPATA = 25,
    PARIVESHA = 26,
    INDRACHAPA = 27,
    UPAKETU = 28,
    KAALA = 29,
    MRITYU = 30,
    ARTHAPRAHARA = 31,
    YAMAGHANTAKA = 32,
  }

  export enum Sphutas { 
    IGNORESPHUTA = 0,
    PRANASPHUTA = 1,
    DEHASPHUTA = 2,
    MRITYUSPHUTA = 3,
    SOOKSHAMATRISPHUTA = 4,
    TITHISPHUTA = 5,
    YOGASUNMOONSPHUTA = 6,
    RAHUTITHISPHUTA = 7,
    KSETHRASPHUTA = 8,
    BEEJASPHUTA = 9,
    TRISPHUTA = 10,
    CHATUSSPHUTA = 11,
    PANCHASPHUTA = 12,
    YOGASPHUTA = 13,
    AVAYOGASPHUTA = 14,
  }

  export enum Rasis { 
    IGNORERASIS = 0,
    ARIES = 1,
    TAURUS = 2,
    GEMINI = 3,
    CANCER = 4,
    LEO = 5,
    VIRGO = 6,
    LIBRA = 7,
    SCORPIO = 8,
    SAGGITARIUS = 9,
    CAPRICON = 10,
    AQUARIUS = 11,
    PISCES = 12,
  }

  export enum Stars { 
    IGNORESTARS = 0,
    ASWINI = 1,
    BHARANI = 2,
    KRITTIKA = 3,
    ROHINI = 4,
    MRIGASIRA = 5,
    ARUDRA = 6,
    PUNARVASU = 7,
    PUSHYAMI = 8,
    ASLESHA = 9,
    MAKHA = 10,
    POORVAPHALGUNI = 11,
    UTTARAPHALGUNI = 12,
    HASTA = 13,
    CHITRA = 14,
    SWATI = 15,
    VISHAKA = 16,
    ANURADHA = 17,
    JYESHTA = 18,
    MOOLA = 19,
    POORVASHADA = 20,
    UTTARASHADA = 21,
    SRAVANA = 22,
    DHANISHTA = 23,
    SATABHISHA = 24,
    POORVABHADRA = 25,
    UTTARABADRA = 26,
    REVATHI = 27,
  }
}

export class Metadata extends jspb.Message {
  getUid(): string;
  setUid(value: string): Metadata;

  getName(): string;
  setName(value: string): Metadata;

  getOwner(): string;
  setOwner(value: string): Metadata;

  getTenancy(): string;
  setTenancy(value: string): Metadata;

  getParentuid(): string;
  setParentuid(value: string): Metadata;

  getDataencryption(): boolean;
  setDataencryption(value: boolean): Metadata;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Metadata;

  getAnnotationsMap(): jspb.Map<string, string>;
  clearAnnotationsMap(): Metadata;

  getCreatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedat(value?: google_protobuf_timestamp_pb.Timestamp): Metadata;
  hasCreatedat(): boolean;
  clearCreatedat(): Metadata;

  getUpdatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedat(value?: google_protobuf_timestamp_pb.Timestamp): Metadata;
  hasUpdatedat(): boolean;
  clearUpdatedat(): Metadata;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metadata.AsObject;
  static toObject(includeInstance: boolean, msg: Metadata): Metadata.AsObject;
  static serializeBinaryToWriter(message: Metadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metadata;
  static deserializeBinaryFromReader(message: Metadata, reader: jspb.BinaryReader): Metadata;
}

export namespace Metadata {
  export type AsObject = {
    uid: string,
    name: string,
    owner: string,
    tenancy: string,
    parentuid: string,
    dataencryption: boolean,
    labelsMap: Array<[string, string]>,
    annotationsMap: Array<[string, string]>,
    createdat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class TobData extends jspb.Message {
  getD(): number;
  setD(value: number): TobData;

  getM(): number;
  setM(value: number): TobData;

  getY(): number;
  setY(value: number): TobData;

  getHour(): number;
  setHour(value: number): TobData;

  getMin(): number;
  setMin(value: number): TobData;

  getSec(): number;
  setSec(value: number): TobData;

  getZone(): string;
  setZone(value: string): TobData;

  getDatestring(): string;
  setDatestring(value: string): TobData;

  getTobtime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTobtime(value?: google_protobuf_timestamp_pb.Timestamp): TobData;
  hasTobtime(): boolean;
  clearTobtime(): TobData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TobData.AsObject;
  static toObject(includeInstance: boolean, msg: TobData): TobData.AsObject;
  static serializeBinaryToWriter(message: TobData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TobData;
  static deserializeBinaryFromReader(message: TobData, reader: jspb.BinaryReader): TobData;
}

export namespace TobData {
  export type AsObject = {
    d: number,
    m: number,
    y: number,
    hour: number,
    min: number,
    sec: number,
    zone: string,
    datestring: string,
    tobtime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class PobData extends jspb.Message {
  getCity(): string;
  setCity(value: string): PobData;

  getState(): string;
  setState(value: string): PobData;

  getCountry(): string;
  setCountry(value: string): PobData;

  getLat(): number;
  setLat(value: number): PobData;

  getLong(): number;
  setLong(value: number): PobData;

  getPlanet(): string;
  setPlanet(value: string): PobData;

  getYuga(): string;
  setYuga(value: string): PobData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PobData.AsObject;
  static toObject(includeInstance: boolean, msg: PobData): PobData.AsObject;
  static serializeBinaryToWriter(message: PobData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PobData;
  static deserializeBinaryFromReader(message: PobData, reader: jspb.BinaryReader): PobData;
}

export namespace PobData {
  export type AsObject = {
    city: string,
    state: string,
    country: string,
    lat: number,
    pb_long: number,
    planet: string,
    yuga: string,
  }
}

export class BirthData extends jspb.Message {
  getTob(): TobData | undefined;
  setTob(value?: TobData): BirthData;
  hasTob(): boolean;
  clearTob(): BirthData;

  getPob(): PobData | undefined;
  setPob(value?: PobData): BirthData;
  hasPob(): boolean;
  clearPob(): BirthData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BirthData.AsObject;
  static toObject(includeInstance: boolean, msg: BirthData): BirthData.AsObject;
  static serializeBinaryToWriter(message: BirthData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BirthData;
  static deserializeBinaryFromReader(message: BirthData, reader: jspb.BinaryReader): BirthData;
}

export namespace BirthData {
  export type AsObject = {
    tob?: TobData.AsObject,
    pob?: PobData.AsObject,
  }
}

export class PanchangTransits extends jspb.Message {
  getName(): number;
  setName(value: number): PanchangTransits;

  getTransit(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTransit(value?: google_protobuf_timestamp_pb.Timestamp): PanchangTransits;
  hasTransit(): boolean;
  clearTransit(): PanchangTransits;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangTransits.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangTransits): PanchangTransits.AsObject;
  static serializeBinaryToWriter(message: PanchangTransits, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangTransits;
  static deserializeBinaryFromReader(message: PanchangTransits, reader: jspb.BinaryReader): PanchangTransits;
}

export namespace PanchangTransits {
  export type AsObject = {
    name: number,
    transit?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class PanchangKalam extends jspb.Message {
  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): PanchangKalam;
  hasStarttime(): boolean;
  clearStarttime(): PanchangKalam;

  getEndtime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEndtime(value?: google_protobuf_timestamp_pb.Timestamp): PanchangKalam;
  hasEndtime(): boolean;
  clearEndtime(): PanchangKalam;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangKalam.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangKalam): PanchangKalam.AsObject;
  static serializeBinaryToWriter(message: PanchangKalam, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangKalam;
  static deserializeBinaryFromReader(message: PanchangKalam, reader: jspb.BinaryReader): PanchangKalam;
}

export namespace PanchangKalam {
  export type AsObject = {
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    endtime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

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

export class Panchang extends jspb.Message {
  getYear(): number;
  setYear(value: number): Panchang;

  getMonth(): number;
  setMonth(value: number): Panchang;

  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): Panchang;
  hasStarttime(): boolean;
  clearStarttime(): Panchang;

  getEphimerisList(): Array<SunMoonEphimeris>;
  setEphimerisList(value: Array<SunMoonEphimeris>): Panchang;
  clearEphimerisList(): Panchang;
  addEphimeris(value?: SunMoonEphimeris, index?: number): SunMoonEphimeris;

  getStarsList(): Array<PanchangTransits>;
  setStarsList(value: Array<PanchangTransits>): Panchang;
  clearStarsList(): Panchang;
  addStars(value?: PanchangTransits, index?: number): PanchangTransits;

  getYogasList(): Array<PanchangTransits>;
  setYogasList(value: Array<PanchangTransits>): Panchang;
  clearYogasList(): Panchang;
  addYogas(value?: PanchangTransits, index?: number): PanchangTransits;

  getKaranasList(): Array<PanchangTransits>;
  setKaranasList(value: Array<PanchangTransits>): Panchang;
  clearKaranasList(): Panchang;
  addKaranas(value?: PanchangTransits, index?: number): PanchangTransits;

  getTithisList(): Array<PanchangTransits>;
  setTithisList(value: Array<PanchangTransits>): Panchang;
  clearTithisList(): Panchang;
  addTithis(value?: PanchangTransits, index?: number): PanchangTransits;

  getRahukalam(): PanchangKalam | undefined;
  setRahukalam(value?: PanchangKalam): Panchang;
  hasRahukalam(): boolean;
  clearRahukalam(): Panchang;

  getGulikalam(): PanchangKalam | undefined;
  setGulikalam(value?: PanchangKalam): Panchang;
  hasGulikalam(): boolean;
  clearGulikalam(): Panchang;

  getYamagandam(): PanchangKalam | undefined;
  setYamagandam(value?: PanchangKalam): Panchang;
  hasYamagandam(): boolean;
  clearYamagandam(): Panchang;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Panchang.AsObject;
  static toObject(includeInstance: boolean, msg: Panchang): Panchang.AsObject;
  static serializeBinaryToWriter(message: Panchang, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Panchang;
  static deserializeBinaryFromReader(message: Panchang, reader: jspb.BinaryReader): Panchang;
}

export namespace Panchang {
  export type AsObject = {
    year: number,
    month: number,
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    ephimerisList: Array<SunMoonEphimeris.AsObject>,
    starsList: Array<PanchangTransits.AsObject>,
    yogasList: Array<PanchangTransits.AsObject>,
    karanasList: Array<PanchangTransits.AsObject>,
    tithisList: Array<PanchangTransits.AsObject>,
    rahukalam?: PanchangKalam.AsObject,
    gulikalam?: PanchangKalam.AsObject,
    yamagandam?: PanchangKalam.AsObject,
  }
}

export class PlanetDetails extends jspb.Message {
  getName(): string;
  setName(value: string): PlanetDetails;

  getLongitude(): number;
  setLongitude(value: number): PlanetDetails;

  getLatitude(): number;
  setLatitude(value: number): PlanetDetails;

  getDeclination(): number;
  setDeclination(value: number): PlanetDetails;

  getSpeed(): number;
  setSpeed(value: number): PlanetDetails;

  getDistance(): number;
  setDistance(value: number): PlanetDetails;

  getStar(): number;
  setStar(value: number): PlanetDetails;

  getPaada(): number;
  setPaada(value: number): PlanetDetails;

  getRasi(): number;
  setRasi(value: number): PlanetDetails;

  getBhava(): number;
  setBhava(value: number): PlanetDetails;

  getBhavaH(): number;
  setBhavaH(value: number): PlanetDetails;

  getExalted(): boolean;
  setExalted(value: boolean): PlanetDetails;

  getDeblited(): boolean;
  setDeblited(value: boolean): PlanetDetails;

  getOwnhouse(): boolean;
  setOwnhouse(value: boolean): PlanetDetails;

  getFriendhouse(): boolean;
  setFriendhouse(value: boolean): PlanetDetails;

  getEnemyhouse(): boolean;
  setEnemyhouse(value: boolean): PlanetDetails;

  getMooltrikona(): boolean;
  setMooltrikona(value: boolean): PlanetDetails;

  getAmsa2(): number;
  setAmsa2(value: number): PlanetDetails;

  getAmsa9(): number;
  setAmsa9(value: number): PlanetDetails;

  getAmsa10(): number;
  setAmsa10(value: number): PlanetDetails;

  getAmsa27(): number;
  setAmsa27(value: number): PlanetDetails;

  getAmsa60(): number;
  setAmsa60(value: number): PlanetDetails;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlanetDetails.AsObject;
  static toObject(includeInstance: boolean, msg: PlanetDetails): PlanetDetails.AsObject;
  static serializeBinaryToWriter(message: PlanetDetails, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlanetDetails;
  static deserializeBinaryFromReader(message: PlanetDetails, reader: jspb.BinaryReader): PlanetDetails;
}

export namespace PlanetDetails {
  export type AsObject = {
    name: string,
    longitude: number,
    latitude: number,
    declination: number,
    speed: number,
    distance: number,
    star: number,
    paada: number,
    rasi: number,
    bhava: number,
    bhavaH: number,
    exalted: boolean,
    deblited: boolean,
    ownhouse: boolean,
    friendhouse: boolean,
    enemyhouse: boolean,
    mooltrikona: boolean,
    amsa2: number,
    amsa9: number,
    amsa10: number,
    amsa27: number,
    amsa60: number,
  }
}

export class Planets extends jspb.Message {
  getEphimeris(): string;
  setEphimeris(value: string): Planets;

  getPlanetsmapMap(): jspb.Map<string, PlanetDetails>;
  clearPlanetsmapMap(): Planets;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Planets.AsObject;
  static toObject(includeInstance: boolean, msg: Planets): Planets.AsObject;
  static serializeBinaryToWriter(message: Planets, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Planets;
  static deserializeBinaryFromReader(message: Planets, reader: jspb.BinaryReader): Planets;
}

export namespace Planets {
  export type AsObject = {
    ephimeris: string,
    planetsmapMap: Array<[string, PlanetDetails.AsObject]>,
  }
}

export class AstroObject extends jspb.Message {
  getMetadata(): Metadata | undefined;
  setMetadata(value?: Metadata): AstroObject;
  hasMetadata(): boolean;
  clearMetadata(): AstroObject;

  getBirthdata(): BirthData | undefined;
  setBirthdata(value?: BirthData): AstroObject;
  hasBirthdata(): boolean;
  clearBirthdata(): AstroObject;

  getPanchang(): Panchang | undefined;
  setPanchang(value?: Panchang): AstroObject;
  hasPanchang(): boolean;
  clearPanchang(): AstroObject;

  getPlanets(): Planets | undefined;
  setPlanets(value?: Planets): AstroObject;
  hasPlanets(): boolean;
  clearPlanets(): AstroObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AstroObject.AsObject;
  static toObject(includeInstance: boolean, msg: AstroObject): AstroObject.AsObject;
  static serializeBinaryToWriter(message: AstroObject, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AstroObject;
  static deserializeBinaryFromReader(message: AstroObject, reader: jspb.BinaryReader): AstroObject;
}

export namespace AstroObject {
  export type AsObject = {
    metadata?: Metadata.AsObject,
    birthdata?: BirthData.AsObject,
    panchang?: Panchang.AsObject,
    planets?: Planets.AsObject,
  }
}

export class PanchangGetRequest extends jspb.Message {
  getBirthdata(): BirthData | undefined;
  setBirthdata(value?: BirthData): PanchangGetRequest;
  hasBirthdata(): boolean;
  clearBirthdata(): PanchangGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangGetRequest): PanchangGetRequest.AsObject;
  static serializeBinaryToWriter(message: PanchangGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangGetRequest;
  static deserializeBinaryFromReader(message: PanchangGetRequest, reader: jspb.BinaryReader): PanchangGetRequest;
}

export namespace PanchangGetRequest {
  export type AsObject = {
    birthdata?: BirthData.AsObject,
  }
}

export class PanchangResponse extends jspb.Message {
  getAstro(): AstroObject | undefined;
  setAstro(value?: AstroObject): PanchangResponse;
  hasAstro(): boolean;
  clearAstro(): PanchangResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PanchangResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PanchangResponse): PanchangResponse.AsObject;
  static serializeBinaryToWriter(message: PanchangResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PanchangResponse;
  static deserializeBinaryFromReader(message: PanchangResponse, reader: jspb.BinaryReader): PanchangResponse;
}

export namespace PanchangResponse {
  export type AsObject = {
    astro?: AstroObject.AsObject,
  }
}

export class UserUpdateRequest extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): UserUpdateRequest;

  getBirthdata(): BirthData | undefined;
  setBirthdata(value?: BirthData): UserUpdateRequest;
  hasBirthdata(): boolean;
  clearBirthdata(): UserUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserUpdateRequest): UserUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UserUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserUpdateRequest;
  static deserializeBinaryFromReader(message: UserUpdateRequest, reader: jspb.BinaryReader): UserUpdateRequest;
}

export namespace UserUpdateRequest {
  export type AsObject = {
    uuid: string,
    birthdata?: BirthData.AsObject,
  }
}

export class UserUpdateResponse extends jspb.Message {
  getAstro(): AstroObject | undefined;
  setAstro(value?: AstroObject): UserUpdateResponse;
  hasAstro(): boolean;
  clearAstro(): UserUpdateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserUpdateResponse): UserUpdateResponse.AsObject;
  static serializeBinaryToWriter(message: UserUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserUpdateResponse;
  static deserializeBinaryFromReader(message: UserUpdateResponse, reader: jspb.BinaryReader): UserUpdateResponse;
}

export namespace UserUpdateResponse {
  export type AsObject = {
    astro?: AstroObject.AsObject,
  }
}

export class UserGetRequest extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): UserGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserGetRequest): UserGetRequest.AsObject;
  static serializeBinaryToWriter(message: UserGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserGetRequest;
  static deserializeBinaryFromReader(message: UserGetRequest, reader: jspb.BinaryReader): UserGetRequest;
}

export namespace UserGetRequest {
  export type AsObject = {
    uuid: string,
  }
}

export class UserGetResponse extends jspb.Message {
  getAstro(): AstroObject | undefined;
  setAstro(value?: AstroObject): UserGetResponse;
  hasAstro(): boolean;
  clearAstro(): UserGetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserGetResponse): UserGetResponse.AsObject;
  static serializeBinaryToWriter(message: UserGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserGetResponse;
  static deserializeBinaryFromReader(message: UserGetResponse, reader: jspb.BinaryReader): UserGetResponse;
}

export namespace UserGetResponse {
  export type AsObject = {
    astro?: AstroObject.AsObject,
  }
}

export class UserListRequest extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): UserListRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserListRequest): UserListRequest.AsObject;
  static serializeBinaryToWriter(message: UserListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserListRequest;
  static deserializeBinaryFromReader(message: UserListRequest, reader: jspb.BinaryReader): UserListRequest;
}

export namespace UserListRequest {
  export type AsObject = {
    uuid: string,
  }
}

export class UserListResponse extends jspb.Message {
  getAstroList(): Array<AstroObject>;
  setAstroList(value: Array<AstroObject>): UserListResponse;
  clearAstroList(): UserListResponse;
  addAstro(value?: AstroObject, index?: number): AstroObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserListResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserListResponse): UserListResponse.AsObject;
  static serializeBinaryToWriter(message: UserListResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserListResponse;
  static deserializeBinaryFromReader(message: UserListResponse, reader: jspb.BinaryReader): UserListResponse;
}

export namespace UserListResponse {
  export type AsObject = {
    astroList: Array<AstroObject.AsObject>,
  }
}


import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as astro_metas_pb from '../astro/metas_pb';


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

  getCharttype(): Metadata.ChartType;
  setCharttype(value: Metadata.ChartType): Metadata;

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
    charttype: Metadata.ChartType,
  }

  export enum ChartType { 
    IGNORECHARTTYPE = 0,
    SELF = 1,
    GLOBAL = 2,
    CLIENT = 3,
    TEST = 4,
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

  getMooltrikona(): boolean;
  setMooltrikona(value: boolean): PlanetDetails;

  getDeblited(): boolean;
  setDeblited(value: boolean): PlanetDetails;

  getOwnhouse(): boolean;
  setOwnhouse(value: boolean): PlanetDetails;

  getFriendhouse(): boolean;
  setFriendhouse(value: boolean): PlanetDetails;

  getEnemyhouse(): boolean;
  setEnemyhouse(value: boolean): PlanetDetails;

  getNeutralhouse(): boolean;
  setNeutralhouse(value: boolean): PlanetDetails;

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
    mooltrikona: boolean,
    deblited: boolean,
    ownhouse: boolean,
    friendhouse: boolean,
    enemyhouse: boolean,
    neutralhouse: boolean,
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

  getBirthdata(): astro_metas_pb.BirthData | undefined;
  setBirthdata(value?: astro_metas_pb.BirthData): AstroObject;
  hasBirthdata(): boolean;
  clearBirthdata(): AstroObject;

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
    birthdata?: astro_metas_pb.BirthData.AsObject,
    planets?: Planets.AsObject,
  }
}

export class UserCreateRequest extends jspb.Message {
  getBirthdata(): astro_metas_pb.BirthData | undefined;
  setBirthdata(value?: astro_metas_pb.BirthData): UserCreateRequest;
  hasBirthdata(): boolean;
  clearBirthdata(): UserCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserCreateRequest): UserCreateRequest.AsObject;
  static serializeBinaryToWriter(message: UserCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserCreateRequest;
  static deserializeBinaryFromReader(message: UserCreateRequest, reader: jspb.BinaryReader): UserCreateRequest;
}

export namespace UserCreateRequest {
  export type AsObject = {
    birthdata?: astro_metas_pb.BirthData.AsObject,
  }
}

export class UserCreateResponse extends jspb.Message {
  getAstro(): AstroObject | undefined;
  setAstro(value?: AstroObject): UserCreateResponse;
  hasAstro(): boolean;
  clearAstro(): UserCreateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserCreateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserCreateResponse): UserCreateResponse.AsObject;
  static serializeBinaryToWriter(message: UserCreateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserCreateResponse;
  static deserializeBinaryFromReader(message: UserCreateResponse, reader: jspb.BinaryReader): UserCreateResponse;
}

export namespace UserCreateResponse {
  export type AsObject = {
    astro?: AstroObject.AsObject,
  }
}

export class UserUpdateRequest extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): UserUpdateRequest;

  getBirthdata(): astro_metas_pb.BirthData | undefined;
  setBirthdata(value?: astro_metas_pb.BirthData): UserUpdateRequest;
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
    birthdata?: astro_metas_pb.BirthData.AsObject,
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


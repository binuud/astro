import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class SubDashaMeta extends jspb.Message {
  getName(): string;
  setName(value: string): SubDashaMeta;

  getDuration(): number;
  setDuration(value: number): SubDashaMeta;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubDashaMeta.AsObject;
  static toObject(includeInstance: boolean, msg: SubDashaMeta): SubDashaMeta.AsObject;
  static serializeBinaryToWriter(message: SubDashaMeta, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubDashaMeta;
  static deserializeBinaryFromReader(message: SubDashaMeta, reader: jspb.BinaryReader): SubDashaMeta;
}

export namespace SubDashaMeta {
  export type AsObject = {
    name: string,
    duration: number,
  }
}

export class DashaMeta extends jspb.Message {
  getName(): string;
  setName(value: string): DashaMeta;

  getYearinseconds(): number;
  setYearinseconds(value: number): DashaMeta;

  getDurationyears(): number;
  setDurationyears(value: number): DashaMeta;

  getStartdasa(): number;
  setStartdasa(value: number): DashaMeta;

  getDasadegrees(): number;
  setDasadegrees(value: number): DashaMeta;

  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): DashaMeta;
  hasStarttime(): boolean;
  clearStarttime(): DashaMeta;

  getSubdashametasList(): Array<SubDashaMeta>;
  setSubdashametasList(value: Array<SubDashaMeta>): DashaMeta;
  clearSubdashametasList(): DashaMeta;
  addSubdashametas(value?: SubDashaMeta, index?: number): SubDashaMeta;

  getType(): DashaMeta.Types;
  setType(value: DashaMeta.Types): DashaMeta;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaMeta.AsObject;
  static toObject(includeInstance: boolean, msg: DashaMeta): DashaMeta.AsObject;
  static serializeBinaryToWriter(message: DashaMeta, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaMeta;
  static deserializeBinaryFromReader(message: DashaMeta, reader: jspb.BinaryReader): DashaMeta;
}

export namespace DashaMeta {
  export type AsObject = {
    name: string,
    yearinseconds: number,
    durationyears: number,
    startdasa: number,
    dasadegrees: number,
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    subdashametasList: Array<SubDashaMeta.AsObject>,
    type: DashaMeta.Types,
  }

  export enum Types { 
    IGNORE = 0,
    VIMSHOTTARI = 1,
    ASHTOTTARI = 2,
  }
}

export class DashaItem extends jspb.Message {
  getName(): string;
  setName(value: string): DashaItem;

  getLevel(): number;
  setLevel(value: number): DashaItem;

  getStarttime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStarttime(value?: google_protobuf_timestamp_pb.Timestamp): DashaItem;
  hasStarttime(): boolean;
  clearStarttime(): DashaItem;

  getEndtime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEndtime(value?: google_protobuf_timestamp_pb.Timestamp): DashaItem;
  hasEndtime(): boolean;
  clearEndtime(): DashaItem;

  getDashasList(): Array<DashaItem>;
  setDashasList(value: Array<DashaItem>): DashaItem;
  clearDashasList(): DashaItem;
  addDashas(value?: DashaItem, index?: number): DashaItem;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaItem.AsObject;
  static toObject(includeInstance: boolean, msg: DashaItem): DashaItem.AsObject;
  static serializeBinaryToWriter(message: DashaItem, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaItem;
  static deserializeBinaryFromReader(message: DashaItem, reader: jspb.BinaryReader): DashaItem;
}

export namespace DashaItem {
  export type AsObject = {
    name: string,
    level: number,
    starttime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    endtime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    dashasList: Array<DashaItem.AsObject>,
  }
}

export class DashaSystem extends jspb.Message {
  getMeta(): DashaMeta | undefined;
  setMeta(value?: DashaMeta): DashaSystem;
  hasMeta(): boolean;
  clearMeta(): DashaSystem;

  getDashasList(): Array<DashaItem>;
  setDashasList(value: Array<DashaItem>): DashaSystem;
  clearDashasList(): DashaSystem;
  addDashas(value?: DashaItem, index?: number): DashaItem;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaSystem.AsObject;
  static toObject(includeInstance: boolean, msg: DashaSystem): DashaSystem.AsObject;
  static serializeBinaryToWriter(message: DashaSystem, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaSystem;
  static deserializeBinaryFromReader(message: DashaSystem, reader: jspb.BinaryReader): DashaSystem;
}

export namespace DashaSystem {
  export type AsObject = {
    meta?: DashaMeta.AsObject,
    dashasList: Array<DashaItem.AsObject>,
  }
}

export class DashaGetReqeust extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): DashaGetReqeust;

  getLevel(): number;
  setLevel(value: number): DashaGetReqeust;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaGetReqeust.AsObject;
  static toObject(includeInstance: boolean, msg: DashaGetReqeust): DashaGetReqeust.AsObject;
  static serializeBinaryToWriter(message: DashaGetReqeust, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaGetReqeust;
  static deserializeBinaryFromReader(message: DashaGetReqeust, reader: jspb.BinaryReader): DashaGetReqeust;
}

export namespace DashaGetReqeust {
  export type AsObject = {
    uuid: string,
    level: number,
  }
}

export class DashaGetResponse extends jspb.Message {
  getUserUuid(): string;
  setUserUuid(value: string): DashaGetResponse;

  getDashasystem(): DashaSystem | undefined;
  setDashasystem(value?: DashaSystem): DashaGetResponse;
  hasDashasystem(): boolean;
  clearDashasystem(): DashaGetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DashaGetResponse): DashaGetResponse.AsObject;
  static serializeBinaryToWriter(message: DashaGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaGetResponse;
  static deserializeBinaryFromReader(message: DashaGetResponse, reader: jspb.BinaryReader): DashaGetResponse;
}

export namespace DashaGetResponse {
  export type AsObject = {
    userUuid: string,
    dashasystem?: DashaSystem.AsObject,
  }
}

export class DashaUpdateReqeust extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): DashaUpdateReqeust;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaUpdateReqeust.AsObject;
  static toObject(includeInstance: boolean, msg: DashaUpdateReqeust): DashaUpdateReqeust.AsObject;
  static serializeBinaryToWriter(message: DashaUpdateReqeust, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaUpdateReqeust;
  static deserializeBinaryFromReader(message: DashaUpdateReqeust, reader: jspb.BinaryReader): DashaUpdateReqeust;
}

export namespace DashaUpdateReqeust {
  export type AsObject = {
    uuid: string,
  }
}

export class DashaUpdateResponse extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): DashaUpdateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashaUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DashaUpdateResponse): DashaUpdateResponse.AsObject;
  static serializeBinaryToWriter(message: DashaUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashaUpdateResponse;
  static deserializeBinaryFromReader(message: DashaUpdateResponse, reader: jspb.BinaryReader): DashaUpdateResponse;
}

export namespace DashaUpdateResponse {
  export type AsObject = {
    uuid: string,
  }
}


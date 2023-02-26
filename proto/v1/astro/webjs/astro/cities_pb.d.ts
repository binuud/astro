import * as jspb from 'google-protobuf'



export class CityInfo extends jspb.Message {
  getId(): number;
  setId(value: number): CityInfo;

  getName(): string;
  setName(value: string): CityInfo;

  getStateId(): number;
  setStateId(value: number): CityInfo;

  getStateCode(): string;
  setStateCode(value: string): CityInfo;

  getStateName(): string;
  setStateName(value: string): CityInfo;

  getCountryId(): number;
  setCountryId(value: number): CityInfo;

  getCountryCode(): string;
  setCountryCode(value: string): CityInfo;

  getCountryName(): string;
  setCountryName(value: string): CityInfo;

  getLatitude(): number;
  setLatitude(value: number): CityInfo;

  getLongitude(): number;
  setLongitude(value: number): CityInfo;

  getWikidataid(): string;
  setWikidataid(value: string): CityInfo;

  getIndex(): number;
  setIndex(value: number): CityInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CityInfo.AsObject;
  static toObject(includeInstance: boolean, msg: CityInfo): CityInfo.AsObject;
  static serializeBinaryToWriter(message: CityInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CityInfo;
  static deserializeBinaryFromReader(message: CityInfo, reader: jspb.BinaryReader): CityInfo;
}

export namespace CityInfo {
  export type AsObject = {
    id: number,
    name: string,
    stateId: number,
    stateCode: string,
    stateName: string,
    countryId: number,
    countryCode: string,
    countryName: string,
    latitude: number,
    longitude: number,
    wikidataid: string,
    index: number,
  }
}

export class CitiesGetRequest extends jspb.Message {
  getId(): number;
  setId(value: number): CitiesGetRequest;

  getName(): string;
  setName(value: string): CitiesGetRequest;

  getStateId(): number;
  setStateId(value: number): CitiesGetRequest;

  getStateCode(): string;
  setStateCode(value: string): CitiesGetRequest;

  getStateName(): string;
  setStateName(value: string): CitiesGetRequest;

  getCountryId(): number;
  setCountryId(value: number): CitiesGetRequest;

  getCountryCode(): string;
  setCountryCode(value: string): CitiesGetRequest;

  getCountryName(): string;
  setCountryName(value: string): CitiesGetRequest;

  getIndex(): number;
  setIndex(value: number): CitiesGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CitiesGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CitiesGetRequest): CitiesGetRequest.AsObject;
  static serializeBinaryToWriter(message: CitiesGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CitiesGetRequest;
  static deserializeBinaryFromReader(message: CitiesGetRequest, reader: jspb.BinaryReader): CitiesGetRequest;
}

export namespace CitiesGetRequest {
  export type AsObject = {
    id: number,
    name: string,
    stateId: number,
    stateCode: string,
    stateName: string,
    countryId: number,
    countryCode: string,
    countryName: string,
    index: number,
  }
}

export class CitiesGetResponse extends jspb.Message {
  getDetails(): CityInfo | undefined;
  setDetails(value?: CityInfo): CitiesGetResponse;
  hasDetails(): boolean;
  clearDetails(): CitiesGetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CitiesGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CitiesGetResponse): CitiesGetResponse.AsObject;
  static serializeBinaryToWriter(message: CitiesGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CitiesGetResponse;
  static deserializeBinaryFromReader(message: CitiesGetResponse, reader: jspb.BinaryReader): CitiesGetResponse;
}

export namespace CitiesGetResponse {
  export type AsObject = {
    details?: CityInfo.AsObject,
  }
}

export class CitiesSearchRequest extends jspb.Message {
  getId(): number;
  setId(value: number): CitiesSearchRequest;

  getName(): string;
  setName(value: string): CitiesSearchRequest;

  getStateId(): number;
  setStateId(value: number): CitiesSearchRequest;

  getStateCode(): string;
  setStateCode(value: string): CitiesSearchRequest;

  getStateName(): string;
  setStateName(value: string): CitiesSearchRequest;

  getCountryId(): number;
  setCountryId(value: number): CitiesSearchRequest;

  getCountryCode(): string;
  setCountryCode(value: string): CitiesSearchRequest;

  getCountryName(): string;
  setCountryName(value: string): CitiesSearchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CitiesSearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CitiesSearchRequest): CitiesSearchRequest.AsObject;
  static serializeBinaryToWriter(message: CitiesSearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CitiesSearchRequest;
  static deserializeBinaryFromReader(message: CitiesSearchRequest, reader: jspb.BinaryReader): CitiesSearchRequest;
}

export namespace CitiesSearchRequest {
  export type AsObject = {
    id: number,
    name: string,
    stateId: number,
    stateCode: string,
    stateName: string,
    countryId: number,
    countryCode: string,
    countryName: string,
  }
}

export class CitiesSearchResponse extends jspb.Message {
  getCitiesList(): Array<CityInfo>;
  setCitiesList(value: Array<CityInfo>): CitiesSearchResponse;
  clearCitiesList(): CitiesSearchResponse;
  addCities(value?: CityInfo, index?: number): CityInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CitiesSearchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CitiesSearchResponse): CitiesSearchResponse.AsObject;
  static serializeBinaryToWriter(message: CitiesSearchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CitiesSearchResponse;
  static deserializeBinaryFromReader(message: CitiesSearchResponse, reader: jspb.BinaryReader): CitiesSearchResponse;
}

export namespace CitiesSearchResponse {
  export type AsObject = {
    citiesList: Array<CityInfo.AsObject>,
  }
}


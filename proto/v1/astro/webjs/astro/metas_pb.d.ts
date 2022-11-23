import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


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

export class GeoLocation extends jspb.Message {
  getLat(): number;
  setLat(value: number): GeoLocation;

  getLong(): number;
  setLong(value: number): GeoLocation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GeoLocation.AsObject;
  static toObject(includeInstance: boolean, msg: GeoLocation): GeoLocation.AsObject;
  static serializeBinaryToWriter(message: GeoLocation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GeoLocation;
  static deserializeBinaryFromReader(message: GeoLocation, reader: jspb.BinaryReader): GeoLocation;
}

export namespace GeoLocation {
  export type AsObject = {
    lat: number,
    pb_long: number,
  }
}

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

  export enum Weekdays { 
    SUNDAY = 0,
    MONDAY = 1,
    TUESDAY = 2,
    WEDNESDAY = 3,
    THURSDAY = 4,
    FRIDAY = 5,
    SATURDAY = 6,
  }

  export enum Planets { 
    LAGNA = 0,
    SUN = 1,
    MOON = 2,
    MARS = 3,
    MERCURY = 4,
    JUPITER = 5,
    VENUS = 6,
    SATURN = 7,
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

  export enum Tithis { 
    IGNORE_TITHI = 0,
    S_PRATHIMA = 1,
    S_DWITIYA = 2,
    S_TRITIYA = 3,
    S_CHATURTHI = 4,
    S_PANCHAMI = 5,
    S_SHASTI = 6,
    S_SAPTAMI = 7,
    S_ASHTAMI = 8,
    S_NAVAMI = 9,
    S_DASHAMI = 10,
    S_EKADASHI = 11,
    S_DWADASHI = 12,
    S_THRAYODASHI = 13,
    S_CHATURDASHI = 14,
    PURNIMA = 15,
    K_PRATHAMA = 16,
    K_DWITIYA = 17,
    K_TRITIYA = 18,
    K_CHATURTHI = 19,
    K_PANCHAMI = 20,
    K_SHASTI = 21,
    K_SAPTAMI = 22,
    K_ASHTAMI = 23,
    K_NAVAMI = 24,
    K_DASHAMI = 25,
    K_EKADASHI = 26,
    K_DWADASHI = 27,
    K_THRAYODASHI = 28,
    K_CHATURDASHI = 29,
    AMAVASYA = 30,
  }
}

export class AstroMetas extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AstroMetas.AsObject;
  static toObject(includeInstance: boolean, msg: AstroMetas): AstroMetas.AsObject;
  static serializeBinaryToWriter(message: AstroMetas, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AstroMetas;
  static deserializeBinaryFromReader(message: AstroMetas, reader: jspb.BinaryReader): AstroMetas;
}

export namespace AstroMetas {
  export type AsObject = {
  }

  export enum Nature { 
    IGNORENATURE = 0,
    BENEFIC = 1,
    MALEFIC = 2,
    CHOOSY = 3,
    BOTH = 4,
    AVOID = 5,
  }

  export enum MetaTags { 
    IGNORE = 0,
    GENDER_MALE = 1,
    GENDER_FEMALE = 2,
    GENDER_NEUTRAL = 3,
    PADDING_4 = 4,
    PADDING_5 = 5,
    PADDING_6 = 6,
    PADDING_7 = 7,
    PADDING_8 = 8,
    PADDING_9 = 9,
    ELEMENT_AIR = 10,
    ELEMENT_FIRE = 11,
    ELEMENT_WATER = 12,
    ELEMENT_EARTH = 13,
    QUALITY_FIXED = 14,
    QUALITY_MOVEABLE = 15,
    QUALITY_MUTABLE = 16,
    QUALITY_CARDINAL = 17,
    PADDING_18 = 18,
    PADDING_19 = 19,
    STATE_COMMON = 20,
    STATE_LAME = 21,
    BLIND_DAY = 22,
    BLIND_NIGHT = 23,
    DEAF_MORNING = 24,
    DEAF_MIDDAY = 25,
    DEAF_NIGHT = 26,
    PADDING_27 = 27,
    PADDING_28 = 28,
    PADDING_29 = 29,
    TITHI_NANDA = 30,
    TITHI_BHADRA = 31,
    TITHI_JAYA = 32,
    TITHI_RIKTHA = 33,
    TITHI_POORNA = 34,
    TRIDOSHA_VAATHA = 35,
    TRIDOSHA_PITHA = 36,
    TRIDOSHA_KAPHA = 37,
    PADDING_38 = 38,
    PADDING_39 = 39,
    BHAVA_TRIKONA = 40,
    BHAVA_KENDRA = 41,
    BHAVA_UPACHAYA = 42,
    BHAVA_DUSTHANA = 43,
    BHAVA_TRISHADAYAS = 44,
    BHAVA_MARAKASTHANA = 45,
    PADDING_46 = 46,
    PADDING_47 = 47,
    PADDING_48 = 48,
    PADDING_49 = 49,
    RASI_EXALTED = 50,
    RASI_MOOLATRIKONA = 51,
    RASI_DEBLITED = 52,
  }

  export enum PanchangTags { 
    IGNORE_PANCHANGTAG = 0,
    KRISHNAPAKSHA = 1,
    SHUKLAPAKSHA = 2,
    AMAVASYA = 3,
    POORNIMA = 4,
    SANKATAHARACHATURTI = 5,
    SHASTI = 6,
    EKADASI = 7,
    AKSHAYATRITIYA = 8,
    KRITHIKA = 9,
    PRADOSHAM = 10,
    SIVARATRI = 11,
    MAHASIVARATRI = 12,
    MARRIAGEMUHURTA = 13,
  }

  export enum AstroTypes { 
    IGNORE_ASTRO_TYPE = 0,
    PLANET = 1,
    RASI = 2,
    PLANET_CHAAYA = 3,
    STAR = 5,
    SPECIAL_STAR = 6,
  }
}

export class AstroMetadata extends jspb.Message {
  getUid(): string;
  setUid(value: string): AstroMetadata;

  getName(): string;
  setName(value: string): AstroMetadata;

  getOwner(): string;
  setOwner(value: string): AstroMetadata;

  getTenancy(): string;
  setTenancy(value: string): AstroMetadata;

  getParentuid(): string;
  setParentuid(value: string): AstroMetadata;

  getType(): AstroMetas.AstroTypes;
  setType(value: AstroMetas.AstroTypes): AstroMetadata;

  getTagsList(): Array<AstroMetas.MetaTags>;
  setTagsList(value: Array<AstroMetas.MetaTags>): AstroMetadata;
  clearTagsList(): AstroMetadata;
  addTags(value: AstroMetas.MetaTags, index?: number): AstroMetadata;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): AstroMetadata;

  getAnnotationsMap(): jspb.Map<string, string>;
  clearAnnotationsMap(): AstroMetadata;

  getDescription(): string;
  setDescription(value: string): AstroMetadata;

  getCreatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedat(value?: google_protobuf_timestamp_pb.Timestamp): AstroMetadata;
  hasCreatedat(): boolean;
  clearCreatedat(): AstroMetadata;

  getUpdatedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedat(value?: google_protobuf_timestamp_pb.Timestamp): AstroMetadata;
  hasUpdatedat(): boolean;
  clearUpdatedat(): AstroMetadata;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AstroMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: AstroMetadata): AstroMetadata.AsObject;
  static serializeBinaryToWriter(message: AstroMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AstroMetadata;
  static deserializeBinaryFromReader(message: AstroMetadata, reader: jspb.BinaryReader): AstroMetadata;
}

export namespace AstroMetadata {
  export type AsObject = {
    uid: string,
    name: string,
    owner: string,
    tenancy: string,
    parentuid: string,
    type: AstroMetas.AstroTypes,
    tagsList: Array<AstroMetas.MetaTags>,
    labelsMap: Array<[string, string]>,
    annotationsMap: Array<[string, string]>,
    description: string,
    createdat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class StarWeekdayYoga extends jspb.Message {
  getName(): string;
  setName(value: string): StarWeekdayYoga;

  getYoga(): StarWeekdayYoga.StarWeekdayYogaTypes;
  setYoga(value: StarWeekdayYoga.StarWeekdayYogaTypes): StarWeekdayYoga;

  getStarsList(): Array<AstroConstants.Stars>;
  setStarsList(value: Array<AstroConstants.Stars>): StarWeekdayYoga;
  clearStarsList(): StarWeekdayYoga;
  addStars(value: AstroConstants.Stars, index?: number): StarWeekdayYoga;

  getNature(): AstroMetas.Nature;
  setNature(value: AstroMetas.Nature): StarWeekdayYoga;

  getDesc(): string;
  setDesc(value: string): StarWeekdayYoga;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StarWeekdayYoga.AsObject;
  static toObject(includeInstance: boolean, msg: StarWeekdayYoga): StarWeekdayYoga.AsObject;
  static serializeBinaryToWriter(message: StarWeekdayYoga, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StarWeekdayYoga;
  static deserializeBinaryFromReader(message: StarWeekdayYoga, reader: jspb.BinaryReader): StarWeekdayYoga;
}

export namespace StarWeekdayYoga {
  export type AsObject = {
    name: string,
    yoga: StarWeekdayYoga.StarWeekdayYogaTypes,
    starsList: Array<AstroConstants.Stars>,
    nature: AstroMetas.Nature,
    desc: string,
  }

  export enum StarWeekdayYogaTypes { 
    IGNORESTARYOGA = 0,
    ANANDA = 1,
    KAALDANDA = 2,
    DHUMRA = 3,
    DHATA = 4,
    SAUMYA = 5,
    DHWANSHA = 6,
    KETU = 7,
    SRIVATSA = 8,
    VAJRA = 9,
    MUDGARA = 10,
    CHHATRA = 11,
    MITRA = 12,
    MANAS = 13,
    PADMA = 14,
    LUMBA = 15,
    UTPAATA = 16,
    MRITYU = 17,
    KAANA = 18,
    SIDDHI = 19,
    SUBHA = 20,
    AMRITA = 21,
    MUSHALA = 22,
    GADA = 23,
    MATANGA = 24,
    RAKSHA = 25,
    CHARA = 26,
    SUSTHIRA = 27,
    PRAVARDHAMAMA = 28,
  }
}

export class StarThithiWeekdayYoga extends jspb.Message {
  getYoga(): StarThithiWeekdayYoga.StarThithiWeekdayYogaTypes;
  setYoga(value: StarThithiWeekdayYoga.StarThithiWeekdayYogaTypes): StarThithiWeekdayYoga;

  getNature(): AstroMetas.Nature;
  setNature(value: AstroMetas.Nature): StarThithiWeekdayYoga;

  getDesc(): string;
  setDesc(value: string): StarThithiWeekdayYoga;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StarThithiWeekdayYoga.AsObject;
  static toObject(includeInstance: boolean, msg: StarThithiWeekdayYoga): StarThithiWeekdayYoga.AsObject;
  static serializeBinaryToWriter(message: StarThithiWeekdayYoga, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StarThithiWeekdayYoga;
  static deserializeBinaryFromReader(message: StarThithiWeekdayYoga, reader: jspb.BinaryReader): StarThithiWeekdayYoga;
}

export namespace StarThithiWeekdayYoga {
  export type AsObject = {
    yoga: StarThithiWeekdayYoga.StarThithiWeekdayYogaTypes,
    nature: AstroMetas.Nature,
    desc: string,
  }

  export enum StarThithiWeekdayYogaTypes { 
    IGNORESTARTHITHIWEEKDAYYOGA = 0,
    TRIPUSHKARAYOGA = 1,
    DWIPUSHKARAYOGA = 2,
  }
}

export class YogaIdentifier extends jspb.Message {
  getType(): YogaIdentifier.YogaTypes;
  setType(value: YogaIdentifier.YogaTypes): YogaIdentifier;

  getYogastarweekday(): StarWeekdayYoga | undefined;
  setYogastarweekday(value?: StarWeekdayYoga): YogaIdentifier;
  hasYogastarweekday(): boolean;
  clearYogastarweekday(): YogaIdentifier;

  getYogastarthithiweekday(): StarThithiWeekdayYoga | undefined;
  setYogastarthithiweekday(value?: StarThithiWeekdayYoga): YogaIdentifier;
  hasYogastarthithiweekday(): boolean;
  clearYogastarthithiweekday(): YogaIdentifier;

  getYogaCase(): YogaIdentifier.YogaCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YogaIdentifier.AsObject;
  static toObject(includeInstance: boolean, msg: YogaIdentifier): YogaIdentifier.AsObject;
  static serializeBinaryToWriter(message: YogaIdentifier, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YogaIdentifier;
  static deserializeBinaryFromReader(message: YogaIdentifier, reader: jspb.BinaryReader): YogaIdentifier;
}

export namespace YogaIdentifier {
  export type AsObject = {
    type: YogaIdentifier.YogaTypes,
    yogastarweekday?: StarWeekdayYoga.AsObject,
    yogastarthithiweekday?: StarThithiWeekdayYoga.AsObject,
  }

  export enum YogaTypes { 
    IGNOREYOGATYPES = 0,
    STARWEEKDAYYOGA = 1,
    STARTHITHIWEEKDAYYOGA = 2,
  }

  export enum YogaCase { 
    YOGA_NOT_SET = 0,
    YOGASTARWEEKDAY = 2,
    YOGASTARTHITHIWEEKDAY = 3,
  }
}


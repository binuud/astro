/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

export enum AstroConstantsWeekdays {
  Sunday = "Sunday",
  Monday = "Monday",
  Tuesday = "Tuesday",
  Wednesday = "Wednesday",
  Thursday = "Thursday",
  Friday = "Friday",
  Saturday = "Saturday",
}

export enum AstroConstantsPlanets {
  Lagna = "Lagna",
  Sun = "Sun",
  Moon = "Moon",
  Mars = "Mars",
  Mercury = "Mercury",
  Jupiter = "Jupiter",
  Venus = "Venus",
  Saturn = "Saturn",
  Rahu = "Rahu",
  Ketu = "Ketu",
  Uranus = "Uranus",
  Neptune = "Neptune",
  Pluto = "Pluto",
  Maandi = "Maandi",
  Gulika = "Gulika",
  BhavaLagna = "BhavaLagna",
  HoraLagna = "HoraLagna",
  GhatiLagna = "GhatiLagna",
  VighatiLagna = "VighatiLagna",
  VamadaLagna = "VamadaLagna",
  SreeLagna = "SreeLagna",
  PranapadaLagna = "PranapadaLagna",
  InduLagna = "InduLagna",
  BrighuBindu = "BrighuBindu",
  Dhooma = "Dhooma",
  Vyatipata = "Vyatipata",
  Parivesha = "Parivesha",
  IndraChapa = "IndraChapa",
  Upaketu = "Upaketu",
  Kaala = "Kaala",
  Mrityu = "Mrityu",
  ArthaPrahara = "ArthaPrahara",
  YamaGhantaka = "YamaGhantaka",
}

export enum AstroConstantsSphutas {
  IgnoreSphuta = "IgnoreSphuta",
  PranaSphuta = "PranaSphuta",
  DehaSphuta = "DehaSphuta",
  MrityuSphuta = "MrityuSphuta",
  SookshamaTriSphuta = "SookshamaTriSphuta",
  TithiSphuta = "TithiSphuta",
  YogaSunMoonSphuta = "YogaSunMoonSphuta",
  RahuTithiSphuta = "RahuTithiSphuta",
  KsethraSphuta = "KsethraSphuta",
  BeejaSphuta = "BeejaSphuta",
  TriSphuta = "TriSphuta",
  ChatusSphuta = "ChatusSphuta",
  PanchaSphuta = "PanchaSphuta",
  YogaSphuta = "YogaSphuta",
  AvayogaSphuta = "AvayogaSphuta",
}

export enum AstroConstantsRasis {
  IgnoreRasis = "IgnoreRasis",
  Aries = "Aries",
  Taurus = "Taurus",
  Gemini = "Gemini",
  Cancer = "Cancer",
  Leo = "Leo",
  Virgo = "Virgo",
  Libra = "Libra",
  Scorpio = "Scorpio",
  Saggitarius = "Saggitarius",
  Capricon = "Capricon",
  Aquarius = "Aquarius",
  Pisces = "Pisces",
}

export enum AstroConstantsStars {
  IgnoreStars = "IgnoreStars",
  Aswini = "Aswini",
  Bharani = "Bharani",
  Krittika = "Krittika",
  Rohini = "Rohini",
  Mrigasira = "Mrigasira",
  Arudra = "Arudra",
  Punarvasu = "Punarvasu",
  Pushyami = "Pushyami",
  Aslesha = "Aslesha",
  Makha = "Makha",
  Poorvaphalguni = "Poorvaphalguni",
  Uttaraphalguni = "Uttaraphalguni",
  Hasta = "Hasta",
  Chitra = "Chitra",
  Swati = "Swati",
  Vishaka = "Vishaka",
  Anuradha = "Anuradha",
  Jyeshta = "Jyeshta",
  Moola = "Moola",
  Poorvashada = "Poorvashada",
  Uttarashada = "Uttarashada",
  Sravana = "Sravana",
  Dhanishta = "Dhanishta",
  Satabhisha = "Satabhisha",
  Poorvabhadra = "Poorvabhadra",
  Uttarabadra = "Uttarabadra",
  Revathi = "Revathi",
}

export enum AstroConstantsTithis {
  Ignore_tithi = "Ignore_tithi",
  S_Prathima = "S_Prathima",
  S_Dwitiya = "S_Dwitiya",
  S_Tritiya = "S_Tritiya",
  S_Chaturthi = "S_Chaturthi",
  S_Panchami = "S_Panchami",
  S_Shasti = "S_Shasti",
  S_Saptami = "S_Saptami",
  S_Ashtami = "S_Ashtami",
  S_Navami = "S_Navami",
  S_Dashami = "S_Dashami",
  S_Ekadashi = "S_Ekadashi",
  S_Dwadashi = "S_Dwadashi",
  S_Thrayodashi = "S_Thrayodashi",
  S_Chaturdashi = "S_Chaturdashi",
  Purnima = "Purnima",
  K_Prathama = "K_Prathama",
  K_Dwitiya = "K_Dwitiya",
  K_Tritiya = "K_Tritiya",
  K_Chaturthi = "K_Chaturthi",
  K_Panchami = "K_Panchami",
  K_Shasti = "K_Shasti",
  K_Saptami = "K_Saptami",
  K_Ashtami = "K_Ashtami",
  K_Navami = "K_Navami",
  K_Dashami = "K_Dashami",
  K_Ekadashi = "K_Ekadashi",
  K_Dwadashi = "K_Dwadashi",
  K_Thrayodashi = "K_Thrayodashi",
  K_Chaturdashi = "K_Chaturdashi",
  Amavasya = "Amavasya",
}

export enum AstroMetasNature {
  IgnoreNature = "IgnoreNature",
  Benefic = "Benefic",
  Malefic = "Malefic",
  Choosy = "Choosy",
  Both = "Both",
  Avoid = "Avoid",
}

export enum AstroMetasMetaTags {
  Ignore = "Ignore",
  Gender_male = "Gender_male",
  Gender_female = "Gender_female",
  Gender_neutral = "Gender_neutral",
  Padding_4 = "Padding_4",
  Padding_5 = "Padding_5",
  Padding_6 = "Padding_6",
  Padding_7 = "Padding_7",
  Padding_8 = "Padding_8",
  Padding_9 = "Padding_9",
  Element_air = "Element_air",
  Element_fire = "Element_fire",
  Element_water = "Element_water",
  Element_earth = "Element_earth",
  Quality_fixed = "Quality_fixed",
  Quality_moveable = "Quality_moveable",
  Quality_mutable = "Quality_mutable",
  Quality_cardinal = "Quality_cardinal",
  Padding_18 = "Padding_18",
  Padding_19 = "Padding_19",
  State_common = "State_common",
  State_lame = "State_lame",
  Blind_day = "Blind_day",
  Blind_night = "Blind_night",
  Deaf_morning = "Deaf_morning",
  Deaf_midday = "Deaf_midday",
  Deaf_night = "Deaf_night",
  Padding_27 = "Padding_27",
  Padding_28 = "Padding_28",
  Padding_29 = "Padding_29",
  Tithi_nanda = "Tithi_nanda",
  Tithi_bhadra = "Tithi_bhadra",
  Tithi_jaya = "Tithi_jaya",
  Tithi_riktha = "Tithi_riktha",
  Tithi_poorna = "Tithi_poorna",
  Tridosha_vaatha = "Tridosha_vaatha",
  Tridosha_pitha = "Tridosha_pitha",
  Tridosha_kapha = "Tridosha_kapha",
  Padding_38 = "Padding_38",
  Padding_39 = "Padding_39",
  Bhava_trikona = "Bhava_trikona",
  Bhava_kendra = "Bhava_kendra",
  Bhava_upachaya = "Bhava_upachaya",
  Bhava_dusthana = "Bhava_dusthana",
  Bhava_trishadayas = "Bhava_trishadayas",
  Bhava_marakasthana = "Bhava_marakasthana",
  Padding_46 = "Padding_46",
  Padding_47 = "Padding_47",
  Padding_48 = "Padding_48",
  Padding_49 = "Padding_49",
  Rasi_exalted = "Rasi_exalted",
  Rasi_moolatrikona = "Rasi_moolatrikona",
  Rasi_deblited = "Rasi_deblited",
}

export enum AstroMetasPanchangTags {
  Ignore_panchangTag = "Ignore_panchangTag",
  KrishnaPaksha = "KrishnaPaksha",
  ShuklaPaksha = "ShuklaPaksha",
  Amavasya = "Amavasya",
  Poornima = "Poornima",
  SankataharaChaturti = "SankataharaChaturti",
  Shasti = "Shasti",
  Ekadasi = "Ekadasi",
  AkshayaTritiya = "AkshayaTritiya",
  Krithika = "Krithika",
  Pradosham = "Pradosham",
  Sivaratri = "Sivaratri",
  Mahasivaratri = "Mahasivaratri",
  MarriageMuhurta = "MarriageMuhurta",
}

export enum AstroMetasAstroTypes {
  Ignore_astro_type = "Ignore_astro_type",
  Planet = "Planet",
  Rasi = "Rasi",
  Planet_chaaya = "Planet_chaaya",
  Star = "Star",
  Special_star = "Special_star",
}

export enum StarWeekdayYogaStarWeekdayYogaTypes {
  IgnoreStarYoga = "IgnoreStarYoga",
  Ananda = "Ananda",
  Kaaldanda = "Kaaldanda",
  Dhumra = "Dhumra",
  Dhata = "Dhata",
  Saumya = "Saumya",
  Dhwansha = "Dhwansha",
  Ketu = "Ketu",
  Srivatsa = "Srivatsa",
  Vajra = "Vajra",
  Mudgara = "Mudgara",
  Chhatra = "Chhatra",
  Mitra = "Mitra",
  Manas = "Manas",
  Padma = "Padma",
  Lumba = "Lumba",
  Utpaata = "Utpaata",
  Mrityu = "Mrityu",
  Kaana = "Kaana",
  Siddhi = "Siddhi",
  Subha = "Subha",
  Amrita = "Amrita",
  Mushala = "Mushala",
  Gada = "Gada",
  Matanga = "Matanga",
  Raksha = "Raksha",
  Chara = "Chara",
  Susthira = "Susthira",
  Pravardhamama = "Pravardhamama",
}

export enum StarThithiWeekdayYogaStarThithiWeekdayYogaTypes {
  IgnoreStarThithiWeekdayYoga = "IgnoreStarThithiWeekdayYoga",
  TriPushkaraYoga = "TriPushkaraYoga",
  DwipushkaraYoga = "DwipushkaraYoga",
}

export enum YogaIdentifierYogaTypes {
  IgnoreYogaTypes = "IgnoreYogaTypes",
  StarWeekdayYoga = "StarWeekdayYoga",
  StarThithiWeekdayYoga = "StarThithiWeekdayYoga",
}

export type TobData = {
  d?: number
  m?: number
  y?: number
  hour?: number
  min?: number
  sec?: number
  zone?: string
  dateString?: string
  tobtime?: GoogleProtobufTimestamp.Timestamp
}

export type PobData = {
  city?: string
  state?: string
  country?: string
  lat?: number
  long?: number
  planet?: string
  yuga?: string
}

export type BirthData = {
  tob?: TobData
  pob?: PobData
}

export type GeoLocation = {
  lat?: number
  long?: number
}

export type AstroConstants = {
}

export type AstroMetas = {
}

export type AstroMetadata = {
  uid?: string
  name?: string
  owner?: string
  tenancy?: string
  parentUid?: string
  type?: AstroMetasAstroTypes
  tags?: AstroMetasMetaTags[]
  labels?: {[key: string]: string}
  annotations?: {[key: string]: string}
  description?: string
  createdAt?: GoogleProtobufTimestamp.Timestamp
  updatedAt?: GoogleProtobufTimestamp.Timestamp
}

export type StarWeekdayYoga = {
  name?: string
  yoga?: StarWeekdayYogaStarWeekdayYogaTypes
  nature?: AstroMetasNature
  desc?: string
}

export type StarThithiWeekdayYoga = {
  yoga?: StarThithiWeekdayYogaStarThithiWeekdayYogaTypes
  nature?: AstroMetasNature
  desc?: string
}


type BaseYogaIdentifier = {
  type?: YogaIdentifierYogaTypes
}

export type YogaIdentifier = BaseYogaIdentifier
  & OneOf<{ yogaStarWeekday: StarWeekdayYoga; yogaStarThithiWeekday: StarThithiWeekdayYoga }>
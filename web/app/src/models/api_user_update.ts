import { API_Success_Error } from "./api_error";

export class API_Avatar_Resp extends API_Success_Error {
  avatarUrl?: string;

  static createFrom(source: object | string = {}) {
    return new API_Avatar_Resp(source);
  }

  constructor(source: object | string = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.avatarUrl = source["avatarUrl"];
  }
}

export class API_Bio_Resp extends API_Success_Error {}

export class API_Name_Resp extends API_Success_Error {}

export class API_Type_Resp extends API_Success_Error {}

import { API_Success_Error } from "./api_error";

export class API_Avatar_Resp extends API_Success_Error {
  avatarUrl?: string;

  static createFrom(source: any = {}) {
    return new API_Avatar_Resp(source);
  }

  constructor(source: any = {}) {
    super(source)
    if ('string' === typeof source) source = JSON.parse(source);
    this.error = source["error"];
  }
}

export class API_Bio_Resp extends API_Success_Error { }

export class API_Name_Resp extends API_Success_Error { }
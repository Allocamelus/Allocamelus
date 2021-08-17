export class API_Success_Error {
  success: boolean;
  error?: string;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Success_Error(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.success = source["success"];
    this.error = source["error"];
  }
}

export class API_Success_ID_Error extends API_Success_Error {
  id?: number;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Success_ID_Error(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    super(source)
    if ('string' === typeof source) source = JSON.parse(source);
    this.id = source["id"];
  }
}

export class API_Error {
  error: string;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Error(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.error = source["error"]
  }
}
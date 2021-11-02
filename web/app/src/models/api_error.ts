export class API_Success_Error {
  success: boolean;
  error?: string;

  static createFrom(source: object | string = {}) {
    return new API_Success_Error(source);
  }

  constructor(source: object | string = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.success = source["success"];
    this.error = source["error"];
  }
}

export class API_Success_ID_Error extends API_Success_Error {
  id?: number;

  static createFrom(source: object | string = {}) {
    return new API_Success_ID_Error(source);
  }

  constructor(source: object | string = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.id = source["id"];
  }
}

export class API_Error {
  error: string;

  static createFrom(source: object | string = {}) {
    return new API_Error(source);
  }

  constructor(source: object | string = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.error = source["error"];
  }
}

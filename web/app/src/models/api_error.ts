export class API_Success_Error {
  success: boolean;
  error?: string;

  static createFrom(source: Partial<API_Success_Error> = {}) {
    return new API_Success_Error(source);
  }

  constructor(source: Partial<API_Success_Error> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.success = source["success"];
    this.error = source["error"];
  }
}

export class API_Success_ID_Error extends API_Success_Error {
  id?: number;

  static createFrom(source: Partial<API_Success_ID_Error> = {}) {
    return new API_Success_ID_Error(source);
  }

  constructor(source: Partial<API_Success_ID_Error> = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.id = source["id"];
  }
}

export class API_Error {
  error: string;

  static createFrom(source: Partial<API_Error> = {}) {
    return new API_Error(source);
  }

  constructor(source: Partial<API_Error> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.error = source["error"];
  }
}

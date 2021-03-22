export class API_Error {
  error: string;

  static createFrom(source: any = {}) {
    return new API_Error(source);
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.error = source["error"]
  }
}
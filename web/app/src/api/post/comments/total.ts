import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";

export class API_Total {
  total: number;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Total(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.total = source["total"];
  }
}

export async function get(postId: number | string): Promise<API_Total> {
  return v1.get(`post/${postId}/comments`)
    .then(r => {
      if (r.data.error == undefined) {
        return API_Total.createFrom(r.data)
      } else {
        throw new API_Error(r.data)
      }
    })
}
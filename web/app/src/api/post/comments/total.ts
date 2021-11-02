import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";

export class API_Total {
  total: number;

  static createFrom(source: Partial<API_Total> = {}) {
    return new API_Total(source);
  }

  constructor(source: Partial<API_Total> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.total = source["total"];
  }
}

export function total(postId: number | string): Promise<API_Total> {
  return v1.get(`post/${postId}/comments/total`).then((r) => {
    if (r.data.error == undefined) {
      return API_Total.createFrom(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

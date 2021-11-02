import v1 from "../v1";
import { API_Posts } from "../../models/api_post";
import { API_Error } from "../../models/api_error";

export function get(pageNum: number): Promise<API_Posts> {
  if (pageNum == undefined) {
    pageNum = 1;
  }
  return v1.get("/posts?p=" + pageNum).then((r) => {
    if (r.data.error == undefined) {
      return API_Posts.createFrom(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

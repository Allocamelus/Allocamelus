import v1 from "../v1";
import { API_Post } from "../../models/api_post";
import { API_Error } from "../../models/api_error";

export function get(postId: number | string): Promise<API_Post> {
  return v1.get(`/post/${postId}`).then((r) => {
    if (r.data.error == undefined) {
      return API_Post.createFrom(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

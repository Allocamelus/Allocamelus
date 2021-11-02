import { API_Error } from "../../../models/api_error";
import { notNull } from "../../../models/responses";
import v1 from "../../v1";
import { API_Comments } from "../comments/get";

export function replies(
  postId: number | string,
  commentId: number,
  pageNum = 0
): Promise<API_Comments> {
  return v1
    .get(`post/${postId}/comment/${commentId}/replies?p=${pageNum}`)
    .then((r) => {
      if (notNull(r.data.error)) {
        throw new API_Error(r.data);
      }
      return API_Comments.createFrom(r.data);
    });
}

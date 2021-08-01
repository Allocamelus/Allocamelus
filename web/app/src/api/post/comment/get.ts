import v1 from "../../v1";
import { API_Comment } from '../../../models/api_comment'
import { API_Error } from "../../../models/api_error";

export async function get(postId: number | string, commentId: number | string): Promise<API_Comment> {
  return v1.get(`post/${postId}/comment/${commentId}`)
    .then(r => {
      if (r.data.error == undefined) {
        return API_Comment.createFrom(r.data)
      } else {
        throw new API_Error(r.data)
      }
    })
}
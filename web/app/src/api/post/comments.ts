import v1 from "../v1";
import { API_Comments } from '../../models/api_comment'
import { API_Error } from "../../models/api_error";

export async function get(postId: number | string): Promise<API_Comments> {
  return v1.get(`post/${postId}/comments`)
    .then(r => {
      if (r.data.error == undefined) {
        return API_Comments.createFrom(r.data)
      } else {
        throw new API_Error(r.data)
      }
    })
}
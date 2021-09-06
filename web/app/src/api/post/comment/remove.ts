import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";

export default function remove(postId: number, commentId: number): Promise<API_Error> {
  return v1.delete(`post/${postId}/comment/${commentId}/delete`)
    .then(r => {
      if (r.status !== 204) {
        return API_Error.createFrom(r.data)
      } else {
        return new API_Error()
      }
    })
}

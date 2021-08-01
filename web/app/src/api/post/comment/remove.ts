import v1 from "../../v1";
import { API_Success_Error } from "../../../models/api_error";

export default function remove(postId: number, commentId: number) {
  return v1.post(`post/${postId}/comment/${commentId}/delete`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}

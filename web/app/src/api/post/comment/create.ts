import v1 from "../../v1";
import { API_Success_ID_Error } from "../../../models/api_error";

export default function create(postId: number, replyTo: number, content: string) {
  return v1.post(`post/${postId}/comment`,
    JSON.stringify({
      replyTo: replyTo,
      content: content
    }), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {
      return API_Success_ID_Error.createFrom(r.data)
    })
}

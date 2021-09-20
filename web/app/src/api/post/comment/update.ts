import v1 from "../../v1";
import { API_Success_Error } from "../../../models/api_error";

export default function update(
  postId: number,
  commentId: number,
  content: string
): Promise<API_Success_Error> {
  return v1
    .post(
      `post/${postId}/comment/${commentId}/update`,
      JSON.stringify({
        content: content,
      }),
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    )
    .then((r) => {
      return API_Success_Error.createFrom(r.data);
    });
}

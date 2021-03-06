import v1 from "../v1"
import { API_Posts } from "../../models/api_posts"
import { API_Error } from "../../models/api_error"

export async function posts(userName: string, pageNum: number): Promise<API_Posts> {
  return v1.get(`/user/${userName}/posts?p=${pageNum}`)
    .then(r => {
      if (r.data.error == undefined) {
        return API_Posts.createFrom(r.data)
      } else {
        throw new API_Error(r.data)
      }
    })
}

import v1 from "../v1";
import { Post } from '../../models/post_gen'
export async function get(postId: any) {
  return v1.get("/post/" + postId)
    .then(r => {
      if (r.data.error == undefined) {
        return Post.createFrom(r.data)
      }
    })
}
import v1 from "../v1";
import { Posts } from '../../models/posts'

export async function get(pageNum: number) {
  if (pageNum == undefined) {
    pageNum = 1
  }
  return v1.get("/posts?p=" + pageNum)
    .then(r => {
      if (r.data.error == undefined) {
        return Posts.createFrom(r.data)
      }
    })
}
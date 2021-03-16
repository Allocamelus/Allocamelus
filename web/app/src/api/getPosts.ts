import v1 from "./v1";
import { List } from '../models/post_gen'
export async function getPosts(pageNum: number) {
  if (pageNum == undefined) {
    pageNum = 1
  }
  return v1.get("/posts?p=" + pageNum)
    .then(r => {
      if (r.data.error == undefined) {
        return List.createFrom(r.data)
      }
    })
}
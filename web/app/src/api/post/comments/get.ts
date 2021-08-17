import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import ordered_list from "../../../models/ordered_list";
import { API_Comment } from "../comment"

export class API_Comments extends ordered_list {
  comments: { [key: number]: API_Comment };

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Comments(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    super(source)
    if ('string' === typeof source) source = JSON.parse(source);
    this.comments = source["comments"];
  }

  // Method
  comment(commentId: number): API_Comment {
    return API_Comment.createFrom(this.comments[commentId]);
  }
  appendComment(c: API_Comment) {
    console.log(this.total());
    this.comments[c.id] = c
    this.order[this.total()] = c.id
  }
}

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
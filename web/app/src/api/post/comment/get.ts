import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import { GEN_Comment, GEN_User } from "../../../models/go_structs_gen";

export class API_Comment {
  comment: GEN_Comment;
  user: GEN_User;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Comment(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.comment = GEN_Comment.createFrom(source["comment"]);
    this.user = GEN_User.createFrom(source["user"])
  }
}

export async function get(postId: number | string, commentId: number | string): Promise<API_Comment> {
  return v1.get(`post/${postId}/comment/${commentId}`)
    .then(r => {
      if (r.data.error == undefined) {
        return API_Comment.createFrom(r.data)
      } else {
        throw new API_Error(r.data)
      }
    })
}
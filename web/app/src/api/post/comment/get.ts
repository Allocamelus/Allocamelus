import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import { GEN_User } from "../../../models/go_structs_gen";
import { API_Comment } from "./index";

export class API_Comment_User {
  comment: API_Comment;
  user: GEN_User;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Comment_User(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.comment = API_Comment.createFrom(source["comment"]);
    this.user = GEN_User.createFrom(source["user"])
  }
}

export async function get(postId: number | string, commentId: number | string): Promise<API_Comment_User> {
  return v1.get(`post/${postId}/comment/${commentId}`)
    .then(r => {
      if (r.data.error == undefined) {
        return API_Comment_User.createFrom(r.data)
      } else {
        throw new API_Error(r.data)
      }
    })
}
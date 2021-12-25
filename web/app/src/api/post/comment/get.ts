import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import { User } from "../../../models/user";
import { API_Comment } from "./index";

export class API_Comment_User {
  comment: API_Comment;
  user: User;

  static createFrom(source: Partial<API_Comment_User> = {}) {
    return new API_Comment_User(source);
  }

  constructor(source: Partial<API_Comment_User> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.comment = API_Comment.createFrom(source["comment"]);
    this.user = new User(source["user"]);
  }
}

export function get(
  postId: number | string,
  commentId: number | string
): Promise<API_Comment_User> {
  return v1.get(`post/${postId}/comment/${commentId}`).then((r) => {
    if (r.data.error == undefined) {
      return API_Comment_User.createFrom(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

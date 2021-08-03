import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import ordered_list from "../../../models/ordered_list";
import { GEN_Comment } from "../../../models/go_structs_gen";

export class API_Comments extends ordered_list {
  comments: { [key: number]: GEN_Comment };

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Comments(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    super(source)
    if ('string' === typeof source) source = JSON.parse(source);
    this.comments = source["comments"];
  }

  // Method
  comment(commentId: number): GEN_Comment {
    return GEN_Comment.createFrom(this.comments[commentId]);
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
import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import ordered_list from "../../../models/ordered_list";
import { API_Comment } from "../comment";

export type Ordered_API_Comments = { [key: number]: API_Comment };

export class API_Comments extends ordered_list {
  comments: Ordered_API_Comments;

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Comments(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    super(source);
    if ("string" === typeof source) source = JSON.parse(source);
    this.comments = source["comments"];
  }

  // Methods
  comment(commentId: number): API_Comment {
    if (Object.hasOwnProperty.call(this.comments, commentId)) {
      // Convert comments to API_Comment class if not
      if (!(this.comments[commentId] instanceof API_Comment)) {
        this.comments[commentId] = new API_Comment(this.comments[commentId]);
      }
      return this.comments[commentId];
    }
  }

  appendComment(c: API_Comment) {
    this.comments[c.id] = c;
    this.order[this.total()] = c.id;
  }
}

export async function get(
  postId: number | string,
  pageNum = 0
): Promise<API_Comments> {
  return v1.get(`post/${postId}/comments?p=${pageNum}`).then((r) => {
    if (r.data.error == undefined) {
      return API_Comments.createFrom(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

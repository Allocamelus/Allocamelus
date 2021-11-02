import v1 from "../../v1";
import { API_Error } from "../../../models/api_error";
import ordered_list from "../../../models/ordered_list";
import { API_Comment } from "../comment";

export type Ordered_API_Comments = { [key: number]: API_Comment };

export class API_Comments extends ordered_list {
  comments: Ordered_API_Comments;

  static createFrom(source: object | string = {}) {
    return new API_Comments(source);
  }

  constructor(source: object | string = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
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

  delComment(id: number | string) {
    if (Object.prototype.hasOwnProperty.call(this.comments, id)) {
      delete this.comments[id];
      // Remove comment id from order
      let removed = false;
      for (let k in this.order) {
        if (this.order[k] == id) {
          delete this.order[k];
          removed = true;
        } else if (removed) {
          // Shift order by one after removed id
          this.order[`${Number(k).valueOf() - 1}`] = this.order[k];
        }
      }
    }
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

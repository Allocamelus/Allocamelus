import { GEN_User, GEN_Comment } from './go_structs_gen'
import ordered_list from './ordered_list';


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
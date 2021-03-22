import { User } from "./user_gen";
import { Post } from './post_gen'

export class API_Post {
  post: Post;
  user: User;

  static createFrom(source: any = {}) {
    return new API_Post(source);
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.post = Post.createFrom(source["post"]);
    this.user = User.createFrom(source["user"])
  }
}
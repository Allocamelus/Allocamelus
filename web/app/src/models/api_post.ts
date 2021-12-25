import { User } from "./user";
import { Post } from "./post";
import ordered_list from "./ordered_list";

export class API_Post {
  post: Post;
  user: User;

  static createFrom(source: Partial<API_Post> = {}) {
    return new API_Post(source);
  }

  constructor(source: Partial<API_Post> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.post = new Post(source["post"]);
    this.user = new User(source["user"]);
  }
}
export class API_Posts extends ordered_list {
  posts: { [key: number]: Post };

  static createFrom(source: Partial<API_Posts> = {}) {
    return new API_Posts(source);
  }

  constructor(source: Partial<API_Posts> = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.posts = source["posts"] || [];
  }

  // Method
  post(postId: number): Post {
    return new Post(this.posts[postId]);
  }
}

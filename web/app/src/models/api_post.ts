import { GEN_User, GEN_Post } from "./go_structs_gen";
import ordered_list from "./ordered_list";

export class API_Post {
  post: GEN_Post;
  user: GEN_User;

  static createFrom(source: Partial<API_Post> = {}) {
    return new API_Post(source);
  }

  constructor(source: Partial<API_Post> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.post = GEN_Post.createFrom(source["post"]);
    this.user = GEN_User.createFrom(source["user"]);
  }
}
export class API_Posts extends ordered_list {
  posts: { [key: number]: GEN_Post };

  static createFrom(source: Partial<API_Posts> = {}) {
    return new API_Posts(source);
  }

  constructor(source: Partial<API_Posts> = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.posts = source["posts"];
  }

  // Method
  post(postId: number): GEN_Post {
    return GEN_Post.createFrom(this.posts[postId]);
  }
}

import { GEN_User, GEN_Post } from "./go_structs_gen";

export class API_Posts {
  posts: { [key: number]: GEN_User };
  users: { [key: number]: GEN_Post };
  order: { [key: number]: number };


  static createFrom(source: any = {}) {
    return new API_Posts(source);
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.posts = source["posts"];
    this.users = source["users"];
    this.order = source["order"]
  }

  // Method
  user(userId: number) {
    return GEN_User.createFrom(this.users[userId]);
  }
  post(postId: number) {
    return GEN_Post.createFrom(this.posts[postId]);
  }
}
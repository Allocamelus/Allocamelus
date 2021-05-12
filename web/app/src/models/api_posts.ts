import { GEN_User, GEN_Post } from "./go_structs_gen";

export class API_Posts {
  posts: { [key: number]: GEN_Post };
  users: { [key: number]: GEN_User };
  order: { [key: number]: number };


  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Posts(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
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
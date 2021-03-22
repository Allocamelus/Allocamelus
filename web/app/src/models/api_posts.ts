import { User } from "./user_gen";
import { Post } from './post_gen'

export class API_Posts {
  posts: { [key: number]: Post };
  users: { [key: number]: User };
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
    return this.users[userId];
  }
  post(postId: number) {
    return this.posts[postId];
  }
}
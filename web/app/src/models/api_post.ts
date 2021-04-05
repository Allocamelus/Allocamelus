import { GEN_User, GEN_Post } from './go_structs_gen'

export class API_Post {
  post: GEN_Post;
  user: GEN_User;

  static createFrom(source: any = {}) {
    return new API_Post(source);
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.post = GEN_Post.createFrom(source["post"]);
    this.user = GEN_User.createFrom(source["user"])
  }
}
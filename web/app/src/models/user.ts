export const Unverified = 0;
export const Private = 1;
export const Public = 2;

export class FollowStruct {
  following: boolean;
  requested: boolean;

  constructor(source: Partial<FollowStruct> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.following = source["following"] || false;
    this.requested = source["requested"] || false;
  }
}
export class User {
  id: number;
  userName: string;
  name: string;
  email?: string;
  avatar: string;
  bio?: string;
  selfFollow: FollowStruct;
  userFollow: FollowStruct;
  followers: number;
  type: number;
  created?: number;

  constructor(source: Partial<User> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.id = source["id"] || 0;
    this.userName = source["userName"] || "";
    this.name = source["name"] || "";
    this.email = source["email"];
    this.avatar = source["avatar"] || "";
    this.bio = source["bio"];
    this.selfFollow = new FollowStruct(source["selfFollow"] || {});
    this.userFollow = new FollowStruct(source["userFollow"] || {});

    this.followers = source["followers"] || 0;
    this.type = source["type"] || 0;
    this.created = source["created"];
  }
}

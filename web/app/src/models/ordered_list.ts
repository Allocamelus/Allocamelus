import { User } from "./user";

export class user_list {
  users: { [key: number]: User };

  static createFrom(source: Partial<user_list> = {}) {
    return new user_list(source);
  }

  constructor(source: Partial<user_list> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.users = source["users"] || [];

    if (source["users"] != undefined) {
      for (const [key, value] of Object.entries(this.users)) {
        this.users[Number(key).valueOf()] = new User(value);
      }
    }
  }

  // Methods
  user(userId: number): User {
    if (Object.hasOwnProperty.call(this.users, userId)) {
      // Convert users to User class if not
      if (!(this.users[userId] instanceof User)) {
        this.users[userId] = new User(this.users[userId]);
      }
      return this.users[userId];
    }
  }
  appendUser(u: User): void {
    if (this.users === null) {
      this.users = {};
    }
    if (!(u.id in this.users)) {
      this.users[u.id] = u;
    }
  }
}
export default class ordered_list extends user_list {
  order: { [key: number]: number };

  static createFrom(source: Partial<ordered_list> = {}) {
    return new ordered_list(source);
  }

  constructor(source: Partial<ordered_list> = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.order = source["order"] || [];
  }

  // Method
  total(): number {
    if (this.order == undefined || this.order === null) {
      return 0;
    }
    return Object.keys(this.order).length;
  }
}

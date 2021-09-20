import { GEN_User } from "./go_structs_gen";

export class user_list {
  users: { [key: number]: GEN_User };

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new user_list(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ("string" === typeof source) source = JSON.parse(source);

    this.users = source["users"];

    if (this.users != undefined) {
      for (const [key, value] of Object.entries(this.users)) {
        this.users[key] = GEN_User.createFrom(value);
      }
    }
  }

  // Methods
  user(userId: number): GEN_User {
    if (Object.hasOwnProperty.call(this.users, userId)) {
      // Convert users to GEN_User class if not
      if (!(this.users[userId] instanceof GEN_User)) {
        this.users[userId] = new GEN_User(this.users[userId]);
      }
      return this.users[userId];
    }
  }
  appendUser(u: GEN_User) {
    if (!(u.id in this.users)) {
      this.users[u.id] = u;
    }
  }
}
export default class ordered_list extends user_list {
  order: { [key: number]: number };

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new ordered_list(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    super(source);
    if ("string" === typeof source) source = JSON.parse(source);
    this.order = source["order"];
  }

  // Method
  total(): number {
    if (this.order == undefined || this.order === null) {
      return 0;
    }
    return Object.keys(this.order).length;
  }
}

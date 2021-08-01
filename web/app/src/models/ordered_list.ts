import { GEN_User } from './go_structs_gen'

export default class ordered_list {
  users: { [key: number]: GEN_User };
  order: { [key: number]: number };


  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new ordered_list(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.users = source["users"];
    this.order = source["order"]
  }

  // Method
  user(userId: number): GEN_User {
    return GEN_User.createFrom(this.users[userId]);
  }
  total(): number {
    if (this.order == undefined || this.order === null) {
      return 0
    }
    return Object.keys(this.order).length
  }
}
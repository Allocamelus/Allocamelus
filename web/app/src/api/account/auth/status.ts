import v1 from "../../v1";
import { GEN_User } from "../../../models/go_structs_gen";


export class Status {
  loggedIn: boolean;
  user?: GEN_User;


  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new Status(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.loggedIn = source["loggedIn"];
    this.user = source["user"];
  }
}

export async function status() {
  return v1.get("account/auth/status").then(r => {
    return Status.createFrom(r.data)
  })
}
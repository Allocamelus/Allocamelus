import v1 from "../../v1";
import { GEN_User } from "../../../models/go_structs_gen";


export class Status {
  loggedIn: Boolean;
  user?: GEN_User;


  static createFrom(source: any = {}) {
    return new Status(source);
  }

  constructor(source: any = {}) {
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
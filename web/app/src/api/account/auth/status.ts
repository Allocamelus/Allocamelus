import v1 from "../../v1";
import { GEN_User } from "../../../models/go_structs_gen";

export class Status {
  loggedIn: boolean;
  user?: GEN_User;

  static createFrom(source: Partial<Status> = {}) {
    return new Status(source);
  }

  constructor(source: Partial<Status> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.loggedIn = source["loggedIn"];
    this.user = source["user"];
  }
}

export function status(): Promise<Status> {
  return v1.get("account/auth/status").then((r) => {
    return Status.createFrom(r.data);
  });
}

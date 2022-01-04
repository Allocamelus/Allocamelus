import v1 from "../../v1";
import { User } from "../../../models/user";

export class Status {
  loggedIn: boolean;
  user?: User;

  static createFrom(source: Partial<Status> = {}) {
    return new Status(source);
  }

  constructor(source: Partial<Status> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.loggedIn = source["loggedIn"] || false;
    this.user = source["user"];
  }
}

export function status(): Promise<Status> {
  return v1.get("account/auth/status").then((r) => {
    return Status.createFrom(r.data);
  });
}

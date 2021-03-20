import v1 from "../../v1";
import { Session } from "../../../models/user_gen";

export async function status() {
  return v1.get("account/auth/status").then(r => {
    return Session.createFrom(r.data)
  })
}
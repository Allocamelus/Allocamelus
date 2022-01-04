import v1 from "../v1";
import { User } from "../../models/user";
import { API_Error } from "../../models/api_error";

export function get(userName: string): Promise<User> {
  return v1.get("/user/" + userName).then((r) => {
    if (r.data.error == undefined) {
      return new User(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

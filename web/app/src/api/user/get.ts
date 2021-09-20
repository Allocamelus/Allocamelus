import v1 from "../v1";
import { GEN_User } from "../../models/go_structs_gen";
import { API_Error } from "../../models/api_error";

export async function get(userName: string): Promise<GEN_User> {
  return v1.get("/user/" + userName).then((r) => {
    if (r.data.error == undefined) {
      return GEN_User.createFrom(r.data);
    } else {
      throw new API_Error(r.data);
    }
  });
}

import v1 from "../../v1";
import { API_Name_Resp } from "../../../models/api_user_update";

export async function name(userName: string, name: string) {
  return v1.post(`/user/${userName}/update/name`,
    JSON.stringify({ name: name }), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {
      return API_Name_Resp.createFrom(r.data)
    })
}
import v1 from "../../v1";
import { API_Bio_Resp } from "../../../models/api_user_update";

export async function bio(userName: string, bio: string): Promise<API_Bio_Resp> {
  return v1.post(`/user/${userName}/update/bio`,
    JSON.stringify({ bio: bio }), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {
      return API_Bio_Resp.createFrom(r.data)
    })
}
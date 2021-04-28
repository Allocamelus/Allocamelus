import v1 from "../../v1";
import { API_Success_Error } from "../../../models/api_error";

export async function remove(userName: string) {
  return v1.delete(`/user/${userName}/follow`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}
export async function post(userName: string) {
  return v1.post(`/user/${userName}/follow`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}
import v1 from "../../v1";
import { API_Type_Resp } from "../../../models/api_user_update";

export const TYPE_PRIVATE = 1;
export const TYPE_PUBLIC = 2;

export function type(userName: string, type: number): Promise<API_Type_Resp> {
  return v1
    .post(`/user/${userName}/update/type`, JSON.stringify({ type: type }), {
      headers: {
        "Content-Type": "application/json",
      },
    })
    .then((r) => {
      return API_Type_Resp.createFrom(r.data);
    });
}

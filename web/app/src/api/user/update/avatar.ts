import v1 from "../../v1";
import { API_Avatar_Resp } from "../../../models/api_user_update";
import { AxiosResponse } from "axios";

export function avatar(userName: string, file: File): Promise<API_Avatar_Resp> {
  const formData = new FormData();
  formData.append("avatar", file);
  return v1
    .post(`/user/${userName}/update/avatar`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
    .then((r) => {
      return API_Avatar_Resp.createFrom(r.data);
    });
}

export function removeAvatar(
  userName: string
): Promise<AxiosResponse<any, any>> {
  return v1.delete(`/user/${userName}/update/avatar`);
}

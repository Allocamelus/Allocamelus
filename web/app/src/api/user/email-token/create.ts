import v1 from "../../v1";
import { API_Success_Error } from "../../../models/api_error";

export function create(
  email: string,
  captcha: string
): Promise<API_Success_Error> {
  return v1
    .post(
      "/user/email-token",
      JSON.stringify({ email: email, captcha: captcha }),
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    )
    .then((r) => {
      return API_Success_Error.createFrom(r.data);
    });
}

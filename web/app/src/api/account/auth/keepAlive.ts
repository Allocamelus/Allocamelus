import { AxiosResponse } from "axios";
import v1 from "../../v1";

export function keepAlive(): Promise<AxiosResponse<any, any>> {
  return v1.post("account/auth/keep-alive");
}

import { AxiosResponse } from "axios";
import v1 from "../../v1";

export async function keepAlive(): Promise<AxiosResponse<any>> {
  return v1.post("account/auth/keep-alive")
}
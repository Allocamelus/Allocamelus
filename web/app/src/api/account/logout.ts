import { AxiosResponse } from "axios";
import v1 from "../v1";

export function logout(): Promise<AxiosResponse<any, any>> {
  return v1.delete("account/logout");
}

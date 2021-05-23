import { AxiosResponse } from "axios";
import v1 from "../v1";

export async function logout(): Promise<AxiosResponse<any>> {
  return v1.delete("account/logout")
}
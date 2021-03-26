import v1 from "../../v1";
import { API_MetaCaptchaSiteKeys } from "../../../models/api_meta";

export async function siteKeys() {
  return v1.get("meta/captcha/site-keys").then(r => {
    return API_MetaCaptchaSiteKeys.createFrom(r.data)
  })
}
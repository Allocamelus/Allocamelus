import v1 from "../v1";
import { API_AuthRequest, API_AuthResp, API_AuthA10Token } from "../../models/api_account_gen";

const a10 = "allocamelus"

export async function auth(request: API_AuthRequest) {
  return v1.post("account/auth", JSON.stringify(request))
    .then(r => {
      return API_AuthResp.createFrom(r)
    })
}

export async function authA10(token: API_AuthA10Token) {
  return auth(API_AuthRequest.createFrom({
    with: a10,
    token: JSON.stringify(token)
  }))
}
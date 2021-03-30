import v1 from "../v1";
import { API_AuthRequest, API_AuthResp, API_AuthA10Token } from "../../models/api_account_gen";

const a9s = "allocamelus"

export async function auth(request: API_AuthRequest) {
  return v1.post("account/auth",
    JSON.stringify(request), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {      
      return API_AuthResp.createFrom(r.data)
    })
}

export async function authA10(token: API_AuthA10Token) {
  return auth(API_AuthRequest.createFrom({
    with: a9s,
    token: JSON.stringify(token)
  }))
}
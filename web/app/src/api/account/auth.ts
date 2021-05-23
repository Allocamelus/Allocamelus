import v1 from "../v1";
import { GEN_AuthRequest, GEN_AuthResp, GEN_AuthA10Token } from "../../models/go_structs_gen";

const a9s = "allocamelus"

export async function auth(request: GEN_AuthRequest): Promise<GEN_AuthResp> {
  return v1.post("account/auth",
    JSON.stringify(request), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {      
      return GEN_AuthResp.createFrom(r.data)
    })
}

export async function authA10(token: GEN_AuthA10Token): Promise<GEN_AuthResp> {
  return auth(GEN_AuthRequest.createFrom({
    with: a9s,
    token: JSON.stringify(token)
  }))
}
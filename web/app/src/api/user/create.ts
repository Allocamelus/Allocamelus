import v1 from "../v1"
import { API_CreateA10Token, API_CreateRequest, API_CreateResp } from "../../models/GEN_User_gen";


const a9s = "allocamelus"

export async function create(request: API_CreateRequest) {
  return v1.post("user",
    JSON.stringify(request), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {      
      return API_CreateResp.createFrom(r.data)
    })
}

export async function createA9s(token: API_CreateA10Token) {
  return create(API_CreateRequest.createFrom({
    with: a9s,
    token: JSON.stringify(token)
  }))
}
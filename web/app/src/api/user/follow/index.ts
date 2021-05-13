import v1 from "../../v1";
import { API_Success_Error } from "../../../models/api_error";
import { GEN_User } from "../../../models/go_structs_gen";

export async function remove(userName: string) {
  return v1.delete(`/user/${userName}/follow`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}
export async function post(userName: string) {
  return v1.post(`/user/${userName}/follow`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}

export async function decline(userName: string) {
  return v1.delete(`/user/${userName}/follow/decline`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}
export async function accept(userName: string) {
  return v1.post(`/user/${userName}/follow/accept`)
    .then(r => {
      return API_Success_Error.createFrom(r.data)
    })
}

export class API_Requests {
  requests: { [key: number]: number };
  users: { [key: number]: GEN_User };


  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Requests(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.requests = source["requests"];
    this.users = source["users"];
  }

  // Method
  user(userId: number) {
    return GEN_User.createFrom(this.users[userId]);
  }
}

export async function requests() {
  return v1.get(`/user/follow/requests`)
    .then(r => {
      return API_Requests.createFrom(r.data)
    })
}
import v1 from "../../v1";
import { API_Success_Error } from "../../../models/api_error";
import { User } from "../../../models/user";

export function remove(userName: string): Promise<API_Success_Error> {
  return v1.delete(`/user/${userName}/follow`).then((r) => {
    return API_Success_Error.createFrom(r.data);
  });
}
export function post(userName: string): Promise<API_Success_Error> {
  return v1.post(`/user/${userName}/follow`).then((r) => {
    return API_Success_Error.createFrom(r.data);
  });
}

export function decline(userName: string): Promise<API_Success_Error> {
  return v1.delete(`/user/${userName}/follow/decline`).then((r) => {
    return API_Success_Error.createFrom(r.data);
  });
}
export function accept(userName: string): Promise<API_Success_Error> {
  return v1.post(`/user/${userName}/follow/accept`).then((r) => {
    return API_Success_Error.createFrom(r.data);
  });
}

export class API_Requests {
  requests: { [key: number]: number };
  users: { [key: number]: User };

  static createFrom(source: Partial<API_Requests> = {}) {
    return new API_Requests(source);
  }

  constructor(source: Partial<API_Requests> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.requests = source["requests"] || [];
    this.users = source["users"] || [];
  }

  // Method
  user(userId: number): User {
    return new User(this.users[userId]);
  }
}

export function requests(): Promise<API_Requests> {
  return v1.get(`/user/follow/requests`).then((r) => {
    return API_Requests.createFrom(r.data);
  });
}

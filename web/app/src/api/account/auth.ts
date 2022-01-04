import { API_Success_Error } from "../../models/api_error";
import { User } from "../../models/user";
import v1 from "../v1";

export class AuthResponse extends API_Success_Error {
  privateArmored?: string;
  user?: User;
  captcha?: string;

  constructor(source: Partial<AuthResponse> = {}) {
    super(source);
    if ("string" === typeof source) source = JSON.parse(source);
    this.privateArmored = source["privateArmored"];
    this.user = new User(source["user"]);
    this.captcha = source["captcha"];
  }
}

export interface AuthRequest {
  userName: string;
  authKey: string;
  remember: boolean;
  captcha?: string;
}

export function auth(request: AuthRequest): Promise<AuthResponse> {
  return v1
    .post("account/auth", JSON.stringify(request), {
      headers: {
        "Content-Type": "application/json",
      },
    })
    .then((r) => {
      return new AuthResponse(r.data);
    });
}

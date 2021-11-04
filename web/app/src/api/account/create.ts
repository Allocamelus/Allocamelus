import v1 from "../v1";

export class CreateResp {
  success: boolean;
  errors?: any;

  constructor(source: Partial<CreateResp> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.success = source["success"];
    this.errors = source["errors"];
  }
}

export class CreateRequest {
  userName: string;
  email: string;
  password: Password;
  captcha: string;

  constructor(source: Partial<CreateRequest> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.userName = source["userName"];
    this.email = source["email"];
    this.password = new Password(source["password"]);
    this.captcha = source["captcha"];
  }
}
export class Password {
  salt: string;
  hash: string;

  constructor(source: Partial<Password> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.salt = source["salt"];
    this.hash = source["hash"];
  }
}

export function create(request: CreateRequest): Promise<CreateResp> {
  return v1
    .post("account", JSON.stringify(request), {
      headers: {
        "Content-Type": "application/json",
      },
    })
    .then((r) => {
      return new CreateResp(r.data);
    });
}

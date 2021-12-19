import v1 from "../v1";

export class SaltResponse {
  salt: string;
  error?: string;

  constructor(source: Partial<SaltResponse> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    if (source["salt"] == undefined) source["salt"] = "";
    this.salt = source["salt"];
    this.error = source["error"];
  }
}

export function salt(userName: string): Promise<SaltResponse> {
  return v1
    .post(
      "account/salt",
      JSON.stringify({
        userName: userName,
      }),
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    )
    .then((r) => {
      return new SaltResponse(r.data);
    });
}

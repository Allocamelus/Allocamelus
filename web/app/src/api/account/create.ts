import v1 from "../v1";

export class CreateResp {
  success: boolean;
  errors?: errors | Array<string>;

  constructor(source: Partial<CreateResp> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.success = source["success"] || false;
    if (source["errors"] != undefined) {
      this.errors = Array.isArray(source["errors"])
        ? source["errors"]
        : new errors(source["errors"]);
    }
  }
}

class errors {
  userName: string;
  email: string;

  constructor(source: Partial<errors> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.userName = source["userName"] || "";
    this.email = source["email"] || "";
  }
}

export interface CreateRequest {
  userName: string;
  email: string;
  auth: authParts;
  key: Key;
  captcha: string;
}
export interface authParts {
  salt: string;
  hash: string;
}

export interface Key {
  // PublicArmored armored PGP public key
  publicArmored: string;
  // PrivateArmored armored PGP private key encrypted with passphrase
  privateArmored: string;
  // RecoveryHash hash of recovery key
  recoveryHash: string;
  // RecoveryArmored PGP private key encrypted with recovery key
  recoveryArmored: string;
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

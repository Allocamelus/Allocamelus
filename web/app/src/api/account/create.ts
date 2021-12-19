import { API_Success_Error } from "../../models/api_error";
import v1 from "../v1";

export class CreateResp {
  success: boolean;
  errors?: any;

  constructor(source: Partial<CreateResp> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.success = source["success"] || false;
    this.errors = source["errors"];
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
  return new Promise(async (resolve) => {
    resolve(
      new CreateResp(
        await v1.post("account", JSON.stringify(request), {
          headers: {
            "Content-Type": "application/json",
          },
        })
      )
    );
    return;
  });
}

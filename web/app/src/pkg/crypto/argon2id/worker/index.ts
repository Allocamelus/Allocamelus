import { expose } from "comlink";
import { argon2id as argon2idHash } from "hash-wasm";
import { argon2idCost } from "../argon2id";
import { IDataType } from "hash-wasm/dist/lib/util";

export interface Argon2id {
  hash: typeof hash;
}

export interface hashReturn {
  encodedHash: string;
  err: any;
}

const argon2id: Argon2id = {
  hash,
};

function hash(
  password: IDataType,
  salt: IDataType,
  cost: argon2idCost
): Promise<hashReturn> {
  return new Promise(async (resolve) => {
    let out: hashReturn = { encodedHash: "", err: null };
    try {
      out.encodedHash = await argon2idHash({
        password: password,
        salt: salt,
        iterations: cost.time,
        hashLength: cost.keyLen,
        memorySize: cost.memory,
        parallelism: cost.threads,
        outputType: "encoded",
      });
    } catch (error) {
      out.err = error;
    }

    resolve(out);
    return;
  });
}

expose(argon2id);

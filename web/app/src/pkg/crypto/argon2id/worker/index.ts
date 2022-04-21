import { expose } from "comlink";
import { argon2id as argon2idHash } from "hash-wasm";
import { argon2idCost } from "../argon2id";
import { IDataType } from "hash-wasm/dist/lib/util";
import { NullError } from "@/models/Error";

export interface Argon2idWorker {
  hash: typeof hash;
}

interface hashReturn {
  encodedHash: string;
  err: NullError<any>;
}

async function hash(
  password: IDataType,
  salt: IDataType,
  cost: argon2idCost
): Promise<hashReturn> {
  const out: hashReturn = { encodedHash: "", err: null };
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
  return out;
}

const argon2id: Argon2idWorker = {
  hash,
};

expose(argon2id);

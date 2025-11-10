import argon2idWorker from "./worker?worker&inline";
import { wrap } from "comlink";
import { Buffer } from "buffer";

import { argon2idCost, argon2idEncoded } from "./argon2id";
import { IDataType } from "hash-wasm/dist/lib/util";
import { Argon2idWorker } from "./worker";
import { NullError } from "@/models/Error";

const argon2id = wrap<Argon2idWorker>(new argon2idWorker());

/**
 * Parse encoded argon2id hash
 *
 * $argon2id$v={version}$m={memory},t={time},p={threads}${base64(salt)}${base64(key)}
 *
 * @param encodedHash Key is optional
 */
export function parse(encodedHash: string): argon2idEncoded {
  const slice = encodedHash.split("$");
  const costSlice = slice[3].replace(/[mtp=\s]/g, "").split(",");

  const encoded = new argon2idEncoded({
    cost: new argon2idCost({
      memory: parseInt(costSlice[0]),
      time: parseInt(costSlice[1]),
      threads: parseInt(costSlice[2]),
      saltLen: Buffer.from(slice[4], "base64").length,
    }),
    encoded: "",
    salt: slice[4],
    version: parseInt(slice[2].replace(/[v=\s]/g, "")),
    key: "",
  });

  if (slice.length > 5) {
    encoded.key = slice[5];
    slice.pop();
    encoded.cost.keyLen = Buffer.from(encoded.key, "base64").length;
  }
  encoded.encoded = slice.join("$");
  return encoded;
}
interface hashReturn {
  hash: argon2idEncoded;
  err: NullError<any>;
}

/**
 * Hash password with argon2id
 * user provides salt
 *
 * @param password
 * @param salt
 * @param cost
 */
export async function hashSalt(
  password: IDataType,
  salt: IDataType,
  cost: argon2idCost
): Promise<hashReturn> {
  // Set defaults for empty cost values
  cost = new argon2idCost(cost);
  cost.FillEmpty();

  // Normalize password
  if (typeof password === "string") {
    password = password.normalize();
  }

  const { encodedHash, err } = await argon2id.hash(password, salt, cost);
  if (err !== null) {
    return {
      hash: new argon2idEncoded(),
      err,
    };
  }
  return { hash: parse(encodedHash), err };
}

/**
 * Hash password with argon2id
 *
 * @param password
 * @param cost
 */
export function hash(
  password: IDataType,
  cost: argon2idCost
): Promise<hashReturn> {
  // Set defaults for empty cost values
  cost = new argon2idCost(cost);
  cost.FillEmpty();

  // Generate salt
  const salt = new Uint8Array(cost.saltLen);
  window.crypto.getRandomValues(salt);

  return hashSalt(password, salt, cost);
}

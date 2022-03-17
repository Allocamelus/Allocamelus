import argon2idWorker from "./worker?worker&inline";
import { wrap } from "comlink";
import { Buffer } from "buffer";

import { argon2idCost, argon2idEncoded } from "./argon2id";
import { IDataType } from "hash-wasm/dist/lib/util";
import { Argon2idWorker } from "./worker";

const argon2id = wrap<Argon2idWorker>(new argon2idWorker());

/**
 * Parse encoded argon2id hash
 *
 * $argon2id$v={version}$m={memory},t={time},p={threads}${base64(salt)}${base64(key)}
 *
 * @param encodedHash Key is optional
 */
export function parse(encodedHash: string): argon2idEncoded {
  let slice = encodedHash.split("$");
  let costSlice = slice[3].replace(/[mtp=\s]/g, "").split(",");

  let encoded: argon2idEncoded = {
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
  };

  if (slice.length > 5) {
    encoded.key = slice[5];
    slice.pop();
    encoded.cost.keyLen = Buffer.from(encoded.key, "base64").length;
  }
  encoded.encoded = slice.join("$");
  return encoded;
}

/**
 * Hash password with argon2id
 * user provides salt
 *
 * @param password
 * @param salt
 * @param cost
 */
export function hashSalt(
  password: IDataType,
  salt: IDataType,
  cost: argon2idCost
): Promise<argon2idEncoded> {
  return new Promise(async (resolve, reject) => {
    // Set defaults for empty cost values
    cost = new argon2idCost(cost);
    cost.FillEmpty();

    // Normalize password
    if (typeof password === "string") {
      password = password.normalize();
    }

    let { encodedHash, err } = await argon2id.hash(password, salt, cost);
    if (err !== null) {
      reject(err);
      return;
    }

    resolve(parse(encodedHash));
    return;
  });
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
): Promise<argon2idEncoded> {
  // Set defaults for empty cost values
  cost = new argon2idCost(cost);
  cost.FillEmpty();

  // Generate salt
  const salt = new Uint8Array(cost.saltLen);
  window.crypto.getRandomValues(salt);

  return hashSalt(password, salt, cost);
}

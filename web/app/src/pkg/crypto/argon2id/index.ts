import { argon2id } from "hash-wasm";
import { Buffer } from "buffer";

import { argon2idCost, argon2idEncoded } from "./argon2id";
import { IDataType } from "hash-wasm/dist/lib/util";

/**
 * Parse encoded argon2id hash
 *
 * $argon2id$v={version}$m={memory},t={time},p={threads}${base64(salt)}${base64(key)}
 *
 * @param encodedHash Key is optional
 */
export function parse(encodedHash: string): Promise<argon2idEncoded> {
  return new Promise((resolve) => {
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
    resolve(encoded);
    return;
  });
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
  return new Promise((resolve, reject) => {
    // Set defaults for empty cost values
    cost = new argon2idCost(cost);
    cost.FillEmpty();

    // Normalize password
    if (typeof password === "string") {
      password = password.normalize();
    }

    argon2id({
      password: password,
      salt: salt,
      iterations: cost.time,
      hashLength: cost.keyLen,
      memorySize: cost.memory,
      parallelism: cost.threads,
      outputType: "encoded",
    })
      .then(async (encodedHash) => {
        resolve(await parse(encodedHash));
        return;
      })
      .catch((e) => {
        reject(e);
        return;
      });
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

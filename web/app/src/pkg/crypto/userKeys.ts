import { Buffer } from "buffer";
import { exportKey } from "./aesgcm-tools";
import { hash as hashA2, hashSalt, parse } from "./argon2id";
import { argon2idCost, argon2idEncoded } from "./argon2id/argon2id";
import { create } from "./recoveryKey";
import { blake2bB64 } from "./blake2b";
import { genKey, decryptKey, encryptKey } from "./pgp";
import { NullError } from "@/models/Error";

/**
 * userKey
 *
 * @var {number} created unix time
 * @var {string} keyAuthHash key hashed
 * @var {string} keySaltEncoded
 * @var {string} publicArmored
 * @var {string} privateArmored encrypted
 * @var {string} recoveryArmored encrypted
 * @var {string} recoveryHash
 */
export interface userKey {
  keyAuthHash: string;
  keySaltEncoded: string;
  publicArmored: string;
  privateArmored: string;
  recoveryArmored: string;
  recoveryHash: string;
}
export class userKey {
  keyAuthHash: string;
  keySaltEncoded: string;
  publicArmored: string;
  privateArmored: string;
  recoveryArmored: string;
  recoveryHash: string;
  constructor(source: Partial<userKey> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.keyAuthHash = source["keyAuthHash"] || "";
    this.keySaltEncoded = source["keySaltEncoded"] || "";
    this.publicArmored = source["publicArmored"] || "";
    this.privateArmored = source["privateArmored"] || "";
    this.recoveryArmored = source["recoveryArmored"] || "";
    this.recoveryHash = source["recoveryHash"] || "";
  }
}

const b2bAuthKey = "B2b User Authentication Key-Info";
const b2bPgpKey = "B2b Pretty Good Privacy Key-Info";

async function genHashKeys(password: string): Promise<{
  saltEncoded: string;
  authKey: string;
  pgpPassphrase: string;
  err: NullError<any>;
}> {
  // Hash password
  const { hash, err } = await hashA2(
    password,
    new argon2idCost({
      keyLen: 32,
      saltLen: 32,
      memory: 128 * 1024,
      threads: 2,
      time: 3,
    })
  );
  if (err !== null) {
    return {
      saltEncoded: "",
      authKey: "",
      pgpPassphrase: "",
      err,
    };
  }

  const { authKey, pgpPassphrase } = await deriveKeys(hash);

  return {
    saltEncoded: hash.encoded,
    authKey: authKey,
    pgpPassphrase: pgpPassphrase,
    err: null,
  };
}

/**
 * genKeys Generates a new pgp key, recovery key, and password hash
 *
 * @param {string} username
 * @param {string} password un-hashed password
 * @return {Promise<userKey>} msg pretty readable error
 */
export async function genKeys(
  username: string,
  password: string
): Promise<{ userKey: userKey; recoveryKey: string; err: NullError<any> }> {
  const recoveryKey = create();

  const recoveryKeyArray = exportKey((await recoveryKey).key);
  const recoveryHash = blake2bB64(await recoveryKeyArray, 512);

  const keys = await genHashKeys(password);
  if (keys.err !== null) {
    return {
      userKey: new userKey(),
      recoveryKey: "",
      err: keys.err,
    };
  }

  const pgpKey = await genKey(username, keys.pgpPassphrase);

  const privateKey = decryptKey(pgpKey.armoredPrivate, keys.pgpPassphrase);
  const recoveryArmored = encryptKey(
    await privateKey,
    username,
    Buffer.from(await recoveryKeyArray).toString("base64")
  );

  return {
    userKey: new userKey({
      keyAuthHash: keys.authKey,
      keySaltEncoded: keys.saltEncoded,
      publicArmored: pgpKey.armoredPublic,
      privateArmored: pgpKey.armoredPrivate,
      recoveryArmored: await recoveryArmored,
      recoveryHash: await recoveryHash,
    }),
    recoveryKey: (await recoveryKey).encoded,
    err: null,
  };
}

async function deriveKeys(
  keyEncoded: argon2idEncoded
): Promise<{ authKey: string; pgpPassphrase: string }> {
  const key = Buffer.from(keyEncoded.key, "base64");

  const authKey = blake2bB64(key, 512, b2bAuthKey);
  const pgpPassphrase = blake2bB64(key, 512, b2bPgpKey);

  return {
    authKey: await authKey,
    pgpPassphrase: await pgpPassphrase,
  };
}

export async function getKeys(
  password: string,
  saltEncoded: string
): Promise<{ authKey: string; pgpPassphrase: string; err: NullError<any> }> {
  const salt = parse(saltEncoded);
  salt.cost.FillEmpty();

  const { hash, err } = await hashSalt(
    password,
    Buffer.from(salt.salt, "base64"),
    salt.cost
  );
  if (err !== null) {
    return {
      authKey: "",
      pgpPassphrase: "",
      err: err,
    };
  }

  const { authKey, pgpPassphrase } = await deriveKeys(hash);

  return {
    authKey: authKey,
    pgpPassphrase: pgpPassphrase,
    err: null,
  };
}

export default genKeys;

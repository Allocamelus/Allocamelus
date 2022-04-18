import { Buffer } from "buffer";
import { exportKey } from "./aesgcm-tools";
import { hash, hashSalt, parse } from "./argon2id";
import { argon2idCost, argon2idEncoded } from "./argon2id/argon2id";
import { create } from "./recoveryKey";
import { blake2bB64 } from "./blake2b";
import { genKey, decryptKey, encryptKey } from "./pgp";

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

const b2bAuthKey = "B2b User Authentication Key-Info";
const b2bPgpKey = "B2b Pretty Good Privacy Key-Info";

function genHashKeys(
  password: string
): Promise<{ saltEncoded: string; authKey: string; pgpPassphrase: string }> {
  return new Promise(async (resolve) => {
    // Hash password
    const keyEncoded = await hash(
      password,
      new argon2idCost({
        keyLen: 32,
        saltLen: 32,
        memory: 128 * 1024,
        threads: 2,
        time: 3,
      })
    );

    const { authKey, pgpPassphrase } = await deriveKeys(keyEncoded);

    resolve({
      saltEncoded: keyEncoded.encoded,
      authKey: authKey,
      pgpPassphrase: pgpPassphrase,
    });
    return;
  });
}

/**
 * genKeys Generates a new pgp key, recovery key, and password hash
 *
 * @param {string} username
 * @param {string} password un-hashed password
 * @return {Promise<userKey>} msg pretty readable error
 */
export function genKeys(
  username: string,
  password: string
): Promise<{ userKey: userKey; recoveryKey: string }> {
  return new Promise(async (resolve) => {
    const recoveryKey = create();

    const recoveryKeyArray = exportKey((await recoveryKey).key);
    const recoveryHash = blake2bB64(await recoveryKeyArray, 512);

    const keys = await genHashKeys(password);

    const pgpKey = await genKey(username, keys.pgpPassphrase);

    const privateKey = decryptKey(pgpKey.armoredPrivate, keys.pgpPassphrase);
    const recoveryArmored = encryptKey(
      await privateKey,
      username,
      Buffer.from(await recoveryKeyArray).toString("base64")
    );

    resolve({
      userKey: {
        keyAuthHash: keys.authKey,
        keySaltEncoded: keys.saltEncoded,
        publicArmored: pgpKey.armoredPublic,
        privateArmored: pgpKey.armoredPrivate,
        recoveryArmored: await recoveryArmored,
        recoveryHash: await recoveryHash,
      },
      recoveryKey: (await recoveryKey).encoded,
    });
    return;
  });
}

function deriveKeys(
  keyEncoded: argon2idEncoded
): Promise<{ authKey: string; pgpPassphrase: string }> {
  return new Promise(async (resolve) => {
    const key = Buffer.from(keyEncoded.key, "base64");

    const authKey = blake2bB64(key, 512, b2bAuthKey);
    const pgpPassphrase = blake2bB64(key, 512, b2bPgpKey);

    resolve({
      authKey: await authKey,
      pgpPassphrase: await pgpPassphrase,
    });
    return;
  });
}

export function getKeys(
  password: string,
  saltEncoded: string
): Promise<{ authKey: string; pgpPassphrase: string }> {
  return new Promise(async (resolve) => {
    const salt = parse(saltEncoded);
    salt.cost.FillEmpty();

    const hash = await hashSalt(
      password,
      Buffer.from(salt.salt, "base64"),
      salt.cost
    );

    const { authKey, pgpPassphrase } = await deriveKeys(hash);

    resolve({
      authKey: authKey,
      pgpPassphrase: pgpPassphrase,
    });
    return;
  });
}

export default genKeys;

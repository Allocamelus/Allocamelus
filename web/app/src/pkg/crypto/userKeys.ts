import { Buffer } from "buffer";
import { UnixTime } from "../time";
import { hash } from "./argon2id";
import { argon2idCost } from "./argon2id/argon2id";
import { encode, decode } from "./backupKey";
import { blake2bB64 } from "./blake2b";
import { genKey } from "./pgp";

/**
 * userKey
 *
 * @var {number} created unix time
 * @var {string} keyAuthHash key hashed
 * @var {string} keySaltEncoded
 * @var {string} publicKey
 * @var {string} privateKey
 */
export interface userKey {
  created: number;
  keyAuthHash: string;
  keySaltEncoded: string;
  publicKey: string;
  privateKey: string;
}

const b2bAuthKey = "B2b User Authentication Key-Info";
const b2bPgpKey = "B2b Pretty Good Privacy Key-Info";

/**
 * userKeys Generates a new pgp key, backup key, and password hash
 *
 * @param {string} username
 * @param {string} password un-hashed password
 * @return {Promise<userKey>} msg pretty readable error
 */
export function userKeys(
  username: string,
  password: string
): Promise<{ userKey: userKey; backupKey: string }> {
  return new Promise(async (resolve) => {
    let now = UnixTime();

    let derivedKeys = deriveKeys(password);

    let pgpKey = genKey(username, (await derivedKeys).pgpPassphrase);

    let backupKey = encode(
      Buffer.from((await derivedKeys).pgpPassphrase, "base64")
    );

    resolve({
      userKey: {
        created: now,
        keyAuthHash: (await derivedKeys).authKey,
        keySaltEncoded: (await derivedKeys).saltEncoded,
        publicKey: (await pgpKey).armoredPublic,
        privateKey: (await pgpKey).armoredPrivate,
      },
      backupKey: await backupKey,
    });
    return;
  });
}

function deriveKeys(
  password: string
): Promise<{ saltEncoded: string; authKey: string; pgpPassphrase: string }> {
  return new Promise(async (resolve) => {
    // Hash password
    let keyEncoded = await hash(
      password,
      new argon2idCost({
        keyLen: 32,
        saltLen: 32,
        memory: 128 * 1024,
        threads: 2,
        time: 3,
      })
    );

    let key = Buffer.from(keyEncoded.key, "base64");

    let authKey = blake2bB64(key, 512, b2bAuthKey);
    let pgpPassphrase = blake2bB64(key, 256, b2bPgpKey);

    resolve({
      saltEncoded: keyEncoded.encoded,
      authKey: await authKey,
      pgpPassphrase: await pgpPassphrase,
    });
    return;
  });
}

/**
 * decodeBackupKey decodes backup key into pgp passphrase
 * @param backupKey
 * @returns
 */
export function decodeBackupKey(backupKey: string): Promise<string> {
  return new Promise(async (resolve) => {
    let key = await decode(backupKey);

    resolve(Buffer.from(key).toString("base64"));
    return;
  });
}

export default userKeys;

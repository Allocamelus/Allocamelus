import { Buffer } from "buffer";
import { decrypt, encrypt, exportKey } from "./aesgcm-tools";
import { hash } from "./argon2id";
import { argon2idCost } from "./argon2id/argon2id";
import { create, decode } from "./recoveryKey";
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
    let recoveryKey = create();

    let recoveryKeyArray = exportKey((await recoveryKey).key);
    let recoveryHash = blake2bB64(await recoveryKeyArray, 512);

    let keys = await deriveKeys(password);

    let pgpKey = await genKey(username, keys.pgpPassphrase);

    let privateKey = decryptKey(pgpKey.armoredPrivate, keys.pgpPassphrase);
    let recoveryArmored = encryptKey(
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
    let pgpPassphrase = blake2bB64(key, 512, b2bPgpKey);

    resolve({
      saltEncoded: keyEncoded.encoded,
      authKey: await authKey,
      pgpPassphrase: await pgpPassphrase,
    });
    return;
  });
}

export default genKeys;

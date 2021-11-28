import { Buffer } from "buffer";
import { blake2b } from "hash-wasm";
import { UnixTime } from "../time";
import { decrypt, encrypt, exportKey } from "./aesgcm-tools";
import { hash } from "./argon2id";
import { argon2idCost } from "./argon2id/argon2id";
import { create, decode } from "./backupKey";
import { blake2bB64 } from "./blake2b";
import { genKey } from "./pgp";

/**
 * userKey
 *
 * @var {number} created unix time
 * @var {string} keyAuthHash key hashed
 * @var {string} keySaltEncoded
 * @var {string} publicKey
 * @var {string} privateKey encrypted
 * @var {string} passphrase encrypted
 * @var {string} backupKeyHash
 */
export interface userKey {
  created: number;
  keyAuthHash: string;
  keySaltEncoded: string;
  publicKey: string;
  privateKey: string;
  passphrase: string;
  backupKeyHash: string;
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

    let keys = deriveKeys(password);

    let pgpKey = genKey(username, (await keys).pgpPassphrase);

    let backupKey = create();

    let encryptedPassphrase = encryptPGPPassphrase(
      (await backupKey).key,
      (await keys).pgpPassphrase
    );

    let backupKeyHash = blake2bB64(await exportKey((await backupKey).key), 512);

    resolve({
      userKey: {
        created: now,
        keyAuthHash: (await keys).authKey,
        keySaltEncoded: (await keys).saltEncoded,
        publicKey: (await pgpKey).armoredPublic,
        privateKey: (await pgpKey).armoredPrivate,
        passphrase: await encryptedPassphrase,
        backupKeyHash: await backupKeyHash,
      },
      backupKey: (await backupKey).encoded,
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

/**
 * encryptPGPPassphrase encrypts pgp passphrase with backup key
 * @param backupKey
 * @param pgpPassphrase
 * @returns
 */
export function encryptPGPPassphrase(
  backupKey: CryptoKey,
  pgpPassphrase: string
): Promise<string> {
  return new Promise(async (resolve) => {
    resolve(
      Buffer.from(
        await encrypt(backupKey, Buffer.from(pgpPassphrase, "base64"))
      ).toString("base64")
    );
    return;
  });
}

/**
 * decryptBackupKey decrypts pgp passphrase with backup key
 * @param backupKey
 * @param pgpPassphrase
 * @returns
 */
export function decryptPGPPassphrase(
  backupKey: string,
  pgpPassphrase: string
): Promise<string> {
  return new Promise(async (resolve) => {
    let key = await decode(backupKey);

    resolve(
      Buffer.from(
        await decrypt(key, Buffer.from(pgpPassphrase, "base64"))
      ).toString("base64")
    );
    return;
  });
}

export default userKeys;

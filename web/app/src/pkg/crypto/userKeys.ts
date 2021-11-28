import { Buffer } from "buffer";
import { decrypt, encrypt, exportKey } from "./aesgcm-tools";
import { hash } from "./argon2id";
import { argon2idCost } from "./argon2id/argon2id";
import { create, decode } from "./recoveryKey";
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
  keyAuthHash: string;
  keySaltEncoded: string;
  publicArmored: string;
  privateArmored: string;
  passphrase: string;
  recoveryHash: string;
}

const b2bAuthKey = "B2b User Authentication Key-Info";
const b2bPgpKey = "B2b Pretty Good Privacy Key-Info";

/**
 * userKeys Generates a new pgp key, recovery key, and password hash
 *
 * @param {string} username
 * @param {string} password un-hashed password
 * @return {Promise<userKey>} msg pretty readable error
 */
export function userKeys(
  username: string,
  password: string
): Promise<{ userKey: userKey; recoveryKey: string }> {
  return new Promise(async (resolve) => {
    let keys = deriveKeys(password);

    let pgpKey = genKey(username, (await keys).pgpPassphrase);

    let recoveryKey = create();

    let encryptedPassphrase = encryptPGPPassphrase(
      (await recoveryKey).key,
      (await keys).pgpPassphrase
    );

    let recoveryHash = blake2bB64(
      await exportKey((await recoveryKey).key),
      512
    );

    resolve({
      userKey: {
        keyAuthHash: (await keys).authKey,
        keySaltEncoded: (await keys).saltEncoded,
        publicArmored: (await pgpKey).armoredPublic,
        privateArmored: (await pgpKey).armoredPrivate,
        passphrase: await encryptedPassphrase,
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

/**
 * encryptPGPPassphrase encrypts pgp passphrase with recovery key
 * @param recoveryKey
 * @param pgpPassphrase
 * @returns
 */
export function encryptPGPPassphrase(
  recoveryKey: CryptoKey,
  pgpPassphrase: string
): Promise<string> {
  return new Promise(async (resolve) => {
    resolve(
      Buffer.from(
        await encrypt(recoveryKey, Buffer.from(pgpPassphrase, "base64"))
      ).toString("base64")
    );
    return;
  });
}

/**
 * decryptBackupKey decrypts pgp passphrase with recovery key
 * @param recoveryKey
 * @param pgpPassphrase
 * @returns
 */
export function decryptPGPPassphrase(
  recoveryKey: string,
  pgpPassphrase: string
): Promise<string> {
  return new Promise(async (resolve) => {
    let key = await decode(recoveryKey);

    resolve(
      Buffer.from(
        await decrypt(key, Buffer.from(pgpPassphrase, "base64"))
      ).toString("base64")
    );
    return;
  });
}

export default userKeys;

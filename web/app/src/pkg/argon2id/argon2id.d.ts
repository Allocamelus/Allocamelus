export interface argon2idCost {
  // time t= (default: 3)
  time?: number;
  // memory m= (default: 128 * 1024 // 128MB)
  memory?: number;
  // threads p= (default: 2)
  threads?: number;
  // Key length (default: 32 // 256 bits)
  keyLen?: number;
  // Salt length (default: 32 // 256 bits)
  saltLen?: number;
}

export interface argon2idEncoded {
  // Argon2id version v=
  version: number;
  // Cost
  cost: argon2idCost;
  // Encoded with no key $argon2id$v={version}$m={memory},t={time},p={threads}${base64(salt)}
  encoded: string;
  // Salt (Base64 encoded)
  salt: string;
  // Key (Base64 encoded)
  key: string;
  // KeyHash blake2b hash of key (Base64 encoded)
  keyHash: string;
}

export interface argon2id {
  loaded: boolean;
  hash(password: string, cost: string): argon2idEncoded | string;
  hashSalt(
    password: string,
    salt: string,
    cost: string
  ): argon2idEncoded | string;
  parse(encodedHash: string): argon2idEncoded | string;
}

export declare namespace argon2id {
  // loaded set by wasm when all functions have been loaded
  let loaded: boolean;

  let loading: boolean;

  /**
   * Hash password with argon2id
   *
   * @param password base64 encoded
   * @param cost JSON argon2idCost
   */
  function hash(password: string, cost: string): argon2idEncoded | string;
  /**
   * Hash password with argon2id
   * user provides salt
   *
   * @param password base64 encoded
   * @param salt base64 encoded
   * @param cost JSON argon2idCost
   */
  function hashSalt(
    password: string,
    salt: string,
    cost: string
  ): argon2idEncoded | string;

  /**
   * Parse encoded argon2id hash
   *
   * $argon2id$v={version}$m={memory},t={time},p={threads}${base64(salt)}${base64(key)}
   *
   * @param encodedHash Key is optional
   */
  function parse(encodedHash: string): argon2idEncoded | string;
}

declare global {
  interface Window {
    argon2id?: argon2id & typeof argon2id;
  }
}

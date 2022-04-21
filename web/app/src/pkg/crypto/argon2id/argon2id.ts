export interface argon2idCost {
  // time t= (default: 3)
  time: number;
  // memory m= (default: 128 * 1024 // 128MB)
  memory: number;
  // threads p= (default: 2)
  threads: number;
  // Key length (default: 32 // 256 bits)
  keyLen: number;
  // Salt length (default: 32 // 256 bits)
  saltLen: number;
  // Sets defaults
  FillEmpty(): void;
}

export class argon2idCost {
  // time t= (default: 3)
  time: number;
  // memory m= (default: 128 * 1024 // 128MB)
  memory: number;
  // threads p= (default: 2)
  threads: number;
  // Key length (default: 32 // 256 bits)
  keyLen: number;
  // Salt length (default: 32 // 256 bits)
  saltLen: number;

  constructor(source: Partial<argon2idCost> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.time = source["time"] || 0;
    this.memory = source["memory"] || 0;
    this.threads = source["threads"] || 0;
    this.keyLen = source["keyLen"] || 0;
    this.saltLen = source["saltLen"] || 0;
  }

  // Sets defaults
  FillEmpty(): void {
    if (!this.time || this.time <= 0) {
      this.time = 3;
    }
    if (!this.memory || this.memory <= 0) {
      this.memory = 128 * 1024;
    }
    if (!this.threads || this.threads <= 0) {
      this.threads = 2;
    }
    if (!this.keyLen || this.keyLen <= 0) {
      this.keyLen = 32;
    }
    if (!this.saltLen || this.saltLen <= 0) {
      this.saltLen = 32;
    }
  }
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
}
export class argon2idEncoded {
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

  constructor(source: Partial<argon2idEncoded> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.version = source["version"] || 0;
    this.cost = new argon2idCost(source["cost"]);
    this.encoded = source["encoded"] || "";
    this.salt = source["salt"] || "";
    this.key = source["key"] || "";
  }
}

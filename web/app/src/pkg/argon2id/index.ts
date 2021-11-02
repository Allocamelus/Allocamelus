import { argon2idCost, argon2idEncoded } from "./argon2id";
import { Buffer } from "buffer";
import "./wasm/wasm_exec";

export function load(): Promise<boolean> {
  return new Promise((resolve, reject) => {
    if (window.argon2id == undefined) {
      window.argon2id = {
        loaded: false,
        loading: false,
        hash: null,
        hashSalt: null,
        parse: null,
      };
    }
    if (window.argon2id.loading) {
      // wait for argon2id to load
      (() => {
        return new Promise((r) => {
          document.addEventListener("argon2id-load", () => {
            r(true);
          });
        });
      })();
      resolve(true);
      return;
    }

    if (window.argon2id.loaded) {
      resolve(true);
      return;
    }

    let loadEvent = new Event("argon2id-load");
    window.argon2id.loading = true;

    let go = new window.Go();
    let wasmUrl = new URL("./wasm/argon2id.wasm", import.meta.url).href;

    WebAssembly.instantiateStreaming(
      fetch(wasmUrl),
      go.importObject
    ).then(function (obj) {
      go.run(obj.instance);
    });

    document.dispatchEvent(loadEvent);

    resolve(true);
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
  password: string,
  cost: argon2idCost
): Promise<argon2idEncoded> {
  return new Promise((resolve, reject) => {
    load().then(() => {
      let enc = window.argon2id.hash(
        Buffer.from(password).toString("base64"),
        JSON.stringify(cost)
      );
      if (typeof enc === "string") {
        reject(enc);
        return;
      }
      resolve(enc);
    });
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
  password: string,
  salt: string,
  cost: argon2idCost
): Promise<argon2idEncoded> {
  return new Promise((resolve, reject) => {
    load().then(() => {
      let enc = window.argon2id.hashSalt(
        Buffer.from(password).toString("base64"),
        Buffer.from(salt).toString("base64"),
        JSON.stringify(cost)
      );
      if (typeof enc === "string") {
        reject(enc);
        return;
      }
      resolve(enc);
    });
  });
}

/**
 * Parse encoded argon2id hash
 *
 * $argon2id$v={version}$m={memory},t={time},p={threads}${base64(salt)}${base64(key)}
 *
 * @param encodedHash Key is optional
 */
export function parse(encodedHash: string): Promise<argon2idEncoded> {
  return new Promise((resolve, reject) => {
    load().then(() => {
      let enc = window.argon2id.parse(encodedHash);
      if (typeof enc === "string") {
        reject(enc);
        return;
      }
      resolve(enc);
    });
  });
}

import { base58btc } from "multiformats/bases/base58";
import { importKey, newKey } from "./aesgcm-tools";

const splitNum = 6,
  splitRune = "-";

export function create(): Promise<{ key: CryptoKey; encoded: string }> {
  return new Promise(async (resolve) => {
    let key = await newKey(256);
    resolve({ key: key, encoded: await encode(key) });
    return;
  });
}

export function decode(encodedKey: string): Promise<CryptoKey> {
  return new Promise(async (resolve) => {
    encodedKey = encodedKey.replace(splitRune, "");
    resolve(await importKey(base58btc.decode(encodedKey)));
    return;
  });
}

// Adds - every splitNumber
function encode(key: CryptoKey): Promise<string> {
  return new Promise(async (resolve) => {
    let encodedKey: string = "";

    let keyStr = await base58btc.baseEncode(
      new Uint8Array(await window.crypto.subtle.exportKey("raw", key))
    );
    let before = splitNum - 1;
    let last = keyStr.length - 1;

    for (let i = 0; i < keyStr.length; i++) {
      let rune = keyStr[i];
      encodedKey = encodedKey + rune;
      if ((i + 4) % splitNum == before && i != last) {
        encodedKey = encodedKey + splitRune;
      }
    }
    resolve(encodedKey);
    return;
  });
}

import { base58btc } from "multiformats/bases/base58";
import { importKey, exportKey, newKey } from "./aesgcm-tools";

const splitNum = 6,
  splitRune = "-";

export function create(): Promise<{ key: CryptoKey; encoded: string }> {
  return new Promise(async (resolve) => {
    let key = await newKey(256);
    resolve({ key: key, encoded: await encode(key) });
    return;
  });
}

export function decode(splitKey: string): Promise<CryptoKey> {
  return new Promise(async (resolve) => {
    let encodedKey = splitKey.replace(new RegExp(splitRune, "g"), "");
    resolve(await importKey(base58btc.baseDecode(encodedKey)));
    return;
  });
}

// Adds - every splitNumber
function encode(key: CryptoKey): Promise<string> {
  return new Promise(async (resolve) => {
    let encodedKey: string = "";

    let keyStr = base58btc.baseEncode(await exportKey(key));
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

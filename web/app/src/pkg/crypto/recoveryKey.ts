import { base58btc } from "multiformats/bases/base58";
import { importKey, exportKey, newKey } from "./aesgcm-tools";

const splitNum = 6,
  splitRune = "-";

// Adds - every splitNumber
async function encode(key: CryptoKey): Promise<string> {
  let encodedKey = "";

  const keyStr = base58btc.baseEncode(await exportKey(key));
  const before = splitNum - 1;
  const last = keyStr.length - 1;

  for (let i = 0; i < keyStr.length; i++) {
    const rune = keyStr[i];
    encodedKey = encodedKey + rune;
    if ((i + 4) % splitNum == before && i != last) {
      encodedKey = encodedKey + splitRune;
    }
  }

  return encodedKey;
}

export async function create(): Promise<{ key: CryptoKey; encoded: string }> {
  const key = await newKey(256);
  return { key: key, encoded: await encode(key) };
}

export async function decode(splitKey: string): Promise<CryptoKey> {
  const encodedKey = splitKey.replace(new RegExp(splitRune, "g"), "");

  return await importKey(base58btc.baseDecode(encodedKey));
}

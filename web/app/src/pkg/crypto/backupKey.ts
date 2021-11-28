import { base58btc } from "multiformats/bases/base58";

const splitNum = 6,
  splitRune = "-";

export function decode(splitKey: string): Promise<Uint8Array> {
  return new Promise(async (resolve) => {
    let encodedKey = splitKey.replace(new RegExp(splitRune, "g"), "");
    resolve(base58btc.baseDecode(encodedKey));
    return;
  });
}

// Adds - every splitNumber
export function encode(key: Uint8Array): Promise<string> {
  return new Promise(async (resolve) => {
    let encodedKey: string = "";

    let keyStr = base58btc.baseEncode(key);
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

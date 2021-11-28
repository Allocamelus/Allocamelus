import { blake2b } from "hash-wasm";
import { IDataType } from "hash-wasm/dist/lib/util";
import { Buffer } from "buffer";

/**
 * blake2bB64 wrapper for blake2b
 *
 * @param data - Input data (string, Buffer or TypedArray)
 * @param bits
 * Number of output bits, which has to be a number divisible by 8, between 8 and 512. Defaults to 512.
 * @param key - Optional key (string, Buffer or TypedArray). Maximum length is 64 bytes.
 * @returns - Computed hash as a base64 string
 */
export function blake2bB64(
  data: IDataType,
  bits?: number,
  key?: IDataType
): Promise<string> {
  return new Promise(async (resolve) => {
    resolve(
      Buffer.from(await blake2b(data, bits, key), "hex").toString("base64")
    );
    return;
  });
}

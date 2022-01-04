import { importPKCS8, KeyLike } from "jose";
import { PrivateKey } from "openpgp";

export function PrivateKeyToKeyLike(pk: PrivateKey): Promise<KeyLike> {
  return importPKCS8(pk.armor(), "EdDSA");
}

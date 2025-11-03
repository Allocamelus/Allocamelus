import { importPKCS8, CryptoKey } from "jose";
import { PrivateKey } from "openpgp";

export function PrivateKeyToKeyLike(pk: PrivateKey): Promise<CryptoKey> {
  return importPKCS8(pk.armor(), "EdDSA");
}

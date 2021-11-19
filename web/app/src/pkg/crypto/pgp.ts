import { generateKey } from "openpgp";

export interface pgpKey {
  armoredPrivate?: string;
  armoredPublic?: string;
  armoredRevocation?: string;
}

export function genkey(name: string, passphrase: string): Promise<pgpKey> {
  return new Promise(async (resolve) => {
    generateKey({
      type: "ecc", // Type of the key, defaults to ECC
      curve: "curve25519", // ECC curve name, defaults to curve25519
      userIDs: [{ name: name }], // you can pass multiple user IDs
      passphrase: passphrase, // protects the private key
      format: "armored", // output key format, defaults to 'armored' (other options: 'binary' or 'object')
    }).then((skp) => {
      let keys: pgpKey = {
        armoredPrivate: skp.privateKey,
        armoredPublic: skp.publicKey,
        armoredRevocation: skp.revocationCertificate,
      };

      resolve(keys);
      return;
    });
  });
}

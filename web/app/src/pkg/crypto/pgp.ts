import {
  generateKey,
  decryptKey as decryptPrivateKey,
  readPrivateKey,
  reformatKey,
  PrivateKey,
} from "openpgp";

export interface pgpKey {
  armoredPrivate: string;
  armoredPublic: string;
  armoredRevocation: string;
}

export function genKey(name: string, passphrase: string): Promise<pgpKey> {
  return new Promise((resolve) => {
    generateKey({
      type: "ecc", // Type of the key, defaults to ECC
      curve: "curve25519", // ECC curve name, defaults to curve25519
      userIDs: [{ name: name }], // you can pass multiple user IDs
      passphrase: passphrase, // protects the private key
      format: "armored", // output key format, defaults to 'armored' (other options: 'binary' or 'object')
    }).then((skp) => {
      resolve({
        armoredPrivate: skp.privateKey,
        armoredPublic: skp.publicKey,
        armoredRevocation: skp.revocationCertificate,
      });
      return;
    });
  });
}

export async function decryptKey(
  key: string,
  passphrase: string
): Promise<PrivateKey> {
  return await decryptPrivateKey({
    privateKey: await readPrivateKey({ armoredKey: key }),
    passphrase,
  });
}

export async function encryptKey(
  key: PrivateKey,
  name: string,
  passphrase: string
): Promise<string> {
  const armoredKey = await reformatKey({
    privateKey: key,
    userIDs: [{ name: name }],
    passphrase,
    format: "armored",
  });

  return armoredKey.privateKey;
}

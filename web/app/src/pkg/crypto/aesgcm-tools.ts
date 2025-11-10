export async function newKey(length?: number): Promise<CryptoKey> {
  if (
    length === undefined ||
    (length != 128 && length != 192 && length != 256)
  ) {
    length = 256;
  }
  return await window.crypto.subtle.generateKey(
    {
      name: "AES-GCM",
      length: length,
    },
    true,
    ["encrypt", "decrypt"]
  );
}

export async function exportKey(key: CryptoKey): Promise<Uint8Array> {
  return new Uint8Array(await window.crypto.subtle.exportKey("raw", key));
}

export async function importKey(key: ArrayBuffer): Promise<CryptoKey> {
  return await window.crypto.subtle.importKey("raw", key, "AES-GCM", true, [
    "encrypt",
    "decrypt",
  ]);
}

export async function encrypt(
  key: CryptoKey,
  message: Uint8Array
): Promise<Uint8Array> {
  const iv = window.crypto.getRandomValues(new Uint8Array(12));

  const cipherText = new Uint8Array(
    await window.crypto.subtle.encrypt(
      {
        name: "AES-GCM",
        iv: iv,
      },
      key,
      message
    )
  );

  const cipher = new Uint8Array(iv.length + cipherText.length);
  cipher.set(iv, 0);
  cipher.set(cipherText, iv.length);

  return cipher;
}

export async function decrypt(
  key: CryptoKey,
  cipherText: Uint8Array
): Promise<Uint8Array> {
  const iv = cipherText.slice(0, 12);

  const plainText = new Uint8Array(
    await window.crypto.subtle.decrypt(
      {
        name: "AES-GCM",
        iv: iv,
      },
      key,
      cipherText.slice(12)
    )
  );

  return plainText;
}

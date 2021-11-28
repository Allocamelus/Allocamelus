export function newKey(length?: number): Promise<CryptoKey> {
  return new Promise(async (resolve) => {
    if (
      length === undefined ||
      (length != 128 && length != 192 && length != 256)
    ) {
      length = 256;
    }
    resolve(
      await window.crypto.subtle.generateKey(
        {
          name: "AES-GCM",
          length: length,
        },
        true,
        ["encrypt", "decrypt"]
      )
    );
    return;
  });
}

export function exportKey(key: CryptoKey): Promise<Uint8Array> {
  return new Promise(async (resolve) => {
    resolve(new Uint8Array(await window.crypto.subtle.exportKey("raw", key)));
    return;
  });
}

export function importKey(key: ArrayBuffer): Promise<CryptoKey> {
  return new Promise(async (resolve) => {
    resolve(
      await window.crypto.subtle.importKey("raw", key, "AES-GCM", true, [
        "encrypt",
        "decrypt",
      ])
    );
    return;
  });
}

export function encrypt(
  key: CryptoKey,
  message: Uint8Array
): Promise<Uint8Array> {
  return new Promise(async (resolve) => {
    let iv = window.crypto.getRandomValues(new Uint8Array(12));

    let cipherText = new Uint8Array(
      await window.crypto.subtle.encrypt(
        {
          name: "AES-GCM",
          iv: iv,
        },
        key,
        message
      )
    );

    let cipher = new Uint8Array(iv.length + cipherText.length);
    cipher.set(iv, 0);
    cipher.set(cipherText, iv.length);

    resolve(cipher);
    return;
  });
}

export function decrypt(
  key: CryptoKey,
  cipherText: Uint8Array
): Promise<Uint8Array> {
  return new Promise(async (resolve) => {
    let iv = cipherText.slice(0, 12);

    let plainText = new Uint8Array(
      await window.crypto.subtle.decrypt(
        {
          name: "AES-GCM",
          iv: iv,
        },
        key,
        cipherText.slice(12)
      )
    );

    resolve(plainText);
    return;
  });
}

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

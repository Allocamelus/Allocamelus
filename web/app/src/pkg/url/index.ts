export function FullURL(url: string, base?: string | URL): string {
  try {
    // Try building url with base
    // base must be a valid see new URL(url [, base])
    // @ https://developer.mozilla.org/en-US/docs/Web/API/URL/URL
    url = new URL(url, base).href;
  } catch (_) {
    return url;
  }
  return url;
}

export function local(url: string) {
  try {
    const testUrl = new URL(url, window.location.origin);
    if (testUrl.host === window.location.host) {
      return true;
    }
  } catch (_) {
    return true;
  }
  return false;
}

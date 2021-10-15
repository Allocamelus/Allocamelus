export function FullURL(url: string, base?: string | URL) {
  try {
    // Try building url with base
    // base must be a valid see new URL(url [, base])
    // @ https://developer.mozilla.org/en-US/docs/Web/API/URL/URL
    url = new URL(url, base).href;
  } finally {
    return url;
  }
}

export default FullURL
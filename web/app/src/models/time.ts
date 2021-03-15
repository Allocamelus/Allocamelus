export function UnixTime(append?: number) {
  if (append != undefined) {
    return Math.floor(Date.now() / 1000) + append
  }
  return Math.floor(Date.now() / 1000)
}

// MinToSec minutes in seconds
export function MinToSec(minutes: number) {
  return minutes * 60
}
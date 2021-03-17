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
// HoursToSec hours in seconds
export function HoursToSec(hours: number) {
  return MinToSec(hours * 60)
}
// DaysToSec hours in seconds
export function DaysToSec(days: number) {
  return HoursToSec(days * 24)
}
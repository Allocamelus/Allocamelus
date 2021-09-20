export function UnixTime(append?: number): number {
  if (append != undefined) {
    return Math.floor(Date.now() / 1000) + append;
  }
  return Math.floor(Date.now() / 1000);
}

export function UnixToDate(time: number): Date {
  return new Date(time * 1000);
}

// SecToMs seconds to millisecond
export function SecToMs(seconds: number): number {
  return seconds * 1000;
}

// MinToSec minutes to seconds
export function MinToSec(minutes: number): number {
  return minutes * 60;
}
// HoursToSec hours to seconds
export function HoursToSec(hours: number): number {
  return MinToSec(hours * 60);
}
// DaysToSec hours to seconds
export function DaysToSec(days: number): number {
  return HoursToSec(days * 24);
}

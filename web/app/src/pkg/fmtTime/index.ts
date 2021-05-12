import { UnixTime } from "../time"

export const times = {
  Second: 1,
  Minute: 60,
  Hour: 3600,
  Day: 86400,
  Month: 2628288,
  Year: 31557600,
}

export default (time: number) => {
  var since = UnixTime(-time)
  if (since >= times.Year) {
    return fmtTime(since, times.Year, " Year")
  } else if (since >= times.Month) {
    return fmtTime(since, times.Month, " Month")
  } else if (since >= times.Day) {
    return fmtTime(since, times.Day, " Day")
  } else if (since >= times.Hour) {
    return fmtTime(since, times.Hour, " Hour")
  } else if (since >= times.Minute) {
    return fmtTime(since, times.Minute, " Minute")
  } else if (since >= times.Second) {
    return fmtTime(since, times.Second, " Second")
  } else {
    return "Just Now"
  }
}

export function fmtTime(t: number, tDuration: number, postFix: string, short = false) {
  // time from now divided by the provide duration
  // Example: t=600 so tD...=60 (1 Minute) so t/tD... = 10 (minutes)
  var sinceIn = Math.round(t / tDuration)

  // If >= 5 seconds ago
  if (tDuration == 1 && 5 >= sinceIn) {
    return "Just Now"
  }

  var fmtTime = `${sinceIn}${postFix}`
  if (!short) {
    if (sinceIn > 1) {
      fmtTime += "s"
    }
    fmtTime += " Ago"
  }
  return fmtTime
}

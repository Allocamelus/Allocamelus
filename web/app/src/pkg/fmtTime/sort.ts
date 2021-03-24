import { UnixTime, UnixToDate } from "../time"
import { times, fmtTime } from "./index";

const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]

// as sort as possible
export default (time: number) => {
  var since = UnixTime(-time)
  if (since >= times.Year) {
    return MDY(time)
  } else if (since >= times.Month) {
    return MD(time)
  } else if (since >= times.Day) {
    return fmtTime(since, times.Day, "d", true)
  } else if (since >= times.Hour) {
    return fmtTime(since, times.Hour, "h", true)
  } else if (since >= times.Minute) {
    return fmtTime(since, times.Minute, "m", true)
  } else if (since >= times.Second) {
    return fmtTime(since, times.Second, "s", true)
  } else {
    return "Just Now"
  }
}

function convertToDate(time: number | Date) {
  if (typeof time == 'number') {
    time = UnixToDate(time)
  }
  return time
}

function paddedTime(date: Date) {
   var min = String(date.getUTCMinutes())
   if (min.length == 1) {
     min = `0${min}`
   }
   return `${date.getUTCHours()}:${min}`
}

export function MDY_HM(time: number | Date) {
  time = convertToDate(time)
  return `${MDY(time)} UTC ${paddedTime(time)}`
}

export function MDY(time: number | Date) {
  time = convertToDate(time)
  return `${MD(time)}, ${time.getUTCFullYear()}`
}

export function MD(time: number | Date) {
  time = convertToDate(time)
  return `${months[time.getUTCMonth()]} ${time.getUTCDate()}`
}
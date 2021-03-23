import { UnixTime } from "../time"
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

export function MDY(time: number) {
	var t = new Date(time*1000)
	return MD(time) + ", " + t.getFullYear()
}

export function MD(time: number) {
	var t = new Date(time*1000)
	return months[t.getMonth()] + " " + t.getDate()
}
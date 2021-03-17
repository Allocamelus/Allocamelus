export const ErrRequired = "required",
  ErrMinLength = "min-length",
  ErrMaxLength = "max-length",
  ErrRegex = "regex"

/**
 * ErrMsg Returns a pretty readable error message 
 *
 * @param {string} err one of the above const
 * @param {string} extra extra text to add to error
 * placement depends on error
 * @param {string} type for length checkers
 * @return {string} msg pretty readable error
 */
export function ErrMsg(err, extra = "", type = "characters") {
  var msg = errMsg(err)
  if (err == ErrRegex && (extra == undefined || extra.length <= 0)) {
    msg = "Failed Regex Check"
  } else {
    msg += extra
  }
  if (err == ErrMinLength || err == ErrMaxLength) {
    msg += " " + type
  }
  return msg
}

function errMsg(err) {
  switch (err) {
    case ErrRequired:
      return "Required"
    case ErrMinLength:
      return "Minimum of "
    case ErrMaxLength:
      return "Maximum of "
    default:
      return ""
  }
}
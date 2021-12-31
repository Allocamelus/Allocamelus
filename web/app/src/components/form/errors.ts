export const ErrRequired = "required",
  ErrMinLength = "min-length",
  ErrMaxLength = "max-length",
  ErrRegex = "regex";

function errMsg(err: string): string {
  switch (err) {
    case ErrRequired:
      return "Required";
    case ErrMinLength:
      return "Minimum";
    case ErrMaxLength:
      return "Maximum";
    default:
      return "";
  }
}

/**
 * ErrMsg Returns a pretty readable error message
 *
 * @param {string} err one of the above const
 * @param {string} extra extra text to add to error
 * placement depends on error
 * @param {string} type for length checkers
 * @return {string} msg pretty readable error
 */
export function ErrMsg(
  err: string,
  extra = "",
  type = "character count"
): string {
  var msg = errMsg(err);

  if (err === ErrMinLength || err === ErrMaxLength) {
    if (type.length > 0) {
      msg += ` ${type}`;
    }
    msg += " of ";
  }

  if (err === ErrRegex && (typeof extra === "undefined" || extra.length <= 0)) {
    msg = "Failed Regex Check";
  } else {
    msg += extra;
  }
  return msg;
}

export const InvalidCharacters = "Invalid Character(s)";
export const SomethingWentWrong = "Something went wrong, Try again later";

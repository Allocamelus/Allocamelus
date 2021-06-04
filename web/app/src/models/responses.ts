import { Html404Func, Html403Func, HtmlSomethingWentWrong } from '../components/htmlErrors'
const invalidLength = "invalid-length",
  invalidChars = "invalid-characters",
  taken = "taken"

export const Shared = {
  A10: "allocamelus",
  InvalidAuthToken: "invalid-auth-token",
  InvalidWith: "invalid-with-value",
  InvalidCaptcha: "invalid-captcha",
  InsufficientPerms: "insufficient-permissions",
  NotFound: "not-found",
  Unauthorized403: "unauthorized-403",
  File: {
    ContentType: "invalid-content-type",
    Size: "invalid-file-size"
  },
}

export const Account = {
  Auth: {
    InvalidUsernamePassword: "invalid-username-password",
    UnverifiedEmail: "unverified-email",
    Authenticated: "already-authenticated",
    // Persistent Auth Failed
    AuthToken: "persistent-auth-failed",
  }
}
export const Post = {
  Create: {
    Unauthorized: "unauthorized"
  },
  Validate: {
    Content: {
      Length: invalidLength + "-min0-max65500"
    },
  }
}
export const User = {
  Create: {
    InvalidCreateToken: "invalid-create-token",
    LoggedIn: "logged-in",
  },
  EmailToken: {
    Validate: {
      Invalid: "invalid-token",
      Expired: "expired-token"
    }
  },
  Validate: {
    UserName: {
      Length: invalidLength + "-min5-max64",
      Taken: taken,
    },
    Name: {
      Length: invalidLength + "-min1-max128",
    },
    Email: {
      Invalid: "invalid-email"
    },
    Bio: {
      Length: invalidLength + "-min0-max255",
    },
    Password: {
      Length: invalidLength + "-min8-max1024",
      Strength: "weak-password",
    },
    Invalid: invalidChars,
  }
}

export const Responses = {
  Shared,
  Account,
  Post,
  User
  // TODO: Add rest of the api codes
}

export default Responses

// RespToError response to a more human readable error
// TODO: Add All error text
export function RespToError(resp: string): string {
  switch (resp) {
    case Responses.User.Validate.Bio.Length:
      return "Invalid Length 0-255 Characters"
    case Responses.User.Validate.Email.Invalid:
      return "Invalid Email"
    case Responses.User.Validate.Invalid:
      return "Invalid Characters"
    case Responses.User.Validate.Name.Length:
      return "Invalid Length 1-128 Characters"
    case Responses.User.Validate.Password.Length:
      return "Invalid Length 8-1024 Characters"
    case Responses.User.Validate.Password.Strength:
      return "Weak Password"
    case Responses.User.Validate.UserName.Length:
      return "Invalid Length 5-64 Characters"
    case Responses.User.Validate.UserName.Taken:
      return "Username Taken"
    case Responses.Shared.File.ContentType:
      return "Unsupported Image Type"
    case Responses.Shared.File.Size:
      return "Image size to large"
  }
  return ""
}

// RespToHtml response to a html readable error
// TODO: Add All error text
export function RespToHtml(resp: string): string {
  switch (resp) {
    case Responses.Shared.NotFound:
      return Html404Func()
    case Responses.Shared.Unauthorized403:
      return Html403Func()
  }
  return HtmlSomethingWentWrong
}
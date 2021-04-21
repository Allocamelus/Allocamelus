const invalidLength = "invalid-length",
  invalidChars = "invalid-characters",
  taken = "taken"

export const Responses = {
  Shared: {
    A10: "allocamelus",
    InvalidAuthToken: "invalid-auth-token",
    InvalidWith: "invalid-with-value",
    InvalidCaptcha: "invalid-captcha",
    NotFound: "not-found",
    File: {
      ContentType: "invalid-content-type",
      Size: "invalid-file-size"
    },
  },
  Account: {
    Auth: {
      InvalidUsernamePassword: "invalid-username-password",
      UnverifiedEmail: "unverified-email",
      Authenticated: "already-authenticated",
      // Persistent Auth Failed
      AuthToken: "persistent-auth-failed",
    }
  },
  Post: {
    Create: {
      Unauthorized: "unauthorized",
      InsufficientPerms: "insufficient-permissions"
    },
    Validate: {
      Content: {
        Length: invalidLength + "-min0-max65500"
      },
    }
  },
  User: {
    Create: {
      InvalidCreateToken: "invalid-create-token",
      LoggedIn: "logged-in",
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
  // TODO: Add rest of the api codes
}

export default Responses

// RespToError response to a more human readable error
// TODO: Add All error text
export function RespToError(resp: string) {
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
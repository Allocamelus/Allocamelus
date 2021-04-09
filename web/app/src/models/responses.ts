const invalidLength = "invalid-length",
  invalidChars = "invalid-characters",
  taken = "taken"

export default {
  Shared: {
    A10: "allocamelus",
    InvalidAuthToken: "invalid-auth-token",
    InvalidWith: "invalid-with-value",
    InvalidCaptcha: "invalid-captcha",
    NotFound: "not-found",
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
    }
  },
  User: {
    Create: {
      InvalidCreateToken: "invalid-create-token",
      LoggedIn: "logged-in",
    },
    Avatar: {
      ContentType: "invalid-content-type",
      FileSize: "invalid-file-size"
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
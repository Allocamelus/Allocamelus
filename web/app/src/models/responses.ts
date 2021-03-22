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
      InvalidCreateToken: "invalid-create-token"
    }
  }
  // TODO: Add rest of the api codes
}
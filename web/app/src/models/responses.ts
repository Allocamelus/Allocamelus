module.exports = {
  API_Shared: {
    A10: "allocamelus",
    InvalidAuthToken: "invalid-auth-token",
    InvalidWith: "invalid-with-value",
    InvalidCaptcha: "invalid-captcha",
  },
  API_Account_Auth: {
    InvalidUsernamePassword: "invalid-username-password",
    UnverifiedEmail: "unverified-email",
    Authenticated: "already-authenticated",
    // Persistent Auth Failed
    AuthToken: "persistent-auth-failed",
  },
  API_Post_Create: {
    Unauthorized: "unauthorized",
    InsufficientPerms: "insufficient-permissions"
  },
  API_User_Create: {
    InvalidCreateToken: "invalid-create-token"
  }
  // TODO: Add rest of the api codes
}
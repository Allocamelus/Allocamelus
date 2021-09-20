import {
  Html404Func,
  Html403Func,
  HtmlSomethingWentWrong,
} from "../components/htmlErrors";

function invalidLenErrMaker(min: number, max: number) {
  return `invalid-length-min${min}-max${max}`;
}

function invalidLenTxtMaker(min: number, max: number) {
  return `Invalid Length ${min}-${max} Characters`;
}

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
    Size: "invalid-file-size",
  },
  InvalidChars: "invalid-characters",
};

export const Account = {
  Auth: {
    InvalidUsernamePassword: "invalid-username-password",
    UnverifiedEmail: "unverified-email",
    Authenticated: "already-authenticated",
    // Persistent Auth Failed
    AuthToken: "persistent-auth-failed",
  },
};
export const Post = {
  Create: {
    Unauthorized: "unauthorized",
  },
  Validate: {
    Content: {
      Length: invalidLenErrMaker(0, 65500),
    },
  },
};

export const Comment = {
  Validate: {
    Content: {
      Length: invalidLenErrMaker(2, 4096),
    },
    Invalid: Shared.InvalidChars,
  },
};

export const User = {
  Create: {
    InvalidCreateToken: "invalid-create-token",
    LoggedIn: "logged-in",
  },
  EmailToken: {
    Validate: {
      Invalid: "invalid-token",
      Expired: "expired-token",
    },
  },
  Validate: {
    UserName: {
      Length: invalidLenErrMaker(5, 64),
      Taken: "taken",
    },
    Name: {
      Length: invalidLenErrMaker(1, 128),
    },
    Email: {
      Invalid: "invalid-email",
    },
    Bio: {
      Length: invalidLenErrMaker(0, 255),
    },
    Password: {
      Length: invalidLenErrMaker(8, 1024),
      Strength: "weak-password",
    },
    Invalid: Shared.InvalidChars,
  },
};

export const Responses = {
  Shared,
  Account,
  Post,
  Comment,
  User,
  // TODO: Add rest of the api codes
};

export default Responses;

// RespToError response to a more human readable error
// TODO: Add All error text
export function RespToError(resp: string): string {
  switch (resp) {
    // Shared errors
    case Shared.NotFound:
      return "Error: 404 Not Found";
    case Shared.Unauthorized403:
      return "Error: 403 Forbidden";

    // User errors
    case User.Validate.Bio.Length:
      return invalidLenTxtMaker(0, 255);
    case User.Validate.Email.Invalid:
      return "Invalid Email";
    case User.Validate.Name.Length:
      return invalidLenTxtMaker(1, 128);
    case User.Validate.Password.Length:
      return invalidLenTxtMaker(8, 1024);
    case User.Validate.Password.Strength:
      return "Weak Password";
    case User.Validate.UserName.Length:
      return invalidLenTxtMaker(5, 64);
    case User.Validate.UserName.Taken:
      return "Username Taken";

    // Comment errors
    case Comment.Validate.Content.Length:
      return invalidLenTxtMaker(2, 4096);

    // Shared errors
    case Shared.File.ContentType:
      return "Unsupported Image Type";
    case Shared.File.Size:
      return "Image size to large";
    case Shared.InvalidChars:
      return "Invalid Characters";
  }
  return "";
}

export function notNull(err: string): boolean {
  return err !== undefined && err !== null && err !== "";
}

// RespToHtml response to a html readable error
// TODO: Add All error text
export function RespToHtml(resp: string): string {
  switch (resp) {
    case Shared.NotFound:
      return Html404Func();
    case Shared.Unauthorized403:
      return Html403Func();
  }
  return HtmlSomethingWentWrong;
}

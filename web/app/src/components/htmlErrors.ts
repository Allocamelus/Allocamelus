import sanitizeHtml from '../pkg/sanitize'

export function htmlErrBuilder(first: string, second = ""): string {
  let err = `<strong>`;
  err += first;
  err += `</strong>`;
  if (second.length > 0) {
    err += `<br>` + second;
  }
  return err;
}

export const HtmlSomethingWentWrong = htmlErrBuilder(
  `Something went wrong`,
  `Try again later`
)

function getPath(path = ""): string {
  if (path.length == 0) {
    path = window.location.pathname
  }
  return sanitizeHtml(path) + " "
}

export function Html403Func(path = ""): string {
  path = getPath(path)
  return htmlErrBuilder(
    `Error: 403 Forbidden`,
    `This page is Private ${path}`
  )
}

export function Html404Func(path = ""): string {
  path = getPath(path)
  return htmlErrBuilder(
    `Error: 404`,
    path + `Not Found`
  )
}


export const Html403 = Html403Func()
export const Html404 = Html404Func()

export const HtmlLoadingCaptcha = htmlErrBuilder("Loading captcha...");
import sanitizeHtml from '../pkg/sanitize'

export function htmlErrBuilder(first: string, second: string = "") {
  var err = `<strong>`;
  err += first;
  err += `</strong>`;
  if (second.length > 0) {
    err += `<br>` + second;
  }
  return err;
}

export const HtmlSomthingWentWrong = htmlErrBuilder(
  `Something went wrong`,
  `Try again later`
)


export const Html404Func = (path: string = "") => {
  if (path.length > 0) {
    path = sanitizeHtml(path) + " "
  }
  return htmlErrBuilder(
    `Error: 404`,
    path + `Not Found`
  )
}

export const Html404 = Html404Func()

export const HtmlLoadingCaptcha = htmlErrBuilder("Loading captcha...");
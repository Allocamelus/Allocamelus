import v1 from "../../v1";

export async function status() {
  return v1.get("account/auth/status").then(r => {
    if (r.data.loggedIn === true) {
      return true
    }
    return false
  })
}
import v1 from "../../v1";

export async function keepAlive() {
  return v1.post("account/auth/keep-alive")
}
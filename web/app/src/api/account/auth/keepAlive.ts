import v1 from "../../v1";

export async function keepAlive() {
  v1.post("account/auth/keep-alive")
}
import v1 from "../../v1";

export function keepAlive() {
  v1.post("account/auth/keep-alive");
}
